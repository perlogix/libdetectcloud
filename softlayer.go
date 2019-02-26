package detectcloud

import (
	"net/http"
)

// https://sldn.softlayer.com/blog/jarteche/getting-started-user-data-and-post-provisioning-scripts
// https://github.com/bodenr/cci/wiki/SL-user-metadata
func detectSoftlayer() string {
	resp, err := hc.Get("https://api.service.softlayer.com/rest/v3/SoftLayer_Resource_Metadata/UserMetadata.txt")
	if err == nil && (resp.StatusCode == http.StatusOK || resp.StatusCode == http.StatusNotFound) {
		return "SoftLayer"
	}
	return ""
}
