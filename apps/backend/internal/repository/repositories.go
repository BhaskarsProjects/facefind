package repository

import "github.com/bhaskarvb28/facefind/internal/server"

type Repositories struct {
	User   *UserRepository
	Folder *FolderRepository
	Photo *PhotoRepository
	Face *FaceRepository
}

func NewRepositories(s *server.Server) *Repositories {
	return &Repositories{
		User:   NewUserRepository(s),
		Folder: NewFolderRepository(s),
		Photo: NewPhotoRepository(s),
		Face: NewFaceRepository(s),
	}
}
