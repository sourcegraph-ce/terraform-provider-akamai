---
layout: "akamai"
page_title: "Akamai: ExportConfiguration"
sidebar_current: "docs-akamai-data-appsec-export-configuration"
description: |-
 ExportConfiguration
---

# akamai_appsec_export_configuration

Use `akamai_appsec_export_configuration` data source to retrieve a export_configuration id.

## Example Usage

Basic usage:

```hcl
provider "akamai" {
  appsec_section = "appsec"
}

data "akamai_appsec_config" "appsecconfigedge" {
  name = "Example for EDGE"
}

resource "akamai_appsec_export_configuration" "appsecexportconfiguration" {
   configid = data.akamai_appsec_config.appsecconfigedge.configid
   version  = data.akamai_appsec_config.appsecconfigedge.latestversion 
}


```

## Argument Reference

The following arguments are supported:

* `configid`- (Required) The Configuration ID

* `version` - (Required) The Version Number of configuration

# Attributes Reference

The following are the return attributes:

* `json` - Export of Configuration data

