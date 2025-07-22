package utils

import (
	"errors"
	"time"
)

var formats = []string{
	"2006-01-02",
	"02.01.2006",
	"01/02/2006",
	//добавить формат "02 Jan 2006"?
}

func ParseDateOfBirth(dateStr string) (time.Time, error) {
	for _, layout := range formats {
		dob, err := time.Parse(layout, dateStr)
		if err == nil {
			return dob, nil
		}
	}
	return time.Time{}, errors.New("unsupported date format")
}
