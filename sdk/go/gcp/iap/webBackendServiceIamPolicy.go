// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package iap

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// ## Import
//
// For all import syntaxes, the "resource in question" can take any of the following forms* projects/{{project}}/iap_web/compute/services/{{name}} * {{project}}/{{name}} * {{name}} Any variables not passed in the import command will be taken from the provider configuration. Identity-Aware Proxy webbackendservice IAM resources can be imported using the resource identifiers, role, and member. IAM member imports use space-delimited identifiersthe resource in question, the role, and the member identity, e.g.
//
// ```sh
//  $ pulumi import gcp:iap/webBackendServiceIamPolicy:WebBackendServiceIamPolicy editor "projects/{{project}}/iap_web/compute/services/{{web_backend_service}} roles/iap.httpsResourceAccessor user:jane@example.com"
// ```
//
//  IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.
//
// ```sh
//  $ pulumi import gcp:iap/webBackendServiceIamPolicy:WebBackendServiceIamPolicy editor "projects/{{project}}/iap_web/compute/services/{{web_backend_service}} roles/iap.httpsResourceAccessor"
// ```
//
//  IAM policy imports use the identifier of the resource in question, e.g.
//
// ```sh
//  $ pulumi import gcp:iap/webBackendServiceIamPolicy:WebBackendServiceIamPolicy editor projects/{{project}}/iap_web/compute/services/{{web_backend_service}}
// ```
//
//  -> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
//
// full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type WebBackendServiceIamPolicy struct {
	pulumi.CustomResourceState

	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringOutput `pulumi:"policyData"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// Used to find the parent resource to bind the IAM policy to
	WebBackendService pulumi.StringOutput `pulumi:"webBackendService"`
}

// NewWebBackendServiceIamPolicy registers a new resource with the given unique name, arguments, and options.
func NewWebBackendServiceIamPolicy(ctx *pulumi.Context,
	name string, args *WebBackendServiceIamPolicyArgs, opts ...pulumi.ResourceOption) (*WebBackendServiceIamPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PolicyData == nil {
		return nil, errors.New("invalid value for required argument 'PolicyData'")
	}
	if args.WebBackendService == nil {
		return nil, errors.New("invalid value for required argument 'WebBackendService'")
	}
	var resource WebBackendServiceIamPolicy
	err := ctx.RegisterResource("gcp:iap/webBackendServiceIamPolicy:WebBackendServiceIamPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWebBackendServiceIamPolicy gets an existing WebBackendServiceIamPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWebBackendServiceIamPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebBackendServiceIamPolicyState, opts ...pulumi.ResourceOption) (*WebBackendServiceIamPolicy, error) {
	var resource WebBackendServiceIamPolicy
	err := ctx.ReadResource("gcp:iap/webBackendServiceIamPolicy:WebBackendServiceIamPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WebBackendServiceIamPolicy resources.
type webBackendServiceIamPolicyState struct {
	// (Computed) The etag of the IAM policy.
	Etag *string `pulumi:"etag"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData *string `pulumi:"policyData"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	// Used to find the parent resource to bind the IAM policy to
	WebBackendService *string `pulumi:"webBackendService"`
}

type WebBackendServiceIamPolicyState struct {
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringPtrInput
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	// Used to find the parent resource to bind the IAM policy to
	WebBackendService pulumi.StringPtrInput
}

func (WebBackendServiceIamPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*webBackendServiceIamPolicyState)(nil)).Elem()
}

type webBackendServiceIamPolicyArgs struct {
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData string `pulumi:"policyData"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	// Used to find the parent resource to bind the IAM policy to
	WebBackendService string `pulumi:"webBackendService"`
}

// The set of arguments for constructing a WebBackendServiceIamPolicy resource.
type WebBackendServiceIamPolicyArgs struct {
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	// Used to find the parent resource to bind the IAM policy to
	WebBackendService pulumi.StringInput
}

func (WebBackendServiceIamPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webBackendServiceIamPolicyArgs)(nil)).Elem()
}

type WebBackendServiceIamPolicyInput interface {
	pulumi.Input

	ToWebBackendServiceIamPolicyOutput() WebBackendServiceIamPolicyOutput
	ToWebBackendServiceIamPolicyOutputWithContext(ctx context.Context) WebBackendServiceIamPolicyOutput
}

func (*WebBackendServiceIamPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((*WebBackendServiceIamPolicy)(nil))
}

func (i *WebBackendServiceIamPolicy) ToWebBackendServiceIamPolicyOutput() WebBackendServiceIamPolicyOutput {
	return i.ToWebBackendServiceIamPolicyOutputWithContext(context.Background())
}

func (i *WebBackendServiceIamPolicy) ToWebBackendServiceIamPolicyOutputWithContext(ctx context.Context) WebBackendServiceIamPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebBackendServiceIamPolicyOutput)
}

type WebBackendServiceIamPolicyOutput struct {
	*pulumi.OutputState
}

func (WebBackendServiceIamPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WebBackendServiceIamPolicy)(nil))
}

func (o WebBackendServiceIamPolicyOutput) ToWebBackendServiceIamPolicyOutput() WebBackendServiceIamPolicyOutput {
	return o
}

func (o WebBackendServiceIamPolicyOutput) ToWebBackendServiceIamPolicyOutputWithContext(ctx context.Context) WebBackendServiceIamPolicyOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(WebBackendServiceIamPolicyOutput{})
}
