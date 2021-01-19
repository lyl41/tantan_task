package dao

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestGetAllUsers(t *testing.T) {

	gotRes, err := GetAllUsers()
	j, _ := json.Marshal(gotRes)
	fmt.Println(err, string(j))

}
