// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package sql

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Creates a new Google SQL Database Instance. For more information, see the [official documentation](https://cloud.google.com/sql/),
// or the [JSON API](https://cloud.google.com/sql/docs/admin-api/v1beta4/instances).
//
// > **NOTE on `sql.DatabaseInstance`:** - First-generation instances have been
// deprecated and should no longer be created, see [upgrade docs](https://cloud.google.com/sql/docs/mysql/upgrade-2nd-gen)
// for more details.
// To upgrade your First-generation instance, update your config that the instance has
// * `settings.ip_configuration.ipv4_enabled=true`
// * `settings.backup_configuration.enabled=true`
// * `settings.backup_configuration.binary_log_enabled=true`.\
//   Apply the config, then upgrade the instance in the console as described in the documentation.
//   Once upgraded, update the following attributes in your config to the correct value according to
//   the above documentation:
// * `region`
// * `databaseVersion` (if applicable)
// * `tier`\
//   Remove any fields that are not applicable to Second-generation instances:
// * `settings.crash_safe_replication`
// * `settings.replication_type`
// * `settings.authorized_gae_applications`
//   And change values to appropriate values for Second-generation instances for:
// * `activationPolicy` ("ON_DEMAND" is no longer an option)
// * `pricingPlan` ("PER_USE" is now the only valid option)
//   Change `settings.backup_configuration.enabled` attribute back to its desired value and apply as necessary.
//
// > **NOTE on `sql.DatabaseInstance`:** - Second-generation instances include a
// default 'root'@'%' user with no password. This user will be deleted by the provider on
// instance creation. You should use `sql.User` to define a custom user with
// a restricted host and strong password.
//
// ## Example Usage
// ### SQL Second Generation Instance
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v3/go/gcp/sql"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := sql.NewDatabaseInstance(ctx, "master", &sql.DatabaseInstanceArgs{
// 			DatabaseVersion: pulumi.String("POSTGRES_11"),
// 			Region:          pulumi.String("us-central1"),
// 			Settings: &sql.DatabaseInstanceSettingsArgs{
// 				Tier: pulumi.String("db-f1-micro"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Private IP Instance
// > **NOTE**: For private IP instance setup, note that the `sql.DatabaseInstance` does not actually interpolate values from `servicenetworking.Connection`. You must explicitly add a `dependsOn`reference as shown below.
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v3/go/gcp/compute"
// 	"github.com/pulumi/pulumi-gcp/sdk/v3/go/gcp/servicenetworking"
// 	"github.com/pulumi/pulumi-gcp/sdk/v3/go/gcp/sql"
// 	"github.com/pulumi/pulumi-random/sdk/v2/go/random"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		privateNetwork, err := compute.NewNetwork(ctx, "privateNetwork", nil, pulumi.Provider(google_beta))
// 		if err != nil {
// 			return err
// 		}
// 		privateIpAddress, err := compute.NewGlobalAddress(ctx, "privateIpAddress", &compute.GlobalAddressArgs{
// 			Purpose:      pulumi.String("VPC_PEERING"),
// 			AddressType:  pulumi.String("INTERNAL"),
// 			PrefixLength: pulumi.Int(16),
// 			Network:      privateNetwork.ID(),
// 		}, pulumi.Provider(google_beta))
// 		if err != nil {
// 			return err
// 		}
// 		privateVpcConnection, err := servicenetworking.NewConnection(ctx, "privateVpcConnection", &servicenetworking.ConnectionArgs{
// 			Network: privateNetwork.ID(),
// 			Service: pulumi.String("servicenetworking.googleapis.com"),
// 			ReservedPeeringRanges: pulumi.StringArray{
// 				privateIpAddress.Name,
// 			},
// 		}, pulumi.Provider(google_beta))
// 		if err != nil {
// 			return err
// 		}
// 		_, err = random.NewRandomId(ctx, "dbNameSuffix", &random.RandomIdArgs{
// 			ByteLength: pulumi.Int(4),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = sql.NewDatabaseInstance(ctx, "instance", &sql.DatabaseInstanceArgs{
// 			Region: pulumi.String("us-central1"),
// 			Settings: &sql.DatabaseInstanceSettingsArgs{
// 				Tier: pulumi.String("db-f1-micro"),
// 				IpConfiguration: &sql.DatabaseInstanceSettingsIpConfigurationArgs{
// 					Ipv4Enabled:    pulumi.Bool(false),
// 					PrivateNetwork: privateNetwork.ID(),
// 				},
// 			},
// 		}, pulumi.Provider(google_beta), pulumi.DependsOn([]pulumi.Resource{
// 			privateVpcConnection,
// 		}))
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type DatabaseInstance struct {
	pulumi.CustomResourceState

	// The connection name of the instance to be used in
	// connection strings. For example, when connecting with [Cloud SQL Proxy](https://cloud.google.com/sql/docs/mysql/connect-admin-proxy).
	ConnectionName pulumi.StringOutput `pulumi:"connectionName"`
	// The MySQL, PostgreSQL or
	// SQL Server (beta) version to use. Supported values include `MYSQL_5_6`,
	// `MYSQL_5_7`, `POSTGRES_9_6`,`POSTGRES_10`, `POSTGRES_11`, `POSTGRES_12`, `SQLSERVER_2017_STANDARD`,
	// `SQLSERVER_2017_ENTERPRISE`, `SQLSERVER_2017_EXPRESS`, `SQLSERVER_2017_WEB`.
	// [Database Version Policies](https://cloud.google.com/sql/docs/sqlserver/db-versions)
	// includes an up-to-date reference of supported versions.
	DatabaseVersion pulumi.StringPtrOutput `pulumi:"databaseVersion"`
	// The full path to the encryption key used for the CMEK disk encryption.  Setting
	// up disk encryption currently requires manual steps outside of this provider.
	// The provided key must be in the same region as the SQL instance.  In order
	// to use this feature, a special kind of service account must be created and
	// granted permission on this key.  This step can currently only be done
	// manually, please see [this step](https://cloud.google.com/sql/docs/mysql/configure-cmek#service-account).
	// That service account needs the `Cloud KMS > Cloud KMS CryptoKey Encrypter/Decrypter` role on your
	// key - please see [this step](https://cloud.google.com/sql/docs/mysql/configure-cmek#grantkey).
	EncryptionKeyName pulumi.StringOutput `pulumi:"encryptionKeyName"`
	// The first IPv4 address of any type assigned.
	FirstIpAddress pulumi.StringOutput                  `pulumi:"firstIpAddress"`
	IpAddresses    DatabaseInstanceIpAddressArrayOutput `pulumi:"ipAddresses"`
	// The name of the instance that will act as
	// the master in the replication setup. Note, this requires the master to have
	// `binaryLogEnabled` set, as well as existing backups.
	MasterInstanceName pulumi.StringOutput `pulumi:"masterInstanceName"`
	// A name for this whitelist entry.
	Name pulumi.StringOutput `pulumi:"name"`
	// The first private (`PRIVATE`) IPv4 address assigned.
	PrivateIpAddress pulumi.StringOutput `pulumi:"privateIpAddress"`
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The first public (`PRIMARY`) IPv4 address assigned.
	PublicIpAddress pulumi.StringOutput `pulumi:"publicIpAddress"`
	// The region the instance will sit in. Note, Cloud SQL is not
	// available in all regions - choose from one of the options listed [here](https://cloud.google.com/sql/docs/mysql/instance-locations).
	// A valid region must be provided to use this resource. If a region is not provided in the resource definition,
	// the provider region will be used instead, but this will be an apply-time error for instances if the provider
	// region is not supported with Cloud SQL. If you choose not to provide the `region` argument for this resource,
	// make sure you understand this.
	Region pulumi.StringOutput `pulumi:"region"`
	// The configuration for replication. The
	// configuration is detailed below.
	ReplicaConfiguration DatabaseInstanceReplicaConfigurationOutput `pulumi:"replicaConfiguration"`
	// Initial root password. Required for MS SQL Server, ignored by MySQL and PostgreSQL.
	RootPassword pulumi.StringPtrOutput `pulumi:"rootPassword"`
	// The URI of the created resource.
	SelfLink     pulumi.StringOutput                `pulumi:"selfLink"`
	ServerCaCert DatabaseInstanceServerCaCertOutput `pulumi:"serverCaCert"`
	// The service account email address assigned to the
	// instance.
	ServiceAccountEmailAddress pulumi.StringOutput `pulumi:"serviceAccountEmailAddress"`
	// The settings to use for the database. The
	// configuration is detailed below.
	Settings DatabaseInstanceSettingsOutput `pulumi:"settings"`
}

// NewDatabaseInstance registers a new resource with the given unique name, arguments, and options.
func NewDatabaseInstance(ctx *pulumi.Context,
	name string, args *DatabaseInstanceArgs, opts ...pulumi.ResourceOption) (*DatabaseInstance, error) {
	if args == nil || args.Settings == nil {
		return nil, errors.New("missing required argument 'Settings'")
	}
	if args == nil {
		args = &DatabaseInstanceArgs{}
	}
	var resource DatabaseInstance
	err := ctx.RegisterResource("gcp:sql/databaseInstance:DatabaseInstance", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDatabaseInstance gets an existing DatabaseInstance resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDatabaseInstance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DatabaseInstanceState, opts ...pulumi.ResourceOption) (*DatabaseInstance, error) {
	var resource DatabaseInstance
	err := ctx.ReadResource("gcp:sql/databaseInstance:DatabaseInstance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DatabaseInstance resources.
type databaseInstanceState struct {
	// The connection name of the instance to be used in
	// connection strings. For example, when connecting with [Cloud SQL Proxy](https://cloud.google.com/sql/docs/mysql/connect-admin-proxy).
	ConnectionName *string `pulumi:"connectionName"`
	// The MySQL, PostgreSQL or
	// SQL Server (beta) version to use. Supported values include `MYSQL_5_6`,
	// `MYSQL_5_7`, `POSTGRES_9_6`,`POSTGRES_10`, `POSTGRES_11`, `POSTGRES_12`, `SQLSERVER_2017_STANDARD`,
	// `SQLSERVER_2017_ENTERPRISE`, `SQLSERVER_2017_EXPRESS`, `SQLSERVER_2017_WEB`.
	// [Database Version Policies](https://cloud.google.com/sql/docs/sqlserver/db-versions)
	// includes an up-to-date reference of supported versions.
	DatabaseVersion *string `pulumi:"databaseVersion"`
	// The full path to the encryption key used for the CMEK disk encryption.  Setting
	// up disk encryption currently requires manual steps outside of this provider.
	// The provided key must be in the same region as the SQL instance.  In order
	// to use this feature, a special kind of service account must be created and
	// granted permission on this key.  This step can currently only be done
	// manually, please see [this step](https://cloud.google.com/sql/docs/mysql/configure-cmek#service-account).
	// That service account needs the `Cloud KMS > Cloud KMS CryptoKey Encrypter/Decrypter` role on your
	// key - please see [this step](https://cloud.google.com/sql/docs/mysql/configure-cmek#grantkey).
	EncryptionKeyName *string `pulumi:"encryptionKeyName"`
	// The first IPv4 address of any type assigned.
	FirstIpAddress *string                     `pulumi:"firstIpAddress"`
	IpAddresses    []DatabaseInstanceIpAddress `pulumi:"ipAddresses"`
	// The name of the instance that will act as
	// the master in the replication setup. Note, this requires the master to have
	// `binaryLogEnabled` set, as well as existing backups.
	MasterInstanceName *string `pulumi:"masterInstanceName"`
	// A name for this whitelist entry.
	Name *string `pulumi:"name"`
	// The first private (`PRIVATE`) IPv4 address assigned.
	PrivateIpAddress *string `pulumi:"privateIpAddress"`
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The first public (`PRIMARY`) IPv4 address assigned.
	PublicIpAddress *string `pulumi:"publicIpAddress"`
	// The region the instance will sit in. Note, Cloud SQL is not
	// available in all regions - choose from one of the options listed [here](https://cloud.google.com/sql/docs/mysql/instance-locations).
	// A valid region must be provided to use this resource. If a region is not provided in the resource definition,
	// the provider region will be used instead, but this will be an apply-time error for instances if the provider
	// region is not supported with Cloud SQL. If you choose not to provide the `region` argument for this resource,
	// make sure you understand this.
	Region *string `pulumi:"region"`
	// The configuration for replication. The
	// configuration is detailed below.
	ReplicaConfiguration *DatabaseInstanceReplicaConfiguration `pulumi:"replicaConfiguration"`
	// Initial root password. Required for MS SQL Server, ignored by MySQL and PostgreSQL.
	RootPassword *string `pulumi:"rootPassword"`
	// The URI of the created resource.
	SelfLink     *string                       `pulumi:"selfLink"`
	ServerCaCert *DatabaseInstanceServerCaCert `pulumi:"serverCaCert"`
	// The service account email address assigned to the
	// instance.
	ServiceAccountEmailAddress *string `pulumi:"serviceAccountEmailAddress"`
	// The settings to use for the database. The
	// configuration is detailed below.
	Settings *DatabaseInstanceSettings `pulumi:"settings"`
}

type DatabaseInstanceState struct {
	// The connection name of the instance to be used in
	// connection strings. For example, when connecting with [Cloud SQL Proxy](https://cloud.google.com/sql/docs/mysql/connect-admin-proxy).
	ConnectionName pulumi.StringPtrInput
	// The MySQL, PostgreSQL or
	// SQL Server (beta) version to use. Supported values include `MYSQL_5_6`,
	// `MYSQL_5_7`, `POSTGRES_9_6`,`POSTGRES_10`, `POSTGRES_11`, `POSTGRES_12`, `SQLSERVER_2017_STANDARD`,
	// `SQLSERVER_2017_ENTERPRISE`, `SQLSERVER_2017_EXPRESS`, `SQLSERVER_2017_WEB`.
	// [Database Version Policies](https://cloud.google.com/sql/docs/sqlserver/db-versions)
	// includes an up-to-date reference of supported versions.
	DatabaseVersion pulumi.StringPtrInput
	// The full path to the encryption key used for the CMEK disk encryption.  Setting
	// up disk encryption currently requires manual steps outside of this provider.
	// The provided key must be in the same region as the SQL instance.  In order
	// to use this feature, a special kind of service account must be created and
	// granted permission on this key.  This step can currently only be done
	// manually, please see [this step](https://cloud.google.com/sql/docs/mysql/configure-cmek#service-account).
	// That service account needs the `Cloud KMS > Cloud KMS CryptoKey Encrypter/Decrypter` role on your
	// key - please see [this step](https://cloud.google.com/sql/docs/mysql/configure-cmek#grantkey).
	EncryptionKeyName pulumi.StringPtrInput
	// The first IPv4 address of any type assigned.
	FirstIpAddress pulumi.StringPtrInput
	IpAddresses    DatabaseInstanceIpAddressArrayInput
	// The name of the instance that will act as
	// the master in the replication setup. Note, this requires the master to have
	// `binaryLogEnabled` set, as well as existing backups.
	MasterInstanceName pulumi.StringPtrInput
	// A name for this whitelist entry.
	Name pulumi.StringPtrInput
	// The first private (`PRIVATE`) IPv4 address assigned.
	PrivateIpAddress pulumi.StringPtrInput
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The first public (`PRIMARY`) IPv4 address assigned.
	PublicIpAddress pulumi.StringPtrInput
	// The region the instance will sit in. Note, Cloud SQL is not
	// available in all regions - choose from one of the options listed [here](https://cloud.google.com/sql/docs/mysql/instance-locations).
	// A valid region must be provided to use this resource. If a region is not provided in the resource definition,
	// the provider region will be used instead, but this will be an apply-time error for instances if the provider
	// region is not supported with Cloud SQL. If you choose not to provide the `region` argument for this resource,
	// make sure you understand this.
	Region pulumi.StringPtrInput
	// The configuration for replication. The
	// configuration is detailed below.
	ReplicaConfiguration DatabaseInstanceReplicaConfigurationPtrInput
	// Initial root password. Required for MS SQL Server, ignored by MySQL and PostgreSQL.
	RootPassword pulumi.StringPtrInput
	// The URI of the created resource.
	SelfLink     pulumi.StringPtrInput
	ServerCaCert DatabaseInstanceServerCaCertPtrInput
	// The service account email address assigned to the
	// instance.
	ServiceAccountEmailAddress pulumi.StringPtrInput
	// The settings to use for the database. The
	// configuration is detailed below.
	Settings DatabaseInstanceSettingsPtrInput
}

func (DatabaseInstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseInstanceState)(nil)).Elem()
}

type databaseInstanceArgs struct {
	// The MySQL, PostgreSQL or
	// SQL Server (beta) version to use. Supported values include `MYSQL_5_6`,
	// `MYSQL_5_7`, `POSTGRES_9_6`,`POSTGRES_10`, `POSTGRES_11`, `POSTGRES_12`, `SQLSERVER_2017_STANDARD`,
	// `SQLSERVER_2017_ENTERPRISE`, `SQLSERVER_2017_EXPRESS`, `SQLSERVER_2017_WEB`.
	// [Database Version Policies](https://cloud.google.com/sql/docs/sqlserver/db-versions)
	// includes an up-to-date reference of supported versions.
	DatabaseVersion *string `pulumi:"databaseVersion"`
	// The full path to the encryption key used for the CMEK disk encryption.  Setting
	// up disk encryption currently requires manual steps outside of this provider.
	// The provided key must be in the same region as the SQL instance.  In order
	// to use this feature, a special kind of service account must be created and
	// granted permission on this key.  This step can currently only be done
	// manually, please see [this step](https://cloud.google.com/sql/docs/mysql/configure-cmek#service-account).
	// That service account needs the `Cloud KMS > Cloud KMS CryptoKey Encrypter/Decrypter` role on your
	// key - please see [this step](https://cloud.google.com/sql/docs/mysql/configure-cmek#grantkey).
	EncryptionKeyName *string `pulumi:"encryptionKeyName"`
	// The name of the instance that will act as
	// the master in the replication setup. Note, this requires the master to have
	// `binaryLogEnabled` set, as well as existing backups.
	MasterInstanceName *string `pulumi:"masterInstanceName"`
	// A name for this whitelist entry.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The region the instance will sit in. Note, Cloud SQL is not
	// available in all regions - choose from one of the options listed [here](https://cloud.google.com/sql/docs/mysql/instance-locations).
	// A valid region must be provided to use this resource. If a region is not provided in the resource definition,
	// the provider region will be used instead, but this will be an apply-time error for instances if the provider
	// region is not supported with Cloud SQL. If you choose not to provide the `region` argument for this resource,
	// make sure you understand this.
	Region *string `pulumi:"region"`
	// The configuration for replication. The
	// configuration is detailed below.
	ReplicaConfiguration *DatabaseInstanceReplicaConfiguration `pulumi:"replicaConfiguration"`
	// Initial root password. Required for MS SQL Server, ignored by MySQL and PostgreSQL.
	RootPassword *string `pulumi:"rootPassword"`
	// The settings to use for the database. The
	// configuration is detailed below.
	Settings DatabaseInstanceSettings `pulumi:"settings"`
}

// The set of arguments for constructing a DatabaseInstance resource.
type DatabaseInstanceArgs struct {
	// The MySQL, PostgreSQL or
	// SQL Server (beta) version to use. Supported values include `MYSQL_5_6`,
	// `MYSQL_5_7`, `POSTGRES_9_6`,`POSTGRES_10`, `POSTGRES_11`, `POSTGRES_12`, `SQLSERVER_2017_STANDARD`,
	// `SQLSERVER_2017_ENTERPRISE`, `SQLSERVER_2017_EXPRESS`, `SQLSERVER_2017_WEB`.
	// [Database Version Policies](https://cloud.google.com/sql/docs/sqlserver/db-versions)
	// includes an up-to-date reference of supported versions.
	DatabaseVersion pulumi.StringPtrInput
	// The full path to the encryption key used for the CMEK disk encryption.  Setting
	// up disk encryption currently requires manual steps outside of this provider.
	// The provided key must be in the same region as the SQL instance.  In order
	// to use this feature, a special kind of service account must be created and
	// granted permission on this key.  This step can currently only be done
	// manually, please see [this step](https://cloud.google.com/sql/docs/mysql/configure-cmek#service-account).
	// That service account needs the `Cloud KMS > Cloud KMS CryptoKey Encrypter/Decrypter` role on your
	// key - please see [this step](https://cloud.google.com/sql/docs/mysql/configure-cmek#grantkey).
	EncryptionKeyName pulumi.StringPtrInput
	// The name of the instance that will act as
	// the master in the replication setup. Note, this requires the master to have
	// `binaryLogEnabled` set, as well as existing backups.
	MasterInstanceName pulumi.StringPtrInput
	// A name for this whitelist entry.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The region the instance will sit in. Note, Cloud SQL is not
	// available in all regions - choose from one of the options listed [here](https://cloud.google.com/sql/docs/mysql/instance-locations).
	// A valid region must be provided to use this resource. If a region is not provided in the resource definition,
	// the provider region will be used instead, but this will be an apply-time error for instances if the provider
	// region is not supported with Cloud SQL. If you choose not to provide the `region` argument for this resource,
	// make sure you understand this.
	Region pulumi.StringPtrInput
	// The configuration for replication. The
	// configuration is detailed below.
	ReplicaConfiguration DatabaseInstanceReplicaConfigurationPtrInput
	// Initial root password. Required for MS SQL Server, ignored by MySQL and PostgreSQL.
	RootPassword pulumi.StringPtrInput
	// The settings to use for the database. The
	// configuration is detailed below.
	Settings DatabaseInstanceSettingsInput
}

func (DatabaseInstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseInstanceArgs)(nil)).Elem()
}
