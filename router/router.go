package router

import (
	"git.medlinker.com/service/grpcwrapper"
	"google.golang.org/grpc/reflection"

	"hello/app/medchan/api"
	storeapi "hello/app/store/api"
	serprotoc "hello/server"
)

// RegistergRPCServers 注册所有的 gRPC 服务
func RegistergRPCServers(engi *grpcwrapper.Engine) {

	// 注册channel相关服务
	serprotoc.RegisterMedTestSerServer(engi.GetRawServer(), new(api.MedChanOutServer))

	// 注册 store 相关服务
	serprotoc.RegisterStoreSerServer(engi.GetRawServer(), new(storeapi.StoreOutServer))

	// DONE(@yeqiang) 增加grpc调试工具
	reflection.Register(engi.GetRawServer())
}
