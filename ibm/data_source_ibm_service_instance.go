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

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceIBMServiceInstance() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceIBMServiceInstanceRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Description: "Service instance name for example, speech_to_text",
				Type:        schema.TypeString,
				Required:    true,
			},

			"space_guid": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The guid of the space in which the instance is present",
			},

			"credentials": {
				Description: "The service broker-provided credentials to use this service.",
				Type:        schema.TypeMap,
				Sensitive:   true,
				Computed:    true,
			},

			"service_keys": {
				Description: "Service keys asociated with the service instance",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The service key name",
						},
						"credentials": {
							Type:        schema.TypeMap,
							Computed:    true,
							Sensitive:   true,
							Description: "The service key credential details like port, username etc",
						},
					},
				},
			},

			"service_plan_guid": {
				Description: "The uniquie identifier of the service offering plan type",
				Type:        schema.TypeString,
				Computed:    true,
			},
		},
	}
}

func dataSourceIBMServiceInstanceRead(d *schema.ResourceData, meta interface{}) error {
	cfClient, err := meta.(ClientSession).MccpAPI()
	if err != nil {
		return err
	}
	siAPI := cfClient.ServiceInstances()
	name := d.Get("name").(string)
	spaceGUID := d.Get("space_guid").(string)
	inst, err := siAPI.FindByNameInSpace(spaceGUID, name)
	if err != nil {
		return err
	}

	serviceInstance, err := siAPI.Get(inst.GUID, 1)
	if err != nil {
		return fmt.Errorf("Error retrieving service: %s", err)
	}

	d.SetId(serviceInstance.Metadata.GUID)
	serviceKeys := serviceInstance.Entity.ServiceKeys
	d.Set("credentials", Flatten(serviceInstance.Entity.Credentials))
	d.Set("service_keys", flattenServiceInstanceCredentials(serviceKeys))
	d.Set("service_plan_guid", serviceInstance.Entity.ServicePlanGUID)

	return nil
}
