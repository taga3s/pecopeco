package user

import (
	"context"
	"net/http"
	"time"

	"github.com/ayanami77/pecopeco-service/internal/presentation/responder"
	"github.com/ayanami77/pecopeco-service/internal/presentation/util/httputil"
	userUsecase "github.com/ayanami77/pecopeco-service/internal/usecase/user"
	"github.com/ayanami77/pecopeco-service/internal/util/jwt"
	"github.com/ayanami77/pecopeco-service/pkg/validator"
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

	validate := validator.Get()
	if err := validate.Struct(&params); err != nil {
		responder.ReturnStatusBadRequest(w, err)
		return
	}

	inputDto := userUsecase.LoginUseCaseInputDto{
		ID:    params.ID,
		Name:  params.Name,
		Email: params.Email,
	}
	outputDto, err := h.loginUsecase.Run(ctx, inputDto)
	if err != nil {
		responder.ReturnStatusInternalServerError(w, err)
		return
	}

	response := LoginResponse{
		ID:    outputDto.ID,
		Name:  outputDto.Name,
		Email: outputDto.Email,
	}

	// jwtの生成
	accessToken, err := jwt.Generate(outputDto.ID)
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
	userID, err := jwt.GetUserIDFromToken(accessToken)
	if err != nil {
		responder.ReturnStatusUnauthorized(w, err)
		return
	}

	dto, err := h.findUserUsecase.Run(ctx, userID)
	if err != nil {
		responder.ReturnStatusInternalServerError(w, err)
		return
	}

	response := LoginResponse{
		ID:    dto.ID,
		Name:  dto.Name,
		Email: dto.Email,
	}
	responder.ReturnStatusOK(w, response)
}
