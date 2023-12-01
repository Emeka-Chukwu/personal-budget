package cmd

import (
	"database/sql"
	"fmt"
	"net/http"
	"personal-budget/token"
	"personal-budget/util"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type Server struct {
	config     util.Config
	conn       *sql.DB
	router     *gin.Engine
	tokenMaker token.Maker
}

//// server serves out http request for our backend service

func NewServer(config util.Config, conn *sql.DB) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	server := &Server{tokenMaker: tokenMaker, config: config, conn: conn}

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

	router.POST("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("Personal Budget app ruuning at %s", server.config.HTTPServerAddress),
		})
	})

	// router.Use(authMiddleware(server.tokenMaker))

	// authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))
	// authRoutes.POST("/accounts", server.createAccount)

	server.router = router
}
