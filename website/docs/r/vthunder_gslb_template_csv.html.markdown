---
layout: "vthunder"
page_title: "vthunder: vthunder_gslb_template_csv"
sidebar_current: "docs-vthunder-resource-gslb-template-csv"
description: |-
	Provides details about vthunder gslb template csv resource for A10
---

# vthunder\_gslb\_template\_csv

`vthunder_gslb_template_csv` Provides details about vthunder gslb template csv
## Example Usage


```hcl
provider "vthunder" {
  address  = "129.213.82.65"
  username = "admin"
  password = "admin"
}

resource "vthunder_gslb_template_csv" "GslbTemplateTest" {
	csv_name = "a"
	user_tag = "a" 
}
```

## Argument Reference

* `csv_name` - Specify name of csv template
* `delim_char` - enter a delimiter character, default “,”
* `delim_num` - enter a delimiter number, default 44 (“,”)
* `ipv6_enable` - Support IPv6 IP ranges
* `user_tag` - Customized tag
* `uuid` - uuid of the object
* `csv_type` - ‘ip-from’: Beginning address of IP range or subnet; ‘ip-to-mask’: Ending address of IP range or Mask; ‘continent’: Continent; ‘country’: Country; ‘state’: State or province; ‘city’: City;
* `field` - Field index number (Index of Field)

