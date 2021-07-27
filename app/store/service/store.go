package service

import (
	"context"
	"hello/dao"
	"hello/model"
	"hello/pkg/db"
	"hello/pkg/util"
	"hello/protocol"
	"hello/repository"
)

type StoreService struct {
	store db.IDao
}

type StoreServer interface {
	GetStoreList(ctx context.Context, req *protocol.StoreListReq) (resp *protocol.StoreListRes, err error)
}

func New() StoreServer {
	return &StoreService{
		store: dao.NewStoreDetail(),
	}
}

func (ss *StoreService) GetStoreList(ctx context.Context, req *protocol.StoreListReq) (resp *protocol.StoreListRes, err error) {
	list := make([]*repository.StoreDetail, 0)
	resp = new(protocol.StoreListRes)

	err = ss.store.Get(ctx, &dao.StoreDetail{
		Base:      dao.BindCommonReq(req.BaseRequest),
		Id:        req.Filter.Id,
		ChannelId: req.Filter.ChannelId,
		Title:     req.Filter.Title,
		Cellphone: req.Filter.Cellphone,
	}, &list)

	resp.List = make([]*protocol.Store, 0, len(list))

	for _, value := range list {
		pStore, err := ss.toStoreInfo(ctx, value)
		if nil != err {
			model.LogError(ctx, "GetStoreList", "transform timestamp error", "created_at? updated_at? deleted_at？", value, err)
		}
		resp.List = append(resp.List, pStore)
	}

	return resp, err
}

func (ss *StoreService) toStoreInfo(ctx context.Context, item *repository.StoreDetail) (res *protocol.Store, err error) {
	storeData := new(protocol.Store)

	if err := util.CopyStruct(storeData, item); nil != err {
		model.LogError(ctx, "toStoreInfo", "copy struct data", "unknown error?", item, err)
		return storeData, err
	}

	storeData.CreatedAt = item.CreatedAt.Unix()
	storeData.UpdatedAt = item.UpdatedAt.Unix()
	storeData.DeletedAt = item.DeletedAt.Unix()

	return storeData, err
}
