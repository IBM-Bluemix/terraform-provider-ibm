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

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccIBMCisGLBDataSource_basic(t *testing.T) {
	node := "data.ibm_cis_global_load_balancers.test"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCisGLBDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(node, "cis_glb.0.id"),
					resource.TestCheckResourceAttrSet(node, "cis_glb.0.glb_id"),
				),
			},
		},
	})
}

func testAccCheckIBMCisGLBDataSourceConfig() string {
	// status filter defaults to empty
	return testAccCheckCisGlbConfigCisDSBasic("test", cisDomainStatic) + fmt.Sprintf(`
	data "ibm_cis_global_load_balancers" "test" {
		cis_id     = ibm_cis_global_load_balancer.test.cis_id
		domain_id  = ibm_cis_global_load_balancer.test.domain_id
	  }`)
}
