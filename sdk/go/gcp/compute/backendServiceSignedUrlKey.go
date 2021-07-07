// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A key for signing Cloud CDN signed URLs for Backend Services.
//
// To get more information about BackendServiceSignedUrlKey, see:
//
// * [API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/backendServices)
// * How-to Guides
//     * [Using Signed URLs](https://cloud.google.com/cdn/docs/using-signed-urls/)
//
// > **Warning:** All arguments including `keyValue` will be stored in the raw
// state as plain-text.
//
// ## Example Usage
// ### Backend Service Signed Url Key
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/compute"
// 	"github.com/pulumi/pulumi-random/sdk/v4/go/random"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		urlSignature, err := random.NewRandomId(ctx, "urlSignature", &random.RandomIdArgs{
// 			ByteLength: pulumi.Int(16),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		webserver, err := compute.NewInstanceTemplate(ctx, "webserver", &compute.InstanceTemplateArgs{
// 			MachineType: pulumi.String("e2-medium"),
// 			NetworkInterfaces: compute.InstanceTemplateNetworkInterfaceArray{
// 				&compute.InstanceTemplateNetworkInterfaceArgs{
// 					Network: pulumi.String("default"),
// 				},
// 			},
// 			Disks: compute.InstanceTemplateDiskArray{
// 				&compute.InstanceTemplateDiskArgs{
// 					SourceImage: pulumi.String("debian-cloud/debian-9"),
// 					AutoDelete:  pulumi.Bool(true),
// 					Boot:        pulumi.Bool(true),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		webservers, err := compute.NewInstanceGroupManager(ctx, "webservers", &compute.InstanceGroupManagerArgs{
// 			Versions: compute.InstanceGroupManagerVersionArray{
// 				&compute.InstanceGroupManagerVersionArgs{
// 					InstanceTemplate: webserver.ID(),
// 					Name:             pulumi.String("primary"),
// 				},
// 			},
// 			BaseInstanceName: pulumi.String("webserver"),
// 			Zone:             pulumi.String("us-central1-f"),
// 			TargetSize:       pulumi.Int(1),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = compute.NewHttpHealthCheck(ctx, "_default", &compute.HttpHealthCheckArgs{
// 			RequestPath:      pulumi.String("/"),
// 			CheckIntervalSec: pulumi.Int(1),
// 			TimeoutSec:       pulumi.Int(1),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleBackend, err := compute.NewBackendService(ctx, "exampleBackend", &compute.BackendServiceArgs{
// 			Description: pulumi.String("Our company website"),
// 			PortName:    pulumi.String("http"),
// 			Protocol:    pulumi.String("HTTP"),
// 			TimeoutSec:  pulumi.Int(10),
// 			EnableCdn:   pulumi.Bool(true),
// 			Backends: compute.BackendServiceBackendArray{
// 				&compute.BackendServiceBackendArgs{
// 					Group: webservers.InstanceGroup,
// 				},
// 			},
// 			HealthChecks: pulumi.String{
// 				_default.ID(),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = compute.NewBackendServiceSignedUrlKey(ctx, "backendKey", &compute.BackendServiceSignedUrlKeyArgs{
// 			KeyValue:       urlSignature.B64Url,
// 			BackendService: exampleBackend.Name,
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
// This resource does not support import.
type BackendServiceSignedUrlKey struct {
	pulumi.CustomResourceState

	// The backend service this signed URL key belongs.
	BackendService pulumi.StringOutput `pulumi:"backendService"`
	// 128-bit key value used for signing the URL. The key value must be a
	// valid RFC 4648 Section 5 base64url encoded string.
	// **Note**: This property is sensitive and will not be displayed in the plan.
	KeyValue pulumi.StringOutput `pulumi:"keyValue"`
	// Name of the signed URL key.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
}

// NewBackendServiceSignedUrlKey registers a new resource with the given unique name, arguments, and options.
func NewBackendServiceSignedUrlKey(ctx *pulumi.Context,
	name string, args *BackendServiceSignedUrlKeyArgs, opts ...pulumi.ResourceOption) (*BackendServiceSignedUrlKey, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BackendService == nil {
		return nil, errors.New("invalid value for required argument 'BackendService'")
	}
	if args.KeyValue == nil {
		return nil, errors.New("invalid value for required argument 'KeyValue'")
	}
	var resource BackendServiceSignedUrlKey
	err := ctx.RegisterResource("gcp:compute/backendServiceSignedUrlKey:BackendServiceSignedUrlKey", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBackendServiceSignedUrlKey gets an existing BackendServiceSignedUrlKey resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBackendServiceSignedUrlKey(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BackendServiceSignedUrlKeyState, opts ...pulumi.ResourceOption) (*BackendServiceSignedUrlKey, error) {
	var resource BackendServiceSignedUrlKey
	err := ctx.ReadResource("gcp:compute/backendServiceSignedUrlKey:BackendServiceSignedUrlKey", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BackendServiceSignedUrlKey resources.
type backendServiceSignedUrlKeyState struct {
	// The backend service this signed URL key belongs.
	BackendService *string `pulumi:"backendService"`
	// 128-bit key value used for signing the URL. The key value must be a
	// valid RFC 4648 Section 5 base64url encoded string.
	// **Note**: This property is sensitive and will not be displayed in the plan.
	KeyValue *string `pulumi:"keyValue"`
	// Name of the signed URL key.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
}

type BackendServiceSignedUrlKeyState struct {
	// The backend service this signed URL key belongs.
	BackendService pulumi.StringPtrInput
	// 128-bit key value used for signing the URL. The key value must be a
	// valid RFC 4648 Section 5 base64url encoded string.
	// **Note**: This property is sensitive and will not be displayed in the plan.
	KeyValue pulumi.StringPtrInput
	// Name of the signed URL key.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
}

func (BackendServiceSignedUrlKeyState) ElementType() reflect.Type {
	return reflect.TypeOf((*backendServiceSignedUrlKeyState)(nil)).Elem()
}

type backendServiceSignedUrlKeyArgs struct {
	// The backend service this signed URL key belongs.
	BackendService string `pulumi:"backendService"`
	// 128-bit key value used for signing the URL. The key value must be a
	// valid RFC 4648 Section 5 base64url encoded string.
	// **Note**: This property is sensitive and will not be displayed in the plan.
	KeyValue string `pulumi:"keyValue"`
	// Name of the signed URL key.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
}

// The set of arguments for constructing a BackendServiceSignedUrlKey resource.
type BackendServiceSignedUrlKeyArgs struct {
	// The backend service this signed URL key belongs.
	BackendService pulumi.StringInput
	// 128-bit key value used for signing the URL. The key value must be a
	// valid RFC 4648 Section 5 base64url encoded string.
	// **Note**: This property is sensitive and will not be displayed in the plan.
	KeyValue pulumi.StringInput
	// Name of the signed URL key.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
}

func (BackendServiceSignedUrlKeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*backendServiceSignedUrlKeyArgs)(nil)).Elem()
}

type BackendServiceSignedUrlKeyInput interface {
	pulumi.Input

	ToBackendServiceSignedUrlKeyOutput() BackendServiceSignedUrlKeyOutput
	ToBackendServiceSignedUrlKeyOutputWithContext(ctx context.Context) BackendServiceSignedUrlKeyOutput
}

func (*BackendServiceSignedUrlKey) ElementType() reflect.Type {
	return reflect.TypeOf((*BackendServiceSignedUrlKey)(nil))
}

func (i *BackendServiceSignedUrlKey) ToBackendServiceSignedUrlKeyOutput() BackendServiceSignedUrlKeyOutput {
	return i.ToBackendServiceSignedUrlKeyOutputWithContext(context.Background())
}

func (i *BackendServiceSignedUrlKey) ToBackendServiceSignedUrlKeyOutputWithContext(ctx context.Context) BackendServiceSignedUrlKeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackendServiceSignedUrlKeyOutput)
}

func (i *BackendServiceSignedUrlKey) ToBackendServiceSignedUrlKeyPtrOutput() BackendServiceSignedUrlKeyPtrOutput {
	return i.ToBackendServiceSignedUrlKeyPtrOutputWithContext(context.Background())
}

func (i *BackendServiceSignedUrlKey) ToBackendServiceSignedUrlKeyPtrOutputWithContext(ctx context.Context) BackendServiceSignedUrlKeyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackendServiceSignedUrlKeyPtrOutput)
}

type BackendServiceSignedUrlKeyPtrInput interface {
	pulumi.Input

	ToBackendServiceSignedUrlKeyPtrOutput() BackendServiceSignedUrlKeyPtrOutput
	ToBackendServiceSignedUrlKeyPtrOutputWithContext(ctx context.Context) BackendServiceSignedUrlKeyPtrOutput
}

type backendServiceSignedUrlKeyPtrType BackendServiceSignedUrlKeyArgs

func (*backendServiceSignedUrlKeyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**BackendServiceSignedUrlKey)(nil))
}

func (i *backendServiceSignedUrlKeyPtrType) ToBackendServiceSignedUrlKeyPtrOutput() BackendServiceSignedUrlKeyPtrOutput {
	return i.ToBackendServiceSignedUrlKeyPtrOutputWithContext(context.Background())
}

func (i *backendServiceSignedUrlKeyPtrType) ToBackendServiceSignedUrlKeyPtrOutputWithContext(ctx context.Context) BackendServiceSignedUrlKeyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackendServiceSignedUrlKeyPtrOutput)
}

// BackendServiceSignedUrlKeyArrayInput is an input type that accepts BackendServiceSignedUrlKeyArray and BackendServiceSignedUrlKeyArrayOutput values.
// You can construct a concrete instance of `BackendServiceSignedUrlKeyArrayInput` via:
//
//          BackendServiceSignedUrlKeyArray{ BackendServiceSignedUrlKeyArgs{...} }
type BackendServiceSignedUrlKeyArrayInput interface {
	pulumi.Input

	ToBackendServiceSignedUrlKeyArrayOutput() BackendServiceSignedUrlKeyArrayOutput
	ToBackendServiceSignedUrlKeyArrayOutputWithContext(context.Context) BackendServiceSignedUrlKeyArrayOutput
}

type BackendServiceSignedUrlKeyArray []BackendServiceSignedUrlKeyInput

func (BackendServiceSignedUrlKeyArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*BackendServiceSignedUrlKey)(nil))
}

func (i BackendServiceSignedUrlKeyArray) ToBackendServiceSignedUrlKeyArrayOutput() BackendServiceSignedUrlKeyArrayOutput {
	return i.ToBackendServiceSignedUrlKeyArrayOutputWithContext(context.Background())
}

func (i BackendServiceSignedUrlKeyArray) ToBackendServiceSignedUrlKeyArrayOutputWithContext(ctx context.Context) BackendServiceSignedUrlKeyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackendServiceSignedUrlKeyArrayOutput)
}

// BackendServiceSignedUrlKeyMapInput is an input type that accepts BackendServiceSignedUrlKeyMap and BackendServiceSignedUrlKeyMapOutput values.
// You can construct a concrete instance of `BackendServiceSignedUrlKeyMapInput` via:
//
//          BackendServiceSignedUrlKeyMap{ "key": BackendServiceSignedUrlKeyArgs{...} }
type BackendServiceSignedUrlKeyMapInput interface {
	pulumi.Input

	ToBackendServiceSignedUrlKeyMapOutput() BackendServiceSignedUrlKeyMapOutput
	ToBackendServiceSignedUrlKeyMapOutputWithContext(context.Context) BackendServiceSignedUrlKeyMapOutput
}

type BackendServiceSignedUrlKeyMap map[string]BackendServiceSignedUrlKeyInput

func (BackendServiceSignedUrlKeyMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*BackendServiceSignedUrlKey)(nil))
}

func (i BackendServiceSignedUrlKeyMap) ToBackendServiceSignedUrlKeyMapOutput() BackendServiceSignedUrlKeyMapOutput {
	return i.ToBackendServiceSignedUrlKeyMapOutputWithContext(context.Background())
}

func (i BackendServiceSignedUrlKeyMap) ToBackendServiceSignedUrlKeyMapOutputWithContext(ctx context.Context) BackendServiceSignedUrlKeyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackendServiceSignedUrlKeyMapOutput)
}

type BackendServiceSignedUrlKeyOutput struct {
	*pulumi.OutputState
}

func (BackendServiceSignedUrlKeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BackendServiceSignedUrlKey)(nil))
}

func (o BackendServiceSignedUrlKeyOutput) ToBackendServiceSignedUrlKeyOutput() BackendServiceSignedUrlKeyOutput {
	return o
}

func (o BackendServiceSignedUrlKeyOutput) ToBackendServiceSignedUrlKeyOutputWithContext(ctx context.Context) BackendServiceSignedUrlKeyOutput {
	return o
}

func (o BackendServiceSignedUrlKeyOutput) ToBackendServiceSignedUrlKeyPtrOutput() BackendServiceSignedUrlKeyPtrOutput {
	return o.ToBackendServiceSignedUrlKeyPtrOutputWithContext(context.Background())
}

func (o BackendServiceSignedUrlKeyOutput) ToBackendServiceSignedUrlKeyPtrOutputWithContext(ctx context.Context) BackendServiceSignedUrlKeyPtrOutput {
	return o.ApplyT(func(v BackendServiceSignedUrlKey) *BackendServiceSignedUrlKey {
		return &v
	}).(BackendServiceSignedUrlKeyPtrOutput)
}

type BackendServiceSignedUrlKeyPtrOutput struct {
	*pulumi.OutputState
}

func (BackendServiceSignedUrlKeyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BackendServiceSignedUrlKey)(nil))
}

func (o BackendServiceSignedUrlKeyPtrOutput) ToBackendServiceSignedUrlKeyPtrOutput() BackendServiceSignedUrlKeyPtrOutput {
	return o
}

func (o BackendServiceSignedUrlKeyPtrOutput) ToBackendServiceSignedUrlKeyPtrOutputWithContext(ctx context.Context) BackendServiceSignedUrlKeyPtrOutput {
	return o
}

type BackendServiceSignedUrlKeyArrayOutput struct{ *pulumi.OutputState }

func (BackendServiceSignedUrlKeyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]BackendServiceSignedUrlKey)(nil))
}

func (o BackendServiceSignedUrlKeyArrayOutput) ToBackendServiceSignedUrlKeyArrayOutput() BackendServiceSignedUrlKeyArrayOutput {
	return o
}

func (o BackendServiceSignedUrlKeyArrayOutput) ToBackendServiceSignedUrlKeyArrayOutputWithContext(ctx context.Context) BackendServiceSignedUrlKeyArrayOutput {
	return o
}

func (o BackendServiceSignedUrlKeyArrayOutput) Index(i pulumi.IntInput) BackendServiceSignedUrlKeyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) BackendServiceSignedUrlKey {
		return vs[0].([]BackendServiceSignedUrlKey)[vs[1].(int)]
	}).(BackendServiceSignedUrlKeyOutput)
}

type BackendServiceSignedUrlKeyMapOutput struct{ *pulumi.OutputState }

func (BackendServiceSignedUrlKeyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]BackendServiceSignedUrlKey)(nil))
}

func (o BackendServiceSignedUrlKeyMapOutput) ToBackendServiceSignedUrlKeyMapOutput() BackendServiceSignedUrlKeyMapOutput {
	return o
}

func (o BackendServiceSignedUrlKeyMapOutput) ToBackendServiceSignedUrlKeyMapOutputWithContext(ctx context.Context) BackendServiceSignedUrlKeyMapOutput {
	return o
}

func (o BackendServiceSignedUrlKeyMapOutput) MapIndex(k pulumi.StringInput) BackendServiceSignedUrlKeyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) BackendServiceSignedUrlKey {
		return vs[0].(map[string]BackendServiceSignedUrlKey)[vs[1].(string)]
	}).(BackendServiceSignedUrlKeyOutput)
}

func init() {
	pulumi.RegisterOutputType(BackendServiceSignedUrlKeyOutput{})
	pulumi.RegisterOutputType(BackendServiceSignedUrlKeyPtrOutput{})
	pulumi.RegisterOutputType(BackendServiceSignedUrlKeyArrayOutput{})
	pulumi.RegisterOutputType(BackendServiceSignedUrlKeyMapOutput{})
}
