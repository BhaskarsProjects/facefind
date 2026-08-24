package repository

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/bhaskarvb28/facefind/internal/model/face"
	"github.com/bhaskarvb28/facefind/internal/server"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type FaceRepository struct {
	server *server.Server
}

func NewFaceRepository(
	server *server.Server,
) *FaceRepository {
	return &FaceRepository{
		server: server,
	}
}

func (r *FaceRepository) Create(
	ctx context.Context,
	photoID uuid.UUID,
	x1 float32,
	y1 float32,
	x2 float32,
	y2 float32,
	confidence float32,
	landmarks [][]float32,
	embedding []float32,
) error {
	landmarksJSON, err := json.Marshal(landmarks)
	if err != nil {
		return fmt.Errorf("marshal landmarks: %w", err)
	}

	embeddingValue := make([]byte, 0, len(embedding)*8+2)

	embeddingValue = append(
		embeddingValue,
		'[',
	)

	for i, value := range embedding {
		if i > 0 {
			embeddingValue = append(
				embeddingValue,
				',',
			)
		}

		embeddingValue = append(
			embeddingValue,
			fmt.Sprintf("%f", value)...,
		)
	}

	embeddingValue = append(
		embeddingValue,
		']',
	)

	const query = `
		INSERT INTO faces (
			photo_id,
			x1,
			y1,
			x2,
			y2,
			confidence,
			landmarks,
			embedding
		)
		VALUES (
			@photo_id,
			@x1,
			@y1,
			@x2,
			@y2,
			@confidence,
			@landmarks,
			@embedding
		)
	`

	args := pgx.NamedArgs{
		"photo_id":   photoID,
		"x1":         x1,
		"y1":         y1,
		"x2":         x2,
		"y2":         y2,
		"confidence": confidence,
		"landmarks":  landmarksJSON,
		"embedding":  string(embeddingValue),
	}

	_, err = r.server.DB.Pool.Exec(
		ctx,
		query,
		args,
	)
	if err != nil {
		return fmt.Errorf("create face: %w", err)
	}

	return nil
}

func (r *FaceRepository) SearchSimilar(
    ctx context.Context,
    folderID uuid.UUID,
    embedding []float32,
    limit int,
    maxDistance float32,
) ([]face.FaceMatch, error) {
	if len(embedding) != 512 {
		return nil, fmt.Errorf(
			"embedding must contain 512 values, got %d",
			len(embedding),
		)
	}

	if limit <= 0 {
		return nil, fmt.Errorf(
			"limit must be greater than zero",
		)
	}

	embeddingValue := make([]byte, 0, len(embedding)*8+2)

	embeddingValue = append(
		embeddingValue,
		'[',
	)

	for i, value := range embedding {
		if i > 0 {
			embeddingValue = append(
				embeddingValue,
				',',
			)
		}

		embeddingValue = append(
			embeddingValue,
			fmt.Sprintf("%f", value)...,
		)
	}

	embeddingValue = append(
		embeddingValue,
		']',
	)

	const query = `
		SELECT
			f.id,
			f.photo_id,
			f.confidence,
			f.embedding <=> @embedding AS distance,
			p.storage_key
		FROM faces f
		JOIN photos p
			ON p.id = f.photo_id
		WHERE
			p.folder_id = @folder_id
			AND f.embedding <=> @embedding <= @max_distance
		ORDER BY f.embedding <=> @embedding
		LIMIT @limit
	`

	args := pgx.NamedArgs{
		"embedding":   string(embeddingValue),
		"folder_id":   folderID,
		"limit":       limit,
		"max_distance": maxDistance,
	}

	rows, err := r.server.DB.Pool.Query(
		ctx,
		query,
		args,
	)
	if err != nil {
		return nil, fmt.Errorf(
			"search similar faces: %w",
			err,
		)
	}
	defer rows.Close()

	matches := make([]face.FaceMatch, 0)

	for rows.Next() {
		var match face.FaceMatch

		if err := rows.Scan(
			&match.ID,
			&match.PhotoID,
			&match.Confidence,
			&match.Distance,
			&match.StorageKey,
		); err != nil {
			return nil, fmt.Errorf(
				"scan face match: %w",
				err,
			)
		}

		matches = append(
			matches,
			match,
		)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf(
			"iterate face matches: %w",
			err,
		)
	}

	return matches, nil
}