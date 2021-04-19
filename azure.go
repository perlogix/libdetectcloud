package libdetectcloud

import (
	"net/http"
)

func detectAzure() string {
	resp, err := hc.Get("http://169.254.169.254/metadata/v1/InstanceInfo")
	if err == nil && resp.StatusCode == http.StatusOK {
		return "Microsoft Azure"
	}
	return ""
}
