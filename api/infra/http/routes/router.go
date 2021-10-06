package routes

import (
	composes "go_clean_api/api/infra/composes"
	"go_clean_api/api/infra/http/middlewares"

	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		main.POST("/signup", composes.SignUpCompose)
		main.POST("/login", composes.LoginCompose)
		main.POST("/logout", composes.LogoutCompose)

		users := main.Group("users")
		{
			users.GET("/:id", middlewares.Authorize(), composes.GetUserCompose)
		}
	}

	return router
}
