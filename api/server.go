package api

import (
	"net/http"

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

	router.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, struct {
			Status string `json:"test"`
		}{Status: "ok"})
	})
	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccounts)

	server.router = router
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
