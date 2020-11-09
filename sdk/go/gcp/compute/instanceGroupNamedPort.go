// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Mange the named ports setting for a managed instance group without
// managing the group as whole. This resource is primarily intended for use
// with GKE-generated groups that shouldn't otherwise be managed by other
// tools.
//
// To get more information about InstanceGroupNamedPort, see:
//
// * [API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/instanceGroup)
// * How-to Guides
//     * [Official Documentation](https://cloud.google.com/compute/docs/instance-groups/)
//
// ## Example Usage
//
// ## Import
//
// InstanceGroupNamedPort can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import gcp:compute/instanceGroupNamedPort:InstanceGroupNamedPort default projects/{{project}}/zones/{{zone}}/instanceGroups/{{group}}/{{port}}/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:compute/instanceGroupNamedPort:InstanceGroupNamedPort default {{project}}/{{zone}}/{{group}}/{{port}}/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:compute/instanceGroupNamedPort:InstanceGroupNamedPort default {{zone}}/{{group}}/{{port}}/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:compute/instanceGroupNamedPort:InstanceGroupNamedPort default {{group}}/{{port}}/{{name}}
// ```
type InstanceGroupNamedPort struct {
	pulumi.CustomResourceState

	// The name of the instance group.
	Group pulumi.StringOutput `pulumi:"group"`
	// The name for this named port. The name must be 1-63 characters
	// long, and comply with RFC1035.
	Name pulumi.StringOutput `pulumi:"name"`
	// The port number, which can be a value between 1 and 65535.
	Port pulumi.IntOutput `pulumi:"port"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The zone of the instance group.
	Zone pulumi.StringOutput `pulumi:"zone"`
}

// NewInstanceGroupNamedPort registers a new resource with the given unique name, arguments, and options.
func NewInstanceGroupNamedPort(ctx *pulumi.Context,
	name string, args *InstanceGroupNamedPortArgs, opts ...pulumi.ResourceOption) (*InstanceGroupNamedPort, error) {
	if args == nil || args.Group == nil {
		return nil, errors.New("missing required argument 'Group'")
	}
	if args == nil || args.Port == nil {
		return nil, errors.New("missing required argument 'Port'")
	}
	if args == nil {
		args = &InstanceGroupNamedPortArgs{}
	}
	var resource InstanceGroupNamedPort
	err := ctx.RegisterResource("gcp:compute/instanceGroupNamedPort:InstanceGroupNamedPort", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInstanceGroupNamedPort gets an existing InstanceGroupNamedPort resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInstanceGroupNamedPort(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InstanceGroupNamedPortState, opts ...pulumi.ResourceOption) (*InstanceGroupNamedPort, error) {
	var resource InstanceGroupNamedPort
	err := ctx.ReadResource("gcp:compute/instanceGroupNamedPort:InstanceGroupNamedPort", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering InstanceGroupNamedPort resources.
type instanceGroupNamedPortState struct {
	// The name of the instance group.
	Group *string `pulumi:"group"`
	// The name for this named port. The name must be 1-63 characters
	// long, and comply with RFC1035.
	Name *string `pulumi:"name"`
	// The port number, which can be a value between 1 and 65535.
	Port *int `pulumi:"port"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The zone of the instance group.
	Zone *string `pulumi:"zone"`
}

type InstanceGroupNamedPortState struct {
	// The name of the instance group.
	Group pulumi.StringPtrInput
	// The name for this named port. The name must be 1-63 characters
	// long, and comply with RFC1035.
	Name pulumi.StringPtrInput
	// The port number, which can be a value between 1 and 65535.
	Port pulumi.IntPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The zone of the instance group.
	Zone pulumi.StringPtrInput
}

func (InstanceGroupNamedPortState) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceGroupNamedPortState)(nil)).Elem()
}

type instanceGroupNamedPortArgs struct {
	// The name of the instance group.
	Group string `pulumi:"group"`
	// The name for this named port. The name must be 1-63 characters
	// long, and comply with RFC1035.
	Name *string `pulumi:"name"`
	// The port number, which can be a value between 1 and 65535.
	Port int `pulumi:"port"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The zone of the instance group.
	Zone *string `pulumi:"zone"`
}

// The set of arguments for constructing a InstanceGroupNamedPort resource.
type InstanceGroupNamedPortArgs struct {
	// The name of the instance group.
	Group pulumi.StringInput
	// The name for this named port. The name must be 1-63 characters
	// long, and comply with RFC1035.
	Name pulumi.StringPtrInput
	// The port number, which can be a value between 1 and 65535.
	Port pulumi.IntInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The zone of the instance group.
	Zone pulumi.StringPtrInput
}

func (InstanceGroupNamedPortArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceGroupNamedPortArgs)(nil)).Elem()
}
