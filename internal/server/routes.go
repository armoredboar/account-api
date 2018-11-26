package accountApi

import (
	"github.com/armoredboar/account-api/internal/server/controllers"

	"github.com/gin-gonic/gin"
)

func SetRoutes(engine *gin.Engine) {

	v1 := engine.Group("/account-api/v1")
	{
		v1.GET("/health-check", controllers.HealthCheckEndpoint)
		v1.GET("/account/check", controllers.CheckAccountEndpoint)
		v1.POST("/account", controllers.CreateAccountEndpoint)
		v1.GET("/account/activate", controllers.ValidateActivationKeyEndpoint)
		v1.POST("/account/resend-activation-email", controllers.ResendActivationEmailEndpoint)
		v1.POST("/account/change-activation-email", controllers.ChangeActivationEmailEndpoint)

		v1.GET("/oauth/authorize", controllers.Authorize)
		v1.POST("/oauth/access-token", controllers.AccessToken)
	}
}
