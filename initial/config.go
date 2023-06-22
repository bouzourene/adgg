package initial

import (
	"os"
	"strings"

	"github.com/bouzourene/adgg/tools"
)

// List of default groups to guard.
// They are only used on the first run to
// populate the config file.
// The user is free to choose ay group
// by editing the "config/groups.txt" file.
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

// If config does not yet exist, create a basic one
func LoadInitialConfig() {

	// Get logger
	logger := tools.GetLogger()

	if _, err := os.Stat("./config/groups.txt"); err == nil {
		// If config file exists, log that information
		logger.Info("Config already exists")
	} else {
		// Else, create config file, and log that information
		logger.Info("Creating default config")
		value := strings.Join(defaultGroups, "\n")

		os.WriteFile(
			"./config/groups.txt",
			[]byte(value),
			os.FileMode(int(0777)),
		)
	}
}
