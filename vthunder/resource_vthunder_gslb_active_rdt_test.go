package vthunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

var TEST_GSLB_ACTIVE_RDT_RESOURCE = `
resource "vthunder_gslb_active_rdt" "GslbTest" {
	domain = "a"
}
`

//Acceptance test
func TestAccGslbActiveRdt_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_GSLB_ACTIVE_RDT_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("vthunder_gslb_active_rdt.GslbTest", "domain", "a"),
				),
			},
		},
	})
}
