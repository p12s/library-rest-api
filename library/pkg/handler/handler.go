package handler

import (
	"time"

	"github.com/gin-gonic/gin"
	cache "github.com/p12s/in-memory-cache"
	_ "github.com/p12s/library-rest-api/library/docs"
	"github.com/p12s/library-rest-api/library/pkg/service"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// Handler - struct contains service
type Handler struct {
	services *service.Service
	cache    *cache.Cache
}

const CACHE_CLEAN_EXPIRED_PERIOD = 1 * time.Minute

// NewHandler - constructor
func NewHandler(services *service.Service) *Handler {
	return &Handler{
		services: services,
		cache:    cache.New(CACHE_CLEAN_EXPIRED_PERIOD),
	}
}

// InitRoutes - routes
func (h *Handler) InitRoutes() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

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
