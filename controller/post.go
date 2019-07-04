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
)

var genericResponse = utils.GenericResponse{
	Data: "Hey",
}

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

func DeleteSinglePost(c echo.Context) error {
	id := c.Param("id")
	postId, err := strconv.ParseUint(id, 10, 64)

	if err != nil {
		fmt.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid post id", err)
	}

	post, deletionError := model.DeletePost(c.Get("db").(*gorm.DB), postId)

	if deletionError != nil {
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

func GetUserPosts(c echo.Context) error {

	id := c.Param("id")
	userId, err := strconv.ParseUint(id, 10, 32)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid user id.", err)
	}

	_, userErr := model.FindUserById(c.Get("db").(*gorm.DB), userId)

	if userErr != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "User does not exists")
	}

	posts, err := model.FindPostsByAuthor(c.Get("db").(*gorm.DB), userId)

	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "No posts found.")
	}

	return c.JSON(http.StatusOK, posts)

}

func PostComment(c echo.Context) error {
	return c.JSON(http.StatusOK, genericResponse)
}

func EditComment(c echo.Context) error {
	return c.JSON(http.StatusOK, genericResponse)
}

func DeleteComment(c echo.Context) error {
	return c.JSON(http.StatusOK, genericResponse)
}

func GetComments(c echo.Context) error {
	return c.JSON(http.StatusOK, genericResponse)
}

func LikePost(c echo.Context) error {
	return c.JSON(http.StatusOK, genericResponse)
}

func UnlikePost(c echo.Context) error {
	return c.JSON(http.StatusOK, genericResponse)
}

func GetLikes(c echo.Context) error {
	return c.JSON(http.StatusOK, genericResponse)
}
