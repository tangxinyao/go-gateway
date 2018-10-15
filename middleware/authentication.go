package middleware

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"go-gateway/global"
	"go-gateway/domain"
)

// Authentication middleware must be placed after the JWT middleware
func Authentication(next echo.HandlerFunc) echo.HandlerFunc {
	return func(context echo.Context) error {
		// Get JWT from pre-middleware
		token := context.Get("user").(*jwt.Token)
		claims := token.Claims.(jwt.MapClaims)
		id := claims["id"].(string)

		// Get raw data from redis
		raw, err := global.RedisClient.Get(global.PermissionPrefix + id).Bytes()
		if err != nil {
			global.MyLogger.Println(err)
			return err
		}

		// Convert binary data to a certain struct
		var permissions []domain.Permission
		err = json.Unmarshal(raw, &permissions)
		if err != nil {
			global.MyLogger.Println(err)
			return err
		}

		// Check whether this user having permissions
		req := context.Request()
		forbidden := true

		for i := 0; i < len(permissions); i++ {
			permission := permissions[i]
			if permission.URL == req.URL.Path && permission.Method == req.Method {
				forbidden = false
				break
			}
		}
		if forbidden {
			return echo.ErrForbidden
		}

		return next(context)
	}
}
