# detectcloud

http.Client timeout is set to `120ms`. Sometimes hitting the metadata service to fast will return unknown instead of the cloud provider detected.

    package main
    
    import (
    	"fmt"
    	"gitlab.com/taskfitio/lib/detectcloud"
    )
    
    func main() {
    
	    // detectcloud.Detect() will return a string
	    // Amazon Web Services, Microsoft Azure, Digital Ocean
	    // Google Compute Engine, OpenStack, SoftLayer, Vultr
	    // unknown

    	fmt.Println(detectcloud.Detect())
    
    }
