package faker

import (
	"fmt"

	"github.com/bxcodec/faker/v3"
)

type MeetingAudit struct {
	ID                uint32 `gorm:"primary_key;column:id" json:"id"`
	Subject           string `gorm:"column:subject" json:"subject"`
	Project           string `gorm:"column:project" json:"project"`
	Type              string `gorm:"column:type" json:"type"`
	StartTime         int64  `gorm:"column:start_time" json:"start_time"`
	EndTime           int64  `gorm:"column:end_time" json:"end_time"`
	Content           string `gorm:"column:content" json:"content"`
	Speaker           string `gorm:"column:speaker" json:"speaker"`
	SpeakerPhone      string `gorm:"column:speaker_phone" json:"speaker_phone"`
	Guests            string `gorm:"column:guests" json:"guests"`
	SChannelMeetingId uint32 `gorm:"column:meeting_system_id" json:"meeting_system_id"`
	WatchUrl          string `gorm:"column:watch_url" json:"watch_url"`
	SpeakerPosterUrl  string `gorm:"column:speaker_poster_url" json:"speaker_poster_url"`
	Creator           string `gorm:"column:creator" json:"creator"`
	CreatorPhone      string `gorm:"column:creator_phone" json:"creator_phone"`
	RefuseReason      string `gorm:"column:refuse_reason" json:"refuse_reason"`
	PostTime          uint32 `gorm:"column:post_time" json:"post_time"`

	Channel          string `gorm:"column:channel" json:"channel"`
	ChannelMeetingId uint32 `gorm:"column:third_id" json:"third_id"`
	StandardStatus   uint8  `gorm:"column:standard_status" json:"standard_status"`
	RealStartTime    uint32 `gorm:"column:real_start_time" json:"real_start_time"`
	RealEndTime      uint32 `gorm:"column:real_end_time" json:"real_end_time"`
	SyncStatus       uint8  `gorm:"column:sync_status" json:"sync_status"`
	MeetingStatus    uint8  `gorm:"column:meeting_status" json:"meeting_status"`

	Status     uint8 `gorm:"column:status" json:"status"`
	CreateTime int64 `gorm:"column:create_time;autoCreateTime" json:"create_time"`
	ModifyTime int64 `gorm:"column:modify_time;autoUpdateTime" json:"modify_time"`
}

type MeetingAuditTags struct {
	ID                uint32 `faker:"long"`
	Subject           string `faker:"sentence"`
	Project           string `gorm:"column:project" json:"project"`
	Type              string `gorm:"column:type" json:"type"`
	StartTime         int64  `gorm:"column:start_time" json:"start_time"`
	EndTime           int64  `gorm:"column:end_time" json:"end_time"`
	Content           string `gorm:"column:content" json:"content"`
	Speaker           string `gorm:"column:speaker" json:"speaker"`
	SpeakerPhone      string `gorm:"column:speaker_phone" json:"speaker_phone"`
	Guests            string `gorm:"column:guests" json:"guests"`
	SChannelMeetingId uint32 `gorm:"column:meeting_system_id" json:"meeting_system_id"`
	WatchUrl          string `gorm:"column:watch_url" json:"watch_url"`
	SpeakerPosterUrl  string `gorm:"column:speaker_poster_url" json:"speaker_poster_url"`
	Creator           string `gorm:"column:creator" json:"creator"`
	CreatorPhone      string `gorm:"column:creator_phone" json:"creator_phone"`
	RefuseReason      string `gorm:"column:refuse_reason" json:"refuse_reason"`
	PostTime          uint32 `gorm:"column:post_time" json:"post_time"`

	Channel          string `gorm:"column:channel" json:"channel"`
	ChannelMeetingId uint32 `gorm:"column:third_id" json:"third_id"`
	StandardStatus   uint8  `gorm:"column:standard_status" json:"standard_status"`
	RealStartTime    uint32 `gorm:"column:real_start_time" json:"real_start_time"`
	RealEndTime      uint32 `gorm:"column:real_end_time" json:"real_end_time"`
	SyncStatus       uint8  `gorm:"column:sync_status" json:"sync_status"`
	MeetingStatus    uint8  `gorm:"column:meeting_status" json:"meeting_status"`

	Status     uint8 `gorm:"column:status" json:"status"`
	CreateTime int64 `gorm:"column:create_time;autoCreateTime" json:"create_time"`
	ModifyTime int64 `gorm:"column:modify_time;autoUpdateTime" json:"modify_time"`
}

func Example_withTags() {

	a := MeetingAuditTags{}
	err := faker.FakeData(&a)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v", a)

}
