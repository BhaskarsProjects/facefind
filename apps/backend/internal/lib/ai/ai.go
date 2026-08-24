package ai

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"mime/multipart"
	"net/http"
)

type EmbedResponse struct {
	Confidence         float32   `json:"confidence"`
	EmbeddingDimension int       `json:"embedding_dimension"`
	Embedding          []float32 `json:"embedding"`
}

type ProcessResponse struct {
	PhotoID   string `json:"photo_id"`
	FaceCount int    `json:"face_count"`
}

func (c *MLClient) Embed(
	ctx context.Context,
	image []byte,
	filename string,
) (*EmbedResponse, error) {
	var body bytes.Buffer

	writer := multipart.NewWriter(&body)

	part, err := writer.CreateFormFile(
		"image",
		filename,
	)
	if err != nil {
		return nil, fmt.Errorf(
			"create image form: %w",
			err,
		)
	}

	if _, err := part.Write(image); err != nil {
		return nil, fmt.Errorf(
			"write image: %w",
			err,
		)
	}

	if err := writer.Close(); err != nil {
		return nil, fmt.Errorf(
			"close multipart writer: %w",
			err,
		)
	}

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		c.baseURL+"/embed",
		&body,
	)
	if err != nil {
		return nil, fmt.Errorf(
			"create ML request: %w",
			err,
		)
	}

	req.Header.Set(
		"Content-Type",
		writer.FormDataContentType(),
	)

	req.Header.Set(
		"Authorization",
		"Bearer "+c.token,
	)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf(
			"call ML service: %w",
			err,
		)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(
			"ML service returned status %d",
			resp.StatusCode,
		)
	}

	var result EmbedResponse

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf(
			"decode ML response: %w",
			err,
		)
	}

	if result.EmbeddingDimension != 512 {
		return nil, fmt.Errorf(
			"invalid embedding dimension: %d",
			result.EmbeddingDimension,
		)
	}

	if len(result.Embedding) != 512 {
		return nil, fmt.Errorf(
			"invalid embedding length: %d",
			len(result.Embedding),
		)
	}

	return &result, nil
}

func (c *MLClient) Process(
	ctx context.Context,
	photoID string,
	storageKey string,
) (*ProcessResponse, error) {
	requestBody := struct {
		PhotoID    string `json:"photo_id"`
		StorageKey string `json:"storage_key"`
	}{
		PhotoID:    photoID,
		StorageKey: storageKey,
	}

	body, err := json.Marshal(requestBody)
	if err != nil {
		return nil, fmt.Errorf(
			"marshal process request: %w",
			err,
		)
	}

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		c.baseURL+"/process",
		bytes.NewReader(body),
	)
	if err != nil {
		return nil, fmt.Errorf(
			"create ML process request: %w",
			err,
		)
	}

	req.Header.Set(
		"Content-Type",
		"application/json",
	)

	req.Header.Set(
		"Authorization",
		"Bearer "+c.token,
	)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf(
			"call ML process service: %w",
			err,
		)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(
			"ML process service returned status %d",
			resp.StatusCode,
		)
	}

	var result ProcessResponse

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf(
			"decode ML process response: %w",
			err,
		)
	}

	if result.PhotoID != photoID {
		return nil, fmt.Errorf(
			"ML process response photo ID mismatch: got %s, expected %s",
			result.PhotoID,
			photoID,
		)
	}

	return &result, nil
}