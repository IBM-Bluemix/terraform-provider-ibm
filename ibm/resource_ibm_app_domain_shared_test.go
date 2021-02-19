/*
* IBM Confidential
* Object Code Only Source Materials
* 5747-SM3
* (c) Copyright IBM Corp. 2017,2021
*
* The source code for this program is not published or otherwise divested
* of its trade secrets, irrespective of what has been deposited with the
* U.S. Copyright Office.
 */

package ibm

import (
	"fmt"
	"testing"

	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/IBM-Cloud/bluemix-go/api/mccp/mccpv2"
)

func TestAccIBMAppDomainShared_Basic(t *testing.T) {
	var conf mccpv2.SharedDomainFields
	name := fmt.Sprintf("terraform%d.com", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMAppDomainShared_basic(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMAppDomainSharedExists("ibm_app_domain_shared.domain", &conf),
					resource.TestCheckResourceAttr("ibm_app_domain_shared.domain", "name", name),
				),
			},
		},
	})
}

func TestAccIBMAppDomainShared_With_Tags(t *testing.T) {
	var conf mccpv2.SharedDomainFields
	name := fmt.Sprintf("terraform%d.com", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMAppDomainShared_with_tags(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMAppDomainSharedExists("ibm_app_domain_shared.domain", &conf),
					resource.TestCheckResourceAttr("ibm_app_domain_shared.domain", "name", name),
					resource.TestCheckResourceAttr("ibm_app_domain_shared.domain", "tags.#", "2"),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMAppDomainShared_with_updated_tags(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMAppDomainSharedExists("ibm_app_domain_shared.domain", &conf),
					resource.TestCheckResourceAttr("ibm_app_domain_shared.domain", "name", name),
					resource.TestCheckResourceAttr("ibm_app_domain_shared.domain", "tags.#", "3"),
				),
			},
		},
	})
}

func testAccCheckIBMAppDomainSharedExists(n string, obj *mccpv2.SharedDomainFields) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		cfClient, err := testAccProvider.Meta().(ClientSession).MccpAPI()
		if err != nil {
			return err
		}
		sharedDomainGUID := rs.Primary.ID

		shdomain, err := cfClient.SharedDomains().Get(sharedDomainGUID)
		if err != nil {
			return err
		}

		*obj = *shdomain
		return nil
	}
}

func testAccCheckIBMAppDomainSharedDestroy(s *terraform.State) error {
	cfClient, err := testAccProvider.Meta().(ClientSession).MccpAPI()
	if err != nil {
		return err
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_app_domain_shared" {
			continue
		}

		sharedDomainGUID := rs.Primary.ID

		// Try to find the shared domain
		_, err := cfClient.SharedDomains().Get(sharedDomainGUID)

		if err == nil {
			return fmt.Errorf("CF shared domain still exists: %s", rs.Primary.ID)
		} else if !strings.Contains(err.Error(), "404") {
			return fmt.Errorf("Error waiting for CF shared domain (%s) to be destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}

func testAccCheckIBMAppDomainShared_basic(name string) string {
	return fmt.Sprintf(`
	
		resource "ibm_app_domain_shared" "domain" {
			name = "%s"
		}
	`, name)
}

func testAccCheckIBMAppDomainShared_with_tags(name string) string {
	return fmt.Sprintf(`
	
		resource "ibm_app_domain_shared" "domain" {
			name = "%s"
			tags = ["one", "two"]
		}
	`, name)
}

func testAccCheckIBMAppDomainShared_with_updated_tags(name string) string {
	return fmt.Sprintf(`
	
		resource "ibm_app_domain_shared" "domain" {
			name = "%s"
			tags = ["one", "two", "three"]
		}
	`, name)
}
