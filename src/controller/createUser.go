package controller

import (
	"fmt"

	"github.com/Dener-Fernandes/crud-go/src/configuration/validation"
	"github.com/Dener-Fernandes/crud-go/src/controller/model/request"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest) ; err != nil {
			errRest := validation.ValidateUserError(err)
		
		c.JSON(errRest.Code, errRest)
		return
	}

	fmt.Println(userRequest)
}