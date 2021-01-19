package handler

import (
	"strconv"
	"tantan_task/src/common"
	"tantan_task/src/model/dao"
	"tantan_task/src/model/table"
)

type userRelation struct {
	UserID string `json:"user_id"`
	State  string `json:"state"`
	Type   string `json:"type"`
}

func GetUserRelation(uid int64) (resp []*userRelation, err error) {
	data, err := dao.GetUserRelation(uid)
	if err != nil {
		return
	}
	resp = make([]*userRelation, 0)
	for _, v := range data {
		state, anotherUid := handleRelation(uid, v)
		if state != "" {
			resp = append(resp, &userRelation{
				UserID: strconv.FormatInt(anotherUid, 10),
				State:  state,
				Type:   "relation",
			})
		}
	}
	return
}

func handleRelation(uid int64, v *table.Relation) (state string, anotherUid int64) {
	if v.UID1 == uid {
		anotherUid = v.UID2
		if v.State&common.Uid1LikeUid2 > 0 {
			state = "liked"
			if v.State&common.Uid2LikeUid1 > 0 {
				state = "matched"
			}
		} else if v.State&common.Uid1DisLikeUid2 > 0 {
			state = "disliked"
		}
	} else if v.UID2 == uid {
		anotherUid = v.UID1
		if v.State&common.Uid2LikeUid1 > 0 {
			state = "liked"
			if v.State&common.Uid1LikeUid2 > 0 {
				state = "matched"
			}
		} else if v.State&common.Uid2DisLikeUid1 > 0 {
			state = "disliked"
		}
	}
	return
}
