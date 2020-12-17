// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package diagflow

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Represents an entity type. Entity types serve as a tool for extracting parameter values from natural language queries.
//
// To get more information about EntityType, see:
//
// * [API documentation](https://cloud.google.com/dialogflow/docs/reference/rest/v2/projects.agent.entityTypes)
// * How-to Guides
//     * [Official Documentation](https://cloud.google.com/dialogflow/docs/)
//
// ## Example Usage
// ### Dialogflow Entity Type Basic
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/diagflow"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		basicAgent, err := diagflow.NewAgent(ctx, "basicAgent", &diagflow.AgentArgs{
// 			DisplayName:         pulumi.String("example_agent"),
// 			DefaultLanguageCode: pulumi.String("en"),
// 			TimeZone:            pulumi.String("America/New_York"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = diagflow.NewEntityType(ctx, "basicEntityType", &diagflow.EntityTypeArgs{
// 			DisplayName: pulumi.String(""),
// 			Kind:        pulumi.String("KIND_MAP"),
// 			Entities: diagflow.EntityTypeEntityArray{
// 				&diagflow.EntityTypeEntityArgs{
// 					Value: pulumi.String("value1"),
// 					Synonyms: pulumi.StringArray{
// 						pulumi.String("synonym1"),
// 						pulumi.String("synonym2"),
// 					},
// 				},
// 				&diagflow.EntityTypeEntityArgs{
// 					Value: pulumi.String("value2"),
// 					Synonyms: pulumi.StringArray{
// 						pulumi.String("synonym3"),
// 						pulumi.String("synonym4"),
// 					},
// 				},
// 			},
// 		}, pulumi.DependsOn([]pulumi.Resource{
// 			basicAgent,
// 		}))
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
// EntityType can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import gcp:diagflow/entityType:EntityType default {{name}}
// ```
type EntityType struct {
	pulumi.CustomResourceState

	// The name of this entity type to be displayed on the console.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// Enables fuzzy entity extraction during classification.
	EnableFuzzyExtraction pulumi.BoolPtrOutput `pulumi:"enableFuzzyExtraction"`
	// The collection of entity entries associated with the entity type.
	// Structure is documented below.
	Entities EntityTypeEntityArrayOutput `pulumi:"entities"`
	// Indicates the kind of entity type.
	// * KIND_MAP: Map entity types allow mapping of a group of synonyms to a reference value.
	// * KIND_LIST: List entity types contain a set of entries that do not map to reference values. However, list entity
	//   types can contain references to other entity types (with or without aliases).
	// * KIND_REGEXP: Regexp entity types allow to specify regular expressions in entries values.
	//   Possible values are `KIND_MAP`, `KIND_LIST`, and `KIND_REGEXP`.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// The unique identifier of the entity type. Format: projects/<Project ID>/agent/entityTypes/<Entity type ID>.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
}

// NewEntityType registers a new resource with the given unique name, arguments, and options.
func NewEntityType(ctx *pulumi.Context,
	name string, args *EntityTypeArgs, opts ...pulumi.ResourceOption) (*EntityType, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	if args.Kind == nil {
		return nil, errors.New("invalid value for required argument 'Kind'")
	}
	var resource EntityType
	err := ctx.RegisterResource("gcp:diagflow/entityType:EntityType", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEntityType gets an existing EntityType resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEntityType(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EntityTypeState, opts ...pulumi.ResourceOption) (*EntityType, error) {
	var resource EntityType
	err := ctx.ReadResource("gcp:diagflow/entityType:EntityType", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EntityType resources.
type entityTypeState struct {
	// The name of this entity type to be displayed on the console.
	DisplayName *string `pulumi:"displayName"`
	// Enables fuzzy entity extraction during classification.
	EnableFuzzyExtraction *bool `pulumi:"enableFuzzyExtraction"`
	// The collection of entity entries associated with the entity type.
	// Structure is documented below.
	Entities []EntityTypeEntity `pulumi:"entities"`
	// Indicates the kind of entity type.
	// * KIND_MAP: Map entity types allow mapping of a group of synonyms to a reference value.
	// * KIND_LIST: List entity types contain a set of entries that do not map to reference values. However, list entity
	//   types can contain references to other entity types (with or without aliases).
	// * KIND_REGEXP: Regexp entity types allow to specify regular expressions in entries values.
	//   Possible values are `KIND_MAP`, `KIND_LIST`, and `KIND_REGEXP`.
	Kind *string `pulumi:"kind"`
	// The unique identifier of the entity type. Format: projects/<Project ID>/agent/entityTypes/<Entity type ID>.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
}

type EntityTypeState struct {
	// The name of this entity type to be displayed on the console.
	DisplayName pulumi.StringPtrInput
	// Enables fuzzy entity extraction during classification.
	EnableFuzzyExtraction pulumi.BoolPtrInput
	// The collection of entity entries associated with the entity type.
	// Structure is documented below.
	Entities EntityTypeEntityArrayInput
	// Indicates the kind of entity type.
	// * KIND_MAP: Map entity types allow mapping of a group of synonyms to a reference value.
	// * KIND_LIST: List entity types contain a set of entries that do not map to reference values. However, list entity
	//   types can contain references to other entity types (with or without aliases).
	// * KIND_REGEXP: Regexp entity types allow to specify regular expressions in entries values.
	//   Possible values are `KIND_MAP`, `KIND_LIST`, and `KIND_REGEXP`.
	Kind pulumi.StringPtrInput
	// The unique identifier of the entity type. Format: projects/<Project ID>/agent/entityTypes/<Entity type ID>.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
}

func (EntityTypeState) ElementType() reflect.Type {
	return reflect.TypeOf((*entityTypeState)(nil)).Elem()
}

type entityTypeArgs struct {
	// The name of this entity type to be displayed on the console.
	DisplayName string `pulumi:"displayName"`
	// Enables fuzzy entity extraction during classification.
	EnableFuzzyExtraction *bool `pulumi:"enableFuzzyExtraction"`
	// The collection of entity entries associated with the entity type.
	// Structure is documented below.
	Entities []EntityTypeEntity `pulumi:"entities"`
	// Indicates the kind of entity type.
	// * KIND_MAP: Map entity types allow mapping of a group of synonyms to a reference value.
	// * KIND_LIST: List entity types contain a set of entries that do not map to reference values. However, list entity
	//   types can contain references to other entity types (with or without aliases).
	// * KIND_REGEXP: Regexp entity types allow to specify regular expressions in entries values.
	//   Possible values are `KIND_MAP`, `KIND_LIST`, and `KIND_REGEXP`.
	Kind string `pulumi:"kind"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
}

// The set of arguments for constructing a EntityType resource.
type EntityTypeArgs struct {
	// The name of this entity type to be displayed on the console.
	DisplayName pulumi.StringInput
	// Enables fuzzy entity extraction during classification.
	EnableFuzzyExtraction pulumi.BoolPtrInput
	// The collection of entity entries associated with the entity type.
	// Structure is documented below.
	Entities EntityTypeEntityArrayInput
	// Indicates the kind of entity type.
	// * KIND_MAP: Map entity types allow mapping of a group of synonyms to a reference value.
	// * KIND_LIST: List entity types contain a set of entries that do not map to reference values. However, list entity
	//   types can contain references to other entity types (with or without aliases).
	// * KIND_REGEXP: Regexp entity types allow to specify regular expressions in entries values.
	//   Possible values are `KIND_MAP`, `KIND_LIST`, and `KIND_REGEXP`.
	Kind pulumi.StringInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
}

func (EntityTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*entityTypeArgs)(nil)).Elem()
}

type EntityTypeInput interface {
	pulumi.Input

	ToEntityTypeOutput() EntityTypeOutput
	ToEntityTypeOutputWithContext(ctx context.Context) EntityTypeOutput
}

func (*EntityType) ElementType() reflect.Type {
	return reflect.TypeOf((*EntityType)(nil))
}

func (i *EntityType) ToEntityTypeOutput() EntityTypeOutput {
	return i.ToEntityTypeOutputWithContext(context.Background())
}

func (i *EntityType) ToEntityTypeOutputWithContext(ctx context.Context) EntityTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EntityTypeOutput)
}

func (i *EntityType) ToEntityTypePtrOutput() EntityTypePtrOutput {
	return i.ToEntityTypePtrOutputWithContext(context.Background())
}

func (i *EntityType) ToEntityTypePtrOutputWithContext(ctx context.Context) EntityTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EntityTypePtrOutput)
}

type EntityTypePtrInput interface {
	pulumi.Input

	ToEntityTypePtrOutput() EntityTypePtrOutput
	ToEntityTypePtrOutputWithContext(ctx context.Context) EntityTypePtrOutput
}

type entityTypePtrType EntityTypeArgs

func (*entityTypePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EntityType)(nil))
}

func (i *entityTypePtrType) ToEntityTypePtrOutput() EntityTypePtrOutput {
	return i.ToEntityTypePtrOutputWithContext(context.Background())
}

func (i *entityTypePtrType) ToEntityTypePtrOutputWithContext(ctx context.Context) EntityTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EntityTypeOutput).ToEntityTypePtrOutput()
}

// EntityTypeArrayInput is an input type that accepts EntityTypeArray and EntityTypeArrayOutput values.
// You can construct a concrete instance of `EntityTypeArrayInput` via:
//
//          EntityTypeArray{ EntityTypeArgs{...} }
type EntityTypeArrayInput interface {
	pulumi.Input

	ToEntityTypeArrayOutput() EntityTypeArrayOutput
	ToEntityTypeArrayOutputWithContext(context.Context) EntityTypeArrayOutput
}

type EntityTypeArray []EntityTypeInput

func (EntityTypeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EntityType)(nil))
}

func (i EntityTypeArray) ToEntityTypeArrayOutput() EntityTypeArrayOutput {
	return i.ToEntityTypeArrayOutputWithContext(context.Background())
}

func (i EntityTypeArray) ToEntityTypeArrayOutputWithContext(ctx context.Context) EntityTypeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EntityTypeArrayOutput)
}

// EntityTypeMapInput is an input type that accepts EntityTypeMap and EntityTypeMapOutput values.
// You can construct a concrete instance of `EntityTypeMapInput` via:
//
//          EntityTypeMap{ "key": EntityTypeArgs{...} }
type EntityTypeMapInput interface {
	pulumi.Input

	ToEntityTypeMapOutput() EntityTypeMapOutput
	ToEntityTypeMapOutputWithContext(context.Context) EntityTypeMapOutput
}

type EntityTypeMap map[string]EntityTypeInput

func (EntityTypeMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]EntityType)(nil))
}

func (i EntityTypeMap) ToEntityTypeMapOutput() EntityTypeMapOutput {
	return i.ToEntityTypeMapOutputWithContext(context.Background())
}

func (i EntityTypeMap) ToEntityTypeMapOutputWithContext(ctx context.Context) EntityTypeMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EntityTypeMapOutput)
}

type EntityTypeOutput struct {
	*pulumi.OutputState
}

func (EntityTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EntityType)(nil))
}

func (o EntityTypeOutput) ToEntityTypeOutput() EntityTypeOutput {
	return o
}

func (o EntityTypeOutput) ToEntityTypeOutputWithContext(ctx context.Context) EntityTypeOutput {
	return o
}

func (o EntityTypeOutput) ToEntityTypePtrOutput() EntityTypePtrOutput {
	return o.ToEntityTypePtrOutputWithContext(context.Background())
}

func (o EntityTypeOutput) ToEntityTypePtrOutputWithContext(ctx context.Context) EntityTypePtrOutput {
	return o.ApplyT(func(v EntityType) *EntityType {
		return &v
	}).(EntityTypePtrOutput)
}

type EntityTypePtrOutput struct {
	*pulumi.OutputState
}

func (EntityTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EntityType)(nil))
}

func (o EntityTypePtrOutput) ToEntityTypePtrOutput() EntityTypePtrOutput {
	return o
}

func (o EntityTypePtrOutput) ToEntityTypePtrOutputWithContext(ctx context.Context) EntityTypePtrOutput {
	return o
}

type EntityTypeArrayOutput struct{ *pulumi.OutputState }

func (EntityTypeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EntityType)(nil))
}

func (o EntityTypeArrayOutput) ToEntityTypeArrayOutput() EntityTypeArrayOutput {
	return o
}

func (o EntityTypeArrayOutput) ToEntityTypeArrayOutputWithContext(ctx context.Context) EntityTypeArrayOutput {
	return o
}

func (o EntityTypeArrayOutput) Index(i pulumi.IntInput) EntityTypeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) EntityType {
		return vs[0].([]EntityType)[vs[1].(int)]
	}).(EntityTypeOutput)
}

type EntityTypeMapOutput struct{ *pulumi.OutputState }

func (EntityTypeMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]EntityType)(nil))
}

func (o EntityTypeMapOutput) ToEntityTypeMapOutput() EntityTypeMapOutput {
	return o
}

func (o EntityTypeMapOutput) ToEntityTypeMapOutputWithContext(ctx context.Context) EntityTypeMapOutput {
	return o
}

func (o EntityTypeMapOutput) MapIndex(k pulumi.StringInput) EntityTypeOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) EntityType {
		return vs[0].(map[string]EntityType)[vs[1].(string)]
	}).(EntityTypeOutput)
}

func init() {
	pulumi.RegisterOutputType(EntityTypeOutput{})
	pulumi.RegisterOutputType(EntityTypePtrOutput{})
	pulumi.RegisterOutputType(EntityTypeArrayOutput{})
	pulumi.RegisterOutputType(EntityTypeMapOutput{})
}
