package main

import (
	"fmt"
	"time"
)

func NextWorkday(start time.Time) time.Time {
	currentDay := start.Format("Monday")
	switch currentDay {
	case "Saturday":
		return start.AddDate(0, 0, 2)
	case "Friday":
		return start.AddDate(0, 0, 3)
	default:
		return start.AddDate(0, 0, 1)
	}
}

func TimeAgo(pastTime time.Time) string {
	duration := time.Since(pastTime)

	switch {
	case duration < time.Minute:
		seconds := int(duration.Seconds())
		if seconds == 1 {
			return "1 second ago"
		}
		return fmt.Sprintf("%d seconds ago", seconds)
	case duration < time.Hour:
		minutes := int(duration.Minutes())
		if minutes == 1 {
			return "1 minute ago"
		}
		return fmt.Sprintf("%d minutes ago", minutes)
	case duration < time.Hour*24:
		hours := int(duration.Hours())
		if hours == 1 {
			return "1 hour ago"
		}
		return fmt.Sprintf("%d hours ago", hours)
	case duration < time.Hour*24*30:
		days := int(duration.Hours() / 24)
		if days == 1 {
			return "1 day ago"
		}
		return fmt.Sprintf("%d days ago", days)
	case duration < time.Hour*24*365:
		months := int(duration.Hours() / (24 * 30))
		if months == 1 {
			return "1 month ago"
		}
		return fmt.Sprintf("%d months ago", months)
	default:
		years := int(duration.Hours() / (24 * 365))
		if years == 1 {
			return "1 year ago"
		}
		return fmt.Sprintf("%d years ago", years)
	}
}

func ParseStringToTime(dateString, format string) (time.Time, error) {
	return time.Parse(format, dateString)
}

func FormatTimeToString(timestamp time.Time, format string) string {
	return timestamp.Format(format)
}

func TimeDifference(start, end time.Time) time.Duration {
	return end.Sub(start)
}

func TickTack() {
	ticker := time.Tick(1 * time.Second)
	for t := range ticker {
		fmt.Println("Tick at", t)
	}
}

func main() {
	NextWorkday(time.Now())
}
