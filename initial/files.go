package initial

import (
	"os"

	"github.com/bouzourene/adgg/tools"
)

// If config folder dot not yet exist, create one
func CreateConfigFolder() {

	// Get logger
	logger := tools.GetLogger()

	if _, err := os.Stat("./config"); err == nil {
		// If config folder exists, log that information
		logger.Info("Config folder already exists")
	} else {
		// Else, create config folder, and log that information
		logger.Info("Creating config folder")

		if err := os.Mkdir("./config", os.ModePerm); err != nil {
			logger.Fatal(err.Error())
		}
	}
}

// If data folder does not yet exist, create one
func CreateDataFolder() {

	// Get logger
	logger := tools.GetLogger()

	if _, err := os.Stat("./data"); err == nil {
		// If data folder exists, log that information
		logger.Info("Data folder already exists")
	} else {
		// Else, create data folder, and log that information
		logger.Info("Creating data folder")

		if err := os.Mkdir("./data", os.ModePerm); err != nil {
			logger.Fatal(err.Error())
		}
	}
}
