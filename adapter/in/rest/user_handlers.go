package rest

import (
	"context"
	"encoding/json"

	"net/http"
	"news-api/adapter/in/auth"
	inport "news-api/application/port/in"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type UserHandlers struct {
	userUseCase      inport.UsersUseCase
	recommendUseCase inport.RecommendUseCase
}

func NewUserHandlers(userUseCase inport.UsersUseCase, recommendUseCase inport.RecommendUseCase) *UserHandlers {
	return &UserHandlers{userUseCase: userUseCase, recommendUseCase: recommendUseCase}
}

func (u *UserHandlers) GetAll(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	usersList, err := u.userUseCase.GetAll()
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(APIResponse[any]{
			StatusCode: 500,
			Message:    "Unknown err",
		})

	}
	json.NewEncoder(response).Encode(APIResponse[[]*inport.User]{
		StatusCode: 200,
		Message:    "Ok",
		Data:       usersList,
	})
}

func (u *UserHandlers) Insert(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	var user inport.CreateUserPayload
	err := json.NewDecoder(request.Body).Decode(&user)
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(response).Encode(APIResponse[any]{
			StatusCode: 400,
			Message:    "Bad request"})
		return
	}
	err = u.userUseCase.Insert(&user)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(APIResponse[any]{
			StatusCode: 500,
			Message:    "Unknown err",
		})
		return
	}
	existUser, err := u.userUseCase.GetUserByAuthID(user.AuthID)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(APIResponse[any]{
			StatusCode: 500,
			Message:    "Unknown err",
		})
		return
	}
	if err = u.recommendUseCase.InsertUser(context.Background(), existUser.ID); err != nil {
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
		Data:       user,
	})
}

func (u *UserHandlers) Update(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	var user inport.UpdateUserPayload
	err := json.NewDecoder(request.Body).Decode(&user)
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(response).Encode(APIResponse[any]{
			StatusCode: 400,
			Message:    "Bad request",
		})
		return
	}
	id := chi.URLParam(request, "id")
	user.ID, err = uuid.Parse(id)
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(response).Encode(APIResponse[any]{
			StatusCode: 400,
			Message:    "Bad request"})
		return
	}
	err = u.userUseCase.Update(&user)
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

func (u *UserHandlers) Login(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	var user inport.CreateUserPayload
	err := json.NewDecoder(request.Body).Decode(&user)
	var accessToken string
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(response).Encode(APIResponse[any]{
			StatusCode: 400,
			Message:    "Bad request"})
		return
	}
	if existUser, _ := u.userUseCase.GetUserByAuthID(user.AuthID); existUser == nil {
		err = u.userUseCase.Insert(&user)
		if err != nil {
			response.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(response).Encode(APIResponse[any]{
				StatusCode: 500,
				Message:    "Unknown err",
			})
			return
		}
		existUser, err = u.userUseCase.GetUserByAuthID(user.AuthID)
		if err != nil {
			response.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(response).Encode(APIResponse[any]{
				StatusCode: 500,
				Message:    "Unknown err",
			})
			return
		}
		if err = u.recommendUseCase.InsertUser(context.Background(), existUser.ID); err != nil {
			response.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(response).Encode(APIResponse[any]{
				StatusCode: 500,
				Message:    "Unknown err",
			})
			return
		}
		accessToken, err = auth.GenerateJWT(existUser)
		if err != nil {
			response.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(response).Encode(APIResponse[any]{
				StatusCode: 500,
				Message:    "Unknown err",
			})
			return
		}
	} else {
		accessToken, err = auth.GenerateJWT(existUser)
		if err != nil {
			response.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(response).Encode(APIResponse[any]{
				StatusCode: 500,
				Message:    "Unknown err",
			})
			return
		}
	}
	json.NewEncoder(response).Encode(APIResponseLogin{
		StatusCode:  200,
		Message:     "Ok",
		AccessToken: accessToken,
	})
}

func (u *UserHandlers) AdminLogin(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	var user inport.AdminLoginPayload
	err := json.NewDecoder(request.Body).Decode(&user)
	var accessToken string
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(response).Encode(APIResponse[any]{
			StatusCode: 400,
			Message:    "Bad request"})
		return
	}
	existUser, err := u.userUseCase.GetAdmin(user.Email, user.Password)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(APIResponse[any]{
			StatusCode: 500,
			Message:    "Unknown err",
		})
		return
	}
	if existUser == nil {
		response.WriteHeader(400)
		json.NewEncoder(response).Encode(APIResponse[any]{
			StatusCode: 400,
			Message:    "Bad request",
		})
		return
	}
	accessToken, err = auth.GenerateJWT(existUser)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(APIResponse[any]{
			StatusCode: 500,
			Message:    "Unknown err",
		})
		return
	}

	json.NewEncoder(response).Encode(APIResponseLogin{
		StatusCode:  200,
		Message:     "Ok",
		AccessToken: accessToken,
	})
}

func (u *UserHandlers) Like(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	newsId := chi.URLParam(request, "newsId")
	user := request.Context().Value("user").(inport.UpdateUserPayload)
	err := u.userUseCase.Like(&inport.Like{
		UserId: user.ID.String(),
		NewsId: newsId,
	})
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(APIResponse[any]{
			StatusCode: 500,
			Message:    "Unknown err",
		})
		return
	}
	json.NewEncoder(response).Encode(APIResponse[any]{
		StatusCode: 201,
		Message:    "Created",
	})
}

func (u *UserHandlers) Unlike(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	newsId := chi.URLParam(request, "newsId")
	user := request.Context().Value("user").(inport.UpdateUserPayload)
	err := u.userUseCase.Unlike(&inport.Like{
		UserId: user.ID.String(),
		NewsId: newsId,
	})
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

func (u *UserHandlers) Dislike(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	newsId := chi.URLParam(request, "newsId")
	user := request.Context().Value("user").(inport.UpdateUserPayload)
	err := u.userUseCase.DisLike(&inport.Like{
		UserId: user.ID.String(),
		NewsId: newsId,
	})
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

func (u *UserHandlers) UnDislike(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	newsId := chi.URLParam(request, "newsId")
	user := request.Context().Value("user").(inport.UpdateUserPayload)
	err := u.userUseCase.UnDisLike(&inport.Like{
		UserId: user.ID.String(),
		NewsId: newsId,
	})
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
