package routes

import (
	"example/clean-arch/internal/delivery/http/handler"
	"example/clean-arch/internal/usecase"

	"github.com/gin-gonic/gin"
)

type RoutesConfig struct {
	Router *gin.Engine
	UserUC usecase.UserUsecase
	AuthUC usecase.AuthUsecase
	// Add more usecases as needed
}

func RegisterRoutes(cfg RoutesConfig) {
	// Group API version

	// Register user routes
	handler.NewUserHandler(cfg.Router, cfg.UserUC)

	// Register auth routes
	handler.NewAuthHandler(cfg.Router, cfg.AuthUC)
}
