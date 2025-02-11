package rest

import (
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"preproj/internal/models"
	"strconv"
)

func (h *Handler) createUser(c *gin.Context) {
	ctx := c.Request.Context()
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		slog.Error("failed bind JSON", slog.Any("error", err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, err := h.services.User.Create(ctx, &user)
	if err != nil {
		slog.Error("failed to create user", slog.Any("error", err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	user.ID = id
	c.JSON(http.StatusCreated, user)
}

func (h *Handler) getUserById(c *gin.Context) {
	ctx := c.Request.Context()
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		slog.Error("failed parse request", slog.Any("error", err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.services.User.GetByID(ctx, id)
	if err != nil {
		slog.Error("failed to get user", slog.Any("error", err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *Handler) getAllUsers(c *gin.Context) {
	ctx := c.Request.Context()
	users, err := h.services.User.GetAll(ctx)
	if err != nil {
		slog.Error("failed to get all users", slog.Any("error", err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if len(users) == 0 {
		slog.Debug("users array is empty")
		c.JSON(http.StatusOK, []models.User{})
		return
	}
	c.JSON(http.StatusOK, users)
}

func (h *Handler) updateUser(c *gin.Context) {
	ctx := c.Request.Context()
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		slog.Error("failed to parse request", slog.Any("error", err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		slog.Error("failed to bind JSON", slog.Any("error", err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user.ID = id

	updatedUserID, err := h.services.User.Update(ctx, user)
	if err != nil {
		slog.Error("failed to update user", slog.Any("err", err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user updated successfully", "id": updatedUserID})
}

func (h *Handler) deleteUser(c *gin.Context) {
	ctx := c.Request.Context()
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		slog.Error("failed to parse request", slog.Any("error", err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.services.User.Delete(ctx, id)
	if err != nil {
		slog.Error("failed to delete user", slog.Any("error", err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user deleted successfully"})
}
