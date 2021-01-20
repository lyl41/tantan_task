package dao

import (
	"fmt"
	"github.com/go-pg/pg/v10"
	"github.com/pkg/errors"
	"tantan_task/src/common"
	"tantan_task/src/model/table"
)

func QueryRelationLock(tx *pg.Tx, uid1, uid2 int64) (state int, exists bool, err error) {
	data := &table.Relation{}
	exists = true
	if err = tx.Model(data).Where("uid1 = ?", uid1).
		Where("uid2 = ?", uid2).For("UPDATE").Select(); err != nil {
		if err == pg.ErrNoRows {
			err = nil
			exists = false
			return
		}
		err = errors.Wrap(err, "")
		return
	}
	state = data.State
	return
}

func GetUserRelation(uid int64) (data []*table.Relation, err error) {
	err = GetDb().Model(&data).Where("uid1=?", uid).
		WhereOr("uid2=?", uid).Select()
	if err != nil {
		if err == pg.ErrNoRows {
			err = nil
			return
		}
		return
	}
	return
}

func UpdateRelation(uid1, uid2 int64, relation string) (res *table.Relation, err error) {
	tx, err := GetDb().Begin()
	if err != nil {
		err = errors.Wrap(err, "")
		return
	}
	ok := false
	defer func() {
		if ok {
			err = tx.Commit()
		} else {
			tx.Rollback()
		}
		tx.Close()
	}()

	curState, exists, err := QueryRelationLock(tx, min(uid1, uid2), max(uid1, uid2))
	if err != nil {
		return
	}
	newState := curState
	switch relation {
	case common.Like:
		if uid1 < uid2 {
			newState |= common.Uid1LikeUid2     // set like symbol
			newState &= ^common.Uid1DisLikeUid2 // remove dislike symbol
		} else {
			newState |= common.Uid2LikeUid1
			newState &= ^common.Uid2DisLikeUid1
		}
	case common.DisLike:
		if uid1 < uid2 {
			newState |= common.Uid1DisLikeUid2
			newState &= ^common.Uid1LikeUid2
		} else {
			newState |= common.Uid2DisLikeUid1
			newState &= ^common.Uid2LikeUid1
		}
	}
	data := &table.Relation{
		UID1:  min(uid1, uid2),
		UID2:  max(uid1, uid2),
		State: newState,
	}
	fmt.Println("exist", exists)
	if !exists { // no record, insert
		_, err = tx.Model(data).Insert()
		if err != nil {
			return
		}
	} else { // update
		_, err = tx.Model(data).Column("state").
			Where("uid1 = ?", min(uid1, uid2)).
			Where("uid2 = ?", max(uid1, uid2)).
			Update()
		if err != nil {
			return
		}
	}

	ok = true
	res = &table.Relation{
		UID1:  min(uid1, uid2),
		UID2:  max(uid1, uid2),
		State: newState,
	}
	return
}

func min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}
