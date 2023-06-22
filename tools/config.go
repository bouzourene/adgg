package tools

import (
	"fmt"

	"github.com/spf13/viper"
)

func ConfigValue(key string) string {

	// Get logger
	logger := GetLogger()

	// Set config file as .env, whicih is an environnement file
	// that is excluded from vcs and should by copied from .env.example
	viper.SetConfigFile(".env")

	// Read config file
	err := viper.ReadInConfig()
	if err != nil {
		logger.Fatal(fmt.Sprintf("Error while reading config file %s", err.Error()))
	}

	// Trying to find the key provided through the function param
	value, ok := viper.Get(key).(string)
	if !ok {
		// If key does not exists, throw fatal error
		// We do not have (at this time) optionnal settings
		logger.Fatal(fmt.Sprintf("Connot get config key %s", key))
	}

	// Return value (k-v) as string
	return value
}
