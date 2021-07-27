package repository

import "time"

type StoreDetail struct {
	Id          int32     `gorm:"column:id" json:"json"`
	ChannelId   int32     `gorm:"column:channel_id" json:"channel_id"`
	Title       string    `gorm:"column:title" json:"title"`
	Address     string    `gorm:"column:address" json:"address"`
	CellPhone   string    `gorm:"column:cellphone" json:"cellPhone"`
	CreatorName string    `gorm:"column:creator_name" json:"creator_name"`
	StoreStatus int32     `gorm:"column:store_status" json:"store_status"`
	CreatedAt   time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt   time.Time `gorm:"column:deleted_at" json:"deleted_at"`
	Db
}

// TableName sets the insert table name for this struct type
func (b *StoreDetail) TableName() string {
	return "store"
}
