package date

import (
	"fmt"
	"strings"
	"time"
)

var nilTime = (time.Time{}).UnixNano()

const DateTimeFormat = "02-01-2006 15:04:05"
const DateFormat = "02-01-2006"
const DateTimeISOFormat = "2006-01-02T15:04:05.000Z"
const DateISOFormat = "2006-01-02"

type DateISO struct {
	time.Time
}

func (ct *DateISO) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")
	if s == "null" {
		ct.Time = time.Time{}
		return
	}
	loc, _ := time.LoadLocation("Asia/Jakarta")
	ct.Time, err = time.ParseInLocation(DateISOFormat, s, loc)
	return
}

func (ct *DateISO) MarshalJSON() ([]byte, error) {
	if ct.Time.UnixNano() == nilTime {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf("\"%s\"", ct.Time.Format(DateISOFormat))), nil
}

func (ct *DateISO) ToStringFormat() string {
	return ct.Time.Format(DateFormat)
}

func (ct *DateISO) ToStringISOFormat() string {
	return ct.Time.Format(DateISOFormat)
}

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

func (ct *Date) ToStringFormat() string {
	return ct.Time.Format(DateFormat)
}

func (ct *Date) ToStringISOFormat() string {
	return ct.Time.Format(DateISOFormat)
}

//

type DateTimeISO struct {
	time.Time
}

func (ct *DateTimeISO) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")
	if s == "null" {
		ct.Time = time.Time{}
		return
	}
	loc, _ := time.LoadLocation("Asia/Jakarta")
	ct.Time, err = time.ParseInLocation(DateTimeISOFormat, s, loc)
	return
}

func (ct *DateTimeISO) MarshalJSON() ([]byte, error) {
	if ct.Time.UnixNano() == nilTime {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf("\"%s\"", ct.Time.Format(DateTimeISOFormat))), nil
}

type DateTime struct {
	time.Time
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
