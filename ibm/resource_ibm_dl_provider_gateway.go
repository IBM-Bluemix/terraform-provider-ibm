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
	"github.com/hashicorp/terraform-plugin-sdk/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
	"os"
)

const (
	customerAccountID = "customer_account_id"
)

func resourceIBMDLProviderGateway() *schema.Resource {
	return &schema.Resource{
		Create:   resourceIBMdlProviderGatewayCreate,
		Read:     resourceIBMdlProviderGatewayRead,
		Delete:   resourceIBMdlProviderGatewayDelete,
		Exists:   resourceIBMdlProviderGatewayExists,
		Update:   resourceIBMdlProviderGatewayUpdate,
		Importer: &schema.ResourceImporter{},
		CustomizeDiff: customdiff.Sequence(
			func(diff *schema.ResourceDiff, v interface{}) error {
				return resourceTagsCustomizeDiff(diff)
			},
		),

		Schema: map[string]*schema.Schema{
			dlBgpAsn: {
				Type:        schema.TypeInt,
				Required:    true,
				ForceNew:    true,
				Description: "BGP ASN",
			},
			dlBgpCerCidr: {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
				Description: "BGP customer edge router CIDR",
			},
			dlBgpIbmCidr: {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
				Description: "BGP IBM CIDR",
			},
			dlPort: {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Gateway port",
			},
			dlCreatedAt: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date and time resource was created",
			},
			dlName: {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     false,
				Description:  "The unique user-defined name for this gateway",
				ValidateFunc: InvokeValidator("ibm_dl_provider_gateway", dlName),
				// ValidateFunc: validateRegexpLen(1, 63, "^([a-zA-Z]|[a-zA-Z][-_a-zA-Z0-9]*[a-zA-Z0-9])$"),
			},

			dlSpeedMbps: {
				Type:        schema.TypeInt,
				Required:    true,
				ForceNew:    false,
				Description: "Gateway speed in megabits per second",
			},
			customerAccountID: {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				Description:  "Customer IBM Cloud account ID for the new gateway. A gateway object containing the pending create request will become available in the specified account.",
				ValidateFunc: InvokeValidator("ibm_dl_provider_gateway", customerAccountID),
			},
			dlOperationalStatus: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Gateway operational status",
			},
			dlType: {
				Type:        schema.TypeString,
				Description: "Gateway type",
				Computed:    true,
			},
			dlProviderAPIManaged: {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Indicates whether gateway was created through a provider portal",
			},
			dlVlan: {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "VLAN allocated for this gateway",
			},
			dlCrn: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The CRN (Cloud Resource Name) of this gateway",
			},
			dlTags: {
				Type:        schema.TypeSet,
				Optional:    true,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Set:         resourceIBMVPCHash,
				Description: "Tags for the direct link gateway",
			},
			ResourceControllerURL: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The URL of the IBM Cloud dashboard that can be used to explore and view details about this instance",
			},

			ResourceName: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The name of the resource",
			},

			ResourceCRN: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The crn of the resource",
			},

			ResourceStatus: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The status of the resource",
			},

			ResourceGroupName: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The resource group name in which resource is provisioned",
			},
		},
	}
}

func resourceIBMDLProviderGatewayValidator() *ResourceValidator {

	validateSchema := make([]ValidateSchema, 2)

	validateSchema = append(validateSchema,
		ValidateSchema{
			Identifier:                 customerAccountID,
			ValidateFunctionIdentifier: ValidateRegexpLen,
			Type:                       TypeString,
			Required:                   true,
			Regexp:                     `^[0-9a-f]+$`,
			MinValueLength:             1,
			MaxValueLength:             32})
	validateSchema = append(validateSchema,
		ValidateSchema{
			Identifier:                 dlName,
			ValidateFunctionIdentifier: ValidateRegexpLen,
			Type:                       TypeString,
			Required:                   true,
			Regexp:                     `^([a-zA-Z]|[a-zA-Z][-_a-zA-Z0-9]*[a-zA-Z0-9])$`,
			MinValueLength:             1,
			MaxValueLength:             63})

	ibmISDLGatewayResourceValidator := ResourceValidator{ResourceName: "ibm_dl_provider_gateway", Schema: validateSchema}
	return &ibmISDLGatewayResourceValidator
}

func resourceIBMdlProviderGatewayCreate(d *schema.ResourceData, meta interface{}) error {
	directLink, err := directlinkProviderClient(meta)
	if err != nil {
		return err
	}
	name := d.Get(dlName).(string)
	speed := int64(d.Get(dlSpeedMbps).(int))
	custAccountID := d.Get(customerAccountID).(string)
	bgpAsn := int64(d.Get(dlBgpAsn).(int))
	var portID string
	portID = d.Get(dlPort).(string)
	portIdentity, _ := directLink.NewProviderGatewayPortIdentity(portID)
	gatewayOptions := directLink.NewCreateProviderGatewayOptions(bgpAsn, custAccountID, name, portIdentity, speed)
	if _, ok := d.GetOk(dlBgpIbmCidr); ok {
		bgpIbmCidr := d.Get(dlBgpIbmCidr).(string)
		gatewayOptions.BgpIbmCidr = &bgpIbmCidr

	}
	if _, ok := d.GetOk(dlBgpCerCidr); ok {
		bgpCerCidr := d.Get(dlBgpCerCidr).(string)
		gatewayOptions.BgpCerCidr = &bgpCerCidr

	}

	gateway, response, err := directLink.CreateProviderGateway(gatewayOptions)
	if err != nil {
		log.Printf("[DEBUG] Create Direct Link Provider Gateway err %s\n%s", err, response)
		return err
	}
	d.SetId(*gateway.ID)

	log.Printf("[INFO] Created Direct Link Provider Gateway : %s", *gateway.ID)

	v := os.Getenv("IC_ENV_TAGS")
	if _, ok := d.GetOk(dlTags); ok || v != "" {
		oldList, newList := d.GetChange(dlTags)
		err = UpdateTagsUsingCRN(oldList, newList, meta, *gateway.Crn)
		if err != nil {
			log.Printf(
				"Error on create of resource direct link Provider gateway (%s) tags: %s", d.Id(), err)
		}
	}

	return resourceIBMdlProviderGatewayRead(d, meta)
}

func resourceIBMdlProviderGatewayRead(d *schema.ResourceData, meta interface{}) error {
	dtype := d.Get(dlType).(string)
	log.Printf("[INFO] Inside resourceIBMdlGatewayRead: %s", dtype)

	directLink, err := directlinkProviderClient(meta)
	if err != nil {
		return err
	}

	ID := d.Id()

	getOptions := directLink.NewGetProviderGatewayOptions(ID)

	log.Printf("[INFO] Calling getgateway api: %s", dtype)

	instance, response, err := directLink.GetProviderGateway(getOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error Getting Direct Link Gateway (%s Template): %s\n%s", dtype, err, response)
	}
	if instance.ID != nil {
		d.Set("id", *instance.ID)
	}
	if instance.Name != nil {
		d.Set(dlName, *instance.Name)
	}
	if instance.Crn != nil {
		d.Set(dlCrn, *instance.Crn)
	}
	if instance.BgpAsn != nil {
		d.Set(dlBgpAsn, *instance.BgpAsn)
	}
	if instance.BgpIbmCidr != nil {
		d.Set(dlBgpIbmCidr, *instance.BgpIbmCidr)
	}
	if instance.BgpIbmAsn != nil {
		d.Set(dlBgpIbmAsn, *instance.BgpIbmAsn)
	}

	if instance.BgpCerCidr != nil {
		d.Set(dlBgpCerCidr, *instance.BgpCerCidr)
	}
	if instance.ProviderApiManaged != nil {
		d.Set(dlProviderAPIManaged, *instance.ProviderApiManaged)
	}
	if instance.Type != nil {
		d.Set(dlType, *instance.Type)
	}
	if instance.SpeedMbps != nil {
		d.Set(dlSpeedMbps, *instance.SpeedMbps)
	}
	if instance.OperationalStatus != nil {
		d.Set(dlOperationalStatus, *instance.OperationalStatus)
	}
	if instance.BgpStatus != nil {
		d.Set(dlBgpStatus, *instance.BgpStatus)
	}

	if instance.Vlan != nil {
		d.Set(dlVlan, *instance.Vlan)
	}
	if instance.CustomerAccountID != nil {
		d.Set(customerAccountID, *instance.CustomerAccountID)
	}
	if instance.Port != nil {
		d.Set(dlPort, *instance.Port.ID)
	}

	if instance.CreatedAt != nil {
		d.Set(dlCreatedAt, instance.CreatedAt.String())
	}
	tags, err := GetTagsUsingCRN(meta, *instance.Crn)
	if err != nil {
		log.Printf(
			"Error on get of resource direct link gateway (%s) tags: %s", d.Id(), err)
	}
	d.Set(dlTags, tags)
	controller, err := getBaseController(meta)
	if err != nil {
		return err
	}
	d.Set(ResourceControllerURL, controller+"/interconnectivity/direct-link")
	d.Set(ResourceName, *instance.Name)
	d.Set(ResourceCRN, *instance.Crn)
	d.Set(ResourceStatus, *instance.OperationalStatus)
	return nil
}

func resourceIBMdlProviderGatewayUpdate(d *schema.ResourceData, meta interface{}) error {

	directLink, err := directlinkProviderClient(meta)
	if err != nil {
		return err
	}

	ID := d.Id()
	getOptions := directLink.NewGetProviderGatewayOptions(ID)

	log.Printf("[INFO] Calling getgateway provider api")

	instance, response, err := directLink.GetProviderGateway(getOptions)

	updateGatewayOptionsModel := directLink.NewUpdateProviderGatewayOptions(ID)

	if d.HasChange(dlTags) {
		oldList, newList := d.GetChange(dlTags)
		err = UpdateTagsUsingCRN(oldList, newList, meta, *instance.Crn)
		if err != nil {
			log.Printf(
				"Error on update of resource direct link gateway dedicated (%s) tags: %s", *instance.ID, err)
		}
	}

	if d.HasChange(dlName) {
		name := d.Get(dlName).(string)
		updateGatewayOptionsModel.Name = &name
	}
	if d.HasChange(dlSpeedMbps) {
		speed := int64(d.Get(dlSpeedMbps).(int))
		updateGatewayOptionsModel.SpeedMbps = &speed
	}

	_, response, err = directLink.UpdateProviderGateway(updateGatewayOptionsModel)
	if err != nil {
		log.Printf("[DEBUG] Update Direct Link Provider Gateway  err %s\n%s", err, response)
		return err
	}

	return resourceIBMdlProviderGatewayRead(d, meta)
}

func resourceIBMdlProviderGatewayDelete(d *schema.ResourceData, meta interface{}) error {

	directLink, err := directlinkProviderClient(meta)
	if err != nil {
		return err
	}

	ID := d.Id()
	delOptions := directLink.NewDeleteProviderGatewayOptions(ID)
	_, response, err := directLink.DeleteProviderGateway(delOptions)
	if err != nil {
		if response != nil && response.StatusCode != 404 {
			return nil
		}
		log.Printf("Error deleting Direct Link  Provider Gateway: %s %s ", response, err)
	}

	d.SetId("")
	return nil
}

func resourceIBMdlProviderGatewayExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	directLink, err := directlinkProviderClient(meta)
	if err != nil {
		return false, err
	}

	ID := d.Id()

	getOptions := directLink.NewGetProviderGatewayOptions(ID)
	_, response, err := directLink.GetProviderGateway(getOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return false, nil
		}
		return false, fmt.Errorf("Error Getting Direct Link Provider Gateway : %s\n%s", err, response)
	}
	return true, nil
}
