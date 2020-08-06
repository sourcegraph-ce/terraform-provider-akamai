package akamai

import (
	"log"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccAkamaiSelectableHostnames_basic(t *testing.T) {
	dataSourceName := "data.akamai_appsec_selectable_hostnames.appsecselectablehostnames"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAkamaiSelectableHostnamesDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAkamaiSelectableHostnamesConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(dataSourceName, "id"),
				),
			},
		},
	})
}

func testAccAkamaiSelectableHostnamesConfig() string {
	return `
provider "akamai" {
  appsec_section = "default"
}

data "akamai_appsec_config" "appsecconfigedge" {
  name = "Example for EDGE"
  
}



output "configsedge" {
  value = data.akamai_appsec_config.appsecconfigedge.configid
}

data "akamai_appsec_selectable_hostnames" "appsecselectablehostnames" {
    configid = data.akamai_appsec_config.appsecconfigedge.configid
    version = data.akamai_appsec_config.appsecconfigedge.latestversion   
}


`
}

func testAccCheckAkamaiSelectableHostnamesDestroy(s *terraform.State) error {
	log.Printf("[DEBUG] [Akamai SelectableHostnames] SelectableHostnames Destroy skipped ")
	return nil
}
