package dao

import (
	"github.com/go-pg/pg/v10"
	"github.com/pkg/errors"
	"tantan_task/src/model/table"
)

func GetAllUsers() (res []*table.Users, err error) {
	if err = GetDb().Model(&res).Select(); err != nil {
		err = errors.Wrap(err, "")
		return
	}
	return
}

func GetUserByID(id int64) (res *table.Users, err error) {
	if id <= 0 {
		err = errors.Wrap(errors.New("id <= 0"), "")
		return
	}
	tmp := new(table.Users)
	if err = GetDb().Model(tmp).Where("id = ?", id).Select(); err != nil {
		if err == pg.ErrNoRows {
			err = nil
			return
		}
		err = errors.Wrap(err, "")
		return
	}
	res = tmp
	return
}
