package api

import (
	"context"
	"hello/app/store/service"
	"hello/protocol"
	"hello/server"
)

type StoreOutServer struct {
	server.UnimplementedStoreSerServer

	storeServer service.StoreServer
}

func (ss *StoreOutServer) New() *StoreOutServer {
	return &StoreOutServer{
		storeServer: service.New(),
	}
}

func (ss *StoreOutServer) GetStoreList(cxt context.Context, req *protocol.StoreListReq) (resp *protocol.StoreListRes, err error) {
	return ss.New().storeServer.GetStoreList(cxt, req)
}
