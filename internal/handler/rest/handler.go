package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"preproj/internal/service"
)

type Handler struct {
	services *service.Service
	cache    *service.CacheService
}

func NewHandler(services *service.Service, cache *service.CacheService) *Handler {
	return &Handler{services: services, cache: cache}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.Use(gin.Recovery())
	router.Use(gin.Logger())

	api := router.Group("/api/v1")
	ttl := viper.GetDuration("ttl.HTTP")
	{
		users := api.Group("/users")
		{
			users.POST("/", h.createUser)
			users.GET("/:id", cacheMiddleware(h.cache, ttl), h.getUserById)
			users.PUT("/:id", h.updateUser)
			users.DELETE("/:id", h.deleteUser)
			users.GET("/", cacheMiddleware(h.cache, ttl), h.getAllUsers)

			products := users.Group("/products")
			{
				products.POST("/", h.createProduct)
				products.GET("/:id", cacheMiddleware(h.cache, ttl), h.getProductById)
				products.PUT("/:id", h.updateProduct)
				products.DELETE("/:id", h.deleteProduct)
				products.GET("/", cacheMiddleware(h.cache, ttl), h.getAllProducts)
				products.GET("/:userID", cacheMiddleware(h.cache, ttl), h.getAllProductsByUserID)
			}
		}
	}
	return router
}
