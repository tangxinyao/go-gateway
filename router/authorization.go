package router

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"net/http"
	"time"
	"go-gateway/global"
	"go-gateway/domain"
)

type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterReq struct {
	Username string   `json:"username"`
	Password string   `json:"password"`
	Roles    []string `json:"roles"`
}

func Login(c echo.Context) error {
	// get request
	req := new(LoginReq)
	if err := c.Bind(req); err != nil {
		global.MyLogger.Println(err)
		return domain.StatusInternalServerError(c, err)
	}

	// get user
	user, err := domain.FindUserByUsernameAndPassword(req.Username, req.Password)
	if err != nil {
		global.MyLogger.Println(err)
		return domain.StatusUnauthorized(c, err)
	}

	// get permissions
	permissions, _ := domain.FindPermissionsByUserId(user.ID)
	pRaw, err := json.Marshal(permissions)
	if err != nil {
		global.MyLogger.Println(err)
		return domain.StatusInternalServerError(c, err)
	}
	global.RedisClient.Set(global.PermissionPrefix+user.ID, string(pRaw), time.Hour*6)

	// generate jwt
	t, err := createToken(user.ID)
	if err != nil {
		global.MyLogger.Println(err)
		return domain.StatusInternalServerError(c, err)
	}
	global.MyLogger.Println(t)
	resp := map[string]interface{}{
		"token":       t,
		"permissions": permissions,
	}
	return domain.StatusOK(c, resp)
}

func Register(c echo.Context) error {
	req := new(RegisterReq)
	if err := c.Bind(req); err != nil {
		return domain.StatusInternalServerError(c, err)
	}

	user, err := domain.CreateUser(req.Username, req.Password, req.Roles)
	if err != nil {
		return domain.StatusInternalServerError(c, err)
	}

	// get permissions
	permissions, _ := domain.FindPermissionsByUserId(user.ID)
	pRaw, err := json.Marshal(permissions)
	if err != nil {
		return domain.StatusInternalServerError(c, err)
	}
	global.RedisClient.Set(global.PermissionPrefix+user.ID, string(pRaw), time.Hour*6)

	// generate jwt
	t, err := createToken(user.ID)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return domain.StatusOK(c, map[string]interface{}{
		"token":       t,
		"permissions": permissions,
	})
}

func createToken(id string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = id
	claims["exp"] = time.Now().Add(time.Hour * 6).Unix()

	t, err := token.SignedString([]byte("secret"))
	return t, err
}
