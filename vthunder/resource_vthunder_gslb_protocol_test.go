package vthunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

var TEST_GSLB_PROTOCOL_RESOURCE = `
resource "vthunder_gslb_protocol" "GslbTest" {
	auto_detect = "1"
	enable_list {
		type= "controller" 
}
}
`

//Acceptance test
func TestAccGslbProtocol_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_GSLB_PROTOCOL_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("vthunder_gslb_protocol.GslbTest", "auto_detect", "1"),
					resource.TestCheckResourceAttr("vthunder_gslb_protocol.GslbTest", "enable_list.0.type", "controller"),
				),
			},
		},
	})
}
