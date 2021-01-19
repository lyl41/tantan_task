package dao

import (
	"github.com/pkg/errors"
	"tantan_task/src/model/table"
)

func AddUser(id int64, name string) (err error) {
	data := &table.Users{
		ID:   id,
		Name: name,
	}
	if _, err = GetDb().Model(data).Insert(); err != nil {
		err = errors.Wrap(err, "")
		return
	}
	return
}
