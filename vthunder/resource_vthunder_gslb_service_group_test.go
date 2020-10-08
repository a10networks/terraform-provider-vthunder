package vthunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

var TEST_GSLB_SERVICE_GROUP_RESOURCE = `
resource "vthunder_gslb_service_group" "GslbTest" {
	service_group_name = "a"
	user_tag = "a" 
}
`

//Acceptance test
func TestAccGslbServiceGroup_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_GSLB_SERVICE_GROUP_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("vthunder_gslb_service_group.GslbTest", "service_group_name", "a"),
					resource.TestCheckResourceAttr("vthunder_gslb_service_group.GslbTest", "user_tag", "a"),
				),
			},
		},
	})
}
