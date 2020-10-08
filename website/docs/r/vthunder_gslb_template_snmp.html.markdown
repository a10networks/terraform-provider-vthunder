---
layout: "vthunder"
page_title: "vthunder: vthunder_gslb_template_snmp"
sidebar_current: "docs-vthunder-resource-gslb-template-snmp"
description: |-
	Provides details about vthunder gslb template snmp resource for A10
---

# vthunder\_gslb\_template\_snmp

`vthunder_gslb_template_snmp` Provides details about vthunder gslb template snmp
## Example Usage


```hcl
provider "vthunder" {
  address  = "129.213.82.65"
  username = "admin"
  password = "admin"
}

resource "vthunder_gslb_template_snmp" "GslbTemplateTest" {
	snmp_name = "a"
	user_tag = "a" 
}
```

## Argument Reference

* `auth_key` - Specify authentication key (Specify key)
* `auth_proto` - ‘sha’: SHA; ‘md5’: MD5;
* `community` - Specify community for version 2c (Community name)
* `context_engine_id` - Specify context engine ID
* `context_name` - Specify context name
* `host` - Specify host (Host name or ip address)
* `interface` - Specify Interface ID
* `interval` - Specify interval, default is 3 (Interval, unit: second, default is 3)
* `oid` - Specify OID
* `port` - Specify port, default is 161 (Port Number, default is 161)
* `priv_key` - Specify privacy key (Specify key)
* `priv_proto` - ‘aes’: AES; ‘des’: DES;
* `security_engine_id` - Specify security engine ID
* `security_level` - ‘no-auth’: No authentication; ‘auth-no-priv’: Authentication, but no privacy; ‘auth-priv’: Authentication and privacy;
* `snmp_name` - Specify name of snmp template
* `user_tag` - Customized tag
* `username` - Specify username (User name)
* `uuid` - uuid of the object
* `version` - ‘v1’: Version 1; ‘v2c’: Version 2c; ‘v3’: Version 3;

