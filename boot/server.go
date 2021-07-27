package boot

import (
	"hello/model"
	"hello/router"
	"log"

	"git.medlinker.com/golang/xerror"
	"git.medlinker.com/service/grpcwrapper"
	glog "git.medlinker.com/service/grpcwrapper/log"
)

type Boot struct {
	engine *grpcwrapper.Engine
}

func NewServer() *Boot {
	return &Boot{}
}

func (b *Boot) Serve() (err error) {
	b.engine = grpcwrapper.Default()
	b.engine.Use(model.GRPCLogInit, nil)
	if err = b.engine.InitServer(); err != nil {
		err = xerror.Wrap(err, "bootServer.engine.InitServer")
		return
	}

	router.RegistergRPCServers(b.engine)

	serverConfig := model.GetConfig().Server

	glog.SetLogger(model.GRpcLog)
	model.GRpcLog.Infof("server run %s\n", serverConfig.Addr)
	log.Printf("server run %s\n", serverConfig.Addr)
	if err = b.engine.Run(serverConfig.Addr); err != nil {
		return
	}

	return
}

func (b *Boot) Stop() (err error) {
	b.engine.GetRawServer().Stop()
	return
}
