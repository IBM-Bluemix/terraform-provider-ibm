package ibm

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/softlayer/softlayer-go/services"
)

func resourceIBMNetworkVlanSpan() *schema.Resource {
	return &schema.Resource{
		Create:   resourceIBMNetworkVlanSpanCreate,
		Read:     resourceIBMNetworkVlanSpanRead,
		Update:   resourceIBMNetworkVlanSpanUpdate,
		Delete:   resourceIBMNetworkVlanSpanDelete,
		Importer: &schema.ResourceImporter{},

		Timeouts: &schema.ResourceTimeout{
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"vlan_spanning": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateAllowedStringValue([]string{"off", "on"}),
			},
		},
	}
}

func resourceIBMNetworkVlanSpanRead(d *schema.ResourceData, meta interface{}) error {
	sess := meta.(ClientSession).SoftLayerSession()
	service := services.GetAccountService(sess)

	vlanSpan, err := service.GetNetworkVlanSpan()

	if err != nil {
		return fmt.Errorf("Error retrieving vlan: %s", err)
	}

	if *vlanSpan.EnabledFlag == true {
		d.Set("vlan_spanning", "on")
	} else {
		d.Set("vlan_spanning", "off")
	}

	return nil
}

func resourceIBMNetworkVlanSpanCreate(d *schema.ResourceData, meta interface{}) error {
	sess := meta.(ClientSession).SoftLayerSession()
	service := services.GetAccountService(sess)

	rnd := rand.Intn(8999999) + 1000000
	vlanSpanning := d.Get("vlan_spanning").(string)

	var enabled bool
	if vlanSpanning == "on" {
		enabled = true
	} else {
		enabled = false
	}

	_, err := service.SetVlanSpan(&enabled)
	if err != nil {
		return fmt.Errorf("Error settinging VLAN Spanning %s", err)
	}

	d.SetId(strconv.Itoa(rnd))
	return resourceIBMNetworkVlanSpanRead(d, meta)
}

func resourceIBMNetworkVlanSpanUpdate(d *schema.ResourceData, meta interface{}) error {
	sess := meta.(ClientSession).SoftLayerSession()
	service := services.GetAccountService(sess)
	vlanSpanning := d.Get("vlan_spanning").(string)

	var enabled bool
	if vlanSpanning == "on" {
		enabled = true
	} else {
		enabled = false
	}

	_, err := service.SetVlanSpan(&enabled)
	if err != nil {
		return fmt.Errorf("Error settinging VLAN Spanning %s", err)
	}

	return resourceIBMNetworkVlanSpanRead(d, meta)
}

func resourceIBMNetworkVlanSpanDelete(d *schema.ResourceData, meta interface{}) error {
	// Leave VLAN Spanning setting in current state
	d.SetId("")

	return nil
}
