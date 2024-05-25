package outAdapter

import (
	"context"
	"github.com/google/uuid"
	outport "news-api/application/port/out"
	db "news-api/internal/db"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserAdapter struct {
	pool *pgxpool.Pool
}

func NewUserAdapter(pool *pgxpool.Pool) *UserAdapter {
	return &UserAdapter{pool: pool}
}

func (u *UserAdapter) GetAll() ([]db.User, error) {
	query := db.New(u.pool)

	return query.GetAllUsers(context.Background())

}

func (u *UserAdapter) Insert(user outport.User) error {
	query := db.New(u.pool)
	err := query.InsertUser(context.Background(), db.InsertUserParams{
		ID: pgtype.UUID{
			Bytes: user.ID,
			Valid: true,
		},
		AuthID: user.AuthID,
		Email: pgtype.Text{
			String: user.Email,
			Valid:  true,
		},
		Name: pgtype.Text{
			String: user.Name,
			Valid:  true,
		},
		Role: pgtype.Text{
			String: user.Role,
			Valid:  true,
		},
		ImageUrl: pgtype.Text{
			String: user.ImageUrl,
			Valid:  true,
		},
	})
	return err
}

func (u *UserAdapter) Update(user outport.User) error {
	query := db.New(u.pool)
	err := query.UpdateUser(context.Background(), db.UpdateUserParams{
		Name: pgtype.Text{
			String: user.Name,
			Valid:  true,
		},
		ImageUrl: pgtype.Text{
			String: user.ImageUrl,
			Valid:  true,
		},
		ID: pgtype.UUID{
			Bytes: user.ID,
			Valid: true,
		},
	})
	return err
}

func (u *UserAdapter) GetByAuthID(authID string) (outUser outport.User, err error) {
	query := db.New(u.pool)
	dbUser, err := query.GetUserByAuthID(context.Background(), authID)
	if err != nil {
		return outport.User{}, err
	}
	outUser = outport.User{
		ID:        dbUser.ID.Bytes,
		AuthID:    dbUser.AuthID,
		Email:     dbUser.Email.String,
		CreatedAt: dbUser.CreatedAt.Time,
		Name:      dbUser.Name.String,
		Role:      dbUser.Role.String,
		ImageUrl:  dbUser.ImageUrl.String,
	}
	return outUser, nil
}

func (u *UserAdapter) Like(like outport.Like) error {
	query := db.New(u.pool)
	return query.InsertLike(context.Background(), db.InsertLikeParams{
		NewsID: pgtype.UUID{
			Bytes: uuid.MustParse(like.NewsID),
			Valid: true,
		},
		UserID: pgtype.UUID{
			Bytes: uuid.MustParse(like.UserID),
			Valid: true,
		},
	})
}

func (u *UserAdapter) Unlike(like outport.Like) error {
	query := db.New(u.pool)
	return query.DeleteLike(context.Background(), db.DeleteLikeParams{
		NewsID: pgtype.UUID{
			Bytes: uuid.MustParse(like.NewsID),
			Valid: true,
		},
		UserID: pgtype.UUID{
			Bytes: uuid.MustParse(like.UserID),
			Valid: true,
		},
	})
}

func (u *UserAdapter) DisLike(like outport.Like) error {
	query := db.New(u.pool)
	return query.InsertDisLike(context.Background(), db.InsertDisLikeParams{
		NewsID: pgtype.UUID{
			Bytes: uuid.MustParse(like.NewsID),
			Valid: true,
		},
		UserID: pgtype.UUID{
			Bytes: uuid.MustParse(like.UserID),
			Valid: true,
		},
	})
}

func (u *UserAdapter) UnDisLike(like outport.Like) error {
	query := db.New(u.pool)
	return query.DeleteDisLike(context.Background(), db.DeleteDisLikeParams{
		NewsID: pgtype.UUID{
			Bytes: uuid.MustParse(like.NewsID),
			Valid: true,
		},
		UserID: pgtype.UUID{
			Bytes: uuid.MustParse(like.UserID),
			Valid: true,
		},
	})
}
