package akamai

import (
	"testing"

	appsec "github.com/akamai/AkamaiOPEN-edgegrid-golang/appsec-v1"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccAkamaiConfigurationClone_basic(t *testing.T) {
	dataSourceName := "resource_akamai_appsec_configuration_cloneconfigurationclone"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		//CheckDestroy: testAccCheckAkamaiConfigurationCloneDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAkamaiConfigurationCloneConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(dataSourceName, "id"),
				),
			},
		},
	})
}

func testAccAkamaiConfigurationCloneConfig() string {
	return `
provider "akamai" {
  appsec_section = "appsec"
}
 data "akamai_appsec_configuration" "appsecconfigedge" {
  name = "Example for EDGE"
  
}



output "configsedge" {
  value = data.akamai_appsec_configuration.appsecconfigedge.configid
}


resource "akamai_appsec_configuration_clone" "appsecconfigurationclone" {
    configid = data.akamai_appsec_configuration.appsecconfigedge.configid
    createfromversion = data.akamai_appsec_configuration.appsecconfigedge.latestversion 
    ruleupdate  = false
   }


`
}

func testAccCheckAkamaiConfigurationCloneExists(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "resource_akamai_appsec_configuration_clone" {
			continue
		}
		//rname := rs.Primary.ID
		ccresp := appsec.NewConfigurationCloneResponse()
		err = ccresp.GetConfigurationClone("TEST")

		if err != nil {
			return err
		}
	}
	return nil
}
