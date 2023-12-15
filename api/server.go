package api

import (
	"fmt"

	db "github.com/PfMartin/wegonice-api/db/sqlc"
	"github.com/PfMartin/wegonice-api/token"
	"github.com/PfMartin/wegonice-api/util"
	"github.com/gin-gonic/gin"
)

type Server struct {
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %v", err)

	}

	server := &Server{config: config, store: store, tokenMaker: tokenMaker}

	server.SetupRouter()

	return server, nil
}

func (server *Server) SetupRouter() {
	router := gin.Default()

	router.POST("/users", server.createUser)
	// router.POST("/users/login", server.loginUser)
	router.POST("/token/renew_access", server.renewAccessToken)

	server.router = router
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
