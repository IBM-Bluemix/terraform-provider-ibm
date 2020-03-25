package ibm

import (
	"fmt"
	"log"
	"strings"
	"time"

	//"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	iserrors "github.ibm.com/Bluemix/riaas-go-client/errors"
)

const (
	pdnsInstanceID = "instance_id"
	pdnsZoneName   = "zone_name"
)

func resourceIBMPrivateDNSZone() *schema.Resource {
	return &schema.Resource{
		Create:   resourceIBMPrivateDnsZoneCreate,
		Read:     resourceIBMPrivateDnsZoneRead,
		Update:   resourceIBMPrivateDnsZoneUpdate,
		Delete:   resourceIBMPrivateDnsZoneDelete,
		Exists:   resourceIBMPrivateDnsZoneExists,
		Importer: &schema.ResourceImporter{},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},

		CustomizeDiff: customdiff.Sequence(
			func(diff *schema.ResourceDiff, v interface{}) error {
				return resourceTagsCustomizeDiff(diff)
			},
		),

		Schema: map[string]*schema.Schema{
			pdnsInstanceID: {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: false,
			},

			pdnsZoneName: {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: false,
			},
		},
	}
}

func resourceIBMPrivateDnsZoneCreate(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(ClientSession).PrivateDnsClientSession()
	if err != nil {
		return err
	}

	instanceID := d.Get(pdnsInstanceID).(string)
	zoneName := d.Get(pdnsZoneName).(string)
	createZoneOptions := sess.NewCreateDnszoneOptions(instanceID, zoneName)
	createZoneOptions.SetDescription("zone description")
	createZoneOptions.SetLabel("zone_label")
	response, _, err := sess.CreateDnszone(createZoneOptions)
	if err != nil {
		return err
	}

	d.SetId(fmt.Sprintf("%s/%s", *response.InstanceID, *response.ID))
	if err != nil {
		log.Printf("[DEBUG]  err %s", isErrorToString(err))
		return err
	}

	log.Printf("[DEBUG] TEST5")

	return resourceIBMPrivateDnsZoneRead(d, meta)
}

func resourceIBMPrivateDnsZoneRead(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(ClientSession).PrivateDnsClientSession()
	if err != nil {
		return err
	}

	id_set := strings.Split(d.Id(), "/")
	getZoneOptions := sess.NewGetDnszoneOptions(id_set[0], id_set[1])
	response, _, reqErr := sess.GetDnszone(getZoneOptions)
	if reqErr == nil {
		return err
	}

	d.Set("id", response.ID)
	d.Set("instance_id", response.InstanceID)

	return nil
}

func resourceIBMPrivateDnsZoneUpdate(d *schema.ResourceData, meta interface{}) error {
	return resourceIBMPrivateDnsZoneRead(d, meta)
}

func resourceIBMPrivateDnsZoneDelete(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(ClientSession).PrivateDnsClientSession()
	if err != nil {
		return err
	}

	id_set := strings.Split(d.Id(), "/")

	deleteZoneOptions := sess.NewDeleteDnszoneOptions(id_set[0], id_set[1])
	_, reqErr := sess.DeleteDnszone(deleteZoneOptions)
	if reqErr == nil {
		return reqErr
	}

	d.SetId("")
	return nil
}

func resourceIBMPrivateDnsZoneExists(d *schema.ResourceData, meta interface{}) (bool, error) {

	sess, err := meta.(ClientSession).PrivateDnsClientSession()
	if err != nil {
		return false, err
	}

	id_set := strings.Split(d.Id(), "/")
	getZoneOptions := sess.NewGetDnszoneOptions(id_set[0], id_set[1])
	_, _, err = sess.GetDnszone(getZoneOptions)
	if err != nil {
		iserror, ok := err.(iserrors.RiaasError)
		if ok {
			if len(iserror.Payload.Errors) == 1 &&
				iserror.Payload.Errors[0].Code == "not_found" {
				return false, nil
			}
		}
		return false, err
	}
	return true, nil
}
