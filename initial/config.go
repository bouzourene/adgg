package initial

import (
	"os"
	"strings"

	"github.com/bouzourene/adgg/tools"
)

var defaultGroups = []string{
	"Account Operators",
	"Administrators",
	"Backup Operators",
	"Domain Admins",
	"Domain Controllers",
	"Enterprise Admins",
	"Enterprise Read-only Domain Controllers",
	"Group Policy Creator Owners",
	"Incoming Forest Trust Builders",
	"Microsoft Exchange Servers",
	"Network Configuration Operators",
	"Power Users",
	"Print Operators",
	"Read-only Domain Controllers",
	"Replicators",
	"Schema Admins",
	"Server Operators",
}

func LoadInitialConfig() {
	logger := tools.GetLogger()

	if _, err := os.Stat("./config/groups.txt"); err == nil {
		logger.Info("Config already exists")
	} else {
		logger.Info("Creating default config")
		value := strings.Join(defaultGroups, "\n")

		os.WriteFile(
			"./config/groups.txt",
			[]byte(value),
			os.FileMode(int(0777)),
		)
	}
}
