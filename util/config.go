package util

import "github.com/spf13/viper"

// Config stores all configuration of the application.
// The values are read by viper from a config file or environment variables.
type Config struct {
	APPPORT               string `mapstructure:"APP_PORT"`
	POSTGRESCONTAINERNAME string `mapstructure:"POSTGRES_CONTAINER_NAME"`
	POSTGRESUSER          string `mapstructure:"POSTGRES_USER"`
	POSTGRESPASSWORD      string `mapstructure:"POSTGRES_PASSWORD"`
	POSTGRESDB            string `mapstructure:"POSTGRES_DB"`
	POSTGRESPORT          string `mapstructure:"POSTGRES_PORT"`
}

// LoadConfig loads the configuration from a config file or environment variables.
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("local")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
