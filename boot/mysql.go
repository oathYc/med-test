package boot

import (
	"hello/model"
	"hello/pkg/db"
	"hello/pkg/db/config"
	"hello/pkg/db/wrapper"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func bootMysql() (err error) {
	cfg := model.GetConfig().Database

	model.DbWrapper = make(map[string]*wrapper.Wrapper)

	for name, database := range cfg {
		model.DbWrapper[name], err = db.GetConnectAdapter("").Connect(&config.Config{
			Driver:          database.Driver,
			Source:          database.Source,
			ConnMaxLifeTime: database.ConnMaxLifeTime,
			Log: config.Log{
				Mode: database.Log.Mode,
			},
		})

		if nil != err {
			return err
		}
	}

	return nil
}
