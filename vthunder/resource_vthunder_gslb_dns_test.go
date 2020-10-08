package vthunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

var TEST_GSLB_DNS_RESOURCE = `
resource "vthunder_gslb_dns" "GslbTest" {
	action = "none" 
}
`

//Acceptance test
func TestAccGslbDns_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_GSLB_DNS_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("vthunder_gslb_dns.GslbTest", "action", "none"),
				),
			},
		},
	})
}
