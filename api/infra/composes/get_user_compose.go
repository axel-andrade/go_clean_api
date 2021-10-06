package composes

import (
	"github.com/gin-gonic/gin"
)

func GetUserCompose(c *gin.Context) {

	c.JSON(200, "test auth")
}
