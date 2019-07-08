package libdetectcloud

import (
	"net/http"
)

// https://azure.microsoft.com/en-us/blog/what-just-happened-to-my-vm-in-vm-metadata-service/
func detectAzure() string {
	resp, err := hc.Get("http://169.254.169.254/metadata/v1/InstanceInfo")
	if err == nil && resp.StatusCode == http.StatusOK {
		return "Microsoft Azure"
	}
	return ""
}
