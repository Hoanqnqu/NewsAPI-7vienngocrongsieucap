package service

import (
	"context"
	inport "news-api/application/port/in"
	outport "news-api/application/port/out"

	"github.com/google/uuid"
)

type NewsService struct {
	newsPort             outport.NewsPort
	recommendationSystem outport.RecommendationSystem
}

func NewNewsService(newsPort outport.NewsPort, recommendationSystem outport.RecommendationSystem) *NewsService {
	return &NewsService{newsPort: newsPort, recommendationSystem: recommendationSystem}
}
func (g *NewsService) GetAll() ([]*inport.News, error) {
	newsList, err := g.newsPort.GetAll()
	if err != nil {
		return nil, err
	}
	return func() ([]*inport.News, error) {
		result := make([]*inport.News, len(newsList))
		for i, v := range newsList {
			result[i] = MapNews(v)
		}
		return result, nil
	}()
}

func (g *NewsService) Insert(news *inport.CreateNewsPayload) error {
	id := uuid.New()
	err := g.newsPort.Insert(outport.News{
		ID:          id,
		Title:       news.Title,
		Content:     news.Content,
		Description: news.Description,
		Author:      news.Author,
		Url:         news.URL,
		ImageUrl:    news.ImageURL,
		PublishAt:   news.PublishAt,
		Categories:  news.Categories,
	})

	if err == nil {
		return g.recommendationSystem.InsertNews(context.Background(), id, news.Categories)
	}
	return nil
}

func (g *NewsService) GetNewsByID(newsID, userID string) (*inport.News, error) {
	news, isLiked, isDislike, err := g.newsPort.GetNewsByID(newsID, userID)
	if err != nil {
		return nil, err
	}
	convertedNews := MapNews(*news)
	convertedNews.IsLiked = isLiked
	convertedNews.IsDisliked = isDislike
	return convertedNews, nil
}

func (g *NewsService) Update(news *inport.UpdateNewsPayload) error {
	err := g.newsPort.Update(outport.News{
		ID:          news.ID,
		Title:       news.Title,
		Content:     news.Content,
		Description: news.Description,
		Author:      news.Author,
		Url:         news.URL,
		ImageUrl:    news.ImageURL,
		PublishAt:   news.PublishAt,
		Categories:  news.Categories,
	})
	if err == nil {
		return g.recommendationSystem.InsertNews(context.Background(), news.ID, news.Categories)
	}
	return err
}

func (g *NewsService) SearchNews(keyword string) ([]*inport.News, error) {
	newsList, err := g.newsPort.SearchNews(keyword)
	if err != nil {
		return nil, err
	}
	return func() ([]*inport.News, error) {
		result := make([]*inport.News, len(newsList))
		for i, v := range newsList {
			result[i] = MapNews(v)
		}
		return result, nil
	}()
}
