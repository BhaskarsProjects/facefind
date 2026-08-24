package face

type CreateFacesRequest struct {
	Faces []CreateFaceInput `json:"faces"`
}

type CreateFaceInput struct {
	X1         float32     `json:"x1"`
	Y1         float32     `json:"y1"`
	X2         float32     `json:"x2"`
	Y2         float32     `json:"y2"`
	Confidence float32     `json:"confidence"`
	Landmarks  [][]float32 `json:"landmarks"`
	Embedding  []float32   `json:"embedding"`
}