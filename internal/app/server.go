package app

import (
	"github.com/gorilla/mux"
	"github.com/ikhwankhaleed/morent/config"
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"
)

type Server struct {
	router *mux.Router
	db     *sqlx.DB
}

func NewServer(cfg *config.Config) *Server {
	db, err := sqlx.Connect("postgres", cfg.DBSource)
	if err != nil {
		log.Fatalf("[APP] failed to connect with database : %v", err)
	}

	router := mux.NewRouter()

	//init here

	//routes here
	router.HandleFunc("/register").Methods("POST")

	return &Server{
		router: router,
		db:     db,
	}
}

func (s *Server) Run() error {
	log.Println("Starting server on :8080")
	return http.ListenAndServe(":8080", s.router)
}
