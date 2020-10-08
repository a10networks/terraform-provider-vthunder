package vthunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

var TEST_GSLB_PROTOCOL_LIMIT_RESOURCE = `
resource "vthunder_gslb_protocol_limit" "GslbProtocolTest" {
	ardt_query = "0" 
}
`

//Acceptance test
func TestAccGslbProtocolLimit_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_GSLB_PROTOCOL_LIMIT_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("vthunder_gslb_protocol_limit.GslbProtocolTest", "ardt_query", "0"),
				),
			},
		},
	})
}
