package server

import (
	"net/http"
	"restapi/control"

	"github.com/gorilla/mux"
)

func createHandler() (handler *mux.Router) {

	handler = mux.NewRouter()

	handler.HandleFunc("/", control.HelloWorld).Methods(http.MethodGet)
	handler.HandleFunc("/api/user/register", control.Auth).Methods(http.MethodPost)
	handler.HandleFunc("/api/user/login", control.Login).Methods(http.MethodPost)
	handler.HandleFunc("/api/posts", control.VerifiedHelloWorld).Methods(http.MethodGet)

	return
}
