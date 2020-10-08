package vthunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

var TEST_GSLB_SYSTEM_RESOURCE = `
resource "vthunder_gslb_system" "GslbTest" {
	ttl = "1" 
}
`

//Acceptance test
func TestAccGslbSystem_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_GSLB_SYSTEM_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("vthunder_gslb_system.GslbTest", "ttl", "1"),
				),
			},
		},
	})
}
