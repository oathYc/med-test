package dao

import (
	"context"
	"hello/pkg/db"
	"hello/pkg/db/builder"
	"hello/repository"

	"hello/pkg/util"
)

type MedChanDetail struct {
	Base
	repository.MedChanDetail

	// search
	KeyId      int32
	ModifyTime int64
}

type MedChanShow interface {
	db.IDao
}

func NewMedChanDetail() db.IDao {
	bd := new(MedChanDetail)
	dao := repository.RegisterDao(bd)
	bd.Dao = dao
	return dao
}

func (db *MedChanDetail) BuildFilterQuery(filter db.IQuery) builder.IBuilder {
	build := db.GetBuilder()
	medChanDetailCon, ok := filter.(*MedChanDetail)
	if !ok {
		return build
	}

	// KeyId
	if keyId := medChanDetailCon.KeyId; keyId > 0 {
		build = build.Where("id > ?", keyId)
	}

	// ModifyTime
	if modifyTime := medChanDetailCon.ModifyTime; modifyTime > 0 {
		build = build.Where("created_at < ?", util.ToFormatTime(modifyTime, true))
	}

	return build
}

func (db *MedChanDetail) Count(ctx context.Context, param db.IQuery) int {
	panic("implement me")
}
