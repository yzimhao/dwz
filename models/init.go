package models

import (
	"sync"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"xorm.io/xorm"
)

var (
	engine           *xorm.Engine
	klineTableMap    sync.Map
	tradeLogTableMap sync.Map
)

func InitDbEngine(driver, dsn string, debug bool) {
	conn, err := xorm.NewEngine(driver, dsn)
	if err != nil {
		logrus.Panic(err)
	}
	engine = conn

	if debug {
		engine.ShowSQL(true)
	}
}

func DbEngine() *xorm.Engine {
	return engine
}
