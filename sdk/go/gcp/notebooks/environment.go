// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package notebooks

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// A Cloud AI Platform Notebook environment.
//
// To get more information about Environment, see:
//
// * [API documentation](https://cloud.google.com/ai-platform/notebooks/docs/reference/rest)
// * How-to Guides
//     * [Official Documentation](https://cloud.google.com/ai-platform-notebooks)
//
// ## Example Usage
//
// ## Import
//
// Environment can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import gcp:notebooks/environment:Environment default projects/{{project}}/locations/{{location}}/environments/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:notebooks/environment:Environment default {{project}}/{{location}}/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:notebooks/environment:Environment default {{location}}/{{name}}
// ```
type Environment struct {
	pulumi.CustomResourceState

	// Use a container image to start the notebook instance.
	// Structure is documented below.
	ContainerImage EnvironmentContainerImagePtrOutput `pulumi:"containerImage"`
	// Instance creation time
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// A brief description of this environment.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Display name of this environment for the UI.
	DisplayName pulumi.StringPtrOutput `pulumi:"displayName"`
	// A reference to the zone where the machine resides.
	Location pulumi.StringOutput `pulumi:"location"`
	// The name specified for the Environment instance.
	// Format: projects/{project_id}/locations/{location}/environments/{environmentId}
	Name pulumi.StringOutput `pulumi:"name"`
	// Path to a Bash script that automatically runs after a notebook instance fully boots up.
	// The path must be a URL or Cloud Storage path. Example: "gs://path-to-file/file-name"
	PostStartupScript pulumi.StringPtrOutput `pulumi:"postStartupScript"`
	// The name of the Google Cloud project that this VM image belongs to.
	// Format: projects/{project_id}
	Project pulumi.StringOutput `pulumi:"project"`
	// Use a Compute Engine VM image to start the notebook instance.
	// Structure is documented below.
	VmImage EnvironmentVmImagePtrOutput `pulumi:"vmImage"`
}

// NewEnvironment registers a new resource with the given unique name, arguments, and options.
func NewEnvironment(ctx *pulumi.Context,
	name string, args *EnvironmentArgs, opts ...pulumi.ResourceOption) (*Environment, error) {
	if args == nil || args.Location == nil {
		return nil, errors.New("missing required argument 'Location'")
	}
	if args == nil {
		args = &EnvironmentArgs{}
	}
	var resource Environment
	err := ctx.RegisterResource("gcp:notebooks/environment:Environment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEnvironment gets an existing Environment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEnvironment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EnvironmentState, opts ...pulumi.ResourceOption) (*Environment, error) {
	var resource Environment
	err := ctx.ReadResource("gcp:notebooks/environment:Environment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Environment resources.
type environmentState struct {
	// Use a container image to start the notebook instance.
	// Structure is documented below.
	ContainerImage *EnvironmentContainerImage `pulumi:"containerImage"`
	// Instance creation time
	CreateTime *string `pulumi:"createTime"`
	// A brief description of this environment.
	Description *string `pulumi:"description"`
	// Display name of this environment for the UI.
	DisplayName *string `pulumi:"displayName"`
	// A reference to the zone where the machine resides.
	Location *string `pulumi:"location"`
	// The name specified for the Environment instance.
	// Format: projects/{project_id}/locations/{location}/environments/{environmentId}
	Name *string `pulumi:"name"`
	// Path to a Bash script that automatically runs after a notebook instance fully boots up.
	// The path must be a URL or Cloud Storage path. Example: "gs://path-to-file/file-name"
	PostStartupScript *string `pulumi:"postStartupScript"`
	// The name of the Google Cloud project that this VM image belongs to.
	// Format: projects/{project_id}
	Project *string `pulumi:"project"`
	// Use a Compute Engine VM image to start the notebook instance.
	// Structure is documented below.
	VmImage *EnvironmentVmImage `pulumi:"vmImage"`
}

type EnvironmentState struct {
	// Use a container image to start the notebook instance.
	// Structure is documented below.
	ContainerImage EnvironmentContainerImagePtrInput
	// Instance creation time
	CreateTime pulumi.StringPtrInput
	// A brief description of this environment.
	Description pulumi.StringPtrInput
	// Display name of this environment for the UI.
	DisplayName pulumi.StringPtrInput
	// A reference to the zone where the machine resides.
	Location pulumi.StringPtrInput
	// The name specified for the Environment instance.
	// Format: projects/{project_id}/locations/{location}/environments/{environmentId}
	Name pulumi.StringPtrInput
	// Path to a Bash script that automatically runs after a notebook instance fully boots up.
	// The path must be a URL or Cloud Storage path. Example: "gs://path-to-file/file-name"
	PostStartupScript pulumi.StringPtrInput
	// The name of the Google Cloud project that this VM image belongs to.
	// Format: projects/{project_id}
	Project pulumi.StringPtrInput
	// Use a Compute Engine VM image to start the notebook instance.
	// Structure is documented below.
	VmImage EnvironmentVmImagePtrInput
}

func (EnvironmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*environmentState)(nil)).Elem()
}

type environmentArgs struct {
	// Use a container image to start the notebook instance.
	// Structure is documented below.
	ContainerImage *EnvironmentContainerImage `pulumi:"containerImage"`
	// A brief description of this environment.
	Description *string `pulumi:"description"`
	// Display name of this environment for the UI.
	DisplayName *string `pulumi:"displayName"`
	// A reference to the zone where the machine resides.
	Location string `pulumi:"location"`
	// The name specified for the Environment instance.
	// Format: projects/{project_id}/locations/{location}/environments/{environmentId}
	Name *string `pulumi:"name"`
	// Path to a Bash script that automatically runs after a notebook instance fully boots up.
	// The path must be a URL or Cloud Storage path. Example: "gs://path-to-file/file-name"
	PostStartupScript *string `pulumi:"postStartupScript"`
	// The name of the Google Cloud project that this VM image belongs to.
	// Format: projects/{project_id}
	Project *string `pulumi:"project"`
	// Use a Compute Engine VM image to start the notebook instance.
	// Structure is documented below.
	VmImage *EnvironmentVmImage `pulumi:"vmImage"`
}

// The set of arguments for constructing a Environment resource.
type EnvironmentArgs struct {
	// Use a container image to start the notebook instance.
	// Structure is documented below.
	ContainerImage EnvironmentContainerImagePtrInput
	// A brief description of this environment.
	Description pulumi.StringPtrInput
	// Display name of this environment for the UI.
	DisplayName pulumi.StringPtrInput
	// A reference to the zone where the machine resides.
	Location pulumi.StringInput
	// The name specified for the Environment instance.
	// Format: projects/{project_id}/locations/{location}/environments/{environmentId}
	Name pulumi.StringPtrInput
	// Path to a Bash script that automatically runs after a notebook instance fully boots up.
	// The path must be a URL or Cloud Storage path. Example: "gs://path-to-file/file-name"
	PostStartupScript pulumi.StringPtrInput
	// The name of the Google Cloud project that this VM image belongs to.
	// Format: projects/{project_id}
	Project pulumi.StringPtrInput
	// Use a Compute Engine VM image to start the notebook instance.
	// Structure is documented below.
	VmImage EnvironmentVmImagePtrInput
}

func (EnvironmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*environmentArgs)(nil)).Elem()
}
