---
layout: "ibm"
page_title: "IBM: ibm_schematics_workspace"
sidebar_current: "docs-ibm-schematics-workspace"
description: |-
  Get information about the terraform State store values of a specific template in a Schematics Workspace .
---

# ibm\_schematics_workspace


Import details of a schematics workspace as a read-only data source. You can then reference the argument fields of the data source in other resources within the same configuration using interpolation syntax.


## Example Usage

```hcl
data "ibm_schematics_worksapce" "test" {
  workspace_id = "my-worspace-id"
}
```

## Argument Reference

The following arguments are supported:

* `workspace_id` - (Required, string) The ID of the Schematics workspace.

## Attribute Reference

The following attributes are exported:

* `name` - The name of the workspace.
* `types` - The Terraform version supported types.
* `status` - The status of workspace.
* `is_frozen` - The frozen flag for the workspace.
* `is_locked` -  The locked flag of the workspace.
* `template_id` - The ID of the templates that are present in the workspace.
* `tags` - The tags suppoprted by workspace.
* `resource_group` - The resource group associated with the workspace.
* `resource_controller_url` - The URL of the IBM Cloud dashboard that can be used to explore and view details about the workspace.

