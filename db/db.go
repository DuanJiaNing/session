package db

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"xorm.io/core"

	"session/conf"
	"session/log"
)

var (
	engine *xorm.Engine
)

func Engine() (*xorm.Engine, error) {
	if engine == nil {
		err := prepareEngine()
		if err != nil {
			return nil, err
		}
	}

	return engine, nil
}

func prepareEngine() error {
	var err error
	dataSource := getDataSource()
	engine, err = xorm.NewEngine("mysql", dataSource)
	if err != nil {
		log.Errorf("fail to create db engine: %v", err)
		return err
	}
	err = engine.Ping()
	if err != nil {
		log.Errorf("fail to ping db: %v", err)
		return err
	}

	engine.SetMapper(core.SnakeMapper{})

	engine.ShowSQL(true)
	engine.ShowExecTime(true)

	return nil
}

func getDataSource() string {
	// "user:password@tcp(localhost:5555)/dbname?tls=skip-verify&autocommit=true"
	dsn := conf.DataSource()
	return fmt.Sprintf("%v:%v@tcp(%v:%v)/session", dsn.Username, dsn.Password, dsn.Host, dsn.Port)
}
