package routes

import (
    "github.com/gin-gonic/gin"
    "github.com/toanmars/crud-generator/handlers"
    "github.com/toanmars/crud-generator/services"
    "github.com/toanmars/crud-generator/repositories"
    "gorm.io/gorm"
)

func SetupUserRoutes(r *gin.Engine, db *gorm.DB) {
    repo := repositories.NewUserRepository(db)
    service := services.NewUserService(repo)
    handler := handlers.NewUserHandler(service)

    r.GET("/user/:id", handler.GetByID)
    r.POST("/user", handler.Create)
    r.PUT("/user/:id", handler.Update)
    r.DELETE("/user/:id", handler.Delete)
}