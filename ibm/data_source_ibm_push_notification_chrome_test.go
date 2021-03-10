package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMDataSourcePNApplicationChrome_Basic(t *testing.T) {
	name := fmt.Sprintf("terraform_PN_%d", acctest.RandIntRange(10, 100))
	serverKey := fmt.Sprint(acctest.RandString(45))         // dummy value                       //dummy value
	websiteURL := "http://webpushnotificaton.mybluemix.net" // dummy url
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMDataSourcePNApplicationChromeConfig(name, serverKey, websiteURL),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_pn_application_chrome.chrome", "server_key"),
					resource.TestCheckResourceAttrSet("data.ibm_pn_application_chrome.chrome", "website_url"),
				),
			},
		},
	})
}

func testAccCheckIBMDataSourcePNApplicationChromeConfig(name, serverKey, websiteURL string) string {
	return fmt.Sprintf(`
		resource "ibm_resource_instance" "push_notification"{
			name     = "%s"
			location = "us-south"
			service  = "imfpush"
			plan     = "lite"
		}
		resource "ibm_pn_application_chrome" "application_chrome" {
			server_key            = "%s"
			website_url           = "%s"
			service_instance_guid = ibm_resource_instance.push_notification.guid
		}
		data "ibm_pn_application_chrome" "chrome" {
			service_instance_guid = ibm_pn_application_chrome.application_chrome.id
		}`, name, serverKey, websiteURL)
}
