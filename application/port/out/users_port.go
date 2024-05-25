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

type Like struct {
	UserID string
	NewsID string
}

type UsersPort interface {
	GetAll() ([]db.User, error)
	Insert(user User) error
	Update(user User) error
	GetByAuthID(authID string) (user User, err error)
	Like(like Like) error
	Unlike(like Like) error
	DisLike(like Like) error
	UnDisLike(like Like) error
}
