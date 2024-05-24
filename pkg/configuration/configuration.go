package configuration

import (
	"fmt"

	"github.com/spf13/viper"
)

func LoadConfig(config interface{}) error {
	viper.AutomaticEnv()
	
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("Error while reading config: %w", err)
	}

	var appEnv = viper.GetString("APP_ENV")

	if appEnv != "production" {
		viper.SetConfigName("app.env." + appEnv)
		if err := viper.MergeInConfig(); err != nil {
			return fmt.Errorf("Error while merginf config for env %s %w", appEnv, err)
		}
	}

	err := viper.Unmarshal(config)
	if err != nil {
		return fmt.Errorf("Could not unmarshal config: %w", err)
	}

	return nil
}
