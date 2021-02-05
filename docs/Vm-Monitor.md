# Azure Monitor and Metric Alerts for Storage Accounts

Metric alerts in Azure Monitor provide a way to get notified when one of your metrics crosses a threshold.

### Integration

Module can be used for the following environments:

- Sandbox
- Development
- Production

### Required

- Subscription for one of the above environments
- Resource Group
- ID of Monitored Resource
- Action Group ID

# Usage

<!--- BEGIN_TF_DOCS --->
## Providers

| Name | Version |
|------|---------|
| azurerm | >= 2.20 |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:-----:|
| action\_group\_id | ID of the associated action group | `string` | n/a | yes |
| alert\_criteria | Alert Criteria: dynamic nested blocks | <pre>map(object({<br>    metric_name      = string<br>    metric_namespace = string<br>    operator         = string<br>    aggregation      = string<br>    threshold        = number<br>  }))</pre> | n/a | yes |
| metric\_alert\_name | Name of the metric alert | `string` | n/a | yes |
| resource\_group\_name | Name of the resource group | `string` | n/a | yes |
| resource\_name | Name of the monitored resource | `string` | n/a | yes |
| scopes | IDs of the monitored resources | `list(string)` | n/a | yes |

## Outputs

No output.
<!--- END_TF_DOCS --->
