// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package gameservices

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get the rollout state.
//
// https://cloud.google.com/game-servers/docs/reference/rest/v1beta/GameServerDeploymentRollout
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/gameservices"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := gameservices.LookupGameServerDeploymentRollout(ctx, &gameservices.LookupGameServerDeploymentRolloutArgs{
// 			DeploymentId: "tf-test-deployment-s8sn12jt2c",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupGameServerDeploymentRollout(ctx *pulumi.Context, args *LookupGameServerDeploymentRolloutArgs, opts ...pulumi.InvokeOption) (*LookupGameServerDeploymentRolloutResult, error) {
	var rv LookupGameServerDeploymentRolloutResult
	err := ctx.Invoke("gcp:gameservices/getGameServerDeploymentRollout:getGameServerDeploymentRollout", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getGameServerDeploymentRollout.
type LookupGameServerDeploymentRolloutArgs struct {
	// The deployment to get the rollout state from. Only 1 rollout must be associated with each deployment.
	DeploymentId string `pulumi:"deploymentId"`
}

// A collection of values returned by getGameServerDeploymentRollout.
type LookupGameServerDeploymentRolloutResult struct {
	DefaultGameServerConfig   string                                                   `pulumi:"defaultGameServerConfig"`
	DeploymentId              string                                                   `pulumi:"deploymentId"`
	GameServerConfigOverrides []GetGameServerDeploymentRolloutGameServerConfigOverride `pulumi:"gameServerConfigOverrides"`
	// The provider-assigned unique ID for this managed resource.
	Id   string `pulumi:"id"`
	Name string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project string `pulumi:"project"`
}

func LookupGameServerDeploymentRolloutApply(ctx *pulumi.Context, args LookupGameServerDeploymentRolloutApplyInput, opts ...pulumi.InvokeOption) LookupGameServerDeploymentRolloutResultOutput {
	return args.ToLookupGameServerDeploymentRolloutApplyOutput().ApplyT(func(v LookupGameServerDeploymentRolloutArgs) (LookupGameServerDeploymentRolloutResult, error) {
		r, err := LookupGameServerDeploymentRollout(ctx, &v, opts...)
		return *r, err

	}).(LookupGameServerDeploymentRolloutResultOutput)
}

// LookupGameServerDeploymentRolloutApplyInput is an input type that accepts LookupGameServerDeploymentRolloutApplyArgs and LookupGameServerDeploymentRolloutApplyOutput values.
// You can construct a concrete instance of `LookupGameServerDeploymentRolloutApplyInput` via:
//
//          LookupGameServerDeploymentRolloutApplyArgs{...}
type LookupGameServerDeploymentRolloutApplyInput interface {
	pulumi.Input

	ToLookupGameServerDeploymentRolloutApplyOutput() LookupGameServerDeploymentRolloutApplyOutput
	ToLookupGameServerDeploymentRolloutApplyOutputWithContext(context.Context) LookupGameServerDeploymentRolloutApplyOutput
}

// A collection of arguments for invoking getGameServerDeploymentRollout.
type LookupGameServerDeploymentRolloutApplyArgs struct {
	// The deployment to get the rollout state from. Only 1 rollout must be associated with each deployment.
	DeploymentId pulumi.StringInput `pulumi:"deploymentId"`
}

func (LookupGameServerDeploymentRolloutApplyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGameServerDeploymentRolloutArgs)(nil)).Elem()
}

func (i LookupGameServerDeploymentRolloutApplyArgs) ToLookupGameServerDeploymentRolloutApplyOutput() LookupGameServerDeploymentRolloutApplyOutput {
	return i.ToLookupGameServerDeploymentRolloutApplyOutputWithContext(context.Background())
}

func (i LookupGameServerDeploymentRolloutApplyArgs) ToLookupGameServerDeploymentRolloutApplyOutputWithContext(ctx context.Context) LookupGameServerDeploymentRolloutApplyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LookupGameServerDeploymentRolloutApplyOutput)
}

// A collection of arguments for invoking getGameServerDeploymentRollout.
type LookupGameServerDeploymentRolloutApplyOutput struct{ *pulumi.OutputState }

func (LookupGameServerDeploymentRolloutApplyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGameServerDeploymentRolloutArgs)(nil)).Elem()
}

func (o LookupGameServerDeploymentRolloutApplyOutput) ToLookupGameServerDeploymentRolloutApplyOutput() LookupGameServerDeploymentRolloutApplyOutput {
	return o
}

func (o LookupGameServerDeploymentRolloutApplyOutput) ToLookupGameServerDeploymentRolloutApplyOutputWithContext(ctx context.Context) LookupGameServerDeploymentRolloutApplyOutput {
	return o
}

// The deployment to get the rollout state from. Only 1 rollout must be associated with each deployment.
func (o LookupGameServerDeploymentRolloutApplyOutput) DeploymentId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGameServerDeploymentRolloutArgs) string { return v.DeploymentId }).(pulumi.StringOutput)
}

// A collection of values returned by getGameServerDeploymentRollout.
type LookupGameServerDeploymentRolloutResultOutput struct{ *pulumi.OutputState }

func (LookupGameServerDeploymentRolloutResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGameServerDeploymentRolloutResult)(nil)).Elem()
}

func (o LookupGameServerDeploymentRolloutResultOutput) ToLookupGameServerDeploymentRolloutResultOutput() LookupGameServerDeploymentRolloutResultOutput {
	return o
}

func (o LookupGameServerDeploymentRolloutResultOutput) ToLookupGameServerDeploymentRolloutResultOutputWithContext(ctx context.Context) LookupGameServerDeploymentRolloutResultOutput {
	return o
}

func (o LookupGameServerDeploymentRolloutResultOutput) DefaultGameServerConfig() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGameServerDeploymentRolloutResult) string { return v.DefaultGameServerConfig }).(pulumi.StringOutput)
}

func (o LookupGameServerDeploymentRolloutResultOutput) DeploymentId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGameServerDeploymentRolloutResult) string { return v.DeploymentId }).(pulumi.StringOutput)
}

func (o LookupGameServerDeploymentRolloutResultOutput) GameServerConfigOverrides() GetGameServerDeploymentRolloutGameServerConfigOverrideArrayOutput {
	return o.ApplyT(func(v LookupGameServerDeploymentRolloutResult) []GetGameServerDeploymentRolloutGameServerConfigOverride {
		return v.GameServerConfigOverrides
	}).(GetGameServerDeploymentRolloutGameServerConfigOverrideArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupGameServerDeploymentRolloutResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGameServerDeploymentRolloutResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupGameServerDeploymentRolloutResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGameServerDeploymentRolloutResult) string { return v.Name }).(pulumi.StringOutput)
}

// The ID of the project in which the resource belongs.
// If it is not provided, the provider project is used.
func (o LookupGameServerDeploymentRolloutResultOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGameServerDeploymentRolloutResult) string { return v.Project }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupGameServerDeploymentRolloutApplyOutput{})
	pulumi.RegisterOutputType(LookupGameServerDeploymentRolloutResultOutput{})
}
