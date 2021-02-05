# Azure - Storage Container  Module

## Introduction

This module will create storage containers to a storage account.<br />
<br />


<!--- BEGIN_TF_DOCS --->
## Providers

| Name | Version |
|------|---------|
| azurerm | >= 2.24.0 |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:-----:|
| containers | List of containers to create and their access levels. | `list(object({ name = string, access_type = string }))` | `[]` | no |
| resource\_group | Name of Azure Resource Group the storage account exist in. | `string` | n/a | yes |
| storage\_account\_name | Azure storage account name. | `string` | n/a | yes |
| tags | Map of tags. | `map` | n/a | yes |

## Outputs

No output.
<!--- END_TF_DOCS --->
