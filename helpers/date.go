package helpers

import (
	"fmt"
	"strings"
	"time"
)

type DateTime struct {
	time.Time
}

const DateTimeFormat = "02-01-2006 15:04:05"
const DateFormat = "02-01-2006"

// time database to Asia/Jakarta
func DateTimeToString(dateTime time.Time) string {
	locName := GetEnvAndValidate("TZ")
	loc, _ := time.LoadLocation(locName)

	return dateTime.In(loc).Format(DateTimeFormat)
}

func (ct *DateTime) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")
	if s == "null" {
		ct.Time = time.Time{}
		return
	}
	loc, _ := time.LoadLocation("Asia/Jakarta")
	ct.Time, err = time.ParseInLocation(DateTimeFormat, s, loc)
	return
}

func (ct *DateTime) MarshalJSON() ([]byte, error) {
	if ct.Time.UnixNano() == nilTime {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf("\"%s\"", ct.Time.Format(DateTimeFormat))), nil
}

var nilTime = (time.Time{}).UnixNano()

//

type Date struct {
	time.Time
}

func (ct *Date) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")
	if s == "null" {
		ct.Time = time.Time{}
		return
	}
	loc, _ := time.LoadLocation("Asia/Jakarta")
	ct.Time, err = time.ParseInLocation(DateTimeFormat, s, loc)
	return
}

func (ct *Date) MarshalJSON() ([]byte, error) {
	if ct.Time.UnixNano() == nilTime {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf("\"%s\"", ct.Time.Format(DateTimeFormat))), nil
}
