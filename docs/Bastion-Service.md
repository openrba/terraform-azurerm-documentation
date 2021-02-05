
# terraform-azurerm-bastion-service
Azure bastion as a service

### please note this will not work with [virtual network module](https://github.com/Azure-Terraform/terraform-azurerm-virtual-network) because it will not create the necesary nsg rules, further troubleshooting is required.

<!--- BEGIN_TF_DOCS --->
## Providers

| Name | Version |
|------|---------|
| azurerm | n/a |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:-----:|
| location | n/a | `string` | `""` | no |
| names | n/a | `map(string)` | `{}` | no |
| resource\_group\_name | n/a | `string` | `""` | no |
| security\_group\_name | n/a | `string` | `""` | no |
| subnet\_id | n/a | `string` | `""` | no |
| tags | n/a | `map(string)` | `{}` | no |

## Outputs

| Name | Description |
|------|-------------|
| nic\_id | n/a |
<!--- END_TF_DOCS --->
