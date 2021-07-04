package routes

import (
	"Chatting/api"
	"github.com/labstack/echo/v4"
)

func GetPostApiRoutes(e *echo.Echo, postController *api.PostController){
	v1 := e.Group("/api/v1")
	{
		v1.GET("/posts",postController.GetAllPost)
		v1.GET("/posts:id",postController.GetPost)
		v1.POST("/posts",postController.SavePost)
		v1.DELETE("/posts:id",postController.UpdatePost)
		v1.PUT("/posts:id",postController.UpdatePost)
	}
}