package user

import (
	"fmt"
	"net/http"

	"github.com/harranali/task-manager-api/utils"
)

type Handler struct {
	srv Service
}

func NewHandler(srv Service) *Handler {
	return &Handler{
		srv: srv,
	}
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	// TODO validatae
	loginRequest := LoginRequest{
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
	}
	// get user from the database
	user, err := h.srv.GetByEmail(loginRequest.Email)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusBadRequest, "invalid email or password")
		return
	}
	// verify the password
	if err := h.srv.VerifyUserPassword(user, loginRequest); err != nil {
		utils.WriteErrorResponse(w, http.StatusBadRequest, "invalid email or password")
		return
	}

	// generate token
	token, err := h.srv.GenerateToken(user)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusInternalServerError, "something went wrong")
		return
	}
	// return token
	utils.WriteSuccessResponse(w, http.StatusOK, TokenResponse{
		Token: token,
	})
}

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	// TODO validate
	registerRequest := RegisterRequest{
		Name:     r.FormValue("name"),
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
	}
	// check if user already exist
	_, err := h.srv.GetByEmail(registerRequest.Email)
	if err == nil {
		utils.WriteErrorResponse(w, http.StatusBadRequest, "user already registered")
		return
	}
	// register
	user, err := h.srv.Register(registerRequest)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.WriteSuccessResponse(w, http.StatusOK, user)
}

func (h *Handler) Logout(w http.ResponseWriter, r *http.Request) {
	// TODO handle logout
	fmt.Println("handle logout")
}
