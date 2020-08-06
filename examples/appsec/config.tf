provider "akamai" {
  edgerc = "~/.edgerc"
  appsec_section = "global"
}

data "akamai_appsec_configuration" "appsecconfig" {
  name = "Akamai Tools"
  
}



output "configs" {
  value = data.akamai_appsec_configuration.appsecconfig.configid
}

data "akamai_appsec_configuration" "appsecconfigedge" {
  name = "Example for EDGE"
  
}

resource "akamai_appsec_configuration_clone" "appsecconfigurationclone" {
    configid = data.akamai_appsec_configuration.appsecconfigedge.configid
    createfromversion = data.akamai_appsec_configuration.appsecconfigedge.latestversion 
    ruleupdate  = false
   }

/*
data "akamai_appsec_selectable_hostnames" "appsecselectablehostnames" {
    configid = data.akamai_appsec_config.appsecconfigedge.configid
    version = data.akamai_appsec_config.appsecconfigedge.latestversion   
}
*/

output "configsedge" {
  value = data.akamai_appsec_configuration.appsecconfigedge.configid
}

output "configsedgelatestversion" {
  value = data.akamai_appsec_configuration.appsecconfigedge.latestversion
}

/*
data "akamai_appsec_export_configuration" "export" {
  configid = 3644 //data.akamai_appsec_config.appsecconfigedge.configid
  version = 1 //data.akamai_appsec_config.appsecconfigedge.latestversion
  
}*/
/*
output "exportconfig" {
  value = data.akamai_appsec_export_configuration.export.json
}*/

resource "akamai_appsec_selected_hostnames" "appsecselectedhostnames" {
    configid = data.akamai_appsec_configuration.appsecconfigedge.configid
    version = data.akamai_appsec_configuration.appsecconfigedge.latestversion 
    hostnames = ["*.example.net","example.com","m.example.com"]  
}