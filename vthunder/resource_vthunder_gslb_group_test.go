package vthunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

var TEST_GSLB_GROUP_RESOURCE = `
resource "vthunder_gslb_group" "GslbTest" {
	name = "a"
	user_tag = "a" 
}
`

//Acceptance test
func TestAccGslbGroup_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_GSLB_GROUP_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("vthunder_gslb_group.GslbTest", "name", "a"),
					resource.TestCheckResourceAttr("vthunder_gslb_group.GslbTest", "user_tag", "a"),
				),
			},
		},
	})
}
