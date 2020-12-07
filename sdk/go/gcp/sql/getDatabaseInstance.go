// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package sql

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Use this data source to get information about a Cloud SQL instance
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/sql"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := sql.LookupDatabaseInstance(ctx, &sql.LookupDatabaseInstanceArgs{
// 			Name: google_sql_database_instance.Master.Name,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupDatabaseInstance(ctx *pulumi.Context, args *LookupDatabaseInstanceArgs, opts ...pulumi.InvokeOption) (*LookupDatabaseInstanceResult, error) {
	var rv LookupDatabaseInstanceResult
	err := ctx.Invoke("gcp:sql/getDatabaseInstance:getDatabaseInstance", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDatabaseInstance.
type LookupDatabaseInstanceArgs struct {
	// The name of the instance.
	Name string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	Project *string `pulumi:"project"`
}

// A collection of values returned by getDatabaseInstance.
type LookupDatabaseInstanceResult struct {
	// The connection name of the instance to be used in connection strings.
	ConnectionName string `pulumi:"connectionName"`
	// The MySQL, PostgreSQL or SQL Server (beta) version to use.
	DatabaseVersion    string `pulumi:"databaseVersion"`
	DeletionProtection bool   `pulumi:"deletionProtection"`
	// The full path to the encryption key used for the CMEK disk encryption.
	EncryptionKeyName string `pulumi:"encryptionKeyName"`
	// The first IPv4 address of any type assigned.
	FirstIpAddress string `pulumi:"firstIpAddress"`
	// The provider-assigned unique ID for this managed resource.
	Id          string                         `pulumi:"id"`
	IpAddresses []GetDatabaseInstanceIpAddress `pulumi:"ipAddresses"`
	// The name of the instance that will act as
	// the master in the replication setup.
	MasterInstanceName string `pulumi:"masterInstanceName"`
	// A name for this whitelist entry.
	Name string `pulumi:"name"`
	// The first private (`PRIVATE`) IPv4 address assigned.
	PrivateIpAddress string  `pulumi:"privateIpAddress"`
	Project          *string `pulumi:"project"`
	// The first public (`PRIMARY`) IPv4 address assigned.
	PublicIpAddress string `pulumi:"publicIpAddress"`
	Region          string `pulumi:"region"`
	// The configuration for replication. The
	// configuration is detailed below.
	ReplicaConfigurations []GetDatabaseInstanceReplicaConfiguration `pulumi:"replicaConfigurations"`
	// Initial root password. Required for MS SQL Server, ignored by MySQL and PostgreSQL.
	RootPassword string `pulumi:"rootPassword"`
	// The URI of the created resource.
	SelfLink      string                            `pulumi:"selfLink"`
	ServerCaCerts []GetDatabaseInstanceServerCaCert `pulumi:"serverCaCerts"`
	// The service account email address assigned to the instance.
	ServiceAccountEmailAddress string `pulumi:"serviceAccountEmailAddress"`
	// The settings to use for the database. The
	// configuration is detailed below.
	Settings []GetDatabaseInstanceSetting `pulumi:"settings"`
}
