// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package dataloss

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Allows creation of custom info types.
//
// To get more information about StoredInfoType, see:
//
// * [API documentation](https://cloud.google.com/dlp/docs/reference/rest/v2/projects.storedInfoTypes)
// * How-to Guides
//     * [Official Documentation](https://cloud.google.com/dlp/docs/creating-stored-infotypes)
//
// ## Example Usage
// ### Dlp Stored Info Type Dictionary
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/dataloss"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := dataloss.NewPreventionStoredInfoType(ctx, "dictionary", &dataloss.PreventionStoredInfoTypeArgs{
// 			Description: pulumi.String("Description"),
// 			Dictionary: &dataloss.PreventionStoredInfoTypeDictionaryArgs{
// 				WordList: &dataloss.PreventionStoredInfoTypeDictionaryWordListArgs{
// 					Words: pulumi.StringArray{
// 						pulumi.String("word"),
// 						pulumi.String("word2"),
// 					},
// 				},
// 			},
// 			DisplayName: pulumi.String("Displayname"),
// 			Parent:      pulumi.String("projects/my-project-name"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Dlp Stored Info Type Large Custom Dictionary
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/dataloss"
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/storage"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		bucket, err := storage.NewBucket(ctx, "bucket", &storage.BucketArgs{
// 			ForceDestroy: pulumi.Bool(true),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		object, err := storage.NewBucketObject(ctx, "object", &storage.BucketObjectArgs{
// 			Bucket: bucket.Name,
// 			Source: pulumi.NewFileAsset("./test-fixtures/dlp/words.txt"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = dataloss.NewPreventionStoredInfoType(ctx, "large", &dataloss.PreventionStoredInfoTypeArgs{
// 			Parent:      pulumi.String("projects/my-project-name"),
// 			Description: pulumi.String("Description"),
// 			DisplayName: pulumi.String("Displayname"),
// 			LargeCustomDictionary: &dataloss.PreventionStoredInfoTypeLargeCustomDictionaryArgs{
// 				CloudStorageFileSet: &dataloss.PreventionStoredInfoTypeLargeCustomDictionaryCloudStorageFileSetArgs{
// 					Url: pulumi.All(bucket.Name, object.Name).ApplyT(func(_args []interface{}) (string, error) {
// 						bucketName := _args[0].(string)
// 						objectName := _args[1].(string)
// 						return fmt.Sprintf("%v%v%v%v", "gs://", bucketName, "/", objectName), nil
// 					}).(pulumi.StringOutput),
// 				},
// 				OutputPath: &dataloss.PreventionStoredInfoTypeLargeCustomDictionaryOutputPathArgs{
// 					Path: bucket.Name.ApplyT(func(name string) (string, error) {
// 						return fmt.Sprintf("%v%v%v", "gs://", name, "/output/dictionary.txt"), nil
// 					}).(pulumi.StringOutput),
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
// StoredInfoType can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import gcp:dataloss/preventionStoredInfoType:PreventionStoredInfoType default {{parent}}/storedInfoTypes/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:dataloss/preventionStoredInfoType:PreventionStoredInfoType default {{parent}}/{{name}}
// ```
type PreventionStoredInfoType struct {
	pulumi.CustomResourceState

	// A description of the info type.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Dictionary which defines the rule.
	// Structure is documented below.
	Dictionary PreventionStoredInfoTypeDictionaryPtrOutput `pulumi:"dictionary"`
	// User set display name of the info type.
	DisplayName pulumi.StringPtrOutput `pulumi:"displayName"`
	// Dictionary which defines the rule.
	// Structure is documented below.
	LargeCustomDictionary PreventionStoredInfoTypeLargeCustomDictionaryPtrOutput `pulumi:"largeCustomDictionary"`
	// Name describing the field.
	Name pulumi.StringOutput `pulumi:"name"`
	// The parent of the info type in any of the following formats:
	// * `projects/{{project}}`
	// * `projects/{{project}}/locations/{{location}}`
	// * `organizations/{{organization_id}}`
	// * `organizations/{{organization_id}}/locations/{{location}}`
	Parent pulumi.StringOutput `pulumi:"parent"`
	// Regular expression which defines the rule.
	// Structure is documented below.
	Regex PreventionStoredInfoTypeRegexPtrOutput `pulumi:"regex"`
}

// NewPreventionStoredInfoType registers a new resource with the given unique name, arguments, and options.
func NewPreventionStoredInfoType(ctx *pulumi.Context,
	name string, args *PreventionStoredInfoTypeArgs, opts ...pulumi.ResourceOption) (*PreventionStoredInfoType, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Parent == nil {
		return nil, errors.New("invalid value for required argument 'Parent'")
	}
	var resource PreventionStoredInfoType
	err := ctx.RegisterResource("gcp:dataloss/preventionStoredInfoType:PreventionStoredInfoType", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPreventionStoredInfoType gets an existing PreventionStoredInfoType resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPreventionStoredInfoType(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PreventionStoredInfoTypeState, opts ...pulumi.ResourceOption) (*PreventionStoredInfoType, error) {
	var resource PreventionStoredInfoType
	err := ctx.ReadResource("gcp:dataloss/preventionStoredInfoType:PreventionStoredInfoType", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PreventionStoredInfoType resources.
type preventionStoredInfoTypeState struct {
	// A description of the info type.
	Description *string `pulumi:"description"`
	// Dictionary which defines the rule.
	// Structure is documented below.
	Dictionary *PreventionStoredInfoTypeDictionary `pulumi:"dictionary"`
	// User set display name of the info type.
	DisplayName *string `pulumi:"displayName"`
	// Dictionary which defines the rule.
	// Structure is documented below.
	LargeCustomDictionary *PreventionStoredInfoTypeLargeCustomDictionary `pulumi:"largeCustomDictionary"`
	// Name describing the field.
	Name *string `pulumi:"name"`
	// The parent of the info type in any of the following formats:
	// * `projects/{{project}}`
	// * `projects/{{project}}/locations/{{location}}`
	// * `organizations/{{organization_id}}`
	// * `organizations/{{organization_id}}/locations/{{location}}`
	Parent *string `pulumi:"parent"`
	// Regular expression which defines the rule.
	// Structure is documented below.
	Regex *PreventionStoredInfoTypeRegex `pulumi:"regex"`
}

type PreventionStoredInfoTypeState struct {
	// A description of the info type.
	Description pulumi.StringPtrInput
	// Dictionary which defines the rule.
	// Structure is documented below.
	Dictionary PreventionStoredInfoTypeDictionaryPtrInput
	// User set display name of the info type.
	DisplayName pulumi.StringPtrInput
	// Dictionary which defines the rule.
	// Structure is documented below.
	LargeCustomDictionary PreventionStoredInfoTypeLargeCustomDictionaryPtrInput
	// Name describing the field.
	Name pulumi.StringPtrInput
	// The parent of the info type in any of the following formats:
	// * `projects/{{project}}`
	// * `projects/{{project}}/locations/{{location}}`
	// * `organizations/{{organization_id}}`
	// * `organizations/{{organization_id}}/locations/{{location}}`
	Parent pulumi.StringPtrInput
	// Regular expression which defines the rule.
	// Structure is documented below.
	Regex PreventionStoredInfoTypeRegexPtrInput
}

func (PreventionStoredInfoTypeState) ElementType() reflect.Type {
	return reflect.TypeOf((*preventionStoredInfoTypeState)(nil)).Elem()
}

type preventionStoredInfoTypeArgs struct {
	// A description of the info type.
	Description *string `pulumi:"description"`
	// Dictionary which defines the rule.
	// Structure is documented below.
	Dictionary *PreventionStoredInfoTypeDictionary `pulumi:"dictionary"`
	// User set display name of the info type.
	DisplayName *string `pulumi:"displayName"`
	// Dictionary which defines the rule.
	// Structure is documented below.
	LargeCustomDictionary *PreventionStoredInfoTypeLargeCustomDictionary `pulumi:"largeCustomDictionary"`
	// The parent of the info type in any of the following formats:
	// * `projects/{{project}}`
	// * `projects/{{project}}/locations/{{location}}`
	// * `organizations/{{organization_id}}`
	// * `organizations/{{organization_id}}/locations/{{location}}`
	Parent string `pulumi:"parent"`
	// Regular expression which defines the rule.
	// Structure is documented below.
	Regex *PreventionStoredInfoTypeRegex `pulumi:"regex"`
}

// The set of arguments for constructing a PreventionStoredInfoType resource.
type PreventionStoredInfoTypeArgs struct {
	// A description of the info type.
	Description pulumi.StringPtrInput
	// Dictionary which defines the rule.
	// Structure is documented below.
	Dictionary PreventionStoredInfoTypeDictionaryPtrInput
	// User set display name of the info type.
	DisplayName pulumi.StringPtrInput
	// Dictionary which defines the rule.
	// Structure is documented below.
	LargeCustomDictionary PreventionStoredInfoTypeLargeCustomDictionaryPtrInput
	// The parent of the info type in any of the following formats:
	// * `projects/{{project}}`
	// * `projects/{{project}}/locations/{{location}}`
	// * `organizations/{{organization_id}}`
	// * `organizations/{{organization_id}}/locations/{{location}}`
	Parent pulumi.StringInput
	// Regular expression which defines the rule.
	// Structure is documented below.
	Regex PreventionStoredInfoTypeRegexPtrInput
}

func (PreventionStoredInfoTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*preventionStoredInfoTypeArgs)(nil)).Elem()
}

type PreventionStoredInfoTypeInput interface {
	pulumi.Input

	ToPreventionStoredInfoTypeOutput() PreventionStoredInfoTypeOutput
	ToPreventionStoredInfoTypeOutputWithContext(ctx context.Context) PreventionStoredInfoTypeOutput
}

func (*PreventionStoredInfoType) ElementType() reflect.Type {
	return reflect.TypeOf((*PreventionStoredInfoType)(nil))
}

func (i *PreventionStoredInfoType) ToPreventionStoredInfoTypeOutput() PreventionStoredInfoTypeOutput {
	return i.ToPreventionStoredInfoTypeOutputWithContext(context.Background())
}

func (i *PreventionStoredInfoType) ToPreventionStoredInfoTypeOutputWithContext(ctx context.Context) PreventionStoredInfoTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PreventionStoredInfoTypeOutput)
}

func (i *PreventionStoredInfoType) ToPreventionStoredInfoTypePtrOutput() PreventionStoredInfoTypePtrOutput {
	return i.ToPreventionStoredInfoTypePtrOutputWithContext(context.Background())
}

func (i *PreventionStoredInfoType) ToPreventionStoredInfoTypePtrOutputWithContext(ctx context.Context) PreventionStoredInfoTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PreventionStoredInfoTypePtrOutput)
}

type PreventionStoredInfoTypePtrInput interface {
	pulumi.Input

	ToPreventionStoredInfoTypePtrOutput() PreventionStoredInfoTypePtrOutput
	ToPreventionStoredInfoTypePtrOutputWithContext(ctx context.Context) PreventionStoredInfoTypePtrOutput
}

type preventionStoredInfoTypePtrType PreventionStoredInfoTypeArgs

func (*preventionStoredInfoTypePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PreventionStoredInfoType)(nil))
}

func (i *preventionStoredInfoTypePtrType) ToPreventionStoredInfoTypePtrOutput() PreventionStoredInfoTypePtrOutput {
	return i.ToPreventionStoredInfoTypePtrOutputWithContext(context.Background())
}

func (i *preventionStoredInfoTypePtrType) ToPreventionStoredInfoTypePtrOutputWithContext(ctx context.Context) PreventionStoredInfoTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PreventionStoredInfoTypeOutput).ToPreventionStoredInfoTypePtrOutput()
}

// PreventionStoredInfoTypeArrayInput is an input type that accepts PreventionStoredInfoTypeArray and PreventionStoredInfoTypeArrayOutput values.
// You can construct a concrete instance of `PreventionStoredInfoTypeArrayInput` via:
//
//          PreventionStoredInfoTypeArray{ PreventionStoredInfoTypeArgs{...} }
type PreventionStoredInfoTypeArrayInput interface {
	pulumi.Input

	ToPreventionStoredInfoTypeArrayOutput() PreventionStoredInfoTypeArrayOutput
	ToPreventionStoredInfoTypeArrayOutputWithContext(context.Context) PreventionStoredInfoTypeArrayOutput
}

type PreventionStoredInfoTypeArray []PreventionStoredInfoTypeInput

func (PreventionStoredInfoTypeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PreventionStoredInfoType)(nil))
}

func (i PreventionStoredInfoTypeArray) ToPreventionStoredInfoTypeArrayOutput() PreventionStoredInfoTypeArrayOutput {
	return i.ToPreventionStoredInfoTypeArrayOutputWithContext(context.Background())
}

func (i PreventionStoredInfoTypeArray) ToPreventionStoredInfoTypeArrayOutputWithContext(ctx context.Context) PreventionStoredInfoTypeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PreventionStoredInfoTypeArrayOutput)
}

// PreventionStoredInfoTypeMapInput is an input type that accepts PreventionStoredInfoTypeMap and PreventionStoredInfoTypeMapOutput values.
// You can construct a concrete instance of `PreventionStoredInfoTypeMapInput` via:
//
//          PreventionStoredInfoTypeMap{ "key": PreventionStoredInfoTypeArgs{...} }
type PreventionStoredInfoTypeMapInput interface {
	pulumi.Input

	ToPreventionStoredInfoTypeMapOutput() PreventionStoredInfoTypeMapOutput
	ToPreventionStoredInfoTypeMapOutputWithContext(context.Context) PreventionStoredInfoTypeMapOutput
}

type PreventionStoredInfoTypeMap map[string]PreventionStoredInfoTypeInput

func (PreventionStoredInfoTypeMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]PreventionStoredInfoType)(nil))
}

func (i PreventionStoredInfoTypeMap) ToPreventionStoredInfoTypeMapOutput() PreventionStoredInfoTypeMapOutput {
	return i.ToPreventionStoredInfoTypeMapOutputWithContext(context.Background())
}

func (i PreventionStoredInfoTypeMap) ToPreventionStoredInfoTypeMapOutputWithContext(ctx context.Context) PreventionStoredInfoTypeMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PreventionStoredInfoTypeMapOutput)
}

type PreventionStoredInfoTypeOutput struct {
	*pulumi.OutputState
}

func (PreventionStoredInfoTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PreventionStoredInfoType)(nil))
}

func (o PreventionStoredInfoTypeOutput) ToPreventionStoredInfoTypeOutput() PreventionStoredInfoTypeOutput {
	return o
}

func (o PreventionStoredInfoTypeOutput) ToPreventionStoredInfoTypeOutputWithContext(ctx context.Context) PreventionStoredInfoTypeOutput {
	return o
}

func (o PreventionStoredInfoTypeOutput) ToPreventionStoredInfoTypePtrOutput() PreventionStoredInfoTypePtrOutput {
	return o.ToPreventionStoredInfoTypePtrOutputWithContext(context.Background())
}

func (o PreventionStoredInfoTypeOutput) ToPreventionStoredInfoTypePtrOutputWithContext(ctx context.Context) PreventionStoredInfoTypePtrOutput {
	return o.ApplyT(func(v PreventionStoredInfoType) *PreventionStoredInfoType {
		return &v
	}).(PreventionStoredInfoTypePtrOutput)
}

type PreventionStoredInfoTypePtrOutput struct {
	*pulumi.OutputState
}

func (PreventionStoredInfoTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PreventionStoredInfoType)(nil))
}

func (o PreventionStoredInfoTypePtrOutput) ToPreventionStoredInfoTypePtrOutput() PreventionStoredInfoTypePtrOutput {
	return o
}

func (o PreventionStoredInfoTypePtrOutput) ToPreventionStoredInfoTypePtrOutputWithContext(ctx context.Context) PreventionStoredInfoTypePtrOutput {
	return o
}

type PreventionStoredInfoTypeArrayOutput struct{ *pulumi.OutputState }

func (PreventionStoredInfoTypeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PreventionStoredInfoType)(nil))
}

func (o PreventionStoredInfoTypeArrayOutput) ToPreventionStoredInfoTypeArrayOutput() PreventionStoredInfoTypeArrayOutput {
	return o
}

func (o PreventionStoredInfoTypeArrayOutput) ToPreventionStoredInfoTypeArrayOutputWithContext(ctx context.Context) PreventionStoredInfoTypeArrayOutput {
	return o
}

func (o PreventionStoredInfoTypeArrayOutput) Index(i pulumi.IntInput) PreventionStoredInfoTypeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PreventionStoredInfoType {
		return vs[0].([]PreventionStoredInfoType)[vs[1].(int)]
	}).(PreventionStoredInfoTypeOutput)
}

type PreventionStoredInfoTypeMapOutput struct{ *pulumi.OutputState }

func (PreventionStoredInfoTypeMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]PreventionStoredInfoType)(nil))
}

func (o PreventionStoredInfoTypeMapOutput) ToPreventionStoredInfoTypeMapOutput() PreventionStoredInfoTypeMapOutput {
	return o
}

func (o PreventionStoredInfoTypeMapOutput) ToPreventionStoredInfoTypeMapOutputWithContext(ctx context.Context) PreventionStoredInfoTypeMapOutput {
	return o
}

func (o PreventionStoredInfoTypeMapOutput) MapIndex(k pulumi.StringInput) PreventionStoredInfoTypeOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) PreventionStoredInfoType {
		return vs[0].(map[string]PreventionStoredInfoType)[vs[1].(string)]
	}).(PreventionStoredInfoTypeOutput)
}

func init() {
	pulumi.RegisterOutputType(PreventionStoredInfoTypeOutput{})
	pulumi.RegisterOutputType(PreventionStoredInfoTypePtrOutput{})
	pulumi.RegisterOutputType(PreventionStoredInfoTypeArrayOutput{})
	pulumi.RegisterOutputType(PreventionStoredInfoTypeMapOutput{})
}
