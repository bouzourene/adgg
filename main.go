package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/bouzourene/adgg/initial"
	"github.com/bouzourene/adgg/tools"
	"github.com/go-ldap/ldap/v3"
)

func main() {
	// Get logger & init LDAP connection to AD
	logger := tools.GetLogger()
	ldapConn := tools.GetLdap()

	// Create required folders & config files
	initial.CreateConfigFolder()
	initial.CreateDataFolder()
	initial.LoadInitialConfig()

	// Read the config file, which contains the groups to guard
	// File is generated auto but the user is free to add/remove groups
	// This file should never be overwritten
	readFile, err := os.Open("./config/groups.txt")
	if err != nil {
		fmt.Println(err)
	}

	// Create a buffer with the config file
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	// Split the config file in different lines
	// We'll have a slice with each group to check
	var fileLines []string
	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	// Close config file I/O, we have everything we need
	readFile.Close()

	// For each group to watch
	for _, groupCn := range fileLines {

		// Create LDAP search filter
		ldapFilter := fmt.Sprintf(
			"(&(objectClass=group)(cn=%s))",
			groupCn,
		)

		// Create LDAP search query
		searchRequest := ldap.NewSearchRequest(
			tools.ConfigValue("LDAP_BASE"), ldap.ScopeWholeSubtree,
			ldap.NeverDerefAliases, 0, 0, false,
			ldapFilter, []string{"cn", "member"}, nil,
		)

		// Run query
		sr, err := ldapConn.Search(searchRequest)
		if err != nil {
			logger.Fatal(err.Error())
		}

		// For each LDAP result
		// Should be only one, but it's easier to handle it that way
		for _, entry := range sr.Entries {

			// Get group CN, which will be our key for this group
			key := entry.GetAttributeValue("cn")

			// Get current members of this group (and sort alpha)
			members := entry.GetAttributeValues("member")
			sort.Strings(members)

			// Create a string with the member slice (each member separated by a ";")
			value := ""
			for _, member := range members {
				if value == "" {
					value = member
				} else {
					value = fmt.Sprintf("%s;%s", value, member)
				}
			}

			// Get the filename for this group
			// This file will be used to store current data
			// and detect changes
			filename := fmt.Sprintf("./data/%s.txt", key)

			// If file already exists, do a diff and alert if needed
			// Else, juste create the file with initial data
			if _, err := os.Stat(filename); err == nil {

				// Read current data file for this group
				buff, err := os.ReadFile(filename)
				if err != nil {
					logger.Fatal(err.Error())
				}

				// Convert bytes to string for easier comparaison
				currentMembers := string(buff)

				// If old data is different than new data
				// That means we detected a change and need to act on it
				if currentMembers != value {

					// Overwrite old data file with the new values
					os.WriteFile(filename, []byte(value), os.FileMode(int(0777)))

					// Get old and new string data into slices
					oldMembers := strings.Split(currentMembers, ";")
					newMembers := strings.Split(value, ";")

					// Use the difference method to detect added or removed groups from the slices
					remGroups := tools.Difference(oldMembers, newMembers)
					addGroups := tools.Difference(newMembers, oldMembers)

					// Convert the groups diff back to strings
					// We'll use this data for the warning email
					remGroupsStr := strings.Join(remGroups, "; ")
					addGroupsStr := strings.Join(addGroups, "; ")

					// Generate the mail & send it
					subject, body := tools.FormatMail(key, addGroupsStr, remGroupsStr)
					tools.SendMail(subject, body)
				}
			} else {
				// Create new file with initial data
				os.WriteFile(filename, []byte(value), os.FileMode(int(0777)))
			}
		}
	}
}
