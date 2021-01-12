// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package redis

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// A Google Cloud Redis instance.
//
// To get more information about Instance, see:
//
// * [API documentation](https://cloud.google.com/memorystore/docs/redis/reference/rest/)
// * How-to Guides
//     * [Official Documentation](https://cloud.google.com/memorystore/docs/redis/)
//
// ## Example Usage
// ### Redis Instance Basic
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/redis"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := redis.NewInstance(ctx, "cache", &redis.InstanceArgs{
// 			MemorySizeGb: pulumi.Int(1),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Redis Instance Full
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/compute"
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/redis"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		redis_network, err := compute.LookupNetwork(ctx, &compute.LookupNetworkArgs{
// 			Name: "redis-test-network",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = redis.NewInstance(ctx, "cache", &redis.InstanceArgs{
// 			Tier:                  pulumi.String("STANDARD_HA"),
// 			MemorySizeGb:          pulumi.Int(1),
// 			LocationId:            pulumi.String("us-central1-a"),
// 			AlternativeLocationId: pulumi.String("us-central1-f"),
// 			AuthorizedNetwork:     pulumi.String(redis_network.Id),
// 			RedisVersion:          pulumi.String("REDIS_4_0"),
// 			DisplayName:           pulumi.String("Test Instance"),
// 			ReservedIpRange:       pulumi.String("192.168.0.0/29"),
// 			Labels: pulumi.StringMap{
// 				"my_key":    pulumi.String("my_val"),
// 				"other_key": pulumi.String("other_val"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// Instance can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import gcp:redis/instance:Instance default projects/{{project}}/locations/{{region}}/instances/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:redis/instance:Instance default {{project}}/{{region}}/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:redis/instance:Instance default {{region}}/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:redis/instance:Instance default {{name}}
// ```
type Instance struct {
	pulumi.CustomResourceState

	// Only applicable to STANDARD_HA tier which protects the instance
	// against zonal failures by provisioning it across two zones.
	// If provided, it must be a different zone from the one provided in
	// [locationId].
	AlternativeLocationId pulumi.StringOutput `pulumi:"alternativeLocationId"`
	// Optional. Indicates whether OSS Redis AUTH is enabled for the
	// instance. If set to "true" AUTH is enabled on the instance.
	// Default value is "false" meaning AUTH is disabled.
	AuthEnabled pulumi.BoolPtrOutput `pulumi:"authEnabled"`
	// AUTH String set on the instance. This field will only be populated if authEnabled is true.
	AuthString pulumi.StringOutput `pulumi:"authString"`
	// The full name of the Google Compute Engine network to which the
	// instance is connected. If left unspecified, the default network
	// will be used.
	AuthorizedNetwork pulumi.StringOutput `pulumi:"authorizedNetwork"`
	// The connection mode of the Redis instance.
	// Default value is `DIRECT_PEERING`.
	// Possible values are `DIRECT_PEERING` and `PRIVATE_SERVICE_ACCESS`.
	ConnectMode pulumi.StringPtrOutput `pulumi:"connectMode"`
	// The time the instance was created in RFC3339 UTC "Zulu" format, accurate to nanoseconds.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// The current zone where the Redis endpoint is placed. For Basic Tier instances, this will always be the same as the
	// [locationId] provided by the user at creation time. For Standard Tier instances, this can be either [locationId] or
	// [alternativeLocationId] and can change after a failover event.
	CurrentLocationId pulumi.StringOutput `pulumi:"currentLocationId"`
	// An arbitrary and optional user-provided name for the instance.
	DisplayName pulumi.StringPtrOutput `pulumi:"displayName"`
	// Hostname or IP address of the exposed Redis endpoint used by clients to connect to the service.
	Host pulumi.StringOutput `pulumi:"host"`
	// Resource labels to represent user provided metadata.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// The zone where the instance will be provisioned. If not provided,
	// the service will choose a zone for the instance. For STANDARD_HA tier,
	// instances will be created across two zones for protection against
	// zonal failures. If [alternativeLocationId] is also provided, it must
	// be different from [locationId].
	LocationId pulumi.StringOutput `pulumi:"locationId"`
	// Redis memory size in GiB.
	MemorySizeGb pulumi.IntOutput `pulumi:"memorySizeGb"`
	// The ID of the instance or a fully qualified identifier for the instance.
	Name pulumi.StringOutput `pulumi:"name"`
	// Output only. Cloud IAM identity used by import / export operations to transfer data to/from Cloud Storage. Format is
	// "serviceAccount:". The value may change over time for a given instance so should be checked before each import/export
	// operation.
	PersistenceIamIdentity pulumi.StringOutput `pulumi:"persistenceIamIdentity"`
	// The port number of the exposed Redis endpoint.
	Port pulumi.IntOutput `pulumi:"port"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// Redis configuration parameters, according to http://redis.io/topics/config.
	// Please check Memorystore documentation for the list of supported parameters:
	// https://cloud.google.com/memorystore/docs/redis/reference/rest/v1/projects.locations.instances#Instance.FIELDS.redis_configs
	RedisConfigs pulumi.StringMapOutput `pulumi:"redisConfigs"`
	// The version of Redis software. If not provided, latest supported
	// version will be used. Currently, the supported values are:
	// - REDIS_5_0 for Redis 5.0 compatibility
	// - REDIS_4_0 for Redis 4.0 compatibility
	// - REDIS_3_2 for Redis 3.2 compatibility
	RedisVersion pulumi.StringOutput `pulumi:"redisVersion"`
	// The name of the Redis region of the instance.
	Region pulumi.StringOutput `pulumi:"region"`
	// The CIDR range of internal addresses that are reserved for this
	// instance. If not provided, the service will choose an unused /29
	// block, for example, 10.0.0.0/29 or 192.168.0.0/29. Ranges must be
	// unique and non-overlapping with existing subnets in an authorized
	// network.
	ReservedIpRange pulumi.StringOutput `pulumi:"reservedIpRange"`
	// The service tier of the instance. Must be one of these values:
	// - BASIC: standalone instance
	// - STANDARD_HA: highly available primary/replica instances
	//   Default value is `BASIC`.
	//   Possible values are `BASIC` and `STANDARD_HA`.
	Tier pulumi.StringPtrOutput `pulumi:"tier"`
}

// NewInstance registers a new resource with the given unique name, arguments, and options.
func NewInstance(ctx *pulumi.Context,
	name string, args *InstanceArgs, opts ...pulumi.ResourceOption) (*Instance, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.MemorySizeGb == nil {
		return nil, errors.New("invalid value for required argument 'MemorySizeGb'")
	}
	var resource Instance
	err := ctx.RegisterResource("gcp:redis/instance:Instance", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInstance gets an existing Instance resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInstance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InstanceState, opts ...pulumi.ResourceOption) (*Instance, error) {
	var resource Instance
	err := ctx.ReadResource("gcp:redis/instance:Instance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Instance resources.
type instanceState struct {
	// Only applicable to STANDARD_HA tier which protects the instance
	// against zonal failures by provisioning it across two zones.
	// If provided, it must be a different zone from the one provided in
	// [locationId].
	AlternativeLocationId *string `pulumi:"alternativeLocationId"`
	// Optional. Indicates whether OSS Redis AUTH is enabled for the
	// instance. If set to "true" AUTH is enabled on the instance.
	// Default value is "false" meaning AUTH is disabled.
	AuthEnabled *bool `pulumi:"authEnabled"`
	// AUTH String set on the instance. This field will only be populated if authEnabled is true.
	AuthString *string `pulumi:"authString"`
	// The full name of the Google Compute Engine network to which the
	// instance is connected. If left unspecified, the default network
	// will be used.
	AuthorizedNetwork *string `pulumi:"authorizedNetwork"`
	// The connection mode of the Redis instance.
	// Default value is `DIRECT_PEERING`.
	// Possible values are `DIRECT_PEERING` and `PRIVATE_SERVICE_ACCESS`.
	ConnectMode *string `pulumi:"connectMode"`
	// The time the instance was created in RFC3339 UTC "Zulu" format, accurate to nanoseconds.
	CreateTime *string `pulumi:"createTime"`
	// The current zone where the Redis endpoint is placed. For Basic Tier instances, this will always be the same as the
	// [locationId] provided by the user at creation time. For Standard Tier instances, this can be either [locationId] or
	// [alternativeLocationId] and can change after a failover event.
	CurrentLocationId *string `pulumi:"currentLocationId"`
	// An arbitrary and optional user-provided name for the instance.
	DisplayName *string `pulumi:"displayName"`
	// Hostname or IP address of the exposed Redis endpoint used by clients to connect to the service.
	Host *string `pulumi:"host"`
	// Resource labels to represent user provided metadata.
	Labels map[string]string `pulumi:"labels"`
	// The zone where the instance will be provisioned. If not provided,
	// the service will choose a zone for the instance. For STANDARD_HA tier,
	// instances will be created across two zones for protection against
	// zonal failures. If [alternativeLocationId] is also provided, it must
	// be different from [locationId].
	LocationId *string `pulumi:"locationId"`
	// Redis memory size in GiB.
	MemorySizeGb *int `pulumi:"memorySizeGb"`
	// The ID of the instance or a fully qualified identifier for the instance.
	Name *string `pulumi:"name"`
	// Output only. Cloud IAM identity used by import / export operations to transfer data to/from Cloud Storage. Format is
	// "serviceAccount:". The value may change over time for a given instance so should be checked before each import/export
	// operation.
	PersistenceIamIdentity *string `pulumi:"persistenceIamIdentity"`
	// The port number of the exposed Redis endpoint.
	Port *int `pulumi:"port"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// Redis configuration parameters, according to http://redis.io/topics/config.
	// Please check Memorystore documentation for the list of supported parameters:
	// https://cloud.google.com/memorystore/docs/redis/reference/rest/v1/projects.locations.instances#Instance.FIELDS.redis_configs
	RedisConfigs map[string]string `pulumi:"redisConfigs"`
	// The version of Redis software. If not provided, latest supported
	// version will be used. Currently, the supported values are:
	// - REDIS_5_0 for Redis 5.0 compatibility
	// - REDIS_4_0 for Redis 4.0 compatibility
	// - REDIS_3_2 for Redis 3.2 compatibility
	RedisVersion *string `pulumi:"redisVersion"`
	// The name of the Redis region of the instance.
	Region *string `pulumi:"region"`
	// The CIDR range of internal addresses that are reserved for this
	// instance. If not provided, the service will choose an unused /29
	// block, for example, 10.0.0.0/29 or 192.168.0.0/29. Ranges must be
	// unique and non-overlapping with existing subnets in an authorized
	// network.
	ReservedIpRange *string `pulumi:"reservedIpRange"`
	// The service tier of the instance. Must be one of these values:
	// - BASIC: standalone instance
	// - STANDARD_HA: highly available primary/replica instances
	//   Default value is `BASIC`.
	//   Possible values are `BASIC` and `STANDARD_HA`.
	Tier *string `pulumi:"tier"`
}

type InstanceState struct {
	// Only applicable to STANDARD_HA tier which protects the instance
	// against zonal failures by provisioning it across two zones.
	// If provided, it must be a different zone from the one provided in
	// [locationId].
	AlternativeLocationId pulumi.StringPtrInput
	// Optional. Indicates whether OSS Redis AUTH is enabled for the
	// instance. If set to "true" AUTH is enabled on the instance.
	// Default value is "false" meaning AUTH is disabled.
	AuthEnabled pulumi.BoolPtrInput
	// AUTH String set on the instance. This field will only be populated if authEnabled is true.
	AuthString pulumi.StringPtrInput
	// The full name of the Google Compute Engine network to which the
	// instance is connected. If left unspecified, the default network
	// will be used.
	AuthorizedNetwork pulumi.StringPtrInput
	// The connection mode of the Redis instance.
	// Default value is `DIRECT_PEERING`.
	// Possible values are `DIRECT_PEERING` and `PRIVATE_SERVICE_ACCESS`.
	ConnectMode pulumi.StringPtrInput
	// The time the instance was created in RFC3339 UTC "Zulu" format, accurate to nanoseconds.
	CreateTime pulumi.StringPtrInput
	// The current zone where the Redis endpoint is placed. For Basic Tier instances, this will always be the same as the
	// [locationId] provided by the user at creation time. For Standard Tier instances, this can be either [locationId] or
	// [alternativeLocationId] and can change after a failover event.
	CurrentLocationId pulumi.StringPtrInput
	// An arbitrary and optional user-provided name for the instance.
	DisplayName pulumi.StringPtrInput
	// Hostname or IP address of the exposed Redis endpoint used by clients to connect to the service.
	Host pulumi.StringPtrInput
	// Resource labels to represent user provided metadata.
	Labels pulumi.StringMapInput
	// The zone where the instance will be provisioned. If not provided,
	// the service will choose a zone for the instance. For STANDARD_HA tier,
	// instances will be created across two zones for protection against
	// zonal failures. If [alternativeLocationId] is also provided, it must
	// be different from [locationId].
	LocationId pulumi.StringPtrInput
	// Redis memory size in GiB.
	MemorySizeGb pulumi.IntPtrInput
	// The ID of the instance or a fully qualified identifier for the instance.
	Name pulumi.StringPtrInput
	// Output only. Cloud IAM identity used by import / export operations to transfer data to/from Cloud Storage. Format is
	// "serviceAccount:". The value may change over time for a given instance so should be checked before each import/export
	// operation.
	PersistenceIamIdentity pulumi.StringPtrInput
	// The port number of the exposed Redis endpoint.
	Port pulumi.IntPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// Redis configuration parameters, according to http://redis.io/topics/config.
	// Please check Memorystore documentation for the list of supported parameters:
	// https://cloud.google.com/memorystore/docs/redis/reference/rest/v1/projects.locations.instances#Instance.FIELDS.redis_configs
	RedisConfigs pulumi.StringMapInput
	// The version of Redis software. If not provided, latest supported
	// version will be used. Currently, the supported values are:
	// - REDIS_5_0 for Redis 5.0 compatibility
	// - REDIS_4_0 for Redis 4.0 compatibility
	// - REDIS_3_2 for Redis 3.2 compatibility
	RedisVersion pulumi.StringPtrInput
	// The name of the Redis region of the instance.
	Region pulumi.StringPtrInput
	// The CIDR range of internal addresses that are reserved for this
	// instance. If not provided, the service will choose an unused /29
	// block, for example, 10.0.0.0/29 or 192.168.0.0/29. Ranges must be
	// unique and non-overlapping with existing subnets in an authorized
	// network.
	ReservedIpRange pulumi.StringPtrInput
	// The service tier of the instance. Must be one of these values:
	// - BASIC: standalone instance
	// - STANDARD_HA: highly available primary/replica instances
	//   Default value is `BASIC`.
	//   Possible values are `BASIC` and `STANDARD_HA`.
	Tier pulumi.StringPtrInput
}

func (InstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceState)(nil)).Elem()
}

type instanceArgs struct {
	// Only applicable to STANDARD_HA tier which protects the instance
	// against zonal failures by provisioning it across two zones.
	// If provided, it must be a different zone from the one provided in
	// [locationId].
	AlternativeLocationId *string `pulumi:"alternativeLocationId"`
	// Optional. Indicates whether OSS Redis AUTH is enabled for the
	// instance. If set to "true" AUTH is enabled on the instance.
	// Default value is "false" meaning AUTH is disabled.
	AuthEnabled *bool `pulumi:"authEnabled"`
	// AUTH String set on the instance. This field will only be populated if authEnabled is true.
	AuthString *string `pulumi:"authString"`
	// The full name of the Google Compute Engine network to which the
	// instance is connected. If left unspecified, the default network
	// will be used.
	AuthorizedNetwork *string `pulumi:"authorizedNetwork"`
	// The connection mode of the Redis instance.
	// Default value is `DIRECT_PEERING`.
	// Possible values are `DIRECT_PEERING` and `PRIVATE_SERVICE_ACCESS`.
	ConnectMode *string `pulumi:"connectMode"`
	// An arbitrary and optional user-provided name for the instance.
	DisplayName *string `pulumi:"displayName"`
	// Resource labels to represent user provided metadata.
	Labels map[string]string `pulumi:"labels"`
	// The zone where the instance will be provisioned. If not provided,
	// the service will choose a zone for the instance. For STANDARD_HA tier,
	// instances will be created across two zones for protection against
	// zonal failures. If [alternativeLocationId] is also provided, it must
	// be different from [locationId].
	LocationId *string `pulumi:"locationId"`
	// Redis memory size in GiB.
	MemorySizeGb int `pulumi:"memorySizeGb"`
	// The ID of the instance or a fully qualified identifier for the instance.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// Redis configuration parameters, according to http://redis.io/topics/config.
	// Please check Memorystore documentation for the list of supported parameters:
	// https://cloud.google.com/memorystore/docs/redis/reference/rest/v1/projects.locations.instances#Instance.FIELDS.redis_configs
	RedisConfigs map[string]string `pulumi:"redisConfigs"`
	// The version of Redis software. If not provided, latest supported
	// version will be used. Currently, the supported values are:
	// - REDIS_5_0 for Redis 5.0 compatibility
	// - REDIS_4_0 for Redis 4.0 compatibility
	// - REDIS_3_2 for Redis 3.2 compatibility
	RedisVersion *string `pulumi:"redisVersion"`
	// The name of the Redis region of the instance.
	Region *string `pulumi:"region"`
	// The CIDR range of internal addresses that are reserved for this
	// instance. If not provided, the service will choose an unused /29
	// block, for example, 10.0.0.0/29 or 192.168.0.0/29. Ranges must be
	// unique and non-overlapping with existing subnets in an authorized
	// network.
	ReservedIpRange *string `pulumi:"reservedIpRange"`
	// The service tier of the instance. Must be one of these values:
	// - BASIC: standalone instance
	// - STANDARD_HA: highly available primary/replica instances
	//   Default value is `BASIC`.
	//   Possible values are `BASIC` and `STANDARD_HA`.
	Tier *string `pulumi:"tier"`
}

// The set of arguments for constructing a Instance resource.
type InstanceArgs struct {
	// Only applicable to STANDARD_HA tier which protects the instance
	// against zonal failures by provisioning it across two zones.
	// If provided, it must be a different zone from the one provided in
	// [locationId].
	AlternativeLocationId pulumi.StringPtrInput
	// Optional. Indicates whether OSS Redis AUTH is enabled for the
	// instance. If set to "true" AUTH is enabled on the instance.
	// Default value is "false" meaning AUTH is disabled.
	AuthEnabled pulumi.BoolPtrInput
	// AUTH String set on the instance. This field will only be populated if authEnabled is true.
	AuthString pulumi.StringPtrInput
	// The full name of the Google Compute Engine network to which the
	// instance is connected. If left unspecified, the default network
	// will be used.
	AuthorizedNetwork pulumi.StringPtrInput
	// The connection mode of the Redis instance.
	// Default value is `DIRECT_PEERING`.
	// Possible values are `DIRECT_PEERING` and `PRIVATE_SERVICE_ACCESS`.
	ConnectMode pulumi.StringPtrInput
	// An arbitrary and optional user-provided name for the instance.
	DisplayName pulumi.StringPtrInput
	// Resource labels to represent user provided metadata.
	Labels pulumi.StringMapInput
	// The zone where the instance will be provisioned. If not provided,
	// the service will choose a zone for the instance. For STANDARD_HA tier,
	// instances will be created across two zones for protection against
	// zonal failures. If [alternativeLocationId] is also provided, it must
	// be different from [locationId].
	LocationId pulumi.StringPtrInput
	// Redis memory size in GiB.
	MemorySizeGb pulumi.IntInput
	// The ID of the instance or a fully qualified identifier for the instance.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// Redis configuration parameters, according to http://redis.io/topics/config.
	// Please check Memorystore documentation for the list of supported parameters:
	// https://cloud.google.com/memorystore/docs/redis/reference/rest/v1/projects.locations.instances#Instance.FIELDS.redis_configs
	RedisConfigs pulumi.StringMapInput
	// The version of Redis software. If not provided, latest supported
	// version will be used. Currently, the supported values are:
	// - REDIS_5_0 for Redis 5.0 compatibility
	// - REDIS_4_0 for Redis 4.0 compatibility
	// - REDIS_3_2 for Redis 3.2 compatibility
	RedisVersion pulumi.StringPtrInput
	// The name of the Redis region of the instance.
	Region pulumi.StringPtrInput
	// The CIDR range of internal addresses that are reserved for this
	// instance. If not provided, the service will choose an unused /29
	// block, for example, 10.0.0.0/29 or 192.168.0.0/29. Ranges must be
	// unique and non-overlapping with existing subnets in an authorized
	// network.
	ReservedIpRange pulumi.StringPtrInput
	// The service tier of the instance. Must be one of these values:
	// - BASIC: standalone instance
	// - STANDARD_HA: highly available primary/replica instances
	//   Default value is `BASIC`.
	//   Possible values are `BASIC` and `STANDARD_HA`.
	Tier pulumi.StringPtrInput
}

func (InstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceArgs)(nil)).Elem()
}

type InstanceInput interface {
	pulumi.Input

	ToInstanceOutput() InstanceOutput
	ToInstanceOutputWithContext(ctx context.Context) InstanceOutput
}

func (*Instance) ElementType() reflect.Type {
	return reflect.TypeOf((*Instance)(nil))
}

func (i *Instance) ToInstanceOutput() InstanceOutput {
	return i.ToInstanceOutputWithContext(context.Background())
}

func (i *Instance) ToInstanceOutputWithContext(ctx context.Context) InstanceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceOutput)
}

func (i *Instance) ToInstancePtrOutput() InstancePtrOutput {
	return i.ToInstancePtrOutputWithContext(context.Background())
}

func (i *Instance) ToInstancePtrOutputWithContext(ctx context.Context) InstancePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstancePtrOutput)
}

type InstancePtrInput interface {
	pulumi.Input

	ToInstancePtrOutput() InstancePtrOutput
	ToInstancePtrOutputWithContext(ctx context.Context) InstancePtrOutput
}

type instancePtrType InstanceArgs

func (*instancePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Instance)(nil))
}

func (i *instancePtrType) ToInstancePtrOutput() InstancePtrOutput {
	return i.ToInstancePtrOutputWithContext(context.Background())
}

func (i *instancePtrType) ToInstancePtrOutputWithContext(ctx context.Context) InstancePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceOutput).ToInstancePtrOutput()
}

// InstanceArrayInput is an input type that accepts InstanceArray and InstanceArrayOutput values.
// You can construct a concrete instance of `InstanceArrayInput` via:
//
//          InstanceArray{ InstanceArgs{...} }
type InstanceArrayInput interface {
	pulumi.Input

	ToInstanceArrayOutput() InstanceArrayOutput
	ToInstanceArrayOutputWithContext(context.Context) InstanceArrayOutput
}

type InstanceArray []InstanceInput

func (InstanceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Instance)(nil))
}

func (i InstanceArray) ToInstanceArrayOutput() InstanceArrayOutput {
	return i.ToInstanceArrayOutputWithContext(context.Background())
}

func (i InstanceArray) ToInstanceArrayOutputWithContext(ctx context.Context) InstanceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceArrayOutput)
}

// InstanceMapInput is an input type that accepts InstanceMap and InstanceMapOutput values.
// You can construct a concrete instance of `InstanceMapInput` via:
//
//          InstanceMap{ "key": InstanceArgs{...} }
type InstanceMapInput interface {
	pulumi.Input

	ToInstanceMapOutput() InstanceMapOutput
	ToInstanceMapOutputWithContext(context.Context) InstanceMapOutput
}

type InstanceMap map[string]InstanceInput

func (InstanceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Instance)(nil))
}

func (i InstanceMap) ToInstanceMapOutput() InstanceMapOutput {
	return i.ToInstanceMapOutputWithContext(context.Background())
}

func (i InstanceMap) ToInstanceMapOutputWithContext(ctx context.Context) InstanceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceMapOutput)
}

type InstanceOutput struct {
	*pulumi.OutputState
}

func (InstanceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Instance)(nil))
}

func (o InstanceOutput) ToInstanceOutput() InstanceOutput {
	return o
}

func (o InstanceOutput) ToInstanceOutputWithContext(ctx context.Context) InstanceOutput {
	return o
}

func (o InstanceOutput) ToInstancePtrOutput() InstancePtrOutput {
	return o.ToInstancePtrOutputWithContext(context.Background())
}

func (o InstanceOutput) ToInstancePtrOutputWithContext(ctx context.Context) InstancePtrOutput {
	return o.ApplyT(func(v Instance) *Instance {
		return &v
	}).(InstancePtrOutput)
}

type InstancePtrOutput struct {
	*pulumi.OutputState
}

func (InstancePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Instance)(nil))
}

func (o InstancePtrOutput) ToInstancePtrOutput() InstancePtrOutput {
	return o
}

func (o InstancePtrOutput) ToInstancePtrOutputWithContext(ctx context.Context) InstancePtrOutput {
	return o
}

type InstanceArrayOutput struct{ *pulumi.OutputState }

func (InstanceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Instance)(nil))
}

func (o InstanceArrayOutput) ToInstanceArrayOutput() InstanceArrayOutput {
	return o
}

func (o InstanceArrayOutput) ToInstanceArrayOutputWithContext(ctx context.Context) InstanceArrayOutput {
	return o
}

func (o InstanceArrayOutput) Index(i pulumi.IntInput) InstanceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Instance {
		return vs[0].([]Instance)[vs[1].(int)]
	}).(InstanceOutput)
}

type InstanceMapOutput struct{ *pulumi.OutputState }

func (InstanceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Instance)(nil))
}

func (o InstanceMapOutput) ToInstanceMapOutput() InstanceMapOutput {
	return o
}

func (o InstanceMapOutput) ToInstanceMapOutputWithContext(ctx context.Context) InstanceMapOutput {
	return o
}

func (o InstanceMapOutput) MapIndex(k pulumi.StringInput) InstanceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Instance {
		return vs[0].(map[string]Instance)[vs[1].(string)]
	}).(InstanceOutput)
}

func init() {
	pulumi.RegisterOutputType(InstanceOutput{})
	pulumi.RegisterOutputType(InstancePtrOutput{})
	pulumi.RegisterOutputType(InstanceArrayOutput{})
	pulumi.RegisterOutputType(InstanceMapOutput{})
}
