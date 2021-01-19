package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"tantan_task/src/api"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(handleErr)

	e.GET("/users", api.GetAllUsers)
	e.POST("/users", api.AddUser)
	e.GET("/users/:user_id/relationships", api.GetUserRelation)
	e.PUT("/users/:user_id/relationships/:other_user_id", api.UpdateRelation)

	if err := e.Start(":10532"); err != nil {
		panic(err)
	}
}

func handleErr(handlerFunc echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if err := handlerFunc(c); err != nil {
			if !c.Response().Committed {
				c.JSON(http.StatusInternalServerError, err.Error())
			}
		}
		return nil
	}
}
