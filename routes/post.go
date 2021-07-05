package routes

import (
	"Chatting/api"
	m "Chatting/middleware"
	"github.com/labstack/echo/v4"
)

func GetPostApiRoutes(e *echo.Echo, postController *api.PostController){
	v1 := e.Group("/api/v1/posts",m.AuthToken)
	{
		v1.GET("",postController.GetAllPost)
		v1.GET("/:id",postController.GetPost)
		v1.POST("",postController.SavePost)
		v1.DELETE("/:id",postController.DeletePost)
		v1.PUT("/:id",postController.UpdatePost)
	}
}