package tool

import (
	"github.com/go-xorm/xorm"
	"CloudRestaurant/model"
	_ "github.com/go-sql-driver/mysql"
)

var DbEngine *Orm

type Orm struct {
	*xorm.Engine
}

func OrmEngine(cfg *Config) (*Orm, error) {
	database := cfg.Database
	conn := database.User + ":" + database.Password + "@tcp(" + database.Host + ":" + database.Port + ")/" + database.DbName + "?charset=" + database.Charset
	engine, err := xorm.NewEngine(database.Driver, conn)
	if err != nil {
		return nil, err
	}

	engine.ShowSQL(database.ShowSql)

	err = engine.Sync2(new(model.SmsCode),
		new(model.Member))
	if err != nil {
		return nil, err
	}

	orm := new(Orm)
	orm.Engine = engine
	DbEngine = orm

	return orm, nil
}
