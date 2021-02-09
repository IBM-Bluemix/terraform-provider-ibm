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

	"github.com/IBM-Cloud/bluemix-go/api/iampap/iampapv1"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// Data source to find all the policies for a user in a particular account
func dataSourceIBMIAMUserPolicy() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceIBMIAMUserPolicyRead,

		Schema: map[string]*schema.Schema{
			"ibm_id": {
				Description: "The ibm id or email of user",
				Type:        schema.TypeString,
				Required:    true,
			},
			"policies": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"roles": {
							Type:        schema.TypeList,
							Computed:    true,
							Elem:        &schema.Schema{Type: schema.TypeString},
							Description: "Role names of the policy definition",
						},
						"resources": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"service": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Service name of the policy definition",
									},
									"resource_instance_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Elem:        &schema.Schema{Type: schema.TypeString},
										Description: "ID of resource instance of the policy definition",
									},
									"region": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Region of the policy definition",
									},
									"resource_type": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Resource type of the policy definition",
									},
									"resource": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Resource of the policy definition",
									},
									"resource_group_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "ID of the resource group.",
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func dataSourceIBMIAMUserPolicyRead(d *schema.ResourceData, meta interface{}) error {
	iampapClient, err := meta.(ClientSession).IAMPAPAPI()
	if err != nil {
		return err
	}

	userEmail := d.Get("ibm_id").(string)

	userDetails, err := meta.(ClientSession).BluemixUserDetails()
	if err != nil {
		return err
	}

	accountID := userDetails.userAccount

	ibmUniqueID, err := getIBMUniqueId(accountID, userEmail, meta)
	if err != nil {
		return err
	}

	policies, err := iampapClient.V1Policy().List(iampapv1.SearchParams{
		AccountID: accountID,
		IAMID:     ibmUniqueID,
		Type:      iampapv1.AccessPolicyType,
	})
	if err != nil {
		return err
	}

	userPolicies := make([]map[string]interface{}, 0, len(policies))
	for _, policy := range policies {
		roles := make([]string, len(policy.Roles))
		for i, role := range policy.Roles {
			roles[i] = role.Name
		}
		resources := flattenPolicyResource(policy.Resources)
		p := map[string]interface{}{
			"id":        fmt.Sprintf("%s/%s", userEmail, policy.ID),
			"roles":     roles,
			"resources": resources,
		}
		userPolicies = append(userPolicies, p)
	}
	d.SetId(userEmail)
	d.Set("policies", userPolicies)

	return nil
}
