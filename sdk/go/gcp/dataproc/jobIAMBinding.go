// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package dataproc

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Three different resources help you manage IAM policies on dataproc jobs. Each of these resources serves a different use case:
//
// * `dataproc.JobIAMPolicy`: Authoritative. Sets the IAM policy for the job and replaces any existing policy already attached.
// * `dataproc.JobIAMBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the job are preserved.
// * `dataproc.JobIAMMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the job are preserved.
//
// > **Note:** `dataproc.JobIAMPolicy` **cannot** be used in conjunction with `dataproc.JobIAMBinding` and `dataproc.JobIAMMember` or they will fight over what your policy should be. In addition, be careful not to accidentally unset ownership of the job as `dataproc.JobIAMPolicy` replaces the entire policy.
//
// > **Note:** `dataproc.JobIAMBinding` resources **can be** used in conjunction with `dataproc.JobIAMMember` resources **only if** they do not grant privilege to the same role.
//
// ## Import
//
// Job IAM resources can be imported using the project, region, job id, role and/or member.
//
// ```sh
//  $ pulumi import gcp:dataproc/jobIAMBinding:JobIAMBinding editor "projects/{project}/regions/{region}/jobs/{job_id}"
// ```
//
// ```sh
//  $ pulumi import gcp:dataproc/jobIAMBinding:JobIAMBinding editor "projects/{project}/regions/{region}/jobs/{job_id} roles/editor"
// ```
//
// ```sh
//  $ pulumi import gcp:dataproc/jobIAMBinding:JobIAMBinding editor "projects/{project}/regions/{region}/jobs/{job_id} roles/editor user:jane@example.com"
// ```
//
//  -> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
//
// full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type JobIAMBinding struct {
	pulumi.CustomResourceState

	Condition JobIAMBindingConditionPtrOutput `pulumi:"condition"`
	// (Computed) The etag of the jobs's IAM policy.
	Etag    pulumi.StringOutput      `pulumi:"etag"`
	JobId   pulumi.StringOutput      `pulumi:"jobId"`
	Members pulumi.StringArrayOutput `pulumi:"members"`
	// The project in which the job belongs. If it
	// is not provided, the provider will use a default.
	Project pulumi.StringOutput `pulumi:"project"`
	// The region in which the job belongs. If it
	// is not provided, the provider will use a default.
	Region pulumi.StringOutput `pulumi:"region"`
	// The role that should be applied. Only one
	// `dataproc.JobIAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringOutput `pulumi:"role"`
}

// NewJobIAMBinding registers a new resource with the given unique name, arguments, and options.
func NewJobIAMBinding(ctx *pulumi.Context,
	name string, args *JobIAMBindingArgs, opts ...pulumi.ResourceOption) (*JobIAMBinding, error) {
	if args == nil || args.JobId == nil {
		return nil, errors.New("missing required argument 'JobId'")
	}
	if args == nil || args.Members == nil {
		return nil, errors.New("missing required argument 'Members'")
	}
	if args == nil || args.Role == nil {
		return nil, errors.New("missing required argument 'Role'")
	}
	if args == nil {
		args = &JobIAMBindingArgs{}
	}
	var resource JobIAMBinding
	err := ctx.RegisterResource("gcp:dataproc/jobIAMBinding:JobIAMBinding", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetJobIAMBinding gets an existing JobIAMBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetJobIAMBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *JobIAMBindingState, opts ...pulumi.ResourceOption) (*JobIAMBinding, error) {
	var resource JobIAMBinding
	err := ctx.ReadResource("gcp:dataproc/jobIAMBinding:JobIAMBinding", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering JobIAMBinding resources.
type jobIAMBindingState struct {
	Condition *JobIAMBindingCondition `pulumi:"condition"`
	// (Computed) The etag of the jobs's IAM policy.
	Etag    *string  `pulumi:"etag"`
	JobId   *string  `pulumi:"jobId"`
	Members []string `pulumi:"members"`
	// The project in which the job belongs. If it
	// is not provided, the provider will use a default.
	Project *string `pulumi:"project"`
	// The region in which the job belongs. If it
	// is not provided, the provider will use a default.
	Region *string `pulumi:"region"`
	// The role that should be applied. Only one
	// `dataproc.JobIAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role *string `pulumi:"role"`
}

type JobIAMBindingState struct {
	Condition JobIAMBindingConditionPtrInput
	// (Computed) The etag of the jobs's IAM policy.
	Etag    pulumi.StringPtrInput
	JobId   pulumi.StringPtrInput
	Members pulumi.StringArrayInput
	// The project in which the job belongs. If it
	// is not provided, the provider will use a default.
	Project pulumi.StringPtrInput
	// The region in which the job belongs. If it
	// is not provided, the provider will use a default.
	Region pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `dataproc.JobIAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringPtrInput
}

func (JobIAMBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*jobIAMBindingState)(nil)).Elem()
}

type jobIAMBindingArgs struct {
	Condition *JobIAMBindingCondition `pulumi:"condition"`
	JobId     string                  `pulumi:"jobId"`
	Members   []string                `pulumi:"members"`
	// The project in which the job belongs. If it
	// is not provided, the provider will use a default.
	Project *string `pulumi:"project"`
	// The region in which the job belongs. If it
	// is not provided, the provider will use a default.
	Region *string `pulumi:"region"`
	// The role that should be applied. Only one
	// `dataproc.JobIAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role string `pulumi:"role"`
}

// The set of arguments for constructing a JobIAMBinding resource.
type JobIAMBindingArgs struct {
	Condition JobIAMBindingConditionPtrInput
	JobId     pulumi.StringInput
	Members   pulumi.StringArrayInput
	// The project in which the job belongs. If it
	// is not provided, the provider will use a default.
	Project pulumi.StringPtrInput
	// The region in which the job belongs. If it
	// is not provided, the provider will use a default.
	Region pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `dataproc.JobIAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringInput
}

func (JobIAMBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*jobIAMBindingArgs)(nil)).Elem()
}

type JobIAMBindingInput interface {
	pulumi.Input

	ToJobIAMBindingOutput() JobIAMBindingOutput
	ToJobIAMBindingOutputWithContext(ctx context.Context) JobIAMBindingOutput
}

func (JobIAMBinding) ElementType() reflect.Type {
	return reflect.TypeOf((*JobIAMBinding)(nil)).Elem()
}

func (i JobIAMBinding) ToJobIAMBindingOutput() JobIAMBindingOutput {
	return i.ToJobIAMBindingOutputWithContext(context.Background())
}

func (i JobIAMBinding) ToJobIAMBindingOutputWithContext(ctx context.Context) JobIAMBindingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobIAMBindingOutput)
}

type JobIAMBindingOutput struct {
	*pulumi.OutputState
}

func (JobIAMBindingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JobIAMBindingOutput)(nil)).Elem()
}

func (o JobIAMBindingOutput) ToJobIAMBindingOutput() JobIAMBindingOutput {
	return o
}

func (o JobIAMBindingOutput) ToJobIAMBindingOutputWithContext(ctx context.Context) JobIAMBindingOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(JobIAMBindingOutput{})
}
