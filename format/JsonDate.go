package format

import (
	"fmt"
	"time"
)

// JsonTimeDateOnly 使用这个类型，可以在序列化json的时候转成yyyy-MM-dd格式的字符串
type JsonTimeDateOnly time.Time

// JsonTimeDateTime 使用这个类型，可以在序列化json的时候转成yyyy-MM-dd HH:mm:ss 格式的字符串
type JsonTimeDateTime time.Time

func (jt JsonTimeDateOnly) MarshalJSON() ([]byte, error) {
	dateOnlyStr := fmt.Sprintf("\"%s\"", time.Time(jt).Format(time.DateOnly))
	return []byte(dateOnlyStr), nil
}

func (jt JsonTimeDateTime) MarshalJSON() ([]byte, error) {
	dateOnlyStr := fmt.Sprintf("\"%s\"", time.Time(jt).Format(time.DateTime))
	return []byte(dateOnlyStr), nil
}
