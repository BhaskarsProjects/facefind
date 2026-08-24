package session

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

const (
	sessionPrefix = "session:"
	sessionTTL    = 7 * 24 * time.Hour
)

type SessionService struct {
	redis *redis.Client
}

func NewSessionService(redis *redis.Client) *SessionService {
	return &SessionService{
		redis: redis,
	}
}

func (s *SessionService) Create(
	ctx context.Context,
	userID string,
) (string, error) {

	sessionID, err := generateSessionID()
	if err != nil {
		return "", fmt.Errorf("failed to generate session ID: %w", err)
	}

	key := sessionPrefix + sessionID

	// fmt.Printf(
	// 	"[SESSION] creating session: key=%s userID=%s ttl=%s\n",
	// 	key,
	// 	userID,
	// 	sessionTTL,
	// )

	if err := s.redis.Set(
		ctx,
		key,
		userID,
		sessionTTL,
	).Err(); err != nil {
		return "", fmt.Errorf("failed to store session: %w", err)
	}

	// fmt.Printf(
	// 	"[SESSION] session created successfully: key=%s\n",
	// 	key,
	// )

	return sessionID, nil
}

func generateSessionID() (string, error) {
	b := make([]byte, 32)

	if _, err := rand.Read(b); err != nil {
		return "", err
	}

	return hex.EncodeToString(b), nil
}
