package api

import (
	db "backend_masterclass/db/sqlc"

	"github.com/gin-gonic/gin"
	
	"github.com/go-playground/validator/v10"

	"github.com/gin-gonic/gin/binding"
)

// Server serves HTTP requests for our banking service
type Server struct{
	store db.Store
	router *gin.Engine
}

// NewServer creates a new HTTP server and setup routing.
func NewServer(store db.Store) *Server{
	server := &Server{store: store}
	router := gin.Default()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}

	router.POST("/users", server.createUser)

	router.POST("/account", server.createAccount)
	router.GET("/account/:id", server.getAccount)
	router.GET("/accounts", server.listAccounts)

	router.POST("/transfers", server.createTransfer)
	
	server.router = router
	return server
}

// Start runs the HTTP server on a specific address
func (server *Server) Start(address string) error{
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error":err.Error()}
}
