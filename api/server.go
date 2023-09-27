package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	db "github.com/yuqihu1103/SimpleBank/db/sqlc"
)

// Server for http services of the simple bank
type Server struct {
	store  db.Store
	router *gin.Engine
}

// Creates a new http server and set up routing
func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}

	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccount)
	router.PUT("/accounts/:id", server.updateAccount)
	router.DELETE("/accounts/:id", server.deleteAccount)

	router.POST("/transfers", server.createTransfer)

	server.router = router
	return server
}

// Runs the http server on a specific address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

// genereate response corresponding to an error
func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
