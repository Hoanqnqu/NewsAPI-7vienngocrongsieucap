package outport

import (
	"news-api/internal/db"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID
	AuthID    string
	Email     string
	Name      string
	Role      string
	ImageUrl  string
	CreatedAt time.Time
}

type UsersPort interface {
	GetAll() ([]db.User, error)
	Insert(user User) error
	Update(user User) error
	GetByAuthID(authID string) (user User, err error)
}
