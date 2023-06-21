package tools

import (
	"crypto/tls"
	"strings"

	"github.com/go-ldap/ldap/v3"
)

func GetLdap() *ldap.Conn {
	logger := GetLogger()

	url := ConfigValue("LDAP_URL")
	userDn := ConfigValue("LDAP_USER")
	password := ConfigValue("LDAP_PASS")

	var conn *ldap.Conn
	var err error

	url = strings.ToLower(url)
	isTls := strings.HasPrefix(url, "ldaps://")

	if isTls {
		conn, err = ldap.DialURL(url, ldap.DialWithTLSConfig(&tls.Config{
			InsecureSkipVerify: true,
		}))
	} else {
		conn, err = ldap.DialURL(url)
	}

	if err != nil {
		logger.Fatal(err.Error())
	}

	err = conn.Bind(userDn, password)
	if err != nil {
		logger.Fatal(err.Error())
	}

	return conn
}
