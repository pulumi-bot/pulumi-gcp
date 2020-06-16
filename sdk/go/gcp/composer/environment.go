// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package composer

import (
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// An environment for running orchestration tasks.
//
// Environments run Apache Airflow software on Google infrastructure.
//
// To get more information about Environments, see:
//
// * [API documentation](https://cloud.google.com/composer/docs/reference/rest/)
// * How-to Guides
//     * [Official Documentation](https://cloud.google.com/composer/docs)
//     * [Configuring Shared VPC for Composer Environments](https://cloud.google.com/composer/docs/how-to/managing/configuring-shared-vpc)
// * [Apache Airflow Documentation](http://airflow.apache.org/)
//
// > **Warning:** We **STRONGLY** recommend  you read the [GCP guides](https://cloud.google.com/composer/docs/how-to)
//   as the Environment resource requires a long deployment process and involves several layers of GCP infrastructure,
//   including a Kubernetes Engine cluster, Cloud Storage, and Compute networking resources. Due to limitations of the API,
//   This provider will not be able to automatically find or manage many of these underlying resources. In particular:
//   * It can take up to one hour to create or update an environment resource. In addition, GCP may only detect some
//     errors in configuration when they are used (e.g. ~40-50 minutes into the creation process), and is prone to limited
//     error reporting. If you encounter confusing or uninformative errors, please verify your configuration is valid
//     against GCP Cloud Composer before filing bugs against this provider.
//   * **Environments create Google Cloud Storage buckets that do not get cleaned up automatically** on environment
//     deletion. [More about Composer's use of Cloud Storage](https://cloud.google.com/composer/docs/concepts/cloud-storage).
//
// ## Example Usage
//
// ### Basic Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v3/go/gcp/composer"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err = composer.NewEnvironment(ctx, "test", &composer.EnvironmentArgs{
// 			Region: pulumi.String("us-central1"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ### With Software (Airflow) Config
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v3/go/gcp/composer"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err = composer.NewEnvironment(ctx, "test", &composer.EnvironmentArgs{
// 			Config: &composer.EnvironmentConfigArgs{
// 				SoftwareConfig: &composer.EnvironmentConfigSoftwareConfigArgs{
// 					AirflowConfigOverrides: map[string]interface{}{
// 						"core-loadExample": "True",
// 					},
// 					EnvVariables: map[string]interface{}{
// 						"FOO": "bar",
// 					},
// 					PypiPackages: map[string]interface{}{
// 						"numpy": "",
// 						"scipy": "==1.1.0",
// 					},
// 				},
// 			},
// 			Region: pulumi.String("us-central1"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type Environment struct {
	pulumi.CustomResourceState

	// Configuration parameters for this environment  Structure is documented below.
	Config EnvironmentConfigOutput `pulumi:"config"`
	// User-defined labels for this environment. The labels map can contain
	// no more than 64 entries. Entries of the labels map are UTF8 strings
	// that comply with the following restrictions:
	// Label keys must be between 1 and 63 characters long and must conform
	// to the following regular expression: `a-z?`.
	// Label values must be between 0 and 63 characters long and must
	// conform to the regular expression `(a-z?)?`.
	// No more than 64 labels can be associated with a given environment.
	// Both keys and values must be <= 128 bytes in size.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// Name of the environment
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The location or Compute Engine region for the environment.
	Region pulumi.StringPtrOutput `pulumi:"region"`
}

// NewEnvironment registers a new resource with the given unique name, arguments, and options.
func NewEnvironment(ctx *pulumi.Context,
	name string, args *EnvironmentArgs, opts ...pulumi.ResourceOption) (*Environment, error) {
	if args == nil {
		args = &EnvironmentArgs{}
	}
	var resource Environment
	err := ctx.RegisterResource("gcp:composer/environment:Environment", name, args, &resource, opts...)
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
	err := ctx.ReadResource("gcp:composer/environment:Environment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Environment resources.
type environmentState struct {
	// Configuration parameters for this environment  Structure is documented below.
	Config *EnvironmentConfig `pulumi:"config"`
	// User-defined labels for this environment. The labels map can contain
	// no more than 64 entries. Entries of the labels map are UTF8 strings
	// that comply with the following restrictions:
	// Label keys must be between 1 and 63 characters long and must conform
	// to the following regular expression: `a-z?`.
	// Label values must be between 0 and 63 characters long and must
	// conform to the regular expression `(a-z?)?`.
	// No more than 64 labels can be associated with a given environment.
	// Both keys and values must be <= 128 bytes in size.
	Labels map[string]string `pulumi:"labels"`
	// Name of the environment
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The location or Compute Engine region for the environment.
	Region *string `pulumi:"region"`
}

type EnvironmentState struct {
	// Configuration parameters for this environment  Structure is documented below.
	Config EnvironmentConfigPtrInput
	// User-defined labels for this environment. The labels map can contain
	// no more than 64 entries. Entries of the labels map are UTF8 strings
	// that comply with the following restrictions:
	// Label keys must be between 1 and 63 characters long and must conform
	// to the following regular expression: `a-z?`.
	// Label values must be between 0 and 63 characters long and must
	// conform to the regular expression `(a-z?)?`.
	// No more than 64 labels can be associated with a given environment.
	// Both keys and values must be <= 128 bytes in size.
	Labels pulumi.StringMapInput
	// Name of the environment
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The location or Compute Engine region for the environment.
	Region pulumi.StringPtrInput
}

func (EnvironmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*environmentState)(nil)).Elem()
}

type environmentArgs struct {
	// Configuration parameters for this environment  Structure is documented below.
	Config *EnvironmentConfig `pulumi:"config"`
	// User-defined labels for this environment. The labels map can contain
	// no more than 64 entries. Entries of the labels map are UTF8 strings
	// that comply with the following restrictions:
	// Label keys must be between 1 and 63 characters long and must conform
	// to the following regular expression: `a-z?`.
	// Label values must be between 0 and 63 characters long and must
	// conform to the regular expression `(a-z?)?`.
	// No more than 64 labels can be associated with a given environment.
	// Both keys and values must be <= 128 bytes in size.
	Labels map[string]string `pulumi:"labels"`
	// Name of the environment
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The location or Compute Engine region for the environment.
	Region *string `pulumi:"region"`
}

// The set of arguments for constructing a Environment resource.
type EnvironmentArgs struct {
	// Configuration parameters for this environment  Structure is documented below.
	Config EnvironmentConfigPtrInput
	// User-defined labels for this environment. The labels map can contain
	// no more than 64 entries. Entries of the labels map are UTF8 strings
	// that comply with the following restrictions:
	// Label keys must be between 1 and 63 characters long and must conform
	// to the following regular expression: `a-z?`.
	// Label values must be between 0 and 63 characters long and must
	// conform to the regular expression `(a-z?)?`.
	// No more than 64 labels can be associated with a given environment.
	// Both keys and values must be <= 128 bytes in size.
	Labels pulumi.StringMapInput
	// Name of the environment
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The location or Compute Engine region for the environment.
	Region pulumi.StringPtrInput
}

func (EnvironmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*environmentArgs)(nil)).Elem()
}
