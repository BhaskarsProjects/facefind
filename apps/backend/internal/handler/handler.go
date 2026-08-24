package handler

import (
	"github.com/bhaskarvb28/facefind/internal/server"
	"github.com/bhaskarvb28/facefind/internal/service"
)

type Handlers struct {
	Auth   *AuthHandler
	Folder *FolderHandler
	Photo  *PhotoHandler
	Face *FaceHandler
}

func NewHandlers(
	s *server.Server,
	services *service.Services,
) *Handlers {
	return &Handlers{
		Auth:   NewAuthHandler(s, services.Auth),
		Folder: NewFolderHandler(s, services.Folder),
		Photo:  NewPhotoHandler(s, services.Photo),
		Face: NewFaceHandler(services.Face),
	}
}
