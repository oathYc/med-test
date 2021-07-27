package api

import (
	"context"
	"hello/app/medchan/service"
	"hello/protocol"
	"hello/server"
)

type MedChanOutServer struct {
	server.UnimplementedMedTestSerServer
	// channel 服务
	medChanServer service.MedChanServer
}

func (mcs *MedChanOutServer) New() *MedChanOutServer {
	return &MedChanOutServer{
		medChanServer: service.New(),
	}
}

func (mcs *MedChanOutServer) GetMedChanList(ctx context.Context, req *protocol.MedChanListReq) (res *protocol.MedChanListRes, err error) {
	return mcs.New().medChanServer.GetChannelList(ctx, req)
}
