package boot

import (
	"hello/model"

	"git.medlinker.com/golang/xerror"
	"git.medlinker.com/golang/xlog"
	"git.medlinker.com/golang/xlog/support/logrus"
	"git.medlinker.com/service/grpcwrapper/log"
)

func bootLogger() error {
	logCfg := model.GetConfig().Log

	path := logCfg.Dir
	if path == "" {
		return xerror.New("log.dir invalid")
	}

	model.GRpcLog = logrus.NewLoggerWith(path, logCfg.Category, xlog.DailyStack)
	model.GRpcLog.SetStdPrint(logCfg.StdPrint)
	if len(logCfg.Level) > 0 {
		model.GRpcLog.SetLevelByString(logCfg.Level)
	}

	// grpc wrapper log
	log.SetLogger(model.GRpcLog)

	return nil
}
