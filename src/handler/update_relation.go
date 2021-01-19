package handler

import (
	"fmt"
	"strconv"
	"tantan_task/src/model/dao"
)

func UpdateRelation(uid1, uid2 int64, relation string) (resp userRelation, err error) {
	info, err := dao.GetUserByID(uid1)
	if err != nil {
		return
	}
	if info == nil {
		err = fmt.Errorf("user_id not exist: %d", uid1)
		return
	}
	info, err = dao.GetUserByID(uid2)
	if err != nil {
		return
	}
	if info == nil {
		err = fmt.Errorf("user_id not exist: %d", uid2)
		return
	}
	res, err := dao.UpdateRelation(uid1, uid2, relation)
	if err != nil {
		return
	}
	stateText, anotherUid := handleRelation(uid1, res)
	resp = userRelation{
		UserID: strconv.FormatInt(anotherUid, 10),
		State:  stateText,
		Type:   "relationship",
	}
	return
}
