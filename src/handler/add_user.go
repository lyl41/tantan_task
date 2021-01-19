package handler

import (
	"math/rand"
	"strconv"
	"tantan_task/src/model/dao"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func AddUser(name string) (resp userData, err error) {
	id := rand.Int63()
	if err = dao.AddUser(id, name); err != nil {
		return
	}
	resp = userData{
		ID:   strconv.FormatInt(id, 10),
		Name: name,
		Type: "user",
	}
	return
}
