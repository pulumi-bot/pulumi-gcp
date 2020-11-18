// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package servicedirectory

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// A container for `services`. Namespaces allow administrators to group services
// together and define permissions for a collection of services.
//
// To get more information about Namespace, see:
//
// * [API documentation](https://cloud.google.com/service-directory/docs/reference/rest/v1beta1/projects.locations.namespaces)
// * How-to Guides
//     * [Configuring a namespace](https://cloud.google.com/service-directory/docs/configuring-service-directory#configuring_a_namespace)
//
// ## Example Usage
type Namespace struct {
	pulumi.CustomResourceState

	// Resource labels associated with this Namespace. No more than 64 user
	// labels can be associated with a given resource. Label keys and values can
	// be no longer than 63 characters.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// The location for the Namespace.
	// A full list of valid locations can be found by running
	// `gcloud beta service-directory locations list`.
	Location pulumi.StringOutput `pulumi:"location"`
	// The resource name for the namespace in the format 'projects/*/locations/*/namespaces/*'.
	Name pulumi.StringOutput `pulumi:"name"`
	// The Resource ID must be 1-63 characters long, including digits,
	// lowercase letters or the hyphen character.
	NamespaceId pulumi.StringOutput `pulumi:"namespaceId"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
}

// NewNamespace registers a new resource with the given unique name, arguments, and options.
func NewNamespace(ctx *pulumi.Context,
	name string, args *NamespaceArgs, opts ...pulumi.ResourceOption) (*Namespace, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Location == nil {
		return nil, errors.New("invalid value for required argument 'Location'")
	}
	if args.NamespaceId == nil {
		return nil, errors.New("invalid value for required argument 'NamespaceId'")
	}
	var resource Namespace
	err := ctx.RegisterResource("gcp:servicedirectory/namespace:Namespace", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNamespace gets an existing Namespace resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNamespace(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NamespaceState, opts ...pulumi.ResourceOption) (*Namespace, error) {
	var resource Namespace
	err := ctx.ReadResource("gcp:servicedirectory/namespace:Namespace", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Namespace resources.
type namespaceState struct {
	// Resource labels associated with this Namespace. No more than 64 user
	// labels can be associated with a given resource. Label keys and values can
	// be no longer than 63 characters.
	Labels map[string]string `pulumi:"labels"`
	// The location for the Namespace.
	// A full list of valid locations can be found by running
	// `gcloud beta service-directory locations list`.
	Location *string `pulumi:"location"`
	// The resource name for the namespace in the format 'projects/*/locations/*/namespaces/*'.
	Name *string `pulumi:"name"`
	// The Resource ID must be 1-63 characters long, including digits,
	// lowercase letters or the hyphen character.
	NamespaceId *string `pulumi:"namespaceId"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
}

type NamespaceState struct {
	// Resource labels associated with this Namespace. No more than 64 user
	// labels can be associated with a given resource. Label keys and values can
	// be no longer than 63 characters.
	Labels pulumi.StringMapInput
	// The location for the Namespace.
	// A full list of valid locations can be found by running
	// `gcloud beta service-directory locations list`.
	Location pulumi.StringPtrInput
	// The resource name for the namespace in the format 'projects/*/locations/*/namespaces/*'.
	Name pulumi.StringPtrInput
	// The Resource ID must be 1-63 characters long, including digits,
	// lowercase letters or the hyphen character.
	NamespaceId pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
}

func (NamespaceState) ElementType() reflect.Type {
	return reflect.TypeOf((*namespaceState)(nil)).Elem()
}

type namespaceArgs struct {
	// Resource labels associated with this Namespace. No more than 64 user
	// labels can be associated with a given resource. Label keys and values can
	// be no longer than 63 characters.
	Labels map[string]string `pulumi:"labels"`
	// The location for the Namespace.
	// A full list of valid locations can be found by running
	// `gcloud beta service-directory locations list`.
	Location string `pulumi:"location"`
	// The Resource ID must be 1-63 characters long, including digits,
	// lowercase letters or the hyphen character.
	NamespaceId string `pulumi:"namespaceId"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
}

// The set of arguments for constructing a Namespace resource.
type NamespaceArgs struct {
	// Resource labels associated with this Namespace. No more than 64 user
	// labels can be associated with a given resource. Label keys and values can
	// be no longer than 63 characters.
	Labels pulumi.StringMapInput
	// The location for the Namespace.
	// A full list of valid locations can be found by running
	// `gcloud beta service-directory locations list`.
	Location pulumi.StringInput
	// The Resource ID must be 1-63 characters long, including digits,
	// lowercase letters or the hyphen character.
	NamespaceId pulumi.StringInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
}

func (NamespaceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*namespaceArgs)(nil)).Elem()
}

type NamespaceInput interface {
	pulumi.Input

	ToNamespaceOutput() NamespaceOutput
	ToNamespaceOutputWithContext(ctx context.Context) NamespaceOutput
}

func (Namespace) ElementType() reflect.Type {
	return reflect.TypeOf((*Namespace)(nil)).Elem()
}

func (i Namespace) ToNamespaceOutput() NamespaceOutput {
	return i.ToNamespaceOutputWithContext(context.Background())
}

func (i Namespace) ToNamespaceOutputWithContext(ctx context.Context) NamespaceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NamespaceOutput)
}

type NamespaceOutput struct {
	*pulumi.OutputState
}

func (NamespaceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NamespaceOutput)(nil)).Elem()
}

func (o NamespaceOutput) ToNamespaceOutput() NamespaceOutput {
	return o
}

func (o NamespaceOutput) ToNamespaceOutputWithContext(ctx context.Context) NamespaceOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(NamespaceOutput{})
}
