package model

import (
	"context"
	"encoding/json"
	"unsafe"

	"git.medlinker.com/golang/xlog"
	"git.medlinker.com/service/common/utils"
	"github.com/davecgh/go-spew/spew"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

var (
	// logger
	GRpcLog xlog.XLogger
)

func GetTraceId(ctx context.Context) string {
	traceId := utils.ValueOfGRPCContext(ctx, utils.GrpcCtxTraceIdKey)
	return traceId
}

// grpc log
func GRPCLogInit(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	traceId := utils.ValueOfGRPCContext(ctx, utils.GrpcCtxTraceIdKey)
	if traceId == "" {
		traceId = utils.TraceId()
	}
	ip := utils.GetIpForGRPCContext(ctx)
	if Log != nil {
		bs, err := json.Marshal(req)
		if err != nil {
			GRpcLog.Infof("[%s][%s][req: %#v][method: %s][user: %s]", traceId, ip, req, info.FullMethod, utils.ValueOfGRPCContext(ctx, utils.GrpcCtxUserKey))
		} else {
			GRpcLog.Infof("[%s][%s][req: %s][method: %s][user: %s]", traceId, ip, *(*string)(unsafe.Pointer(&bs)), info.FullMethod, utils.ValueOfGRPCContext(ctx, utils.GrpcCtxUserKey))
		}
	}
	ctx = context.WithValue(ctx, utils.GrpcCtxTraceIdKey, traceId)
	user := utils.GetUserForGRPCContext(ctx)
	if user != nil {
		ctx = context.WithValue(ctx, utils.GrpcCtxUserKey, user)
	}
	resp, err = handler(ctx, req)
	if err != nil && Log != nil {
		GRpcLog.Infof("[%s][err: %s]", traceId, err.Error())
	}
	_ = grpc.SendHeader(ctx, metadata.Pairs(utils.GrpcCtxTraceIdKey, traceId))
	return
}

// 日志统一格式-error
func LogError(ctx context.Context, method string, msg string, errMsg string, args ...interface{}) {
	traceId := GetTraceId(ctx)
	GRpcLog.Errorf("trace_id: [%s]; method: [%s]; msg: [%s]; err: [%s]; args: [%v];", traceId, method, msg, errMsg, spew.Sdump(args...))
}

// 日志统一格式-info
func LogInfo(ctx context.Context, method string, msg string, args ...interface{}) {
	traceId := GetTraceId(ctx)
	GRpcLog.Infof("trace_id: [%s]; method: [%s]; msg: [%s]; args: [%v];", traceId, method, msg, spew.Sdump(args...))
}

// 日志统一格式-debug
func LogDebug(ctx context.Context, method string, msg string, args ...interface{}) {
	traceId := GetTraceId(ctx)
	GRpcLog.Debugf("trace_id: [%s]; method: [%s]; msg: [%s]; args: [%v];", traceId, method, msg, spew.Sdump(args...))
}

