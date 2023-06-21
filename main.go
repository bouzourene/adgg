package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"

	"github.com/bouzourene/adgg/tools"
	"github.com/go-ldap/ldap/v3"
)

func main() {
	logger := tools.GetLogger()
	ldapConn := tools.GetLdap()

	readFile, err := os.Open("./config/groups.txt")
	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var fileLines []string
	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	readFile.Close()

	for _, groupCn := range fileLines {
		searchRequest := ldap.NewSearchRequest(
			tools.ConfigValue("LDAP_BASE"),
			ldap.ScopeWholeSubtree,
			ldap.NeverDerefAliases,
			0,
			0,
			false,
			fmt.Sprintf(
				"(&(objectClass=group)(cn=%s))",
				groupCn,
			),
			[]string{"cn", "member"},
			nil,
		)

		sr, err := ldapConn.Search(searchRequest)
		if err != nil {
			logger.Fatal(err.Error())
		}

		for _, entry := range sr.Entries {
			key := entry.GetAttributeValue("cn")
			value := ""

			members := entry.GetAttributeValues("member")
			sort.Strings(members)

			for _, member := range members {
				if value == "" {
					value = member
				} else {
					value = fmt.Sprintf("%s;%s", value, member)
				}
			}

			filename := fmt.Sprintf("./data/%s.txt", key)

			if _, err := os.Stat(filename); err == nil {
				buff, err := ioutil.ReadFile(filename)
				if err != nil {
					logger.Fatal(err.Error())
				}

				currentMembers := string(buff)

				if currentMembers != value {
					ioutil.WriteFile(filename, []byte(value), os.FileMode(int(0777)))

					oldMembers := strings.Split(currentMembers, ";")
					newMembers := strings.Split(value, ";")

					remGroups := difference(oldMembers, newMembers)
					addGroups := difference(newMembers, oldMembers)

					remGroupsStr := strings.Join(remGroups, ", ")
					addGroupsStr := strings.Join(addGroups, ", ")

					tools.SendMail(
						fmt.Sprintf(
							"[ADGG] Change in AD group: %s",
							key,
						),
						fmt.Sprintf(`
Changes detected in AD group [%s]
- Members added: %s
- Members removed: %s

This mail was sent by ADGG (Active Directory Groups Guard)`,
							key,
							addGroupsStr,
							remGroupsStr,
						),
					)
				}
			} else {
				ioutil.WriteFile(filename, []byte(value), os.FileMode(int(0777)))
			}
		}
	}
}

// https://stackoverflow.com/questions/19374219/how-to-find-the-difference-between-two-slices-of-strings
// difference returns the elements in `a` that aren't in `b`.
func difference(a, b []string) []string {
	mb := make(map[string]struct{}, len(b))
	for _, x := range b {
		mb[x] = struct{}{}
	}
	var diff []string
	for _, x := range a {
		if _, found := mb[x]; !found {
			diff = append(diff, x)
		}
	}
	return diff
}
