package ibm

import (
	//"errors"
	"fmt"
	"testing"

	//"regexp"

	v1 "github.com/IBM-Cloud/bluemix-go/api/cis/cisv1"
	//"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccCisGlb_Basic(t *testing.T) {
	// multiple instances of this config would conflict but we only use it once
	t.Parallel()

	// cis_domain as test environment variable
	var glb v1.Glb

	//rnd := acctest.RandString(10)
	name := "ibm_cis_global_load_balancer." + "test"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		// Remove check destroy as this occurs after the CIS instance is deleted and fails with an auth error
		//CheckDestroy: testAccCheckCisGlbDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCisGlbConfigBasic("test", cis_domain),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCisGlbExists(name, &glb),
					// dont check that specified values are set, this will be evident by lack of plan diff
					// some values will get empty values
					//resource.TestCheckResourceAttr(name, "pop_pools.#", "0"),
					//resource.TestCheckResourceAttr(name, "region_pools.#", "0"),
					// also expect api to generate some values
					resource.TestCheckResourceAttr(name, "proxied", "false"), // default value
				),
			},
		},
	})
}

func TestAccCisGlb_SessionAffinity(t *testing.T) {
	t.Parallel()
	var glb v1.Glb
	//rnd := acctest.RandString(10)
	name := "ibm_cis_global_load_balancer." + "test"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		//CheckDestroy: testAccCheckCisGlbDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCisGlbConfigSessionAffinity("test", cis_domain),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCisGlbExists(name, &glb),
					// explicitly verify that our session_affinity has been set
					resource.TestCheckResourceAttr(name, "session_affinity", "cookie"),
					// dont check that other specified values are set, this will be evident by lack
					// of plan diff some values will get empty values
					//resource.TestCheckResourceAttr(name, "pop_pools.#", "0"),
					//resource.TestCheckResourceAttr(name, "region_pools.#", "0"),
				),
			},
		},
	})
}

func testAccCheckCisGlbExists(n string, glb *v1.Glb) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Load Balancer ID is set")
		}
		cisClient, err := testAccProvider.Meta().(ClientSession).CisAPI()
		foundGlb, err := cisClient.Glbs().GetGlb(rs.Primary.Attributes["cis_id"], rs.Primary.Attributes["domain_id"], rs.Primary.ID)
		if err != nil {
			return err
		}

		glb = foundGlb

		return nil
	}
}

func testAccCheckCisGlbConfigBasic(id string, cis_domain string) string {
	return testAccCheckCisPoolConfigFullySpecified(id, cis_domain) + fmt.Sprintf(`
resource "ibm_cis_global_load_balancer" "%[1]s" {
  cis_id = "${ibm_cis.instance.id}"
  domain_id = "${ibm_cis_domain.%[1]s.id}"
  name = "%[2]s"
  fallback_pool_id = "${ibm_cis_origin_pool.%[1]s.id}"
  default_pool_ids = ["${ibm_cis_origin_pool.%[1]s.id}"]
}`, id, cis_domain)
}

func testAccCheckCisGlbConfigSessionAffinity(id string, cis_domain string) string {
	return testAccCheckCisPoolConfigFullySpecified(id, cis_domain) + fmt.Sprintf(`
resource "ibm_cis_global_load_balancer" "%[1]s" {
  cis_id = "${ibm_cis.instance.id}"
  domain_id = "${ibm_cis_domain.%[1]s.id}"
  name = "%[2]s"
  fallback_pool_id = "${ibm_cis_origin_pool.%[1]s.id}"
  default_pool_ids = ["${ibm_cis_origin_pool.%[1]s.id}"]
  session_affinity = "cookie"
}`, id, cis_domain)
}
