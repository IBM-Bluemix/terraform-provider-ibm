package ibm

import (
	"errors"
	"github.com/hashicorp/terraform/helper/schema"
	st "github.ibm.com/Bluemix/power-go-client/clients/instance"
	"github.ibm.com/Bluemix/power-go-client/helpers"
	"log"
	"time"
)

func resourceIBMPIVolumeAttach() *schema.Resource {
	return &schema.Resource{
		Create: resourceIBMPIVolumeAttachCreate,
		Read:   resourceIBMPIVolumeAttachRead,
		Update: resourceIBMPIVolumeAttachUpdate,
		Delete: resourceIBMPIVolumeAttachDelete,
		//Exists:   resourceIBMPowerVolumeExists,
		Importer: &schema.ResourceImporter{},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(60 * time.Minute),
			Delete: schema.DefaultTimeout(60 * time.Minute),
		},

		Schema: map[string]*schema.Schema{

			"volumeattachid": {
				Type:     schema.TypeString,
				Computed: true,
				ForceNew: true,
				Optional: true,
			},

			helpers.PICloudInstanceId: {
				Type:        schema.TypeString,
				Required:    true,
				Description: " Cloud Instance ID - This is the service_instance_id.",
			},

			helpers.PIVolumeAttachName: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the volume to attach. Note these  volumes should have been created",
			},

			helpers.PIInstanceName: {
				Type:     schema.TypeString,
				Required: true,
			},

			helpers.PIVolumeAttachStatus: {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			helpers.PIVolumeShareable: {
				Type:     schema.TypeBool,
				Computed: true,
				Optional: true,
			},
		},
	}
}

func resourceIBMPIVolumeAttachCreate(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(ClientSession).IBMPISession()
	if err != nil {
		return err
	}

	name := d.Get(helpers.PIVolumeAttachName).(string)
	servername := d.Get(helpers.PIInstanceName).(string)
	powerinstanceid := d.Get(helpers.PICloudInstanceId).(string)

	client := st.NewIBMPIVolumeClient(sess, powerinstanceid)

	//log.Print("Now doing a get with the volumename %s  ", name)
	volinfo, err := client.Get(name, powerinstanceid)

	if err != nil {
		return errors.New("The volume cannot be attached since it's not available")
		log.Printf(" The volume that is being attached is not available ")
	}
	//log.Print("The volume info is %s", volinfo)

	if volinfo.State == helpers.PIVolumeAllowableAttachStatus {

		return errors.New("The volume cannot be attached in the current state. The volume must be in the *available* state. No other states are permissible")
	}

	resp, err := client.Attach(servername, name, powerinstanceid)

	if err != nil {
		return err
	}
	log.Printf("Printing the resp %+v", resp)

	d.SetId(*volinfo.VolumeID)
	if err != nil {
		log.Printf("[DEBUG]  err %s", isErrorToString(err))
		return err
	}

	_, err = isWaitForIBMPIVolumeAvailable(client, d.Id(), powerinstanceid, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return err
	}
	//return nil
	return resourceIBMPIVolumeAttachRead(d, meta)
}

func resourceIBMPIVolumeAttachRead(d *schema.ResourceData, meta interface{}) error {
	sess, _ := meta.(ClientSession).IBMPISession()
	powerinstanceid := d.Get(helpers.PICloudInstanceId).(string)

	client := st.NewIBMPIVolumeClient(sess, powerinstanceid)
	vol, err := client.Get(d.Id(), powerinstanceid)
	if err != nil {
		return err
	}

	//d.SetId(vol.ID.String())
	d.Set(helpers.PIVolumeAttachName, vol.Name)
	d.Set(helpers.PIVolumeSize, vol.Size)
	d.Set(helpers.PIVolumeShareable, vol.Shareable)
	return nil
}

func resourceIBMPIVolumeAttachUpdate(d *schema.ResourceData, meta interface{}) error {

	sess, _ := meta.(ClientSession).IBMPISession()
	powerinstanceid := d.Get(helpers.PICloudInstanceId).(string)
	client := st.NewIBMPIVolumeClient(sess, powerinstanceid)

	name := ""
	if d.HasChange(helpers.PIVolumeAttachName) {
		name = d.Get(helpers.PIVolumeAttachName).(string)
	}

	size := float64(d.Get(helpers.PIVolumeSize).(float64))
	shareable := bool(d.Get(helpers.PIVolumeShareable).(bool))

	volrequest, err := client.Update(d.Id(), name, size, shareable, powerinstanceid)
	if err != nil {
		return err
	}

	_, err = isWaitForIBMPIVolumeAvailable(client, *volrequest.VolumeID, powerinstanceid, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return err
	}

	return resourceIBMPIVolumeRead(d, meta)
}

func resourceIBMPIVolumeAttachDelete(d *schema.ResourceData, meta interface{}) error {

	sess, _ := meta.(ClientSession).IBMPISession()
	powerinstanceid := d.Get(helpers.PICloudInstanceId).(string)
	client := st.NewIBMPIVolumeClient(sess, powerinstanceid)

	err := client.Delete(d.Id(), powerinstanceid)
	if err != nil {
		return err
	}

	// wait for power volume states to be back as available. if it's attached it will be in-use
	d.SetId("")
	return nil
}
