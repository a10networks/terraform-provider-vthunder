package vthunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

var TEST_GSLB_SERVICE_IP_RESOURCE = `
resource "vthunder_gslb_service_ip" "GslbTest" {
	node_name = "a"
	ipv6_address = "2003::1" 
}
`

//Acceptance test
func TestAccGslbServiceIp_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_GSLB_SERVICE_IP_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("vthunder_gslb_service_ip.GslbTest", "node_name", "a"),
					resource.TestCheckResourceAttr("vthunder_gslb_service_ip.GslbTest", "ipv6_address", "2003::1"),
				),
			},
		},
	})
}
