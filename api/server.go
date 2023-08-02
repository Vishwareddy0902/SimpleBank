package api

import (
	"fmt"

	db "github.com/Vishwareddy0902/Simple-Bank/db/sqlc"
	"github.com/Vishwareddy0902/Simple-Bank/token"
	"github.com/Vishwareddy0902/Simple-Bank/util"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type Server struct {
	config     util.Config
	store      *db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

func NewServer(config util.Config, store *db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TOKEN_SYMMETRIC_KEY)
	if err != nil {
		return nil, fmt.Errorf("connot create token maker: %w", err)
	}
	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	router := gin.Default()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}

	router.POST("/users", server.createUser)
	router.POST("/users/login", server.loginUser)
	authRouter := router.Group("/").Use(authMiddleware(tokenMaker))
	authRouter.POST("/accounts", server.createAccount)
	authRouter.GET("/accounts/:id", server.getAccount)
	authRouter.GET("/accounts", server.listAccounts)
	authRouter.DELETE("/accounts/:id", server.deleteAccount)
	authRouter.POST("/transfers", server.createTransfer)

	server.router = router

	return server, nil

}

func (server *Server) StartServer(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
