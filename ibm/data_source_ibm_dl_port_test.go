// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMDLPortDataSource_basic(t *testing.T) {
	name := "dl_port"
	resName := "data.ibm_dl_port.test_dl_port"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMDLPortDataSourceConfig(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resName, "port_id"),
				),
			},
		},
	})
}

func testAccCheckIBMDLPortDataSourceConfig(name string) string {
	return fmt.Sprintf(`
	   data "ibm_dl_ports" "test_%s" {
	   }
	   data "ibm_dl_port" "test_dl_port" {
		   port_id = data.ibm_dl_ports.test_%s.ports[0].port_id
	   }
	  `, name, name)
}
