---
layout: "vthunder"
page_title: "vthunder: vthunder_gslb_geo_location"
sidebar_current: "docs-vthunder-resource-gslb-geo-location"
description: |-
	Provides details about vthunder gslb geo location resource for A10
---

# vthunder\_gslb\_geo\_location

`vthunder_gslb_geo_location` Provides details about vthunder gslb geo location
## Example Usage


```hcl
provider "vthunder" {
  address  = "129.213.82.65"
  username = "admin"
  password = "admin"
}

resource "vthunder_gslb_geo_location" "GslbTest" {
	geo_locn_obj_name = "a"
	user_tag = "a" 
}
```

## Argument Reference

* `geo_locn_obj_name` - Specify geo-location name, section range is (1-15)
* `user_tag` - Customized tag
* `uuid` - uuid of the object
* `first_ip_address` - Specify IP information (Specify IP address)
* `first_ipv6_address` - Specify IPv6 address
* `geol_ipv4_mask` - Specify IPv4 mask
* `geol_ipv6_mask` - Specify IPv6 mask
* `ip_addr2` - Specify IP address range
* `ipv6_addr2` - Specify IPv6 address range

