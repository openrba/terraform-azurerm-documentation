# terraform-azurerm-redis

This Terraform module deploys a new redis cache in Azure.

## Example Usage

```hcl
provider "azurerm" {
  features {}
}
terraform{
backend "azurerm" {
    resource_group_name   = "tstate"
    storage_account_name  = "tstate23743"
    container_name        = "tstate"
    key                   = "terraform.tfstate"
  }
}
# Azure Cache for Redis
module "redis" {
  source              = "github.com/openrba/terraform-azurerm-redis.git?ref=v1.1.0"
  location            = "eastus"
  redis_name          = "redis-postgresql"
  resource_group_name = "example-resource-group-name"
  sku_name            = "Standard"
  app_name            = "postgresql"
}

```

## Argument Reference

The following arguments are supported.

- **location** - (Required) The location of where the redis cache is being deployed.
- **redis_name** - (Required) The redis name for the redis cache, example: redis-infrastructure-sandbox-eastus-${redis_name}.
- **sku_name** - (Required) The SKU of Redis to use. Possible values are Basic, Standard and Premium.
- **app_name** - (Required) The app name for the resource group, example: rg-infrastructure-sandbox-eastus-${app_name}.

## Remote State Reference
 Prior to running terraform init run:
 
 export ARM_ACCESS_KEY=${key}
 
 The key is in the storage account: tstate23743 
 
https://docs.microsoft.com/en-us/azure/developer/terraform/store-state-in-azure-storage
