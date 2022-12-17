package time_handler

import "time"

func ConvertTimeToString(date *time.Time) string {
	if date != nil {
		return date.String()
	}
	return ""
}
