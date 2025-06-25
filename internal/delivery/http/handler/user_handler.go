package handler

import (
	"example/clean-arch/internal/usecase"

	"github.com/gin-gonic/gin"
)

type UserHadler struct {
	uc usecase.UserUsecase
}

func NewUserHandler(g *gin.Engine, uc usecase.UserUsecase) {
	h := &UserHadler{uc}

	g.GET("/users", h.All)
}

func (h *UserHadler) All(c *gin.Context) {
	users, err := h.uc.All()
	if err != nil {
		c.JSON(500, gin.H{"error": "failed to fetch users"})
		return
	}

	c.JSON(200, users)
}
