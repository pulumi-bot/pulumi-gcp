// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package storage

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The hmacKeys resource represents an HMAC key within Cloud Storage. The resource
// consists of a secret and HMAC key metadata. HMAC keys can be used as credentials
// for service accounts.
//
// To get more information about HmacKey, see:
//
// * [API documentation](https://cloud.google.com/storage/docs/json_api/v1/projects/hmacKeys)
// * How-to Guides
//     * [Official Documentation](https://cloud.google.com/storage/docs/authentication/managing-hmackeys)
//
// > **Warning:** All arguments including the `secret` value will be stored in the raw
// state as plain-text. [Read more about secrets in state](https://www.pulumi.com/docs/intro/concepts/programming-model/#secrets).
// On import, the `secret` value will not be retrieved.
//
// > **Warning:** All arguments including `secret` will be stored in the raw
// state as plain-text. [Read more about secrets in state](https://www.pulumi.com/docs/intro/concepts/programming-model/#secrets).
//
// ## Example Usage
// ### Storage Hmac Key
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/serviceAccount"
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/storage"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		serviceAccount, err := serviceAccount.NewAccount(ctx, "serviceAccount", &serviceAccount.AccountArgs{
// 			AccountId: pulumi.String("my-svc-acc"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = storage.NewHmacKey(ctx, "key", &storage.HmacKeyArgs{
// 			ServiceAccountEmail: serviceAccount.Email,
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
// HmacKey can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import gcp:storage/hmacKey:HmacKey default projects/{{project}}/hmacKeys/{{access_id}}
// ```
//
// ```sh
//  $ pulumi import gcp:storage/hmacKey:HmacKey default {{project}}/{{access_id}}
// ```
//
// ```sh
//  $ pulumi import gcp:storage/hmacKey:HmacKey default {{access_id}}
// ```
type HmacKey struct {
	pulumi.CustomResourceState

	// The access ID of the HMAC Key.
	AccessId pulumi.StringOutput `pulumi:"accessId"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// HMAC secret key material.
	Secret pulumi.StringOutput `pulumi:"secret"`
	// The email address of the key's associated service account.
	ServiceAccountEmail pulumi.StringOutput `pulumi:"serviceAccountEmail"`
	// The state of the key. Can be set to one of ACTIVE, INACTIVE.
	// Default value is `ACTIVE`.
	// Possible values are `ACTIVE` and `INACTIVE`.
	State pulumi.StringPtrOutput `pulumi:"state"`
	// 'The creation time of the HMAC key in RFC 3339 format. '
	TimeCreated pulumi.StringOutput `pulumi:"timeCreated"`
	// 'The last modification time of the HMAC key metadata in RFC 3339 format.'
	Updated pulumi.StringOutput `pulumi:"updated"`
}

// NewHmacKey registers a new resource with the given unique name, arguments, and options.
func NewHmacKey(ctx *pulumi.Context,
	name string, args *HmacKeyArgs, opts ...pulumi.ResourceOption) (*HmacKey, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ServiceAccountEmail == nil {
		return nil, errors.New("invalid value for required argument 'ServiceAccountEmail'")
	}
	var resource HmacKey
	err := ctx.RegisterResource("gcp:storage/hmacKey:HmacKey", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetHmacKey gets an existing HmacKey resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetHmacKey(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *HmacKeyState, opts ...pulumi.ResourceOption) (*HmacKey, error) {
	var resource HmacKey
	err := ctx.ReadResource("gcp:storage/hmacKey:HmacKey", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering HmacKey resources.
type hmacKeyState struct {
	// The access ID of the HMAC Key.
	AccessId *string `pulumi:"accessId"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// HMAC secret key material.
	Secret *string `pulumi:"secret"`
	// The email address of the key's associated service account.
	ServiceAccountEmail *string `pulumi:"serviceAccountEmail"`
	// The state of the key. Can be set to one of ACTIVE, INACTIVE.
	// Default value is `ACTIVE`.
	// Possible values are `ACTIVE` and `INACTIVE`.
	State *string `pulumi:"state"`
	// 'The creation time of the HMAC key in RFC 3339 format. '
	TimeCreated *string `pulumi:"timeCreated"`
	// 'The last modification time of the HMAC key metadata in RFC 3339 format.'
	Updated *string `pulumi:"updated"`
}

type HmacKeyState struct {
	// The access ID of the HMAC Key.
	AccessId pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// HMAC secret key material.
	Secret pulumi.StringPtrInput
	// The email address of the key's associated service account.
	ServiceAccountEmail pulumi.StringPtrInput
	// The state of the key. Can be set to one of ACTIVE, INACTIVE.
	// Default value is `ACTIVE`.
	// Possible values are `ACTIVE` and `INACTIVE`.
	State pulumi.StringPtrInput
	// 'The creation time of the HMAC key in RFC 3339 format. '
	TimeCreated pulumi.StringPtrInput
	// 'The last modification time of the HMAC key metadata in RFC 3339 format.'
	Updated pulumi.StringPtrInput
}

func (HmacKeyState) ElementType() reflect.Type {
	return reflect.TypeOf((*hmacKeyState)(nil)).Elem()
}

type hmacKeyArgs struct {
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The email address of the key's associated service account.
	ServiceAccountEmail string `pulumi:"serviceAccountEmail"`
	// The state of the key. Can be set to one of ACTIVE, INACTIVE.
	// Default value is `ACTIVE`.
	// Possible values are `ACTIVE` and `INACTIVE`.
	State *string `pulumi:"state"`
}

// The set of arguments for constructing a HmacKey resource.
type HmacKeyArgs struct {
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The email address of the key's associated service account.
	ServiceAccountEmail pulumi.StringInput
	// The state of the key. Can be set to one of ACTIVE, INACTIVE.
	// Default value is `ACTIVE`.
	// Possible values are `ACTIVE` and `INACTIVE`.
	State pulumi.StringPtrInput
}

func (HmacKeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*hmacKeyArgs)(nil)).Elem()
}

type HmacKeyInput interface {
	pulumi.Input

	ToHmacKeyOutput() HmacKeyOutput
	ToHmacKeyOutputWithContext(ctx context.Context) HmacKeyOutput
}

func (*HmacKey) ElementType() reflect.Type {
	return reflect.TypeOf((*HmacKey)(nil))
}

func (i *HmacKey) ToHmacKeyOutput() HmacKeyOutput {
	return i.ToHmacKeyOutputWithContext(context.Background())
}

func (i *HmacKey) ToHmacKeyOutputWithContext(ctx context.Context) HmacKeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HmacKeyOutput)
}

type HmacKeyOutput struct {
	*pulumi.OutputState
}

func (HmacKeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HmacKey)(nil))
}

func (o HmacKeyOutput) ToHmacKeyOutput() HmacKeyOutput {
	return o
}

func (o HmacKeyOutput) ToHmacKeyOutputWithContext(ctx context.Context) HmacKeyOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(HmacKeyOutput{})
}
