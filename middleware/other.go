package middleware

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"go-gateway/global"
)

var JWT = middleware.JWTWithConfig(global.JWTConfig)

var CORS = middleware.CORSWithConfig(middleware.CORSConfig{
	AllowOrigins: []string{"*"},
	AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization, "X-Id"},
})
