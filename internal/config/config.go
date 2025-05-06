// ❤ ❤ ยังไม่ได้ เติม API token นะจ๊ะ อุอิ ง่ายๆ ก็คือยังไม่สมบูรณ์ นั้นแหละ ถ้า .env แก้แล้วจะมาแก้ให้ เพราะมันดึงจาก .env เอาคราวๆประมาณนี้ไปก่อนนะ❤ ❤
package config

import (
	"os"
)

type Config struct {
	POSTGRESHOST           string
	POSTGRESDB             string
	POSTGRESUSER           string
	POSTGRESPASSWORD       string
	POSTGRESPORT           string
	PGADMINDEFAULTEMAIL    string
	PGADMINDEFAULTPASSWORD string
	PGADMINPORT            string
}

func LoadConfig() Config {
	return Config{
		POSTGRESHOST:           getEnv("POSTGRES_HOST", "localhost"),
		POSTGRESDB:             getEnv("POSTGRES_DB", "postgres"),
		POSTGRESUSER:           getEnv("POSTGRES_USER", "postgres"),
		POSTGRESPASSWORD:       getEnv("POSTGRES_PASSWORD", "postgres123"),
		POSTGRESPORT:           getEnv("POSTGRES_PORT", "5432"),
		PGADMINDEFAULTEMAIL:    getEnv("PGADMIN_DEFAULT_EMAIL", "admin@admin.com"),
		PGADMINDEFAULTPASSWORD: getEnv("PGADMIN_DEFAULT_PASSWORD", "admin123"),
		PGADMINPORT:            getEnv("PGADMIN_PORT", "5050"),
	}
}

func getEnv(key, fallback string) string {
	v := os.Getenv(key)
	if v == "" {
		return fallback
	}
	return v
}
