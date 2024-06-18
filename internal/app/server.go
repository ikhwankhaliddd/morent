package app

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/ikhwankhaleed/morent/config"
	"github.com/ikhwankhaleed/morent/internal/handlers"
	"github.com/ikhwankhaleed/morent/internal/repositories"
	"github.com/ikhwankhaleed/morent/internal/services"
	"github.com/ikhwankhaleed/morent/internal/utils"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"strings"
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
	authService := services.NewAuthService()
	userHandler := handlers.NewUserHandler(userService, authService)

	//routes here
	v1 := e.Group("/v1")
	{
		users := v1.Group("/users")
		{
			users.POST("/register", userHandler.RegisterUser)
			users.POST("/login", userHandler.Login)
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

func authMiddleware(authService services.AuthService, userService services.UserService) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ctx := c.Request().Context()
			authHeader := c.Request().Header.Get("Authorization")

			if !strings.Contains(authHeader, "Bearer") {
				response := utils.CreateResponse("Unauthorized", http.StatusUnauthorized, nil, nil)
				return c.JSON(http.StatusUnauthorized, response)
			}

			tokenString := ""
			arrayToken := strings.Split(authHeader, " ")
			if len(arrayToken) == 2 {
				tokenString = arrayToken[1]
			}

			token, err := authService.ValidateToken(tokenString)
			if err != nil {
				response := utils.CreateResponse("Unauthorized", http.StatusUnauthorized, nil, nil)
				return c.JSON(http.StatusUnauthorized, response)
			}

			payload, ok := token.Claims.(jwt.MapClaims)
			if !ok || !token.Valid {
				response := utils.CreateResponse("Unauthorized", http.StatusUnauthorized, nil, nil)
				return c.JSON(http.StatusUnauthorized, response)
			}

			userEmail := payload["user_email"].(string)
			user, err := userService.GetUserByEmail(ctx, userEmail)
			if err != nil {
				response := utils.CreateResponse("Unauthorized", http.StatusUnauthorized, nil, nil)
				return c.JSON(http.StatusUnauthorized, response)
			}

			c.Set("currentUser", user)
			return next(c)
		}
	}
}
