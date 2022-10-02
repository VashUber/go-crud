package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	router *mux.Router
}

func New() *Server {
	s := &Server{
		router: mux.NewRouter(),
	}
	s.InitRoutes()
	return s
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *Server) InitRoutes() {
	s.router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		resp, err := json.Marshal("samara")
		if err != nil {
			fmt.Println(err)
		}
		w.Write(resp)
	}).Methods("GET")
}
