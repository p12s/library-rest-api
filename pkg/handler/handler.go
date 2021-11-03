package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/p12s/library-rest-api/pkg/service"
)

// Handler - struct containes service
type Handler struct {
	services *service.Service
}

// NewHandler - constructor
func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

// InitRoutes - routes
func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("/health", h.health)

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api", h.userIdentity)
	{
		book := api.Group("/book")
		{
			book.POST("/", h.createBook)
			book.GET("/", h.getAllBooks)
			book.GET("/:id", h.getBookById)
			book.PUT("/:id", h.updateBook)
			book.DELETE("/:id", h.deleteBook)
		}
	}
	return router
}
