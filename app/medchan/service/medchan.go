package service

import (
	"context"
	"hello/dao"
	"hello/model"
	"hello/pkg/db"
	"hello/pkg/util"
	MedTest "hello/protocol"
	"hello/repository"
)

type MedChanService struct {
	medChannel db.IDao
}

type MedChanServer interface {
	GetChannelList(ctx context.Context, req *MedTest.MedChanListReq) (res *MedTest.MedChanListRes, err error)
}

func New() MedChanServer {
	return &MedChanService{
		medChannel: dao.NewMedChanDetail(),
	}
}

// 转时间格式
func (ds *MedChanService) toMedChan(ctx context.Context, dbMedChan *repository.MedChanDetail) (*MedTest.MedChan, error) {
	medChanData := new(MedTest.MedChan)
	if err := util.CopyStruct(medChanData, dbMedChan); nil != err {
		model.LogError(ctx, "toDoctorBroker", "CopyStruct err", err.Error(), dbMedChan)
		return medChanData, err
	}

	medChanData.CreatedAt = dbMedChan.CreatedAt.Unix()
	medChanData.UpdatedAt = dbMedChan.UpdatedAt.Unix()
	medChanData.DeletedAt = dbMedChan.DeletedAt.Unix()

	return medChanData, nil
}

// 渠道列表
func (ds *MedChanService) GetChannelList(ctx context.Context, req *MedTest.MedChanListReq) (*MedTest.MedChanListRes, error) {
	res := new(MedTest.MedChanListRes)
	list := make([]*repository.MedChanDetail, 0)

	err := ds.medChannel.Get(ctx, &dao.MedChanDetail{
		Base:       dao.BindCommonReq(req.BaseRequest),
		KeyId:      req.GetFilter().GetId(),
		ModifyTime: req.GetFilter().GetModifyTime(),
	}, &list)

	res.List = make([]*MedTest.MedChan, 0, len(list))
	for _, item := range list {
		pDoctorInfo, err := ds.toMedChan(ctx, item)
		if nil != err {
			model.LogError(ctx, "GetDoctorList", "toDoctorInfo err", err.Error(), item)
			return res, err
		}
		res.List = append(res.List, pDoctorInfo)
	}

	return res, err
}
