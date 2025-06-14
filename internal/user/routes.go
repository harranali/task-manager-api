package user

import "net/http"

func RegisterRoutes(mux *http.ServeMux) {
	repo := NewRepository()
	srv := NewService(repo)
	h := NewHandler(srv)
	mux.HandleFunc("POST /login", h.Login)
	mux.HandleFunc("POST /register", h.Register)
	mux.HandleFunc("POST /logout", h.Logout)
}
