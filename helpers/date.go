package helpers

import (
	"fmt"
	"strings"
	"time"
)

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

	locName := GetEnvAndValidate("TZ")
	loc, _ := time.LoadLocation(locName)
	dateF, err := time.ParseInLocation(DateFormat, *date, loc)

	if err != nil {
		panic(err)
	}

	return &Date{
		dateF,
	}
}

func StringToDateUTC(date *string) *Date {
	if date == nil {
		return nil
	}

	locName := GetEnvAndValidate("TZ")
	loc, _ := time.LoadLocation(locName)
	dateF, err := time.ParseInLocation(DateFormat, *date, loc)

	if err != nil {
		panic(err)
	}

	return &Date{
		dateF.UTC(),
	}
}

func StringToDateTimeUTC(date *string) *DateTime {
	if date == nil {
		return nil
	}

	locName := GetEnvAndValidate("TZ")
	loc, _ := time.LoadLocation(locName)
	dateF, err := time.ParseInLocation(DateFormat, *date, loc)

	if err != nil {
		panic(err)
	}

	return &DateTime{
		dateF.UTC(),
	}
}

func TimeNow() time.Time {
	locName := GetEnvAndValidate("TZ")
	loc, _ := time.LoadLocation(locName)

	return time.Now().In(loc)
}

func TimeNowUTC() time.Time {
	locName := GetEnvAndValidate("TZ")
	loc, _ := time.LoadLocation(locName)

	return time.Now().In(loc).UTC()
}

func StringDateNow() string {
	return TimeNow().Format(DateFormat)
}

func StringDateTimeNow() string {
	return TimeNow().Format(DateTimeFormat)
}

func DateToStringUTC(date time.Time) string {
	locName := GetEnvAndValidate("TZ")
	loc, _ := time.LoadLocation(locName)

	return date.In(loc).UTC().Format(DateTimeFormat)
}

//

func BeginningOfDay() time.Time {
	timeNow := TimeNow()

	y, m, d := timeNow.Date()
	return time.Date(y, m, d, 0, 0, 0, 0, timeNow.Location())
}

func BeginningOfDayUTC() time.Time {
	timeNow := TimeNow()

	y, m, d := timeNow.Date()
	return time.Date(y, m, d, 0, 0, 0, 0, timeNow.Location()).UTC()
}

func BeginningOfDayDateTime() DateTime {
	timeNow := TimeNow()

	y, m, d := timeNow.Date()
	return DateTime{
		time.Date(y, m, d, 0, 0, 0, 0, timeNow.Location()),
	}
}

func BeginningOfDayDateTimeUTC() DateTime {
	timeNow := TimeNow()

	y, m, d := timeNow.Date()
	return DateTime{
		time.Date(y, m, d, 0, 0, 0, 0, timeNow.Location()).UTC(),
	}
}

func EndOfDay() time.Time {
	timeNow := TimeNow()

	y, m, d := timeNow.Date()
	return time.Date(y, m, d, 23, 59, 59, int(time.Second-time.Nanosecond), timeNow.Location())
}

func EndOfDayUTC() time.Time {
	timeNow := TimeNow()

	y, m, d := timeNow.Date()
	return time.Date(y, m, d, 23, 59, 59, int(time.Second-time.Nanosecond), timeNow.Location()).UTC()
}

func EndOfDayDateTime() DateTime {
	timeNow := TimeNow()

	y, m, d := timeNow.Date()
	return DateTime{
		time.Date(y, m, d, 23, 59, 59, int(time.Second-time.Nanosecond), timeNow.Location()),
	}
}

func EndOfDayDateTimeUTC() DateTime {
	timeNow := TimeNow()

	y, m, d := timeNow.Date()
	return DateTime{
		time.Date(y, m, d, 23, 59, 59, int(time.Second-time.Nanosecond), timeNow.Location()).UTC(),
	}
}

//

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
