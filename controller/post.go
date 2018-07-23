package controller

import (
	"fmt"
	"github.com/ankibahuguna/social/model"
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

func GetAllPosts(c echo.Context) error {
	posts, err := model.FindPosts(c.Get("db").(*gorm.DB))
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "No posts found.")
	}
	return c.JSON(http.StatusOK, posts)
}

func GetSinglePost(c echo.Context) error {
	id := c.Param("id")

	postId, err := strconv.ParseUint(id, 10, 64)

	if err != nil {
		fmt.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid post id.", err)
	}

	post, err := model.FindPostById(c.Get("db").(*gorm.DB), postId)

	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Post doesn't exist")
	}

	return c.JSON(http.StatusOK, post)
}

func CreateNewPost(c echo.Context) error {
	post := &model.Post{}

	if bindingErr := c.Bind(post); bindingErr != nil {
		return bindingErr
	}

	userToken := c.Get("user").(*jwt.Token)
	claims := userToken.Claims.(jwt.MapClaims)
	userId := uint64(claims["uid"].(float64))

	user, userErr := model.FindUserById(c.Get("db").(*gorm.DB), userId)
	if userErr != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Something went wrong")
	}

	post.User = user

	if postErr := post.Create(c.Get("db").(*gorm.DB)); postErr != nil {
		fmt.Println(postErr)
		return echo.NewHTTPError(http.StatusInternalServerError, "Could not save post.")
	}

	return c.JSON(http.StatusCreated, post)

}
