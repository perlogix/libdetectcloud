package libdetectcloud

import (
	"regexp"

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

var hypervisors = map[string][]string{
	"VMWare":     []string{"vmware", "vm ware"},
	"Xen":        []string{"xen"},
	"Hyper-V":    []string{"microsoft", "hyperv", "hyper-v", "hyper v"},
	"VirtualBox": []string{"virtual box", "virtual-box", "virtualbox"},
	"KVM":        []string{"virt-manager", "virt manager", "virtmanager", "hvm", "kvm"},
	"Parallels":  []string{"parallels"},
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
	for h, s := range hypervisors {
		for _, i := range s {
			re, err := regexp.Compile(`(?igm).*` + i + `.*`)
			if err != nil {
				return b, nil
			}
			if re.MatchString(b.Manufacturer) || re.MatchString(b.ProductName) {
				b.IsVM = true
				b.Hypervisor = h
				return b, nil
			}
		}
	}
	return b, nil
}
