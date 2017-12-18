package ibm

import (
	"fmt"
	"testing"

	"github.com/apache/incubator-openwhisk-client-go/whisk"
	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	"github.com/IBM-Bluemix/bluemix-go/bmxerror"
)

func TestAccCloudFunctionsRule_Basic(t *testing.T) {
	var conf whisk.Rule
	name := fmt.Sprintf("terraform_%d", acctest.RandInt())
	actionName := fmt.Sprintf("terraform_%d", acctest.RandInt())
	triggerName := fmt.Sprintf("terraform_%d", acctest.RandInt())
	updatedTriggerName := fmt.Sprintf("terraform_%d", acctest.RandInt())

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudFunctionsRuleDestroy,
		Steps: []resource.TestStep{

			resource.TestStep{
				Config: testAccCheckCloudFunctionsRuleCreate(actionName, triggerName, name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckCloudFunctionsRuleExists("ibm_cloud_functions_rule.rule", &conf),
					resource.TestCheckResourceAttr("ibm_cloud_functions_rule.rule", "name", name),
					resource.TestCheckResourceAttr("ibm_cloud_functions_rule.rule", "version", "0.0.1"),
					resource.TestCheckResourceAttr("ibm_cloud_functions_rule.rule", "publish", "false"),
					resource.TestCheckResourceAttr("ibm_cloud_functions_rule.rule", "trigger_name", triggerName),
				),
			},

			resource.TestStep{
				Config: testAccCheckCloudFunctionsRuleUpdate(updatedTriggerName, name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckCloudFunctionsRuleExists("ibm_cloud_functions_rule.rule", &conf),
					resource.TestCheckResourceAttr("ibm_cloud_functions_rule.rule", "name", name),
					resource.TestCheckResourceAttr("ibm_cloud_functions_rule.rule", "version", "0.0.2"),
					resource.TestCheckResourceAttr("ibm_cloud_functions_rule.rule", "publish", "false"),
					resource.TestCheckResourceAttr("ibm_cloud_functions_rule.rule", "trigger_name", updatedTriggerName),
					resource.TestCheckResourceAttr("ibm_cloud_functions_rule.rule", "action_name", "/whisk.system/cloudant/delete-attachment"),
				),
			},
		},
	})
}

func TestAccCloudFunctionsRule_Import(t *testing.T) {
	var conf whisk.Rule
	name := fmt.Sprintf("terraform_%d", acctest.RandInt())
	triggeName := fmt.Sprintf("terraform_%d", acctest.RandInt())
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudFunctionsRuleDestroy,
		Steps: []resource.TestStep{

			resource.TestStep{
				Config: testAccCheckCloudFunctionsRuleImport(triggeName, name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckCloudFunctionsRuleExists("ibm_cloud_functions_rule.import", &conf),
					resource.TestCheckResourceAttr("ibm_cloud_functions_rule.import", "name", name),
					resource.TestCheckResourceAttr("ibm_cloud_functions_rule.import", "version", "0.0.1"),
					resource.TestCheckResourceAttr("ibm_cloud_functions_rule.import", "publish", "false"),
				),
			},

			resource.TestStep{
				ResourceName:      "ibm_cloud_functions_rule.import",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckCloudFunctionsRuleExists(n string, obj *whisk.Rule) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		client, err := testAccProvider.Meta().(ClientSession).CloudFunctionsClient()
		if err != nil {
			return err
		}
		name := rs.Primary.ID

		rule, _, err := client.Rules.Get(name)
		if err != nil {
			return err
		}

		*obj = *rule
		return nil
	}
}

func testAccCheckCloudFunctionsRuleDestroy(s *terraform.State) error {
	client, err := testAccProvider.Meta().(ClientSession).CloudFunctionsClient()
	if err != nil {
		return err
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_cloud_functions_rule" {
			continue
		}

		name := rs.Primary.ID
		_, _, err := client.Rules.Get(name)

		if err != nil {
			if apierr, ok := err.(bmxerror.RequestFailure); ok && apierr.StatusCode() != 404 {
				return fmt.Errorf("Error waiting for IBM Cloud Functions Rule (%s) to be destroyed: %s", rs.Primary.ID, err)
			}
		}
	}
	return nil
}

func testAccCheckCloudFunctionsRuleCreate(actionName, triggerName, name string) string {
	return fmt.Sprintf(`
		resource "ibm_cloud_functions_action" "action" {
			name = "%s"		  
			exec = {
			  kind = "nodejs:6"
			  code = "${file("test-fixtures/hellonode.js")}"
			}
		  }
		  resource "ibm_cloud_functions_trigger" "trigger" {
			name = "%s"
			feed = [
				{
					  name = "/whisk.system/alarms/alarm"
					  parameters = <<EOF
					[
						{
							"key":"cron",
							"value":"0 */2 * * *"
						}
					]
                EOF
			 },
		 ]

		 user_defined_annotations = <<EOF
		 [
	 {
		 "key":"sample trigger",
		 "value":"Trigger for hello action"
	 }
		 ]
		 EOF
	}
	resource "ibm_cloud_functions_rule" "rule" {
		name = "%s"
		trigger_name = "${ibm_cloud_functions_trigger.trigger.name}"
		action_name = "${ibm_cloud_functions_action.action.name}"
	  
	  }
`, actionName, triggerName, name)

}

func testAccCheckCloudFunctionsRuleUpdate(updatedTriggerName, name string) string {
	return fmt.Sprintf(`
		resource "ibm_cloud_functions_trigger" "triggerUpdated" {
			name = "%s"
		 user_defined_annotations = <<EOF
		 [
	 {
		 "key":"sample trigger",
		 "value":"Trigger for hello action"
	 }
		 ]
		 EOF
	}
	resource "ibm_cloud_functions_rule" "rule" {
		name = "%s"
		trigger_name = "${ibm_cloud_functions_trigger.triggerUpdated.name}"
		action_name = "/whisk.system/cloudant/delete-attachment"
	  
	  }
`, updatedTriggerName, name)

}

func testAccCheckCloudFunctionsRuleImport(triggerName, name string) string {
	return fmt.Sprintf(`
		resource "ibm_cloud_functions_trigger" "trigger" {
			name = "%s"
		 user_defined_annotations = <<EOF
		 [
	 {
		 "key":"sample trigger",
		 "value":"Trigger for hello action"
	 }
		 ]
		 EOF
	}
	resource "ibm_cloud_functions_rule" "import" {
		name = "%s"
		trigger_name = "${ibm_cloud_functions_trigger.trigger.name}"
		action_name = "/whisk.system/cloudant/delete-attachment"
	  
}
`, triggerName, name)

}
