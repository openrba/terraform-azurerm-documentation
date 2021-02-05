## Azure - Kubernetes Velero backup and restore module

### Introduction

This module will install Velero module into an AKS cluster. This module is to backup and restore aks clusters, set schedules to automatically kickoff backups at recurring intervals.

<!--- BEGIN_TF_DOCS --->
## Requirements

No requirements.

## Providers

| Name | Version |
|------|---------|
| azurerm | n/a |
| helm | n/a |
| kubernetes | n/a |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| backup\_storage\_account | Name of backup storage account in azure | `string` | `""` | no |
| backup\_storage\_container | Name of backup container | `string` | `"velero"` | no |
| helm\_release\_name | helm release name | `string` | `"velero"` | no |
| helm\_repository | velero helm repository url | `string` | `"https://vmware-tanzu.github.io/helm-charts/"` | no |
| resource\_group\_location | Name of backup resource group location | `string` | `""` | no |
| resource\_group\_name | Name of backup resource group | `string` | `""` | no |
| velero\_chart\_version | Velero helm chart version to use | `string` | `"2.13.3"` | no |
| velero\_namespace | Kubernetes namespace in which to deploy Velero | `string` | `"velero"` | no |

## Outputs

No output.

<!--- END_TF_DOCS --->