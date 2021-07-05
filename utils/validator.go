package utils

import (
	"Chatting/exception"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"net/http"
)

var Val = validator.New()

type CustomValidator struct{
	Validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error{
	if err := cv.Validator.Struct(i); err != nil{
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return nil
}

func Validate(i interface{}) error {
	if err := Val.Struct(i); err != nil{
		return err
	}
	return nil
}

func BindAndValidate(c echo.Context, i interface{}) error {
	if err := c.Bind(i); err != nil {
		return exception.BadRequestException(err.Error())
	}
	if err := c.Validate(i); err != nil {
		return exception.BadRequestException(err.Error())
	}
	return nil
}