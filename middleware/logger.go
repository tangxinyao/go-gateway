package middleware

import (
	"github.com/labstack/echo"
	"go-gateway/global"
	"time"
)

// Record the log information
func Logger(next echo.HandlerFunc) echo.HandlerFunc {
	return func(context echo.Context) error {
		current := time.Now()
		err := next(context)
		elapsed := time.Now().Sub(current)
		if err != nil {
			global.MyLogger.Println(elapsed)
			return err
		}
		global.MyLogger.Println(elapsed)
		return nil
	}
}
