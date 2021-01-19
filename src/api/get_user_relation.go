package api

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"tantan_task/src/handler"
)

func GetUserRelation(c echo.Context) (err error) {
	us := c.Param("user_id")
	fmt.Println(us)
	uid, err := strconv.ParseInt(us, 10, 64)
	if err != nil {
		return
	}
	if uid <= 0 {
		return c.JSON(http.StatusBadRequest, "user_id <= 0")
	}
	resp, err := handler.GetUserRelation(uid)
	if err != nil {
		return
	}
	return c.JSON(http.StatusOK, resp)
}
