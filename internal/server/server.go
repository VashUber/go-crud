package server

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	router *mux.Router
}

type User struct {
	Name string `json:"name"`
	Age  uint8  `json:"age"`
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
	s.router.HandleFunc("/", s.getUser).Methods("GET")
}

func (s *Server) getUser(w http.ResponseWriter, r *http.Request) {
	json, _ := json.Marshal(User{
		Name: "Jmishenko",
		Age:  54,
	})
	w.Write(json)
}
