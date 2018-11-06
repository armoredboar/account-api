package accountApi

import (
	"github.com/armoredboar/account-api/internal/server/controllers"
	"github.com/gin-gonic/gin"
)

func SetRoutes(engine *gin.Engine) {

	v1 := engine.Group("/api")
	{
		v1.GET("/health-check", controllers.HealthCheckEndpoint)
		v1.GET("/account/check", controllers.CheckAccountEndpoint)
		v1.POST("/account", controllers.CreateAccountEndpoint)
		v1.GET("/account/activate", controllers.ValidateActivationCodeEndpoint)
		v1.GET("/account/resend-activation", controllers.ResendActivationCodeEndpoint)
	}
}