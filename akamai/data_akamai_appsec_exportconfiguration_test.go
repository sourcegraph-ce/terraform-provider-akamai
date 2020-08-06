package akamai

import (
	"log"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccAkamaiExportConfiguration_basic(t *testing.T) {
	dataSourceName := "data.akamai_appsec_export_configuration.appsecexportconfiguration"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAkamaiExportConfigurationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAkamaiExportConfigurationConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(dataSourceName, "id"),
				),
			},
		},
	})
}

func testAccAkamaiExportConfigurationConfig() string {
	return `
provider "akamai" {
  appsec_section = "default"
}

data "akamai_appsec_configuration" "appsecconfigedge" {
  name = "Example for EDGE"
}

data "akamai_appsec_export_configuration" "appsecexportconfiguration" {
   configid = data.akamai_appsec_configuration.appsecconfigedge.configid
   version  = data.akamai_appsec_configuration.appsecconfigedge.latestversion 
}


`
}

func testAccCheckAkamaiExportConfigurationDestroy(s *terraform.State) error {
	log.Printf("[DEBUG] [Akamai ExportConfiguration] ExportConfiguration Destroy skipped ")
	return nil
}
