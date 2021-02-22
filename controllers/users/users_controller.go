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

func CreateUser(c *gin.Context) {

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

	result, saveErr := services.CreateUser(user)

	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}

	c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context) {

	userId, err := strconv.ParseInt(c.Param("user_id"), 10, 64)

	if err != nil {
		err := errors.NewBadRequestError(err.Error())
		c.JSON(err.Status, err)
		return
	}

	user, getErr := services.GetUser(userId)

	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}

	c.JSON(http.StatusCreated, user)

}
