package ibm

import (
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	st "github.ibm.com/Bluemix/power-go-client/clients/instance"
	"github.ibm.com/Bluemix/power-go-client/helpers"
	"github.ibm.com/Bluemix/power-go-client/power/client/p_cloud_p_vm_instances"
	"github.ibm.com/Bluemix/power-go-client/power/models"

	"log"
	"time"
)

func resourceIBMPIInstance() *schema.Resource {
	return &schema.Resource{
		Create:   resourceIBMPIInstanceCreate,
		Read:     resourceIBMPIInstanceRead,
		Update:   resourceIBMPIInstanceUpdate,
		Delete:   resourceIBMPIInstanceDelete,
		Exists:   resourceIBMPIInstanceExists,
		Importer: &schema.ResourceImporter{},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(60 * time.Minute),
			Delete: schema.DefaultTimeout(60 * time.Minute),
		},

		Schema: map[string]*schema.Schema{

			helpers.PICloudInstanceId: {
				Type:     schema.TypeString,
				Required: true,
			},
			helpers.PIInstanceDiskSize: {
				Type:     schema.TypeInt,
				Computed: true,
			},
			helpers.PIInstanceStatus: {
				Type:     schema.TypeString,
				Computed: true,
			},
			helpers.PIInstanceMigratable: {
				Type:     schema.TypeBool,
				Required: true,
			},
			helpers.PIInstanceMinProc: {
				Type:     schema.TypeFloat,
				Computed: true,
			},
			helpers.PIInstanceNetworkIds: {
				Type:     schema.TypeSet,
				Required: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set:      schema.HashString,
			},

			helpers.PIInstanceVolumeIds: {
				Type:     schema.TypeSet,
				Required: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set:      schema.HashString,
			},

			"addresses": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ip": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"macaddress": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"networkid": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"networkname": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						/*"version": {
							Type:     schema.TypeFloat,
							Computed: true,
						},*/
					},
				},
			},

			helpers.PIInstanceHealthStatus: {
				Type:     schema.TypeString,
				Computed: true,
			},
			helpers.PIInstanceId: {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			helpers.PIInstanceDate: {
				Type:     schema.TypeString,
				Computed: true,
			},
			helpers.PIInstanceImageName: {
				Type:     schema.TypeString,
				Required: true,
			},
			helpers.PIInstanceProcessors: {
				Type:     schema.TypeFloat,
				Required: true,
			},
			helpers.PIInstanceName: {
				Type:     schema.TypeString,
				Required: true,
			},
			helpers.PIInstanceProcType: {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateAllowedStringValue([]string{"dedicated", "shared"}),
			},
			helpers.PIInstanceSSHKeyName: {
				Type:     schema.TypeString,
				Required: true,
			},
			helpers.PIInstanceMemory: {
				Type:     schema.TypeFloat,
				Required: true,
			},
			helpers.PIInstanceSystemType: {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateAllowedStringValue([]string{"any", "s922", "e880"}),
			},
			helpers.PIInstanceReplicants: {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			helpers.PIInstanceReplicationPolicy: {
				Type:     schema.TypeString,
				Optional: true,
			},
			helpers.PIInstanceProgress: {
				Type:        schema.TypeFloat,
				Computed:    true,
				Description: "Progress of the operation",
			},
		},
	}
}

func resourceIBMPIInstanceCreate(d *schema.ResourceData, meta interface{}) error {

	log.Printf("Now in the PowerVMCreate")
	sess, err := meta.(ClientSession).IBMPISession()
	if err != nil {
		return err
	}
	powerinstanceid := d.Get(helpers.PICloudInstanceId).(string)

	name := d.Get(helpers.PIInstanceName).(string)
	sshkey := d.Get(helpers.PIInstanceSSHKeyName).(string)
	mem := d.Get(helpers.PIInstanceMemory).(float64)
	procs := d.Get(helpers.PIInstanceProcessors).(float64)
	migrateable := d.Get(helpers.PIInstanceMigratable).(bool)
	systype := d.Get(helpers.PIInstanceSystemType).(string)
	networks := expandStringList((d.Get(helpers.PIInstanceNetworkIds).(*schema.Set)).List())
	volids := expandStringList((d.Get(helpers.PIInstanceVolumeIds).(*schema.Set)).List())
	replicants := d.Get(helpers.PIInstanceReplicants).(float64)
	if d.Get(helpers.PIInstanceReplicants) == "" {
		replicants = 1
	}
	replicationpolicy := d.Get(helpers.PIInstanceReplicationPolicy).(string)
	if d.Get(helpers.PIInstanceReplicationPolicy) == "" {
		replicationpolicy = "none"
	}

	imageid := d.Get(helpers.PIInstanceImageName).(string)

	processortype := d.Get(helpers.PIInstanceProcType).(string)

	body := &models.PVMInstanceCreate{

		VolumeIds: volids, NetworkIds: networks, Processors: &procs, Memory: &mem, ServerName: ptrToString(name),
		Migratable:              &migrateable,
		SysType:                 ptrToString(systype),
		KeyPairName:             sshkey,
		ImageID:                 ptrToString(imageid),
		ProcType:                ptrToString(processortype),
		Replicants:              replicants,
		ReplicantAffinityPolicy: ptrToString(replicationpolicy),
	}

	client := st.NewIBMPIInstanceClient(sess, powerinstanceid)

	pvm, _, _, err := client.Create(&p_cloud_p_vm_instances.PcloudPvminstancesPostParams{
		Body: body,
	}, powerinstanceid)
	log.Printf("the number of instances is %d", len(*pvm))

	if err != nil {
		log.Printf("[DEBUG]  err %s", isErrorToString(err))
		return err
	}

	log.Printf("Printing the instance info %+v", &pvm)

	truepvmid := (*pvm)[0].PvmInstanceID
	d.SetId(*truepvmid)
	//d.Set("addresses",(*pvm)[0].Addresses)

	log.Printf("Printing the instance id .. after the create ... %s", *truepvmid)

	_, err = isWaitForPIInstanceAvailable(client, *truepvmid, d.Timeout(schema.TimeoutCreate), powerinstanceid)
	if err != nil {
		return err
	}

	return resourceIBMPIInstanceRead(d, meta)
}

func resourceIBMPIInstanceRead(d *schema.ResourceData, meta interface{}) error {

	log.Printf("Calling the PowerInstance Read code..")

	sess, err := meta.(ClientSession).IBMPISession()
	if err != nil {
		return err
	}
	powerinstanceid := d.Get(helpers.PICloudInstanceId).(string)
	powerC := st.NewIBMPIInstanceClient(sess, powerinstanceid)
	powervmdata, err := powerC.Get(d.Id(), powerinstanceid)

	if err != nil {
		return err
	}

	pvminstanceid := *powervmdata.PvmInstanceID

	log.Printf("The Power pvm instance id is %s", pvminstanceid)

	d.SetId(pvminstanceid)
	d.Set("memory", powervmdata.Memory)
	d.Set("processors", powervmdata.Processors)
	d.Set(helpers.PIInstanceStatus, powervmdata.Status)
	d.Set("proctype", powervmdata.ProcType)
	d.Set("migratable", powervmdata.Migratable)
	d.Set(helpers.PIInstanceMinProc, powervmdata.Minproc)
	d.Set(helpers.PIInstanceProgress, powervmdata.Progress)

	if powervmdata.Addresses != nil {
		pvmaddress := make([]map[string]interface{}, len(powervmdata.Addresses))
		for i, pvmip := range powervmdata.Addresses {
			log.Printf("Now entering the powervm address space....")

			p := make(map[string]interface{})
			p["ip"] = pvmip.IP
			p["networkname"] = pvmip.NetworkName
			p["networkid"] = pvmip.NetworkID
			p["macaddress"] = pvmip.MacAddress
			p["type"] = pvmip.Type
			pvmaddress[i] = p
		}
		d.Set("addresses", pvmaddress)

		//log.Printf("Printing the value after the read - this should set it.... %+v", pvmaddress)

	}

	if powervmdata.Health != nil {
		d.Set(helpers.PIInstanceHealthStatus, powervmdata.Health.Status)

	}

	return nil

}

func resourceIBMPIInstanceUpdate(d *schema.ResourceData, meta interface{}) error {

	sess, _ := meta.(ClientSession).IBMPISession()

	name := d.Get(helpers.PIInstanceName).(string)
	mem := d.Get(helpers.PIInstanceMemory).(float64)
	procs := d.Get(helpers.PIInstanceProcessors).(float64)
	migrateable := d.Get(helpers.PIInstanceMigratable).(bool)
	processortype := d.Get(helpers.PIInstanceProcType).(string)
	powerinstanceid := d.Get(helpers.PICloudInstanceId).(string)

	client := st.NewIBMPIInstanceClient(sess, powerinstanceid)

	body := &models.PVMInstanceUpdate{
		Memory:     mem,
		Migratable: &migrateable,
		ProcType:   processortype,
		Processors: procs,
		ServerName: name,
	}

	resp, err := client.Update(d.Id(), powerinstanceid, &p_cloud_p_vm_instances.PcloudPvminstancesPutParams{Body: body})
	if err != nil {
		return err
	}

	log.Printf("Getting the response %s", resp.StatusURL)

	_, err = isWaitForPIInstanceAvailable(client, d.Id(), d.Timeout(schema.TimeoutCreate), powerinstanceid)
	if err != nil {
		return err
	}

	return resourceIBMPIInstanceRead(d, meta)

	return nil
}

func resourceIBMPIInstanceDelete(data *schema.ResourceData, meta interface{}) error {
	sess, _ := meta.(ClientSession).IBMPISession()
	powerinstanceid := data.Get(helpers.PICloudInstanceId).(string)
	client := st.NewIBMPIInstanceClient(sess, powerinstanceid)
	err := client.Delete(data.Id(), powerinstanceid)
	if err != nil {
		return err
	}

	_, err = isWaitForPIInstanceDeleted(client, data.Id(), data.Timeout(schema.TimeoutDelete), powerinstanceid)
	if err != nil {
		return err
	}

	data.SetId("")
	return nil
}

// Exists

func resourceIBMPIInstanceExists(d *schema.ResourceData, meta interface{}) (bool, error) {

	sess, err := meta.(ClientSession).IBMPISession()
	if err != nil {
		return false, err
	}
	id := d.Id()
	powerinstanceid := d.Get(helpers.PICloudInstanceId).(string)
	client := st.NewIBMPIInstanceClient(sess, powerinstanceid)

	instance, err := client.Get(d.Id(), powerinstanceid)
	if err != nil {

		return false, err
	}
	return instance.PvmInstanceID == &id, nil
}

func isWaitForPIInstanceDeleted(client *st.IBMPIInstanceClient, id string, timeout time.Duration, powerinstanceid string) (interface{}, error) {

	log.Printf("Waiting for  (%s) to be deleted.", id)

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"retry", helpers.PIInstanceDeleting},
		Target:     []string{helpers.PIInstanceNotFound},
		Refresh:    isPIInstanceDeleteRefreshFunc(client, id, powerinstanceid),
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForState()
}

func isPIInstanceDeleteRefreshFunc(client *st.IBMPIInstanceClient, id, powerinstanceid string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		pvm, err := client.Get(id, powerinstanceid)
		if err != nil {
			log.Printf("The power vm does not exist")
			return pvm, helpers.PIInstanceNotFound, nil

		}
		return pvm, helpers.PIInstanceNotFound, nil

	}
}

func isWaitForPIInstanceAvailable(client *st.IBMPIInstanceClient, id string, timeout time.Duration, powerinstanceid string) (interface{}, error) {
	log.Printf("Waiting for PIInstance (%s) to be available and sleeping ", id)

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"PENDING", helpers.PIInstanceHealthWarning},
		Target:     []string{"OK", helpers.PIInstanceHealthOk},
		Refresh:    isPIInstanceRefreshFunc(client, id, powerinstanceid),
		Delay:      3 * time.Minute,
		MinTimeout: 30 * time.Second,
		Timeout:    30 * time.Minute,
	}

	return stateConf.WaitForState()
}

func isPIInstanceRefreshFunc(client *st.IBMPIInstanceClient, id, powerinstanceid string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {

		pvm, err := client.Get(id, powerinstanceid)
		if err != nil {
			return nil, "", err
		}

		if pvm.Health.Status == helpers.PIInstanceHealthOk {
			log.Printf("The health status is now ok")
			//if *pvm.Status == "active" ; if *pvm.Addresses[0].IP == nil  {
			return pvm, helpers.PIInstanceHealthOk, nil
			//}
		}

		return pvm, helpers.PIInstanceHealthWarning, nil
	}
}

func checkPIActive(vminstance *models.PVMInstance) bool {

	log.Printf("Calling the check vm status function and the health status is %s", vminstance.Health.Status)
	activeStatus := false

	if vminstance.Health.Status == "OK" {
		//if *vminstance.Status == "active" {
		log.Printf(" The status of the vm is now set to what we want it to be %s", vminstance.Health.Status)
		activeStatus = true

	}
	return activeStatus
}
