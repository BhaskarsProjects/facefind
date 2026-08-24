package service

import (
	"context"
	"fmt"
	"mime/multipart"
	"strings"

	"github.com/bhaskarvb28/facefind/internal/lib/ai"
	"github.com/bhaskarvb28/facefind/internal/lib/aws"
	"github.com/bhaskarvb28/facefind/internal/model/photo"
	"github.com/bhaskarvb28/facefind/internal/repository"
	"github.com/bhaskarvb28/facefind/internal/server"
	"github.com/google/uuid"
)


type PhotoService struct {
    server     *server.Server
    repository *repository.PhotoRepository
    storage    *aws.S3Client
    mlClient   *ai.MLClient
}

func NewPhotoService(
    server *server.Server,
    repository *repository.PhotoRepository,
    storage *aws.S3Client,
    mlClient *ai.MLClient,
) (*PhotoService, error) {
    return &PhotoService{
        server:     server,
        repository: repository,
        storage:    storage,
        mlClient:   mlClient,
    }, nil
}

func (s *PhotoService) Upload(
	ctx context.Context,
	ownerSub string,
	folderID string,
	files []*multipart.FileHeader,
) error {

	// 1. Verify that the folder belongs to the authenticated user.
	if err := s.repository.VerifyFolderOwner(
		ctx,
		ownerSub,
		folderID,
	); err != nil {
		return err
	}

	for _, file := range files {
		if err := s.uploadPhoto(
			ctx,
			folderID,
			file,
		); err != nil {
			return err
		}
	}

	return nil
}

func (s *PhotoService) uploadPhoto(
	ctx context.Context,
	folderID string,
	file *multipart.FileHeader,
) error {

	// 2. Validate content type.
	contentType := strings.ToLower(
		file.Header.Get("Content-Type"),
	)

	switch contentType {
	case "image/jpeg", "image/png", "image/webp":
		// Valid.
	default:
		return fmt.Errorf(
			"unsupported image type: %s",
			contentType,
		)
	}

	// 3. Validate file size.
	const maxFileSize = 25 * 1024 * 1024

	if file.Size > maxFileSize {
		return fmt.Errorf(
			"file %q exceeds 25 MB",
			file.Filename,
		)
	}

	// 4. Open uploaded file.
	src, err := file.Open()
	if err != nil {
		return fmt.Errorf(
			"open %q: %w",
			file.Filename,
			err,
		)
	}
	defer src.Close()

	// 5. Generate photo ID and storage key.
	photoID := uuid.NewString()

	storageKey := fmt.Sprintf(
		"folders/%s/photos/%s",
		folderID,
		photoID,
	)

	// 6. Upload original to R2.
	err = s.storage.Upload(
		ctx,
		s.server.Config.R2.Bucket,
		storageKey,
		src,
		contentType,
		file.Size,
	)
	if err != nil {
		return fmt.Errorf(
			"upload %q to storage: %w",
			file.Filename,
			err,
		)
	}

	// 7. Create database record.
	err = s.repository.Create(
		ctx,
		folderID,
		photoID,
		file.Filename,
		contentType,
		file.Size,
		storageKey,
	)
	if err != nil {
		// R2 upload succeeded but DB insert failed.
		// Remove the object to prevent an orphaned file.
		if deleteErr := s.storage.Delete(
			ctx,
			s.server.Config.R2.Bucket,
			storageKey,
		); deleteErr != nil {
			return fmt.Errorf(
				"create photo record: %w; cleanup storage object: %v",
				err,
				deleteErr,
			)
		}

		return fmt.Errorf(
			"create photo record: %w",
			err,
		)
	}

	_, err = s.mlClient.Process(
		ctx,
		photoID,
		storageKey,
	)
	if err != nil {
		return fmt.Errorf(
			"process photo with AI worker: %w",
			err,
		)
	}


	return nil
}

func (s *PhotoService) GetFolderPhotos(
	ctx context.Context,
	folderID uuid.UUID,
	userSub string,
) ([]*photo.PhotoResponse, error) {

	photos, err := s.repository.GetByFolderID(
		ctx,
		folderID,
		userSub,
	)
	if err != nil {
		return nil, err
	}

	responses := make([]*photo.PhotoResponse, 0, len(photos))

	for _, p := range photos {
		url, err := s.storage.CreatePresignedURL(
			ctx,
			s.server.Config.R2.Bucket,
			p.StorageKey,
		)
		if err != nil {
			return nil, fmt.Errorf(
				"create presigned url for photo %s: %w",
				p.ID,
				err,
			)
		}

		responses = append(responses, &photo.PhotoResponse{
			ID:               p.ID,
			FolderID:         p.FolderID,
			OriginalFilename: p.OriginalFilename,
			MimeType:         p.MimeType,
			SizeBytes:        p.SizeBytes,
			CreatedAt:        p.CreatedAt,
			UpdatedAt:        p.UpdatedAt,
			URL:              url,
		})
	}

	return responses, nil
}

func (s *PhotoService) Delete(
	ctx context.Context,
	folderID uuid.UUID,
	photoID uuid.UUID,
	userSub string,
) error {
	p, err := s.repository.GetByID(
		ctx,
		folderID,
		photoID,
		userSub,
	)
	if err != nil {
		return fmt.Errorf("get photo: %w", err)
	}

	err = s.storage.Delete(
		ctx,
		s.server.Config.R2.Bucket,
		p.StorageKey,
	)
	if err != nil {
		return fmt.Errorf(
			"delete photo from storage: %w",
			err,
		)
	}

	err = s.repository.Delete(
		ctx,
		folderID,
		photoID,
		userSub,
	)
	if err != nil {
		return fmt.Errorf(
			"delete photo record: %w",
			err,
		)
	}

	return nil
}