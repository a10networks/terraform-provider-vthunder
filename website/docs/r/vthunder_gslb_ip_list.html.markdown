---
layout: "vthunder"
page_title: "vthunder: vthunder_gslb_ip_list"
sidebar_current: "docs-vthunder-resource-gslb-ip-list"
description: |-
	Provides details about vthunder gslb ip list resource for A10
---

# vthunder\_gslb\_ip\_list

`vthunder_gslb_ip_list` Provides details about vthunder gslb ip list
## Example Usage


```hcl
provider "vthunder" {
  address  = "129.213.82.65"
  username = "admin"
  password = "admin"
}

resource "vthunder_gslb_ip_list" "GslbTest" {
	gslb_ip_list_obj_name = "a"
	user_tag = "a" 
}
```

## Argument Reference

* `gslb_ip_list_filename` - Load IP List file (IP List filename)
* `gslb_ip_list_obj_name` - Specify IP List name
* `user_tag` - Customized tag
* `uuid` - uuid of the object
* `id` - ID Number
* `ip` - Specify IP address
* `ip_mask` - IP mask

