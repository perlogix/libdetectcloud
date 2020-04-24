package libdetectcloud

import (
	"fmt"
	"log"

	"github.com/digitalocean/go-smbios/smbios"
)

// BIOS type
type BIOS struct {
	Manufacturer string
	ProductName  string
	SerialNumber string
	IsVM         bool
	Hypervisor   string
}

//BIOSInfo function gets information from the SMBIOS
func BIOSInfo() {
	rc, _, err := smbios.Stream()
	if err != nil {
		log.Fatalf("failed to open stream: %v", err)
	}
	// Be sure to close the stream!
	defer rc.Close()

	// Decode SMBIOS structures from the stream.
	d := smbios.NewDecoder(rc)
	ss, err := d.Decode()
	if err != nil {
		log.Fatalf("failed to decode structures: %v", err)
	}
	for _, s := range ss {
		fmt.Println(s)
	}

}
