package config

import "os"

func getEnv(key string, defaultValue string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return defaultValue
}

var DatabaseHost = getEnv("MOKHTASAR_DATABASE_HOST", "db")
var DatabaseUser = getEnv("MOKHTASAR_DATABASE_USER", "root")
var DatabasePass = getEnv("MOKHTASAR_DATABASE_PASS", "changeme")
var DatabaseName = getEnv("MOKHTASAR_DATABASE_NAME", "mydb")
var DatabaseSSLMode = getEnv("MOKHTASAR_DATABASE_SSL_MODE", "disable")
