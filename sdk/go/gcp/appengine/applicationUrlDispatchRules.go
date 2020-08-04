// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package appengine

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Rules to match an HTTP request and dispatch that request to a service.
//
// To get more information about ApplicationUrlDispatchRules, see:
//
// * [API documentation](https://cloud.google.com/appengine/docs/admin-api/reference/rest/v1/apps#UrlDispatchRule)
//
// ## Example Usage
// ### App Engine Application Url Dispatch Rules Basic
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-gcp/sdk/v3/go/gcp/appengine"
// 	"github.com/pulumi/pulumi-gcp/sdk/v3/go/gcp/storage"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		bucket, err := storage.NewBucket(ctx, "bucket", nil)
// 		if err != nil {
// 			return err
// 		}
// 		object, err := storage.NewBucketObject(ctx, "object", &storage.BucketObjectArgs{
// 			Bucket: bucket.Name,
// 			Source: pulumi.NewFileAsset("./test-fixtures/appengine/hello-world.zip"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		adminV3, err := appengine.NewStandardAppVersion(ctx, "adminV3", &appengine.StandardAppVersionArgs{
// 			VersionId: pulumi.String("v3"),
// 			Service:   pulumi.String("admin"),
// 			Runtime:   pulumi.String("nodejs10"),
// 			Entrypoint: &appengine.StandardAppVersionEntrypointArgs{
// 				Shell: pulumi.String("node ./app.js"),
// 			},
// 			Deployment: &appengine.StandardAppVersionDeploymentArgs{
// 				Zip: &appengine.StandardAppVersionDeploymentZipArgs{
// 					SourceUrl: pulumi.All(bucket.Name, object.Name).ApplyT(func(_args []interface{}) (string, error) {
// 						bucketName := _args[0].(string)
// 						objectName := _args[1].(string)
// 						return fmt.Sprintf("%v%v%v%v", "https://storage.googleapis.com/", bucketName, "/", objectName), nil
// 					}).(pulumi.StringOutput),
// 				},
// 			},
// 			EnvVariables: pulumi.StringMap{
// 				"port": pulumi.String("8080"),
// 			},
// 			NoopOnDestroy: pulumi.Bool(true),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = appengine.NewApplicationUrlDispatchRules(ctx, "webService", &appengine.ApplicationUrlDispatchRulesArgs{
// 			DispatchRules: appengine.ApplicationUrlDispatchRulesDispatchRuleArray{
// 				&appengine.ApplicationUrlDispatchRulesDispatchRuleArgs{
// 					Domain:  pulumi.String("*"),
// 					Path:    pulumi.String("/*"),
// 					Service: pulumi.String("default"),
// 				},
// 				&appengine.ApplicationUrlDispatchRulesDispatchRuleArgs{
// 					Domain:  pulumi.String("*"),
// 					Path:    pulumi.String("/admin/*"),
// 					Service: adminV3.Service,
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
type ApplicationUrlDispatchRules struct {
	pulumi.CustomResourceState

	// Rules to match an HTTP request and dispatch that request to a service.  Structure is documented below.
	DispatchRules ApplicationUrlDispatchRulesDispatchRuleArrayOutput `pulumi:"dispatchRules"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
}

// NewApplicationUrlDispatchRules registers a new resource with the given unique name, arguments, and options.
func NewApplicationUrlDispatchRules(ctx *pulumi.Context,
	name string, args *ApplicationUrlDispatchRulesArgs, opts ...pulumi.ResourceOption) (*ApplicationUrlDispatchRules, error) {
	if args == nil || args.DispatchRules == nil {
		return nil, errors.New("missing required argument 'DispatchRules'")
	}
	if args == nil {
		args = &ApplicationUrlDispatchRulesArgs{}
	}
	var resource ApplicationUrlDispatchRules
	err := ctx.RegisterResource("gcp:appengine/applicationUrlDispatchRules:ApplicationUrlDispatchRules", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApplicationUrlDispatchRules gets an existing ApplicationUrlDispatchRules resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApplicationUrlDispatchRules(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApplicationUrlDispatchRulesState, opts ...pulumi.ResourceOption) (*ApplicationUrlDispatchRules, error) {
	var resource ApplicationUrlDispatchRules
	err := ctx.ReadResource("gcp:appengine/applicationUrlDispatchRules:ApplicationUrlDispatchRules", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ApplicationUrlDispatchRules resources.
type applicationUrlDispatchRulesState struct {
	// Rules to match an HTTP request and dispatch that request to a service.  Structure is documented below.
	DispatchRules []ApplicationUrlDispatchRulesDispatchRule `pulumi:"dispatchRules"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
}

type ApplicationUrlDispatchRulesState struct {
	// Rules to match an HTTP request and dispatch that request to a service.  Structure is documented below.
	DispatchRules ApplicationUrlDispatchRulesDispatchRuleArrayInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
}

func (ApplicationUrlDispatchRulesState) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationUrlDispatchRulesState)(nil)).Elem()
}

type applicationUrlDispatchRulesArgs struct {
	// Rules to match an HTTP request and dispatch that request to a service.  Structure is documented below.
	DispatchRules []ApplicationUrlDispatchRulesDispatchRule `pulumi:"dispatchRules"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
}

// The set of arguments for constructing a ApplicationUrlDispatchRules resource.
type ApplicationUrlDispatchRulesArgs struct {
	// Rules to match an HTTP request and dispatch that request to a service.  Structure is documented below.
	DispatchRules ApplicationUrlDispatchRulesDispatchRuleArrayInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
}

func (ApplicationUrlDispatchRulesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationUrlDispatchRulesArgs)(nil)).Elem()
}
