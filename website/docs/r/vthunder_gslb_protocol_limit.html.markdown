---
layout: "vthunder"
page_title: "vthunder: vthunder_gslb_protocol_limit"
sidebar_current: "docs-vthunder-resource-gslb-protocol-limit"
description: |-
	Provides details about vthunder gslb protocol limit resource for A10
---

# vthunder\_gslb\_protocol\_limit

`vthunder_gslb_protocol_limit` Provides details about vthunder gslb protocol limit
## Example Usage


```hcl
provider "vthunder" {
  address  = "129.213.82.65"
  username = "admin"
  password = "admin"
}

resource "vthunder_gslb_protocol_limit" "GslbProtocolTest" {
	ardt_query = "0" 
}
```

## Argument Reference

* `ardt_query` - Query Messages of Active RDT, default is 200 (Number)
* `ardt_response` - Response Messages of Active RDT, default is 1000 (Number)
* `ardt_session` - Sessions of Active RDT, default is 32768 (Number)
* `conn_response` - Response Messages of Connection Load, default is no limit (Number)
* `message` - Amount of Messages, default is 10000 (Number)
* `response` - Amount of Response Messages, default is 3600 (Number)
* `uuid` - uuid of the object

