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

func dataSourceIBMAppDomainPrivate() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceIBMAppDomainPrivateRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Description: "The name of the private domain",
				Type:        schema.TypeString,
				Required:    true,
			},
		},
	}
}

func dataSourceIBMAppDomainPrivateRead(d *schema.ResourceData, meta interface{}) error {
	cfAPI, err := meta.(ClientSession).MccpAPI()
	if err != nil {
		return err
	}
	domainName := d.Get("name").(string)
	prdomain, err := cfAPI.PrivateDomains().FindByName(domainName)
	if err != nil {
		return fmt.Errorf("Error retrieving domain: %s", err)
	}
	d.SetId(prdomain.GUID)
	return nil

}
