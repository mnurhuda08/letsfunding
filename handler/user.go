package handler

import (
	"letsfunding/helper"
	"letsfunding/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	//recive input user
	var input user.RegisterUserInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		response := helper.APIResponse("Wrong Input", http.StatusBadRequest, "Error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	newUser, err := h.userService.RegisterUser(input)

	if err != nil {
		response := helper.APIResponse("Register Failed", http.StatusBadRequest, "Error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	//generate token code

	formatter := user.FormatUser(newUser, "token")

	response := helper.APIResponse("Registerd Success", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}
