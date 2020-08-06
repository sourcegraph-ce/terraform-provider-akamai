// Package papi provides a simple wrapper for the Akamai Property Manager API
package appsec

import (
	"time"

	"github.com/akamai/AkamaiOPEN-edgegrid-golang/edgegrid"
	"github.com/patrickmn/go-cache"
)

// Init sets the PAPI edgegrid Config
func Init(config edgegrid.Config) {

	Config = config
	Profilecache = cache.New(5*time.Minute, 10*time.Minute)
	edgegrid.SetupLogging()
}
