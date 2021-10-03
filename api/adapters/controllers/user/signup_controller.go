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

	result, err := ctrl.Interactor.Execute(&user)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, result)
}
