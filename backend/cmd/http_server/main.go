package main

import (
	"context"
	"log"

	"github.com/Krab1o/taskboard/internal/api"
	"github.com/Krab1o/taskboard/internal/repository"
	"github.com/Krab1o/taskboard/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	ctx := context.Background()
	err := LoadEnv()
	if err != nil {
		log.Fatal(err)
	}

	dbConn, err := DSN()
	if err != nil {
		log.Fatal(err)
	}

	pool, err := pgxpool.New(ctx, dbConn)
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()

	userRepo := repository.NewUserRepository(pool)
	projectRepo := repository.NewProjectRepository(pool)
	taskRepo := repository.NewTaskRepository(pool)

	userService := service.NewUserService(userRepo)
	projectService := service.NewProjectService(projectRepo)
	taskService := service.NewTaskService(taskRepo)

	handler := api.NewHandler(
		userService,
		projectService,
		taskService,
	)

	s := gin.Default()
	setupRoutes(s, handler)
	err = s.Run()
	if err != nil {
		log.Fatal(err)
	}
}
