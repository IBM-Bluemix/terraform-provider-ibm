package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccIBMKpResource_basic(t *testing.T) {
	instanceName := fmt.Sprintf("cos_%d", acctest.RandIntRange(10, 100))
	// bucketName := fmt.Sprintf("bucket", acctest.RandIntRange(10, 100))
	keyName := fmt.Sprintf("key_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMKpResourceConfig(instanceName, keyName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_kp_key.test", "key_name", keyName),
				),
			},
		},
	})
}

func testAccCheckIBMKpResourceConfig(instanceName, KeyName string) string {
	return fmt.Sprintf(`
	resource "ibm_resource_instance" "kp_instance" {
		name              = "%s"
		service           = "kms"
		plan              = "tiered-pricing"
		location          = "us-south"
	  }
	  resource "ibm_kp_key" "test" {
		key_protect_id = "${ibm_resource_instance.kp_instance.guid}"
		key_name = "%s"
		standard_key =  true
	}
	
`, instanceName, KeyName)
}
