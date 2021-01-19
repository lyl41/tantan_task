package api

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"tantan_task/src/common"
	"tantan_task/src/handler"
)

type UpdateRelationReq struct {
	State string `json:"state"`
}

func UpdateRelation(c echo.Context) (err error) {
	u1s := c.Param("user_id")
	u2s := c.Param("other_user_id")
	if u1s == "" || u2s == "" {
		return c.JSON(http.StatusBadRequest, "params err")
	}
	uid1, err := strconv.ParseInt(u1s, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "params err user_id")
	}
	uid2, err := strconv.ParseInt(u2s, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "params err other_user_id")
	}
	if uid1 == uid2 {
		return c.JSON(http.StatusBadRequest, "two user_id can not equal")
	}
	req := new(UpdateRelationReq)
	if err = c.Bind(req); err != nil {
		return
	}
	switch req.State {
	case common.Like, common.DisLike:
	default:
		return c.JSON(http.StatusBadRequest, "state empty")
	}
	resp, err := handler.UpdateRelation(uid1, uid2, req.State)
	if err != nil {
		return
	}
	return c.JSON(http.StatusOK, resp)
}
