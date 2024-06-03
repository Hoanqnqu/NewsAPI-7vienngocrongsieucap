package outport

import (
	"news-api/internal/db"
	"time"

	"github.com/google/uuid"
)

type News struct {
	ID          uuid.UUID
	Title       string
	Content     string
	Description string
	Author      string
	Url         string
	ImageUrl    string
	PublishAt   time.Time
	Categories  []uuid.UUID
}

type NewsPort interface {
	GetAll() ([]db.News, error)
	Insert(news News) error
	Update(news News) error
	GetNewsByID(newsID, userID string) (*db.News, bool, bool, error)
}
