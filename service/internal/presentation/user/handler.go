package user

import (
	"context"
	"net/http"
	"time"

	"github.com/Seiya-Tagami/pecopeco-service/internal/presentation/responder"
	"github.com/Seiya-Tagami/pecopeco-service/internal/presentation/util/httputil"
	userUsecase "github.com/Seiya-Tagami/pecopeco-service/internal/usecase/user"
	"github.com/Seiya-Tagami/pecopeco-service/internal/util/jwt"
	"github.com/Seiya-Tagami/pecopeco-service/pkg/validator"
)

type handler struct {
	findUserUsecase *userUsecase.FindUserUseCase
	loginUsecase    *userUsecase.LoginUsecase
}

func NewHandler(
	findUserUsecase *userUsecase.FindUserUseCase,
	loginUsecase *userUsecase.LoginUsecase,
) handler {
	return handler{
		findUserUsecase: findUserUsecase,
		loginUsecase:    loginUsecase,
	}
}

func (h *handler) Login(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	ctx, cancel := context.WithTimeout(ctx, time.Second*30)
	defer cancel()

	params := LoginParams{}
	if err := httputil.ParseJSONRequestBody(r, &params); err != nil {
		responder.ReturnStatusInternalServerError(w, err)
		return
	}

	validate := validator.GetValidator()
	if err := validate.Struct(&params); err != nil {
		responder.ReturnStatusBadRequest(w, err)
		return
	}

	dto := userUsecase.LoginUseCaseDto{
		ID:    params.ID,
		Name:  params.Name,
		Email: params.Email,
	}
	ud, err := h.loginUsecase.Run(ctx, dto)
	if err != nil {
		responder.ReturnStatusInternalServerError(w, err)
		return
	}

	response := LoginResponse{
		ID:    ud.ID,
		Name:  ud.Name,
		Email: ud.Email,
	}

	// jwtの生成
	accessToken, err := jwt.Generate(ud.ID)
	if err != nil {
		responder.ReturnStatusInternalServerError(w, err)
		return
	}
	jwt.SetHttpHeader(w, accessToken)

	responder.ReturnStatusOK(w, response)
}

func (h *handler) FindUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	ctx, cancel := context.WithTimeout(ctx, time.Second*30)
	defer cancel()

	accessToken := r.Header.Get("Authorization")
	userID, err := jwt.UserIDFromToken(accessToken)
	if err != nil {
		responder.ReturnStatusUnauthorized(w, err)
		return
	}

	dto := userUsecase.FindUserUseCaseDto{
		ID: userID,
	}
	ud, err := h.findUserUsecase.Run(ctx, dto)
	if err != nil {
		responder.ReturnStatusInternalServerError(w, err)
		return
	}

	response := LoginResponse{
		ID:    ud.ID,
		Name:  ud.Name,
		Email: ud.Email,
	}
	responder.ReturnStatusOK(w, response)
}
