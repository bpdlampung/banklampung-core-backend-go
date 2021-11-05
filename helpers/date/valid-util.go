package date

import (
	"fmt"
	"github.com/bpdlampung/banklampung-core-backend-go/errors"
	"github.com/bpdlampung/banklampung-core-backend-go/helpers"
	"time"
)

func TimeToString(time time.Time, format string) string {
	return time.Format(format)
}

func StringToTime(stringTime, format string, locKey *string) (dateTime time.Time) {
	locName := helpers.GetEnvAndValidate("TZ")
	loc, _ := time.LoadLocation(locName)

	if locKey != nil {
		loc, _ = time.LoadLocation(*locKey)
	}

	dateF, err := time.ParseInLocation(format, stringTime, loc)

	if err != nil {
		panic(err)
	}

	return dateF
}

func StringValidation(stringTime, format string) error {
	_, err := time.Parse(format, stringTime)

	if err != nil {
		return errors.BadRequest(fmt.Sprintf("Time not format must %s, your value %s", format, stringTime))
	}

	return nil
}
