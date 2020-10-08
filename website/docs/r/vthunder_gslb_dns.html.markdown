---
layout: "vthunder"
page_title: "vthunder: vthunder_gslb_dns"
sidebar_current: "docs-vthunder-resource-gslb-dns"
description: |-
	Provides details about vthunder gslb dns resource for A10
---

# vthunder\_gslb\_dns

`vthunder_gslb_dns` Provides details about vthunder gslb dns
## Example Usage


```hcl
provider "vthunder" {
  address  = "129.213.82.65"
  username = "admin"
  password = "admin"
}

resource "vthunder_gslb_dns" "GslbTest" {
	action = "none" 
}
```

## Argument Reference

* `action` - ‘none’: No action (default); ‘drop’: Drop query; ‘reject’: Send refuse response; ‘ignore’: Send empty response;
* `logging` - ‘none’: No logging (default); ‘query’: DNS Query; ‘response’: DNS Response; ‘both’: Both DNS Query and Response;
* `template` - Logging template (Logging Template Name)
* `uuid` - uuid of the object
* `counters1` - ‘all’: all; ‘total-query’: Total number of DNS queries received; ‘total-response’: Total number of DNS replies sent to clients; ‘bad-packet-query’: Number of queries with incorrect data length; ‘bad-packet-response’: Number of replies with incorrect data length; ‘bad-header-query’: Number of queries with incorrect header; ‘bad-header-response’: Number of replies with incorrect header; ‘bad-format-query’: Number of queries with incorrect format; ‘bad-format-response’: Number of replies with incorrect format; ‘bad-service-query’: Number of queries with unknown service; ‘bad-service-response’: Number of replies with unknown service; ‘bad-class-query’: Number of queries with incorrect class; ‘bad-class-response’: Number of replies with incorrect class; ‘bad-type-query’: Number of queries with incorrect type; ‘bad-type-response’: Number of replies with incorrect type; ‘no_answer’: Number of replies with unknown server IP;

