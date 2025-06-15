package user

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/harranali/task-manager-api/utils"
)

type Handler struct {
	srv Service
	v   *validator.Validate
}

func NewHandler(srv Service) *Handler {
	return &Handler{
		srv: srv,
		v:   validator.New(),
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
	var registerRequest RegisterRequest
	json.NewDecoder(r.Body).Decode(&registerRequest)
	err := h.v.Struct(registerRequest)
	var errors = make(map[string]string)
	if err != nil {
		vErrors := err.(validator.ValidationErrors)
		for _, fieldError := range vErrors {
			switch fieldError.Tag() {
			case "required":
				errors[fieldError.Field()] = fmt.Sprintf("the %v is required", fieldError.Field())
			case "email":
				errors[fieldError.Field()] = fmt.Sprintf("the %v must be valid email", fieldError.Field())
			}
		}
		utils.WriteErrorResponse(w, http.StatusUnprocessableEntity, errors)
		return
	}
	// check if user already exist
	_, err = h.srv.GetByEmail(registerRequest.Email)
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
