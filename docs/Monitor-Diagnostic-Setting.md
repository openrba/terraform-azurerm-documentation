# terraform-azurerm-monitor-diagnostic-setting
Azure monitor diagnostic setting module may be attached to Azure resources to configure setting such as retention policy. 
See example for usage. 

<!--- BEGIN_TF_DOCS --->
## Providers

| Name | Version |
|------|---------|
| azurerm | n/a |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:-----:|
| ds\_log\_api\_endpoints | Azure monitor diagnostic setting log API endpoints relevent to target resource. Value should be {"api\_endpoint = retention\_days"}. For example: {"MySqlSlowLogs" = 90, "MySqlAuditLogs" = 30} | `map` | `{}` | no |
| ds\_metrics\_retention\_days | Azure monitor diagnostic setting category for retention in days of target resource. | `map` | `{}` | no |
| metric\_category | Metric category | `string` | `"AllMetrics"` | no |
| sa\_resource\_group | Name of resource group the storage account resides in | `string` | n/a | yes |
| storage\_account | Name of storage account | `string` | n/a | yes |
| target\_resource\_id | Target resource id | `string` | n/a | yes |
| target\_resource\_name | Target resource name which will be also be applied as the name of monitor diagnostic setting. | `string` | n/a | yes |

## Outputs

No output.
<!--- END_TF_DOCS --->


