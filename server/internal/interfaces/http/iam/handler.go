package iam

import (
	"encoding/json"
	"net/http"

	iamService "github.com/shojib116/auditflow-api/internal/application/iam"
	"github.com/shojib116/auditflow-api/internal/interfaces/http/middlewares"
	"github.com/shojib116/auditflow-api/internal/interfaces/http/utils"
)

type Handler struct {
	service *iamService.UserService
}

func NewHandler(service *iamService.UserService) *Handler {
	return &Handler{service: service}
}

func (h *Handler) RegisterRoutes(mux *http.ServeMux, mngr *middlewares.Manager) {
	mux.Handle("POST /api/v1/auth/register", http.HandlerFunc(h.HandlerRegister))
}

func (h *Handler) HandlerRegister(w http.ResponseWriter, r *http.Request) {
	var req RegisterUserRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		appError := MapError(err)
		utils.HandleAndLogError(w, r, appError.StatusCode, appError.Message)
		return
	}

	user, err := h.service.RegisterUser(r.Context(), iamService.RegisterRequestInput{
		Email:    req.Email,
		Password: req.Password,
		FullName: req.FullName,
	})
	if err != nil {
		appError := MapError(err)
		utils.HandleAndLogError(w, r, appError.StatusCode, appError.Message)
		return
	}

	resp := RegisterUserResponse{
		ID:       user.ID.String(),
		Email:    string(user.Email),
		FullName: user.FullName,
	}

	utils.SendJSON(w, http.StatusCreated, resp)
}
