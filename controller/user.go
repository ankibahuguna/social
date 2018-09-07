package controller

import (
	"fmt"
	"github.com/ankibahuguna/social/model"
	"github.com/ankibahuguna/social/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
	"time"
)

func GetUsers(c echo.Context) error {
	users, err := model.FindUsers(c.Get("db").(*gorm.DB))
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Users do not exist.")
	}
	return c.JSON(http.StatusOK, users)
}

func SaveUser(c echo.Context) error {
	user := &model.User{}

	if err := c.Bind(user); err != nil {
		return err
	}

	hash, err := utils.HashPassword(user.Password)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Something went wrong.")
	}

	if err := user.EmailExists(c.Get("db").(*gorm.DB)); err == nil {
		return echo.NewHTTPError(http.StatusConflict, "This email is already registered.")
	}

	user.Password = hash
	if err := user.Create(c.Get("db").(*gorm.DB)); err != nil {
		fmt.Println(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Could not register user.")
	}

	tokenStr, err := utils.GenerateJWT(user.ID, time.Duration(24))

	if err != nil {
		return err
	}

	type Response struct {
		User  interface{} `json:"user"`
		Token string      `json:"token"`
	}

	response := Response{
		User:  user,
		Token: tokenStr,
	}

	return c.JSON(http.StatusOK, response)
}

func Login(c echo.Context) error {

	loginPayload := &model.User{}
	if err := c.Bind(loginPayload); err != nil {
		return err
	}
	user := model.User{}

	user, err := model.FindUserByEmail(c.Get("db").(*gorm.DB), loginPayload.Email)

	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "Incorrect email or password")
	}

	if match := utils.CheckPasswordHash(loginPayload.Password, user.Password); match != true {
		return echo.NewHTTPError(http.StatusUnauthorized, "Incorrect email or password")
	}

	tokenStr, err := utils.GenerateJWT(user.ID, time.Duration(24))

	if err != nil {
		return err
	}

	type Response struct {
		User  interface{} `json:"user"`
		Token string      `json:"token"`
	}

	response := Response{
		User:  user,
		Token: tokenStr,
	}

	return c.JSON(http.StatusOK, response)
}

func GetUser(c echo.Context) error {
	id := c.Param("id")

	userId, err := strconv.ParseUint(id, 10, 64)

	if err != nil {
		fmt.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid user id.", err)
	}

	user, err := model.FindUserById(c.Get("db").(*gorm.DB), userId)

	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "User doesn't exist")
	}

	return c.JSON(http.StatusOK, user)
}

func GetUserProfile(c echo.Context) error {
	userToken := c.Get("user").(*jwt.Token)
	claims := userToken.Claims.(jwt.MapClaims)
	userId := uint64(claims["uid"].(float64))

	user, err := model.FindUserById(c.Get("db").(*gorm.DB), userId)

	if err != nil {
		return echo.NewHTTPError(http.StatusForbidden, "Not allowed.")
	}

	return c.JSON(http.StatusOK, user)
}

func ChangePassword(c echo.Context) error {
	return c.JSON(http.StatusOK)
}

func ResetPassword(c echo.Context) error {
	return c.JSON(http.StatusOK)
}

func FollowUser(c echo.Context) error {
	return c.JSON(http.StatusOK)
}

func UnFollowUser(c echo.Context) error {
	return c.JSON(http.StatusOK)
}

func GetUserFollowers(c echo.Context) error {
	return c.JSON(http.StatusOK)
}
