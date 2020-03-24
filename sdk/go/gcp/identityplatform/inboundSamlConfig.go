// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package identityplatform

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Inbound SAML configuration for a Identity Toolkit project.
//
// You must enable the
// [Google Identity Platform](https://console.cloud.google.com/marketplace/details/google-cloud-platform/customer-identity) in
// the marketplace prior to using this resource.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/identity_platform_inbound_saml_config.html.markdown.
type InboundSamlConfig struct {
	pulumi.CustomResourceState

	// Human friendly display name.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// If this config allows users to sign in with the provider.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// SAML IdP configuration when the project acts as the relying party
	IdpConfig InboundSamlConfigIdpConfigOutput `pulumi:"idpConfig"`
	// The name of the InboundSamlConfig resource. Must start with 'saml.' and can only have alphanumeric characters, hyphens,
	// underscores or periods. The part after 'saml.' must also start with a lowercase letter, end with an alphanumeric
	// character, and have at least 2 characters.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// SAML SP (Service Provider) configuration when the project acts as the relying party to receive and accept an
	// authentication assertion issued by a SAML identity provider.
	SpConfig InboundSamlConfigSpConfigOutput `pulumi:"spConfig"`
}

// NewInboundSamlConfig registers a new resource with the given unique name, arguments, and options.
func NewInboundSamlConfig(ctx *pulumi.Context,
	name string, args *InboundSamlConfigArgs, opts ...pulumi.ResourceOption) (*InboundSamlConfig, error) {
	if args == nil || args.DisplayName == nil {
		return nil, errors.New("missing required argument 'DisplayName'")
	}
	if args == nil || args.IdpConfig == nil {
		return nil, errors.New("missing required argument 'IdpConfig'")
	}
	if args == nil || args.SpConfig == nil {
		return nil, errors.New("missing required argument 'SpConfig'")
	}
	if args == nil {
		args = &InboundSamlConfigArgs{}
	}
	var resource InboundSamlConfig
	err := ctx.RegisterResource("gcp:identityplatform/inboundSamlConfig:InboundSamlConfig", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInboundSamlConfig gets an existing InboundSamlConfig resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInboundSamlConfig(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InboundSamlConfigState, opts ...pulumi.ResourceOption) (*InboundSamlConfig, error) {
	var resource InboundSamlConfig
	err := ctx.ReadResource("gcp:identityplatform/inboundSamlConfig:InboundSamlConfig", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering InboundSamlConfig resources.
type inboundSamlConfigState struct {
	// Human friendly display name.
	DisplayName *string `pulumi:"displayName"`
	// If this config allows users to sign in with the provider.
	Enabled *bool `pulumi:"enabled"`
	// SAML IdP configuration when the project acts as the relying party
	IdpConfig *InboundSamlConfigIdpConfig `pulumi:"idpConfig"`
	// The name of the InboundSamlConfig resource. Must start with 'saml.' and can only have alphanumeric characters, hyphens,
	// underscores or periods. The part after 'saml.' must also start with a lowercase letter, end with an alphanumeric
	// character, and have at least 2 characters.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// SAML SP (Service Provider) configuration when the project acts as the relying party to receive and accept an
	// authentication assertion issued by a SAML identity provider.
	SpConfig *InboundSamlConfigSpConfig `pulumi:"spConfig"`
}

type InboundSamlConfigState struct {
	// Human friendly display name.
	DisplayName pulumi.StringPtrInput
	// If this config allows users to sign in with the provider.
	Enabled pulumi.BoolPtrInput
	// SAML IdP configuration when the project acts as the relying party
	IdpConfig InboundSamlConfigIdpConfigPtrInput
	// The name of the InboundSamlConfig resource. Must start with 'saml.' and can only have alphanumeric characters, hyphens,
	// underscores or periods. The part after 'saml.' must also start with a lowercase letter, end with an alphanumeric
	// character, and have at least 2 characters.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// SAML SP (Service Provider) configuration when the project acts as the relying party to receive and accept an
	// authentication assertion issued by a SAML identity provider.
	SpConfig InboundSamlConfigSpConfigPtrInput
}

func (InboundSamlConfigState) ElementType() reflect.Type {
	return reflect.TypeOf((*inboundSamlConfigState)(nil)).Elem()
}

type inboundSamlConfigArgs struct {
	// Human friendly display name.
	DisplayName string `pulumi:"displayName"`
	// If this config allows users to sign in with the provider.
	Enabled *bool `pulumi:"enabled"`
	// SAML IdP configuration when the project acts as the relying party
	IdpConfig InboundSamlConfigIdpConfig `pulumi:"idpConfig"`
	// The name of the InboundSamlConfig resource. Must start with 'saml.' and can only have alphanumeric characters, hyphens,
	// underscores or periods. The part after 'saml.' must also start with a lowercase letter, end with an alphanumeric
	// character, and have at least 2 characters.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// SAML SP (Service Provider) configuration when the project acts as the relying party to receive and accept an
	// authentication assertion issued by a SAML identity provider.
	SpConfig InboundSamlConfigSpConfig `pulumi:"spConfig"`
}

// The set of arguments for constructing a InboundSamlConfig resource.
type InboundSamlConfigArgs struct {
	// Human friendly display name.
	DisplayName pulumi.StringInput
	// If this config allows users to sign in with the provider.
	Enabled pulumi.BoolPtrInput
	// SAML IdP configuration when the project acts as the relying party
	IdpConfig InboundSamlConfigIdpConfigInput
	// The name of the InboundSamlConfig resource. Must start with 'saml.' and can only have alphanumeric characters, hyphens,
	// underscores or periods. The part after 'saml.' must also start with a lowercase letter, end with an alphanumeric
	// character, and have at least 2 characters.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// SAML SP (Service Provider) configuration when the project acts as the relying party to receive and accept an
	// authentication assertion issued by a SAML identity provider.
	SpConfig InboundSamlConfigSpConfigInput
}

func (InboundSamlConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*inboundSamlConfigArgs)(nil)).Elem()
}

