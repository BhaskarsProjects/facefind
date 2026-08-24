package ai

import (
	"net/http"
	"time"

	"github.com/bhaskarvb28/facefind/internal/server"
)

type MLClient struct {
	baseURL string
	token   string
	client  *http.Client
}

func NewMLClient(
	server *server.Server,
) *MLClient {
	return &MLClient{
		baseURL: server.Config.AIWorker.BaseURL,
		client: &http.Client{
			Timeout: 60 * time.Second,
		},
		token: server.Config.AIWorker.Token,
	}
}