package outAdapter

import (
	"context"
	"github.com/google/uuid"
	outport "news-api/application/port/out"
	db "news-api/internal/db"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

type NewsAdapter struct {
	pool *pgxpool.Pool
}

func NewNewsAdapter(pool *pgxpool.Pool) *NewsAdapter {
	return &NewsAdapter{pool: pool}
}

func (u *NewsAdapter) GetAll() ([]db.News, error) {
	query := db.New(u.pool)

	return query.GetAllNews(context.Background())
}

func (u *NewsAdapter) Insert(news outport.News) error {
	query := db.New(u.pool)
	err := query.InsertNews(context.Background(), db.InsertNewsParams{
		ID: pgtype.UUID{
			Bytes: news.ID,
			Valid: true,
		},
		Author: pgtype.Text{
			String: news.Author,
			Valid:  true,
		},
		Title: pgtype.Text{
			String: news.Title,
			Valid:  true,
		},
		Content: pgtype.Text{
			String: news.Content,
			Valid:  true,
		},
		Description: pgtype.Text{
			String: news.Description,
			Valid:  true,
		},
		Url: pgtype.Text{
			String: news.Url,
			Valid:  true,
		},
		ImageUrl: pgtype.Text{
			String: news.ImageUrl,
			Valid:  true,
		},
		PublishAt: pgtype.Timestamp{
			Time:  news.PublishAt,
			Valid: true,
		},
	})
	if err != nil {
		for _, v := range news.Categories {
			err = query.InsertHasCategory(context.Background(), db.InsertHasCategoryParams{
				NewsID: pgtype.UUID{
					Bytes: news.ID,
					Valid: true,
				},
				CategoryID: pgtype.UUID{
					Bytes: v,
					Valid: true,
				},
			})
			if err != nil {
				return err
			}
		}
	}
	return err
}

func (u *NewsAdapter) Update(news outport.News) error {
	query := db.New(u.pool)

	err := query.UpdateNews(context.Background(), db.UpdateNewsParams{
		ID: pgtype.UUID{
			Bytes: news.ID,
			Valid: true,
		},
		Author: pgtype.Text{
			String: news.Author,
			Valid:  true,
		},
		Title: pgtype.Text{
			String: news.Title,
			Valid:  true,
		},
		Content: pgtype.Text{
			String: news.Content,
			Valid:  true,
		},
		Description: pgtype.Text{
			String: news.Description,
			Valid:  true,
		},
		Url: pgtype.Text{
			String: news.Url,
			Valid:  true,
		},
		ImageUrl: pgtype.Text{
			String: news.ImageUrl,
			Valid:  true,
		},
		PublishAt: pgtype.Timestamp{
			Time:  news.PublishAt,
			Valid: true,
		},
	})
	return err
}

func (u *NewsAdapter) GetNewsByID(newsID string, userID string) (news *db.News, isLiked bool, isDisliked bool, err error) {
	query := db.New(u.pool)
	_news, err := query.GetNews(context.Background(), pgtype.UUID{
		Bytes: uuid.MustParse(newsID),
		Valid: true,
	})
	if err != nil {
		return nil, false, false, err
	}
	news = &_news
	_, err = query.GetLike(context.Background(), db.GetLikeParams{
		NewsID: pgtype.UUID{
			Bytes: uuid.MustParse(newsID),
			Valid: true,
		},
		UserID: pgtype.UUID{
			Bytes: uuid.MustParse(userID),
			Valid: true,
		},
	})
	if err == nil {
		isLiked = true
	}
	_, err = query.GetDislike(context.Background(), db.GetDislikeParams{
		NewsID: pgtype.UUID{
			Bytes: uuid.MustParse(newsID),
			Valid: true,
		},
		UserID: pgtype.UUID{
			Bytes: uuid.MustParse(userID),
			Valid: true,
		},
	})
	if err == nil {
		isDisliked = true
	}
	return
}
