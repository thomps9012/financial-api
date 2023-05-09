package config

import "os"

func ENV(key string) string {
	return os.Getenv(key)
}
