package main

import (
	"github.com/toanmars/crud-generator/routes"

	"github.com/gin-gonic/gin"
	"github.com/toanmars/crud-generator/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// @title CRUD Generator API
// @version 1.0
// @description Đây là một API được tạo tự động bởi CRUD Generator.
// @host localhost:8080
// @BasePath /
func main() {
	// Kết nối database
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	// Tạo bảng
	db.AutoMigrate(&models.User{})

	// Khởi tạo router
	r := gin.Default()

	// Thiết lập routes
	routes.SetupUserRoutes(r, db)

	// Chạy server
	r.Run(":8080")
}
