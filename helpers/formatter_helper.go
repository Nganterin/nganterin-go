package helpers

import (
	"fmt"
	"time"
)

func FormatFileSize(size int64) string {
	if size < 1024 {
		return fmt.Sprintf("%d B", size)
	} else if size < 1024*1024 {
		return fmt.Sprintf("%d KB", size/1024)
	}
	return fmt.Sprintf("%d MB", size/(1024*1024))
}

func FormatDateStringToTimeTime(date string) (*time.Time, error) {
	layout := "2006-01-02"

	parsedDate, err := time.Parse(layout, date)
	if err != nil {
		fmt.Println("Error parsing date:", err)
		return nil, err
	}

	return &parsedDate, nil
}

func FormatMoneyAlt(amount float64) string {
	switch {
	case amount >= 1_000_000_000:
		return fmt.Sprintf("%.2fB", amount/1_000_000_000)
	case amount >= 1_000_000:
		return fmt.Sprintf("%.2fM", amount/1_000_000)
	case amount >= 1_000:
		return fmt.Sprintf("%.2fK", amount/1_000)
	default:
		return fmt.Sprintf("%.2f", amount)
	}
}