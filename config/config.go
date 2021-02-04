package config

import (
	"github.com/spf13/viper"
)

// keys for database configuration.
const (
	DBName = "db.name"
	DBHost = "db.ip"
	DBPort = "db.port"
	DBUser = "db.user"
	DBPass = "db.pass"

	ServerHost = "server.host"
	ServerPort = "server.port"
)

func init() {
	// env var for db
	_ = viper.BindEnv(DBName, "DB_NAME")
	_ = viper.BindEnv(DBHost, "DB_HOST")
	_ = viper.BindEnv(DBPort, "DB_PORT")
	_ = viper.BindEnv(DBUser, "DB_USER")
	_ = viper.BindEnv(DBPass, "DB_PASS")

	// env var for server
	_ = viper.BindEnv(ServerHost, "SERVER_HOST")
	_ = viper.BindEnv(ServerPort, "SERVER_PORT")

	// defaults
	viper.SetDefault(DBName, "todo")
	viper.SetDefault(DBHost, "localhost")
	viper.SetDefault(DBPort, "27017")

	viper.SetDefault(ServerHost, "127.0.0.1")
	viper.SetDefault(ServerPort, "8080")
}
