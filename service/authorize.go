package service

import (
	"authorize_server/db"
	"authorize_server/dto"
	"authorize_server/model"
	"errors"
	"github.com/golang-jwt/jwt"
	"time"
)

var JwtKey = []byte("123456")

func Signup(signupRequest *dto.SignUpQuest) error {
	now := time.Now()
	userInfo := &model.User{
		Username:  signupRequest.Username,
		Password:  signupRequest.Password,
		CreatedAt: now,
		UpdatedAt: now,
	}
	result := db.DB.Create(userInfo)

	if result.Error != nil {
		return result.Error
	}
	return nil
}

func Login(request *dto.LoginRequest, response *dto.LoginResponse) error {
	user := new(model.User)
	result := db.DB.First(&user, "username = ?", request.Username)
	if result.Error != nil {
		return result.Error
	}
	if request.Password != user.Password {
		return errors.New("password not match")
	}
	expiredIn := 60 * time.Minute
	expireTime := time.Now().Add(expiredIn)

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = expireTime.Unix()
	claims["authorized"] = true
	claims["username"] = user.Username
	claims["userId"] = user.Id
	tokenString, err := token.SignedString(JwtKey)
	if err != nil {
		return err
	}
	response.Scope = request.Scope
	response.ExpiresIn = expireTime.Unix()
	response.AccessToken = tokenString
	response.TokenType = "bearer"
	return nil
}
