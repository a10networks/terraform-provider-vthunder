package vthunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

var TEST_GSLB_IP_LIST_RESOURCE = `
resource "vthunder_gslb_ip_list" "GslbTest" {
	gslb_ip_list_obj_name = "a"
	user_tag = "a" 
}
`

//Acceptance test
func TestAccGslbIpList_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_GSLB_IP_LIST_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("vthunder_gslb_ip_list.GslbTest", "gslb_ip_list_obj_name", "a"),
					resource.TestCheckResourceAttr("vthunder_gslb_ip_list.GslbTest", "user_tag", "a"),
				),
			},
		},
	})
}
