---
layout: "vthunder"
page_title: "vthunder: vthunder_gslb_protocol_enable"
sidebar_current: "docs-vthunder-resource-gslb-protocol-enable"
description: |-
	Provides details about vthunder gslb protocol enable resource for A10
---

# vthunder\_gslb\_protocol\_enable

`vthunder_gslb_protocol_enable` Provides details about vthunder gslb protocol enable
## Example Usage


```hcl
provider "vthunder" {
  address  = "129.213.82.65"
  username = "admin"
  password = "admin"
}

resource "vthunder_gslb_protocol_enable" "GslbProtocolTest" {
	type = "controller" 
}
```

## Argument Reference

* `type` - ‘controller’: Enable/Disable GSLB protocol as GSLB controller; ‘device’: Enable/Disable GSLB protocol as site device;
* `uuid` - uuid of the object

