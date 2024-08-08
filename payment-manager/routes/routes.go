package routes

import (
	"payment-manager/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	r.POST("/send", controllers.Send(db))
	r.POST("/withdraw", controllers.Withdraw(db))
}
