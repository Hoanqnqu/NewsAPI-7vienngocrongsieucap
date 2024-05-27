package service

import (
	inport "news-api/application/port/in"
	outport "news-api/application/port/out"

	"github.com/google/uuid"
)

type UsersService struct {
	usersPort outport.UsersPort
}

func NewUsersService(userPort outport.UsersPort) *UsersService {
	return &UsersService{usersPort: userPort}
}

func (g *UsersService) GetAll() ([]*inport.User, error) {
	usersList, err := g.usersPort.GetAll()
	if err != nil {
		return nil, err
	}
	return func() []*inport.User {
		result := make([]*inport.User, len(usersList))
		for i, v := range usersList {
			result[i] = MapUser(v)
		}
		return result
	}(), nil
}

func (g *UsersService) Insert(user *inport.CreateUserPayload) error {
	return g.usersPort.Insert(outport.User{
		ID:       uuid.New(),
		AuthID:   user.AuthID,
		Email:    user.Email,
		Name:     user.Name,
		Role:     user.Role,
		ImageUrl: user.ImageUrl,
	})
}

func (g *UsersService) Update(user *inport.UpdateUserPayload) error {
	return g.usersPort.Update(outport.User{
		ID:       user.ID,
		AuthID:   user.AuthID,
		Email:    user.Email,
		Name:     user.Name,
		Role:     user.Role,
		ImageUrl: user.ImageUrl,
	})
}
func (g *UsersService) GetUserByAuthID(authID string) (user *inport.UpdateUserPayload, err error) {
	u, err := g.usersPort.GetByAuthID(authID)
	if err != nil {
		return nil, err
	}
	return &inport.UpdateUserPayload{
		ID:       u.ID,
		AuthID:   u.AuthID,
		Email:    u.Email,
		Name:     u.Name,
		Role:     u.Role,
		ImageUrl: u.ImageUrl,
	}, nil
}
func (g *UsersService) GetAdmin(email string, password string) (user *inport.UpdateUserPayload, err error) {
	u, err := g.usersPort.GetAdmin(email, password)
	if err != nil {
		return nil, err
	}
	if u == nil {
		return nil, nil
	}
	return &inport.UpdateUserPayload{
		ID:       u.ID,
		AuthID:   u.AuthID,
		Email:    u.Email,
		Name:     u.Name,
		Role:     u.Role,
		ImageUrl: u.ImageUrl,
	}, nil
}

func (g *UsersService) Like(like *inport.Like) error {
	return g.usersPort.Like(outport.Like{
		UserID: like.UserId,
		NewsID: like.NewsId,
	})
}

func (g *UsersService) Unlike(like *inport.Like) error {
	return g.usersPort.Unlike(outport.Like{
		UserID: like.UserId,
		NewsID: like.NewsId,
	})
}

func (g *UsersService) DisLike(like *inport.Like) error {
	return g.usersPort.DisLike(outport.Like{
		UserID: like.UserId,
		NewsID: like.NewsId,
	})
}

func (g *UsersService) UnDisLike(like *inport.Like) error {
	return g.usersPort.UnDisLike(outport.Like{
		UserID: like.UserId,
		NewsID: like.NewsId,
	})
}
