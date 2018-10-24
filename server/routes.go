package accountApi

import (
	"github.com/eduardojonssen/account-api/server/controllers"
	"github.com/gin-gonic/gin"
)

func SetRoutes(engine *gin.Engine) {

	v1 := engine.Group("/api")
	{
		v1.GET("/health-check", controllers.HealthCheckEndpoint)
		v1.POST("account", controllers.CreateAccountEndpoint)
	}
}
