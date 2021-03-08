provider "ibm" {
  ibmcloud_api_key = var.ibmcloud_api_key
}

// Create is_dedicated_host data source
data "ibm_is_dedicated_host" "is_dedicated_host_instance" {
  id = var.is_dedicated_host_id
}

// Create is_dedicated_hosts data source
data "ibm_is_dedicated_hosts" "is_dedicated_hosts_instance" {
  id = var.is_dedicated_hosts_id
}