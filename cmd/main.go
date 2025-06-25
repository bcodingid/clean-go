package main

// main.go - Entry point for the Clean Architecture example application

import (
	"log"
	"os"

	"example/clean-arch/config"
	"example/clean-arch/internal/infrastructure/pgsql"
	"example/clean-arch/internal/usecase"
	"example/clean-arch/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Println("Error loading .env file")
	}

	// Initialize PostgreSQL database connection
	pgdb, err := pgsql.InitDB()
	if err != nil {
		panic(err)
	}

	pgUserRepo := pgsql.NewUserRepo(pgdb)
	authUsecase := usecase.NewAuthUsecase(pgUserRepo, os.Getenv(config.JwtKey))
	userUsecase := usecase.NewUserUsecase(pgUserRepo)

	// Initialize Gin router
	// register routes
	r := gin.Default()
	routes.RegisterRoutes(routes.RoutesConfig{
		Router: r,
		UserUC: userUsecase,
		AuthUC: authUsecase,
	})

	log.Println("Server running at :8080")
	r.Run(":8080")
}
