package discovery

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func SubdomainCheck(url, wordlist string) {
	var file *os.File
	var err error

	// Open default or given wordlist
	if wordlist == "" {
		file, err = os.Open("C:\\Users\\Stefa\\OneDrive\\IT\\sudosec\\recon-tool\\internal\\discovery\\shortSub.txt")
	} else {
		file, err = os.Open(wordlist)
	}

	if err != nil {
		fmt.Printf("Error: %v", err)
	}

	defer file.Close()

	// Open or creat Outputfile
	outputFile, err := os.Create("subdomains.txt")
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	defer outputFile.Close()

	// scan wordlist-file line by line and test every word included as a subdomain
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		subdomain := scanner.Text()
		host, found := SingleCheck(url, subdomain)

		// Write existing subdomains into outputFile
		if found == true {
			fmt.Printf("Found: %v\n", host)

			_, err := outputFile.WriteString(host + "\n")
			if err != nil {
				fmt.Printf("Error: %v\n", err)
			}
		}
	}

	return
}

func SingleCheck(url, subdomain string) (string, bool) {
	// This function checks a given url for subdomains
	// by adding words from a wordlist to the url
	// If found, func returns host and true
	// If not found, it returns: "no such host"
	host := subdomain + "." + url
	found := false
	_, err := net.LookupHost(host)
	//fmt.Printf("I have checked %v for subdomain %v. Result: %v\n", url, subdomain, check)
	if err != nil {
		fmt.Printf("%v\n", err)
	} else {
		found = true
	}
	return host, found
}

// prüfen, ob mein Vorgehen (Parameter an LookupHost) grundsätzlich funktioniert
// Wordlist integrieren
// Schleife zum Durchlauf der Liste
// Ausgabe der positiven Ergebnisse
