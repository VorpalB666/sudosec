package scan

import (
	"fmt"
	"net"
	"os"
	"time"
)

func PortScan(url string, port int) {
	//fmt.Printf("Scanned %v for port %v\n", url, port)
	address := fmt.Sprintf("%s:%v", url, port)
	protocol := "tcp"

	// Open or creat Outputfile

	outputFile, err := os.OpenFile("port.txt",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}

	conn, err := net.DialTimeout(protocol, address, 2*time.Second)
	if err != nil {
		fmt.Printf("Port %v is closed\n", port)
		return
	} else {
		_, err := fmt.Fprintf(outputFile, "Port %d is open\n", port)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}
	}
	defer conn.Close()
	fmt.Printf("Port %v is open\n", port)
	defer outputFile.Close()
}
