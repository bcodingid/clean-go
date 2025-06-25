package handler

import (
	"example/clean-arch/internal/entity"
	"example/clean-arch/internal/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	uc usecase.AuthUsecase
}

func NewAuthHandler(r *gin.Engine, uc usecase.AuthUsecase) {
	h := &AuthHandler{uc}

	r.POST("/register", h.Register)
	r.POST("/login", h.Login)
}

// Register handler
func (h *AuthHandler) Register(c *gin.Context) {
	params := entity.RegisterParams{}

	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	err := h.uc.Register(params)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "registered"})
}

// Login Handler
func (h *AuthHandler) Login(c *gin.Context) {
	params := entity.LoginParams{}

	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	token, err := h.uc.Login(params)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
