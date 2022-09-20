package router

import (
	"authorize_server/controller"
	"github.com/gin-gonic/gin"
)

func AddAuthorizeRouter(router *gin.Engine) {
	router.GET("/signup", controller.GoToSignUpPage)
	router.POST("/signup", controller.SignUp)
	router.GET("/login", controller.GoToLoginpPage)
	router.POST("/login", controller.Login)
}
