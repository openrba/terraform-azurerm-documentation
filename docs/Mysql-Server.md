# Azure Database for MySQL Server

This repo contains an example Terraform configuration that deploys a MySQL database using Azure.
For more info, please see https://docs.microsoft.com/en-us/azure/mysql/.


<!--- BEGIN_TF_DOCS --->
## Providers

| Name | Version |
|------|---------|
| azurerm | >= 2.25.0 |
| random | >= 2.2.0 |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:-----:|
| access\_list | Access list for MySQL instance. Map off names to cidr ip start/end addresses | <pre>map(object({ start_ip_address = string,<br>  end_ip_address = string }))</pre> | `{}` | no |
| ad\_admin\_login\_name | The login name of the azure ad admin. | `string` | `""` | no |
| administrator\_login | Database administrator login name | `string` | `"az_dbadmin"` | no |
| administrator\_password | Database administrator login name (leave blank to generate random string) | `string` | `""` | no |
| audit\_log\_enabled | The value of this variable is ON or OFF to Allow to audit the log. | `string` | `"ON"` | no |
| auto\_grow\_enabled | Enable/Disable auto-growing of the storage. | `bool` | `false` | no |
| backup\_retention\_days | Backup retention days for the server, supported values are between 7 and 35 days. | `number` | `7` | no |
| character\_set\_server | Use charset\_name as the default server character set. | `string` | `"UTF8MB4"` | no |
| create\_mode | Can be used to restore or replicate existing servers. Possible values are Default, Replica, GeoRestore, and PointInTimeRestore. Defaults to Default | `string` | `"Default"` | no |
| creation\_source\_server\_id | the source server ID to use. use this only when creating a read replica server | `string` | `""` | no |
| database\_defaults | database default charset and collation (only applied to databases managed within this module) | <pre>object({<br>    charset   = string<br>    collation = string<br>  })</pre> | <pre>{<br>  "charset": "utf8",<br>  "collation": "utf8_unicode_ci"<br>}</pre> | no |
| databases | Map of databases to create (keys are database names). Allowed values are the same as for database\_defaults. | `map` | `{}` | no |
| ds\_metrics\_retention\_days | Retention only applies to storage account. Retention policy ranges from 1 to 365 days. If you do not want to apply any retention policy and retain data forever, set retention (days) to 0. | `number` | `0` | no |
| enable\_mysql\_ad\_admin | Set a user or group as the AD administrator for an MySQL server in Azure | `bool` | `false` | no |
| errors | failed connections. | `number` | `0` | no |
| event\_scheduler | Indicates the status of the Event Scheduler. It is always ON for a replica server. | `string` | `"OFF"` | no |
| geo\_redundant\_backup\_enabled | Turn Geo-redundant server backups on/off. This allows you to choose between locally redundant or geo-redundant backup storage in the General Purpose and Memory Optimized tiers. When the backups are stored in geo-redundant backup storage, they are not only stored within the region in which your server is hosted, but are also replicated to a paired data center. This provides better protection and ability to restore your server in a different region in the event of a disaster. This is not supported for the Basic tier. | `bool` | `false` | no |
| infrastructure\_encryption\_enabled | Whether or not infrastructure is encrypted for this server. Defaults to false. Changing this forces a new resource to be created. | `bool` | `false` | no |
| innodb\_autoinc\_lock\_mode | The lock mode to use for generating auto-increment values. | `string` | `"2"` | no |
| innodb\_file\_per\_table | InnoDB stores the data and indexes for each newly created table in a separate .ibd file instead of the system tablespace. It cannot be updated any more for a master/replica server to keep the replication consistency. | `string` | `"ON"` | no |
| join\_buffer\_size | The minimum size of the buffer that is used for plain index scans, range index scans, and joins that do not use indexes and thus perform full table scans. | `string` | `"8000000"` | no |
| latency | max lag accross replicas - lag in bytes of the most lagging replica. replica lag - replica lag in seconds | `number` | `0` | no |
| local\_infile | This variable controls server-side LOCAL capability for LOAD DATA statements. | `string` | `"ON"` | no |
| location | Specifies the supported Azure location to MySQL server resource | `string` | n/a | yes |
| max\_allowed\_packet | The maximum size of one packet or any generated/intermediate string, or any parameter sent by the mysql\_stmt\_send\_long\_data() C API function. | `string` | `"1073741824"` | no |
| max\_connections | The maximum permitted number of simultaneous client connections. value 10-600 | `string` | `"600"` | no |
| max\_heap\_table\_size | This variable sets the maximum size to which user-created MEMORY tables are permitted to grow. | `string` | `"64000000"` | no |
| mysql\_server\_parameters | A map of mysql configuration server parameters to values. https://docs.microsoft.com/en-us/azure/mysql/howto-server-parameters | `map(string)` | `{}` | no |
| mysql\_version | MySQL version | `string` | `"8.0"` | no |
| mysqlauditlogs | produce a log file containing an audit record of server activity. The log contents include when clients connect and disconnect, and what actions they perform while connected, such as which databases and tables they access.. | `number` | `0` | no |
| mysqlslowlogs | The slow query log is a record of SQL queries that took a long time to perform. | `number` | `0` | no |
| names | names to be applied to resources | `map(string)` | n/a | yes |
| performance\_schema | The value of this variable is ON or OFF to indicate whether the Performance Schema is enabled. | `string` | `"ON"` | no |
| query\_store\_capture\_interval | The query store capture interval in minutes. Allows to specify the interval in which the query metrics are aggregated. | `string` | `"15"` | no |
| query\_store\_capture\_mode | The query store capture mode, NONE means do not capture any statements. NOTE: If performance\_schema is OFF, turning on query\_store\_capture\_mode will turn on performance\_schema and a subset of performance schema instruments required for this feature. | `string` | `"ALL"` | no |
| query\_store\_capture\_utility\_queries | Turning ON or OFF to capture all the utility queries that is executing in the system. | `string` | `"YES"` | no |
| query\_store\_retention\_period\_in\_days | The query store capture interval in minutes. Allows to specify the interval in which the query metrics are aggregated. | `string` | `"7"` | no |
| query\_store\_wait\_sampling\_capture\_mode | The query store wait event sampling capture mode, NONE means do not capture any wait events. | `string` | `"ALL"` | no |
| query\_store\_wait\_sampling\_frequency | The query store wait event sampling frequency in seconds. | `string` | `"30"` | no |
| replicate\_wild\_ignore\_table | Creates a replication filter which keeps the slave thread from replicating a statement in which any table matches the given wildcard pattern. To specify more than one table to ignore, use comma-separated list. | `string` | `"mysql.%,tempdb.%"` | no |
| resource\_group\_name | name of the resource group to create the resource | `string` | n/a | yes |
| saturation | backup storage used. cpu percent. io percent. memory percent. server log storage percent. server log storage used. storage limit. storage percent. storage used. | `number` | `0` | no |
| server\_id | identifier appended to server name for more info see https://github.com/openrba/python-azure-naming#azuredbformysql | `string` | n/a | yes |
| service\_endpoints | Creates a virtual network rule in the subnet\_id (values are virtual network subnet ids). | `map(string)` | `{}` | no |
| skip\_show\_database | This prevents people from using the SHOW DATABASES statement if they do not have the SHOW DATABASES privilege. | `string` | `"OFF"` | no |
| sku\_name | Azure database for MySQL sku name | `string` | `"GP_Gen5_2"` | no |
| slow\_query\_log | Enable or disable the slow query log | `string` | `"OFF"` | no |
| sort\_buffer\_size | Each session that must perform a sort allocates a buffer of this size. | `string` | `"2000000"` | no |
| ssl\_enforcement\_enabled | Specifies if SSL should be enforced on connections. Possible values are true and false. | `bool` | `true` | no |
| storage\_account\_access\_key | Specifies the identifier key of the Threat Detection audit storage account. Required if state is Enabled. | `string` | `""` | no |
| storage\_account\_resource\_group | Azure resource group where the storage account resides. | `string` | n/a | yes |
| storage\_endpoint | This blob storage will hold all diagnostic setting audit logs. Required if state is Enabled. | `string` | `""` | no |
| storage\_mb | Max storage allowed for a server | `number` | `"10240"` | no |
| tags | tags to be applied to resources | `map(string)` | n/a | yes |
| threat\_detection\_policy | Threat detection policy configuration.  If not input, threat detection will be disabled. | <pre>object({<br>    enable_threat_detection_policy   = bool<br>    threat_detection_email_addresses = list(string)<br>  })</pre> | n/a | yes |
| tmp\_table\_size | The maximum size of internal in-memory temporary tables. This variable does not apply to user-created MEMORY tables. | `string` | `"64000000"` | no |
| traffic | active connections, network data in across active connections, network data out accross active connections. | `number` | `0` | no |
| transaction\_isolation | The default transaction isolation level. | `string` | `"READ-COMMITTED"` | no |

## Outputs

| Name | Description |
|------|-------------|
| administrator\_login | The mysql instance login for the admin. |
| administrator\_password | The password for the admin account of the MySQL instance. |
| fqdn | The fully qualified domain name of the Azure MySQL Server |
| id | The ID of the MySQL instance. |
| name | The Name of the mysql instance. |
| resource\_group\_name | The name of the resource group in which resources are created |
<!--- END_TF_DOCS --->
