package config

import (
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/echo/v4"
	"net/http"
)

func CORSConfig(e *echo.Echo){
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))
}