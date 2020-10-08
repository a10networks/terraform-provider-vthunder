---
layout: "vthunder"
page_title: "vthunder: vthunder_gslb_active_rdt"
sidebar_current: "docs-vthunder-resource-gslb-active-rdt"
description: |-
	Provides details about vthunder gslb active rdt resource for A10
---

# vthunder\_gslb\_active\_rdt

`vthunder_gslb_active_rdt` Provides details about vthunder gslb active rdt
## Example Usage


```hcl
provider "vthunder" {
  address  = "129.213.82.65"
  username = "admin"
  password = "admin"
}

resource "vthunder_gslb_active_rdt" "GslbTest" {
	domain = "a"
}
```

## Argument Reference

* `domain` - Specify Query Domain (Specify Domain Name)
* `icmp` - Using ICMP
* `interval` - Specify Query Interval, unit: second, default is 1
* `port` - Specify local port to send probe packet, default is 0 (no port)
* `retry` - Specify Retry Count, default is 3
* `sleep` - Specify Sleep Time when query fail, unit: second, default is 3
* `timeout` - Specify Query Timeout, unit: msec, default is 3000
* `track` - Specify Tracking Time, unit: second, default is 60
* `uuid` - uuid of the object

