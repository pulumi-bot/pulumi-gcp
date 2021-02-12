// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package sql

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// ## Import
//
// Database instances can be imported using one of any of these accepted formats
//
// ```sh
//  $ pulumi import gcp:sql/databaseInstance:DatabaseInstance master projects/{{project}}/instances/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:sql/databaseInstance:DatabaseInstance master {{project}}/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:sql/databaseInstance:DatabaseInstance master {{name}}
// ```
//
//  config and set on the server. When importing, double-check that your config has all the fields set that you expect- just seeing no diff isn't sufficient to know that your config could reproduce the imported resource.
type DatabaseInstance struct {
	pulumi.CustomResourceState

	// Configuration for creating a new instance as a clone of another instance.
	Clone DatabaseInstanceClonePtrOutput `pulumi:"clone"`
	// The connection name of the instance to be used in
	// connection strings. For example, when connecting with [Cloud SQL Proxy](https://cloud.google.com/sql/docs/mysql/connect-admin-proxy).
	ConnectionName pulumi.StringOutput `pulumi:"connectionName"`
	// The MySQL, PostgreSQL or
	// SQL Server (beta) version to use. Supported values include `MYSQL_5_6`,
	// `MYSQL_5_7`, `MYSQL_8_0`, `POSTGRES_9_6`,`POSTGRES_10`, `POSTGRES_11`,
	// `POSTGRES_12`, `POSTGRES_13`, `SQLSERVER_2017_STANDARD`,
	// `SQLSERVER_2017_ENTERPRISE`, `SQLSERVER_2017_EXPRESS`, `SQLSERVER_2017_WEB`.
	// [Database Version Policies](https://cloud.google.com/sql/docs/db-versions)
	// includes an up-to-date reference of supported versions.
	DatabaseVersion pulumi.StringPtrOutput `pulumi:"databaseVersion"`
	// Used to block Terraform from deleting a SQL Instance.
	DeletionProtection pulumi.BoolPtrOutput `pulumi:"deletionProtection"`
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
	// The full project ID of the source instance.`
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
	ReplicaConfiguration DatabaseInstanceReplicaConfigurationOutput    `pulumi:"replicaConfiguration"`
	RestoreBackupContext DatabaseInstanceRestoreBackupContextPtrOutput `pulumi:"restoreBackupContext"`
	// Initial root password. Required for MS SQL Server, ignored by MySQL and PostgreSQL.
	RootPassword pulumi.StringPtrOutput `pulumi:"rootPassword"`
	// The URI of the created resource.
	SelfLink      pulumi.StringOutput                     `pulumi:"selfLink"`
	ServerCaCerts DatabaseInstanceServerCaCertArrayOutput `pulumi:"serverCaCerts"`
	// The service account email address assigned to the
	// instance.
	ServiceAccountEmailAddress pulumi.StringOutput `pulumi:"serviceAccountEmailAddress"`
	// The settings to use for the database. The
	// configuration is detailed below. Required if `clone` is not set.
	Settings DatabaseInstanceSettingsOutput `pulumi:"settings"`
}

// NewDatabaseInstance registers a new resource with the given unique name, arguments, and options.
func NewDatabaseInstance(ctx *pulumi.Context,
	name string, args *DatabaseInstanceArgs, opts ...pulumi.ResourceOption) (*DatabaseInstance, error) {
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
	// Configuration for creating a new instance as a clone of another instance.
	Clone *DatabaseInstanceClone `pulumi:"clone"`
	// The connection name of the instance to be used in
	// connection strings. For example, when connecting with [Cloud SQL Proxy](https://cloud.google.com/sql/docs/mysql/connect-admin-proxy).
	ConnectionName *string `pulumi:"connectionName"`
	// The MySQL, PostgreSQL or
	// SQL Server (beta) version to use. Supported values include `MYSQL_5_6`,
	// `MYSQL_5_7`, `MYSQL_8_0`, `POSTGRES_9_6`,`POSTGRES_10`, `POSTGRES_11`,
	// `POSTGRES_12`, `POSTGRES_13`, `SQLSERVER_2017_STANDARD`,
	// `SQLSERVER_2017_ENTERPRISE`, `SQLSERVER_2017_EXPRESS`, `SQLSERVER_2017_WEB`.
	// [Database Version Policies](https://cloud.google.com/sql/docs/db-versions)
	// includes an up-to-date reference of supported versions.
	DatabaseVersion *string `pulumi:"databaseVersion"`
	// Used to block Terraform from deleting a SQL Instance.
	DeletionProtection *bool `pulumi:"deletionProtection"`
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
	// The full project ID of the source instance.`
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
	RestoreBackupContext *DatabaseInstanceRestoreBackupContext `pulumi:"restoreBackupContext"`
	// Initial root password. Required for MS SQL Server, ignored by MySQL and PostgreSQL.
	RootPassword *string `pulumi:"rootPassword"`
	// The URI of the created resource.
	SelfLink      *string                        `pulumi:"selfLink"`
	ServerCaCerts []DatabaseInstanceServerCaCert `pulumi:"serverCaCerts"`
	// The service account email address assigned to the
	// instance.
	ServiceAccountEmailAddress *string `pulumi:"serviceAccountEmailAddress"`
	// The settings to use for the database. The
	// configuration is detailed below. Required if `clone` is not set.
	Settings *DatabaseInstanceSettings `pulumi:"settings"`
}

type DatabaseInstanceState struct {
	// Configuration for creating a new instance as a clone of another instance.
	Clone DatabaseInstanceClonePtrInput
	// The connection name of the instance to be used in
	// connection strings. For example, when connecting with [Cloud SQL Proxy](https://cloud.google.com/sql/docs/mysql/connect-admin-proxy).
	ConnectionName pulumi.StringPtrInput
	// The MySQL, PostgreSQL or
	// SQL Server (beta) version to use. Supported values include `MYSQL_5_6`,
	// `MYSQL_5_7`, `MYSQL_8_0`, `POSTGRES_9_6`,`POSTGRES_10`, `POSTGRES_11`,
	// `POSTGRES_12`, `POSTGRES_13`, `SQLSERVER_2017_STANDARD`,
	// `SQLSERVER_2017_ENTERPRISE`, `SQLSERVER_2017_EXPRESS`, `SQLSERVER_2017_WEB`.
	// [Database Version Policies](https://cloud.google.com/sql/docs/db-versions)
	// includes an up-to-date reference of supported versions.
	DatabaseVersion pulumi.StringPtrInput
	// Used to block Terraform from deleting a SQL Instance.
	DeletionProtection pulumi.BoolPtrInput
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
	// The full project ID of the source instance.`
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
	RestoreBackupContext DatabaseInstanceRestoreBackupContextPtrInput
	// Initial root password. Required for MS SQL Server, ignored by MySQL and PostgreSQL.
	RootPassword pulumi.StringPtrInput
	// The URI of the created resource.
	SelfLink      pulumi.StringPtrInput
	ServerCaCerts DatabaseInstanceServerCaCertArrayInput
	// The service account email address assigned to the
	// instance.
	ServiceAccountEmailAddress pulumi.StringPtrInput
	// The settings to use for the database. The
	// configuration is detailed below. Required if `clone` is not set.
	Settings DatabaseInstanceSettingsPtrInput
}

func (DatabaseInstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseInstanceState)(nil)).Elem()
}

type databaseInstanceArgs struct {
	// Configuration for creating a new instance as a clone of another instance.
	Clone *DatabaseInstanceClone `pulumi:"clone"`
	// The MySQL, PostgreSQL or
	// SQL Server (beta) version to use. Supported values include `MYSQL_5_6`,
	// `MYSQL_5_7`, `MYSQL_8_0`, `POSTGRES_9_6`,`POSTGRES_10`, `POSTGRES_11`,
	// `POSTGRES_12`, `POSTGRES_13`, `SQLSERVER_2017_STANDARD`,
	// `SQLSERVER_2017_ENTERPRISE`, `SQLSERVER_2017_EXPRESS`, `SQLSERVER_2017_WEB`.
	// [Database Version Policies](https://cloud.google.com/sql/docs/db-versions)
	// includes an up-to-date reference of supported versions.
	DatabaseVersion *string `pulumi:"databaseVersion"`
	// Used to block Terraform from deleting a SQL Instance.
	DeletionProtection *bool `pulumi:"deletionProtection"`
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
	// The full project ID of the source instance.`
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
	RestoreBackupContext *DatabaseInstanceRestoreBackupContext `pulumi:"restoreBackupContext"`
	// Initial root password. Required for MS SQL Server, ignored by MySQL and PostgreSQL.
	RootPassword *string `pulumi:"rootPassword"`
	// The settings to use for the database. The
	// configuration is detailed below. Required if `clone` is not set.
	Settings *DatabaseInstanceSettings `pulumi:"settings"`
}

// The set of arguments for constructing a DatabaseInstance resource.
type DatabaseInstanceArgs struct {
	// Configuration for creating a new instance as a clone of another instance.
	Clone DatabaseInstanceClonePtrInput
	// The MySQL, PostgreSQL or
	// SQL Server (beta) version to use. Supported values include `MYSQL_5_6`,
	// `MYSQL_5_7`, `MYSQL_8_0`, `POSTGRES_9_6`,`POSTGRES_10`, `POSTGRES_11`,
	// `POSTGRES_12`, `POSTGRES_13`, `SQLSERVER_2017_STANDARD`,
	// `SQLSERVER_2017_ENTERPRISE`, `SQLSERVER_2017_EXPRESS`, `SQLSERVER_2017_WEB`.
	// [Database Version Policies](https://cloud.google.com/sql/docs/db-versions)
	// includes an up-to-date reference of supported versions.
	DatabaseVersion pulumi.StringPtrInput
	// Used to block Terraform from deleting a SQL Instance.
	DeletionProtection pulumi.BoolPtrInput
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
	// The full project ID of the source instance.`
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
	RestoreBackupContext DatabaseInstanceRestoreBackupContextPtrInput
	// Initial root password. Required for MS SQL Server, ignored by MySQL and PostgreSQL.
	RootPassword pulumi.StringPtrInput
	// The settings to use for the database. The
	// configuration is detailed below. Required if `clone` is not set.
	Settings DatabaseInstanceSettingsPtrInput
}

func (DatabaseInstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseInstanceArgs)(nil)).Elem()
}

type DatabaseInstanceInput interface {
	pulumi.Input

	ToDatabaseInstanceOutput() DatabaseInstanceOutput
	ToDatabaseInstanceOutputWithContext(ctx context.Context) DatabaseInstanceOutput
}

func (*DatabaseInstance) ElementType() reflect.Type {
	return reflect.TypeOf((*DatabaseInstance)(nil))
}

func (i *DatabaseInstance) ToDatabaseInstanceOutput() DatabaseInstanceOutput {
	return i.ToDatabaseInstanceOutputWithContext(context.Background())
}

func (i *DatabaseInstance) ToDatabaseInstanceOutputWithContext(ctx context.Context) DatabaseInstanceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseInstanceOutput)
}

func (i *DatabaseInstance) ToDatabaseInstancePtrOutput() DatabaseInstancePtrOutput {
	return i.ToDatabaseInstancePtrOutputWithContext(context.Background())
}

func (i *DatabaseInstance) ToDatabaseInstancePtrOutputWithContext(ctx context.Context) DatabaseInstancePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseInstancePtrOutput)
}

type DatabaseInstancePtrInput interface {
	pulumi.Input

	ToDatabaseInstancePtrOutput() DatabaseInstancePtrOutput
	ToDatabaseInstancePtrOutputWithContext(ctx context.Context) DatabaseInstancePtrOutput
}

type databaseInstancePtrType DatabaseInstanceArgs

func (*databaseInstancePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DatabaseInstance)(nil))
}

func (i *databaseInstancePtrType) ToDatabaseInstancePtrOutput() DatabaseInstancePtrOutput {
	return i.ToDatabaseInstancePtrOutputWithContext(context.Background())
}

func (i *databaseInstancePtrType) ToDatabaseInstancePtrOutputWithContext(ctx context.Context) DatabaseInstancePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseInstancePtrOutput)
}

// DatabaseInstanceArrayInput is an input type that accepts DatabaseInstanceArray and DatabaseInstanceArrayOutput values.
// You can construct a concrete instance of `DatabaseInstanceArrayInput` via:
//
//          DatabaseInstanceArray{ DatabaseInstanceArgs{...} }
type DatabaseInstanceArrayInput interface {
	pulumi.Input

	ToDatabaseInstanceArrayOutput() DatabaseInstanceArrayOutput
	ToDatabaseInstanceArrayOutputWithContext(context.Context) DatabaseInstanceArrayOutput
}

type DatabaseInstanceArray []DatabaseInstanceInput

func (DatabaseInstanceArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*DatabaseInstance)(nil))
}

func (i DatabaseInstanceArray) ToDatabaseInstanceArrayOutput() DatabaseInstanceArrayOutput {
	return i.ToDatabaseInstanceArrayOutputWithContext(context.Background())
}

func (i DatabaseInstanceArray) ToDatabaseInstanceArrayOutputWithContext(ctx context.Context) DatabaseInstanceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseInstanceArrayOutput)
}

// DatabaseInstanceMapInput is an input type that accepts DatabaseInstanceMap and DatabaseInstanceMapOutput values.
// You can construct a concrete instance of `DatabaseInstanceMapInput` via:
//
//          DatabaseInstanceMap{ "key": DatabaseInstanceArgs{...} }
type DatabaseInstanceMapInput interface {
	pulumi.Input

	ToDatabaseInstanceMapOutput() DatabaseInstanceMapOutput
	ToDatabaseInstanceMapOutputWithContext(context.Context) DatabaseInstanceMapOutput
}

type DatabaseInstanceMap map[string]DatabaseInstanceInput

func (DatabaseInstanceMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*DatabaseInstance)(nil))
}

func (i DatabaseInstanceMap) ToDatabaseInstanceMapOutput() DatabaseInstanceMapOutput {
	return i.ToDatabaseInstanceMapOutputWithContext(context.Background())
}

func (i DatabaseInstanceMap) ToDatabaseInstanceMapOutputWithContext(ctx context.Context) DatabaseInstanceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseInstanceMapOutput)
}

type DatabaseInstanceOutput struct {
	*pulumi.OutputState
}

func (DatabaseInstanceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DatabaseInstance)(nil))
}

func (o DatabaseInstanceOutput) ToDatabaseInstanceOutput() DatabaseInstanceOutput {
	return o
}

func (o DatabaseInstanceOutput) ToDatabaseInstanceOutputWithContext(ctx context.Context) DatabaseInstanceOutput {
	return o
}

func (o DatabaseInstanceOutput) ToDatabaseInstancePtrOutput() DatabaseInstancePtrOutput {
	return o.ToDatabaseInstancePtrOutputWithContext(context.Background())
}

func (o DatabaseInstanceOutput) ToDatabaseInstancePtrOutputWithContext(ctx context.Context) DatabaseInstancePtrOutput {
	return o.ApplyT(func(v DatabaseInstance) *DatabaseInstance {
		return &v
	}).(DatabaseInstancePtrOutput)
}

type DatabaseInstancePtrOutput struct {
	*pulumi.OutputState
}

func (DatabaseInstancePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DatabaseInstance)(nil))
}

func (o DatabaseInstancePtrOutput) ToDatabaseInstancePtrOutput() DatabaseInstancePtrOutput {
	return o
}

func (o DatabaseInstancePtrOutput) ToDatabaseInstancePtrOutputWithContext(ctx context.Context) DatabaseInstancePtrOutput {
	return o
}

type DatabaseInstanceArrayOutput struct{ *pulumi.OutputState }

func (DatabaseInstanceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DatabaseInstance)(nil))
}

func (o DatabaseInstanceArrayOutput) ToDatabaseInstanceArrayOutput() DatabaseInstanceArrayOutput {
	return o
}

func (o DatabaseInstanceArrayOutput) ToDatabaseInstanceArrayOutputWithContext(ctx context.Context) DatabaseInstanceArrayOutput {
	return o
}

func (o DatabaseInstanceArrayOutput) Index(i pulumi.IntInput) DatabaseInstanceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DatabaseInstance {
		return vs[0].([]DatabaseInstance)[vs[1].(int)]
	}).(DatabaseInstanceOutput)
}

type DatabaseInstanceMapOutput struct{ *pulumi.OutputState }

func (DatabaseInstanceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]DatabaseInstance)(nil))
}

func (o DatabaseInstanceMapOutput) ToDatabaseInstanceMapOutput() DatabaseInstanceMapOutput {
	return o
}

func (o DatabaseInstanceMapOutput) ToDatabaseInstanceMapOutputWithContext(ctx context.Context) DatabaseInstanceMapOutput {
	return o
}

func (o DatabaseInstanceMapOutput) MapIndex(k pulumi.StringInput) DatabaseInstanceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) DatabaseInstance {
		return vs[0].(map[string]DatabaseInstance)[vs[1].(string)]
	}).(DatabaseInstanceOutput)
}

func init() {
	pulumi.RegisterOutputType(DatabaseInstanceOutput{})
	pulumi.RegisterOutputType(DatabaseInstancePtrOutput{})
	pulumi.RegisterOutputType(DatabaseInstanceArrayOutput{})
	pulumi.RegisterOutputType(DatabaseInstanceMapOutput{})
}
