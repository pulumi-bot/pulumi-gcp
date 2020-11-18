// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package osconfig

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// An OS Config resource representing a guest configuration policy. These policies represent
// the desired state for VM instance guest environments including packages to install or remove,
// package repository configurations, and software to install.
//
// To get more information about GuestPolicies, see:
//
// * [API documentation](https://cloud.google.com/compute/docs/osconfig/rest)
// * How-to Guides
//     * [Official Documentation](https://cloud.google.com/compute/docs/os-config-management)
//
// ## Example Usage
type GuestPolicies struct {
	pulumi.CustomResourceState

	// Specifies the VM instances that are assigned to this policy. This allows you to target sets
	// or groups of VM instances by different parameters such as labels, names, OS, or zones.
	// If left empty, all VM instances underneath this policy are targeted.
	// At the same level in the resource hierarchy (that is within a project), the service prevents
	// the creation of multiple policies that conflict with each other.
	// For more information, see how the service
	// [handles assignment conflicts](https://cloud.google.com/compute/docs/os-config-management/create-guest-policy#handle-conflicts).
	// Structure is documented below.
	Assignment GuestPoliciesAssignmentOutput `pulumi:"assignment"`
	// Time this guest policy was created. A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example:
	// "2014-10-02T15:01:23.045123456Z".
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Description of the guest policy. Length of the description is limited to 1024 characters.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The etag for this guest policy. If this is provided on update, it must match the server's etag.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The logical name of the guest policy in the project with the following restrictions:
	// * Must contain only lowercase letters, numbers, and hyphens.
	// * Must start with a letter.
	// * Must be between 1-63 characters.
	// * Must end with a number or a letter.
	// * Must be unique within the project.
	GuestPolicyId pulumi.StringOutput `pulumi:"guestPolicyId"`
	// Unique identifier for the recipe. Only one recipe with a given name is installed on an instance.
	// Names are also used to identify resources which helps to determine whether guest policies have conflicts.
	// This means that requests to create multiple recipes with the same name and version are rejected since they
	// could potentially have conflicting assignments.
	Name pulumi.StringOutput `pulumi:"name"`
	// A list of package repositories to configure on the VM instance.
	// This is done before any other configs are applied so they can use these repos.
	// Package repositories are only configured if the corresponding package manager(s) are available.
	// Structure is documented below.
	PackageRepositories GuestPoliciesPackageRepositoryArrayOutput `pulumi:"packageRepositories"`
	// The software packages to be managed by this policy.
	// Structure is documented below.
	Packages GuestPoliciesPackageArrayOutput `pulumi:"packages"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// A list of Recipes to install on the VM instance.
	// Structure is documented below.
	Recipes GuestPoliciesRecipeArrayOutput `pulumi:"recipes"`
	// Last time this guest policy was updated. A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example:
	// "2014-10-02T15:01:23.045123456Z".
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewGuestPolicies registers a new resource with the given unique name, arguments, and options.
func NewGuestPolicies(ctx *pulumi.Context,
	name string, args *GuestPoliciesArgs, opts ...pulumi.ResourceOption) (*GuestPolicies, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Assignment == nil {
		return nil, errors.New("invalid value for required argument 'Assignment'")
	}
	if args.GuestPolicyId == nil {
		return nil, errors.New("invalid value for required argument 'GuestPolicyId'")
	}
	var resource GuestPolicies
	err := ctx.RegisterResource("gcp:osconfig/guestPolicies:GuestPolicies", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGuestPolicies gets an existing GuestPolicies resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGuestPolicies(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GuestPoliciesState, opts ...pulumi.ResourceOption) (*GuestPolicies, error) {
	var resource GuestPolicies
	err := ctx.ReadResource("gcp:osconfig/guestPolicies:GuestPolicies", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GuestPolicies resources.
type guestPoliciesState struct {
	// Specifies the VM instances that are assigned to this policy. This allows you to target sets
	// or groups of VM instances by different parameters such as labels, names, OS, or zones.
	// If left empty, all VM instances underneath this policy are targeted.
	// At the same level in the resource hierarchy (that is within a project), the service prevents
	// the creation of multiple policies that conflict with each other.
	// For more information, see how the service
	// [handles assignment conflicts](https://cloud.google.com/compute/docs/os-config-management/create-guest-policy#handle-conflicts).
	// Structure is documented below.
	Assignment *GuestPoliciesAssignment `pulumi:"assignment"`
	// Time this guest policy was created. A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example:
	// "2014-10-02T15:01:23.045123456Z".
	CreateTime *string `pulumi:"createTime"`
	// Description of the guest policy. Length of the description is limited to 1024 characters.
	Description *string `pulumi:"description"`
	// The etag for this guest policy. If this is provided on update, it must match the server's etag.
	Etag *string `pulumi:"etag"`
	// The logical name of the guest policy in the project with the following restrictions:
	// * Must contain only lowercase letters, numbers, and hyphens.
	// * Must start with a letter.
	// * Must be between 1-63 characters.
	// * Must end with a number or a letter.
	// * Must be unique within the project.
	GuestPolicyId *string `pulumi:"guestPolicyId"`
	// Unique identifier for the recipe. Only one recipe with a given name is installed on an instance.
	// Names are also used to identify resources which helps to determine whether guest policies have conflicts.
	// This means that requests to create multiple recipes with the same name and version are rejected since they
	// could potentially have conflicting assignments.
	Name *string `pulumi:"name"`
	// A list of package repositories to configure on the VM instance.
	// This is done before any other configs are applied so they can use these repos.
	// Package repositories are only configured if the corresponding package manager(s) are available.
	// Structure is documented below.
	PackageRepositories []GuestPoliciesPackageRepository `pulumi:"packageRepositories"`
	// The software packages to be managed by this policy.
	// Structure is documented below.
	Packages []GuestPoliciesPackage `pulumi:"packages"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// A list of Recipes to install on the VM instance.
	// Structure is documented below.
	Recipes []GuestPoliciesRecipe `pulumi:"recipes"`
	// Last time this guest policy was updated. A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example:
	// "2014-10-02T15:01:23.045123456Z".
	UpdateTime *string `pulumi:"updateTime"`
}

type GuestPoliciesState struct {
	// Specifies the VM instances that are assigned to this policy. This allows you to target sets
	// or groups of VM instances by different parameters such as labels, names, OS, or zones.
	// If left empty, all VM instances underneath this policy are targeted.
	// At the same level in the resource hierarchy (that is within a project), the service prevents
	// the creation of multiple policies that conflict with each other.
	// For more information, see how the service
	// [handles assignment conflicts](https://cloud.google.com/compute/docs/os-config-management/create-guest-policy#handle-conflicts).
	// Structure is documented below.
	Assignment GuestPoliciesAssignmentPtrInput
	// Time this guest policy was created. A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example:
	// "2014-10-02T15:01:23.045123456Z".
	CreateTime pulumi.StringPtrInput
	// Description of the guest policy. Length of the description is limited to 1024 characters.
	Description pulumi.StringPtrInput
	// The etag for this guest policy. If this is provided on update, it must match the server's etag.
	Etag pulumi.StringPtrInput
	// The logical name of the guest policy in the project with the following restrictions:
	// * Must contain only lowercase letters, numbers, and hyphens.
	// * Must start with a letter.
	// * Must be between 1-63 characters.
	// * Must end with a number or a letter.
	// * Must be unique within the project.
	GuestPolicyId pulumi.StringPtrInput
	// Unique identifier for the recipe. Only one recipe with a given name is installed on an instance.
	// Names are also used to identify resources which helps to determine whether guest policies have conflicts.
	// This means that requests to create multiple recipes with the same name and version are rejected since they
	// could potentially have conflicting assignments.
	Name pulumi.StringPtrInput
	// A list of package repositories to configure on the VM instance.
	// This is done before any other configs are applied so they can use these repos.
	// Package repositories are only configured if the corresponding package manager(s) are available.
	// Structure is documented below.
	PackageRepositories GuestPoliciesPackageRepositoryArrayInput
	// The software packages to be managed by this policy.
	// Structure is documented below.
	Packages GuestPoliciesPackageArrayInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// A list of Recipes to install on the VM instance.
	// Structure is documented below.
	Recipes GuestPoliciesRecipeArrayInput
	// Last time this guest policy was updated. A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example:
	// "2014-10-02T15:01:23.045123456Z".
	UpdateTime pulumi.StringPtrInput
}

func (GuestPoliciesState) ElementType() reflect.Type {
	return reflect.TypeOf((*guestPoliciesState)(nil)).Elem()
}

type guestPoliciesArgs struct {
	// Specifies the VM instances that are assigned to this policy. This allows you to target sets
	// or groups of VM instances by different parameters such as labels, names, OS, or zones.
	// If left empty, all VM instances underneath this policy are targeted.
	// At the same level in the resource hierarchy (that is within a project), the service prevents
	// the creation of multiple policies that conflict with each other.
	// For more information, see how the service
	// [handles assignment conflicts](https://cloud.google.com/compute/docs/os-config-management/create-guest-policy#handle-conflicts).
	// Structure is documented below.
	Assignment GuestPoliciesAssignment `pulumi:"assignment"`
	// Description of the guest policy. Length of the description is limited to 1024 characters.
	Description *string `pulumi:"description"`
	// The etag for this guest policy. If this is provided on update, it must match the server's etag.
	Etag *string `pulumi:"etag"`
	// The logical name of the guest policy in the project with the following restrictions:
	// * Must contain only lowercase letters, numbers, and hyphens.
	// * Must start with a letter.
	// * Must be between 1-63 characters.
	// * Must end with a number or a letter.
	// * Must be unique within the project.
	GuestPolicyId string `pulumi:"guestPolicyId"`
	// A list of package repositories to configure on the VM instance.
	// This is done before any other configs are applied so they can use these repos.
	// Package repositories are only configured if the corresponding package manager(s) are available.
	// Structure is documented below.
	PackageRepositories []GuestPoliciesPackageRepository `pulumi:"packageRepositories"`
	// The software packages to be managed by this policy.
	// Structure is documented below.
	Packages []GuestPoliciesPackage `pulumi:"packages"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// A list of Recipes to install on the VM instance.
	// Structure is documented below.
	Recipes []GuestPoliciesRecipe `pulumi:"recipes"`
}

// The set of arguments for constructing a GuestPolicies resource.
type GuestPoliciesArgs struct {
	// Specifies the VM instances that are assigned to this policy. This allows you to target sets
	// or groups of VM instances by different parameters such as labels, names, OS, or zones.
	// If left empty, all VM instances underneath this policy are targeted.
	// At the same level in the resource hierarchy (that is within a project), the service prevents
	// the creation of multiple policies that conflict with each other.
	// For more information, see how the service
	// [handles assignment conflicts](https://cloud.google.com/compute/docs/os-config-management/create-guest-policy#handle-conflicts).
	// Structure is documented below.
	Assignment GuestPoliciesAssignmentInput
	// Description of the guest policy. Length of the description is limited to 1024 characters.
	Description pulumi.StringPtrInput
	// The etag for this guest policy. If this is provided on update, it must match the server's etag.
	Etag pulumi.StringPtrInput
	// The logical name of the guest policy in the project with the following restrictions:
	// * Must contain only lowercase letters, numbers, and hyphens.
	// * Must start with a letter.
	// * Must be between 1-63 characters.
	// * Must end with a number or a letter.
	// * Must be unique within the project.
	GuestPolicyId pulumi.StringInput
	// A list of package repositories to configure on the VM instance.
	// This is done before any other configs are applied so they can use these repos.
	// Package repositories are only configured if the corresponding package manager(s) are available.
	// Structure is documented below.
	PackageRepositories GuestPoliciesPackageRepositoryArrayInput
	// The software packages to be managed by this policy.
	// Structure is documented below.
	Packages GuestPoliciesPackageArrayInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// A list of Recipes to install on the VM instance.
	// Structure is documented below.
	Recipes GuestPoliciesRecipeArrayInput
}

func (GuestPoliciesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*guestPoliciesArgs)(nil)).Elem()
}

type GuestPoliciesInput interface {
	pulumi.Input

	ToGuestPoliciesOutput() GuestPoliciesOutput
	ToGuestPoliciesOutputWithContext(ctx context.Context) GuestPoliciesOutput
}

func (GuestPolicies) ElementType() reflect.Type {
	return reflect.TypeOf((*GuestPolicies)(nil)).Elem()
}

func (i GuestPolicies) ToGuestPoliciesOutput() GuestPoliciesOutput {
	return i.ToGuestPoliciesOutputWithContext(context.Background())
}

func (i GuestPolicies) ToGuestPoliciesOutputWithContext(ctx context.Context) GuestPoliciesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GuestPoliciesOutput)
}

type GuestPoliciesOutput struct {
	*pulumi.OutputState
}

func (GuestPoliciesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GuestPoliciesOutput)(nil)).Elem()
}

func (o GuestPoliciesOutput) ToGuestPoliciesOutput() GuestPoliciesOutput {
	return o
}

func (o GuestPoliciesOutput) ToGuestPoliciesOutputWithContext(ctx context.Context) GuestPoliciesOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(GuestPoliciesOutput{})
}
