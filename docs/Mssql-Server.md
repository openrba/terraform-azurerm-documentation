# Azure Database for MS SQL Server

This repo contains an example Terraform configuration that deploys a MySQL database using Azure.
For more info, please see https://docs.microsoft.com/en-us/azure/azure-sql/database/.

## Version compatibility

| Module version    | Terraform version | AzureRM version |
|-------------------|-------------------|-----------------|
| >= 1.x.x          | 0.13.x            | >= 2.2.0        |



## Example Usage
<!--- BEGIN_TF_DOCS --->
## Requirements

| Name | Version |
|------|---------|
| terraform | >= 0.13.0 |
| azurerm | >= 2.25.0 |
| random | >= 2.2.0 |

## Providers

| Name | Version |
|------|---------|
| azurerm | >= 2.25.0 |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| access\_list | Access list for SQL Server instance. Map off names to cidr ip start/end addresses | <pre>map(object({ start_ip_address = string,<br>  end_ip_address = string }))</pre> | `{}` | no |
| account\_replication\_type | The account replication type of the servers that do not use failover groups. Failover groups should always use geo redundant storage with the secondary in the same region as the replica. | `string` | `"LRS"` | no |
| ad\_admin\_login\_name | The login name of the azure ad admin. | `string` | `""` | no |
| administrator\_login | Database administrator login name | `string` | `"az_dbadmin"` | no |
| audit\_log\_enabled | Whether to enable auditing of logs. | `bool` | `true` | no |
| auto\_grow\_enabled | Enable/Disable auto-growing of the storage. | `bool` | `false` | no |
| automatic\_tuning | Contains information about automatic tuning recommendations for a database. Valid values are for Retention policy ranges from 1 to 365 days. If you don't want to apply any retention policy, set retention days to '0' | `number` | `0` | no |
| backup\_retention\_days | Backup retention days for the server, supported values are between 7 and 35 days. | `number` | `7` | no |
| basic | Contains DTU/CPU percentage, DTU/CPU limit, physical data read percentage, log write percentage, Successful/Failed/Blocked by firewall connections, sessions percentage, workers percentage, storage, storage percentage, and XTP storage percentage. Valid values are for Retention policy ranges from 1 to 365 days. If you don't want to apply any retention policy, set retention days to '0' | `number` | `0` | no |
| blocks | Contains information about blocking events on a database. Valid values are for Retention policy ranges from 1 to 365 days. If you don't want to apply any retention policy, set retention days to '0' | `number` | `0` | no |
| create\_mode | Can be used to restore or replicate existing servers. Possible values are Default and Replica. Defaults to Default | `string` | `"Default"` | no |
| creation\_source\_server\_name | the source server name to use. use this only when creating a read replica server | `string` | `""` | no |
| database\_defaults | database default charset and collation (only applied to databases managed within this module) | <pre>object({<br>    collation          = string<br>    license_type       = string<br>    sku_name           = string<br>    max_size_gb        = number<br>    zone_redundant     = bool<br>    read_scale         = bool<br>    read_replica_count = number<br>  })</pre> | <pre>{<br>  "auto_pause_delay_in_minutes": -1,<br>  "collation": "SQL_Latin1_General_CP1_CI_AS",<br>  "license_type": "LicenseIncluded",<br>  "max_size_gb": 4,<br>  "min_capacity": 0.5,<br>  "read_replica_count": 0,<br>  "read_scale": false,<br>  "sku_name": "GP_S_Gen5_1",<br>  "zone_redundant": false<br>}</pre> | no |
| database\_wait\_statistics | Contains information about how much time a database spent waiting on different wait types. Valid values are for Retention policy ranges from 1 to 365 days. If you don't want to apply any retention policy, set retention days to '0' | `number` | `0` | no |
| databases | Map of databases to create (keys are database names). Allowed values are the same as for database\_defaults. | `map(any)` | `{}` | no |
| deadlocks | Contains information about deadlock events on a database. Valid values are for Retention policy ranges from 1 to 365 days. If you don't want to apply any retention policy, set retention days to '0' | `number` | `0` | no |
| enable\_active\_directory\_administrator | Set a user or group as the AD administrator for an SQL Server server in Azure | `bool` | `false` | no |
| enable\_threat\_detection\_policy | Threat detection policy configuration, known in the API as Server Security Alerts Policy. | `bool` | `false` | no |
| error\_log | Contains information about SQL errors on a database. Valid values are for Retention policy ranges from 1 to 365 days. If you don't want to apply any retention policy, set retention days to '0' | `number` | `0` | no |
| geo\_redundant\_backup\_enabled | Turn Geo-redundant server backups on/off. This allows you to choose between locally redundant or geo-redundant backup storage in the General Purpose and Memory Optimized tiers. When the backups are stored in geo-redundant backup storage, they are not only stored within the region in which your server is hosted, but are also replicated to a paired data center. This provides better protection and ability to restore your server in a different region in the event of a disaster. This is not supported for the Basic tier. | `bool` | `false` | no |
| infrastructure\_encryption\_enabled | Whether or not infrastructure is encrypted for this server. Defaults to false. Changing this forces a new resource to be created. | `bool` | `false` | no |
| instance\_and\_app\_advanced | Contains tempdb system database data and log file size and tempdb percent log file used. Valid values are for Retention policy ranges from 1 to 365 days. If you don't want to apply any retention policy, set retention days to '0' | `number` | `0` | no |
| location | Specifies the supported Azure location to SQL Server server resource | `string` | n/a | yes |
| log\_retention\_days | Specifies the number of days to keep in the Threat Detection audit logs | `number` | `7` | no |
| max\_allowed\_packet | The maximum size of one packet or any generated/intermediate string, or any parameter sent by the sql\_stmt\_send\_long\_data() C API function. | `string` | `"1073741824"` | no |
| max\_connections | The maximum permitted number of simultaneous client connections. value 10-600 | `number` | `600` | no |
| max\_heap\_table\_size | This variable sets the maximum size to which user-created MEMORY tables are permitted to grow. | `string` | `"64000000"` | no |
| metric | Retention only applies to storage account. Valid values are for Retention policy ranges from 1 to 365 days. If you don't want to apply any retention policy, set retention days to '0' | `number` | `0` | no |
| names | names to be applied to resources | `map(string)` | n/a | yes |
| performance\_schema | The value of this variable is ON or OFF to indicate whether the Performance Schema is enabled. | `string` | `"ON"` | no |
| query\_store\_capture\_interval | The query store capture interval in minutes. Allows to specify the interval in which the query metrics are aggregated. | `string` | `"15"` | no |
| query\_store\_capture\_mode | The query store capture mode, NONE means do not capture any statements. NOTE: If performance\_schema is OFF, turning on query\_store\_capture\_mode will turn on performance\_schema and a subset of performance schema instruments required for this feature. | `string` | `"ALL"` | no |
| query\_store\_capture\_utility\_queries | Turning ON or OFF to capture all the utility queries that is executing in the system. | `string` | `"YES"` | no |
| query\_store\_retention\_period\_in\_days | The query store capture interval in minutes. Allows to specify the interval in which the query metrics are aggregated. | `number` | `7` | no |
| query\_store\_runtime\_statistics | Contains information about the query runtime statistics such as CPU usage and query duration statistics. Valid values are for Retention policy ranges from 1 to 365 days. If you don't want to apply any retention policy, set retention days to '0' | `number` | `0` | no |
| query\_store\_wait\_sampling\_capture\_mode | The query store wait event sampling capture mode, NONE means do not capture any wait events. | `string` | `"ALL"` | no |
| query\_store\_wait\_sampling\_frequency | The query store wait event sampling frequency in seconds. | `number` | `30` | no |
| query\_store\_wait\_statistics | Contains information about the query wait statistics (what your queries waited on) such are CPU, LOG, and LOCKING. Valid values are for Retention policy ranges from 1 to 365 days. If you don't want to apply any retention policy, set retention days to '0' | `number` | `0` | no |
| replicate\_wild\_ignore\_table | Creates a replication filter which keeps the slave thread from replicating a statement in which any table matches the given wildcard pattern. To specify more than one table to ignore, use comma-separated list. | `string` | `"sql.%,tempdb.%"` | no |
| resource\_group\_name | name of the resource group to create the resource | `string` | n/a | yes |
| server\_id | Specifies the server identifier to SQL Server server resource | `string` | n/a | yes |
| service\_endpoints | Creates a virtual network rule in the subnet\_id (values are virtual network subnet ids). | `map(string)` | `{}` | no |
| sku\_name | Azure database for SQL Server sku name | `string` | `"GP_Gen5_2"` | no |
| slow\_query\_log | Enable or disable the slow query log | `string` | `"OFF"` | no |
| sort\_buffer\_size | Each session that must perform a sort allocates a buffer of this size. | `number` | `2000000` | no |
| sql\_insights | Contains Intelligent Insights into performance for a database. Valid values are for Retention policy ranges from 1 to 365 days. If you don't want to apply any retention policy, set retention days to '0' | `number` | `0` | no |
| sql\_version | SQL Server version | `string` | `"12.0"` | no |
| ssl\_enforcement\_enabled | Specifies if SSL should be enforced on connections. Possible values are true and false. | `bool` | `true` | no |
| storage\_mb | Max storage allowed for a server | `number` | `"10240"` | no |
| tags | tags to be applied to resources | `map(string)` | n/a | yes |
| timeouts | Contains information about timeouts on a database. Valid values are for Retention policy ranges from 1 to 365 days. If you don't want to apply any retention policy, set retention days to '0' | `number` | `0` | no |
| tmp\_table\_size | The maximum size of internal in-memory temporary tables. This variable does not apply to user-created MEMORY tables. | `number` | `64000000` | no |
| transaction\_isolation | The default transaction isolation level. | `string` | `"READ-COMMITTED"` | no |
| workload\_management | Resource Governor enables you to manage SQL Server workloads and resources by specifying limits on resource consumption by incoming requests. Valid values are for Retention policy ranges from 1 to 365 days. If you don't want to apply any retention policy, set retention days to '0' | `number` | `0` | no |

## Outputs

| Name | Description |
|------|-------------|
| administrator\_login | The sql instance login for the admin. |
| administrator\_password | The password for the admin account of the MySQL instance. |
| database\_ids | The resulting database ids |
| fqdn | The fully qualified domain name of the Azure MySQL Server |
| id | The ID of the MySQL instance. |
| name | The Name of the sql instance. |
| resource\_group\_name | The name of the resource group in which resources are created |

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

Deploy the code:

```
terraform init
terraform plan -out azure-mysql-01.tfplan
terraform apply azure-mysql-01.tfplan
```
