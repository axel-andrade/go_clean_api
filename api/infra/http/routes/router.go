package routes

import (
	composes "go_clean_api/api/infra/composes"

	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		main.POST("/signup", composes.SignUpCompose)
		// categories := main.Group("categories")
		// {
		// 	categories.POST("/", composes.SignUpCompose())
		// }
	}

	return router
}
