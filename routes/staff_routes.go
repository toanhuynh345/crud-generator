package routes

import (
    "github.com/gin-gonic/gin"
    "github.com/toanmars/crud-generator/handlers"
    "github.com/toanmars/crud-generator/services"
    "github.com/toanmars/crud-generator/repositories"
    "gorm.io/gorm"
)

func SetupStaffRoutes(r *gin.Engine, db *gorm.DB) {
    repo := repositories.NewStaffRepository(db)
    service := services.NewStaffService(repo)
    handler := handlers.NewStaffHandler(service)

    r.GET("/staff/:id", handler.GetByID)
    r.POST("/staff", handler.Create)
    r.PUT("/staff/:id", handler.Update)
    r.DELETE("/staff/:id", handler.Delete)
}