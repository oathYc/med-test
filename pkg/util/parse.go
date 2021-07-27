package util

import (
	"time"

	"github.com/jinzhu/copier"
)

func ToFormatTime(timeStamp int64, needSecond bool) string {
	layout := "2006-01-02 15:04"
	if needSecond {
		layout = layout + ":05"
	}

	timeObj := time.Unix(timeStamp, 0)
	return timeObj.Format(layout)
}

// copy struct
func CopyStruct(toValue interface{}, fromValue interface{}) (err error) {
	// jsonData, err := json.Marshal(fromValue)
	// if nil != err {
	// 	return
	// }
	// err = json.Unmarshal(jsonData, toValue)
	// if nil != err {
	// 	return
	// }
	// return
	return copier.Copy(toValue, fromValue)
}
