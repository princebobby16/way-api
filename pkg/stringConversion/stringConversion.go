package stringConversion

import (
	"strconv"
	"time"
)

const (
	// dateTimeLayout is the date time layout for the datetime type.
	dateTimeLayout = "2006-01-02 15:04:05"
)

/*
ConvertDateTimeStringToUTCString converts a string version of the datetime type to a string time in ISO 8601 UTC format.
*/
func ConvertDateTimeStringToUTCString(dateTime string) (string, error) {
	if dateTime == "" {
		return dateTime, nil
	}
	middleman, err := time.Parse(dateTimeLayout, dateTime)
	if err != nil {
		return "", err
	}

	return middleman.Format(time.RFC3339), nil
}

/*
ConvertStringToInt converts a string value to an integer.
*/
func ConvertStringToInt(value string) (int, error) {
	number, err := strconv.ParseInt(value, 10, 0)
	if err != nil {
		return 0, err
	}
	return int(number), nil
}
