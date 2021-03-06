package routes

import (
	"Chatting/api"
	"github.com/labstack/echo/v4"
)

func GetUserApiRoutes(e *echo.Echo, userController *api.UserController){
	v1 := e.Group("/api/v1")
	{
		v1.POST("/login",userController.PostLoginUser)
		v1.POST("/join",userController.PostUser)
	}
}
