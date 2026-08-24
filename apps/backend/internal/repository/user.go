package repository

import (
	"context"
	"fmt"

	"github.com/bhaskarvb28/facefind/internal/model/auth"
	"github.com/bhaskarvb28/facefind/internal/server"
	"github.com/jackc/pgx/v5"
)

type UserRepository struct {
	server *server.Server
}

func NewUserRepository(server *server.Server) *UserRepository {
	return &UserRepository{
		server: server,
	}
}

func (r *UserRepository) CreateUser(
	ctx context.Context,
	auth0Sub string,
) error {
	stmt := `
		INSERT INTO users (
			auth0_sub
		)
		VALUES (
			@auth0_sub
		)
	`

	_, err := r.server.DB.Pool.Exec(
		ctx,
		stmt,
		pgx.NamedArgs{
			"auth0_sub": auth0Sub,
		},
	)
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	return nil
}

func (r *UserRepository) GetByAuth0Sub(
	ctx context.Context,
	auth0Sub string,
) (*auth.User, error) {
	stmt := `
		SELECT
			id,
			auth0_sub,
			role,
			status,
			created_at,
			updated_at
		FROM users
		WHERE auth0_sub = @auth0_sub
	`

	rows, err := r.server.DB.Pool.Query(
		ctx,
		stmt,
		pgx.NamedArgs{
			"auth0_sub": auth0Sub,
		},
	)
	if err != nil {
		return nil, fmt.Errorf(
			"failed to query user auth0_sub=%s: %w",
			auth0Sub,
			err,
		)
	}

	user, err := pgx.CollectOneRow(
		rows,
		pgx.RowToStructByName[auth.User],
	)
	if err != nil {
		return nil, fmt.Errorf(
			"failed to find user auth0_sub=%s: %w",
			auth0Sub,
			err,
		)
	}

	return &user, nil
}
