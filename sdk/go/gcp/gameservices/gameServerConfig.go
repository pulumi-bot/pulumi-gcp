// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package gameservices

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// A game server config resource. Configs are global and immutable.
//
// To get more information about GameServerConfig, see:
//
// * [API documentation](https://cloud.google.com/game-servers/docs/reference/rest/v1beta/projects.locations.gameServerDeployments.configs)
// * How-to Guides
//     * [Official Documentation](https://cloud.google.com/game-servers/docs)
//
// ## Example Usage
//
// ## Import
//
// GameServerConfig can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import gcp:gameservices/gameServerConfig:GameServerConfig default projects/{{project}}/locations/{{location}}/gameServerDeployments/{{deployment_id}}/configs/{{config_id}}
// ```
//
// ```sh
//  $ pulumi import gcp:gameservices/gameServerConfig:GameServerConfig default {{project}}/{{location}}/{{deployment_id}}/{{config_id}}
// ```
//
// ```sh
//  $ pulumi import gcp:gameservices/gameServerConfig:GameServerConfig default {{location}}/{{deployment_id}}/{{config_id}}
// ```
type GameServerConfig struct {
	pulumi.CustomResourceState

	// A unique id for the deployment config.
	ConfigId pulumi.StringOutput `pulumi:"configId"`
	// A unique id for the deployment.
	DeploymentId pulumi.StringOutput `pulumi:"deploymentId"`
	// The description of the game server config.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The fleet config contains list of fleet specs. In the Single Cloud, there
	// will be only one.
	// Structure is documented below.
	FleetConfigs GameServerConfigFleetConfigArrayOutput `pulumi:"fleetConfigs"`
	// Set of labels to group by.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// Location of the Deployment.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The name of the ScalingConfig
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// Optional. This contains the autoscaling settings.
	// Structure is documented below.
	ScalingConfigs GameServerConfigScalingConfigArrayOutput `pulumi:"scalingConfigs"`
}

// NewGameServerConfig registers a new resource with the given unique name, arguments, and options.
func NewGameServerConfig(ctx *pulumi.Context,
	name string, args *GameServerConfigArgs, opts ...pulumi.ResourceOption) (*GameServerConfig, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ConfigId == nil {
		return nil, errors.New("invalid value for required argument 'ConfigId'")
	}
	if args.DeploymentId == nil {
		return nil, errors.New("invalid value for required argument 'DeploymentId'")
	}
	if args.FleetConfigs == nil {
		return nil, errors.New("invalid value for required argument 'FleetConfigs'")
	}
	var resource GameServerConfig
	err := ctx.RegisterResource("gcp:gameservices/gameServerConfig:GameServerConfig", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGameServerConfig gets an existing GameServerConfig resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGameServerConfig(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GameServerConfigState, opts ...pulumi.ResourceOption) (*GameServerConfig, error) {
	var resource GameServerConfig
	err := ctx.ReadResource("gcp:gameservices/gameServerConfig:GameServerConfig", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GameServerConfig resources.
type gameServerConfigState struct {
	// A unique id for the deployment config.
	ConfigId *string `pulumi:"configId"`
	// A unique id for the deployment.
	DeploymentId *string `pulumi:"deploymentId"`
	// The description of the game server config.
	Description *string `pulumi:"description"`
	// The fleet config contains list of fleet specs. In the Single Cloud, there
	// will be only one.
	// Structure is documented below.
	FleetConfigs []GameServerConfigFleetConfig `pulumi:"fleetConfigs"`
	// Set of labels to group by.
	Labels map[string]string `pulumi:"labels"`
	// Location of the Deployment.
	Location *string `pulumi:"location"`
	// The name of the ScalingConfig
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// Optional. This contains the autoscaling settings.
	// Structure is documented below.
	ScalingConfigs []GameServerConfigScalingConfig `pulumi:"scalingConfigs"`
}

type GameServerConfigState struct {
	// A unique id for the deployment config.
	ConfigId pulumi.StringPtrInput
	// A unique id for the deployment.
	DeploymentId pulumi.StringPtrInput
	// The description of the game server config.
	Description pulumi.StringPtrInput
	// The fleet config contains list of fleet specs. In the Single Cloud, there
	// will be only one.
	// Structure is documented below.
	FleetConfigs GameServerConfigFleetConfigArrayInput
	// Set of labels to group by.
	Labels pulumi.StringMapInput
	// Location of the Deployment.
	Location pulumi.StringPtrInput
	// The name of the ScalingConfig
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// Optional. This contains the autoscaling settings.
	// Structure is documented below.
	ScalingConfigs GameServerConfigScalingConfigArrayInput
}

func (GameServerConfigState) ElementType() reflect.Type {
	return reflect.TypeOf((*gameServerConfigState)(nil)).Elem()
}

type gameServerConfigArgs struct {
	// A unique id for the deployment config.
	ConfigId string `pulumi:"configId"`
	// A unique id for the deployment.
	DeploymentId string `pulumi:"deploymentId"`
	// The description of the game server config.
	Description *string `pulumi:"description"`
	// The fleet config contains list of fleet specs. In the Single Cloud, there
	// will be only one.
	// Structure is documented below.
	FleetConfigs []GameServerConfigFleetConfig `pulumi:"fleetConfigs"`
	// Set of labels to group by.
	Labels map[string]string `pulumi:"labels"`
	// Location of the Deployment.
	Location *string `pulumi:"location"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// Optional. This contains the autoscaling settings.
	// Structure is documented below.
	ScalingConfigs []GameServerConfigScalingConfig `pulumi:"scalingConfigs"`
}

// The set of arguments for constructing a GameServerConfig resource.
type GameServerConfigArgs struct {
	// A unique id for the deployment config.
	ConfigId pulumi.StringInput
	// A unique id for the deployment.
	DeploymentId pulumi.StringInput
	// The description of the game server config.
	Description pulumi.StringPtrInput
	// The fleet config contains list of fleet specs. In the Single Cloud, there
	// will be only one.
	// Structure is documented below.
	FleetConfigs GameServerConfigFleetConfigArrayInput
	// Set of labels to group by.
	Labels pulumi.StringMapInput
	// Location of the Deployment.
	Location pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// Optional. This contains the autoscaling settings.
	// Structure is documented below.
	ScalingConfigs GameServerConfigScalingConfigArrayInput
}

func (GameServerConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*gameServerConfigArgs)(nil)).Elem()
}

type GameServerConfigInput interface {
	pulumi.Input

	ToGameServerConfigOutput() GameServerConfigOutput
	ToGameServerConfigOutputWithContext(ctx context.Context) GameServerConfigOutput
}

func (*GameServerConfig) ElementType() reflect.Type {
	return reflect.TypeOf((*GameServerConfig)(nil))
}

func (i *GameServerConfig) ToGameServerConfigOutput() GameServerConfigOutput {
	return i.ToGameServerConfigOutputWithContext(context.Background())
}

func (i *GameServerConfig) ToGameServerConfigOutputWithContext(ctx context.Context) GameServerConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GameServerConfigOutput)
}

func (i *GameServerConfig) ToGameServerConfigPtrOutput() GameServerConfigPtrOutput {
	return i.ToGameServerConfigPtrOutputWithContext(context.Background())
}

func (i *GameServerConfig) ToGameServerConfigPtrOutputWithContext(ctx context.Context) GameServerConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GameServerConfigPtrOutput)
}

type GameServerConfigPtrInput interface {
	pulumi.Input

	ToGameServerConfigPtrOutput() GameServerConfigPtrOutput
	ToGameServerConfigPtrOutputWithContext(ctx context.Context) GameServerConfigPtrOutput
}

type gameServerConfigPtrType GameServerConfigArgs

func (*gameServerConfigPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**GameServerConfig)(nil))
}

func (i *gameServerConfigPtrType) ToGameServerConfigPtrOutput() GameServerConfigPtrOutput {
	return i.ToGameServerConfigPtrOutputWithContext(context.Background())
}

func (i *gameServerConfigPtrType) ToGameServerConfigPtrOutputWithContext(ctx context.Context) GameServerConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GameServerConfigOutput).ToGameServerConfigPtrOutput()
}

// GameServerConfigArrayInput is an input type that accepts GameServerConfigArray and GameServerConfigArrayOutput values.
// You can construct a concrete instance of `GameServerConfigArrayInput` via:
//
//          GameServerConfigArray{ GameServerConfigArgs{...} }
type GameServerConfigArrayInput interface {
	pulumi.Input

	ToGameServerConfigArrayOutput() GameServerConfigArrayOutput
	ToGameServerConfigArrayOutputWithContext(context.Context) GameServerConfigArrayOutput
}

type GameServerConfigArray []GameServerConfigInput

func (GameServerConfigArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GameServerConfig)(nil))
}

func (i GameServerConfigArray) ToGameServerConfigArrayOutput() GameServerConfigArrayOutput {
	return i.ToGameServerConfigArrayOutputWithContext(context.Background())
}

func (i GameServerConfigArray) ToGameServerConfigArrayOutputWithContext(ctx context.Context) GameServerConfigArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GameServerConfigArrayOutput)
}

// GameServerConfigMapInput is an input type that accepts GameServerConfigMap and GameServerConfigMapOutput values.
// You can construct a concrete instance of `GameServerConfigMapInput` via:
//
//          GameServerConfigMap{ "key": GameServerConfigArgs{...} }
type GameServerConfigMapInput interface {
	pulumi.Input

	ToGameServerConfigMapOutput() GameServerConfigMapOutput
	ToGameServerConfigMapOutputWithContext(context.Context) GameServerConfigMapOutput
}

type GameServerConfigMap map[string]GameServerConfigInput

func (GameServerConfigMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]GameServerConfig)(nil))
}

func (i GameServerConfigMap) ToGameServerConfigMapOutput() GameServerConfigMapOutput {
	return i.ToGameServerConfigMapOutputWithContext(context.Background())
}

func (i GameServerConfigMap) ToGameServerConfigMapOutputWithContext(ctx context.Context) GameServerConfigMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GameServerConfigMapOutput)
}

type GameServerConfigOutput struct {
	*pulumi.OutputState
}

func (GameServerConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GameServerConfig)(nil))
}

func (o GameServerConfigOutput) ToGameServerConfigOutput() GameServerConfigOutput {
	return o
}

func (o GameServerConfigOutput) ToGameServerConfigOutputWithContext(ctx context.Context) GameServerConfigOutput {
	return o
}

func (o GameServerConfigOutput) ToGameServerConfigPtrOutput() GameServerConfigPtrOutput {
	return o.ToGameServerConfigPtrOutputWithContext(context.Background())
}

func (o GameServerConfigOutput) ToGameServerConfigPtrOutputWithContext(ctx context.Context) GameServerConfigPtrOutput {
	return o.ApplyT(func(v GameServerConfig) *GameServerConfig {
		return &v
	}).(GameServerConfigPtrOutput)
}

type GameServerConfigPtrOutput struct {
	*pulumi.OutputState
}

func (GameServerConfigPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GameServerConfig)(nil))
}

func (o GameServerConfigPtrOutput) ToGameServerConfigPtrOutput() GameServerConfigPtrOutput {
	return o
}

func (o GameServerConfigPtrOutput) ToGameServerConfigPtrOutputWithContext(ctx context.Context) GameServerConfigPtrOutput {
	return o
}

type GameServerConfigArrayOutput struct{ *pulumi.OutputState }

func (GameServerConfigArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GameServerConfig)(nil))
}

func (o GameServerConfigArrayOutput) ToGameServerConfigArrayOutput() GameServerConfigArrayOutput {
	return o
}

func (o GameServerConfigArrayOutput) ToGameServerConfigArrayOutputWithContext(ctx context.Context) GameServerConfigArrayOutput {
	return o
}

func (o GameServerConfigArrayOutput) Index(i pulumi.IntInput) GameServerConfigOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GameServerConfig {
		return vs[0].([]GameServerConfig)[vs[1].(int)]
	}).(GameServerConfigOutput)
}

type GameServerConfigMapOutput struct{ *pulumi.OutputState }

func (GameServerConfigMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]GameServerConfig)(nil))
}

func (o GameServerConfigMapOutput) ToGameServerConfigMapOutput() GameServerConfigMapOutput {
	return o
}

func (o GameServerConfigMapOutput) ToGameServerConfigMapOutputWithContext(ctx context.Context) GameServerConfigMapOutput {
	return o
}

func (o GameServerConfigMapOutput) MapIndex(k pulumi.StringInput) GameServerConfigOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) GameServerConfig {
		return vs[0].(map[string]GameServerConfig)[vs[1].(string)]
	}).(GameServerConfigOutput)
}

func init() {
	pulumi.RegisterOutputType(GameServerConfigOutput{})
	pulumi.RegisterOutputType(GameServerConfigPtrOutput{})
	pulumi.RegisterOutputType(GameServerConfigArrayOutput{})
	pulumi.RegisterOutputType(GameServerConfigMapOutput{})
}
