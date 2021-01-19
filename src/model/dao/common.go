package dao

import (
	"context"
	"encoding/json"
	"github.com/go-pg/pg/v10"
	"io/ioutil"
	"os"
)

var _db *pg.DB

type dbCfg struct {
	Addr     string `json:"addr"`
	User     string `json:"user"`
	Password string `json:"password"`
	Database string `json:"database"`
}

func init() {
	path, ok := os.LookupEnv("CFG_ADDR")
	if path == "" || !ok {
		path = `./src/config/db.json`
	}
	content, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	cfg := new(dbCfg)
	if err = json.Unmarshal(content, cfg); err != nil {
		panic(err)
	}
	_db = pg.Connect(&pg.Options{ // TODO config file
		Addr:     cfg.Addr,
		User:     cfg.User,
		Password: cfg.Password,
		Database: cfg.Database,
	})

	ctx := context.Background()

	if err := _db.Ping(ctx); err != nil {
		panic(err)
	}

}

func GetDb() *pg.DB {
	return _db
}
