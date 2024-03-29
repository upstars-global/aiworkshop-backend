package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "todo/docs"
	"todo/internal/app/todo"
)

// @title Task List App API
// @description This is a sample Task List app server.
// @version 1.0
// @host b6f6001008cb.ngrok.app
// @BasePath /
func main() {
	router := gin.Default()
	router.Use(CORSMiddleware())
	repo := todo.NewInMemoryRepository()
	service := todo.NewTaskService(repo)
	handler := todo.NewTaskHandler(service)

	handler.RegisterRoutes(router)
	url := ginSwagger.URL("/swagger/doc.json") // Адреса документації

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	router.Run(":8085")
}
