package router

import (
	"github.com/labstack/echo"
	"go-gateway/middleware"
)

func SetRouter(e *echo.Echo) {
	authRouter := e.Group("/auth")
	authRouter.POST("/login", Login)
	authRouter.POST("/register", Register)
	redirectRouter := e.Group("/redirect")
	redirectRouter.Use(middleware.JWT, middleware.Proxy())
}
