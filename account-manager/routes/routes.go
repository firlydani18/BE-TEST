package routes

import (
	"account-manager/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	r.POST("/register", controllers.Register(db))
	r.POST("/login", controllers.Login(db))
}
