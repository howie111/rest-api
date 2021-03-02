package users

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/howie111/rest-api/services"
	"github.com/howie111/rest-api/utils/errors"

	"github.com/howie111/rest-api/domains/users"

	"github.com/gin-gonic/gin"
)

func getUserId(userIdParam string) (int64, *errors.RestError) {
	userId, err := strconv.ParseInt(userIdParam, 10, 64)

	if err != nil {
		return 0, errors.NewBadRequestError("user id should be a number")
	}

	return userId, nil
}

func Create(c *gin.Context) {

	user := users.User{}

	bytes, err := ioutil.ReadAll(c.Request.Body)

	if err != nil {
		return
	}
	unmarshalErr := json.Unmarshal(bytes, &user)

	if unmarshalErr != nil {
		resErr := errors.NewBadRequestError(unmarshalErr.Error())
		c.JSON(resErr.Status, resErr)
		return
	}

	result, saveErr := services.UsersService.CreateUser(user)

	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}

	c.JSON(http.StatusCreated, result.Marshall(c.GetHeader("X-Public") == "true"))
}

func Get(c *gin.Context) {

	userId, idErr := getUserId(c.Param("user_id"))

	if idErr != nil {
		c.JSON(idErr.Status, idErr)
	}

	user, getErr := services.UsersService.GetUser(userId)

	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}

	c.JSON(http.StatusCreated, user.Marshall(c.GetHeader("X-Public") == "true"))

}

func Update(c *gin.Context) {

	userId, idErr := getUserId(c.Param("user_id"))

	if idErr != nil {
		c.JSON(idErr.Status, idErr)
	}

	user := users.User{}

	user.Id = userId

	bytes, err := ioutil.ReadAll(c.Request.Body)

	if err != nil {
		return
	}
	unmarshalErr := json.Unmarshal(bytes, &user)

	if unmarshalErr != nil {
		resErr := errors.NewBadRequestError(unmarshalErr.Error())
		c.JSON(resErr.Status, resErr)
		return
	}

	var isPatch bool
	if c.Request.Method == http.MethodPatch {
		isPatch = true
	}

	result, updateErr := services.UsersService.UpdateUser(isPatch, user)

	if updateErr != nil {
		err := errors.NewBadRequestError(updateErr.Error)
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, result.Marshall(c.GetHeader("X-Public") == "true"))

}

func Delete(c *gin.Context) {

	userId, idErr := getUserId(c.Param("user_id"))

	if idErr != nil {
		c.JSON(idErr.Status, idErr)
	}

	deleteErr := services.UsersService.DeleteUser(userId)

	if deleteErr != nil {
		c.JSON(deleteErr.Status, deleteErr)
		return
	}

	c.JSON(http.StatusOK, map[string]string{"status": "deleted"})

}

func Search(c *gin.Context) {
	status := c.Query("status")

	users, err := services.UsersService.SearchUser(status)

	if err != nil {
		c.JSON(err.Status, err)
	}

	c.JSON(http.StatusOK, users.Marshall(c.GetHeader("X-Public") == "true"))

}
