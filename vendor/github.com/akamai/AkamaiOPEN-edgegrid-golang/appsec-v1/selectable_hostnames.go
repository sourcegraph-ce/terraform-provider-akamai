package appsec

import (
	"fmt"

	"github.com/akamai/AkamaiOPEN-edgegrid-golang/client-v1"
	edge "github.com/akamai/AkamaiOPEN-edgegrid-golang/edgegrid"
)

// SelectableHostnames represents a collection of SelectableHostnames
//
// See: SelectableHostnames.GetSelectableHostnames()
// API Docs: // appsec v1
//
// https://developer.akamai.com/api/cloud_security/application_security/v1.html

type SelectableHostnamesResponse struct {
	AvailableSet []struct {
		ActiveInProduction     bool   `json:"activeInProduction"`
		ActiveInStaging        bool   `json:"activeInStaging"`
		ArlInclusion           bool   `json:"arlInclusion"`
		Hostname               string `json:"hostname"`
		ConfigIDInProduction   int    `json:"configIdInProduction,omitempty"`
		ConfigNameInProduction string `json:"configNameInProduction,omitempty"`
	} `json:"availableSet"`
	ConfigID                int  `json:"configId"`
	ConfigVersion           int  `json:"configVersion"`
	ProtectARLInclusionHost bool `json:"protectARLInclusionHost"`
}

// NewSelectableHostnames creates a new *SelectableHostnames
func NewSelectableHostnamesResponse() *SelectableHostnamesResponse {
	return &SelectableHostnamesResponse{}
}

// GetSelectableHostnames populates a *SelectableHostnames with it's related SelectableHostnames
//
// API Docs: // appsec v1
//
// https://developer.akamai.com/api/cloud_security/application_security/v1.html#getselectablehostnames

func (selectablehostnames *SelectableHostnamesResponse) GetSelectableHostnames(correlationid string) error {

	req, err := client.NewRequest(
		Config,
		"GET",
		fmt.Sprintf(
			"/appsec/v1/configs/%d/versions/%d/selectable-hostnames",
			selectablehostnames.ConfigID,
			selectablehostnames.ConfigVersion,
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

	if err = client.BodyJSON(res, selectablehostnames); err != nil {
		return err
	}

	return nil

}
