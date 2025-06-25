package main

// main.go - Entry point for the Clean Architecture example application

import (
	"log"
	"os"

	"example/clean-arch/internal/infrastructure/pgsql"
	"example/clean-arch/internal/usecase"
	"example/clean-arch/routes"

	"github.com/gin-gonic/gin"

	_ "github.com/lib/pq"
)

func main() {

	// Initialize PostgreSQL database connection
	db, err := pgsql.InitDB()
	if err != nil {
		panic(err)
	}

	pgUserRepo := pgsql.NewUserRepo(db)
	authUsecase := usecase.NewAuthUsecase(pgUserRepo, os.Getenv("JWT_SECRET"))
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
