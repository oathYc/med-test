package dao

import (
	MedTest "hello/protocol"

	"hello/pkg/db"
	"hello/pkg/db/builder"
	"hello/repository"
)

type Base struct {
	*repository.Dao
}

func (db *Base) BuildFilterQuery(filter db.IQuery) builder.IBuilder {
	build := db.GetBuilder()
	return build
}

// BindCommonReq
func BindCommonReq(baseRequest *MedTest.BaseRequest) Base {
	return Base{
		Dao: &repository.Dao{
			CommonReq: &db.CommonReq{
				Start: baseRequest.GetStart(),
				Limit: baseRequest.GetLimit(),
				Sorts: baseRequest.GetSorts(),
			},
		},
	}
}
