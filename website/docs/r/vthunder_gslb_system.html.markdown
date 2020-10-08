---
layout: "vthunder"
page_title: "vthunder: vthunder_gslb_system"
sidebar_current: "docs-vthunder-resource-gslb-system"
description: |-
	Provides details about vthunder gslb system resource for A10
---

# vthunder\_gslb\_system

`vthunder_gslb_system` Provides details about vthunder gslb system
## Example Usage


```hcl
provider "vthunder" {
  address  = "129.213.82.65"
  username = "admin"
  password = "admin"
}

resource "vthunder_gslb_system" "GslbTest" {
	ttl = "1" 
}
```

## Argument Reference

* `age_interval` - Interval to age runtime statistics. 0: never age, default is 10 (Time, unit: sec, default is 10)
* `geo_location_iana` - Load built-in IANA table
* `gslb_group` - GSLB Group
* `gslb_service_ip` - GSLB Service-IP
* `gslb_site` - GSLB Site
* `hostname` - System’s Network Name
* `ip_ttl` - TTL of IP packets, default is 0 (IP TTL value, default is 0)
* `module` - Specify Auto Map Module
* `slb_device` - SLB Device
* `slb_server` - SLB Server
* `slb_virtual_server` - SLB Virtual Server
* `ttl` - Specify Auto Map TTL (TTL, default is 300)
* `uuid` - uuid of the object
* `wait` - Disable GSLB until timeout if system is not ready (Time, unit: sec, default is 0)
* `geo_location_load_filename` - Specify file to be loaded
* `template_name` - CSV template to load this file

