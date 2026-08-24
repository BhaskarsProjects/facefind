package service

import (
	"context"
	"fmt"

	"github.com/bhaskarvb28/facefind/internal/lib/ai"
	"github.com/bhaskarvb28/facefind/internal/lib/aws"
	"github.com/bhaskarvb28/facefind/internal/model/face"
	"github.com/bhaskarvb28/facefind/internal/repository"
	"github.com/bhaskarvb28/facefind/internal/server"
	"github.com/google/uuid"
)

type FaceService struct {
	server         *server.Server
	repository     *repository.FaceRepository
	photoRepository *repository.PhotoRepository
	mlClient       *ai.MLClient
	storage        *aws.S3Client
}

func NewFaceService(
	server *server.Server,
	repository *repository.FaceRepository,
	photoRepository *repository.PhotoRepository,
	mlClient *ai.MLClient,
	storage *aws.S3Client,
) *FaceService {
	return &FaceService{
		server:          server,
		repository:      repository,
		photoRepository: photoRepository,
		mlClient:        mlClient,
		storage: storage,
	}
}

func (s *FaceService) Create(
	ctx context.Context,
	photoID uuid.UUID,
	faces []face.CreateFaceInput,
) error {
	exists, err := s.photoRepository.Exists(
		ctx,
		photoID,
	)
	if err != nil {
		return fmt.Errorf("check photo: %w", err)
	}

	if !exists {
		return fmt.Errorf("photo not found")
	}

	for _, detectedFace := range faces {
		if len(detectedFace.Embedding) != 512 {
			return fmt.Errorf(
				"embedding must contain 512 values, got %d",
				len(detectedFace.Embedding),
			)
		}

		if len(detectedFace.Landmarks) != 5 {
			return fmt.Errorf(
				"expected 5 landmarks, got %d",
				len(detectedFace.Landmarks),
			)
		}

		for _, landmark := range detectedFace.Landmarks {
			if len(landmark) != 2 {
				return fmt.Errorf(
					"each landmark must contain 2 values",
				)
			}
		}

		if err := s.repository.Create(
			ctx,
			photoID,
			detectedFace.X1,
			detectedFace.Y1,
			detectedFace.X2,
			detectedFace.Y2,
			detectedFace.Confidence,
			detectedFace.Landmarks,
			detectedFace.Embedding,
		); err != nil {
			return err
		}
	}

	return nil
}

func (s *FaceService) Search(
    ctx context.Context,
    folderID uuid.UUID,
    image []byte,
    filename string,
    limit int,
) ([]face.FaceSearchResult, error) {

    if len(image) == 0 {
        return nil, fmt.Errorf("image is required")
    }

    if limit <= 0 {
        limit = 50
    }

    embedding, err := s.mlClient.Embed(
        ctx,
        image,
        filename,
    )
    if err != nil {
        return nil, fmt.Errorf(
            "generate query embedding: %w",
            err,
        )
    }

    if len(embedding.Embedding) != 512 {
        return nil, fmt.Errorf(
            "embedding must contain 512 values, got %d",
            len(embedding.Embedding),
        )
    }

    const maxDistance float32 = 0.7

    matches, err := s.repository.SearchSimilar(
        ctx,
        folderID,
        embedding.Embedding,
        limit,
        maxDistance,
    )
    if err != nil {
        return nil, fmt.Errorf(
            "search faces: %w",
            err,
        )
    }

    results := make(
        []face.FaceSearchResult,
        0,
        len(matches),
    )

    for _, match := range matches {
		url, err := s.storage.CreatePresignedURL(
			ctx,
			s.server.Config.R2.Bucket,
			match.StorageKey,
		)
		if err != nil {
			return nil, fmt.Errorf(
				"create presigned url for photo %s: %w",
				match.PhotoID,
				err,
			)
		}

		results = append(
			results,
			face.FaceSearchResult{
				FaceID:     match.ID,
				PhotoID:    match.PhotoID,
				Confidence: match.Confidence,
				Distance:   match.Distance,
				URL:        url,
			},
		)
	}

    return results, nil
}