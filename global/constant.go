package global

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/middleware"
)

var PermissionPrefix = "PermissionPrefix"

var SigningKey = []byte("secret")

var JWTConfig = middleware.JWTConfig{
	Claims:     jwt.MapClaims{},
	SigningKey: SigningKey,
}
