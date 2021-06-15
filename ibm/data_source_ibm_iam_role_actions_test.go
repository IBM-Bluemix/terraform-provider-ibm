// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMIAMRoleDataSourceAction_basic(t *testing.T) {
	serviceName := "kms"
	kmsManagerAction := "kms.instancepolicies.read"
	countActions := "1"
	name := fmt.Sprintf("Terraform%d", acctest.RandIntRange(10, 100))
	displayName := fmt.Sprintf("Terraform%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIAMRoleActionConfig(name, displayName, serviceName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_iam_role_actions.test", "service", serviceName),
					resource.TestCheckResourceAttr("ibm_iam_custom_role.customrole", "service", serviceName),
					resource.TestCheckResourceAttr("ibm_iam_custom_role.customrole", "actions.#", countActions),
					resource.TestCheckResourceAttr("ibm_iam_custom_role.customrole", "actions.0", kmsManagerAction),
				),
			},
		},
	})
}

func TestAccIBMIAMRoleDataSourceAction_withServiceSpecificRoleActions(t *testing.T) {
	serviceName := "cloud-object-storage"
	countActionsForObjectWriterRole := "12"
	name := fmt.Sprintf("Terraform%d", acctest.RandIntRange(10, 100))
	displayName := fmt.Sprintf("Terraform%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIAMCustomRoleActionConfig(name, displayName, serviceName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_iam_role_actions.test", "service", serviceName),
					resource.TestCheckResourceAttr("ibm_iam_custom_role.customrole", "service", serviceName),
					resource.TestCheckResourceAttr("ibm_iam_custom_role.customrole", "actions.#", countActionsForObjectWriterRole),
				),
			},
		},
	})
}

func testAccCheckIBMIAMRoleActionConfig(name, displayName, serviceName string) string {
	return fmt.Sprintf(`

data "ibm_iam_role_actions" "test" {
  service = "%s"
}

resource "ibm_iam_custom_role" "customrole" {
    name         = "%s"
    display_name = "%s"
    description  = "Custom Role for test scenario2"
    service = "kms"
    actions      = [data.ibm_iam_role_actions.test["Manager"].18]
}
`, serviceName, name, displayName)
}

func testAccCheckIBMIAMCustomRoleActionConfig(name, displayName, serviceName string) string {
	return fmt.Sprintf(`

data "ibm_iam_role_actions" "test" {
  service = "%s"
}

resource "ibm_iam_custom_role" "customrole" {
    name         = "%s"
    display_name = "%s"
    description  = "Custom Role for ObjectWriter"
    service = "%s"
    actions      = [data.ibm_iam_role_actions.test["Object Writer"]]
}
`, serviceName, name, displayName, serviceName)
}
