package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
	inport "news-api/application/port/in"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type NewsHandlers struct {
	newsUseCase inport.NewsUseCase
}

func NewNewsHandlers(newsUseCase inport.NewsUseCase) *NewsHandlers {
	return &NewsHandlers{newsUseCase: newsUseCase}
}

func (u *NewsHandlers) GetAll(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	newsList, err := u.newsUseCase.GetAll()
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(APIResponse[any]{
			StatusCode: 500,
			Message:    "Unknown err",
		})
	}
	json.NewEncoder(response).Encode(APIResponse[[]*inport.News]{
		StatusCode: 200,
		Message:    "Ok",
		Data:       newsList,
	})
}

func (u *NewsHandlers) Insert(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	var news inport.CreateNewsPayload
	err := json.NewDecoder(request.Body).Decode(&news)
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(response).Encode(APIResponse[any]{
			StatusCode: 400,
			Message:    "Bad request"})
		return
	}
	err = u.newsUseCase.Insert(&news)
	fmt.Println(err)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(APIResponse[any]{
			StatusCode: 500,
			Message:    "Unknown err",
		})
		return
	}
	json.NewEncoder(response).Encode(APIResponse[any]{
		StatusCode: 200,
		Message:    "Ok",
	})
}

func (u *NewsHandlers) Update(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	var news inport.UpdateNewsPayload
	err := json.NewDecoder(request.Body).Decode(&news)
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(response).Encode(APIResponse[any]{
			StatusCode: 400,
			Message:    "Bad request",
		})
		return
	}
	id := chi.URLParam(request, "id")
	news.ID, err = uuid.Parse(id)
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(response).Encode(APIResponse[any]{
			StatusCode: 400,
			Message:    "Bad request"})
		return
	}
	err = u.newsUseCase.Update(&news)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(APIResponse[any]{
			StatusCode: 500,
			Message:    "Unknown err",
		})
		return
	}
	json.NewEncoder(response).Encode(APIResponse[any]{
		StatusCode: 200,
		Message:    "Ok",
	})
}
func (u *NewsHandlers) GetNewsByID(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	newsId := chi.URLParam(request, "newsId")
	user := request.Context().Value("user").(inport.UpdateUserPayload)
	news, err := u.newsUseCase.GetNewsByID(newsId, user.ID.String())
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(response).Encode(APIResponse[any]{
			StatusCode: 404,
			Message:    "Not Found"})
		return
	}

	json.NewEncoder(response).Encode(APIResponse[*inport.News]{
		StatusCode: 200,
		Message:    "Ok",
		Data:       news,
	})
}
