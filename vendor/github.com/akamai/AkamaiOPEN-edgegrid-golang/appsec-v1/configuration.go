package appsec

import (
	"fmt"

	"github.com/akamai/AkamaiOPEN-edgegrid-golang/client-v1"
	edge "github.com/akamai/AkamaiOPEN-edgegrid-golang/edgegrid"
)

// Configuration represents a collection of Configuration
//
// See: Configuration.GetConfiguration()
// API Docs: // appsec v1
//
// https://developer.akamai.com/api/cloud_security/application_security/v1.html

type ConfigurationResponse struct {
	Configurations []struct {
		Description         string   `json:"description,omitempty"`
		FileType            string   `json:"fileType"`
		ID                  int      `json:"id"`
		LatestVersion       int      `json:"latestVersion"`
		Name                string   `json:"name,omitempty"`
		StagingVersion      int      `json:"stagingVersion,omitempty"`
		TargetProduct       string   `json:"targetProduct"`
		ProductionHostnames []string `json:"productionHostnames,omitempty"`
		ProductionVersion   int      `json:"productionVersion,omitempty"`
	} `json:"configurations"`
}

// NewConfiguration creates a new *Configuration
func NewConfigurationResponse() *ConfigurationResponse {
	return &ConfigurationResponse{}
}

// GetConfiguration populates a *Configuration with it's related Configuration
//
// API Docs: // appsec v1
//
// https://developer.akamai.com/api/cloud_security/application_security/v1.html#getconfiguration

func (configuration *ConfigurationResponse) GetConfiguration(correlationid string) error {

	req, err := client.NewRequest(
		Config,
		"GET",
		fmt.Sprintf(
			"/appsec/v1/configs",
		),
		nil,
	)
	if err != nil {
		return err
	}

	edge.PrintHttpRequestCorrelation(req, true, correlationid)

	res, err := client.Do(Config, req)
	if err != nil {
		return err
	}

	edge.PrintHttpResponseCorrelation(res, true, correlationid)

	if client.IsError(res) {
		return client.NewAPIError(res)
	}

	if err = client.BodyJSON(res, configuration); err != nil {
		return err
	}

	return nil

}
