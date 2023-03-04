package handlers

import (
	"backEnd/dto/auth"
	"backEnd/dto/result"
	"backEnd/models"
	"backEnd/pkg/bcrypt"
	jwtToken "backEnd/pkg/jwt"
	"backEnd/repositories"
	"log"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type handlerAuth struct {
	AuthRepository repositories.AuthRepository
}

func HandlerAuth(AuthRepository repositories.AuthRepository) *handlerAuth{
	return &handlerAuth{AuthRepository}
}

func (h *handlerAuth) Register(c echo.Context)error{
	request := new(auth.RegisterRequest)
	if err := c.Bind(request); err != nil{
		return c.JSON(http.StatusBadRequest, result.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}
	validation := validator.New()
	err := validation.Struct(request)
	if err != nil{
		return c.JSON(http.StatusBadRequest, result.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}
	password, err := bcrypt.HashingPassword(request.Password)
	if err != nil{
		return c.JSON(http.StatusBadRequest, result.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}
	user := models.User{
		Name: request.Name,
		Email: request.Email,
		Password: password,
	}
	data, err := h.AuthRepository.Register(user)
	if  err != nil{
		return c.JSON(http.StatusInternalServerError, result.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}
	return c.JSON(http.StatusOK, result.SuccessResult{Status: http.StatusOK,Data: data})
}

func (h *handlerAuth)Login(c echo.Context)error{
	request := new(auth.LoginRequest)
	if err := c.Bind(request); err != nil{
		return c.JSON(http.StatusBadRequest, result.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}
	user := models.User{
		Email: request.Email,
		Password: request.Password,
	}
	user, err := h.AuthRepository.Login(user.Email)
	if err != nil {
		return c.JSON(http.StatusBadRequest, result.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}
	isValid := bcrypt.CheckPasswordHash(request.Password, user.Password)
	if !isValid {
		return c.JSON(http.StatusBadRequest, result.ErrorResult{Status: http.StatusBadRequest, Message: "Wrong Email or Password"})
	}

	claims := jwt.MapClaims{}
	claims["id"] = user.ID
	claims["exp"] = time.Now().Add(time.Hour * 2).Unix()

	token, errGenerateToken := jwtToken.GenerateToken(&claims)
	if errGenerateToken != nil{
		log.Println(errGenerateToken)
		return echo.NewHTTPError(http.StatusUnauthorized)
	}
	loginResponse := auth.AuthResponse{
		Name: user.Name,
		Email: user.Email,
		Password: user.Password,
		Token: token,
	}
	return c.JSON(http.StatusOK,result.SuccessResult{Status: http.StatusOK, Data: loginResponse})
}