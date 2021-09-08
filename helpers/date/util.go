package date

import (
	"github.com/bpdlampung/banklampung-core-backend-go/helpers"
	"time"
)

func TimeNow() time.Time {
	locName := helpers.GetEnvAndValidate("TZ")
	loc, _ := time.LoadLocation(locName)

	return time.Now().In(loc)
}

func TimeNowDateFormat() Date {
	return Date{
		TimeNow(),
	}
}

func TimeNowDateISOFormat() DateISO {
	return DateISO{
		TimeNow(),
	}
}

func TimeNowDateTimeFormat() DateTime {
	return DateTime{
		TimeNow(),
	}
}

func TimeNowDateTimeISOFormat() DateTimeISO {
	return DateTimeISO{
		TimeNow(),
	}
}

func TimeNowUTC() time.Time {
	locName := helpers.GetEnvAndValidate("TZ")
	loc, _ := time.LoadLocation(locName)

	return time.Now().In(loc).UTC()
}

func TimeNowUTCDateFormat() Date {
	return Date{
		TimeNowUTC(),
	}
}

func TimeNowUTCDateISOFormat() DateISO {
	return DateISO{
		TimeNowUTC(),
	}
}

func TimeNowUTCDateTimeFormat() DateTime {
	return DateTime{
		TimeNowUTC(),
	}
}

func TimeNowUTCDateTimeISOFormat() DateTimeISO {
	return DateTimeISO{
		TimeNowUTC(),
	}
}

//

// Date -> DateFormat
// DateTime -> DateTimeFormat
// DateISO -> DateISOFormat
// DateTimeISO -> DateTimeISOFormat

func DateStringToDateFormatUTC(stringDate *string) *helpers.Date {
	if stringDate == nil {
		return nil
	}

	locName := helpers.GetEnvAndValidate("TZ")
	loc, _ := time.LoadLocation(locName)
	dateF, err := time.ParseInLocation(DateFormat, *stringDate, loc)

	if err != nil {
		panic(err)
	}

	return &helpers.Date{
		dateF.UTC(),
	}
}

func DateStringToEndOfDayDateFormatUTC(stringDate *string) *helpers.Date {
	if stringDate == nil {
		return nil
	}

	locName := helpers.GetEnvAndValidate("TZ")
	loc, _ := time.LoadLocation(locName)
	dateF, err := time.ParseInLocation(DateFormat, *stringDate, loc)

	if err != nil {
		panic(err)
	}

	y, m, d := dateF.Date()

	return &helpers.Date{
		time.Date(y, m, d, 23, 59, 59, int(time.Second-time.Nanosecond), dateF.Location()).UTC(),
	}
}
