package akamai

import (
	"fmt"
	"strconv"

	appsec "github.com/akamai/AkamaiOPEN-edgegrid-golang/appsec-v1"
	edge "github.com/akamai/AkamaiOPEN-edgegrid-golang/edgegrid"
	"github.com/akamai/AkamaiOPEN-edgegrid-golang/jsonhooks-v1"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceConfiguration() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceConfigurationRead,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"latestversion": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"configid": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func dataSourceConfigurationRead(d *schema.ResourceData, meta interface{}) error {
	CorrelationID := "[APPSEC][dataSourceConfigurationRead-" + CreateNonce() + "]"

	edge.PrintfCorrelation("[DEBUG]", CorrelationID, "  Read Configuration")

	//edge.PrintfCorrelation("[DEBUG]", CorrelationID, fmt.Sprintf("configid  %d version %d\n", configID, version))
	configuration := appsec.NewConfigurationResponse()
	configName := d.Get("name").(string)

	err := configuration.GetConfiguration(CorrelationID)
	if err != nil {
		edge.PrintfCorrelation("[DEBUG]", CorrelationID, fmt.Sprintf("Error  %v\n", err))
		return nil
	}

	edge.PrintfCorrelation("[DEBUG]", CorrelationID, fmt.Sprintf("Configuration   %v\n", configuration))

	jsonBody, err := jsonhooks.Marshal(configuration)
	if err != nil {
		return err
	}

	d.Set("json", string(jsonBody))

	for _, configval := range configuration.Configurations {

		edge.PrintfCorrelation("[DEBUG]", CorrelationID, fmt.Sprintf("CONFIG value  %v\n", configval.ID))
		if configval.Name == configName {

			d.Set("configid", configval.ID)
			d.Set("latestversion", configval.LatestVersion)
			d.SetId(strconv.Itoa(configval.ID))
		}
	}

	return nil
}
