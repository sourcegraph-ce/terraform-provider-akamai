package akamai

import (
	"fmt"
	"strconv"

	appsec "github.com/akamai/AkamaiOPEN-edgegrid-golang/appsec-v1"
	edge "github.com/akamai/AkamaiOPEN-edgegrid-golang/edgegrid"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// appsec v1
//
// https://developer.akamai.com/api/cloud_security/application_security/v1.html
func resourceConfigurationClone() *schema.Resource {
	return &schema.Resource{
		Create: resourceConfigurationCloneCreate,
		Read:   resourceConfigurationCloneRead,
		Update: resourceConfigurationCloneUpdate,
		Delete: resourceConfigurationCloneDelete,
		Schema: map[string]*schema.Schema{
			"configid": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"createfromversion": {
				Type:             schema.TypeInt,
				Required:         true,
				DiffSuppressFunc: suppressConfigurationCloneVersion,
			},

			"ruleupdate": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"version": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func resourceConfigurationCloneCreate(d *schema.ResourceData, meta interface{}) error {
	CorrelationID := "[APPSEC][resourceConfigurationCloneCreate-" + CreateNonce() + "]"

	edge.PrintfCorrelation("[DEBUG]", CorrelationID, " Creating ConfigurationClone")

	configurationclone := appsec.NewConfigurationCloneResponse()
	configurationclonepost := appsec.NewConfigurationClonePost()
	configurationclone.ConfigID = d.Get("configid").(int)
	configurationclonepost.CreateFromVersion = d.Get("createfromversion").(int)
	edge.PrintfCorrelation("[DEBUG]", CorrelationID, fmt.Sprintf("CREATE NEW VERSION FROM  %d\n", configurationclonepost.CreateFromVersion))
	err := configurationclone.Save(configurationclonepost, CorrelationID)
	if err != nil {
		edge.PrintfCorrelation("[DEBUG]", CorrelationID, fmt.Sprintf("Error  %v\n", err))
		return nil
	}
	//d.SetId(configurationclone.ID)
	d.Set("version", configurationclone.Version)
	return resourceConfigurationCloneRead(d, meta)
}

func resourceConfigurationCloneRead(d *schema.ResourceData, meta interface{}) error {
	CorrelationID := "[APPSEC][resourceConfigurationCloneRead-" + CreateNonce() + "]"
	edge.PrintfCorrelation("[DEBUG]", CorrelationID, "  Read ConfigurationClone")

	configurationclone := appsec.NewConfigurationCloneResponse()
	configurationclone.ConfigID = d.Get("configid").(int)
	configurationclone.Version = d.Get("createfromversion").(int)
	err := configurationclone.GetConfigurationClone(CorrelationID)
	if err != nil {
		edge.PrintfCorrelation("[DEBUG]", CorrelationID, fmt.Sprintf("Error  %v\n", err))
		return nil
	}
	edge.PrintfCorrelation("[DEBUG]", CorrelationID, fmt.Sprintf("NEW VERSION  %d\n", configurationclone.Version))
	//d.Set("version", configurationclone.Version)
	d.SetId(strconv.Itoa(configurationclone.ConfigID))
	return nil
}

func resourceConfigurationCloneDelete(d *schema.ResourceData, meta interface{}) error {
	CorrelationID := "[APPSEC][resourceConfigurationCloneDelete-" + CreateNonce() + "]"
	edge.PrintfCorrelation("[DEBUG]", CorrelationID, "  Deleting ConfigurationClone")
	return schema.Noop(d, meta)
}

func resourceConfigurationCloneUpdate(d *schema.ResourceData, meta interface{}) error {
	CorrelationID := "[APPSEC][resourceConfigurationCloneUpdate-" + CreateNonce() + "]"
	edge.PrintfCorrelation("[DEBUG]", CorrelationID, "  Updating ConfigurationClone")
	return schema.Noop(d, meta)
}
