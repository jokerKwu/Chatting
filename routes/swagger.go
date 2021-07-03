package routes

import (
	"Chatting/api"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func GetSwaggerRoutes(e *echo.Echo){
	e.GET("/api", api.RedirectIndexPage)
	e.GET("/api/*",echoSwagger.WrapHandler)
}
