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
const DateTimeISOFormat = "2006-01-02T15:04:05.000Z"

// time database to Asia/Jakarta
func DateTimeToString(dateTime time.Time) string {
	locName := GetEnvAndValidate("TZ")
	loc, _ := time.LoadLocation(locName)

	return dateTime.In(loc).Format(DateTimeFormat)
}

// date database to Asia/Jakarta
func DateToString(date time.Time) string {
	locName := GetEnvAndValidate("TZ")
	loc, _ := time.LoadLocation(locName)

	return date.In(loc).Format(DateFormat)
}

func StringToDate(date *string) *Date {
	if date == nil {
		return nil
	}

	loc, _ := time.LoadLocation("Asia/Jakarta")
	dateF, err := time.ParseInLocation(DateFormat, *date, loc)

	if err != nil {
		panic(err)
	}

	return &Date{
		dateF,
	}
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
	ct.Time, err = time.ParseInLocation(DateFormat, s, loc)
	return
}

func (ct *Date) MarshalJSON() ([]byte, error) {
	if ct.Time.UnixNano() == nilTime {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf("\"%s\"", ct.Time.Format(DateFormat))), nil
}
