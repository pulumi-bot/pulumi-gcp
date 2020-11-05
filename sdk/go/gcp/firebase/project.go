// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package firebase

import (
	"fmt"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// A Google Cloud Firebase instance. This enables Firebase resources on a given google project.
// Since a FirebaseProject is actually also a GCP Project, a FirebaseProject uses underlying GCP
// identifiers (most importantly, the projectId) as its own for easy interop with GCP APIs.
//
// Once Firebase has been added to a Google Project it cannot be removed.
//
// To get more information about Project, see:
//
// * [API documentation](https://firebase.google.com/docs/projects/api/reference/rest/v1beta1/projects)
// * How-to Guides
//     * [Official Documentation](https://firebase.google.com/)
//
// ## Example Usage
type Project struct {
	pulumi.CustomResourceState

	// The GCP project display name
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The number of the google project that firebase is enabled on.
	ProjectNumber pulumi.StringOutput `pulumi:"projectNumber"`
}

// NewProject registers a new resource with the given unique name, arguments, and options.
func NewProject(ctx *pulumi.Context,
	name string, args *ProjectArgs, opts ...pulumi.ResourceOption) (*Project, error) {
	if args == nil {
		args = &ProjectArgs{}
	}
	var resource Project
	err := ctx.RegisterResource("gcp:firebase/project:Project", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProject gets an existing Project resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProject(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProjectState, opts ...pulumi.ResourceOption) (*Project, error) {
	var resource Project
	err := ctx.ReadResource("gcp:firebase/project:Project", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Project resources.
type projectState struct {
	// The GCP project display name
	DisplayName *string `pulumi:"displayName"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The number of the google project that firebase is enabled on.
	ProjectNumber *string `pulumi:"projectNumber"`
}

type ProjectState struct {
	// The GCP project display name
	DisplayName pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The number of the google project that firebase is enabled on.
	ProjectNumber pulumi.StringPtrInput
}

func (ProjectState) ElementType() reflect.Type {
	return reflect.TypeOf((*projectState)(nil)).Elem()
}

type projectArgs struct {
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
}

// The set of arguments for constructing a Project resource.
type ProjectArgs struct {
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
}

func (ProjectArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*projectArgs)(nil)).Elem()
}
