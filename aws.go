package detectcloud

import (
	"net/http"
)

func detectAWS() string {
	resp, err := hc.Get("http://169.254.169.254/latest/")
	if err == nil && resp.StatusCode == http.StatusOK {
		return "Amazon Web Services"
	}
	return ""
}
