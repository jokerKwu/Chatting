package routes

import (
	"Chatting/api"
	"github.com/labstack/echo/v4"
)

func GetPostApiRoutes(e *echo.Echo, userController *api.UserController){
	v1 := e.Group("/api/v1")
	{
		v1.GET("/posts",userController.GetAllUser)
	}
}