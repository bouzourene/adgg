package tools

import (
	"crypto/tls"
	"strings"

	"github.com/go-ldap/ldap/v3"
)

func GetLdap() *ldap.Conn {

	// Get logger
	logger := GetLogger()

	// Get LDAP settings from .env
	url := ConfigValue("LDAP_URL")
	userDn := ConfigValue("LDAP_USER")
	password := ConfigValue("LDAP_PASS")

	// Define ldap connection and error, ahead of time
	var conn *ldap.Conn
	var err error

	// Set LDAP URL to lower chars
	// Detect TLS by finding out if it starts with ldaps://
	url = strings.ToLower(url)
	isTls := strings.HasPrefix(url, "ldaps://")

	if isTls {
		// If LDAP is TLS, create TLS connection
		// Certificate check is skipped, we don't need that in this usecase anyway
		conn, err = ldap.DialURL(url, ldap.DialWithTLSConfig(&tls.Config{
			InsecureSkipVerify: true,
		}))
	} else {
		// If simple lDAP, just dial the provided URL
		conn, err = ldap.DialURL(url)
	}

	// If error, throw fatal, because we can't work without LDAP
	if err != nil {
		logger.Fatal(err.Error())
	}

	// Try to bind with userDn and password from config
	err = conn.Bind(userDn, password)
	// If error, throw fatal, because we can't work without LDAP
	if err != nil {
		logger.Fatal(err.Error())
	}

	// Return a LDAP connection
	return conn
}
