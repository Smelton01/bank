package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/smelton01/bank/db/sqlc"
)

// Server serves HTTP requests for our service.
type Server struct {
	store  *db.Store
	router *gin.Engine
}

// NewServer creates a new HTTP server and sets up routing.
func NewServer(store *db.Store) *Server {
	server := &Server{store: store}

	server.setupRouter()

	return server
}

func (server *Server) setupRouter() {
	router := gin.Default()

	router.POST("/accounts", server.createAccount)

	server.router = router
}
