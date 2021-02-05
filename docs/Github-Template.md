# Azure Container Registry

This repo contains an example Terraform configuration that deploys an Azure Resource.
For more info, please see https://docs.microsoft.com/en-us/azure/.

## Version compatibility

| Module version    | Terraform version | AzureRM version |
|-------------------|-------------------|-----------------|
| >= 1.x.x          | 0.13.x            | >= 2.3.0        |

## Example Usage
<!--- BEGIN_TF_DOCS --->
<!--- END_TF_DOCS --->

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

### Deploy the code

```
cd examples/sandbox
terraform init
terraform plan -out sandbox-01.tfplan
terraform apply sandbox-01.tfplan
```

### Test the code

```
cd tests
go mod init 'tests'
go test -run TestSandboxExample -v -timeout 30m
```

Or Using Make
```
make test
```


