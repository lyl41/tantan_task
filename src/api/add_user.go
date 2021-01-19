package api

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
	"tantan_task/src/handler"
)

type AddUserReq struct {
	Name string `json:"name"`
}

func AddUser(c echo.Context) (err error) {
	req := new(AddUserReq)
	if err = c.Bind(req); err != nil {
		return
	}
	req.Name = strings.TrimSpace(req.Name)
	if req.Name == "" {
		return c.JSON(http.StatusBadRequest, "name is empty") // TODO
	}
	resp, err := handler.AddUser(req.Name)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key") {
			return c.JSON(http.StatusBadRequest, "name already exist")
		}
		return
	}
	return c.JSON(http.StatusOK, resp)
}
