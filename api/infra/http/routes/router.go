package routes

import (
	composes "go_clean_api/api/infra/composes"
	"go_clean_api/api/infra/http/middlewares"

	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {

	main := router.Group("/")
	{
		main.GET("/healthcheck", func(c *gin.Context) {
			c.JSON(200, gin.H{"status": "OK"})
		})
	}

	v1 := router.Group("api/v1")
	{

		v1.POST("/signup", composes.SignUpCompose)
		v1.POST("/login", composes.LoginCompose)
		v1.POST("/logout", composes.LogoutCompose)

		users := v1.Group("users")
		{
			users.GET("/", middlewares.Authorize(), composes.FindUsersCompose)
		}
	}

	return router
}
