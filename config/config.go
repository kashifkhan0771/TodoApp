package config

import "github.com/spf13/viper"

const (
	DBName     = "db.name"
	DBHost     = "db.ip"
	DBUser     = "db.user"
	DBPassword = "db.password"

	ServerHost = "server.host"
	ServerPort = "server.port"

	LogLevel = "log.level"
)

func init(){
	_ = viper.BindEnv(DBName, "DB_NAME")
	_ = viper.BindEnv(DBHost, "DB_Host")
	_ = viper.BindEnv(DBUser, "DB_USER")
	_ = viper.BindEnv(DBPassword, "DB_PASSWORD")

	_ = viper.BindEnv(ServerHost, "SERVER_HOST")
	_ = viper.BindEnv(ServerPort, "SERVER_PORT")

	viper.SetDefault(DBName, "todo")
	viper.SetDefault(DBHost, "localhost:27017")

	viper.SetDefault(ServerHost, "localhost")
	viper.SetDefault(ServerPort, "8080")
}

// ReadConfig - reads config file and returns it's content
func ReadConfig(filepath string) (*viper.Viper, error) {
	config := viper.New()
	config.AddConfigPath(filepath)

	err := config.ReadInConfig()
	if err != nil {
		_, ok := err.(viper.ConfigFileNotFoundError)
		if ok {
			log().Errorf("config file not found, err: %+v", err)

			return nil, err
		}

		log().Errorf("failed to read config file, err: %+v", err)

		return nil, err
	}

	return config, nil
}
