package tools

import (
	"fmt"

	"github.com/spf13/viper"
)

func ConfigValue(key string) string {
	logger := GetLogger()
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		logger.Fatal(fmt.Sprintf("Error while reading config file %s", err.Error()))
	}

	value, ok := viper.Get(key).(string)
	if !ok {
		logger.Fatal(fmt.Sprintf("Connot get config key %s", key))
	}

	return value
}
