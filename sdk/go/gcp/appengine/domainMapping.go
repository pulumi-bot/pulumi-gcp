// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package appengine

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// A domain serving an App Engine application.
//
// To get more information about DomainMapping, see:
//
// * [API documentation](https://cloud.google.com/appengine/docs/admin-api/reference/rest/v1/apps.domainMappings)
// * How-to Guides
//     * [Official Documentation](https://cloud.google.com/appengine/docs/standard/python/mapping-custom-domains)
//
// ## Example Usage
type DomainMapping struct {
	pulumi.CustomResourceState

	// Relative name of the domain serving the application. Example: example.com.
	DomainName pulumi.StringOutput `pulumi:"domainName"`
	// Full path to the DomainMapping resource in the API. Example: apps/myapp/domainMapping/example.com.
	Name pulumi.StringOutput `pulumi:"name"`
	// Whether the domain creation should override any existing mappings for this domain.
	// By default, overrides are rejected.
	// Default value is `STRICT`.
	// Possible values are `STRICT` and `OVERRIDE`.
	OverrideStrategy pulumi.StringPtrOutput `pulumi:"overrideStrategy"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The resource records required to configure this domain mapping. These records must be added to the domain's DNS
	// configuration in order to serve the application via this domain mapping.
	ResourceRecords DomainMappingResourceRecordArrayOutput `pulumi:"resourceRecords"`
	// SSL configuration for this domain. If unconfigured, this domain will not serve with SSL.
	// Structure is documented below.
	SslSettings DomainMappingSslSettingsPtrOutput `pulumi:"sslSettings"`
}

// NewDomainMapping registers a new resource with the given unique name, arguments, and options.
func NewDomainMapping(ctx *pulumi.Context,
	name string, args *DomainMappingArgs, opts ...pulumi.ResourceOption) (*DomainMapping, error) {
	if args == nil || args.DomainName == nil {
		return nil, errors.New("missing required argument 'DomainName'")
	}
	if args == nil {
		args = &DomainMappingArgs{}
	}
	var resource DomainMapping
	err := ctx.RegisterResource("gcp:appengine/domainMapping:DomainMapping", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDomainMapping gets an existing DomainMapping resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDomainMapping(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DomainMappingState, opts ...pulumi.ResourceOption) (*DomainMapping, error) {
	var resource DomainMapping
	err := ctx.ReadResource("gcp:appengine/domainMapping:DomainMapping", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DomainMapping resources.
type domainMappingState struct {
	// Relative name of the domain serving the application. Example: example.com.
	DomainName *string `pulumi:"domainName"`
	// Full path to the DomainMapping resource in the API. Example: apps/myapp/domainMapping/example.com.
	Name *string `pulumi:"name"`
	// Whether the domain creation should override any existing mappings for this domain.
	// By default, overrides are rejected.
	// Default value is `STRICT`.
	// Possible values are `STRICT` and `OVERRIDE`.
	OverrideStrategy *string `pulumi:"overrideStrategy"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The resource records required to configure this domain mapping. These records must be added to the domain's DNS
	// configuration in order to serve the application via this domain mapping.
	ResourceRecords []DomainMappingResourceRecord `pulumi:"resourceRecords"`
	// SSL configuration for this domain. If unconfigured, this domain will not serve with SSL.
	// Structure is documented below.
	SslSettings *DomainMappingSslSettings `pulumi:"sslSettings"`
}

type DomainMappingState struct {
	// Relative name of the domain serving the application. Example: example.com.
	DomainName pulumi.StringPtrInput
	// Full path to the DomainMapping resource in the API. Example: apps/myapp/domainMapping/example.com.
	Name pulumi.StringPtrInput
	// Whether the domain creation should override any existing mappings for this domain.
	// By default, overrides are rejected.
	// Default value is `STRICT`.
	// Possible values are `STRICT` and `OVERRIDE`.
	OverrideStrategy pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The resource records required to configure this domain mapping. These records must be added to the domain's DNS
	// configuration in order to serve the application via this domain mapping.
	ResourceRecords DomainMappingResourceRecordArrayInput
	// SSL configuration for this domain. If unconfigured, this domain will not serve with SSL.
	// Structure is documented below.
	SslSettings DomainMappingSslSettingsPtrInput
}

func (DomainMappingState) ElementType() reflect.Type {
	return reflect.TypeOf((*domainMappingState)(nil)).Elem()
}

type domainMappingArgs struct {
	// Relative name of the domain serving the application. Example: example.com.
	DomainName string `pulumi:"domainName"`
	// Whether the domain creation should override any existing mappings for this domain.
	// By default, overrides are rejected.
	// Default value is `STRICT`.
	// Possible values are `STRICT` and `OVERRIDE`.
	OverrideStrategy *string `pulumi:"overrideStrategy"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// SSL configuration for this domain. If unconfigured, this domain will not serve with SSL.
	// Structure is documented below.
	SslSettings *DomainMappingSslSettings `pulumi:"sslSettings"`
}

// The set of arguments for constructing a DomainMapping resource.
type DomainMappingArgs struct {
	// Relative name of the domain serving the application. Example: example.com.
	DomainName pulumi.StringInput
	// Whether the domain creation should override any existing mappings for this domain.
	// By default, overrides are rejected.
	// Default value is `STRICT`.
	// Possible values are `STRICT` and `OVERRIDE`.
	OverrideStrategy pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// SSL configuration for this domain. If unconfigured, this domain will not serve with SSL.
	// Structure is documented below.
	SslSettings DomainMappingSslSettingsPtrInput
}

func (DomainMappingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*domainMappingArgs)(nil)).Elem()
}

type DomainMappingInput interface {
	pulumi.Input

	ToDomainMappingOutput() DomainMappingOutput
	ToDomainMappingOutputWithContext(ctx context.Context) DomainMappingOutput
}

func (DomainMapping) ElementType() reflect.Type {
	return reflect.TypeOf((*DomainMapping)(nil)).Elem()
}

func (i DomainMapping) ToDomainMappingOutput() DomainMappingOutput {
	return i.ToDomainMappingOutputWithContext(context.Background())
}

func (i DomainMapping) ToDomainMappingOutputWithContext(ctx context.Context) DomainMappingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainMappingOutput)
}

type DomainMappingOutput struct {
	*pulumi.OutputState
}

func (DomainMappingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DomainMappingOutput)(nil)).Elem()
}

func (o DomainMappingOutput) ToDomainMappingOutput() DomainMappingOutput {
	return o
}

func (o DomainMappingOutput) ToDomainMappingOutputWithContext(ctx context.Context) DomainMappingOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(DomainMappingOutput{})
}
