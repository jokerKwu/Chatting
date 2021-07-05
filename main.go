package main

import (
	"Chatting/api"
	"Chatting/config"
	_ "Chatting/docs"
	"Chatting/routes"
	"Chatting/utils"
	"fmt"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var userController *api.UserController
var postController *api.PostController

// @title Golang CHATTING REST API
// @version 1.0
// @host localhost:8080
// @BasePath /api/v1
func main() {
	e := echo.New()
	e.HTTPErrorHandler = api.ErrorHandler

	//middleware
	e.Validator = &utils.CustomValidator{Validator: validator.New()}
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	config.CORSConfig(e)

	routes.GetUserApiRoutes(e, userController)
	routes.GetPostApiRoutes(e, postController)
	routes.GetSwaggerRoutes(e)

	e.Logger.Fatal(	e.Start(fmt.Sprintf(":%s",config.ServerPort.(string))))
}
