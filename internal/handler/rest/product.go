package rest

import (
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"preproj/internal/models"
	"strconv"
)

func (h *Handler) createProduct(c *gin.Context) {
	ctx := c.Request.Context()
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		slog.Error("failed to bing JSON", slog.Any("error", err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, err := h.services.Product.Create(ctx, &product)
	if err != nil {
		slog.Error("failed to create product", slog.Any("error", err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	product.ID = id
	c.JSON(http.StatusCreated, id)
}

func (h *Handler) getProductById(c *gin.Context) {
	ctx := c.Request.Context()
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		slog.Error("failed to parse request", slog.Any("error", err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product, err := h.services.Product.GetByID(ctx, id)
	if err != nil {
		slog.Error("failed to get user", slog.Any("error", err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if product == nil {
		slog.Debug("product is nil")
		c.JSON(http.StatusNotFound, gin.H{"error": "product not found"})
		return
	}
	c.JSON(http.StatusOK, product)
}

func (h *Handler) updateProduct(c *gin.Context) {
	ctx := c.Request.Context()
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		slog.Error("failed to parse request", slog.Any("error", err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var product models.Product

	if err := c.ShouldBindJSON(&product); err != nil {
		slog.Error("failed to bind JSON", slog.Any("error", err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product.ID = id

	updatedProductID, err := h.services.Product.Update(ctx, product)
	if err != nil {
		slog.Error("failed to update product", slog.Any("error", err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "product updated", "id": updatedProductID})
}

func (h *Handler) deleteProduct(c *gin.Context) {
	ctx := c.Request.Context()
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		slog.Error("failed to parse request", slog.Any("error", err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.services.Product.Delete(ctx, id)
	if err != nil {
		slog.Error("failed to delete product", slog.Any("error", err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "product deleted"})
}

func (h *Handler) getAllProducts(c *gin.Context) {
	ctx := c.Request.Context()
	products, err := h.services.Product.GetAll(ctx)
	if err != nil {
		slog.Error("failed to get all products", slog.Any("error", err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if len(products) == 0 {
		slog.Debug("products array is empty")
		c.JSON(http.StatusOK, []models.Product{})
		return
	}
	c.JSON(http.StatusOK, products)
}

func (h *Handler) getAllProductsByUserID(c *gin.Context) {
	ctx := c.Request.Context()
	userID, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		slog.Error("failed to parse request", slog.Any("error", err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	products, err := h.services.Product.GetAllByUserID(ctx, userID)
	if err != nil {
		slog.Error("failed to get all products by userID", slog.Any("error", err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if len(products) == 0 {
		slog.Debug("all products by userID array is empty")
		c.JSON(http.StatusOK, []models.Product{})
		return
	}
	c.JSON(http.StatusOK, products)
}
