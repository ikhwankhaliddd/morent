package app

import (
	"github.com/ikhwankhaleed/morent/config"
	"github.com/ikhwankhaleed/morent/internal/handlers"
	"github.com/ikhwankhaleed/morent/internal/repositories"
	"github.com/ikhwankhaleed/morent/internal/services"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"log"
)

type Server struct {
	router *echo.Echo
	db     *sqlx.DB
}

func NewServer(cfg *config.Config) *Server {
	db, err := sqlx.Connect("postgres", cfg.DBSource)
	if err != nil {
		log.Fatalf("[APP] failed to connect with database : %v", err)
	}

	e := echo.New()

	//init here
	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	//routes here
	v1 := e.Group("/v1")
	{
		users := v1.Group("/users")
		{
			users.POST("/register", userHandler.RegisterUser)
		}
	}

	return &Server{
		router: e,
		db:     db,
	}
}

func (s *Server) Run() error {
	log.Println("Starting server on :8080")
	return s.router.Start(":8080")
}
