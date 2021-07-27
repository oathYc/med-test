package builder

import (
	"reflect"

	"github.com/jinzhu/gorm"
)

// Gorm连接
type Gorm struct {
	opts options
	Base
	*gorm.DB
}

// 可选参数列表
type options struct {
	isWrite bool
}

// 为可选参数赋值的函数
type ServerOption func(*options)

// 是否为写库
func WithIsWrite(isWrite bool) ServerOption {
	return func(o *options) {
		o.isWrite = isWrite
	}
}

func NewGorm(gdb *gorm.DB, sos ...ServerOption) IBuilder {
	var opts options
	for _, so := range sos {
		so(&opts)
	}
	return &Gorm{
		opts: opts,
		DB:   gdb,
	}
}

// Add
func (g *Gorm) Add(value interface{}) error {
	return g.DB.Create(value).Error
}

// Update
func (g *Gorm) Update(attrs ...interface{}) error {
	return g.DB.Update(attrs).Error
}

// delete
func (g *Gorm) Delete(value interface{}, where ...interface{}) error {
	return g.DB.Delete(value, where...).Error
}

// find
func (g *Gorm) Find(out interface{}, where ...interface{}) error {
	return g.First(out).Error
}

// get
func (g *Gorm) Get(out interface{}, where ...interface{}) error {
	return g.DB.Find(out).Error
}

// count
func (g *Gorm) Count(value interface{}) error {
	return g.DB.Count(value).Error
}

// where
func (g *Gorm) Where(query interface{}, args ...interface{}) IBuilder {
	return NewGorm(g.DB.Where(query, args...))
}

// or
func (g *Gorm) Or(query interface{}, args ...interface{}) IBuilder {
	return NewGorm(g.DB.Or(query, args...))
}

// offset
func (g *Gorm) Offset(offset interface{}) IBuilder {
	return NewGorm(g.DB.Offset(offset))
}

// limit
func (g *Gorm) Limit(limit interface{}) IBuilder {
	return NewGorm(g.DB.Limit(limit))
}

// order
func (g *Gorm) Order(value interface{}, reorder ...bool) IBuilder {
	var db = g.DB
	if reflect.ValueOf(value).Kind() == reflect.Slice {
		if sort, ok := value.([]string); ok {
			for _, s := range sort {
				db = db.Order(s, reorder...)
			}
		}
	} else {
		db = db.Order(value, reorder...)
	}
	return NewGorm(db)
}

// begin
func (g *Gorm) Begin() (IBuilder, error) {
	return NewGorm(g.DB.Begin()), g.DB.Error
}

// rollback
func (g *Gorm) Rollback() error {
	return g.DB.Rollback().Error
}

// commit
func (g *Gorm) Commit() error {
	return g.DB.Commit().Error
}

// Model
func (g *Gorm) Model(value interface{}) IBuilder {
	return NewGorm(g.DB.Model(value))
}

// IsEmpty
func (g *Gorm) IsEmpty(err error) bool {
	var flag bool
	if err == gorm.ErrRecordNotFound {
		flag = true
	}
	return flag
}

// IsWrite
func (g *Gorm) IsWrite() bool {
	return g.opts.isWrite
}

// SetLogger
func (g *Gorm) SetLogger(log logger) IBuilder {
	g.DB.SetLogger(log)
	return NewGorm(g.DB)
}
