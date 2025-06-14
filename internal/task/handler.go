package task

import (
	"net/http"
	"strconv"

	"github.com/harranali/task-manager-api/internal/user"
	"github.com/harranali/task-manager-api/utils"
)

type Handler struct {
	srv     Service
	userSrv user.Service
}

func NewHandler(srv Service, userSrv user.Service) *Handler {
	return &Handler{
		srv:     srv,
		userSrv: userSrv,
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
	// TODO validate
	if r.FormValue("title") == "" {
		utils.WriteErrorResponse(w, http.StatusBadRequest, "invalid task title")
		return
	}
	var createTaskRequest = CreateTaskRequest{
		Title: r.FormValue("title"),
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

// TODO implement Delete(id uint) error
