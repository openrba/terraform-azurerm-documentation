# Azure Action Groups

An action group is a collection of notification preferences defined by the owner of an Azure subscription. Azure Monitor and Service Health alerts use action groups to notify users that an alert has been triggered. Various alerts may use the same action group or different action groups depending on the requirements.

### Integration

Module can be used for the following environments:

- Sandbox
- Development
- Production

### Required

- Subscription for one of the above environments
- Resource Group

# Usage

<!--- BEGIN_TF_DOCS --->
## Providers

| Name | Version |
|------|---------|
| azurerm | >= 2.20 |
| random | >= 2.3 |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:-----:|
| email\_receivers | A map of objects that define the email attributes | <pre>map(object({<br>    name                    = string<br>    email_address           = string<br>    use_common_alert_schema = bool<br>  }))</pre> | n/a | yes |
| resource\_group\_name | Name of the resource group | `string` | n/a | yes |
| short\_name | This display name will be shown as the action group name in email and SMS notifications. Limited to 12 characters | `string` | n/a | yes |

## Outputs

| Name | Description |
|------|-------------|
| action\_group\_id | n/a |
<!--- END_TF_DOCS --->
