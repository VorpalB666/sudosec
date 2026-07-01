package main

// starte aus dem Ordner recon-tool mit: "go run ./cmd/main.go"

import (
	"fmt"
	"sudosec/internal/discovery"
)

func main() {
	fmt.Println("WELCOME TO SUDOSEC")
	discovery.SubdomainCheck("example.com", "")
}
