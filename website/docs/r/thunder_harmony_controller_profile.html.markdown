---
layout: "thunder"
page_title: "thunder: thunder_harmony_controller_profile"
sidebar_current: "docs-thunder-resource-harmony-controller-profile"
description: |-
    Provides details about thunder harmony-controller profile resource for A10
---

# thunder\_harmony\_controller\_profile

`thunder_harmony_controller_profile` provides details about harmony-controller profile
## Example Usage


```hcl
provider "thunder" {
  address  = "192.0.2.65"
  username = "admin"
  password = "admin"
}

resource "thunder_harmony_controller_profile" "profile" {
		host = "192.0.2.65"
		port=8443
		user_name = "terraform@a10networks.com"
		secret_value = "admin123"
		provider2 = "root"
		action = "register"
		use_mgmt_port = 1
		region = "US/West"
		availability_zone = "DC-1"
		thunder_mgmt_ip {
		ip_address = "192.0.2.65"
		}
}
```

## Argument Reference

* `host` - Set harmony controller host adddress
* `port` - Set port for remote Harmony Controller, default is 8443
* `user_name` - user-name for the tenant
* `secret_value` - Specify the password for the user
* `provider2` - provider for the harmony-controller
* `action` - 'register’: Register the device to the controller; 'deregister’: Deregister the device from controller;
* `use_mgmt_port` - Use management port for connections
* `region` - region of the thunder-device
* `availability_zone` - availablity zone of the thunder-device
* `thunder_mgmt_ip` - 
    * `ip_address` - IP address (IPv4 address)
