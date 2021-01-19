package api

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"tantan_task/src/handler"
)

func GetAllUsers(c echo.Context) (err error) {
	res, err := handler.GetAllUsers()
	if err != nil { // TODO  handler err
		return
	}
	return c.JSON(http.StatusOK, res)
}
