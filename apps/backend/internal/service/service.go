package service

import (
	"fmt"

	"github.com/bhaskarvb28/facefind/internal/lib/ai"
	"github.com/bhaskarvb28/facefind/internal/lib/aws"
	"github.com/bhaskarvb28/facefind/internal/repository"
	"github.com/bhaskarvb28/facefind/internal/server"
)

type Services struct {
	Auth   *AuthService
	Folder *FolderService
	Photo  *PhotoService
	Face   *FaceService
}

func NewServices(
	s *server.Server,
	repos *repository.Repositories,
) (*Services, error) {

	authService, err := NewAuthService(s, repos.User)
	if err != nil {
		return nil, err
	}

	folderService, err := NewFolderService(s, repos.Folder)
	if err != nil {
		return nil, err
	}

	r2Client, err := aws.NewS3Client(s)
	if err != nil {
		return nil, fmt.Errorf(
			"failed to create R2 storage: %w",
			err,
		)
	}

	mlClient := ai.NewMLClient(s)

	photoService, err := NewPhotoService(
		s,
		repos.Photo,
		r2Client,
		mlClient,
	)

	if err != nil {
		return nil, err
	}

	faceService := NewFaceService(
		s,
		repos.Face,
		repos.Photo,
		mlClient,
		r2Client,
	)


	return &Services{
		Auth:   authService,
		Folder: folderService,
		Photo:  photoService,
		Face:   faceService,
	}, nil
}