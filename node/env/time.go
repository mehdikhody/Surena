package env

import "os"

func GetTimezone() string {
	timezone := os.Getenv("TZ")
	if timezone == "" {
		timezone = "UTC"
	}

	return timezone
}
