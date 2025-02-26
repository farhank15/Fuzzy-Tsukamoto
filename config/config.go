package config

import (
	"os"
)

func GetDSN() string {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = "host=" + os.Getenv("BLUEPRINT_DB_HOST") +
			" user=" + os.Getenv("BLUEPRINT_DB_USERNAME") +
			" password=" + os.Getenv("BLUEPRINT_DB_PASSWORD") +
			" dbname=" + os.Getenv("BLUEPRINT_DB_DATABASE") +
			" port=" + os.Getenv("BLUEPRINT_DB_PORT") +
			" sslmode=disable TimeZone=Asia/Shanghai"
	}
	return dsn
}
