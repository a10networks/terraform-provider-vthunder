---
layout: "vthunder"
page_title: "vthunder: vthunder_gslb_protocol"
sidebar_current: "docs-vthunder-resource-gslb-protocol"
description: |-
	Provides details about vthunder gslb protocol resource for A10
---

# vthunder\_gslb\_protocol

`vthunder_gslb_protocol` Provides details about vthunder gslb protocol
## Example Usage


```hcl
provider "vthunder" {
  address  = "129.213.82.65"
  username = "admin"
  password = "admin"
}

resource "vthunder_gslb_protocol" "GslbTest" {
	auto_detect = "1"
	enable_list {
		type = "controller"
		} 
}
```

## Argument Reference

* `auto_detect` - Automatically detect SLB Config
* `msg_format_acos_2x` - Run GSLB Protocol in compatible mode with a ACOS 2.x GSLB peer
* `ping_site` - name of site or ip address to ping
* `status_interval` - Specify GSLB Message Protocol update period (The GSLB Protocol update interval (seconds), default is 30)
* `use_mgmt_port` - Use management port for connections
* `use_mgmt_port_for_all_partitions` - Use management port for connections in all L3v Partitions
* `uuid` - uuid of the object
* `ardt_query` - Query Messages of Active RDT, default is 200 (Number)
* `ardt_response` - Response Messages of Active RDT, default is 1000 (Number)
* `ardt_session` - Sessions of Active RDT, default is 32768 (Number)
* `conn_response` - Response Messages of Connection Load, default is no limit (Number)
* `message` - Amount of Messages, default is 10000 (Number)
* `response` - Amount of Response Messages, default is 3600 (Number)
* `type` - ‘controller’: Enable/Disable GSLB protocol as GSLB controller; ‘device’: Enable/Disable GSLB protocol as site device;

