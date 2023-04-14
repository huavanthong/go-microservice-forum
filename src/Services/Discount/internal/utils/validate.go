package utils

import (
	"errors"
	"time"
)

func ValidateTime(dateStr string) bool {
	_, err := time.Parse("02-01-2006", dateStr)
	return err == nil
}

func ValidateStartDateEndDate(startDate, endDate string) error {

	layout := "02-01-2006"

	start, err := time.Parse(layout, startDate)
	if err != nil {
		return errors.New("Invalid start date format")
	}

	end, err := time.Parse(layout, startDate)
	if err != nil {
		return errors.New("Invalid end date format")
	}

	if end.Before(start) {
		return errors.New("End date must be after start date")
	}

	return nil
}
