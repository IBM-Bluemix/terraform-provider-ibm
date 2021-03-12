/**
 * (C) Copyright IBM Corp. 2021.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package ibm

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM/vpc-go-sdk/vpcv1"
)

func dataSourceIbmIsDedicatedHostGroup() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIbmIsDedicatedHostGroupRead,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The unique user-defined name for this dedicated host. If unspecified, the name will be a hyphenated list of randomly-selected words.",
			},
			"class": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The dedicated host profile class for hosts in this group.",
			},
			"created_at": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date and time that the dedicated host group was created.",
			},
			"crn": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The CRN for this dedicated host group.",
			},
			"dedicated_hosts": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The dedicated hosts that are in this dedicated host group.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"crn": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The CRN for this dedicated host.",
						},
						"deleted": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "If present, this property indicates the referenced resource has been deleted and providessome supplementary information.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"more_info": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Link to documentation about deleted resources.",
									},
								},
							},
						},
						"href": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The URL for this dedicated host.",
						},
						"id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique identifier for this dedicated host.",
						},
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique user-defined name for this dedicated host. If unspecified, the name will be a hyphenated list of randomly-selected words.",
						},
						"resource_type": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type of resource referenced.",
						},
					},
				},
			},
			"family": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The dedicated host profile family for hosts in this group.",
			},
			"href": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The URL for this dedicated host group.",
			},
			"resource_group": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The unique identifier of the resource group for this dedicated host group.",
			},
			"resource_type": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The type of resource referenced.",
			},
			"supported_instance_profiles": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Array of instance profiles that can be used by instances placed on this dedicated host group.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"href": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The URL for this virtual server instance profile.",
						},
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The globally unique name for this virtual server instance profile.",
						},
					},
				},
			},
			"zone": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The globally unique name of the zone this dedicated host group resides in.",
			},
		},
	}
}

func dataSourceIbmIsDedicatedHostGroupRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	vpcClient, err := meta.(ClientSession).VpcV1API()
	if err != nil {
		return diag.FromErr(err)
	}

	listDedicatedHostGroupsOptions := &vpcv1.ListDedicatedHostGroupsOptions{}

	dedicatedHostGroupCollection, response, err := vpcClient.ListDedicatedHostGroupsWithContext(context, listDedicatedHostGroupsOptions)
	if err != nil {
		log.Printf("[DEBUG] ListDedicatedHostGroupsWithContext failed %s\n%s", err, response)
		return diag.FromErr(err)
	}

	if len(dedicatedHostGroupCollection.Groups) == 0 {
		return nil
	}
	dedicatedHostGroup := vpcv1.DedicatedHostGroup{}
	name := d.Get("name").(string)
	for _, data := range dedicatedHostGroupCollection.Groups {
		if *data.Name == name {
			dedicatedHostGroup = data
			d.SetId(*dedicatedHostGroup.ID)
			if err = d.Set("class", dedicatedHostGroup.Class); err != nil {
				return diag.FromErr(fmt.Errorf("Error setting class: %s", err))
			}
			if dedicatedHostGroup.CreatedAt != nil {
				if err = d.Set("created_at", dedicatedHostGroup.CreatedAt.String()); err != nil {
					return diag.FromErr(fmt.Errorf("Error setting created_at: %s", err))
				}
			}

			if err = d.Set("crn", dedicatedHostGroup.CRN); err != nil {
				return diag.FromErr(fmt.Errorf("Error setting crn: %s", err))
			}

			if dedicatedHostGroup.DedicatedHosts != nil {
				err = d.Set("dedicated_hosts", dataSourceDedicatedHostGroupFlattenDedicatedHosts(dedicatedHostGroup.DedicatedHosts))
				if err != nil {
					return diag.FromErr(fmt.Errorf("Error setting dedicated_hosts %s", err))
				}
			}
			if err = d.Set("family", dedicatedHostGroup.Family); err != nil {
				return diag.FromErr(fmt.Errorf("Error setting family: %s", err))
			}
			if err = d.Set("href", dedicatedHostGroup.Href); err != nil {
				return diag.FromErr(fmt.Errorf("Error setting href: %s", err))
			}

			if dedicatedHostGroup.ResourceGroup != nil {
				err = d.Set("resource_group", *dedicatedHostGroup.ResourceGroup.ID)
				if err != nil {
					return diag.FromErr(fmt.Errorf("Error setting resource_group %s", err))
				}
			}
			if err = d.Set("resource_type", dedicatedHostGroup.ResourceType); err != nil {
				return diag.FromErr(fmt.Errorf("Error setting resource_type: %s", err))
			}

			if dedicatedHostGroup.SupportedInstanceProfiles != nil {
				err = d.Set("supported_instance_profiles", dataSourceDedicatedHostGroupFlattenSupportedInstanceProfiles(dedicatedHostGroup.SupportedInstanceProfiles))
				if err != nil {
					return diag.FromErr(fmt.Errorf("Error setting supported_instance_profiles %s", err))
				}
			}

			if dedicatedHostGroup.Zone != nil {
				err = d.Set("zone", *dedicatedHostGroup.Zone.Name)
				if err != nil {
					return diag.FromErr(fmt.Errorf("Error setting zone %s", err))
				}
			}
			return nil
		}
	}
	return diag.FromErr(fmt.Errorf("No Dedicated Host Group found with name %s", name))
}

// dataSourceIbmIsDedicatedHostGroupID returns a reasonable ID for the list.
func dataSourceIbmIsDedicatedHostGroupID(d *schema.ResourceData) string {
	return "abcd"
}

func dataSourceDedicatedHostGroupFlattenDedicatedHosts(result []vpcv1.DedicatedHostReference) (dedicatedHosts []map[string]interface{}) {
	for _, dedicatedHostsItem := range result {
		dedicatedHosts = append(dedicatedHosts, dataSourceDedicatedHostGroupDedicatedHostsToMap(dedicatedHostsItem))
	}

	return dedicatedHosts
}

func dataSourceDedicatedHostGroupDedicatedHostsToMap(dedicatedHostsItem vpcv1.DedicatedHostReference) (dedicatedHostsMap map[string]interface{}) {
	dedicatedHostsMap = map[string]interface{}{}

	if dedicatedHostsItem.CRN != nil {
		dedicatedHostsMap["crn"] = dedicatedHostsItem.CRN
	}
	if dedicatedHostsItem.Deleted != nil {
		deletedList := []map[string]interface{}{}
		deletedMap := dataSourceDedicatedHostGroupDedicatedHostsDeletedToMap(*dedicatedHostsItem.Deleted)
		deletedList = append(deletedList, deletedMap)
		dedicatedHostsMap["deleted"] = deletedList
	}
	if dedicatedHostsItem.Href != nil {
		dedicatedHostsMap["href"] = dedicatedHostsItem.Href
	}
	if dedicatedHostsItem.ID != nil {
		dedicatedHostsMap["id"] = dedicatedHostsItem.ID
	}
	if dedicatedHostsItem.Name != nil {
		dedicatedHostsMap["name"] = dedicatedHostsItem.Name
	}
	if dedicatedHostsItem.ResourceType != nil {
		dedicatedHostsMap["resource_type"] = dedicatedHostsItem.ResourceType
	}

	return dedicatedHostsMap
}

func dataSourceDedicatedHostGroupDedicatedHostsDeletedToMap(deletedItem vpcv1.DedicatedHostReferenceDeleted) (deletedMap map[string]interface{}) {
	deletedMap = map[string]interface{}{}

	if deletedItem.MoreInfo != nil {
		deletedMap["more_info"] = deletedItem.MoreInfo
	}

	return deletedMap
}

func dataSourceDedicatedHostGroupFlattenResourceGroup(result vpcv1.ResourceGroupReference) (finalList []map[string]interface{}) {
	finalList = []map[string]interface{}{}
	finalMap := dataSourceDedicatedHostGroupResourceGroupToMap(result)
	finalList = append(finalList, finalMap)

	return finalList
}

func dataSourceDedicatedHostGroupResourceGroupToMap(resourceGroupItem vpcv1.ResourceGroupReference) (resourceGroupMap map[string]interface{}) {
	resourceGroupMap = map[string]interface{}{}

	if resourceGroupItem.Href != nil {
		resourceGroupMap["href"] = resourceGroupItem.Href
	}
	if resourceGroupItem.ID != nil {
		resourceGroupMap["id"] = resourceGroupItem.ID
	}
	if resourceGroupItem.Name != nil {
		resourceGroupMap["name"] = resourceGroupItem.Name
	}

	return resourceGroupMap
}

func dataSourceDedicatedHostGroupFlattenSupportedInstanceProfiles(result []vpcv1.InstanceProfileReference) (supportedInstanceProfiles []map[string]interface{}) {
	for _, supportedInstanceProfilesItem := range result {
		supportedInstanceProfiles = append(supportedInstanceProfiles, dataSourceDedicatedHostGroupSupportedInstanceProfilesToMap(supportedInstanceProfilesItem))
	}

	return supportedInstanceProfiles
}

func dataSourceDedicatedHostGroupSupportedInstanceProfilesToMap(supportedInstanceProfilesItem vpcv1.InstanceProfileReference) (supportedInstanceProfilesMap map[string]interface{}) {
	supportedInstanceProfilesMap = map[string]interface{}{}

	if supportedInstanceProfilesItem.Href != nil {
		supportedInstanceProfilesMap["href"] = supportedInstanceProfilesItem.Href
	}
	if supportedInstanceProfilesItem.Name != nil {
		supportedInstanceProfilesMap["name"] = supportedInstanceProfilesItem.Name
	}

	return supportedInstanceProfilesMap
}

func dataSourceDedicatedHostGroupFlattenZone(result vpcv1.ZoneReference) (finalList []map[string]interface{}) {
	finalList = []map[string]interface{}{}
	finalMap := dataSourceDedicatedHostGroupZoneToMap(result)
	finalList = append(finalList, finalMap)

	return finalList
}

func dataSourceDedicatedHostGroupZoneToMap(zoneItem vpcv1.ZoneReference) (zoneMap map[string]interface{}) {
	zoneMap = map[string]interface{}{}

	if zoneItem.Href != nil {
		zoneMap["href"] = zoneItem.Href
	}
	if zoneItem.Name != nil {
		zoneMap["name"] = zoneItem.Name
	}

	return zoneMap
}
