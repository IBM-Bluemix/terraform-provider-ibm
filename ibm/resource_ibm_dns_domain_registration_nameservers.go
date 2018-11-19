package ibm

import (
	//"errors"
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform/helper/schema"
	//"github.com/softlayer/softlayer-go/datatypes"
	"github.com/softlayer/softlayer-go/services"
	//"github.com/softlayer/softlayer-go/sl"
)

func resourceIBMDNSDomainRegistrationNameservers() *schema.Resource {
	return &schema.Resource{
		Create: resourceIBMDNSDomainRegistrationNSCreate,
		Read:   resourceIBMDNSDomainRegistrationNSRead,
		Update: resourceIBMDNSDomainRegistrationNSUpdate,
		Delete: resourceIBMDNSDomainRegistrationNSDelete,
		Schema: map[string]*schema.Schema{
			"dns_registration_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name_servers": &schema.Schema{
				Description: "Custom name servers for the domain registration",
				Type:        schema.TypeSet,
				Required:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func resourceIBMDNSDomainRegistrationNSCreate(d *schema.ResourceData, meta interface{}) error {
	sess := meta.(ClientSession).SoftLayerSession()
	nService := services.GetDnsDomainRegistrationService(sess)
	dnsId, _ := strconv.Atoi(d.Get("dns_registration_id").(string))
	newNameServers := d.Get("name_servers").(*schema.Set).List()

	// Get existing name servers as these will be replaced later
	dns_domain_nameservers, err := nService.Id(dnsId).
		Mask("nameservers.name").
		GetDomainNameservers()

	if err != nil {
		return fmt.Errorf("Error retrieving domain Registration NSCreate: %s", err)
	}

	if len(dns_domain_nameservers) == 0 {
		return fmt.Errorf("No domain found with id NSCreate [%d]", dnsId)
	}
	oldNameServers := make([]string, len(dns_domain_nameservers[0].Nameservers))
	for i, elem := range dns_domain_nameservers[0].Nameservers {
		oldNameServers[i] = *elem.Name
	}
	//
	log.Printf("Original DNS registration name servers %s\n", oldNameServers)

	// New NS to add, if not found in old list
	var addNs []string
	for _, newNs := range newNameServers {
		found := false
		for _, oldNs := range oldNameServers {
			log.Printf("old %s, new %s", oldNs, newNs)

			if oldNs == newNs {

				found = true
				break
			}
		}
		if found == false {
			addNs = append(addNs, newNs.(string))
		}
	}
	log.Printf("Name servers to add %v\n", addNs)

	// if no name servers to add then, we already have the correct name servers.
	// So return at this point.
	if len(addNs) == 0 {
		d.SetId(fmt.Sprintf("%d", dnsId))
		return resourceIBMDNSDomainRegistrationNSRead(d, meta)
	}

	nsUnlock_res, err := nService.Id(dnsId).
		UnlockDomain()
	if err != nil || nsUnlock_res != true {
		return fmt.Errorf("Error unlocking domain registration record: %s", err)
	}

	nsAdd_res := false
	nsAdd_res, err = nService.Id(dnsId).
		AddNameserversToDomain(addNs)

	if err != nil || nsAdd_res != true {
		return fmt.Errorf("Error Adding name servers to record: %s", err)
	}

	// old NS to delete, if not found in new list
	var delNs []string
	for _, oldNs := range oldNameServers {
		found := false
		for _, newNs := range newNameServers {
			log.Printf("old %s, new %s", oldNs, newNs)
			if oldNs == newNs {
				found = true
				break
			}
		}
		if found == false {
			delNs = append(delNs, oldNs)
		}
	}

	log.Printf("Name servers to delete %v\n", delNs)

	nsDel_res := false
	nsDel_res, err = nService.Id(dnsId).
		RemoveNameserversFromDomain(delNs)

	if err != nil || nsDel_res != true {
		return fmt.Errorf("Error Deleting name servers from record: %s", err)
	}

	_, _ = nService.Id(dnsId).LockDomain()
	// Ignore lock errors as does not impact operation

	d.SetId(fmt.Sprintf("%d", dnsId))
	return resourceIBMDNSDomainRegistrationNSRead(d, meta)
}

func resourceIBMDNSDomainRegistrationNSRead(d *schema.ResourceData, meta interface{}) error {
	sess := meta.(ClientSession).SoftLayerSession()
	dnsId, _ := strconv.Atoi(d.Id())
	//service := services.GetDnsDomainService(sess)

	nService := services.GetDnsDomainRegistrationService(sess)
	dns_domain_nameservers, err := nService.Id(dnsId).
		Mask("nameservers.name").
		GetDomainNameservers()

	if err != nil {
		return fmt.Errorf("Error retrieving domain registration NSReaD: %s", err)
	}

	if len(dns_domain_nameservers) == 0 {
		return fmt.Errorf("No domain found with id [%d]", dnsId)
	}

	log.Printf("list %v\n", dns_domain_nameservers)
	ns := make([]string, len(dns_domain_nameservers[0].Nameservers))
	for i, elem := range dns_domain_nameservers[0].Nameservers {
		ns[i] = *elem.Name
	}

	log.Printf("names %v\n", ns)

	if err != nil {
		return fmt.Errorf("Error retrieving domain registration nameservers: %s", err)
	}

	d.SetId(fmt.Sprintf("%d", dnsId))
	d.Set("name_servers", ns)
	return nil
}

// No delete on IBM Cloud
func resourceIBMDNSDomainRegistrationNSUpdate(d *schema.ResourceData, meta interface{}) error {
	return nil
}

// No delete on IBM Cloud
func resourceIBMDNSDomainRegistrationNSDelete(d *schema.ResourceData, meta interface{}) error {
	d.SetId("")
	return nil
}
