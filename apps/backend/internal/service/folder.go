package service

import (
	"context"

	"github.com/bhaskarvb28/facefind/internal/model/folder"
	"github.com/bhaskarvb28/facefind/internal/repository"
	"github.com/bhaskarvb28/facefind/internal/server"
)

type FolderService struct {
	server     *server.Server
	repository *repository.FolderRepository
}

func NewFolderService(
	server *server.Server,
	repository *repository.FolderRepository,
) (*FolderService, error) {
	return &FolderService{
		server:     server,
		repository: repository,
	}, nil
}

func (s *FolderService) Create(
	ctx context.Context,
	ownerSub string,
	name string,
) (*folder.Folder, error) {
	return s.repository.Create(ctx, ownerSub, name)
}

func (s *FolderService) Get(
	ctx context.Context,
	ownerSub string,
) ([]*folder.Folder, error) {
	return s.repository.Get(ctx, ownerSub)
}

func (s *FolderService) UpdateName(
	ctx context.Context,
	ownerSub string,
	id string,
	name string,
) (*folder.Folder, error) {
	return s.repository.UpdateName(
		ctx,
		ownerSub,
		id,
		name,
	)
}
