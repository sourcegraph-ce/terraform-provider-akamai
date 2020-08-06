package appsec

import (
	"time"

	"github.com/akamai/AkamaiOPEN-edgegrid-golang/edgegrid"
	"github.com/patrickmn/go-cache"
)

var (
	// Config contains the Akamai OPEN Edgegrid API credentials
	// for automatic signing of requests
	Config       edgegrid.Config
	Profilecache = cache.New(5*time.Minute, 10*time.Minute)
)

// GetConfiguration retrieves all Configuration
func GetConfiguration() (*ConfigurationResponse, error) {
	configuration := NewConfigurationResponse()
	if err := configuration.GetConfiguration(""); err != nil {
		return nil, err
	}

	return configuration, nil
}

// GetExportConfiguration retrieves all configuration
func GetExportConfiguration() (*ExportConfigurationResponse, error) {
	configuration := NewExportConfigurationResponse()
	if err := configuration.GetExportConfiguration(""); err != nil {
		return nil, err
	}

	return configuration, nil
}

// GetSelectableHostnames retrieves all SelectableHostnames
func GetSelectableHostnames() (*SelectableHostnamesResponse, error) {
	configuration := NewSelectableHostnamesResponse()
	if err := configuration.GetSelectableHostnames(""); err != nil {
		return nil, err
	}

	return configuration, nil
}
