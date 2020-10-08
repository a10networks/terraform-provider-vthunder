---
layout: "vthunder"
page_title: "vthunder: vthunder_gslb_service_ip"
sidebar_current: "docs-vthunder-resource-gslb-service-ip"
description: |-
	Provides details about vthunder gslb service ip resource for A10
---

# vthunder\_gslb\_service\_ip

`vthunder_gslb_service_ip` Provides details about vthunder gslb service ip
## Example Usage


```hcl
provider "vthunder" {
  address  = "129.213.82.65"
  username = "admin"
  password = "admin"
}

resource "vthunder_gslb_service_ip" "GslbTest" {
	node_name = "a"
	ipv6_address = "2003::1" 
}
```

## Argument Reference

* `action` - ‘enable’: Enable this GSLB server port; ‘disable’: Disable this GSLB server port;
* `external_ip` - External IP address for NAT
* `health_check` - Health Check Monitor (Monitor Name)
* `health_check_disable` - Disable Health Check Monitor
* `health_check_protocol_disable` - Disable GSLB Protocol Health Monitor
* `ip_address` - IP address
* `ipv6` - IPv6 address Mapping
* `ipv6_address` - IPV6 address
* `node_name` - Service-IP Name
* `user_tag` - Customized tag
* `uuid` - uuid of the object
* `follow_port_protocol` - ‘tcp’: TCP Port; ‘udp’: UDP Port;
* `health_check_follow_port` - Specify which port to follow for health status (Port Number)
* `port_num` - Port Number
* `port_proto` - ‘tcp’: TCP Port; ‘udp’: UDP Port;
* `counters1` - ‘all’: all; ‘hits’: Number of times the service IP has been selected; ‘recent’: Recent hits;

