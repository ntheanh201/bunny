package app

import (
	"github.com/spf13/viper"
	"org.idev.bunny/backend/api/enum"
	"org.idev.bunny/backend/common/logger"
)

type appConfig struct {
	Env       enum.Env `mapstructure:"ENV"`
	Port      string   `mapstructure:"PORT"`
	Dsn       string   `mapstructure:"DSN"`
	JWKsUrl   string   `mapstructure:"JWKS_URL"`
	RedisUrl  string   `mapstructure:"REDIS_URL"`
	KafkaHost string   `mapstructure:"KAFKA_HOST"`
	KafkaPort int32    `mapstructure:"KAFKA_PORT"`
}

func LoadConfig() *appConfig {

	log := logger.New("Server", "Load Config")

	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.AllowEmptyEnv(false)
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}

	appConfig := &appConfig{}
	if err := viper.Unmarshal(appConfig); err != nil {
		log.Fatal(err)
	}

	return appConfig
}