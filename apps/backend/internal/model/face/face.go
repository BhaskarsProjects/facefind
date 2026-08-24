package face

import "github.com/google/uuid"

type Face struct {
	ID         uuid.UUID `db:"id"`
	PhotoID    uuid.UUID `db:"photo_id"`
	X1         float32  `db:"x1"`
	Y1         float32  `db:"y1"`
	X2         float32  `db:"x2"`
	Y2         float32  `db:"y2"`
	Confidence float32  `db:"confidence"`
	Landmarks  []byte   `db:"landmarks"`
	Embedding  []float32 `db:"embedding"`
}

type FaceMatch struct {
    ID         uuid.UUID
    PhotoID    uuid.UUID
    Confidence float32
    Distance   float32
    StorageKey string
}

type FaceSearchResult struct {
    FaceID     uuid.UUID `json:"face_id"`
    PhotoID    uuid.UUID `json:"photo_id"`
    Confidence float32   `json:"confidence"`
    Distance   float32   `json:"distance"`
    URL        string    `json:"url"`
}