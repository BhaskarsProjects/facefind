package photo

import (
	"time"

	"github.com/google/uuid"
)

type PhotoResponse struct {
	ID               uuid.UUID `json:"id"`
	FolderID         uuid.UUID `json:"folderId"`
	OriginalFilename string    `json:"originalFilename"`
	MimeType         string    `json:"mimeType"`
	SizeBytes        int64     `json:"sizeBytes"`
	CreatedAt        time.Time `json:"createdAt"`
	UpdatedAt        time.Time `json:"updatedAt"`
	URL              string    `json:"url"`
}