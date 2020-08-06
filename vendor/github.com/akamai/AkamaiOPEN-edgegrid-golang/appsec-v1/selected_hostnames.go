package appsec

import (
	"fmt"

	"github.com/akamai/AkamaiOPEN-edgegrid-golang/client-v1"
	edge "github.com/akamai/AkamaiOPEN-edgegrid-golang/edgegrid"
)

// SelectedHostnames represents a collection of SelectedHostnames
//
// See: SelectedHostnames.GetSelectedHostnames()
// API Docs: // appsec v1
//
// https://developer.akamai.com/api/cloud_security/application_security/v1.html

type SelectedHostnamesResponse struct {
	HostnameList []Hostname `json:"hostnameList"`
}

type Hostname struct {
	Hostname string `json:"hostname"`
}

// NewCpCodes creates a new *CpCodes
func NewSelectedHostnamesResponse() *SelectedHostnamesResponse {
	return &SelectedHostnamesResponse{}
}

// GetSelectedHostnames populates a *SelectedHostnames with it's related SelectedHostnames
//
// API Docs: // appsec v1
//
// https://developer.akamai.com/api/cloud_security/application_security/v1.html#getselectedhostnames

func (selectedhostnames *SelectedHostnamesResponse) GetSelectedHostnames(configid int, version int, correlationid string) error {

	req, err := client.NewRequest(
		Config,
		"GET",
		fmt.Sprintf(
			"/appsec/v1/configs/%d/versions/%d/selected-hostnames",
			configid,
			version,
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

	if err = client.BodyJSON(res, selectedhostnames); err != nil {
		return err
	}

	return nil

}

// Update will update a SelectedHostnames.
//
// API Docs: // appsec v1
//
// https://developer.akamai.com/api/cloud_security/application_security/v1.html#putselectedhostnames
func (selectedhostnames *SelectedHostnamesResponse) UpdateSelectedHostnames(configid int, version int, correlationid string) error {
	req, err := client.NewJSONRequest(
		Config,
		"PUT",
		fmt.Sprintf(
			"/appsec/v1/configs/%d/versions/%d/selected-hostnames",
			configid,
			version,
		),
		selectedhostnames,
		//client.JSONBody{"productId": cpcode.ProductID, "cpcodeName": cpcode.CpcodeName},
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

	return nil
}
