// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"reflect"
	"time"

	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceBareMetalServerNetworkInterfaces() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceIBMISBareMetalServerNetworkInterfacesRead,

		Schema: map[string]*schema.Schema{
			isBareMetalServerID: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The bare metal server identifier",
			},

			//network interface properties
			isBareMetalServerNetworkInterfaces: {
				Type:        schema.TypeList,
				Description: "A list of all network interfaces on a bare metal server. A network interface is an abstract representation of a network interface card and connects a bare metal server to a subnet. While each network interface can attach to only one subnet, multiple network interfaces can be created to attach to multiple subnets. Multiple interfaces may also attach to the same subnet.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						isBareMetalServerNicAllowIPSpoofing: {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Indicates whether source IP spoofing is allowed on this interface. If false, source IP spoofing is prevented on this interface. If true, source IP spoofing is allowed on this interface.",
						},
						isBareMetalServerNicEnableInfraNAT: {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "If true, the VPC infrastructure performs any needed NAT operations. If false, the packet is passed unmodified to/from the network interface, allowing the workload to perform any needed NAT operations.",
						},
						isBareMetalServerNicFloatingIPs: {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The floating IPs associated with this network interface.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									isBareMetalServerNicIpAddress: {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The globally unique IP address",
									},

									isBareMetalServerNicIpCRN: {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The CRN for this floating IP",
									},
									isBareMetalServerNicIpHref: {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The URL for this floating IP",
									},
									isBareMetalServerNicIpID: {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The unique identifier for this floating IP",
									},
									isBareMetalServerNicIpName: {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The unique user-defined name for this floating IP",
									},
								},
							},
						},
						isBareMetalServerNicHref: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The URL for this network interface",
						},
						isBareMetalServerNicInterfaceType: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The network interface type: [ pci, vlan ]",
						},
						isBareMetalServerNicReservedIps: {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The reserved IPs bound to this network interface.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									isBareMetalServerNicIpAddress: {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The globally unique IP address",
									},
									isBareMetalServerNicIpHref: {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The URL for this reserved IP",
									},
									isBareMetalServerNicIpID: {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The unique identifier for this reserved IP",
									},
									isBareMetalServerNicIpName: {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The unique user-defined name for this reserved IP",
									},
									isBareMetalServerNicResourceType: {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The resource type: [ subnet_reserved_ip ]",
									},
								},
							},
						},
						isBareMetalServerNicMacAddress: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The MAC address of the interface. If absent, the value is not known.",
						},
						isBareMetalServerNicName: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The user-defined name for this network interface",
						},
						isBareMetalServerNicPortSpeed: {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The network interface port speed in Mbps",
						},
						isBareMetalServerNicPrimaryIP: {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "title: IPv4, The IP address. ",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									isBareMetalServerNicIpAddress: {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The globally unique IP address",
									},
									isBareMetalServerNicIpHref: {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The URL for this reserved IP",
									},
									isBareMetalServerNicIpID: {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The unique identifier for this reserved IP",
									},
									isBareMetalServerNicIpName: {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The unique user-defined name for this reserved IP",
									},
									isBareMetalServerNicResourceType: {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The resource type: [ subnet_reserved_ip ]",
									},
								},
							},
						},
						isBareMetalServerNicResourceType: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The resource type : [ subnet_reserved_ip ]",
						},

						isBareMetalServerNicSecurityGroups: {
							Type:        schema.TypeSet,
							Computed:    true,
							Elem:        &schema.Schema{Type: schema.TypeString},
							Set:         schema.HashString,
							Description: "Collection of security groups ids",
						},

						isBareMetalServerNicStatus: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of the network interface : [ available, deleting, failed, pending ]",
						},

						isBareMetalServerNicSubnet: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the associated subnet",
						},

						isBareMetalServerNicType: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type of this bare metal server network interface : [ primary, secondary ]",
						},

						isBareMetalServerNicAllowedVlans: {
							Type:        schema.TypeSet,
							Computed:    true,
							Elem:        &schema.Schema{Type: schema.TypeInt},
							Set:         schema.HashInt,
							Description: "Indicates what VLAN IDs (for VLAN type only) can use this physical (PCI type) interface. A given VLAN can only be in the allowed_vlans array for one PCI type adapter per bare metal server.",
						},

						isBareMetalServerNicAllowInterfaceToFloat: {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Indicates if the interface can float to any other server within the same resource_group. The interface will float automatically if the network detects a GARP or RARP on another bare metal server in the resource group. Applies only to vlan type interfaces.",
						},

						isBareMetalServerNicVlan: {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Indicates the 802.1Q VLAN ID tag that must be used for all traffic on this interface",
						},
					},
				},
			},
		},
	}
}

func dataSourceIBMISBareMetalServerNetworkInterfacesRead(d *schema.ResourceData, meta interface{}) error {
	bareMetalServerID := d.Get(isBareMetalServerID).(string)
	err := bmsGetNetworkInterfaces(d, meta, bareMetalServerID)
	if err != nil {
		return err
	}
	return nil
}

func bmsGetNetworkInterfaces(d *schema.ResourceData, meta interface{}, bareMetalServerID string) error {
	sess, err := vpcClient(meta)
	if err != nil {
		return err
	}
	options := &vpcv1.ListBareMetalServerNetworkInterfacesOptions{
		BareMetalServerID: &bareMetalServerID,
	}
	nics := []vpcv1.BareMetalServerNetworkInterfaceIntf{}
	bmsNics, response, err := sess.ListBareMetalServerNetworkInterfaces(options)
	if err != nil || bmsNics == nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error listing Bare Metal Server (%s) network interfaces : %s\n%s", bareMetalServerID, err, response)
	}
	nics = append(nics, bmsNics.NetworkInterfaces...)
	nicsInfo := make([]map[string]interface{}, 0)
	for _, nicIntf := range nics {
		l := map[string]interface{}{}
		switch reflect.TypeOf(nicIntf).String() {
		case "*vpcv1.BareMetalServerNetworkInterfaceByPci":
			{
				nic := nicIntf.(*vpcv1.BareMetalServerNetworkInterfaceByPci)
				l[isBareMetalServerNicAllowIPSpoofing] = *nic.AllowIPSpoofing
				l[isBareMetalServerNicEnableInfraNAT] = *nic.EnableInfrastructureNat
				if nic.FloatingIps != nil {
					floatingIPList := make([]map[string]interface{}, 0)
					for _, ip := range nic.FloatingIps {
						currentIP := map[string]interface{}{
							// isBareMetalServerNicIpHref:    *ip.Href,
							isBareMetalServerNicIpID: *ip.ID,
							// isBareMetalServerNicIpCRN:     *ip.CRN,
							// isBareMetalServerNicIpName:    *ip.Name,
							isBareMetalServerNicIpAddress: *ip.Address,
						}
						floatingIPList = append(floatingIPList, currentIP)
					}
					l[isBareMetalServerNicFloatingIPs] = floatingIPList
				}
				l[isBareMetalServerNicHref] = *nic.Href
				l[isBareMetalServerNicInterfaceType] = *nic.InterfaceType
				if nic.Ips != nil {
					ipsList := make([]map[string]interface{}, 0)
					for _, ip := range nic.Ips {
						currentIP := map[string]interface{}{
							// isBareMetalServerNicIpHref:       *ip.Href,
							isBareMetalServerNicIpID: *ip.ID,
							// isBareMetalServerNicResourceType: *ip.ResourceType,
							// isBareMetalServerNicIpName:       *ip.Name,
							isBareMetalServerNicIpAddress: *ip.Address,
						}
						ipsList = append(ipsList, currentIP)
					}
					l[isBareMetalServerNicReservedIps] = ipsList
				}
				l[isBareMetalServerNicMacAddress] = *nic.MacAddress
				l[isBareMetalServerNicName] = *nic.Name
				if nic.PortSpeed != nil {
					l[isBareMetalServerNicPortSpeed] = *nic.PortSpeed
				}
				primaryIpList := make([]map[string]interface{}, 0)
				if nic.PrimaryIP != nil {
					currentIP := map[string]interface{}{
						// isBareMetalServerNicIpHref:       *nic.PrimaryIP.Href,
						isBareMetalServerNicIpID: *nic.PrimaryIP.ID,
						// isBareMetalServerNicResourceType: *nic.PrimaryIP.ResourceType,
						// isBareMetalServerNicIpName:       *nic.PrimaryIP.Name,
						isBareMetalServerNicIpAddress: *nic.PrimaryIP.Address,
					}
					primaryIpList = append(primaryIpList, currentIP)
				}
				l[isBareMetalServerNicPrimaryIP] = primaryIpList
				l[isBareMetalServerNicResourceType] = *nic.ResourceType
				if nic.SecurityGroups != nil && len(nic.SecurityGroups) != 0 {
					secgrpList := []string{}
					for i := 0; i < len(nic.SecurityGroups); i++ {
						secgrpList = append(secgrpList, string(*(nic.SecurityGroups[i].ID)))
					}
					l[isBareMetalServerNicSecurityGroups] = newStringSet(schema.HashString, secgrpList)
				}

				l[isBareMetalServerNicStatus] = *nic.Status
				l[isBareMetalServerNicSubnet] = *nic.Subnet.ID
				l[isBareMetalServerNicType] = *nic.Type
				if nic.AllowedVlans != nil {
					var out = make([]interface{}, len(nic.AllowedVlans), len(nic.AllowedVlans))
					for i, v := range nic.AllowedVlans {
						out[i] = int(v)
					}
					l[isBareMetalServerNicAllowedVlans] = schema.NewSet(schema.HashInt, out)
				}
			}
		case "*vpcv1.BareMetalServerNetworkInterfaceByVlan":
			{
				nic := nicIntf.(*vpcv1.BareMetalServerNetworkInterfaceByVlan)
				l[isBareMetalServerNicAllowIPSpoofing] = *nic.AllowIPSpoofing
				l[isBareMetalServerNicEnableInfraNAT] = *nic.EnableInfrastructureNat
				if nic.FloatingIps != nil {
					floatingIPList := make([]map[string]interface{}, 0)
					for _, ip := range nic.FloatingIps {
						currentIP := map[string]interface{}{
							// isBareMetalServerNicIpHref:    *ip.Href,
							isBareMetalServerNicIpID: *ip.ID,
							// isBareMetalServerNicIpCRN:     *ip.CRN,
							// isBareMetalServerNicIpName:    *ip.Name,
							isBareMetalServerNicIpAddress: *ip.Address,
						}
						floatingIPList = append(floatingIPList, currentIP)
					}
					l[isBareMetalServerNicFloatingIPs] = floatingIPList
				}
				l[isBareMetalServerNicHref] = *nic.Href
				l[isBareMetalServerNicInterfaceType] = *nic.InterfaceType
				if nic.Ips != nil {
					ipsList := make([]map[string]interface{}, 0)
					for _, ip := range nic.Ips {
						currentIP := map[string]interface{}{
							// isBareMetalServerNicIpHref:       *ip.Href,
							isBareMetalServerNicIpID: *ip.ID,
							// isBareMetalServerNicResourceType: *ip.ResourceType,
							// isBareMetalServerNicIpName:       *ip.Name,
							isBareMetalServerNicIpAddress: *ip.Address,
						}
						ipsList = append(ipsList, currentIP)
					}
					l[isBareMetalServerNicReservedIps] = ipsList
				}
				l[isBareMetalServerNicMacAddress] = *nic.MacAddress
				l[isBareMetalServerNicName] = *nic.Name
				if nic.PortSpeed != nil {
					l[isBareMetalServerNicPortSpeed] = *nic.PortSpeed
				}

				primaryIpList := make([]map[string]interface{}, 0)
				if nic.PrimaryIP != nil {
					currentIP := map[string]interface{}{
						// isBareMetalServerNicIpHref:       *nic.PrimaryIP.Href,
						isBareMetalServerNicIpID: *nic.PrimaryIP.ID,
						// isBareMetalServerNicResourceType: *nic.PrimaryIP.ResourceType,
						// isBareMetalServerNicIpName:       *nic.PrimaryIP.Name,
						isBareMetalServerNicIpAddress: *nic.PrimaryIP.Address,
					}
					primaryIpList = append(primaryIpList, currentIP)
				}
				l[isBareMetalServerNicPrimaryIP] = primaryIpList
				l[isBareMetalServerNicResourceType] = *nic.ResourceType
				if nic.SecurityGroups != nil && len(nic.SecurityGroups) != 0 {
					secgrpList := []string{}
					for i := 0; i < len(nic.SecurityGroups); i++ {
						secgrpList = append(secgrpList, string(*(nic.SecurityGroups[i].ID)))
					}
					l[isBareMetalServerNicSecurityGroups] = newStringSet(schema.HashString, secgrpList)
				}
				l[isBareMetalServerNicStatus] = *nic.Status
				l[isBareMetalServerNicSubnet] = *nic.Subnet.ID
				l[isBareMetalServerNicType] = *nic.Type
				l[isBareMetalServerNicAllowInterfaceToFloat] = *nic.AllowInterfaceToFloat
				l[isBareMetalServerNicVlan] = *nic.Vlan
			}
		}
		nicsInfo = append(nicsInfo, l)
	}
	d.SetId(dataSourceIBMISBareMetalServerNetworkInterfacesID(d))
	d.Set(isBareMetalServerNetworkInterfaces, nicsInfo)
	return nil
}

// dataSourceIBMISBMSProfilesID returns a reasonable ID for a BMS Profile list.
func dataSourceIBMISBareMetalServerNetworkInterfacesID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}
