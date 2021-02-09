/**
 * © Copyright IBM Corporation 2020. All Rights Reserved.
 *
 * Licensed under the Mozilla Public License, version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at https://mozilla.org/MPL/2.0/
 */

package ibm

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccIBMMultiVlanFirewall_Basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMMultiVlanFirewallConfig_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_multi_vlan_firewall.firewall_first", "datacenter", "dal13"),
					resource.TestCheckResourceAttr(
						"ibm_multi_vlan_firewall.firewall_first", "pod", "pod01"),
					resource.TestCheckResourceAttr(
						"ibm_multi_vlan_firewall.firewall_first", "name", "Checkdelete1"),
					resource.TestCheckResourceAttr(
						"ibm_multi_vlan_firewall.firewall_first", "public_vlan_id", "2213543"),
					resource.TestCheckResourceAttr(
						"ibm_multi_vlan_firewall.firewall_first", "firewall_type", "FortiGate Security Appliance"),
					resource.TestCheckResourceAttr(
						"ibm_multi_vlan_firewall.firewall_first", "addon_configuration.#", "3"),
				),
			},
		},
	})
}

func TestAccIBMMultiVlanFirewallHA_Basic(t *testing.T) {
	t.SkipNow()
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMMultiVlanFirewallHAConfig_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_multi_vlan_firewall.firewall_first", "datacenter", "dal13"),
					resource.TestCheckResourceAttr(
						"ibm_multi_vlan_firewall.firewall_first", "pod", "pod01"),
					resource.TestCheckResourceAttr(
						"ibm_multi_vlan_firewall.firewall_first", "name", "Checkdelete1"),
					resource.TestCheckResourceAttr(
						"ibm_multi_vlan_firewall.firewall_first", "public_vlan_id", "2213543"),
					resource.TestCheckResourceAttr(
						"ibm_multi_vlan_firewall.firewall_first", "firewall_type", "FortiGate Firewall Appliance HA Option"),
					resource.TestCheckResourceAttr(
						"ibm_multi_vlan_firewall.firewall_first", "addon_configuration.#", "3"),
				),
			},
		},
	})
}
func TestAccIBMMultiVlanFirewall_InvalidFirewallType(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config:      testAccCheckIBMMultiVlanFirewallFirewallTypeConfig_InvalidFirewallType,
				ExpectError: regexp.MustCompile("must contain a value from"),
			},
		},
	})
}

const testAccCheckIBMMultiVlanFirewallConfig_basic = `
resource "ibm_multi_vlan_firewall" "firewall_first" {
	datacenter = "dal13"
	pod = "pod01"
	name = "Checkdelete1"
	firewall_type = "FortiGate Security Appliance"
	addon_configuration = ["FortiGate Security Appliance - Web Filtering Add-on","FortiGate Security Appliance - NGFW Add-on","FortiGate Security Appliance - AV Add-on"]
	}`

const testAccCheckIBMMultiVlanFirewallHAConfig_basic = `
resource "ibm_multi_vlan_firewall" "firewall_first" {
	datacenter = "dal13"
	pod = "pod01"
	name = "Checkdelete1"
	firewall_type = "FortiGate Firewall Appliance HA Option"
	addon_configuration = ["FortiGate Security Appliance - Web Filtering Add-on (High Availability)","FortiGate Security Appliance - NGFW Add-on (High Availability)","FortiGate Security Appliance - AV Add-on (High Availability)"]
	}`
const testAccCheckIBMMultiVlanFirewallFirewallTypeConfig_InvalidFirewallType = `
	resource "ibm_multi_vlan_firewall" "firewall_first" {
		datacenter = "dal13"
		pod = "pod01"
		name = "Checkdelete1"
		firewall_type = "FortiGate Security Appliance ABC"
		addon_configuration = ["FortiGate Security Appliance - Web Filtering Add-on (High Availability)","FortiGate Security Appliance - NGFW Add-on (High Availability)","FortiGate Security Appliance - AV Add-on (High Availability)"]
		}`
