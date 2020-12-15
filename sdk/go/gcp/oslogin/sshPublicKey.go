// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package oslogin

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The SSH public key information associated with a Google account.
//
// To get more information about SSHPublicKey, see:
//
// * [API documentation](https://cloud.google.com/compute/docs/oslogin/rest/v1/users.sshPublicKeys)
// * How-to Guides
//     * [Official Documentation](https://cloud.google.com/compute/docs/oslogin)
//
// ## Example Usage
//
// ## Import
//
// SSHPublicKey can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import gcp:oslogin/sshPublicKey:SshPublicKey default users/{{user}}/sshPublicKeys/{{fingerprint}}
// ```
//
// ```sh
//  $ pulumi import gcp:oslogin/sshPublicKey:SshPublicKey default {{user}}/{{fingerprint}}
// ```
type SshPublicKey struct {
	pulumi.CustomResourceState

	// An expiration time in microseconds since epoch.
	ExpirationTimeUsec pulumi.StringPtrOutput `pulumi:"expirationTimeUsec"`
	// The SHA-256 fingerprint of the SSH public key.
	Fingerprint pulumi.StringOutput `pulumi:"fingerprint"`
	// Public key text in SSH format, defined by RFC4253 section 6.6.
	Key pulumi.StringOutput `pulumi:"key"`
	// The project ID of the Google Cloud Platform project.
	Project pulumi.StringPtrOutput `pulumi:"project"`
	// The user email.
	User pulumi.StringOutput `pulumi:"user"`
}

// NewSshPublicKey registers a new resource with the given unique name, arguments, and options.
func NewSshPublicKey(ctx *pulumi.Context,
	name string, args *SshPublicKeyArgs, opts ...pulumi.ResourceOption) (*SshPublicKey, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Key == nil {
		return nil, errors.New("invalid value for required argument 'Key'")
	}
	if args.User == nil {
		return nil, errors.New("invalid value for required argument 'User'")
	}
	var resource SshPublicKey
	err := ctx.RegisterResource("gcp:oslogin/sshPublicKey:SshPublicKey", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSshPublicKey gets an existing SshPublicKey resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSshPublicKey(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SshPublicKeyState, opts ...pulumi.ResourceOption) (*SshPublicKey, error) {
	var resource SshPublicKey
	err := ctx.ReadResource("gcp:oslogin/sshPublicKey:SshPublicKey", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SshPublicKey resources.
type sshPublicKeyState struct {
	// An expiration time in microseconds since epoch.
	ExpirationTimeUsec *string `pulumi:"expirationTimeUsec"`
	// The SHA-256 fingerprint of the SSH public key.
	Fingerprint *string `pulumi:"fingerprint"`
	// Public key text in SSH format, defined by RFC4253 section 6.6.
	Key *string `pulumi:"key"`
	// The project ID of the Google Cloud Platform project.
	Project *string `pulumi:"project"`
	// The user email.
	User *string `pulumi:"user"`
}

type SshPublicKeyState struct {
	// An expiration time in microseconds since epoch.
	ExpirationTimeUsec pulumi.StringPtrInput
	// The SHA-256 fingerprint of the SSH public key.
	Fingerprint pulumi.StringPtrInput
	// Public key text in SSH format, defined by RFC4253 section 6.6.
	Key pulumi.StringPtrInput
	// The project ID of the Google Cloud Platform project.
	Project pulumi.StringPtrInput
	// The user email.
	User pulumi.StringPtrInput
}

func (SshPublicKeyState) ElementType() reflect.Type {
	return reflect.TypeOf((*sshPublicKeyState)(nil)).Elem()
}

type sshPublicKeyArgs struct {
	// An expiration time in microseconds since epoch.
	ExpirationTimeUsec *string `pulumi:"expirationTimeUsec"`
	// Public key text in SSH format, defined by RFC4253 section 6.6.
	Key string `pulumi:"key"`
	// The project ID of the Google Cloud Platform project.
	Project *string `pulumi:"project"`
	// The user email.
	User string `pulumi:"user"`
}

// The set of arguments for constructing a SshPublicKey resource.
type SshPublicKeyArgs struct {
	// An expiration time in microseconds since epoch.
	ExpirationTimeUsec pulumi.StringPtrInput
	// Public key text in SSH format, defined by RFC4253 section 6.6.
	Key pulumi.StringInput
	// The project ID of the Google Cloud Platform project.
	Project pulumi.StringPtrInput
	// The user email.
	User pulumi.StringInput
}

func (SshPublicKeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sshPublicKeyArgs)(nil)).Elem()
}

type SshPublicKeyInput interface {
	pulumi.Input

	ToSshPublicKeyOutput() SshPublicKeyOutput
	ToSshPublicKeyOutputWithContext(ctx context.Context) SshPublicKeyOutput
}

func (*SshPublicKey) ElementType() reflect.Type {
	return reflect.TypeOf((*SshPublicKey)(nil))
}

func (i *SshPublicKey) ToSshPublicKeyOutput() SshPublicKeyOutput {
	return i.ToSshPublicKeyOutputWithContext(context.Background())
}

func (i *SshPublicKey) ToSshPublicKeyOutputWithContext(ctx context.Context) SshPublicKeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SshPublicKeyOutput)
}

type SshPublicKeyOutput struct {
	*pulumi.OutputState
}

func (SshPublicKeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SshPublicKey)(nil))
}

func (o SshPublicKeyOutput) ToSshPublicKeyOutput() SshPublicKeyOutput {
	return o
}

func (o SshPublicKeyOutput) ToSshPublicKeyOutputWithContext(ctx context.Context) SshPublicKeyOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(SshPublicKeyOutput{})
}
