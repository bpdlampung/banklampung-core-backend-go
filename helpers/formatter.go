package helpers

import (
	"fmt"
	"github.com/devfeel/mapper"
	"regexp"
	"strconv"
)

func Mapper(fromObj, toObj interface{}) error {
	return mapper.AutoMapper(fromObj, toObj)
}

// SanitizePhone used to sanitize Indonesia's phone number from 08 to +628.
func SanitizePhone(input string) string {
	stripExp, _ := regexp.Compile(`[^0-9]+`)
	stripped := stripExp.ReplaceAllString(input, "")

	preExp, _ := regexp.Compile(`^(\+?62|0)([0-9]*)`)
	matches := preExp.FindStringSubmatch(stripped)
	if len(matches) == 0 {
		return ""
	}
	sanitized := fmt.Sprintf("%s%s", "+62", matches[len(matches)-1])
	return sanitized
}

// UnSanitizePhone used to unsanitize Indonesia's phone number from +628 to 08.
func UnSanitizePhone(input string) string {
	stripExp, _ := regexp.Compile(`[^0-9]+`)
	stripped := stripExp.ReplaceAllString(input, "")

	preExp, _ := regexp.Compile(`^(\+?62|0)([0-9]*)`)
	matches := preExp.FindStringSubmatch(stripped)
	if len(matches) == 0 {
		return ""
	}
	sanitized := fmt.Sprintf("%s%s", "0", matches[len(matches)-1])
	return sanitized
}

// StringToInteger used to get int value from string, without try to catch an error.
func StringToInt(param string) int {
	val, _ := strconv.Atoi(param)
	return val
}

// StringToInteger used to get int value from string, without try to catch an error.
func StringToUint64(param string) uint64 {
	val, _ := strconv.ParseUint(param, 10, 64)
	return val
}
