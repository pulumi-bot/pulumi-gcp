// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package iap

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Contains the data that describes an Identity Aware Proxy owned client.
//
// > **Note:** Only internal org clients can be created via declarative tools. Other types of clients must be
// manually created via the GCP console. This restriction is due to the existing APIs and not lack of support
// in this tool.
//
// To get more information about Client, see:
//
// * [API documentation](https://cloud.google.com/iap/docs/reference/rest/v1/projects.brands.identityAwareProxyClients)
// * How-to Guides
//     * [Setting up IAP Client](https://cloud.google.com/iap/docs/authentication-howto)
//
// > **Warning:** All arguments including `secret` will be stored in the raw
// state as plain-text. [Read more about secrets in state](https://www.pulumi.com/docs/intro/concepts/programming-model/#secrets).
//
// ## Example Usage
type Client struct {
	pulumi.CustomResourceState

	// Identifier of the brand to which this client
	// is attached to. The format is
	// `projects/{project_number}/brands/{brand_id}/identityAwareProxyClients/{client_id}`.
	Brand pulumi.StringOutput `pulumi:"brand"`
	// Output only. Unique identifier of the OAuth client.
	ClientId pulumi.StringOutput `pulumi:"clientId"`
	// Human-friendly name given to the OAuth client.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// Output only. Client secret of the OAuth client.
	Secret pulumi.StringOutput `pulumi:"secret"`
}

// NewClient registers a new resource with the given unique name, arguments, and options.
func NewClient(ctx *pulumi.Context,
	name string, args *ClientArgs, opts ...pulumi.ResourceOption) (*Client, error) {
	if args == nil || args.Brand == nil {
		return nil, errors.New("missing required argument 'Brand'")
	}
	if args == nil || args.DisplayName == nil {
		return nil, errors.New("missing required argument 'DisplayName'")
	}
	if args == nil {
		args = &ClientArgs{}
	}
	var resource Client
	err := ctx.RegisterResource("gcp:iap/client:Client", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetClient gets an existing Client resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetClient(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClientState, opts ...pulumi.ResourceOption) (*Client, error) {
	var resource Client
	err := ctx.ReadResource("gcp:iap/client:Client", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Client resources.
type clientState struct {
	// Identifier of the brand to which this client
	// is attached to. The format is
	// `projects/{project_number}/brands/{brand_id}/identityAwareProxyClients/{client_id}`.
	Brand *string `pulumi:"brand"`
	// Output only. Unique identifier of the OAuth client.
	ClientId *string `pulumi:"clientId"`
	// Human-friendly name given to the OAuth client.
	DisplayName *string `pulumi:"displayName"`
	// Output only. Client secret of the OAuth client.
	Secret *string `pulumi:"secret"`
}

type ClientState struct {
	// Identifier of the brand to which this client
	// is attached to. The format is
	// `projects/{project_number}/brands/{brand_id}/identityAwareProxyClients/{client_id}`.
	Brand pulumi.StringPtrInput
	// Output only. Unique identifier of the OAuth client.
	ClientId pulumi.StringPtrInput
	// Human-friendly name given to the OAuth client.
	DisplayName pulumi.StringPtrInput
	// Output only. Client secret of the OAuth client.
	Secret pulumi.StringPtrInput
}

func (ClientState) ElementType() reflect.Type {
	return reflect.TypeOf((*clientState)(nil)).Elem()
}

type clientArgs struct {
	// Identifier of the brand to which this client
	// is attached to. The format is
	// `projects/{project_number}/brands/{brand_id}/identityAwareProxyClients/{client_id}`.
	Brand string `pulumi:"brand"`
	// Human-friendly name given to the OAuth client.
	DisplayName string `pulumi:"displayName"`
}

// The set of arguments for constructing a Client resource.
type ClientArgs struct {
	// Identifier of the brand to which this client
	// is attached to. The format is
	// `projects/{project_number}/brands/{brand_id}/identityAwareProxyClients/{client_id}`.
	Brand pulumi.StringInput
	// Human-friendly name given to the OAuth client.
	DisplayName pulumi.StringInput
}

func (ClientArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clientArgs)(nil)).Elem()
}

type ClientInput interface {
	pulumi.Input

	ToClientOutput() ClientOutput
	ToClientOutputWithContext(ctx context.Context) ClientOutput
}

func (Client) ElementType() reflect.Type {
	return reflect.TypeOf((*Client)(nil)).Elem()
}

func (i Client) ToClientOutput() ClientOutput {
	return i.ToClientOutputWithContext(context.Background())
}

func (i Client) ToClientOutputWithContext(ctx context.Context) ClientOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClientOutput)
}

type ClientOutput struct {
	*pulumi.OutputState
}

func (ClientOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClientOutput)(nil)).Elem()
}

func (o ClientOutput) ToClientOutput() ClientOutput {
	return o
}

func (o ClientOutput) ToClientOutputWithContext(ctx context.Context) ClientOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ClientOutput{})
}
