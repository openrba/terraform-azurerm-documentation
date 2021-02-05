# Azure - VM Module

## Introduction

This module will deploy a single generic VM. To deploy multiple VMs import the module and use count to set the "index_value" variable.

Where applicable, the inputs will be checked against the RBA naming rules for Azure.

- https://github.com/openrba/python-azure-naming/<br />

<!--- BEGIN_TF_DOCS --->
## Providers

| Name | Version |
|------|---------|
| azurerm | >= 2.25.0 |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:-----:|
| naming_rules | Names to be applied to resources | map(string) | n/a | yes |
| location | Location where azure resources will be deployed | string | n/a | yes |
| resource_group | Resource group for azure resources | string | n/a | yes |
| subscription_id | The id of the subscription to deploy azure resources in | string | n/a | yes |
| names | names to be applied to resources | map(string) | n/a | yes |
| tags | tags to be applied to resources | map(string) | n/a | yes |
| allow_internet_out | Adds a Security Rule to allow traffic out to the internet | bool | false | no |
| index_value | index value to be passed into naming convention | number | 1 | no |
| application_name | Name of application according to python azure naming | string | n/a | yes |
| address_space | CIDRs for virtual network | list(string) | n/a | yes |
| enable_AADLogin | Enables VM addon to allow vm access via Azure Active Directory | bool | true | no |
| machine_size | Type of VM to deploy | string | "Standard_D2s_v3" | no |
| username | Username to login to VM with | string | "lnadmin" | no |
| public_key_path | The location of the users local public key used to ssh with the VM | string | "~/.ssh/id_rsa.pub" | no |
| source_image | Map used to define what source image to use when deploying the VM | map(string) | <pre>{<br>                  publisher = "Canonical"<br>                  offer = "UbuntuServer"<br>                  sku = "16.04-LTS"<br>                  version = "latest"<br>}</pre> | no |
| os_disk | Map used to define what type of disk will be attached to the VM | map(string) | <pre>{<br>                  caching = "ReadWrite"<br>                  storage_account_type = "Standard_LRS"<br>}</pre> | no |

## Outputs

| Name | Description |
|------|-------------|
| vm_name | Name of the created VM |
| vm_public_ip | The public IP assicated with the creaded VM |
| vm_ssh_cmd | The SSH command to use when connecting to the VM via SSH |
| resource_group | The resource group the VM was created in |
<!--- END_TF_DOCS --->
