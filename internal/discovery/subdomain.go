package discovery

import (
	"fmt"
	"net"
)

func SubdomainCheck(url string) {
	// This function checks a given url for subdomains
	// by adding words from a wordlist to the url
	// https://pkg.go.dev/net - Go package net
	// func LookupHost(host string) (addrs []string, err error)
	subdomain := "werden"

	check, err := net.LookupHost(url + "/" + subdomain)
	fmt.Printf("I have checked %v for subdomain %v. Result: %v\n", url, subdomain, check)
	if err != nil {
		fmt.Printf("Problem: %v", err)
	}
}

// prüfen, ob mein Vorgehen (Parameter an LookupHost) grundsätzlich funktioniert
// Wordlist integrieren
// Schleife zum Durchlauf der Liste
// Ausgabe der positiven Ergebnisse
