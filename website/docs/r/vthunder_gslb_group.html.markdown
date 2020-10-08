---
layout: "vthunder"
page_title: "vthunder: vthunder_gslb_group"
sidebar_current: "docs-vthunder-resource-gslb-group"
description: |-
	Provides details about vthunder gslb group resource for A10
---

# vthunder\_gslb\_group

`vthunder_gslb_group` Provides details about vthunder gslb group
## Example Usage


```hcl
provider "vthunder" {
  address  = "129.213.82.65"
  username = "admin"
  password = "admin"
}

resource "vthunder_gslb_group" "GslbTest" {
	name = "a"
	user_tag = "a" 
}
```

## Argument Reference

* `auto_map_learn` - IP Address learned from other controller
* `auto_map_primary` - Primary Controller’s IP address
* `auto_map_smart` - Choose Best IP address
* `config_anywhere` - Every member can do config
* `config_merge` - Merge old master’s config when new one take over
* `config_save` - Accept config-save message from master
* `data_interface` - Data Interface IP Address
* `dns_discover` - Discover member via DNS Protocol
* `enable` - Join GSLB Group
* `learn` - Learn neighbour information from other controllers
* `mgmt_interface` - Management Interface IP Address
* `name` - Specify Group domain name
* `priority` - Specify Local Priority, default is 100
* `standalone` - Run GSLB Group in standalone mode
* `suffix` - Set DNS Suffix (Name)
* `user_tag` - Customized tag
* `uuid` - uuid of the object
* `primary` - Specify Primary controller’s IP address

