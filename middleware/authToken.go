package middleware

import (
	u "Chatting/utils"
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
)

func AuthToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := u.CtxGenerate(c.Request(), "", "")
		// get jwt Token
		accessToken := c.Request().Header.Get("access_token")
		if accessToken == "" {
			return c.JSON(http.StatusBadRequest,errors.New( "no access code in header"))
		}

		// verify & get Data
		tokenData, _, err := u.TokenVerifyAccess(ctx, accessToken, false)
		if err != nil {
			return c.JSON(http.StatusBadRequest, errors.New("token invalid"))
		}
		c.Set("name",tokenData.Name)
		return next(c)
	}
}