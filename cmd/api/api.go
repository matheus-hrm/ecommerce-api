package api

import (
	"database/sql"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/matheus-hrm/trampos/service/user"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db: db,
	}
}

func (s *APIServer) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api").Subrouter()
	userHandler := user.NewHandler()
	userHandler.RegisterRoutes(subrouter)
	log.Println("Server is running on", s.addr)
	return http.ListenAndServe(s.addr, router)
}