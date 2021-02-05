# Azure - Firewall Management

## Introduction

This repository manages all YAML files for Azure firewall rules.<br />
Converting these rules into Terraform resources to manage the firewall between our internal network and Azure.<br />
<br />

<!--- BEGIN_TF_DOCS --->
## Providers

| Name | Version |
|------|---------|
| azurerm | >=2.0.0 |
| random | n/a |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:-----:|
| expressroute\_firewall\_name | n/a | `string` | `"app-expressroute-prod-useast2-firewall"` | no |
| expressroute\_resource\_group | n/a | `string` | `"app-expressroute-prod-useast2"` | no |
| expressroute\_subscription\_id | n/a | `string` | `"f77593b8-c144-4ed2-9038-fa8d1dabc54a"` | no |

## Outputs

No output.
<!--- END_TF_DOCS --->