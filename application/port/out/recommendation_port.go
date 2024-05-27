package outport

import (
	"context"
	"github.com/google/uuid"
)

type RecommendationSystem interface {
	InsertUser(ctx context.Context, id uuid.UUID) error
	InsertNews(ctx context.Context, id uuid.UUID, categories []uuid.UUID) error
}
