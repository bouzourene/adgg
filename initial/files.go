package initial

import (
	"os"

	"github.com/bouzourene/adgg/tools"
)

func CreateConfigFolder() {
	logger := tools.GetLogger()

	if _, err := os.Stat("./config"); err == nil {
		logger.Info("Config folder already exists")
	} else {
		logger.Info("Creating config folder")

		if err := os.Mkdir("./config", os.ModePerm); err != nil {
			logger.Fatal(err.Error())
		}
	}
}

func CreateDataFolder() {
	logger := tools.GetLogger()

	if _, err := os.Stat("./data"); err == nil {
		logger.Info("Data folder already exists")
	} else {
		logger.Info("Creating data folder")

		if err := os.Mkdir("./data", os.ModePerm); err != nil {
			logger.Fatal(err.Error())
		}
	}
}
