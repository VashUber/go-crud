package server

import (
	"encoding/json"
	"io"
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

var users = make(map[string]User)

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
	s.router.HandleFunc("/api/user", s.getUser).Methods("GET")
	s.router.HandleFunc("/api/user", s.createUser).Methods("POST")
}

func (s *Server) createUser(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var user User
	json.Unmarshal(body, &user)
	if _, inc := users[user.Name]; inc {
		w.WriteHeader(http.StatusConflict)
		return
	}
	users[user.Name] = user
}

func (s *Server) getUser(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	if u, ok := users[name]; ok {
		user, _ := json.Marshal(u)
		w.Write(user)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}
