package inport

import (
	"time"

	"github.com/google/uuid"
)

type News struct {
	ID          string     `json:"id"`
	Author      string     `json:"author"`
	Title       string     `json:"title"`
	Content     string     `json:"content"`
	Description string     `json:"description"`
	URL         string     `json:"url"`
	ImageURL    string     `json:"image_url"`
	PublishAt   time.Time  `json:"publish_at"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
	IsLiked     bool       `json:"isLiked"`
	IsDisliked  bool       `json:"isDisliked"`
}
type CreateNewsPayload struct {
	Author      string      `json:"author"`
	Title       string      `json:"title"`
	Content     string      `json:"content"`
	Description string      `json:"description"`
	URL         string      `json:"url"`
	ImageURL    string      `json:"image_url"`
	PublishAt   time.Time   `json:"publish_at"`
	Categories  []uuid.UUID `json:"categories"`
}

type UpdateNewsPayload struct {
	ID          uuid.UUID `json:"id"`
	Author      string    `json:"author"`
	Title       string    `json:"title"`
	Content     string    `json:"content"`
	Description string    `json:"description"`
	URL         string    `json:"url"`
	ImageURL    string    `json:"image_url"`
	PublishAt   time.Time `json:"publish_at"`
}

type NewsUseCase interface {
	GetAll() ([]*News, error)
	Insert(news *CreateNewsPayload) error
	Update(news *UpdateNewsPayload) error
	GetNewsByID(newsID, userID string) (*News, error)
}
