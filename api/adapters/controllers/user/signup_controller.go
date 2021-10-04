package controllers

import (
	interactor "go_clean_api/api/usecases/user/signup"

	"github.com/gin-gonic/gin"
)

// Exemplo request http
// type HTTPSignUpInput = {
// 	params: any;
// 	headers?: any;
// 	body: SignUpRequestDTO;
//   };

type SignUpController struct {
	Interactor interactor.SignUpInteractor
}

func (ctrl *SignUpController) Run(c *gin.Context) {
	var input interactor.SignUpInputDTO

	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid JSON: " + err.Error()})
		return
	}

	output := ctrl.Interactor.Execute(input)

	// Isso não deve ficar na controller e sim na composição
	if output.StatusCode != 201 {
		c.JSON(int(output.StatusCode), gin.H{"error": output.Error})
		return
	}

	c.JSON(201, output.Data)
}
