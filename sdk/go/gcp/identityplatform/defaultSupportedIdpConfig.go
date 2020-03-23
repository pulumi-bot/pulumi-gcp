// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package identityplatform

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Configurations options for authenticating with a the standard set of Identity Toolkit-trusted IDPs.
//
// You must enable the
// [Google Identity Platform](https://console.cloud.google.com/marketplace/details/google-cloud-platform/customer-identity) in
// the marketplace prior to using this resource.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/identity_platform_default_supported_idp_config.html.markdown.
type DefaultSupportedIdpConfig struct {
	pulumi.CustomResourceState

	// OAuth client ID
	ClientId pulumi.StringOutput `pulumi:"clientId"`
	// OAuth client secret
	ClientSecret pulumi.StringOutput `pulumi:"clientSecret"`
	// If this IDP allows the user to sign in
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// ID of the IDP. Possible values include: * 'apple.com' * 'facebook.com' * 'gc.apple.com' * 'github.com' * 'google.com' *
	// 'linkedin.com' * 'microsoft.com' * 'playgames.google.com' * 'twitter.com' * 'yahoo.com'
	IdpId pulumi.StringOutput `pulumi:"idpId"`
	// The name of the DefaultSupportedIdpConfig resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
}

// NewDefaultSupportedIdpConfig registers a new resource with the given unique name, arguments, and options.
func NewDefaultSupportedIdpConfig(ctx *pulumi.Context,
	name string, args *DefaultSupportedIdpConfigArgs, opts ...pulumi.ResourceOption) (*DefaultSupportedIdpConfig, error) {
	if args == nil || args.ClientId == nil {
		return nil, errors.New("missing required argument 'ClientId'")
	}
	if args == nil || args.ClientSecret == nil {
		return nil, errors.New("missing required argument 'ClientSecret'")
	}
	if args == nil || args.IdpId == nil {
		return nil, errors.New("missing required argument 'IdpId'")
	}
	if args == nil {
		args = &DefaultSupportedIdpConfigArgs{}
	}
	var resource DefaultSupportedIdpConfig
	err := ctx.RegisterResource("gcp:identityplatform/defaultSupportedIdpConfig:DefaultSupportedIdpConfig", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDefaultSupportedIdpConfig gets an existing DefaultSupportedIdpConfig resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDefaultSupportedIdpConfig(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DefaultSupportedIdpConfigState, opts ...pulumi.ResourceOption) (*DefaultSupportedIdpConfig, error) {
	var resource DefaultSupportedIdpConfig
	err := ctx.ReadResource("gcp:identityplatform/defaultSupportedIdpConfig:DefaultSupportedIdpConfig", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DefaultSupportedIdpConfig resources.
type defaultSupportedIdpConfigState struct {
	// OAuth client ID
	ClientId *string `pulumi:"clientId"`
	// OAuth client secret
	ClientSecret *string `pulumi:"clientSecret"`
	// If this IDP allows the user to sign in
	Enabled *bool `pulumi:"enabled"`
	// ID of the IDP. Possible values include: * 'apple.com' * 'facebook.com' * 'gc.apple.com' * 'github.com' * 'google.com' *
	// 'linkedin.com' * 'microsoft.com' * 'playgames.google.com' * 'twitter.com' * 'yahoo.com'
	IdpId *string `pulumi:"idpId"`
	// The name of the DefaultSupportedIdpConfig resource
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
}

type DefaultSupportedIdpConfigState struct {
	// OAuth client ID
	ClientId pulumi.StringPtrInput
	// OAuth client secret
	ClientSecret pulumi.StringPtrInput
	// If this IDP allows the user to sign in
	Enabled pulumi.BoolPtrInput
	// ID of the IDP. Possible values include: * 'apple.com' * 'facebook.com' * 'gc.apple.com' * 'github.com' * 'google.com' *
	// 'linkedin.com' * 'microsoft.com' * 'playgames.google.com' * 'twitter.com' * 'yahoo.com'
	IdpId pulumi.StringPtrInput
	// The name of the DefaultSupportedIdpConfig resource
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
}

func (DefaultSupportedIdpConfigState) ElementType() reflect.Type {
	return reflect.TypeOf((*defaultSupportedIdpConfigState)(nil)).Elem()
}

type defaultSupportedIdpConfigArgs struct {
	// OAuth client ID
	ClientId string `pulumi:"clientId"`
	// OAuth client secret
	ClientSecret string `pulumi:"clientSecret"`
	// If this IDP allows the user to sign in
	Enabled *bool `pulumi:"enabled"`
	// ID of the IDP. Possible values include: * 'apple.com' * 'facebook.com' * 'gc.apple.com' * 'github.com' * 'google.com' *
	// 'linkedin.com' * 'microsoft.com' * 'playgames.google.com' * 'twitter.com' * 'yahoo.com'
	IdpId string `pulumi:"idpId"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
}

// The set of arguments for constructing a DefaultSupportedIdpConfig resource.
type DefaultSupportedIdpConfigArgs struct {
	// OAuth client ID
	ClientId pulumi.StringInput
	// OAuth client secret
	ClientSecret pulumi.StringInput
	// If this IDP allows the user to sign in
	Enabled pulumi.BoolPtrInput
	// ID of the IDP. Possible values include: * 'apple.com' * 'facebook.com' * 'gc.apple.com' * 'github.com' * 'google.com' *
	// 'linkedin.com' * 'microsoft.com' * 'playgames.google.com' * 'twitter.com' * 'yahoo.com'
	IdpId pulumi.StringInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
}

func (DefaultSupportedIdpConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*defaultSupportedIdpConfigArgs)(nil)).Elem()
}

