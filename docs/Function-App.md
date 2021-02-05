# Azure Function Terraform Module

Usage:

Example of 'func' module in `main.tf`.

location = eastus  
environment = sandbox  
names = names

Even inline **formatting** in _here_ is possible.  
and some [link](https://domain.com/)

* list item 3
* list item 4

```hcl
module "func" {
  source = "github.com/openrba/terraform-azurerm-function-app"

  id   = "1234567890"
  name = "baz"

  zones = ["us-east-1", "us-west-1"]

  tags = {
    Name         = "baz"
    Created-By   = "first.last@email.com"
    Date-Created = "20180101"
  }
}
```

Here is some trailing text after code block,  
followed by another line of text.

| Name | Description     |
|------|-----------------|
| Foo  | Foo description |
| Bar  | Bar description |

## Quick start

1.Install [Terraform](https://learn.hashicorp.com/tutorials/terraform/install-cli).\  
2.Sign into your [Azure Account](https://docs.microsoft.com/en-us/cli/azure/authenticate-azure-cli?view=azure-cli-latest)

```
# Login with the Azure CLI/bash terminal/powershell by running
az login

# Verify access by running
az account show --output jsonc

# Confirm you are running required/pinned version of terraform
terraform version
```

Deploy the code:

```
terraform init
terraform plan -out azure-mysql-01.tfplan
terraform apply azure-mysql-01.tfplan
```

## Requirements

No requirements.

## Providers

| Name | Version |
|------|---------|
| archive | n/a |
| azurerm | n/a |
| random | n/a |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| docker | block for docker credentials | <pre>object({<br>    image    = string<br>    server   = string<br>    username = string<br>    password = string<br>  })</pre> | n/a | yes |
| function\_runtime | language runtime for the function app | `string` | n/a | yes |
| location | Location for all resources | `string` | n/a | yes |
| names | names to be applied to resources | `map(string)` | n/a | yes |
| resource\_group\_name | Resource group name | `string` | n/a | yes |
| tags | tags to be applied to resources | `map(string)` | n/a | yes |

## Outputs

| Name | Description |
|------|-------------|
| function\_app\_id | The resource id for the function app. |
| function\_hostname | The default hostname associated with the Function App - such as mysite.azurewebsites.net. |

