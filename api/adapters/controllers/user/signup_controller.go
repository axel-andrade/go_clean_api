package controllers

import (
	"go_clean_api/api/entities"
	usecases "go_clean_api/api/usecases/user"

	"github.com/gin-gonic/gin"
)

type SignUpController struct {
	Interactor usecases.SignUpInteractor
}

func (ctrl *SignUpController) Run(c *gin.Context) {
	var user entities.User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid JSON: " + err.Error()})
		return
	}

	output := ctrl.Interactor.Execute(&user)
	if output.StatusCode != 200 {
		c.JSON(int(output.StatusCode), gin.H{"error": output.Error})
		return
	}

	c.JSON(201, output.Data)
}
