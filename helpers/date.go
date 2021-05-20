package helpers

import (
	"fmt"
	"strings"
	"time"
)

type DateTime struct {
	time.Time
}

const dateTimeFormat = "2006-01-02 15:04:05"
const dateFormat = "2006-01-02"

func (ct *DateTime) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")
	if s == "null" {
		ct.Time = time.Time{}
		return
	}
	loc, _ := time.LoadLocation("Asia/Jakarta")
	ct.Time, err = time.ParseInLocation(dateTimeFormat, s, loc)
	return
}

func (ct *DateTime) MarshalJSON() ([]byte, error) {
	if ct.Time.UnixNano() == nilTime {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf("\"%s\"", ct.Time.Format(dateTimeFormat))), nil
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
	ct.Time, err = time.ParseInLocation(dateTimeFormat, s, loc)
	return
}

func (ct *Date) MarshalJSON() ([]byte, error) {
	if ct.Time.UnixNano() == nilTime {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf("\"%s\"", ct.Time.Format(dateTimeFormat))), nil
}
