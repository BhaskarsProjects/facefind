package repository

import (
	"context"
	"net/http"

	"github.com/bhaskarvb28/facefind/internal/model/photo"
	"github.com/bhaskarvb28/facefind/internal/server"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/labstack/echo/v4"
)

type PhotoRepository struct {
	server *server.Server
}

func NewPhotoRepository(
	server *server.Server,
) *PhotoRepository {
	return &PhotoRepository{
		server: server,
	}
}

func (r *PhotoRepository) VerifyFolderOwner(
	ctx context.Context,
	ownerSub string,
	folderID string,
) error {
	var exists bool

	err := r.server.DB.Pool.QueryRow(
		ctx,
		`
		SELECT EXISTS (
			SELECT 1
			FROM folders
			WHERE id = $1
			  AND owner_sub = $2
		)
		`,
		folderID,
		ownerSub,
	).Scan(&exists)

	if err != nil {
		return err
	}

	if !exists {
		return echo.NewHTTPError(
			http.StatusNotFound,
			"folder not found",
		)
	}

	return nil
}

func (r *PhotoRepository) Create(
	ctx context.Context,
	folderID string,
	id string,
	originalFilename string,
	mimeType string,
	sizeBytes int64,
	storageKey string,
) error {
	_, err := r.server.DB.Pool.Exec(
		ctx,
		`
		INSERT INTO photos (
			id,
			folder_id,
			storage_key,
			original_filename,
			mime_type,
			size_bytes
		)
		VALUES ($1, $2, $3, $4, $5, $6)
		`,
		id,
		folderID,
		storageKey,
		originalFilename,
		mimeType,
		sizeBytes,
	)

	return err
}

func (r *PhotoRepository) GetByFolderID(
	ctx context.Context,
	folderID uuid.UUID,
	userSub string,
) ([]*photo.Photo, error) {
	const query = `
		SELECT
			p.id,
			p.folder_id,
			p.storage_key,
			p.original_filename,
			p.mime_type,
			p.size_bytes,
			p.created_at,
			p.updated_at
		FROM photos p
		JOIN folders f
			ON f.id = p.folder_id
		LEFT JOIN folder_members fm
			ON fm.folder_id = f.id
			AND fm.user_sub = @user_sub
		WHERE p.folder_id = @folder_id
		  AND (
				f.owner_sub = @user_sub
				OR fm.user_sub IS NOT NULL
		  )
		ORDER BY p.created_at DESC
	`

	args := pgx.NamedArgs{
		"folder_id": folderID,
		"user_sub":  userSub,
	}

	rows, err := r.server.DB.Pool.Query(ctx, query, args)
	if err != nil {
		return nil, err
	}

	photos, err := pgx.CollectRows(
		rows,
		pgx.RowToAddrOfStructByName[photo.Photo],
	)
	if err != nil {
		return nil, err
	}

	return photos, nil
}

func (r *PhotoRepository) GetByID(
	ctx context.Context,
	folderID uuid.UUID,
	photoID uuid.UUID,
	userSub string,
) (*photo.Photo, error) {
	const query = `
		SELECT
			p.id,
			p.folder_id,
			p.storage_key,
			p.original_filename,
			p.mime_type,
			p.size_bytes,
			p.created_at,
			p.updated_at
		FROM photos p
		JOIN folders f
			ON f.id = p.folder_id
		LEFT JOIN folder_members fm
			ON fm.folder_id = f.id
			AND fm.user_sub = @user_sub
		WHERE p.id = @photo_id
		  AND p.folder_id = @folder_id
		  AND (
				f.owner_sub = @user_sub
				OR (
					fm.user_sub IS NOT NULL
					AND fm.role = 'editor'
				)
		  )
	`

	args := pgx.NamedArgs{
		"photo_id":  photoID,
		"folder_id": folderID,
		"user_sub":  userSub,
	}

	var p photo.Photo

	err := r.server.DB.Pool.QueryRow(
		ctx,
		query,
		args,
	).Scan(
		&p.ID,
		&p.FolderID,
		&p.StorageKey,
		&p.OriginalFilename,
		&p.MimeType,
		&p.SizeBytes,
		&p.CreatedAt,
		&p.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &p, nil
}

func (r *PhotoRepository) Delete(
	ctx context.Context,
	folderID uuid.UUID,
	photoID uuid.UUID,
	userSub string,
) error {
	const query = `
		DELETE FROM photos p
		USING folders f
		LEFT JOIN folder_members fm
			ON fm.folder_id = f.id
			AND fm.user_sub = @user_sub
		WHERE p.id = @photo_id
		  AND p.folder_id = @folder_id
		  AND f.id = p.folder_id
		  AND (
				f.owner_sub = @user_sub
				OR (
					fm.user_sub IS NOT NULL
					AND fm.role = 'editor'
				)
		  )
	`

	args := pgx.NamedArgs{
		"photo_id":  photoID,
		"folder_id": folderID,
		"user_sub":  userSub,
	}

	tag, err := r.server.DB.Pool.Exec(
		ctx,
		query,
		args,
	)
	if err != nil {
		return err
	}

	if tag.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}

	return nil
}

func (r *PhotoRepository) Exists(
	ctx context.Context,
	photoID uuid.UUID,
) (bool, error) {
	const query = `
		SELECT EXISTS (
			SELECT 1
			FROM photos
			WHERE id = @photo_id
		)
	`

	args := pgx.NamedArgs{
		"photo_id": photoID,
	}

	var exists bool

	err := r.server.DB.Pool.QueryRow(
		ctx,
		query,
		args,
	).Scan(&exists)

	return exists, err
}