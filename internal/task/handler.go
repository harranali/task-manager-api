package task

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/harranali/task-manager-api/internal/user"
	"github.com/harranali/task-manager-api/utils"
)

type Handler struct {
	srv     Service
	userSrv user.Service
	v       *validator.Validate
}

func NewHandler(srv Service, userSrv user.Service) *Handler {
	return &Handler{
		srv:     srv,
		userSrv: userSrv,
		v:       validator.New(),
	}
}

func (h *Handler) SaveTask(w http.ResponseWriter, r *http.Request) {
	// get user token
	token, err := h.userSrv.GetUserToken(r)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusUnauthorized, "unauthorized request")
		return
	}
	// get user id
	user, err := h.userSrv.GetUserByToken(token)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusUnauthorized, "unauthorized request")
		return
	}

	var createTaskRequest CreateTaskRequest
	err = json.NewDecoder(r.Body).Decode(&createTaskRequest)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusBadRequest, "invalid request")
		return
	}

	err = h.v.Struct(createTaskRequest)
	if err != nil {
		var errors = make(map[string]string)
		vErrors := err.(validator.ValidationErrors)
		for _, fieldError := range vErrors {
			switch fieldError.Tag() {
			case "required":
				errors[fieldError.Field()] = fmt.Sprintf("the %v is required", fieldError.Field())
			}
		}
		utils.WriteErrorResponse(w, http.StatusUnprocessableEntity, errors)
		return
	}

	task, err := h.srv.Save(createTaskRequest, user.ID)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusInternalServerError, "something went wrong")
		return
	}
	utils.WriteSuccessResponse(w, http.StatusOK, task)
}

func (h *Handler) GetByID(w http.ResponseWriter, r *http.Request) {
	token, err := h.userSrv.GetUserToken(r)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusUnauthorized, "unauthorized request")
		return
	}
	_, err = h.userSrv.GetUserByToken(token)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusUnauthorized, "unauthorized request")
		return
	}
	// get task id from request
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusNotFound, "not found")
		return
	}
	// get task by id from service
	task, err := h.srv.GetById(uint(id))
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusNotFound, "not found")
		return
	}
	// return response
	utils.WriteSuccessResponse(w, http.StatusOK, task)
}

func (h *Handler) GetUserTasks(w http.ResponseWriter, r *http.Request) {
	// get the token
	token, err := h.userSrv.GetUserToken(r)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusUnauthorized, "unauthorized request")
		return
	}
	user, err := h.userSrv.GetUserByToken(token)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusUnauthorized, "unauthorized request")
		return
	}
	tasks, err := h.srv.GetUserTasks(user.ID)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusInternalServerError, "something went wrong")
		return
	}
	utils.WriteSuccessResponse(w, http.StatusOK, tasks)
}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	// authorize
	token, err := h.userSrv.GetUserToken(r)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusUnauthorized, "unauthorized request")
		return
	}
	_, err = h.userSrv.GetUserByToken(token)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusUnauthorized, "unauthorized request")
		return
	}
	// validate input
	isDoneStr := r.FormValue("is_done")
	isDone, err := strconv.ParseBool(isDoneStr)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusBadRequest, "invalid request")
		return
	}
	updateTaskRequest := UpdateTaskRequest{
		Title:  r.FormValue("title"),
		IsDone: isDone,
	}
	// get task
	taskIdStr, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusInternalServerError, "something went wrong")
		return
	}
	// update task
	task, err := h.srv.Update(updateTaskRequest, uint(taskIdStr))
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusInternalServerError, "something went wrong")
		return
	}

	utils.WriteSuccessResponse(w, http.StatusOK, task)
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	//authorize
	token, err := h.userSrv.GetUserToken(r)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusUnauthorized, "unauthorized request")
		return
	}
	_, err = h.userSrv.GetUserByToken(token)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusUnauthorized, "unauthorized request")
		return
	}
	taskID, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusNotFound, "not found")
	}
	err = h.srv.Delete(uint(taskID))
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusInternalServerError, "something went wrong")
		return
	}
	utils.WriteSuccessResponse(w, http.StatusOK, nil)
}
