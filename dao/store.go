package dao

import (
	"context"
	"hello/pkg/db"
	"hello/pkg/db/builder"
	"hello/repository"
)

type StoreDetail struct {
	Base
	repository.StoreDetail

	// search
	Id        uint32
	ChannelId int32
	Title     string
	Cellphone string
}

type NewStoreShow interface {
	db.IDao
}

func NewStoreDetail() db.IDao {
	bd := new(StoreDetail)
	dao := repository.RegisterDao(bd)
	bd.Dao = dao
	return dao
}

func (sd *StoreDetail) BuildFilterQuery(filter db.IQuery) builder.IBuilder {
	build := sd.GetBuilder()
	storeDetailConn, ok := filter.(*StoreDetail)
	if !ok {
		return build
	}

	if id := storeDetailConn.Id; id > 0 {
		build.Where("id > ?", id)
	}

	if channelId := storeDetailConn.ChannelId; channelId > 0 {
		build.Where("channel_id > ?", channelId)
	}

	if title := storeDetailConn.Title; len(title) > 0 {
		build.Where("title like \"%?%\"", title)
	}

	if phone := storeDetailConn.Cellphone; len(phone) > 0 {
		build.Where("cellphone = \"?\"", phone)
	}

	return build
}

func (sd *StoreDetail) Count(ctx context.Context, param db.IQuery) int {
	panic("implement me")
}
