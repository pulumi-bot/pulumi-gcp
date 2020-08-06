// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ml

import (
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Represents a machine learning solution.
//
// A model can have multiple versions, each of which is a deployed, trained model
// ready to receive prediction requests. The model itself is just a container.
//
// ## Example Usage
// ### Ml Model Basic
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v3/go/gcp/ml"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := ml.NewEngineModel(ctx, "_default", &ml.EngineModelArgs{
// 			Description: pulumi.String("My model"),
// 			Regions:     pulumi.String("us-central1"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Ml Model Full
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v3/go/gcp/ml"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := ml.NewEngineModel(ctx, "_default", &ml.EngineModelArgs{
// 			Description: pulumi.String("My model"),
// 			Labels: pulumi.StringMap{
// 				"my_model": pulumi.String("foo"),
// 			},
// 			OnlinePredictionConsoleLogging: pulumi.Bool(true),
// 			OnlinePredictionLogging:        pulumi.Bool(true),
// 			Regions:                        pulumi.String("us-central1"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type EngineModel struct {
	pulumi.CustomResourceState

	// The default version of the model. This version will be used to handle
	// prediction requests that do not specify a version.  Structure is documented below.
	DefaultVersion EngineModelDefaultVersionPtrOutput `pulumi:"defaultVersion"`
	// The description specified for the model when it was created.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// One or more labels that you can add, to organize your models.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// The name specified for the version when it was created.
	Name pulumi.StringOutput `pulumi:"name"`
	// If true, online prediction nodes send stderr and stdout streams to Stackdriver Logging
	OnlinePredictionConsoleLogging pulumi.BoolPtrOutput `pulumi:"onlinePredictionConsoleLogging"`
	// If true, online prediction access logs are sent to StackDriver Logging.
	OnlinePredictionLogging pulumi.BoolPtrOutput `pulumi:"onlinePredictionLogging"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The list of regions where the model is going to be deployed.
	// Currently only one region per model is supported
	Regions pulumi.StringPtrOutput `pulumi:"regions"`
}

// NewEngineModel registers a new resource with the given unique name, arguments, and options.
func NewEngineModel(ctx *pulumi.Context,
	name string, args *EngineModelArgs, opts ...pulumi.ResourceOption) (*EngineModel, error) {
	if args == nil {
		args = &EngineModelArgs{}
	}
	var resource EngineModel
	err := ctx.RegisterResource("gcp:ml/engineModel:EngineModel", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEngineModel gets an existing EngineModel resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEngineModel(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EngineModelState, opts ...pulumi.ResourceOption) (*EngineModel, error) {
	var resource EngineModel
	err := ctx.ReadResource("gcp:ml/engineModel:EngineModel", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EngineModel resources.
type engineModelState struct {
	// The default version of the model. This version will be used to handle
	// prediction requests that do not specify a version.  Structure is documented below.
	DefaultVersion *EngineModelDefaultVersion `pulumi:"defaultVersion"`
	// The description specified for the model when it was created.
	Description *string `pulumi:"description"`
	// One or more labels that you can add, to organize your models.
	Labels map[string]string `pulumi:"labels"`
	// The name specified for the version when it was created.
	Name *string `pulumi:"name"`
	// If true, online prediction nodes send stderr and stdout streams to Stackdriver Logging
	OnlinePredictionConsoleLogging *bool `pulumi:"onlinePredictionConsoleLogging"`
	// If true, online prediction access logs are sent to StackDriver Logging.
	OnlinePredictionLogging *bool `pulumi:"onlinePredictionLogging"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The list of regions where the model is going to be deployed.
	// Currently only one region per model is supported
	Regions *string `pulumi:"regions"`
}

type EngineModelState struct {
	// The default version of the model. This version will be used to handle
	// prediction requests that do not specify a version.  Structure is documented below.
	DefaultVersion EngineModelDefaultVersionPtrInput
	// The description specified for the model when it was created.
	Description pulumi.StringPtrInput
	// One or more labels that you can add, to organize your models.
	Labels pulumi.StringMapInput
	// The name specified for the version when it was created.
	Name pulumi.StringPtrInput
	// If true, online prediction nodes send stderr and stdout streams to Stackdriver Logging
	OnlinePredictionConsoleLogging pulumi.BoolPtrInput
	// If true, online prediction access logs are sent to StackDriver Logging.
	OnlinePredictionLogging pulumi.BoolPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The list of regions where the model is going to be deployed.
	// Currently only one region per model is supported
	Regions pulumi.StringPtrInput
}

func (EngineModelState) ElementType() reflect.Type {
	return reflect.TypeOf((*engineModelState)(nil)).Elem()
}

type engineModelArgs struct {
	// The default version of the model. This version will be used to handle
	// prediction requests that do not specify a version.  Structure is documented below.
	DefaultVersion *EngineModelDefaultVersion `pulumi:"defaultVersion"`
	// The description specified for the model when it was created.
	Description *string `pulumi:"description"`
	// One or more labels that you can add, to organize your models.
	Labels map[string]string `pulumi:"labels"`
	// The name specified for the version when it was created.
	Name *string `pulumi:"name"`
	// If true, online prediction nodes send stderr and stdout streams to Stackdriver Logging
	OnlinePredictionConsoleLogging *bool `pulumi:"onlinePredictionConsoleLogging"`
	// If true, online prediction access logs are sent to StackDriver Logging.
	OnlinePredictionLogging *bool `pulumi:"onlinePredictionLogging"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The list of regions where the model is going to be deployed.
	// Currently only one region per model is supported
	Regions *string `pulumi:"regions"`
}

// The set of arguments for constructing a EngineModel resource.
type EngineModelArgs struct {
	// The default version of the model. This version will be used to handle
	// prediction requests that do not specify a version.  Structure is documented below.
	DefaultVersion EngineModelDefaultVersionPtrInput
	// The description specified for the model when it was created.
	Description pulumi.StringPtrInput
	// One or more labels that you can add, to organize your models.
	Labels pulumi.StringMapInput
	// The name specified for the version when it was created.
	Name pulumi.StringPtrInput
	// If true, online prediction nodes send stderr and stdout streams to Stackdriver Logging
	OnlinePredictionConsoleLogging pulumi.BoolPtrInput
	// If true, online prediction access logs are sent to StackDriver Logging.
	OnlinePredictionLogging pulumi.BoolPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The list of regions where the model is going to be deployed.
	// Currently only one region per model is supported
	Regions pulumi.StringPtrInput
}

func (EngineModelArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*engineModelArgs)(nil)).Elem()
}
