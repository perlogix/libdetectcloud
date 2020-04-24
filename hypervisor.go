package libdetectcloud

import (
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
func BIOSInfo() (*BIOS, error) {
	rc, _, err := smbios.Stream()
	if err != nil {
		return nil, err
	}
	// Be sure to close the stream!
	defer rc.Close()

	// Decode SMBIOS structures from the stream.
	d := smbios.NewDecoder(rc)
	ss, err := d.Decode()
	if err != nil {
		return nil, err
	}
	info := []string{}
	for _, i := range ss {
		if i.Header.Type == 1 {
			info = i.Strings
		}
	}
	b := &BIOS{
		Manufacturer: info[0],
		ProductName:  info[1],
		SerialNumber: info[3],
	}
	return b, nil
}
