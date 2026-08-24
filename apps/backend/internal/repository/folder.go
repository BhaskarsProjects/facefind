package repository

import (
	"context"

	"github.com/bhaskarvb28/facefind/internal/model/folder"
	"github.com/bhaskarvb28/facefind/internal/server"
	"github.com/jackc/pgx/v5"
)

type FolderRepository struct {
	server *server.Server
}

func NewFolderRepository(server *server.Server) *FolderRepository {
	return &FolderRepository{
		server: server,
	}
}

func (r *FolderRepository) Create(
	ctx context.Context,
	ownerSub string,
	name string,
) (*folder.Folder, error) {
	folder := &folder.Folder{}

	err := r.server.DB.Pool.QueryRow(
		ctx,
		`
		INSERT INTO folders (
			owner_sub,
			name
		)
		VALUES (
			@owner_sub,
			@name
		)
		RETURNING
			id,
			owner_sub,
			name,
			created_at,
			updated_at
		`,
		pgx.NamedArgs{
			"owner_sub": ownerSub,
			"name":      name,
		},
	).Scan(
		&folder.ID,
		&folder.OwnerSub,
		&folder.Name,
		&folder.CreatedAt,
		&folder.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return folder, nil
}

func (r *FolderRepository) Get(
	ctx context.Context,
	ownerSub string,
) ([]*folder.Folder, error) {

	rows, err := r.server.DB.Pool.Query(
		ctx,
		`
		SELECT
			id,
			owner_sub,
			name,
			created_at,
			updated_at
		FROM folders
		WHERE owner_sub = @owner_sub
		ORDER BY created_at DESC
		`,
		pgx.NamedArgs{
			"owner_sub": ownerSub,
		},
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	folders := make([]*folder.Folder, 0)

	for rows.Next() {
		f := &folder.Folder{}

		if err := rows.Scan(
			&f.ID,
			&f.OwnerSub,
			&f.Name,
			&f.CreatedAt,
			&f.UpdatedAt,
		); err != nil {
			return nil, err
		}

		folders = append(folders, f)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return folders, nil
}

func (r *FolderRepository) UpdateName(
	ctx context.Context,
	ownerSub string,
	id string,
	name string,
) (*folder.Folder, error) {

	var f folder.Folder

	err := r.server.DB.Pool.QueryRow(
		ctx,
		`
		UPDATE folders
		SET
			name = $1,
			updated_at = NOW()
		WHERE
			id = $2
			AND owner_sub = $3
		RETURNING
			id,
			owner_sub,
			name,
			created_at,
			updated_at
		`,
		name,
		id,
		ownerSub,
	).Scan(
		&f.ID,
		&f.OwnerSub,
		&f.Name,
		&f.CreatedAt,
		&f.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &f, nil
}
