// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package gkehub

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type MembershipAuthority struct {
	// A JSON Web Token (JWT) issuer URI. `issuer` must start with `https://` and // be a valid
	// with length <2000 characters.
	Issuer string `pulumi:"issuer"`
}

// MembershipAuthorityInput is an input type that accepts MembershipAuthorityArgs and MembershipAuthorityOutput values.
// You can construct a concrete instance of `MembershipAuthorityInput` via:
//
//          MembershipAuthorityArgs{...}
type MembershipAuthorityInput interface {
	pulumi.Input

	ToMembershipAuthorityOutput() MembershipAuthorityOutput
	ToMembershipAuthorityOutputWithContext(context.Context) MembershipAuthorityOutput
}

type MembershipAuthorityArgs struct {
	// A JSON Web Token (JWT) issuer URI. `issuer` must start with `https://` and // be a valid
	// with length <2000 characters.
	Issuer pulumi.StringInput `pulumi:"issuer"`
}

func (MembershipAuthorityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MembershipAuthority)(nil)).Elem()
}

func (i MembershipAuthorityArgs) ToMembershipAuthorityOutput() MembershipAuthorityOutput {
	return i.ToMembershipAuthorityOutputWithContext(context.Background())
}

func (i MembershipAuthorityArgs) ToMembershipAuthorityOutputWithContext(ctx context.Context) MembershipAuthorityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MembershipAuthorityOutput)
}

func (i MembershipAuthorityArgs) ToMembershipAuthorityPtrOutput() MembershipAuthorityPtrOutput {
	return i.ToMembershipAuthorityPtrOutputWithContext(context.Background())
}

func (i MembershipAuthorityArgs) ToMembershipAuthorityPtrOutputWithContext(ctx context.Context) MembershipAuthorityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MembershipAuthorityOutput).ToMembershipAuthorityPtrOutputWithContext(ctx)
}

// MembershipAuthorityPtrInput is an input type that accepts MembershipAuthorityArgs, MembershipAuthorityPtr and MembershipAuthorityPtrOutput values.
// You can construct a concrete instance of `MembershipAuthorityPtrInput` via:
//
//          MembershipAuthorityArgs{...}
//
//  or:
//
//          nil
type MembershipAuthorityPtrInput interface {
	pulumi.Input

	ToMembershipAuthorityPtrOutput() MembershipAuthorityPtrOutput
	ToMembershipAuthorityPtrOutputWithContext(context.Context) MembershipAuthorityPtrOutput
}

type membershipAuthorityPtrType MembershipAuthorityArgs

func MembershipAuthorityPtr(v *MembershipAuthorityArgs) MembershipAuthorityPtrInput {
	return (*membershipAuthorityPtrType)(v)
}

func (*membershipAuthorityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MembershipAuthority)(nil)).Elem()
}

func (i *membershipAuthorityPtrType) ToMembershipAuthorityPtrOutput() MembershipAuthorityPtrOutput {
	return i.ToMembershipAuthorityPtrOutputWithContext(context.Background())
}

func (i *membershipAuthorityPtrType) ToMembershipAuthorityPtrOutputWithContext(ctx context.Context) MembershipAuthorityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MembershipAuthorityPtrOutput)
}

type MembershipAuthorityOutput struct{ *pulumi.OutputState }

func (MembershipAuthorityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MembershipAuthority)(nil)).Elem()
}

func (o MembershipAuthorityOutput) ToMembershipAuthorityOutput() MembershipAuthorityOutput {
	return o
}

func (o MembershipAuthorityOutput) ToMembershipAuthorityOutputWithContext(ctx context.Context) MembershipAuthorityOutput {
	return o
}

func (o MembershipAuthorityOutput) ToMembershipAuthorityPtrOutput() MembershipAuthorityPtrOutput {
	return o.ToMembershipAuthorityPtrOutputWithContext(context.Background())
}

func (o MembershipAuthorityOutput) ToMembershipAuthorityPtrOutputWithContext(ctx context.Context) MembershipAuthorityPtrOutput {
	return o.ApplyT(func(v MembershipAuthority) *MembershipAuthority {
		return &v
	}).(MembershipAuthorityPtrOutput)
}

// A JSON Web Token (JWT) issuer URI. `issuer` must start with `https://` and // be a valid
// with length <2000 characters.
func (o MembershipAuthorityOutput) Issuer() pulumi.StringOutput {
	return o.ApplyT(func(v MembershipAuthority) string { return v.Issuer }).(pulumi.StringOutput)
}

type MembershipAuthorityPtrOutput struct{ *pulumi.OutputState }

func (MembershipAuthorityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MembershipAuthority)(nil)).Elem()
}

func (o MembershipAuthorityPtrOutput) ToMembershipAuthorityPtrOutput() MembershipAuthorityPtrOutput {
	return o
}

func (o MembershipAuthorityPtrOutput) ToMembershipAuthorityPtrOutputWithContext(ctx context.Context) MembershipAuthorityPtrOutput {
	return o
}

func (o MembershipAuthorityPtrOutput) Elem() MembershipAuthorityOutput {
	return o.ApplyT(func(v *MembershipAuthority) MembershipAuthority { return *v }).(MembershipAuthorityOutput)
}

// A JSON Web Token (JWT) issuer URI. `issuer` must start with `https://` and // be a valid
// with length <2000 characters.
func (o MembershipAuthorityPtrOutput) Issuer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MembershipAuthority) *string {
		if v == nil {
			return nil
		}
		return &v.Issuer
	}).(pulumi.StringPtrOutput)
}

type MembershipEndpoint struct {
	// If this Membership is a Kubernetes API server hosted on GKE, this is a self link to its GCP resource.
	// Structure is documented below.
	GkeCluster *MembershipEndpointGkeCluster `pulumi:"gkeCluster"`
}

// MembershipEndpointInput is an input type that accepts MembershipEndpointArgs and MembershipEndpointOutput values.
// You can construct a concrete instance of `MembershipEndpointInput` via:
//
//          MembershipEndpointArgs{...}
type MembershipEndpointInput interface {
	pulumi.Input

	ToMembershipEndpointOutput() MembershipEndpointOutput
	ToMembershipEndpointOutputWithContext(context.Context) MembershipEndpointOutput
}

type MembershipEndpointArgs struct {
	// If this Membership is a Kubernetes API server hosted on GKE, this is a self link to its GCP resource.
	// Structure is documented below.
	GkeCluster MembershipEndpointGkeClusterPtrInput `pulumi:"gkeCluster"`
}

func (MembershipEndpointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MembershipEndpoint)(nil)).Elem()
}

func (i MembershipEndpointArgs) ToMembershipEndpointOutput() MembershipEndpointOutput {
	return i.ToMembershipEndpointOutputWithContext(context.Background())
}

func (i MembershipEndpointArgs) ToMembershipEndpointOutputWithContext(ctx context.Context) MembershipEndpointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MembershipEndpointOutput)
}

func (i MembershipEndpointArgs) ToMembershipEndpointPtrOutput() MembershipEndpointPtrOutput {
	return i.ToMembershipEndpointPtrOutputWithContext(context.Background())
}

func (i MembershipEndpointArgs) ToMembershipEndpointPtrOutputWithContext(ctx context.Context) MembershipEndpointPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MembershipEndpointOutput).ToMembershipEndpointPtrOutputWithContext(ctx)
}

// MembershipEndpointPtrInput is an input type that accepts MembershipEndpointArgs, MembershipEndpointPtr and MembershipEndpointPtrOutput values.
// You can construct a concrete instance of `MembershipEndpointPtrInput` via:
//
//          MembershipEndpointArgs{...}
//
//  or:
//
//          nil
type MembershipEndpointPtrInput interface {
	pulumi.Input

	ToMembershipEndpointPtrOutput() MembershipEndpointPtrOutput
	ToMembershipEndpointPtrOutputWithContext(context.Context) MembershipEndpointPtrOutput
}

type membershipEndpointPtrType MembershipEndpointArgs

func MembershipEndpointPtr(v *MembershipEndpointArgs) MembershipEndpointPtrInput {
	return (*membershipEndpointPtrType)(v)
}

func (*membershipEndpointPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MembershipEndpoint)(nil)).Elem()
}

func (i *membershipEndpointPtrType) ToMembershipEndpointPtrOutput() MembershipEndpointPtrOutput {
	return i.ToMembershipEndpointPtrOutputWithContext(context.Background())
}

func (i *membershipEndpointPtrType) ToMembershipEndpointPtrOutputWithContext(ctx context.Context) MembershipEndpointPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MembershipEndpointPtrOutput)
}

type MembershipEndpointOutput struct{ *pulumi.OutputState }

func (MembershipEndpointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MembershipEndpoint)(nil)).Elem()
}

func (o MembershipEndpointOutput) ToMembershipEndpointOutput() MembershipEndpointOutput {
	return o
}

func (o MembershipEndpointOutput) ToMembershipEndpointOutputWithContext(ctx context.Context) MembershipEndpointOutput {
	return o
}

func (o MembershipEndpointOutput) ToMembershipEndpointPtrOutput() MembershipEndpointPtrOutput {
	return o.ToMembershipEndpointPtrOutputWithContext(context.Background())
}

func (o MembershipEndpointOutput) ToMembershipEndpointPtrOutputWithContext(ctx context.Context) MembershipEndpointPtrOutput {
	return o.ApplyT(func(v MembershipEndpoint) *MembershipEndpoint {
		return &v
	}).(MembershipEndpointPtrOutput)
}

// If this Membership is a Kubernetes API server hosted on GKE, this is a self link to its GCP resource.
// Structure is documented below.
func (o MembershipEndpointOutput) GkeCluster() MembershipEndpointGkeClusterPtrOutput {
	return o.ApplyT(func(v MembershipEndpoint) *MembershipEndpointGkeCluster { return v.GkeCluster }).(MembershipEndpointGkeClusterPtrOutput)
}

type MembershipEndpointPtrOutput struct{ *pulumi.OutputState }

func (MembershipEndpointPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MembershipEndpoint)(nil)).Elem()
}

func (o MembershipEndpointPtrOutput) ToMembershipEndpointPtrOutput() MembershipEndpointPtrOutput {
	return o
}

func (o MembershipEndpointPtrOutput) ToMembershipEndpointPtrOutputWithContext(ctx context.Context) MembershipEndpointPtrOutput {
	return o
}

func (o MembershipEndpointPtrOutput) Elem() MembershipEndpointOutput {
	return o.ApplyT(func(v *MembershipEndpoint) MembershipEndpoint { return *v }).(MembershipEndpointOutput)
}

// If this Membership is a Kubernetes API server hosted on GKE, this is a self link to its GCP resource.
// Structure is documented below.
func (o MembershipEndpointPtrOutput) GkeCluster() MembershipEndpointGkeClusterPtrOutput {
	return o.ApplyT(func(v *MembershipEndpoint) *MembershipEndpointGkeCluster {
		if v == nil {
			return nil
		}
		return v.GkeCluster
	}).(MembershipEndpointGkeClusterPtrOutput)
}

type MembershipEndpointGkeCluster struct {
	ResourceLink string `pulumi:"resourceLink"`
}

// MembershipEndpointGkeClusterInput is an input type that accepts MembershipEndpointGkeClusterArgs and MembershipEndpointGkeClusterOutput values.
// You can construct a concrete instance of `MembershipEndpointGkeClusterInput` via:
//
//          MembershipEndpointGkeClusterArgs{...}
type MembershipEndpointGkeClusterInput interface {
	pulumi.Input

	ToMembershipEndpointGkeClusterOutput() MembershipEndpointGkeClusterOutput
	ToMembershipEndpointGkeClusterOutputWithContext(context.Context) MembershipEndpointGkeClusterOutput
}

type MembershipEndpointGkeClusterArgs struct {
	ResourceLink pulumi.StringInput `pulumi:"resourceLink"`
}

func (MembershipEndpointGkeClusterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MembershipEndpointGkeCluster)(nil)).Elem()
}

func (i MembershipEndpointGkeClusterArgs) ToMembershipEndpointGkeClusterOutput() MembershipEndpointGkeClusterOutput {
	return i.ToMembershipEndpointGkeClusterOutputWithContext(context.Background())
}

func (i MembershipEndpointGkeClusterArgs) ToMembershipEndpointGkeClusterOutputWithContext(ctx context.Context) MembershipEndpointGkeClusterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MembershipEndpointGkeClusterOutput)
}

func (i MembershipEndpointGkeClusterArgs) ToMembershipEndpointGkeClusterPtrOutput() MembershipEndpointGkeClusterPtrOutput {
	return i.ToMembershipEndpointGkeClusterPtrOutputWithContext(context.Background())
}

func (i MembershipEndpointGkeClusterArgs) ToMembershipEndpointGkeClusterPtrOutputWithContext(ctx context.Context) MembershipEndpointGkeClusterPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MembershipEndpointGkeClusterOutput).ToMembershipEndpointGkeClusterPtrOutputWithContext(ctx)
}

// MembershipEndpointGkeClusterPtrInput is an input type that accepts MembershipEndpointGkeClusterArgs, MembershipEndpointGkeClusterPtr and MembershipEndpointGkeClusterPtrOutput values.
// You can construct a concrete instance of `MembershipEndpointGkeClusterPtrInput` via:
//
//          MembershipEndpointGkeClusterArgs{...}
//
//  or:
//
//          nil
type MembershipEndpointGkeClusterPtrInput interface {
	pulumi.Input

	ToMembershipEndpointGkeClusterPtrOutput() MembershipEndpointGkeClusterPtrOutput
	ToMembershipEndpointGkeClusterPtrOutputWithContext(context.Context) MembershipEndpointGkeClusterPtrOutput
}

type membershipEndpointGkeClusterPtrType MembershipEndpointGkeClusterArgs

func MembershipEndpointGkeClusterPtr(v *MembershipEndpointGkeClusterArgs) MembershipEndpointGkeClusterPtrInput {
	return (*membershipEndpointGkeClusterPtrType)(v)
}

func (*membershipEndpointGkeClusterPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MembershipEndpointGkeCluster)(nil)).Elem()
}

func (i *membershipEndpointGkeClusterPtrType) ToMembershipEndpointGkeClusterPtrOutput() MembershipEndpointGkeClusterPtrOutput {
	return i.ToMembershipEndpointGkeClusterPtrOutputWithContext(context.Background())
}

func (i *membershipEndpointGkeClusterPtrType) ToMembershipEndpointGkeClusterPtrOutputWithContext(ctx context.Context) MembershipEndpointGkeClusterPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MembershipEndpointGkeClusterPtrOutput)
}

type MembershipEndpointGkeClusterOutput struct{ *pulumi.OutputState }

func (MembershipEndpointGkeClusterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MembershipEndpointGkeCluster)(nil)).Elem()
}

func (o MembershipEndpointGkeClusterOutput) ToMembershipEndpointGkeClusterOutput() MembershipEndpointGkeClusterOutput {
	return o
}

func (o MembershipEndpointGkeClusterOutput) ToMembershipEndpointGkeClusterOutputWithContext(ctx context.Context) MembershipEndpointGkeClusterOutput {
	return o
}

func (o MembershipEndpointGkeClusterOutput) ToMembershipEndpointGkeClusterPtrOutput() MembershipEndpointGkeClusterPtrOutput {
	return o.ToMembershipEndpointGkeClusterPtrOutputWithContext(context.Background())
}

func (o MembershipEndpointGkeClusterOutput) ToMembershipEndpointGkeClusterPtrOutputWithContext(ctx context.Context) MembershipEndpointGkeClusterPtrOutput {
	return o.ApplyT(func(v MembershipEndpointGkeCluster) *MembershipEndpointGkeCluster {
		return &v
	}).(MembershipEndpointGkeClusterPtrOutput)
}
func (o MembershipEndpointGkeClusterOutput) ResourceLink() pulumi.StringOutput {
	return o.ApplyT(func(v MembershipEndpointGkeCluster) string { return v.ResourceLink }).(pulumi.StringOutput)
}

type MembershipEndpointGkeClusterPtrOutput struct{ *pulumi.OutputState }

func (MembershipEndpointGkeClusterPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MembershipEndpointGkeCluster)(nil)).Elem()
}

func (o MembershipEndpointGkeClusterPtrOutput) ToMembershipEndpointGkeClusterPtrOutput() MembershipEndpointGkeClusterPtrOutput {
	return o
}

func (o MembershipEndpointGkeClusterPtrOutput) ToMembershipEndpointGkeClusterPtrOutputWithContext(ctx context.Context) MembershipEndpointGkeClusterPtrOutput {
	return o
}

func (o MembershipEndpointGkeClusterPtrOutput) Elem() MembershipEndpointGkeClusterOutput {
	return o.ApplyT(func(v *MembershipEndpointGkeCluster) MembershipEndpointGkeCluster { return *v }).(MembershipEndpointGkeClusterOutput)
}

func (o MembershipEndpointGkeClusterPtrOutput) ResourceLink() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MembershipEndpointGkeCluster) *string {
		if v == nil {
			return nil
		}
		return &v.ResourceLink
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(MembershipAuthorityOutput{})
	pulumi.RegisterOutputType(MembershipAuthorityPtrOutput{})
	pulumi.RegisterOutputType(MembershipEndpointOutput{})
	pulumi.RegisterOutputType(MembershipEndpointPtrOutput{})
	pulumi.RegisterOutputType(MembershipEndpointGkeClusterOutput{})
	pulumi.RegisterOutputType(MembershipEndpointGkeClusterPtrOutput{})
}