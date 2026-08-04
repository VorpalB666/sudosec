package main

// starte aus dem Ordner recon-tool mit: "go run ./cmd/main.go"

import (
	"fmt"
	"sudosec/internal/discovery"
	"sudosec/internal/scan"
)
func main() {
	url := "example.de"
	ports := []int{20, 21, 22, 23, 25, 53, 80, 110, 119, 123, 143, 161, 194, 443}

	fmt.Println("WELCOME TO SUDOSEC")
	discovery.SubdomainCheck(url, "")

	for i := 1; i < len(ports); i++ {
		scan.PortScan(url, ports[i])
	}
}
