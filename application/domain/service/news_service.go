package service

import (
	inport "news-api/application/port/in"
	outport "news-api/application/port/out"

	"github.com/google/uuid"
)

type NewsService struct {
	newsPort outport.NewsPort
}

func NewNewsService(newsPort outport.NewsPort) *NewsService {
	return &NewsService{newsPort: newsPort}
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
	return g.newsPort.Insert(outport.News{
		ID:          uuid.New(),
		Title:       news.Title,
		Content:     news.Content,
		Description: news.Description,
		Author:      news.Author,
		Url:         news.URL,
		ImageUrl:    news.ImageURL,
		PublishAt:   news.PublishAt,
	})
}

func (g *NewsService) Update(news *inport.UpdateNewsPayload) error {
	return g.newsPort.Update(outport.News{
		ID:          news.ID,
		Title:       news.Title,
		Content:     news.Content,
		Description: news.Description,
		Author:      news.Author,
		Url:         news.URL,
		ImageUrl:    news.ImageURL,
		PublishAt:   news.PublishAt,
	})
}

