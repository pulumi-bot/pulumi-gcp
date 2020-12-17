// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cloudrun

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Resource to hold the state and status of a user's domain mapping.
//
// To get more information about DomainMapping, see:
//
// * [API documentation](https://cloud.google.com/run/docs/reference/rest/v1/projects.locations.domainmappings)
// * How-to Guides
//     * [Official Documentation](https://cloud.google.com/run/docs/mapping-custom-domains)
//
// ## Example Usage
// ### Cloud Run Domain Mapping Basic
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/cloudrun"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		defaultService, err := cloudrun.NewService(ctx, "defaultService", &cloudrun.ServiceArgs{
// 			Location: pulumi.String("us-central1"),
// 			Metadata: &cloudrun.ServiceMetadataArgs{
// 				Namespace: pulumi.String("my-project-name"),
// 			},
// 			Template: &cloudrun.ServiceTemplateArgs{
// 				Spec: &cloudrun.ServiceTemplateSpecArgs{
// 					Containers: cloudrun.ServiceTemplateSpecContainerArray{
// 						&cloudrun.ServiceTemplateSpecContainerArgs{
// 							Image: pulumi.String("gcr.io/cloudrun/hello"),
// 						},
// 					},
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = cloudrun.NewDomainMapping(ctx, "defaultDomainMapping", &cloudrun.DomainMappingArgs{
// 			Location: pulumi.String("us-central1"),
// 			Metadata: &cloudrun.DomainMappingMetadataArgs{
// 				Namespace: pulumi.String("my-project-name"),
// 				Annotations: pulumi.StringMap{
// 					"run.googleapis.com/launch-stage": pulumi.String("BETA"),
// 				},
// 			},
// 			Spec: &cloudrun.DomainMappingSpecArgs{
// 				RouteName: defaultService.Name,
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
// DomainMapping can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import gcp:cloudrun/domainMapping:DomainMapping default locations/{{location}}/namespaces/{{project}}/domainmappings/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:cloudrun/domainMapping:DomainMapping default {{location}}/{{project}}/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:cloudrun/domainMapping:DomainMapping default {{location}}/{{name}}
// ```
type DomainMapping struct {
	pulumi.CustomResourceState

	// The location of the cloud run instance. eg us-central1
	Location pulumi.StringOutput `pulumi:"location"`
	// Metadata associated with this DomainMapping.
	// Structure is documented below.
	Metadata DomainMappingMetadataOutput `pulumi:"metadata"`
	// Name should be a verified domain
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The spec for this DomainMapping.
	// Structure is documented below.
	Spec DomainMappingSpecOutput `pulumi:"spec"`
	// The current status of the DomainMapping.
	Statuses DomainMappingStatusArrayOutput `pulumi:"statuses"`
}

// NewDomainMapping registers a new resource with the given unique name, arguments, and options.
func NewDomainMapping(ctx *pulumi.Context,
	name string, args *DomainMappingArgs, opts ...pulumi.ResourceOption) (*DomainMapping, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Location == nil {
		return nil, errors.New("invalid value for required argument 'Location'")
	}
	if args.Metadata == nil {
		return nil, errors.New("invalid value for required argument 'Metadata'")
	}
	if args.Spec == nil {
		return nil, errors.New("invalid value for required argument 'Spec'")
	}
	var resource DomainMapping
	err := ctx.RegisterResource("gcp:cloudrun/domainMapping:DomainMapping", name, args, &resource, opts...)
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
	err := ctx.ReadResource("gcp:cloudrun/domainMapping:DomainMapping", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DomainMapping resources.
type domainMappingState struct {
	// The location of the cloud run instance. eg us-central1
	Location *string `pulumi:"location"`
	// Metadata associated with this DomainMapping.
	// Structure is documented below.
	Metadata *DomainMappingMetadata `pulumi:"metadata"`
	// Name should be a verified domain
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The spec for this DomainMapping.
	// Structure is documented below.
	Spec *DomainMappingSpec `pulumi:"spec"`
	// The current status of the DomainMapping.
	Statuses []DomainMappingStatus `pulumi:"statuses"`
}

type DomainMappingState struct {
	// The location of the cloud run instance. eg us-central1
	Location pulumi.StringPtrInput
	// Metadata associated with this DomainMapping.
	// Structure is documented below.
	Metadata DomainMappingMetadataPtrInput
	// Name should be a verified domain
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The spec for this DomainMapping.
	// Structure is documented below.
	Spec DomainMappingSpecPtrInput
	// The current status of the DomainMapping.
	Statuses DomainMappingStatusArrayInput
}

func (DomainMappingState) ElementType() reflect.Type {
	return reflect.TypeOf((*domainMappingState)(nil)).Elem()
}

type domainMappingArgs struct {
	// The location of the cloud run instance. eg us-central1
	Location string `pulumi:"location"`
	// Metadata associated with this DomainMapping.
	// Structure is documented below.
	Metadata DomainMappingMetadata `pulumi:"metadata"`
	// Name should be a verified domain
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The spec for this DomainMapping.
	// Structure is documented below.
	Spec DomainMappingSpec `pulumi:"spec"`
}

// The set of arguments for constructing a DomainMapping resource.
type DomainMappingArgs struct {
	// The location of the cloud run instance. eg us-central1
	Location pulumi.StringInput
	// Metadata associated with this DomainMapping.
	// Structure is documented below.
	Metadata DomainMappingMetadataInput
	// Name should be a verified domain
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The spec for this DomainMapping.
	// Structure is documented below.
	Spec DomainMappingSpecInput
}

func (DomainMappingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*domainMappingArgs)(nil)).Elem()
}

type DomainMappingInput interface {
	pulumi.Input

	ToDomainMappingOutput() DomainMappingOutput
	ToDomainMappingOutputWithContext(ctx context.Context) DomainMappingOutput
}

func (*DomainMapping) ElementType() reflect.Type {
	return reflect.TypeOf((*DomainMapping)(nil))
}

func (i *DomainMapping) ToDomainMappingOutput() DomainMappingOutput {
	return i.ToDomainMappingOutputWithContext(context.Background())
}

func (i *DomainMapping) ToDomainMappingOutputWithContext(ctx context.Context) DomainMappingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainMappingOutput)
}

func (i *DomainMapping) ToDomainMappingPtrOutput() DomainMappingPtrOutput {
	return i.ToDomainMappingPtrOutputWithContext(context.Background())
}

func (i *DomainMapping) ToDomainMappingPtrOutputWithContext(ctx context.Context) DomainMappingPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainMappingPtrOutput)
}

type DomainMappingPtrInput interface {
	pulumi.Input

	ToDomainMappingPtrOutput() DomainMappingPtrOutput
	ToDomainMappingPtrOutputWithContext(ctx context.Context) DomainMappingPtrOutput
}

type domainMappingPtrType DomainMappingArgs

func (*domainMappingPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DomainMapping)(nil))
}

func (i *domainMappingPtrType) ToDomainMappingPtrOutput() DomainMappingPtrOutput {
	return i.ToDomainMappingPtrOutputWithContext(context.Background())
}

func (i *domainMappingPtrType) ToDomainMappingPtrOutputWithContext(ctx context.Context) DomainMappingPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainMappingOutput).ToDomainMappingPtrOutput()
}

// DomainMappingArrayInput is an input type that accepts DomainMappingArray and DomainMappingArrayOutput values.
// You can construct a concrete instance of `DomainMappingArrayInput` via:
//
//          DomainMappingArray{ DomainMappingArgs{...} }
type DomainMappingArrayInput interface {
	pulumi.Input

	ToDomainMappingArrayOutput() DomainMappingArrayOutput
	ToDomainMappingArrayOutputWithContext(context.Context) DomainMappingArrayOutput
}

type DomainMappingArray []DomainMappingInput

func (DomainMappingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DomainMapping)(nil))
}

func (i DomainMappingArray) ToDomainMappingArrayOutput() DomainMappingArrayOutput {
	return i.ToDomainMappingArrayOutputWithContext(context.Background())
}

func (i DomainMappingArray) ToDomainMappingArrayOutputWithContext(ctx context.Context) DomainMappingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainMappingArrayOutput)
}

// DomainMappingMapInput is an input type that accepts DomainMappingMap and DomainMappingMapOutput values.
// You can construct a concrete instance of `DomainMappingMapInput` via:
//
//          DomainMappingMap{ "key": DomainMappingArgs{...} }
type DomainMappingMapInput interface {
	pulumi.Input

	ToDomainMappingMapOutput() DomainMappingMapOutput
	ToDomainMappingMapOutputWithContext(context.Context) DomainMappingMapOutput
}

type DomainMappingMap map[string]DomainMappingInput

func (DomainMappingMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]DomainMapping)(nil))
}

func (i DomainMappingMap) ToDomainMappingMapOutput() DomainMappingMapOutput {
	return i.ToDomainMappingMapOutputWithContext(context.Background())
}

func (i DomainMappingMap) ToDomainMappingMapOutputWithContext(ctx context.Context) DomainMappingMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainMappingMapOutput)
}

type DomainMappingOutput struct {
	*pulumi.OutputState
}

func (DomainMappingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DomainMapping)(nil))
}

func (o DomainMappingOutput) ToDomainMappingOutput() DomainMappingOutput {
	return o
}

func (o DomainMappingOutput) ToDomainMappingOutputWithContext(ctx context.Context) DomainMappingOutput {
	return o
}

func (o DomainMappingOutput) ToDomainMappingPtrOutput() DomainMappingPtrOutput {
	return o.ToDomainMappingPtrOutputWithContext(context.Background())
}

func (o DomainMappingOutput) ToDomainMappingPtrOutputWithContext(ctx context.Context) DomainMappingPtrOutput {
	return o.ApplyT(func(v DomainMapping) *DomainMapping {
		return &v
	}).(DomainMappingPtrOutput)
}

type DomainMappingPtrOutput struct {
	*pulumi.OutputState
}

func (DomainMappingPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DomainMapping)(nil))
}

func (o DomainMappingPtrOutput) ToDomainMappingPtrOutput() DomainMappingPtrOutput {
	return o
}

func (o DomainMappingPtrOutput) ToDomainMappingPtrOutputWithContext(ctx context.Context) DomainMappingPtrOutput {
	return o
}

type DomainMappingArrayOutput struct{ *pulumi.OutputState }

func (DomainMappingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DomainMapping)(nil))
}

func (o DomainMappingArrayOutput) ToDomainMappingArrayOutput() DomainMappingArrayOutput {
	return o
}

func (o DomainMappingArrayOutput) ToDomainMappingArrayOutputWithContext(ctx context.Context) DomainMappingArrayOutput {
	return o
}

func (o DomainMappingArrayOutput) Index(i pulumi.IntInput) DomainMappingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DomainMapping {
		return vs[0].([]DomainMapping)[vs[1].(int)]
	}).(DomainMappingOutput)
}

type DomainMappingMapOutput struct{ *pulumi.OutputState }

func (DomainMappingMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]DomainMapping)(nil))
}

func (o DomainMappingMapOutput) ToDomainMappingMapOutput() DomainMappingMapOutput {
	return o
}

func (o DomainMappingMapOutput) ToDomainMappingMapOutputWithContext(ctx context.Context) DomainMappingMapOutput {
	return o
}

func (o DomainMappingMapOutput) MapIndex(k pulumi.StringInput) DomainMappingOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) DomainMapping {
		return vs[0].(map[string]DomainMapping)[vs[1].(string)]
	}).(DomainMappingOutput)
}

func init() {
	pulumi.RegisterOutputType(DomainMappingOutput{})
	pulumi.RegisterOutputType(DomainMappingPtrOutput{})
	pulumi.RegisterOutputType(DomainMappingArrayOutput{})
	pulumi.RegisterOutputType(DomainMappingMapOutput{})
}
