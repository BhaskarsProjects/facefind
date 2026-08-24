package folder

import (
	"time"

	"github.com/google/uuid"
)

type Folder struct {
	ID        uuid.UUID `json:"id"`
	OwnerSub  string    `json:"owner_sub"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
