package vthunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

var TEST_GSLB_GEO_LOCATION_RESOURCE = `
resource "vthunder_gslb_geo_location" "GslbTest" {
	geo_locn_obj_name = "a"
	user_tag = "a" 
}
`

//Acceptance test
func TestAccGslbGeoLocation_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_GSLB_GEO_LOCATION_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("vthunder_gslb_geo_location.GslbTest", "geo_locn_obj_name", "a"),
					resource.TestCheckResourceAttr("vthunder_gslb_geo_location.GslbTest", "user_tag", "a"),
				),
			},
		},
	})
}
