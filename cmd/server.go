package cmd

import (
	"fmt"
	"personal-budget/token"
	"personal-budget/util"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type Server struct {
	config util.Config

	router     *gin.Engine
	tokenMaker token.Maker
}

//// server serves out http request for our backend service

func NewServer(config util.Config) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	server := &Server{tokenMaker: tokenMaker, config: config}

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}
	server.setupRouter()
	return server, nil
}

func errorrResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func (server *Server) setupRouter() {
	router := gin.Default()

	// router.POST("/users", server.createUser)

	// router.Use(authMiddleware(server.tokenMaker))

	// authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))
	// authRoutes.POST("/accounts", server.createAccount)

	server.router = router
}
