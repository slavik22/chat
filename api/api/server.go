package api

import (
	"fmt"
	db "github.com/slavik22/chat/db/sqlc"
	"github.com/slavik22/chat/token"
	"github.com/slavik22/chat/util"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Server struct {
	config   util.Config
	store    db.Store
	router   *echo.Echo
	jwtMaker token.JWTMaker
}

func NewServer(config util.Config, store db.Store) (*Server, error) {

	jwt, err := token.NewJWTMaker(config.SecretKey)

	if err != nil {
		return nil, fmt.Errorf("Cannot create token maker %w", err)
	}

	server := &Server{
		store:    store,
		jwtMaker: *jwt,
		config:   config,
	}

	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	router := echo.New()

	v1 := router.Group("api/v1")

	auth := v1.Group("/auth")
	{
		auth.POST("/login", server.loginUser)
		auth.POST("/register", server.createUser)
	}

	users := v1.Group("/users", authMiddleware(server.jwtMaker))
	{
		users.GET("/", server.getUsers)
		users.GET("/:id", server.getUser)
		users.PUT("/", server.updateUser)

		friends := users.Group("/friends")
		{
			friends.GET("/", server.getFriends)
			friends.POST("/:id", server.addFriend)
			friends.DELETE("/:id", server.removeFriend)
		}
		blackList := users.Group("/blackList")
		{
			blackList.GET("/", server.getBlackList)
			blackList.POST("/:id", server.addBlackList)
			blackList.DELETE("/:id", server.removeBlackList)
		}
	}

	server.router = router
}

func (server *Server) Start(address string) error {
	return server.router.StartServer(&http.Server{Addr: address})
}
