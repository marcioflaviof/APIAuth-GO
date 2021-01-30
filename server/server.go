package server

import (
	"log"
	"net/http"
)

func Start() {
	h := createHandler()

	s := createServer()

	s.Handler = h

	log.Fatal(s.ListenAndServe())
}

func createServer() (server *http.Server) {
	server = &http.Server{
		Addr: ":3000",
	}

	return
}
