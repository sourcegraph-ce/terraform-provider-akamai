package akamai

import (
	"fmt"
	"strconv"

	appsec "github.com/akamai/AkamaiOPEN-edgegrid-golang/appsec-v1"
	edge "github.com/akamai/AkamaiOPEN-edgegrid-golang/edgegrid"
	"github.com/akamai/AkamaiOPEN-edgegrid-golang/jsonhooks-v1"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceExportConfiguration() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceExportConfigurationRead,
		Schema: map[string]*schema.Schema{
			"configid": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"version": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"json": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "JSON Export representation",
			},
		},
	}
}

func dataSourceExportConfigurationRead(d *schema.ResourceData, meta interface{}) error {
	CorrelationID := "[APPSEC][dataSourceExportConfigurationRead-" + CreateNonce() + "]"

	edge.PrintfCorrelation("[DEBUG]", CorrelationID, "  Read ExportConfiguration")

	//edge.PrintfCorrelation("[DEBUG]", CorrelationID, fmt.Sprintf("configid  %d version %d\n", configID, version))
	exportconfiguration := appsec.NewExportConfigurationResponse()
	exportconfiguration.ConfigID = d.Get("configid").(int)
	exportconfiguration.Version = d.Get("version").(int)

	err := exportconfiguration.GetExportConfiguration(CorrelationID)
	if err != nil {
		edge.PrintfCorrelation("[DEBUG]", CorrelationID, fmt.Sprintf("Error  %v\n", err))
		return nil
	}

	edge.PrintfCorrelation("[DEBUG]", CorrelationID, fmt.Sprintf("ExportConfiguration   %v\n", exportconfiguration))

	jsonBody, err := jsonhooks.Marshal(exportconfiguration)
	if err != nil {
		return err
	}

	d.Set("json", string(jsonBody))

	d.SetId(strconv.Itoa(exportconfiguration.ConfigID))

	return nil
}
