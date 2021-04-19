package libdetectcloud

import (
	"net/http"
	"runtime"
	"sync"
	"time"
)

var hc = &http.Client{Timeout: 300 * time.Millisecond}

// Clouds type
type Clouds struct {
	Aws       string
	Azure     string
	Do        string
	Gce       string
	Ost       string
	Sl        string
	Vr        string
	Container string
}

// Detect function
func Detect() string {
	if runtime.GOOS != "darwin" {
		var c Clouds
		var wg sync.WaitGroup
		wg.Add(8)
		go func() {
			defer wg.Done()
			c.Aws = detectAWS()
		}()
		go func() {
			defer wg.Done()
			c.Azure = detectAzure()
		}()
		go func() {
			defer wg.Done()
			c.Do = detectDigitalOcean()
		}()
		go func() {
			defer wg.Done()
			c.Gce = detectGCE()
		}()
		go func() {
			defer wg.Done()
			c.Ost = detectOpenStack()
		}()
		go func() {
			defer wg.Done()
			c.Sl = detectSoftlayer()
		}()
		go func() {
			defer wg.Done()
			c.Vr = detectVultr()
		}()
		go func() {
			defer wg.Done()
			c.Container = detectContainer()
		}()
		wg.Wait()

		if c.Aws != "" {
			return c.Aws
		}
		if c.Azure != "" {
			return c.Azure
		}
		if c.Do != "" {
			return c.Do
		}
		if c.Gce != "" {
			return c.Gce
		}
		if c.Ost != "" {
			return c.Ost
		}
		if c.Sl != "" {
			return c.Sl
		}
		if c.Vr != "" {
			return c.Vr
		}
		if c.Container != "" {
			return c.Container
		}
	}
	return ""
}
