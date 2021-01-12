// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package accesscontextmanager

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Replace all existing Service Perimeters in an Access Policy with the Service Perimeters provided. This is done atomically.
// This is a bulk edit of all Service Perimeters and may override existing Service Perimeters created by `accesscontextmanager.ServicePerimeter`,
// thus causing a permadiff if used alongside `accesscontextmanager.ServicePerimeter` on the same parent.
//
// To get more information about ServicePerimeters, see:
//
// * [API documentation](https://cloud.google.com/access-context-manager/docs/reference/rest/v1/accessPolicies.servicePerimeters)
// * How-to Guides
//     * [Service Perimeter Quickstart](https://cloud.google.com/vpc-service-controls/docs/quickstart)
//
// ## Example Usage
// ### Access Context Manager Service Perimeters Basic
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/accesscontextmanager"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := accesscontextmanager.NewAccessPolicy(ctx, "access_policy", &accesscontextmanager.AccessPolicyArgs{
// 			Parent: pulumi.String("organizations/123456789"),
// 			Title:  pulumi.String("my policy"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = accesscontextmanager.NewServicePerimeters(ctx, "service_perimeter", &accesscontextmanager.ServicePerimetersArgs{
// 			Parent: access_policy.Name.ApplyT(func(name string) (string, error) {
// 				return fmt.Sprintf("%v%v", "accessPolicies/", name), nil
// 			}).(pulumi.StringOutput),
// 			ServicePerimeters: accesscontextmanager.ServicePerimetersServicePerimeterArray{
// 				&accesscontextmanager.ServicePerimetersServicePerimeterArgs{
// 					Name: access_policy.Name.ApplyT(func(name string) (string, error) {
// 						return fmt.Sprintf("%v%v%v", "accessPolicies/", name, "/servicePerimeters/"), nil
// 					}).(pulumi.StringOutput),
// 					Status: &accesscontextmanager.ServicePerimetersServicePerimeterStatusArgs{
// 						RestrictedServices: pulumi.StringArray{
// 							pulumi.String("storage.googleapis.com"),
// 						},
// 					},
// 					Title: pulumi.String(""),
// 				},
// 				&accesscontextmanager.ServicePerimetersServicePerimeterArgs{
// 					Name: access_policy.Name.ApplyT(func(name string) (string, error) {
// 						return fmt.Sprintf("%v%v%v", "accessPolicies/", name, "/servicePerimeters/"), nil
// 					}).(pulumi.StringOutput),
// 					Status: &accesscontextmanager.ServicePerimetersServicePerimeterStatusArgs{
// 						RestrictedServices: pulumi.StringArray{
// 							pulumi.String("bigtable.googleapis.com"),
// 						},
// 					},
// 					Title: pulumi.String(""),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = accesscontextmanager.NewAccessLevel(ctx, "access_level", &accesscontextmanager.AccessLevelArgs{
// 			Basic: &accesscontextmanager.AccessLevelBasicArgs{
// 				Conditions: accesscontextmanager.AccessLevelBasicConditionArray{
// 					&accesscontextmanager.AccessLevelBasicConditionArgs{
// 						DevicePolicy: &accesscontextmanager.AccessLevelBasicConditionDevicePolicyArgs{
// 							OsConstraints: accesscontextmanager.AccessLevelBasicConditionDevicePolicyOsConstraintArray{
// 								&accesscontextmanager.AccessLevelBasicConditionDevicePolicyOsConstraintArgs{
// 									OsType: pulumi.String("DESKTOP_CHROME_OS"),
// 								},
// 							},
// 							RequireScreenLock: pulumi.Bool(false),
// 						},
// 						Regions: pulumi.StringArray{
// 							pulumi.String("CH"),
// 							pulumi.String("IT"),
// 							pulumi.String("US"),
// 						},
// 					},
// 				},
// 			},
// 			Parent: access_policy.Name.ApplyT(func(name string) (string, error) {
// 				return fmt.Sprintf("%v%v", "accessPolicies/", name), nil
// 			}).(pulumi.StringOutput),
// 			Title: pulumi.String("chromeos_no_lock"),
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
// ServicePerimeters can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import gcp:accesscontextmanager/servicePerimeters:ServicePerimeters default {{parent}}/servicePerimeters
// ```
//
// ```sh
//  $ pulumi import gcp:accesscontextmanager/servicePerimeters:ServicePerimeters default {{parent}}
// ```
type ServicePerimeters struct {
	pulumi.CustomResourceState

	// The AccessPolicy this ServicePerimeter lives in.
	// Format: accessPolicies/{policy_id}
	Parent pulumi.StringOutput `pulumi:"parent"`
	// The desired Service Perimeters that should replace all existing Service Perimeters in the Access Policy.
	// Structure is documented below.
	ServicePerimeters ServicePerimetersServicePerimeterArrayOutput `pulumi:"servicePerimeters"`
}

// NewServicePerimeters registers a new resource with the given unique name, arguments, and options.
func NewServicePerimeters(ctx *pulumi.Context,
	name string, args *ServicePerimetersArgs, opts ...pulumi.ResourceOption) (*ServicePerimeters, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Parent == nil {
		return nil, errors.New("invalid value for required argument 'Parent'")
	}
	var resource ServicePerimeters
	err := ctx.RegisterResource("gcp:accesscontextmanager/servicePerimeters:ServicePerimeters", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServicePerimeters gets an existing ServicePerimeters resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServicePerimeters(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServicePerimetersState, opts ...pulumi.ResourceOption) (*ServicePerimeters, error) {
	var resource ServicePerimeters
	err := ctx.ReadResource("gcp:accesscontextmanager/servicePerimeters:ServicePerimeters", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServicePerimeters resources.
type servicePerimetersState struct {
	// The AccessPolicy this ServicePerimeter lives in.
	// Format: accessPolicies/{policy_id}
	Parent *string `pulumi:"parent"`
	// The desired Service Perimeters that should replace all existing Service Perimeters in the Access Policy.
	// Structure is documented below.
	ServicePerimeters []ServicePerimetersServicePerimeter `pulumi:"servicePerimeters"`
}

type ServicePerimetersState struct {
	// The AccessPolicy this ServicePerimeter lives in.
	// Format: accessPolicies/{policy_id}
	Parent pulumi.StringPtrInput
	// The desired Service Perimeters that should replace all existing Service Perimeters in the Access Policy.
	// Structure is documented below.
	ServicePerimeters ServicePerimetersServicePerimeterArrayInput
}

func (ServicePerimetersState) ElementType() reflect.Type {
	return reflect.TypeOf((*servicePerimetersState)(nil)).Elem()
}

type servicePerimetersArgs struct {
	// The AccessPolicy this ServicePerimeter lives in.
	// Format: accessPolicies/{policy_id}
	Parent string `pulumi:"parent"`
	// The desired Service Perimeters that should replace all existing Service Perimeters in the Access Policy.
	// Structure is documented below.
	ServicePerimeters []ServicePerimetersServicePerimeter `pulumi:"servicePerimeters"`
}

// The set of arguments for constructing a ServicePerimeters resource.
type ServicePerimetersArgs struct {
	// The AccessPolicy this ServicePerimeter lives in.
	// Format: accessPolicies/{policy_id}
	Parent pulumi.StringInput
	// The desired Service Perimeters that should replace all existing Service Perimeters in the Access Policy.
	// Structure is documented below.
	ServicePerimeters ServicePerimetersServicePerimeterArrayInput
}

func (ServicePerimetersArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*servicePerimetersArgs)(nil)).Elem()
}

type ServicePerimetersInput interface {
	pulumi.Input

	ToServicePerimetersOutput() ServicePerimetersOutput
	ToServicePerimetersOutputWithContext(ctx context.Context) ServicePerimetersOutput
}

func (*ServicePerimeters) ElementType() reflect.Type {
	return reflect.TypeOf((*ServicePerimeters)(nil))
}

func (i *ServicePerimeters) ToServicePerimetersOutput() ServicePerimetersOutput {
	return i.ToServicePerimetersOutputWithContext(context.Background())
}

func (i *ServicePerimeters) ToServicePerimetersOutputWithContext(ctx context.Context) ServicePerimetersOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServicePerimetersOutput)
}

func (i *ServicePerimeters) ToServicePerimetersPtrOutput() ServicePerimetersPtrOutput {
	return i.ToServicePerimetersPtrOutputWithContext(context.Background())
}

func (i *ServicePerimeters) ToServicePerimetersPtrOutputWithContext(ctx context.Context) ServicePerimetersPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServicePerimetersPtrOutput)
}

type ServicePerimetersPtrInput interface {
	pulumi.Input

	ToServicePerimetersPtrOutput() ServicePerimetersPtrOutput
	ToServicePerimetersPtrOutputWithContext(ctx context.Context) ServicePerimetersPtrOutput
}

type servicePerimetersPtrType ServicePerimetersArgs

func (*servicePerimetersPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ServicePerimeters)(nil))
}

func (i *servicePerimetersPtrType) ToServicePerimetersPtrOutput() ServicePerimetersPtrOutput {
	return i.ToServicePerimetersPtrOutputWithContext(context.Background())
}

func (i *servicePerimetersPtrType) ToServicePerimetersPtrOutputWithContext(ctx context.Context) ServicePerimetersPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServicePerimetersOutput).ToServicePerimetersPtrOutput()
}

// ServicePerimetersArrayInput is an input type that accepts ServicePerimetersArray and ServicePerimetersArrayOutput values.
// You can construct a concrete instance of `ServicePerimetersArrayInput` via:
//
//          ServicePerimetersArray{ ServicePerimetersArgs{...} }
type ServicePerimetersArrayInput interface {
	pulumi.Input

	ToServicePerimetersArrayOutput() ServicePerimetersArrayOutput
	ToServicePerimetersArrayOutputWithContext(context.Context) ServicePerimetersArrayOutput
}

type ServicePerimetersArray []ServicePerimetersInput

func (ServicePerimetersArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ServicePerimeters)(nil))
}

func (i ServicePerimetersArray) ToServicePerimetersArrayOutput() ServicePerimetersArrayOutput {
	return i.ToServicePerimetersArrayOutputWithContext(context.Background())
}

func (i ServicePerimetersArray) ToServicePerimetersArrayOutputWithContext(ctx context.Context) ServicePerimetersArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServicePerimetersArrayOutput)
}

// ServicePerimetersMapInput is an input type that accepts ServicePerimetersMap and ServicePerimetersMapOutput values.
// You can construct a concrete instance of `ServicePerimetersMapInput` via:
//
//          ServicePerimetersMap{ "key": ServicePerimetersArgs{...} }
type ServicePerimetersMapInput interface {
	pulumi.Input

	ToServicePerimetersMapOutput() ServicePerimetersMapOutput
	ToServicePerimetersMapOutputWithContext(context.Context) ServicePerimetersMapOutput
}

type ServicePerimetersMap map[string]ServicePerimetersInput

func (ServicePerimetersMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ServicePerimeters)(nil))
}

func (i ServicePerimetersMap) ToServicePerimetersMapOutput() ServicePerimetersMapOutput {
	return i.ToServicePerimetersMapOutputWithContext(context.Background())
}

func (i ServicePerimetersMap) ToServicePerimetersMapOutputWithContext(ctx context.Context) ServicePerimetersMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServicePerimetersMapOutput)
}

type ServicePerimetersOutput struct {
	*pulumi.OutputState
}

func (ServicePerimetersOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServicePerimeters)(nil))
}

func (o ServicePerimetersOutput) ToServicePerimetersOutput() ServicePerimetersOutput {
	return o
}

func (o ServicePerimetersOutput) ToServicePerimetersOutputWithContext(ctx context.Context) ServicePerimetersOutput {
	return o
}

func (o ServicePerimetersOutput) ToServicePerimetersPtrOutput() ServicePerimetersPtrOutput {
	return o.ToServicePerimetersPtrOutputWithContext(context.Background())
}

func (o ServicePerimetersOutput) ToServicePerimetersPtrOutputWithContext(ctx context.Context) ServicePerimetersPtrOutput {
	return o.ApplyT(func(v ServicePerimeters) *ServicePerimeters {
		return &v
	}).(ServicePerimetersPtrOutput)
}

type ServicePerimetersPtrOutput struct {
	*pulumi.OutputState
}

func (ServicePerimetersPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServicePerimeters)(nil))
}

func (o ServicePerimetersPtrOutput) ToServicePerimetersPtrOutput() ServicePerimetersPtrOutput {
	return o
}

func (o ServicePerimetersPtrOutput) ToServicePerimetersPtrOutputWithContext(ctx context.Context) ServicePerimetersPtrOutput {
	return o
}

type ServicePerimetersArrayOutput struct{ *pulumi.OutputState }

func (ServicePerimetersArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ServicePerimeters)(nil))
}

func (o ServicePerimetersArrayOutput) ToServicePerimetersArrayOutput() ServicePerimetersArrayOutput {
	return o
}

func (o ServicePerimetersArrayOutput) ToServicePerimetersArrayOutputWithContext(ctx context.Context) ServicePerimetersArrayOutput {
	return o
}

func (o ServicePerimetersArrayOutput) Index(i pulumi.IntInput) ServicePerimetersOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ServicePerimeters {
		return vs[0].([]ServicePerimeters)[vs[1].(int)]
	}).(ServicePerimetersOutput)
}

type ServicePerimetersMapOutput struct{ *pulumi.OutputState }

func (ServicePerimetersMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ServicePerimeters)(nil))
}

func (o ServicePerimetersMapOutput) ToServicePerimetersMapOutput() ServicePerimetersMapOutput {
	return o
}

func (o ServicePerimetersMapOutput) ToServicePerimetersMapOutputWithContext(ctx context.Context) ServicePerimetersMapOutput {
	return o
}

func (o ServicePerimetersMapOutput) MapIndex(k pulumi.StringInput) ServicePerimetersOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ServicePerimeters {
		return vs[0].(map[string]ServicePerimeters)[vs[1].(string)]
	}).(ServicePerimetersOutput)
}

func init() {
	pulumi.RegisterOutputType(ServicePerimetersOutput{})
	pulumi.RegisterOutputType(ServicePerimetersPtrOutput{})
	pulumi.RegisterOutputType(ServicePerimetersArrayOutput{})
	pulumi.RegisterOutputType(ServicePerimetersMapOutput{})
}
