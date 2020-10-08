package vthunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

var TEST_GSLB_TEMPLATE_CSV_RESOURCE = `
resource "vthunder_gslb_template_csv" "GslbTemplateTest" {
	csv_name = "a"
	user_tag = "a" 
}
`

//Acceptance test
func TestAccGslbTemplateCsv_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_GSLB_TEMPLATE_CSV_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("vthunder_gslb_template_csv.GslbTemplateTest", "csv_name", "a"),
					resource.TestCheckResourceAttr("vthunder_gslb_template_csv.GslbTemplateTest", "user_tag", "a"),
				),
			},
		},
	})
}
