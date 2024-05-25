package inport

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        string
	AuthID    string
	Email     string
	Name      string
	Role      string
	ImageUrl  string
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}

type CreateUserPayload struct {
	AuthID   string
	Email    string
	Name     string
	Role     string
	ImageUrl string
}

type UpdateUserPayload struct {
	ID       uuid.UUID
	AuthID   string
	Email    string
	Name     string
	Role     string
	ImageUrl string
}

type UsersUseCase interface {
	GetAll() ([]*User, error)
	Insert(user *CreateUserPayload) error
	Update(user *UpdateUserPayload) error
	GetUserByAuthID(authID string) (user *CreateUserPayload, err error)
}
