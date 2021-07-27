package connect

import (
	"time"

	"github.com/jinzhu/gorm"

	"hello/pkg/db/builder"
	"hello/pkg/db/config"
	"hello/pkg/db/wrapper"
)

// Gorm连接
type Gorm struct {
	Base
}

func NewGorm() *Gorm {
	return &Gorm{}
}

// 连接
func (g *Gorm) Connect(config *config.Config) (*wrapper.Wrapper, error) {
	var (
		err error
		gdb *gorm.DB
		w   = new(wrapper.Wrapper)
	)
	gdb, err = gorm.Open(config.Driver, config.Source)
	if err != nil {
		return w, err
	}
	gdb.DB().SetConnMaxLifetime(time.Duration(config.ConnMaxLifeTime) * time.Second)
	gdb.DB().SetMaxIdleConns(config.MaxIdleConns)
	gdb.DB().SetMaxOpenConns(config.MaxOpenConns)
	if config.Log.Mode > 0 {
		gdb.LogMode(true)
	}
	w.Dsn = builder.NewGorm(gdb, builder.WithIsWrite(true))
	//for _, s := range config.Slave {
	//	var gdbs *gorm.DB
	//
	//	gdbs, err = gorm.Open(config.Driver, s.Source)
	//	if err != nil {
	//		return w, err
	//	}
	//	gdbs.DB().SetConnMaxLifetime(time.Duration(config.ConnMaxLifeTime) * time.Second)
	//	gdbs.DB().SetMaxIdleConns(config.MaxIdleConns)
	//	gdbs.DB().SetMaxOpenConns(config.MaxOpenConns)
	//	if config.Log.Mode > 1 {
	//		gdbs.LogMode(true)
	//	}
	//	w.Slave = append(w.Slave, builder.NewGorm(gdbs))
	//}
	return w, err
}
