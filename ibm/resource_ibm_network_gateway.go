package ibm

import (
	"fmt"
	"log"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/softlayer/softlayer-go/datatypes"
	"github.com/softlayer/softlayer-go/filter"
	"github.com/softlayer/softlayer-go/helpers/location"
	"github.com/softlayer/softlayer-go/helpers/product"
	"github.com/softlayer/softlayer-go/services"
	"github.com/softlayer/softlayer-go/session"
	"github.com/softlayer/softlayer-go/sl"
)

const packageKeyName = "NETWORK_GATEWAY_APPLIANCE"

const highAvailability = "HA"

func resourceIBMNetworkGateway() *schema.Resource {
	return &schema.Resource{
		Create:   resourceIBMNetworkGatewayCreate,
		Read:     resourceIBMNetworkGatewayRead,
		Update:   resourceIBMNetworkGatewayUpdate,
		Delete:   resourceIBMNetworkGatewayDelete,
		Exists:   resourceIBMNetworkGatewayExists,
		Importer: &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{

			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the gateway",
			},

			"ssh_key_ids": {
				Type:             schema.TypeList,
				Optional:         true,
				Elem:             &schema.Schema{Type: schema.TypeInt},
				ForceNew:         true,
				DiffSuppressFunc: applyOnce,
			},

			"post_install_script_uri": {
				Type:             schema.TypeString,
				Optional:         true,
				Default:          nil,
				ForceNew:         true,
				DiffSuppressFunc: applyOnce,
			},

			"private_ip_address_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},

			"private_vlan_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},

			"public_ip_address_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},

			"public_ipv6_address_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},

			"public_vlan_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},

			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"members": {
				Type:        schema.TypeList,
				Description: "The hardware members of this network Gateway",
				Required:    true,
				MinItems:    1,
				MaxItems:    2,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"member_id": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"hostname": {
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							DefaultFunc: genID,
							DiffSuppressFunc: func(k, o, n string, d *schema.ResourceData) bool {
								// FIXME: Work around another bug in terraform.
								// When a default function is used with an optional property,
								// terraform will always execute it on apply, even when the property
								// already has a value in the state for it. This causes a false diff.
								// Making the property Computed:true does not make a difference.
								if strings.HasPrefix(o, "terraformed-") && strings.HasPrefix(n, "terraformed-") {
									return true
								}
								return o == n
							},
						},

						"domain": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},

						"notes": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},

						"datacenter": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},

						"network_speed": {
							Type:     schema.TypeInt,
							Optional: true,
							Default:  100,
							ForceNew: true,
						},

						"tcp_monitoring": {
							Type:             schema.TypeBool,
							Optional:         true,
							Default:          false,
							ForceNew:         true,
							DiffSuppressFunc: applyOnce,
						},

						"process_key_name": {
							Type:             schema.TypeString,
							Optional:         true,
							ForceNew:         true,
							Default:          "INTEL_SINGLE_XEON_1270_3_40_2",
							DiffSuppressFunc: applyOnce,
						},

						"os_key_name": {
							Type:             schema.TypeString,
							Optional:         true,
							ForceNew:         true,
							Default:          "OS_VYATTA_5600_5_X_UP_TO_1GBPS_SUBSCRIPTION_EDITION_64_BIT",
							DiffSuppressFunc: applyOnce,
						},

						"redundant_network": {
							Type:             schema.TypeBool,
							Optional:         true,
							Default:          false,
							ForceNew:         true,
							DiffSuppressFunc: applyOnce,
						},
						"unbonded_network": {
							Type:             schema.TypeBool,
							Optional:         true,
							Default:          false,
							ForceNew:         true,
							DiffSuppressFunc: applyOnce,
						},
						"tags": {
							Type:     schema.TypeSet,
							Optional: true,
							ForceNew: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Set:      schema.HashString,
						},
						"public_bandwidth": {
							Type:             schema.TypeInt,
							Optional:         true,
							ForceNew:         true,
							Default:          20000,
							DiffSuppressFunc: applyOnce,
						},
						"memory": {
							Type:     schema.TypeInt,
							Required: true,
							//Sometime memory returns back as different. Since this resource is immutable at this point
							//and memory can't be really updated , suppress the change until we figure out how to handle it
							DiffSuppressFunc: applyOnce,
							ForceNew:         true,
						},
						"storage_groups": {
							Type:     schema.TypeList,
							Optional: true,
							ForceNew: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"array_type_id": {
										Type:     schema.TypeInt,
										Required: true,
									},
									"hard_drives": {
										Type:     schema.TypeList,
										Elem:     &schema.Schema{Type: schema.TypeInt},
										Required: true,
									},
									"array_size": {
										Type:     schema.TypeInt,
										Optional: true,
									},
									"partition_template_id": {
										Type:     schema.TypeInt,
										Optional: true,
									},
								},
							},
							DiffSuppressFunc: applyOnce,
						},

						"ssh_key_ids": {
							Type:             schema.TypeList,
							Optional:         true,
							Elem:             &schema.Schema{Type: schema.TypeInt},
							ForceNew:         true,
							DiffSuppressFunc: applyOnce,
						},

						"post_install_script_uri": {
							Type:             schema.TypeString,
							Optional:         true,
							Default:          nil,
							ForceNew:         true,
							DiffSuppressFunc: applyOnce,
						},

						"user_metadata": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},

						"disk_key_names": {
							Type:             schema.TypeList,
							Optional:         true,
							ForceNew:         true,
							Elem:             &schema.Schema{Type: schema.TypeString},
							DiffSuppressFunc: applyOnce,
						},

						"public_vlan_id": {
							Type:             schema.TypeInt,
							Optional:         true,
							ForceNew:         true,
							Computed:         true,
							DiffSuppressFunc: applyOnce,
						},

						"private_vlan_id": {
							Type:             schema.TypeInt,
							Optional:         true,
							ForceNew:         true,
							Computed:         true,
							DiffSuppressFunc: applyOnce,
						},

						"public_ipv4_address": {
							Type:     schema.TypeString,
							Computed: true,
						},

						"private_ipv4_address": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"ipv6_enabled": {
							Type:     schema.TypeBool,
							Optional: true,
							ForceNew: true,
							Default:  true,
						},

						"ipv6_address": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"private_network_only": {
							Type:     schema.TypeBool,
							Optional: true,
							Default:  false,
							ForceNew: true,
						},
					},
				},
			},

			"associated_vlans": {
				Type:        schema.TypeSet,
				Description: "The VLAN instances associated with this Network Gateway",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"vlan_id": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"network_vlan_id": {
							Type:        schema.TypeInt,
							Description: "The Identifier of the VLAN which is associated",
							Computed:    true,
						},
						"bypass": {
							Type:        schema.TypeBool,
							Description: "Indicates if the VLAN is in bypass or routed modes",
							Default:     nil,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func resourceIBMNetworkGatewayCreate(d *schema.ResourceData, meta interface{}) error {
	sess := meta.(ClientSession).SoftLayerSession()
	members := []gatewayMember{}
	for _, v := range d.Get("members").([]interface{}) {
		m := v.(map[string]interface{})
		members = append(members, m)
	}

	if len(members) == 2 {
		if !areVlanCompatible(members) {
			return fmt.Errorf("Members should have exactly same public and private vlan configuration," +
				"please check public_vlan_id and private_vlan_id property on individual members")
		}
	}
	d.Partial(true)

	//Build order for one member
	order, err := getMonthlyGatewayOrder(members[0], meta)
	if err != nil {
		return fmt.Errorf(
			"Encountered problem trying to get the Gateway order template: %s", err)
	}
	err = setHardwareOptions(members[0], &order.Hardware[0])
	if err != nil {
		return fmt.Errorf(
			"Encountered problem trying to configure Gateway options: %s", err)
	}

	// two members can be ordered together if they have same hardware configuration
	// and differ only in hostname, domain, user_metadata, post_install_script_uri etc
	sameOrder := canBeOrderedTogether(members)

	if sameOrder {
		//Ordering HA
		order.Quantity = sl.Int(2)
		order.Hardware = append(order.Hardware, datatypes.Hardware{
			Hostname: sl.String(members[1]["hostname"].(string)),
			Domain:   sl.String(members[1]["domain"].(string)),
		})
		err = setHardwareOptions(members[1], &order.Hardware[1])
		if err != nil {
			return fmt.Errorf(
				"Encountered problem trying to configure Gateway options: %s", err)
		}
	}
	// Set SSH Key on main order
	ssh_key_ids := d.Get("ssh_key_ids").([]interface{})
	if len(ssh_key_ids) > 0 {
		order.SshKeys = make([]datatypes.Container_Product_Order_SshKeys, 0, len(ssh_key_ids))
		for _, ssh_key_id := range ssh_key_ids {
			sshKeyA := make([]int, 1)
			sshKeyA[0] = ssh_key_id.(int)
			order.SshKeys = append(order.SshKeys, datatypes.Container_Product_Order_SshKeys{
				SshKeyIds: sshKeyA,
			})
		}
	}
	// Set post_install_script_uri on main order
	if v, ok := d.GetOk("post_install_script_uri"); ok {
		order.ProvisionScripts = []string{v.(string)}
	}

	var productOrder datatypes.Container_Product_Order
	productOrder.OrderContainers = []datatypes.Container_Product_Order{order}

	_, err = services.GetProductOrderService(sess).VerifyOrder(&productOrder)
	if err != nil {
		return fmt.Errorf(
			"Encountered problem trying to verify the order: %s", err)
	}
	_, err = services.GetProductOrderService(sess.SetRetries(0)).PlaceOrder(&productOrder, sl.Bool(false))
	if err != nil {
		return fmt.Errorf(
			"Encountered problem trying to place the order: %s", err)
	}

	bm, err := waitForNetworkGatewayMemberProvision(&order.Hardware[0], meta)
	if err != nil {
		return fmt.Errorf(
			"Error waiting for Gateway (%s) to become ready: %s", d.Id(), err)
	}

	id := *bm.(datatypes.Hardware).NetworkGatewayMember.NetworkGatewayId
	d.SetId(fmt.Sprintf("%d", id))
	log.Printf("[INFO] Gateway ID: %s", d.Id())

	member1Id := *bm.(datatypes.Hardware).Id
	members[0]["member_id"] = member1Id
	log.Printf("[INFO] Member 1 ID: %d", member1Id)
	d.SetPartial("members")

	err = setTagsAndNotes(members[0], meta)
	if err != nil {
		return err
	}

	if sameOrder {
		// If we ordered HA and then wait for other member
		bm, err := waitForNetworkGatewayMemberProvision(&order.Hardware[1], meta)
		if err != nil {
			return fmt.Errorf(
				"Error waiting for Gateway (%s) to become ready: %s", d.Id(), err)
		}
		member2Id := *bm.(datatypes.Hardware).Id
		log.Printf("[INFO] Member 2 ID: %d", member2Id)
		members[1]["member_id"] = member2Id
		err = setTagsAndNotes(members[1], meta)
		if err != nil {
			return err
		}
		d.SetPartial("members")
	} else if len(members) == 2 {
		//Add the new gateway which has different configuration than the first
		err := addGatewayMember(id, members[1], meta)
		if err != nil {
			return err
		}
		d.SetPartial("members")
	}

	name := d.Get("name").(string)
	err = updateGatewayName(id, name, meta)
	if err != nil {
		return err
	}

	d.Partial(false)
	return resourceIBMNetworkGatewayRead(d, meta)
}

func resourceIBMNetworkGatewayRead(d *schema.ResourceData, meta interface{}) error {
	service := services.GetNetworkGatewayService(meta.(ClientSession).SoftLayerSession())
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return fmt.Errorf("Not a valid ID, must be an integer: %s", err)
	}
	result, err := service.Id(id).Mask(
		"insideVlans,members,status," +
			"members[hardware],members[hardware[datacenter]]," +
			"members[hardware[primaryNetworkComponent]],members[hardware[backendNetworkComponents,primaryBackendNetworkComponent[redundancyEnabledFlag]," +
			"tagReferences,primaryIpAddress,primaryBackendIpAddress," +
			"primaryNetworkComponent[primaryVersion6IpAddressRecord],privateNetworkOnlyFlag," +
			"powerSupplyCount,primaryNetworkComponent[networkVlan],memoryCapacity,networkVlans[id,vlanNumber]]]",
	).GetObject()
	if err != nil {
		return fmt.Errorf("Error retrieving Network Gateway: %s", err)
	}
	d.Set("name", result.Name)
	d.Set("private_ip_address_id", result.PrivateIpAddressId)
	d.Set("private_vlan_id", result.PrivateVlanId)
	d.Set("public_ip_address_id", result.PublicIpAddressId)
	d.Set("public_ipv6_address_id", result.PublicIpv6AddressId)
	d.Set("public_vlan_id", result.PublicVlanId)
	d.Set("status", result.Status.Name)

	err = d.Set("members", flattenGatewayMembers(d, result.Members))
	if err != nil {
		return err
	}
	d.Set("associated_vlans", flattenGatewayVlans(result.InsideVlans))

	return nil
}

func updateGatewayName(id int, name string, meta interface{}) error {
	sess := meta.(ClientSession).SoftLayerSession()
	service := services.GetNetworkGatewayService(sess)
	_, err := service.Id(id).EditObject(&datatypes.Network_Gateway{
		Name: sl.String(name),
	})
	if err != nil {
		return fmt.Errorf("Couldn't set the gateway name to %s", name)
	}
	return err
}

func addGatewayMember(gwID int, member gatewayMember, meta interface{}) error {
	sess := meta.(ClientSession).SoftLayerSession()
	order, err := getMonthlyGatewayOrder(member, meta)
	if err != nil {
		return fmt.Errorf(
			"Encountered problem trying to get the Gateway order template: %s", err)
	}
	err = setHardwareOptions(member, &order.Hardware[0])
	if err != nil {
		return fmt.Errorf(
			"Encountered problem trying to configure Gateway options: %s", err)
	}

	haOrder := datatypes.Container_Product_Order_Hardware_Server_Gateway_Appliance{}
	haOrder.ContainerIdentifier = order.ContainerIdentifier
	haOrder.Hardware = order.Hardware
	haOrder.PackageId = order.PackageId
	haOrder.Location = order.Location
	haOrder.Prices = order.Prices
	haOrder.ClusterResourceId = sl.Int(gwID)
	haOrder.ClusterOrderType = sl.String(highAvailability)

	_, err = services.GetProductOrderService(sess).VerifyOrder(&haOrder)
	if err != nil {
		return fmt.Errorf(
			"Encountered problem trying to verify the order: %s", err)
	}
	_, err = services.GetProductOrderService(sess.SetRetries(0)).PlaceOrder(&haOrder, sl.Bool(false))
	if err != nil {
		return fmt.Errorf(
			"Encountered problem trying to place the order: %s", err)
	}

	bm, err := waitForNetworkGatewayMemberProvision(&order.Hardware[0], meta)
	if err != nil {
		return fmt.Errorf(
			"Error waiting for Gateway (%d) to become ready: %s", gwID, err)
	}
	id := *bm.(datatypes.Hardware).Id
	log.Printf("[INFO] Newly added member ID: %d", id)
	member["member_id"] = id
	err = setTagsAndNotes(member, meta)
	return err
}

func resourceIBMNetworkGatewayUpdate(d *schema.ResourceData, meta interface{}) error {
	id, _ := strconv.Atoi(d.Id())
	if d.HasChange("name") {
		gwName := d.Get("name").(string)
		err := updateGatewayName(id, gwName, meta)
		if err != nil {
			return err
		}
	}
	return resourceIBMNetworkGatewayRead(d, meta)
}

func resourceIBMNetworkGatewayDelete(d *schema.ResourceData, meta interface{}) error {
	sess := meta.(ClientSession).SoftLayerSession()
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return fmt.Errorf("Not a valid ID, must be an integer: %s", err)
	}
	service := services.GetNetworkGatewayService(sess)
	gw, err := service.Id(id).Mask("members[hardwareId]").GetObject()
	for _, v := range gw.Members {
		m := gatewayMember{
			"member_id": *v.HardwareId,
		}
		err := deleteHardware(m, meta)
		if err != nil {
			return err
		}
	}
	//If both the hardwares have been deleted then gateway will go away as well
	d.SetId("")
	return nil
}

func resourceIBMNetworkGatewayExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	service := services.GetNetworkGatewayService(meta.(ClientSession).SoftLayerSession())

	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return false, fmt.Errorf("Not a valid ID, must be an integer: %s", err)
	}

	result, err := service.Id(id).GetObject()
	if err != nil {
		if apiErr, ok := err.(sl.Error); !ok || apiErr.StatusCode != 404 {
			return false, fmt.Errorf("Error trying to retrieve Network Gateway: %s", err)
		}
	}

	return result.Id != nil && *result.Id == id, nil
}

func getMonthlyGatewayOrder(d dataRetriever, meta interface{}) (datatypes.Container_Product_Order, error) {
	sess := meta.(ClientSession).SoftLayerSession()

	// Validate attributes for network gateway ordering.
	model := packageKeyName

	datacenter := d.Get("datacenter")

	osKeyName := d.Get("os_key_name")

	process_key_name := d.Get("process_key_name")

	dc, err := location.GetDatacenterByName(sess, datacenter.(string), "id")
	if err != nil {
		return datatypes.Container_Product_Order{}, err
	}

	// 1. Find a package id using Gateway package key name.
	pkg, err := getPackageByModelGateway(sess, model)

	if err != nil {
		return datatypes.Container_Product_Order{}, err
	}

	if pkg.Id == nil {
		return datatypes.Container_Product_Order{}, err
	}

	// 2. Get all prices for the package
	items, err := product.GetPackageProducts(sess, *pkg.Id, productItemMaskWithPriceLocationGroupID)
	if err != nil {
		return datatypes.Container_Product_Order{}, err
	}

	// 3. Build price items
	server, err := getItemPriceId(items, "server", process_key_name.(string))
	if err != nil {
		return datatypes.Container_Product_Order{}, err
	}

	os, err := getItemPriceId(items, "os", osKeyName.(string))
	if err != nil {
		return datatypes.Container_Product_Order{}, err
	}

	ram, err := findMemoryItemPriceId(items, d)
	if err != nil {
		return datatypes.Container_Product_Order{}, err
	}

	portSpeed, err := findNetworkItemPriceId(items, d)
	if err != nil {
		return datatypes.Container_Product_Order{}, err
	}

	monitoring, err := getItemPriceId(items, "monitoring", "MONITORING_HOST_PING")
	if err != nil {
		return datatypes.Container_Product_Order{}, err
	}
	if d.Get("tcp_monitoring").(bool) {
		monitoring, err = getItemPriceId(items, "monitoring", "MONITORING_HOST_PING_AND_TCP_SERVICE")
		if err != nil {
			return datatypes.Container_Product_Order{}, err
		}
	}
	// Other common default options
	priIpAddress, err := getItemPriceId(items, "pri_ip_addresses", "1_IP_ADDRESS")
	if err != nil {
		return datatypes.Container_Product_Order{}, err
	}

	remoteManagement, err := getItemPriceId(items, "remote_management", "REBOOT_KVM_OVER_IP")
	if err != nil {
		return datatypes.Container_Product_Order{}, err
	}
	vpnManagement, err := getItemPriceId(items, "vpn_management", "UNLIMITED_SSL_VPN_USERS_1_PPTP_VPN_USER_PER_ACCOUNT")
	if err != nil {
		return datatypes.Container_Product_Order{}, err
	}

	notification, err := getItemPriceId(items, "notification", "NOTIFICATION_EMAIL_AND_TICKET")
	if err != nil {
		return datatypes.Container_Product_Order{}, err
	}
	response, err := getItemPriceId(items, "response", "AUTOMATED_NOTIFICATION")
	if err != nil {
		return datatypes.Container_Product_Order{}, err
	}
	vulnerabilityScanner, err := getItemPriceId(items, "vulnerability_scanner", "NESSUS_VULNERABILITY_ASSESSMENT_REPORTING")
	if err != nil {
		return datatypes.Container_Product_Order{}, err
	}

	// Define an order object using basic paramters.

	order := datatypes.Container_Product_Order{
		ContainerIdentifier: sl.String(d.Get("hostname").(string)),
		Quantity:            sl.Int(1),
		Hardware: []datatypes.Hardware{
			{
				Hostname: sl.String(d.Get("hostname").(string)),
				Domain:   sl.String(d.Get("domain").(string)),
			},
		},
		Location:  sl.String(strconv.Itoa(*dc.Id)),
		PackageId: pkg.Id,
		Prices: []datatypes.Product_Item_Price{
			server,
			os,
			ram,
			portSpeed,
			priIpAddress,
			remoteManagement,
			vpnManagement,
			monitoring,
			notification,
			response,
			vulnerabilityScanner,
		},
	}

	// Add optional price ids.
	// Add public bandwidth

	publicBandwidth := d.Get("public_bandwidth")
	publicBandwidthStr := "BANDWIDTH_" + strconv.Itoa(publicBandwidth.(int)) + "_GB"
	bandwidth, err := getItemPriceId(items, "bandwidth", publicBandwidthStr)
	if err != nil {
		return datatypes.Container_Product_Order{}, err
	}
	order.Prices = append(order.Prices, bandwidth)
	privateNetworkOnly := d.Get("private_network_only").(bool)
	if d.Get("ipv6_enabled").(bool) {
		if privateNetworkOnly {
			return datatypes.Container_Product_Order{}, fmt.Errorf("Unable to configure a public IPv6 address with a private_network_only option")
		}
		keyName := "1_IPV6_ADDRESS"
		price, err := getItemPriceId(items, "pri_ipv6_addresses", keyName)
		if err != nil {
			return datatypes.Container_Product_Order{}, err
		}
		order.Prices = append(order.Prices, price)
	}
	// Add prices of disks.
	disks := d.Get("disk_key_names").([]interface{})
	diskLen := len(disks)
	if diskLen > 0 {
		for i, disk := range disks {
			diskPrice, err := getItemPriceId(items, "disk"+strconv.Itoa(i), disk.(string))
			if err != nil {
				return datatypes.Container_Product_Order{}, err
			}
			order.Prices = append(order.Prices, diskPrice)
		}
	}

	// Add storage_groups for RAID configuration
	diskController, err := getItemPriceId(items, "disk_controller", "DISK_CONTROLLER_NONRAID")
	if err != nil {
		return datatypes.Container_Product_Order{}, err
	}

	if _, ok := d.GetOk("storage_groups"); ok {
		order.StorageGroups = getStorageGroupsFromResourceData(d)
		diskController, err = getItemPriceId(items, "disk_controller", "DISK_CONTROLLER_RAID")
		if err != nil {
			return datatypes.Container_Product_Order{}, err
		}
	}
	order.Prices = append(order.Prices, diskController)

	return order, nil
}

func getPackageByModelGateway(sess *session.Session, model string) (datatypes.Product_Package, error) {
	objectMask := "id,keyName,name,description,isActive,type[keyName],categories[id,name,categoryCode]"
	service := services.GetProductPackageService(sess)
	availableModels := ""
	filterStr := "{\"items\": {\"categories\": {\"categoryCode\": {\"operation\":\"server\"}}},\"type\": {\"keyName\": {\"operation\":\"BARE_METAL_GATEWAY\"}}}"

	// Get package id
	packages, err := service.Mask(objectMask).
		Filter(filterStr).GetAllObjects()
	if err != nil {
		return datatypes.Product_Package{}, err
	}
	for _, pkg := range packages {
		availableModels = availableModels + *pkg.KeyName
		if pkg.Description != nil {
			availableModels = availableModels + " ( " + *pkg.Description + " ), "
		} else {
			availableModels = availableModels + ", "
		}
		if *pkg.KeyName == model {
			return pkg, nil
		}
	}
	return datatypes.Product_Package{}, fmt.Errorf("No Gateway package key name for %s. Available package key name(s) is(are) %s", model, availableModels)
}

func setHardwareOptions(m gatewayMember, hardware *datatypes.Hardware) error {
	public_vlan_id := m.Get("public_vlan_id").(int)
	if public_vlan_id > 0 {
		hardware.PrimaryNetworkComponent = &datatypes.Network_Component{
			NetworkVlan: &datatypes.Network_Vlan{Id: sl.Int(public_vlan_id)},
		}
	}

	private_vlan_id := m.Get("private_vlan_id").(int)
	if private_vlan_id > 0 {
		hardware.PrimaryBackendNetworkComponent = &datatypes.Network_Component{
			NetworkVlan: &datatypes.Network_Vlan{Id: sl.Int(private_vlan_id)},
		}
	}

	if userMetadata, ok := m.GetOk("user_metadata"); ok {
		hardware.UserData = []datatypes.Hardware_Attribute{
			{Value: sl.String(userMetadata.(string))},
		}
	}

	if v, ok := m.GetOk("post_install_script_uri"); ok {
		hardware.PostInstallScriptUri = sl.String(v.(string))
	}

	// Get configured ssh_keys
	ssh_key_ids := m.Get("ssh_key_ids").([]interface{})
	if len(ssh_key_ids) > 0 {
		hardware.SshKeys = make([]datatypes.Security_Ssh_Key, 0, len(ssh_key_ids))
		for _, ssh_key_id := range ssh_key_ids {
			hardware.SshKeys = append(hardware.SshKeys, datatypes.Security_Ssh_Key{
				Id: sl.Int(ssh_key_id.(int)),
			})
		}
	}

	return nil
}

// Network gateways or Bare metal creation does not return a  object with an Id.
// Have to wait on provision date to become available on server that matches
// hostname and domain.
// http://sldn.softlayer.com/blog/bpotter/ordering-bare-metal-servers-using-softlayer-api
func waitForNetworkGatewayMemberProvision(d *datatypes.Hardware, meta interface{}) (interface{}, error) {
	hostname := *d.Hostname
	domain := *d.Domain
	log.Printf("Waiting for Gateway (%s.%s) to be provisioned", hostname, domain)

	stateConf := &resource.StateChangeConf{
		Pending: []string{"retry", "pending"},
		Target:  []string{"provisioned"},
		Refresh: func() (interface{}, string, error) {
			service := services.GetAccountService(meta.(ClientSession).SoftLayerSession())
			bms, err := service.Filter(
				filter.Build(
					filter.Path("hardware.hostname").Eq(hostname),
					filter.Path("hardware.domain").Eq(domain),
				),
			).Mask("id,provisionDate,networkGatewayMember[networkGatewayId]").GetHardware()
			if err != nil {
				return false, "retry", nil
			}

			if len(bms) == 0 || bms[0].ProvisionDate == nil {
				return datatypes.Hardware{}, "pending", nil
			} else {
				return bms[0], "provisioned", nil
			}
		},
		Timeout:        24 * time.Hour,
		Delay:          10 * time.Second,
		MinTimeout:     1 * time.Minute,
		NotFoundChecks: 24 * 60,
	}

	return stateConf.WaitForState()
}

func setTagsAndNotes(m gatewayMember, meta interface{}) error {
	err := setHardwareTags(m["member_id"].(int), m, meta)
	if err != nil {
		return err
	}

	if m["notes"].(string) != "" {
		err := setHardwareNotes(m["member_id"].(int), m, meta)
		if err != nil {
			return err
		}
	}
	return nil
}

//New types to resuse functions from other resources which does the same job
//Essentially mimic schema.ResourceData get functions
type dataRetriever interface {
	Get(string) interface{}
	GetOk(string) (interface{}, bool)
	Id() string
}
type gatewayMember map[string]interface{}

func (m gatewayMember) Get(k string) interface{} {
	if k == "restricted_network" || k == "hourly_billing" {
		//findNetworkItemPriceId is used from bare metal and that looks for this key
		//deleteHardware looks for hourly_billing
		//We won't need this when we support those speed on the gateway
		return false
	}
	return m[k]
}
func (m gatewayMember) GetOk(k string) (i interface{}, ok bool) {
	i, ok = m[k]
	if ok {
		if k == "storage_groups" {
			return i, len(i.([]interface{})) > 0
		}
		if k == "user_metadata" || k == "post_install_script_uri" {
			return i, len(i.(string)) > 0
		}
	}
	return
}

func (m gatewayMember) Id() string {
	return strconv.Itoa(m["member_id"].(int))
}

func areVlanCompatible(m []gatewayMember) bool {
	if m[0]["public_vlan_id"].(int) != m[1]["public_vlan_id"].(int) {
		return false
	}
	if m[0]["private_vlan_id"].(int) != m[1]["private_vlan_id"].(int) {
		return false
	}
	return true
}

func canBeOrderedTogether(members []gatewayMember) bool {
	if len(members) != 2 {
		return false
	}
	m1 := members[0]
	m2 := members[1]
	for k, v := range m1 {
		if k == "hostname" ||
			k == "domain" ||
			k == "notes" ||
			k == "tags" ||
			k == "public_vlan_id" ||
			k == "private_vlan_id" ||
			k == "user_metadata" ||
			k == "post_install_script_uri" {
			continue
		}

		// If other harware configurations are not equal then they can't be ordered together
		// For example different memory
		if !reflect.DeepEqual(v, m2[k]) {
			return false
		}
	}
	return true
}
