package repository

import (
	"context"

	"github.com/jinzhu/gorm"

	medmodel "hello/model"
	"hello/pkg/db"
	"hello/pkg/db/builder"
)

type Dao struct {
	*db.CommonReq
	tabler  db.Tabler
	filter  db.IFilter
	builder builder.IBuilder
}

// NewDao new
func NewDao(model db.IDao) *Dao {
	connect, ok := model.(db.Connecter)
	if !ok {
		return nil
	}
	build := medmodel.DbWrapper[connect.Connection()].Dsn
	table, ok := model.(db.Tabler)
	if !ok {
		return nil
	}
	filter, ok := model.(db.IFilter)
	if !ok {
		return nil
	}
	return &Dao{
		tabler:  table,
		filter:  filter,
		builder: build,
	}
}

// register dao
func RegisterDao(model db.IDao) *Dao {
	return NewDao(model)
}

func (d *Dao) Add(ctx context.Context, value db.Tabler) error {
	return d.builder.SetLogger(&dbLog{ctx: ctx}).Add(value)
}

func (d *Dao) Update(ctx context.Context, param db.IQuery, update map[string]interface{}) error {
	return d.filter.BuildFilterQuery(param).SetLogger(&dbLog{ctx: ctx}).Update(update)
}

func (d *Dao) Delete(ctx context.Context, param db.IQuery) error {
	return d.filter.BuildFilterQuery(param).SetLogger(&dbLog{ctx: ctx}).Delete(d.tabler)
}

func (d *Dao) Find(ctx context.Context, param db.IQuery, result db.Tabler) error {
	return d.filter.BuildFilterQuery(param).SetLogger(&dbLog{ctx: ctx}).Find(result)
}

func (d *Dao) Get(ctx context.Context, param db.IQuery, result interface{}) error {
	return d.filter.Filter(d.filter.BuildFilterQuery(param), param.GetCommonReq()).SetLogger(&dbLog{ctx: ctx}).Get(result)
}

func (d *Dao) Where(query interface{}, args ...interface{}) builder.IBuilder {
	return d.builder.Where(query, args...)
}

func (d *Dao) GetCommonReq() *db.CommonReq {
	if nil == d {
		return &db.CommonReq{}
	}
	return d.CommonReq
}

func (d *Dao) GetBuilder() builder.IBuilder {
	return d.builder
}

// filter
func (d *Dao) Filter(build builder.IBuilder, condition *db.CommonReq) builder.IBuilder {
	if nil != condition {
		// Start
		if start := condition.Start; start > 0 {
			build = build.Offset(start)
		}
		// Limit
		if limit := condition.Limit; limit > 0 {
			build = build.Limit(limit)
		}
		// 排序
		for _, sort := range condition.Sorts {
			if string(sort[0]) == "+" {
				sort = string(sort[1:]) + " ASC"
			} else if string(sort[0]) == "-" {
				sort = string(sort[1:]) + " DESC"
			} else {
				sort = sort + " ASC"
			}
			build = build.Order(sort)
		}
	}

	return build
}

type dbLog struct {
	ctx context.Context
}

func (l *dbLog) Print(v ...interface{}) {
	if l.ctx == nil {
		l.ctx = context.Background()
	}
	traceId := medmodel.GetTraceId(l.ctx)
	medmodel.GRpcLog.Infof("trace_id: [%s]; args: [%v];", traceId, gorm.LogFormatter(v...))
}
