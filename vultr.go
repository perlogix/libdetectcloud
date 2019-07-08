package libdetectcloud

import (
	"net/http"
)

// https://www.vultr.com/metadata/
func detectVultr() string {
	resp, err := hc.Get("http://169.254.169.254/v1.json")
	if err == nil && resp.StatusCode == http.StatusOK {
		return "Vultr"
	}
	return ""
}
