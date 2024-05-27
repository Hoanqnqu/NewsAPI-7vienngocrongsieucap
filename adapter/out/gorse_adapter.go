package outAdapter

import (
	"context"
	"github.com/google/uuid"
	"github.com/zhenghaoz/gorse/client"
)

type GorseAdapter struct {
	gorse *client.GorseClient
}

func NewGorseAdapter(gorse *client.GorseClient) *GorseAdapter {
	return &GorseAdapter{gorse: gorse}
}

func (g *GorseAdapter) InsertUser(ctx context.Context, id uuid.UUID) error {
	_, err := g.gorse.InsertUser(ctx, client.User{
		UserId: id.String(),
	})
	return err
}

func (g *GorseAdapter) InsertNews(ctx context.Context, id uuid.UUID, categories []uuid.UUID) error {
	_, err := g.gorse.InsertItem(ctx, client.Item{
		ItemId:   id.String(),
		IsHidden: false,
		Categories: func() []string {
			result := make([]string, len(categories))
			for i, v := range categories {
				result[i] = v.String()
			}
			return result
		}(),
	})
	return err
}
