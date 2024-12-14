package handler

import (
	"github.com/gin-gonic/gin"
	"preproj/internal/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		users := api.Group("/users")
		{
			users.POST(`/api/v1/users`, h.createUser)
			users.GET(`/api/v1/users/{id}`, h.getUserById)
			users.PUT(`/api/v1/users/{id}`, h.updateUser)
			users.DELETE(`/api/v1/users/{id}`, h.deleteUser)
			users.GET(`/api/v1/users`, h.getAllUsers)

			products := users.Group("/products")
			{
				products.POST(`/api/v1/products`, h.createProduct)
				products.GET(`/api/v1/products/{id}`, h.getProductById)
				products.PUT(`/api/v1/products/{id}`, h.updateProduct)
				products.DELETE(`/api/v1/products/{id}`, h.deleteProduct)
				products.GET(`/api/v1/products`, h.getAllProducts)
			}
		}
	}
	return router
}
