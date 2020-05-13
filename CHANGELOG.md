## 1.5.2 (May 07, 2020)

ENAHANCEMENTS:

* resource : Support for entitlement argument for IKS Classic ROKS cluster (ibm_container_cluster) and worker pool(ibm_container_worker_pool)([#1350](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1350))

* resource : Support for source_resource_group_id and target_resource_group_id([#1364](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1364))

BUG FIXES:

* resource : Error deleting instance with data volume ([#1412](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1412))
* resource : Add force_new true for cidr argument of ibm_is_address_prefix ([#1416](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1416))
* resource : Fix import of ibm_container_cluster ([#1360](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1360))

## 1.5.1 (May 04, 2020)
BUG FIXES:

* resource : Fix VPC subnets created in incorrect resource group([#1398](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1398))

## 1.5.0 (April 29, 2020)
FEATURES:

* New Resource: ([ibm_is_lb_listener_policy_rule](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1147) )
* New Datasource: ([ibm_certificate_manager_certificates](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1277) )

ENAHANCEMENTS:

* resource : Support for auto-generate client_id and client_id for API gateway endpoint([#1390](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1390))
* resource: Support point_in_time_recovery_time and point_in_time_recovery_deployment_id arguments for ICD database([#1259](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1259))
* resource: Support for pending_reclamination for database and CIS instances ([#1242](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1242))


BUG FIXES:

* resource : Fix VPC Load Balancer resource ID is appended to Pool/Listener/Listener Policy ID ([#1359](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1359))
* resource : Fix domainID for CIS firewall resource ([#1201](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1201))
* resource : Fix the update of private dns resource record TTL ([#1331](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1331))
* resource : ibm_container_worker_pool_zone_attachment should wait for ALBs to finish in a new zone ([#1372](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1372))


## 1.4.0 (April 16, 2020)
  
NOTE :  For creating either vpc-classic (generation=1) or vpc-Gen2 (generation=1) IKS cluster, generation parameter needs to be set either in provider block or export via environment variable “IC_GENERATION”. By default the generation value is 2. 

FEATURES:

* New Resource: ([Terraform support for DNS service (beta service ) ibm_dns_zone, ibm_dns_permitted_network, ibm_dns_resource_record](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/958 ))
* New Resource: ([ibm_cis_firewall (lockdown)](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1201) )
* New Resource: ([ibm_lb_listener_policy](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1147) )

ENAHANCEMENTS:

* resource : Add support for resource group argument in ibm_is_network_acl ([#1265](https://github.com/IBM-Cloud/terraform-provider-ibm/issue/1265))
* resource : Support for IKS on Gen-2  (beta service) ([#1321](https://github.com/IBM-Cloud/terraform-provider-ibm/issue/1321))
* resource : Update functionality support for cis resources  ([#1180](https://github.com/IBM-Cloud/terraform-provider-ibm/issue/1180))
* resource : Add support for crn attribute for is_vpc   ([#1315](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/1315) )
* data : Add support for crn attribute for is_vpc   ([#1317](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/1317))

BUG FIXES:

* resource :  Fix the nil pointer exception for ibm_is_lb_listener resource ([#1289](https://github.com/IBM-Cloud/terraform-provider-ibm/issue/1289))


## 1.3.0 (April 02, 2020)

FEATURES:

* New Resource: ([ibm_iam_access_group_dynamic_rule](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/691))
* New Resource: ([ibm_api_gateway_endpoint](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1247))
* New Resource: ([ibm_api_gateway_endpoint_subscription](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1247))
* New DataSource: ([ibm_iam_access_group](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/953))
* New DataSource: ([ibm_api_gateway](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1247))

BUG FIXES:

* resource : Fix the destroy of cloudantnosqldb service([#1242](https://github.com/IBM-Cloud/terraform-provider-ibm/issue/1242)) 
* resource : Fix the ICD service endpoint for osl01([#1158](https://github.com/IBM-Cloud/terraform-provider-ibm/issue/1158)) 


## 1.2.6 (March 26, 2020)

ENHANCEMENTS:

* resource : Added support for cse_source_addresses  attribute for ibm_is_vpc  ([#1165](https://github.com/IBM-Cloud/terraform-provider-ibm/issue/1165)) 
* data : Added support for cse_source_addresses  attribute for ibm_is_vpc ([#1165](https://github.com/IBM-Cloud/terraform-provider-ibm/issue/1165)) 
* resource: Added support for new storage class smart for COS bucket  ([#1184](https://github.com/IBM-Cloud/terraform-provider-ibm/issue/1184))
* resource:  Allow deletion of non-existing resources like is_vpc, is_subnet, is_vpc_address_prefix and is_instance  ([#1229](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/1229))
* resource:  Added support for force_delete argument for ibm_kp_key ([#1214](https://github.com/IBM-Cloud/terraform-provider-ibm/issue/1214))


## 1.2.5 (March 19, 2020)

ENHANCEMENTS:

* Provider : Adapt IAM access resources to v2 version ([#1183](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/1183)) 
* resource: Added support for GUID attribute for ibm_cis and ibm_database ([#1169](https://github.com/IBM-Cloud/terraform-provider-ibm/issue/1169)) 
* data: Added support for GUID attribute for ibm_cis and ibm_database ([#1169](https://github.com/IBM-Cloud/terraform-provider-ibm/issue/1169))

BUG FIXES:

* resources : Updated the status string for `ibm_resource_instance, ibm_database and ibm_cis` to be inline with resource controller API changes ([#1190](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1190)) 
* resource/ibm_compute_bare_metal: Fix the order of provisioning of `bare metal` for processor capacity restriction type and SAP servers ([#1189](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1189)) 
* resource/ibm_resource_instance: Fix the order of provisioning of `block chain` platform service ([#1186](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1186)) 
* resource/ibm_container_cluster: Fix the force new for deprecated `billing` argument([#1187](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1187))  


## 1.2.4 (March 11, 2020)

ENHANCEMENTS:

* Provider: Added new parameter `zone` to support power virtual resources and data sources to work in multi-zone environment ([#1141](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1141))
* resource/ibm_pi_volume: Updated the list of volume types for power virtual volume ([#1149](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/1149))
* resource/ibm_container_vpc_cluster : Added support for `ingress_hostname` and `ingress_secret` attributes ([#1167](https://github.com/IBM-Cloud/terraform-provider-ibm/issue/1167))
* data/ibm_container_vpc_cluster : Added support for `ingress_hostname` and `ingress_secret` attributes ([#1167](https://github.com/IBM-Cloud/terraform-provider-ibm/issue/1167))
* resource/ibm_is_floating_ip : Handle the case when floating IP is deleted manually ([#1160](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/1160))

BUG FIXES:

* resources : Handle the case where the resource might be already deleted (manually) for ibm_iam_access_policies, ibm_iam_authorization_policies, ibm_iam_service_policies ([#1162](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1162))
* resource/ibm_is_inetwork_acl: Fix the order of creation of network acl ([#1123](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1123))
* resource/ibm_container_vpc_cluster: Added new attribute `wait_till` to control the cluster creation. Now user can control the cluster creation until master is ready / any one worker node is ready / ingress_hostname is  
  assigned.  ([#1143](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1143))
* resource/ibm_pi_instance: Fix the timeout configuration for create ([#1178](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/1178))
* doc/ibm_cis_ip_addresses : Fix the description of data source ([#1178](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/1178))

## 1.2.3 (March 03, 2020)

BUG FIXES:

* data/ibm_container_cluster_config : Fix the error to download the cluster config for VPC clusters ([#1150](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/1150))

## 1.2.2 (February 26, 2020)

ENHANCEMENTS:

* resource/ibm_is_vpc: Improved error message for VPC creation ([#1106](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/1106))
* resource/ibm_is_ssh_key: Improved error message for VPC SSH Key creation ([#1105](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/1105))
* resource/ibm_container_cluster : Added gateway feature support for IKS clusters. This feature helps to create a cluster with a gateway worker pool of two gateway worker nodes that are connected to public and private VLANs to provide limited public access, and a compute worker pool of compute worker nodes that are connected to the private VLAN only. 
([#1125](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/1125))
* data/ibm_conatiner_cluster_config : Extended the data source to provide additional attribute like admin_key, admin_certificate, ca_certificate, host and token. This attributes helps to connect to other providers like Kubernetes and Helm without loading cluster config file. ([#895](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/895))

BUG FIXES:

* doc/ibm_certificate_manager_order: Changed the type of rotate_key from string to bool ([#1110](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/1110))
* resource/ibm_is_instance: Fix for updating security group for primary network interface for vpc instance. Now users can add or delete security groups([#1078](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1078))
* doc/ibm_resource_key : Provided an example in the docs as a workaround to create credentials using serviceID parameter ([#1121](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/1121))
* resource/ibm_is_network_acl : Fix for crash during the update of rules. Fix for the order of rules creation. Now users can add or delete rules for network_acl ([#1117](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1117))
* resource/ibm_is_public_gateway : Added support for resource group and tags parameters ([#1102](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1102))
* resource/ibm_is_floating_ip : Added support for tags parameters ([#1131](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/1131))
* resource/ibm_database : Parameters remote_leader_id, key_protect_instance and key_protect_key can’t be updated after creation. ([#1111](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1111))
* example/ibm-key-protect : Updated example to create an authorisation policy between COS and Key Protect instance([#1133](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/1133))
* resource/ibm_resource_group: Removed suppression of error during deletion ([#1108](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/1108))
* resource/ibm_iam_user_invite : Fix for inviting user from IBM Cloud lite account. ([#1114](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1114))


