package domain

import (
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
	"go-gateway/global"
	"go-gateway/service"
	"time"
)

type Base struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (Base *Base) BeforeCreate(scope *gorm.Scope) error {
	id, err := uuid.NewV1()
	if err != nil {
		return err
	}
	scope.SetColumn("id", id.String())
	scope.SetColumn("created_at", time.Now())
	scope.SetColumn("updated_at", time.Now())
	return nil
}

func (Base *Base) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("updated_at", time.Now())
	return nil
}

// 用户
type User struct {
	Base
	Username string `json:"username" gorm:"not null;unique"`
	Password string `json:"password" gorm:"not null"`
	Salt     string `json:"salt"`
	Roles    []Role `json:"roles" gorm:"many2many:user_roles"`
}

// 角色
type Role struct {
	Base
	Name        string       `json:"name" gorm:"not null;unique"`
	Permissions []Permission `json:"permissions" gorm:"many2many:role_permissions"`
}

// 权限
type Permission struct {
	Base
	Name   string `json:"name" gorm:"not null;unique"`
	Method string `json:"method"`
	URL    string `json:"url"`
}

// methods
func CreateUser(username string, password string, roles []string) (*User, error) {
	user := new(User)
	user.Username = username
	salt, err := service.GenerateSalt([]byte(password))
	if err != nil {
		return nil, err
	}

	// set salt and password
	user.Salt = salt
	user.Password = service.HashPassword(password, salt)

	// set roles
	if len(roles) > 0 {
		global.MySQLClient.Where("name in (?)", roles).Find(&user.Roles)
	}
	// insert user
	global.MySQLClient.Create(user)
	return user, nil
}

func FindUserByUsernameAndPassword(username string, password string) (*User, error) {
	user := new(User)
	global.MySQLClient.Where("username = ?", username).Find(&user)
	if user.ID == "" {
		return nil, errors.New("user not found")
	}
	hashPassword := service.HashPassword(password, user.Salt)
	if hashPassword != user.Password {
		return nil, errors.New("password incorrect")
	}
	return user, nil
}

func FindPermissionsByUserId(userId string) ([]Permission, error) {
	var permissions []Permission
	global.MySQLClient.
		Joins("JOIN role_permissions on role_permissions.permission_id = permissions.id").
		Joins("JOIN roles on roles.id = role_permissions.role_id").
		Joins("JOIN user_roles on user_roles.role_id = roles.id").
		Where("user_roles.user_id = ?", userId).
		Find(&permissions)
	return permissions, nil
}
