package middleware

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/levigross/grequests"
	"go-gateway/global"
	"go-gateway/service"
	"go-gateway/domain"
)

func Proxy() echo.MiddlewareFunc {
	redirect := global.RedirectUrl
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req := c.Request()
			scheme := c.Scheme()
			method := req.Method
			url, err := service.FindNoRedirectPrefix(req.RequestURI)
			if err != nil {
				global.MyLogger.Println(err)
				return echo.ErrNotFound
			}
			url = fmt.Sprintf("%s://%s/apis/v1%s", scheme, redirect, url)
			global.MyLogger.Println(url)

			// Get JWT from pre-middleware
			token := c.Get("user").(*jwt.Token)
			claims := token.Claims.(jwt.MapClaims)
			userId := claims["id"].(string)
			options := new(grequests.RequestOptions)
			options.Headers = map[string]string{
				"user_id": userId,
			}
			resp, err := grequests.DoRegularRequest(method, url, options)
			if err != nil {
				global.MyLogger.Println(err)
				return domain.StatusInternalServerError(c, err)
			}

			return c.String(200, resp.String())
		}
	}
}
