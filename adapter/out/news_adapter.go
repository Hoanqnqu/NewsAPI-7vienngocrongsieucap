package outAdapter

import (
	"context"
	"encoding/json"
	outport "news-api/application/port/out"
	db "news-api/internal/db"

	"github.com/google/uuid"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

type NewsAdapter struct {
	pool *pgxpool.Pool
}

func NewNewsAdapter(pool *pgxpool.Pool) *NewsAdapter {
	return &NewsAdapter{pool: pool}
}

func (u *NewsAdapter) GetAll() ([]outport.NewsWithCategory, error) {
	query := db.New(u.pool)
	news, err := query.GetAllNews(context.Background())
	sl := make([]outport.NewsWithCategory, len(news))
	if err != nil {
		return nil, err
	}
	var category_ids []pgtype.UUID
	for i, v := range news {
		sl[i].Author = v.Author
		sl[i].Content = v.Content
		sl[i].Description = v.Description
		sl[i].Title = v.Title
		sl[i].Url = v.Url
		sl[i].ImageUrl = v.ImageUrl
		sl[i].PublishAt = v.PublishAt
		sl[i].ID = v.ID
		err = json.Unmarshal(v.CategoryIds, &category_ids)
		if err != nil {
			return nil, err
		}
		sl[i].Categories = category_ids
	}
	return sl, nil
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

func (u *NewsAdapter) GetNewsByID(newsID string, userID string) (news *outport.NewsWithCategory, isLiked bool, isDisliked bool, err error) {
	query := db.New(u.pool)
	_news, err := query.GetNews(context.Background(), pgtype.UUID{
		Bytes: uuid.MustParse(newsID),
		Valid: true,
	})
	if err != nil {
		return nil, false, false, err
	}
	var category_ids []pgtype.UUID

	news.Author = _news.Author
	news.Content = _news.Content
	news.Description = _news.Description
	news.Title = _news.Title
	news.Url = _news.Url
	news.ImageUrl = _news.ImageUrl
	news.PublishAt = _news.PublishAt
	news.ID = _news.ID

	err = json.Unmarshal(_news.CategoryIds, &category_ids)
	if err != nil {
		return nil, false, false, err
	}
	news.Categories = category_ids

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
