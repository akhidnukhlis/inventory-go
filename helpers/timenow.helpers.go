package helpers

import "time"

func LayoutFormat() string {
	var layoutFormat = "2006-01-02 15:04:05"

	return layoutFormat
}

func BuildTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func BuildFileName() string {
	return time.Now().Format("200601021504")
}