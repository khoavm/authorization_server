package controller

import (
	"authorize_server/dto"
	"authorize_server/service"
	"authorize_server/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SignUp(c *gin.Context) {
	signupRequest := new(dto.SignUpQuest)
	if err := c.Bind(signupRequest); err != nil {
		util.Error(c, 400, err)
		return
	}
	if err := service.Signup(signupRequest); err != nil {
		util.Error(c, 500, err)
		return
	}
	util.Success(c, "signup success!")
	return
}
func Login(c *gin.Context) {
	loginRequest := new(dto.LoginRequest)
	if err := c.Bind(loginRequest); err != nil {
		util.Error(c, 400, err)
		return
	}
	loginResponse := new(dto.LoginResponse)
	err := service.Login(loginRequest, loginResponse)
	if err != nil {
		util.Error(c, 401, err)
		return
	}
	util.Success(c, loginResponse)
	return
}
func GoToSignUpPage(c *gin.Context) {
	c.HTML(http.StatusOK, "signup.tmpl", gin.H{
		"title": "Signup Page",
	})

	return
}

func GoToLoginpPage(c *gin.Context) {
	c.HTML(http.StatusOK, "login.tmpl", gin.H{
		"title": "Login Page",
	})

	return
}
