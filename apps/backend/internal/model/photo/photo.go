package photo

import (
	"time"

	"github.com/google/uuid"
)

type Photo struct {
	ID               uuid.UUID `db:"id" json:"id"`
	FolderID         uuid.UUID `db:"folder_id" json:"folderId"`
	StorageKey       string    `db:"storage_key" json:"storageKey"`
	OriginalFilename string    `db:"original_filename" json:"originalFilename"`
	MimeType         string    `db:"mime_type" json:"mimeType"`
	SizeBytes        int64     `db:"size_bytes" json:"sizeBytes"`
	CreatedAt        time.Time `db:"created_at" json:"createdAt"`
	UpdatedAt        time.Time `db:"updated_at" json:"updatedAt"`
}