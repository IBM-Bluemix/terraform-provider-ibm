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

func TestAccIBMISVPCRoutingTablesDataSource_basic(t *testing.T) {
	node := "data.ibm_is_vpc_routing_tables.test1"
	vpcname := fmt.Sprintf("tf-vpcname-%d", acctest.RandIntRange(100, 200))
	routetablename := fmt.Sprintf("tf-routetable-%d", acctest.RandIntRange(100, 200))
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISVPCRoutingTablesDataSourceConfig(vpcname, routetablename),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(node, "routing_tables.#"),
				),
			},
		},
	})
}

func testAccCheckIBMISVPCRoutingTablesDataSourceConfig(vpcname, routetablename string) string {
	return fmt.Sprintf(`

	resource "ibm_is_vpc" "test_custom_route_vpc" {
  		name = "%s"
	}
    
	resource "ibm_is_vpc_routing_table" "test_route_table" {
  		name = "%s"
  		vpc =  ibm_is_vpc.test_custom_route_vpc.id
	}

	data "ibm_is_vpc_routing_tables" "test1" {
		vpc =  ibm_is_vpc.test_custom_route_vpc.id
	}

	`, vpcname, routetablename)
}
