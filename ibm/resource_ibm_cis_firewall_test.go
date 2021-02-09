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
	"log"
	"testing"

	"github.com/IBM/go-sdk-core/v4/core"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccIBMCisFirewall_Basic(t *testing.T) {

	var record string
	name := "ibm_cis_firewall.lockdowns"

	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMCisFirewallDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCisFirewallLockdownBasic(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMCisFirewallExists(name, &record),
					resource.TestCheckResourceAttr(
						name, "firewall_type", "lockdowns"),
					resource.TestCheckResourceAttr(
						name, "lockdown.0.configurations.0.value", "127.0.0.1"),
				),
			},
			{
				Config: testAccCheckIBMCisFirewallLockdownUpdate(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMCisFirewallExists(name, &record),
					resource.TestCheckResourceAttr(
						name, "firewall_type", "lockdowns"),
					resource.TestCheckResourceAttr(
						name, "lockdown.0.configurations.0.value", "127.0.0.3"),
				),
			},
		},
	})
}

func TestAccIBMCisFirewallAccessRuleBasic(t *testing.T) {

	var record string
	name := "ibm_cis_firewall.access_rules"

	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMCisFirewallDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCisFirewallAccessRuleBasic(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMCisFirewallExists(name, &record),
					resource.TestCheckResourceAttr(
						name, "firewall_type", "access_rules"),
					resource.TestCheckResourceAttr(
						name, "access_rule.0.configuration.0.value", "192.168.1.3"),
				),
			},
			{
				Config: testAccCheckIBMCisFirewallAccessRuleUpdate(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMCisFirewallExists(name, &record),
					resource.TestCheckResourceAttr(
						name, "firewall_type", "access_rules"),
					resource.TestCheckResourceAttr(
						name, "access_rule.0.configuration.0.value", "192.168.1.3"),
				),
			},
		},
	})
}

func TestAccIBMCisFirewallUARuleBasic(t *testing.T) {

	var record string
	name := "ibm_cis_firewall.ua_rules"

	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMCisFirewallDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCisFirewallUARuleBasic(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMCisFirewallExists(name, &record),
					resource.TestCheckResourceAttr(
						name, "firewall_type", "ua_rules"),
					resource.TestCheckResourceAttr(
						name, "ua_rule.0.mode", "block"),
				),
			},
			{
				Config: testAccCheckIBMCisFirewallUARuleUpdate(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMCisFirewallExists(name, &record),
					resource.TestCheckResourceAttr(
						name, "firewall_type", "ua_rules"),
					resource.TestCheckResourceAttr(
						name, "ua_rule.0.mode", "challenge"),
				),
			},
		},
	})
}

func TestAccIBMCisFirewallLockdown_Import(t *testing.T) {
	name := "ibm_cis_firewall.lockdowns"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCisFirewallLockdownBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						name, "firewall_type", "lockdowns"),
					resource.TestCheckResourceAttr(
						name, "lockdown.0.configurations.0.value", "127.0.0.1"),
				),
			},
			{
				ResourceName:      name,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccIBMCisFirewallAccessRule_Import(t *testing.T) {
	name := "ibm_cis_firewall.access_rules"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCisFirewallAccessRuleBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						name, "firewall_type", "access_rules"),
					resource.TestCheckResourceAttr(
						name, "access_rule.0.configuration.0.value", "192.168.1.3"),
				),
			},
			{
				ResourceName:      name,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccIBMCisFirewallUARule_Import(t *testing.T) {
	name := "ibm_cis_firewall.ua_rules"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCisFirewallUARuleBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						name, "firewall_type", "ua_rules"),
					resource.TestCheckResourceAttr(
						name, "ua_rule.0.mode", "block"),
				),
			},
			{
				ResourceName:      name,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIBMCisFirewallDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_cis_firewall" {
			continue
		}
		firewallType, lockdownID, zoneID, crn, _ := convertTfToCisFourVar(rs.Primary.ID)
		if firewallType == cisFirewallTypeLockdowns {
			// Firewall Type : Lockdowns
			cisClient, err := testAccProvider.Meta().(ClientSession).CisLockdownClientSession()
			if err != nil {
				return err
			}

			cisClient.Crn = core.StringPtr(crn)
			cisClient.ZoneIdentifier = core.StringPtr(zoneID)

			opt := cisClient.NewGetLockdownOptions(lockdownID)
			_, _, err = cisClient.GetLockdown(opt)
			if err == nil {
				return fmt.Errorf("%s type rule still exists", firewallType)
			}

		} else if firewallType == cisFirewallTypeAccessRules {

			// Firewall Type : Zone Access firewall rules
			cisClient, err := testAccProvider.Meta().(ClientSession).CisAccessRuleClientSession()
			if err != nil {
				return err
			}
			cisClient.Crn = core.StringPtr(crn)
			cisClient.ZoneIdentifier = core.StringPtr(zoneID)

			opt := cisClient.NewGetZoneAccessRuleOptions(lockdownID)
			_, _, err = cisClient.GetZoneAccessRule(opt)
			if err == nil {
				return fmt.Errorf("%s type rule still exists", firewallType)
			}

		} else if firewallType == cisFirewallTypeUARules {
			// Firewall Type: User Agent access rules
			cisClient, err := testAccProvider.Meta().(ClientSession).CisUARuleClientSession()
			if err != nil {
				return err
			}
			cisClient.Crn = core.StringPtr(crn)
			cisClient.ZoneIdentifier = core.StringPtr(zoneID)

			opt := cisClient.NewGetUserAgentRuleOptions(lockdownID)
			_, _, err = cisClient.GetUserAgentRule(opt)
			if err == nil {
				return fmt.Errorf("%s type rule still exists", firewallType)
			}
		}

	}

	return nil
}

func testAccCheckIBMCisFirewallExists(n string, tfRecordID *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Record ID is set")
		}

		tfRecord := *tfRecordID
		firewallType, lockdownID, zoneID, crn, _ := convertTfToCisFourVar(rs.Primary.ID)
		if firewallType == cisFirewallTypeLockdowns {
			// Firewall Type : Lockdowns
			cisClient, err := testAccProvider.Meta().(ClientSession).CisLockdownClientSession()
			if err != nil {
				return err
			}

			cisClient.Crn = core.StringPtr(crn)
			cisClient.ZoneIdentifier = core.StringPtr(zoneID)

			opt := cisClient.NewGetLockdownOptions(lockdownID)

			result, response, err := cisClient.GetLockdown(opt)
			if err != nil {
				log.Printf("Get zone firewall lockdown failed: %v", response)
				return err
			}
			tfRecord = convertCisToTfFourVar(firewallType, *result.Result.ID, zoneID, crn)
		} else if firewallType == cisFirewallTypeAccessRules {

			// Firewall Type : Zone Access firewall rules
			cisClient, err := testAccProvider.Meta().(ClientSession).CisAccessRuleClientSession()
			if err != nil {
				return err
			}
			cisClient.Crn = core.StringPtr(crn)
			cisClient.ZoneIdentifier = core.StringPtr(zoneID)

			opt := cisClient.NewGetZoneAccessRuleOptions(lockdownID)

			result, response, err := cisClient.GetZoneAccessRule(opt)
			if err != nil {
				log.Printf("Get zone firewall lockdown failed: %v", response)
				return err
			}
			tfRecord = convertCisToTfFourVar(firewallType, *result.Result.ID, zoneID, crn)
		} else if firewallType == cisFirewallTypeUARules {
			// Firewall Type: User Agent access rules
			cisClient, err := testAccProvider.Meta().(ClientSession).CisUARuleClientSession()
			if err != nil {
				return err
			}
			cisClient.Crn = core.StringPtr(crn)
			cisClient.ZoneIdentifier = core.StringPtr(zoneID)

			opt := cisClient.NewGetUserAgentRuleOptions(lockdownID)
			result, response, err := cisClient.GetUserAgentRule(opt)
			if err != nil {
				log.Printf("Get zone user agent rule failed: %v", response)
				return err
			}
			tfRecord = convertCisToTfFourVar(firewallType, *result.Result.ID, zoneID, crn)
		}

		if rs.Primary.ID != tfRecord {
			return fmt.Errorf("Firewall lockdown not found")
		}

		// tfRecord = convertCisToTfFourVar(firewallType, foundRecord.ID, zoneID, cisID)
		*tfRecordID = tfRecord
		return nil
	}
}

func testAccCheckIBMCisFirewallLockdownBasic() string {
	return testAccCheckIBMCisDomainDataSourceConfigBasic1() + fmt.Sprintf(`
	resource "ibm_cis_firewall" "lockdowns" {
		cis_id        = data.ibm_cis.cis.id
		domain_id     = data.ibm_cis_domain.cis_domain.id
		firewall_type = "lockdowns"
		lockdown {
			paused = "false"
			urls   = ["www.cis-terraform.com"]
			configurations {
			target = "ip"
			value  = "127.0.0.1"
			}
		}
	}`)
}

func testAccCheckIBMCisFirewallLockdownUpdate() string {
	return testAccCheckIBMCisDomainDataSourceConfigBasic1() + fmt.Sprintf(`
	resource "ibm_cis_firewall" "lockdowns" {
		cis_id        = data.ibm_cis.cis.id
		domain_id     = data.ibm_cis_domain.cis_domain.id
		firewall_type = "lockdowns"
		lockdown {
			paused = "false"
			urls   = ["www.cis-terraform.com"]
			configurations {
			target = "ip"
			value  = "127.0.0.3"
			}
		}
	}`)
}

func testAccCheckIBMCisFirewallAccessRuleBasic() string {
	return testAccCheckIBMCisDomainDataSourceConfigBasic1() + fmt.Sprintf(`
	resource "ibm_cis_firewall" "access_rules" {
		cis_id        = data.ibm_cis.cis.id
		domain_id     = data.ibm_cis_domain.cis_domain.id
		firewall_type = "access_rules"
		access_rule {
		  mode = "block"
		  notes = "access rule notes"
		  configuration {
			target = "ip"
			value  = "192.168.1.3"
		  }
		}
	}`)
}

func testAccCheckIBMCisFirewallAccessRuleUpdate() string {
	return testAccCheckIBMCisDomainDataSourceConfigBasic1() + fmt.Sprintf(`
	resource "ibm_cis_firewall" "access_rules" {
		cis_id        = data.ibm_cis.cis.id
		domain_id     = data.ibm_cis_domain.cis_domain.id
		firewall_type = "access_rules"
		access_rule {
			mode  = "block"
			notes = "access rule notes update"
			configuration {
				target = "ip"
				value  = "192.168.1.3"
			}
		}
	}`)
}

func testAccCheckIBMCisFirewallUARuleBasic() string {
	return testAccCheckIBMCisDomainDataSourceConfigBasic1() + fmt.Sprintf(`
	resource "ibm_cis_firewall" "ua_rules" {
		cis_id        = data.ibm_cis.cis.id
		domain_id     = data.ibm_cis_domain.cis_domain.id
		firewall_type = "ua_rules"
		ua_rule {
		  mode = "block"
		  configuration {
			target = "ua"
			value  = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_5) AppleWebKit/603.2.4 (KHTML, like Gecko) Version/10.1.1 Safari/603.2.4"
		  }
		}
	}`)
}

func testAccCheckIBMCisFirewallUARuleUpdate() string {
	return testAccCheckIBMCisDomainDataSourceConfigBasic1() + fmt.Sprintf(`
	resource "ibm_cis_firewall" "ua_rules" {
		cis_id        = data.ibm_cis.cis.id
		domain_id     = data.ibm_cis_domain.cis_domain.id
		firewall_type = "ua_rules"
		ua_rule {
		  mode = "challenge"
		  configuration {
			target = "ua"
			value  = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_5) AppleWebKit/603.2.4 (KHTML, like Gecko) Version/10.1.1 Safari/603.2.4"
		  }
		}
	}`)
}
