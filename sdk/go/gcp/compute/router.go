// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Represents a Router resource.
//
// To get more information about Router, see:
//
// * [API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/routers)
// * How-to Guides
//     * [Google Cloud Router](https://cloud.google.com/router/docs/)
//
// ## Example Usage
// ### Router Basic
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/compute"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		foobarNetwork, err := compute.NewNetwork(ctx, "foobarNetwork", &compute.NetworkArgs{
// 			AutoCreateSubnetworks: pulumi.Bool(false),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = compute.NewRouter(ctx, "foobarRouter", &compute.RouterArgs{
// 			Network: foobarNetwork.Name,
// 			Bgp: &compute.RouterBgpArgs{
// 				Asn:           pulumi.Int(64514),
// 				AdvertiseMode: pulumi.String("CUSTOM"),
// 				AdvertisedGroups: pulumi.StringArray{
// 					pulumi.String("ALL_SUBNETS"),
// 				},
// 				AdvertisedIpRanges: compute.RouterBgpAdvertisedIpRangeArray{
// 					&compute.RouterBgpAdvertisedIpRangeArgs{
// 						Range: pulumi.String("1.2.3.4"),
// 					},
// 					&compute.RouterBgpAdvertisedIpRangeArgs{
// 						Range: pulumi.String("6.7.0.0/16"),
// 					},
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// Router can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import gcp:compute/router:Router default projects/{{project}}/regions/{{region}}/routers/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:compute/router:Router default {{project}}/{{region}}/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:compute/router:Router default {{region}}/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:compute/router:Router default {{name}}
// ```
type Router struct {
	pulumi.CustomResourceState

	// BGP information specific to this router.
	// Structure is documented below.
	Bgp RouterBgpPtrOutput `pulumi:"bgp"`
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringOutput `pulumi:"creationTimestamp"`
	// User-specified description for the IP range.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Name of the resource. The name must be 1-63 characters long, and
	// comply with RFC1035. Specifically, the name must be 1-63 characters
	// long and match the regular expression `a-z?`
	// which means the first character must be a lowercase letter, and all
	// following characters must be a dash, lowercase letter, or digit,
	// except the last character, which cannot be a dash.
	Name pulumi.StringOutput `pulumi:"name"`
	// A reference to the network to which this router belongs.
	Network pulumi.StringOutput `pulumi:"network"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// Region where the router resides.
	Region pulumi.StringOutput `pulumi:"region"`
	// The URI of the created resource.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
}

// NewRouter registers a new resource with the given unique name, arguments, and options.
func NewRouter(ctx *pulumi.Context,
	name string, args *RouterArgs, opts ...pulumi.ResourceOption) (*Router, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Network == nil {
		return nil, errors.New("invalid value for required argument 'Network'")
	}
	var resource Router
	err := ctx.RegisterResource("gcp:compute/router:Router", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRouter gets an existing Router resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRouter(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RouterState, opts ...pulumi.ResourceOption) (*Router, error) {
	var resource Router
	err := ctx.ReadResource("gcp:compute/router:Router", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Router resources.
type routerState struct {
	// BGP information specific to this router.
	// Structure is documented below.
	Bgp *RouterBgp `pulumi:"bgp"`
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp *string `pulumi:"creationTimestamp"`
	// User-specified description for the IP range.
	Description *string `pulumi:"description"`
	// Name of the resource. The name must be 1-63 characters long, and
	// comply with RFC1035. Specifically, the name must be 1-63 characters
	// long and match the regular expression `a-z?`
	// which means the first character must be a lowercase letter, and all
	// following characters must be a dash, lowercase letter, or digit,
	// except the last character, which cannot be a dash.
	Name *string `pulumi:"name"`
	// A reference to the network to which this router belongs.
	Network *string `pulumi:"network"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// Region where the router resides.
	Region *string `pulumi:"region"`
	// The URI of the created resource.
	SelfLink *string `pulumi:"selfLink"`
}

type RouterState struct {
	// BGP information specific to this router.
	// Structure is documented below.
	Bgp RouterBgpPtrInput
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringPtrInput
	// User-specified description for the IP range.
	Description pulumi.StringPtrInput
	// Name of the resource. The name must be 1-63 characters long, and
	// comply with RFC1035. Specifically, the name must be 1-63 characters
	// long and match the regular expression `a-z?`
	// which means the first character must be a lowercase letter, and all
	// following characters must be a dash, lowercase letter, or digit,
	// except the last character, which cannot be a dash.
	Name pulumi.StringPtrInput
	// A reference to the network to which this router belongs.
	Network pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// Region where the router resides.
	Region pulumi.StringPtrInput
	// The URI of the created resource.
	SelfLink pulumi.StringPtrInput
}

func (RouterState) ElementType() reflect.Type {
	return reflect.TypeOf((*routerState)(nil)).Elem()
}

type routerArgs struct {
	// BGP information specific to this router.
	// Structure is documented below.
	Bgp *RouterBgp `pulumi:"bgp"`
	// User-specified description for the IP range.
	Description *string `pulumi:"description"`
	// Name of the resource. The name must be 1-63 characters long, and
	// comply with RFC1035. Specifically, the name must be 1-63 characters
	// long and match the regular expression `a-z?`
	// which means the first character must be a lowercase letter, and all
	// following characters must be a dash, lowercase letter, or digit,
	// except the last character, which cannot be a dash.
	Name *string `pulumi:"name"`
	// A reference to the network to which this router belongs.
	Network string `pulumi:"network"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// Region where the router resides.
	Region *string `pulumi:"region"`
}

// The set of arguments for constructing a Router resource.
type RouterArgs struct {
	// BGP information specific to this router.
	// Structure is documented below.
	Bgp RouterBgpPtrInput
	// User-specified description for the IP range.
	Description pulumi.StringPtrInput
	// Name of the resource. The name must be 1-63 characters long, and
	// comply with RFC1035. Specifically, the name must be 1-63 characters
	// long and match the regular expression `a-z?`
	// which means the first character must be a lowercase letter, and all
	// following characters must be a dash, lowercase letter, or digit,
	// except the last character, which cannot be a dash.
	Name pulumi.StringPtrInput
	// A reference to the network to which this router belongs.
	Network pulumi.StringInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// Region where the router resides.
	Region pulumi.StringPtrInput
}

func (RouterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*routerArgs)(nil)).Elem()
}

type RouterInput interface {
	pulumi.Input

	ToRouterOutput() RouterOutput
	ToRouterOutputWithContext(ctx context.Context) RouterOutput
}

func (Router) ElementType() reflect.Type {
	return reflect.TypeOf((*Router)(nil)).Elem()
}

func (i Router) ToRouterOutput() RouterOutput {
	return i.ToRouterOutputWithContext(context.Background())
}

func (i Router) ToRouterOutputWithContext(ctx context.Context) RouterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouterOutput)
}

type RouterOutput struct {
	*pulumi.OutputState
}

func (RouterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RouterOutput)(nil)).Elem()
}

func (o RouterOutput) ToRouterOutput() RouterOutput {
	return o
}

func (o RouterOutput) ToRouterOutputWithContext(ctx context.Context) RouterOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(RouterOutput{})
}
