package handler

import (
	"strconv"
	"tantan_task/src/model/dao"
)

type userData struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

func GetAllUsers() (data []*userData, err error) {
	dbData, err := dao.GetAllUsers()
	if err != nil {
		return
	}
	data = make([]*userData, 0, len(dbData))
	for _, v := range dbData {
		data = append(data, &userData{
			ID:   strconv.FormatInt(v.ID, 10),
			Name: v.Name,
			Type: "user",
		})
	}
	return
}
