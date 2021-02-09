/**
 * © Copyright IBM Corporation 2020. All Rights Reserved.
 *
 * Licensed under the Mozilla Public License, version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at https://mozilla.org/MPL/2.0/
 */

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccIBMISLBSDatasource_basic(t *testing.T) {
	name := fmt.Sprintf("tflb-name-%d", acctest.RandIntRange(10, 100))
	vpcname := fmt.Sprintf("tflb-vpc-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tflb-subnet-name-%d", acctest.RandIntRange(10, 100))
	var lb string

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{

			{
				Config: testDSCheckIBMISLBSConfig(vpcname, subnetname, ISZoneName, ISCIDR, name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISLBExists("ibm_is_lb.testacc_lb", lb),
					resource.TestCheckResourceAttr(
						"data.ibm_is_lb.ds_lb", "name", name),
				),
			},
			{
				Config: testDSCheckIBMISLBSDatasourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_lbs.test_lbs", "load_balancers.0.name"),
				),
			},
		},
	})

}

func testDSCheckIBMISLBSConfig(vpcname, subnetname, zone, cidr, name string) string {
	// status filter defaults to empty
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
	  }
	  resource "ibm_is_subnet" "testacc_subnet" {
		name            = "%s"
		vpc             = ibm_is_vpc.testacc_vpc.id
		zone            = "%s"
		ipv4_cidr_block = "%s"
	  }
	  resource "ibm_is_lb" "testacc_lb" {
		name    = "%s"
		subnets = [ibm_is_subnet.testacc_subnet.id]
	  }
	  data "ibm_is_lb" "ds_lb" {
		name = ibm_is_lb.testacc_lb.name
	  }`, vpcname, subnetname, zone, cidr, name)
}
func testDSCheckIBMISLBSDatasourceConfig() string {
	// status filter defaults to empty
	return fmt.Sprintf(`
      data "ibm_is_lbs" "test_lbs" {
	  }`)
}
