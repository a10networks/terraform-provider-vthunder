---
layout: "vthunder"
page_title: "vthunder: vthunder_gslb_service_group"
sidebar_current: "docs-vthunder-resource-gslb-service-group"
description: |-
	Provides details about vthunder gslb service group resource for A10
---

# vthunder\_gslb\_service\_group

`vthunder_gslb_service_group` Provides details about vthunder gslb service group
## Example Usage


```hcl
provider "vthunder" {
  address  = "129.213.82.65"
  username = "admin"
  password = "admin"
}

resource "vthunder_gslb_service_group" "GslbTest" {
	service_group_name = "a"
	user_tag = "a" 
}
```

## Argument Reference

* `dependency_site` - Dependency on site
* `disable` - Disable all members
* `persistent_aging_time` - Specify aging-time, unit: min, default is 5 (Aging time)
* `persistent_ipv6_mask` - Specify IPv6 mask length, default is 128
* `persistent_mask` - Specify IP mask, default is /32
* `persistent_site` - Persistent based on site
* `service_group_name` - Specify Service Group name
* `user_tag` - Customized tag
* `uuid` - uuid of the object
* `disable_site` - Site name
* `member_name` - Service name

