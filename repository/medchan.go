package repository

import "time"

type MedChanDetail struct {
	Id            int32     `gorm:"column:id" json:"json"`
	Title         string    `gorm:"column:title" json:"title"`
	IsSelf        int32     `gorm:"column:is_self" json:"is_self"`
	Remark        string    `gorm:"column:remark" json:"remark"`
	ChannelStatus int32     `gorm:"column:channel_status" json:"channel_status"`
	StoreCount    int32     `gorm:"column:store_count" json:"store_count"`
	CreatedAt     time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt     time.Time `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt     time.Time `gorm:"column:deleted_at" json:"deleted_at"`
	Db
}

// TableName sets the insert table name for this struct type
func (b *MedChanDetail) TableName() string {
	return "channel"
}
