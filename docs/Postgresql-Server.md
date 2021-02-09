# Azure Database for PostgresSQL Server

This repo contains an example Terraform configuration that deploys a PostgresSQL database using Azure.
For more info, please see https://docs.microsoft.com/en-us/azure/postgresql/.


<!--- BEGIN_TF_DOCS --->
## Providers

| Name | Version |
|------|---------|
| azurerm | n/a |
| random | n/a |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:-----:|
| access\_list | Access list for PostgresSQL instance. Map off names to cidr ip start/end addresses | <pre>map(object({ start_ip_address = string<br>  end_ip_address = string }))</pre> | `{}` | no |
| ad\_admin\_login\_name | The login name of the principal to set as the server administrator. | `string` | `""` | no |
| administrator\_login | Database administrator login name | `string` | `"az_dbadmin"` | no |
| all\_metrics | Retention only applies to storage account. Retention policy ranges from 1 to 365 days. If you do not want to apply any retention policy and retain data forever, set retention (days) to 0. | `number` | `0` | no |
| auto\_grow\_enabled | Enable/Disable auto-growing of the storage. | `bool` | `false` | no |
| backup\_retention\_days | Backup retention days for the server, supported values are between 7 and 35 days. | `number` | `"7"` | no |
| create\_mode | Can be used to restore or replicate existing servers. Possible values are Default, Replica, GeoRestore, and PointInTimeRestore. Defaults to Default | `string` | `"Default"` | no |
| creation\_source\_server\_id | the source server ID to use. use this only when creating a read replica server | `string` | `""` | no |
| database\_defaults | database default charset and collation (for TF managed databases) | <pre>object({<br>    charset   = string<br>    collation = string<br>  })</pre> | <pre>{<br>  "charset": "UTF8",<br>  "collation": "English_United States.1252"<br>}</pre> | no |
| databases | Map of databases to create (keys are database names). Allowed values are the same as for database\_defaults. | `map` | `{}` | no |
| enable\_postgresql\_ad\_admin | Set a user or group as the AD administrator for an postgresql server in Azure | `bool` | `false` | no |
| enable\_threat\_detection\_policy | Threat detection policy configuration, known in the API as Server Security Alerts Policy. | `bool` | `false` | no |
| geo\_redundant\_backup\_enabled | Turn Geo-redundant server backups on/off. This allows you to choose between locally redundant or geo-redundant backup storage in the General Purpose and Memory Optimized tiers. When the backups are stored in geo-redundant backup storage, they are not only stored within the region in which your server is hosted, but are also replicated to a paired data center. This provides better protection and ability to restore your server in a different region in the event of a disaster. This is not supported for the Basic tier. | `string` | `"true"` | no |
| infrastructure\_encryption\_enabled | Whether or not infrastructure is encrypted for this server. Defaults to false. Changing this forces a new resource to be created. | `string` | `"false"` | no |
| location | Azure region | `string` | n/a | yes |
| log\_retention\_days | Specifies the number of days to keep in the Threat Detection audit logs | `string` | `"7"` | no |
| max\_query\_text\_length | Sets the maximum query text length that will be saved; longer queries will be truncated. | `string` | `"6000"` | no |
| names | names to be applied to resources | `map(string)` | n/a | yes |
| postgresql\_config | A map of postgresql additional configuration parameters to values. | `map(string)` | `{}` | no |
| postgresql\_logs | PostgreSQLLogs retention days | `number` | `0` | no |
| postgresql\_version | Specifies the version of PostgreSQL to use. Valid values are 9.5, 9.6, 10, 10.0, and 11. | `string` | `"10.0"` | no |
| query\_capture\_mode | Selects whether parameter placeholders are replaced in parameterized queries. | `string` | `"TOP"` | no |
| query\_store\_runtime\_statistics | Contains information about the query runtime statistics such as CPU usage and query duration statistics. | `number` | `0` | no |
| query\_store\_wait\_statistics | Contains information about the query wait statistics (what your queries waited on) such are CPU, LOG, and LOCKING. | `number` | `0` | no |
| replace\_parameter\_placeholders | Selects whether parameter placeholders are replaced in parameterized queries. | `string` | `"off"` | no |
| resource\_group\_name | name of the resource group to create the resource | `string` | n/a | yes |
| retention\_period\_in\_days | Sets the retention period window in days for pg\_qs - after this time data will be deleted. | `string` | `"7"` | no |
| server\_id | identifier appended to server name for more info see https://github.com/openrba/python-azure-naming#azuredbforpostgresql | `string` | n/a | yes |
| service\_endpoints | Creates a virtual network rule in the subnet\_id (values are virtual network subnet ids). | `map(string)` | `{}` | no |
| sku\_name | Specifies the SKU Name for this PostgreSQL Server. The name of the SKU, follows the tier + family + cores pattern (e.g. B\_Gen4\_1, GP\_Gen5\_8). | `string` | `"GP_Gen5_2"` | no |
| ssl\_enforcement\_enabled | Specifies if SSL should be enforced on connections. Possible values are true and false. | `bool` | `true` | no |
| storage\_mb | Max storage allowed for a server. | `number` | `"10240"` | no |
| tags | tags to be applied to resources | `map(string)` | n/a | yes |
| threat\_detection\_policy | Threat detection policy configuration, known in the API as Server Security Alerts Policy. The threat\_detection\_policy block supports fields documented below. | `string` | `"false"` | no |
| track\_utility | Selects whether utility commands are tracked by pg\_qs. | `string` | `"on"` | no |

## Outputs

| Name | Description |
|------|-------------|
| administrator\_login | The postgresql instance login for the admin. |
| administrator\_password | The password for the admin account of the postgresql instance. |
| fqdn | The fully qualified domain name of the instance Azure SQL Server |
| id | The ID of the postgresql instance. |
| name | The Name of the postgresql instance. |
| resource\_group\_name | The name of the resource group in which resources are created |
<!--- END_TF_DOCS --->
