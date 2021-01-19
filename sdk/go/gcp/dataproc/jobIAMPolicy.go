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
// ## google\_dataproc\_job\_iam\_policy
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/dataproc"
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/organizations"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		admin, err := organizations.LookupIAMPolicy(ctx, &organizations.LookupIAMPolicyArgs{
// 			Bindings: []organizations.GetIAMPolicyBinding{
// 				organizations.GetIAMPolicyBinding{
// 					Role: "roles/editor",
// 					Members: []string{
// 						"user:jane@example.com",
// 					},
// 				},
// 			},
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = dataproc.NewJobIAMPolicy(ctx, "editor", &dataproc.JobIAMPolicyArgs{
// 			Project:    pulumi.String("your-project"),
// 			Region:     pulumi.String("your-region"),
// 			JobId:      pulumi.String("your-dataproc-job"),
// 			PolicyData: pulumi.String(admin.PolicyData),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## google\_dataproc\_job\_iam\_binding
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/dataproc"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := dataproc.NewJobIAMBinding(ctx, "editor", &dataproc.JobIAMBindingArgs{
// 			JobId: pulumi.String("your-dataproc-job"),
// 			Members: pulumi.StringArray{
// 				pulumi.String("user:jane@example.com"),
// 			},
// 			Role: pulumi.String("roles/editor"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## google\_dataproc\_job\_iam\_member
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/dataproc"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := dataproc.NewJobIAMMember(ctx, "editor", &dataproc.JobIAMMemberArgs{
// 			JobId:  pulumi.String("your-dataproc-job"),
// 			Member: pulumi.String("user:jane@example.com"),
// 			Role:   pulumi.String("roles/editor"),
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
// Job IAM resources can be imported using the project, region, job id, role and/or member.
//
// ```sh
//  $ pulumi import gcp:dataproc/jobIAMPolicy:JobIAMPolicy editor "projects/{project}/regions/{region}/jobs/{job_id}"
// ```
//
// ```sh
//  $ pulumi import gcp:dataproc/jobIAMPolicy:JobIAMPolicy editor "projects/{project}/regions/{region}/jobs/{job_id} roles/editor"
// ```
//
// ```sh
//  $ pulumi import gcp:dataproc/jobIAMPolicy:JobIAMPolicy editor "projects/{project}/regions/{region}/jobs/{job_id} roles/editor user:jane@example.com"
// ```
//
//  -> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
//
// full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type JobIAMPolicy struct {
	pulumi.CustomResourceState

	// (Computed) The etag of the jobs's IAM policy.
	Etag  pulumi.StringOutput `pulumi:"etag"`
	JobId pulumi.StringOutput `pulumi:"jobId"`
	// The policy data generated by a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringOutput `pulumi:"policyData"`
	// The project in which the job belongs. If it
	// is not provided, the provider will use a default.
	Project pulumi.StringOutput `pulumi:"project"`
	// The region in which the job belongs. If it
	// is not provided, the provider will use a default.
	Region pulumi.StringOutput `pulumi:"region"`
}

// NewJobIAMPolicy registers a new resource with the given unique name, arguments, and options.
func NewJobIAMPolicy(ctx *pulumi.Context,
	name string, args *JobIAMPolicyArgs, opts ...pulumi.ResourceOption) (*JobIAMPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.JobId == nil {
		return nil, errors.New("invalid value for required argument 'JobId'")
	}
	if args.PolicyData == nil {
		return nil, errors.New("invalid value for required argument 'PolicyData'")
	}
	var resource JobIAMPolicy
	err := ctx.RegisterResource("gcp:dataproc/jobIAMPolicy:JobIAMPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetJobIAMPolicy gets an existing JobIAMPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetJobIAMPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *JobIAMPolicyState, opts ...pulumi.ResourceOption) (*JobIAMPolicy, error) {
	var resource JobIAMPolicy
	err := ctx.ReadResource("gcp:dataproc/jobIAMPolicy:JobIAMPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering JobIAMPolicy resources.
type jobIAMPolicyState struct {
	// (Computed) The etag of the jobs's IAM policy.
	Etag  *string `pulumi:"etag"`
	JobId *string `pulumi:"jobId"`
	// The policy data generated by a `organizations.getIAMPolicy` data source.
	PolicyData *string `pulumi:"policyData"`
	// The project in which the job belongs. If it
	// is not provided, the provider will use a default.
	Project *string `pulumi:"project"`
	// The region in which the job belongs. If it
	// is not provided, the provider will use a default.
	Region *string `pulumi:"region"`
}

type JobIAMPolicyState struct {
	// (Computed) The etag of the jobs's IAM policy.
	Etag  pulumi.StringPtrInput
	JobId pulumi.StringPtrInput
	// The policy data generated by a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringPtrInput
	// The project in which the job belongs. If it
	// is not provided, the provider will use a default.
	Project pulumi.StringPtrInput
	// The region in which the job belongs. If it
	// is not provided, the provider will use a default.
	Region pulumi.StringPtrInput
}

func (JobIAMPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*jobIAMPolicyState)(nil)).Elem()
}

type jobIAMPolicyArgs struct {
	JobId string `pulumi:"jobId"`
	// The policy data generated by a `organizations.getIAMPolicy` data source.
	PolicyData string `pulumi:"policyData"`
	// The project in which the job belongs. If it
	// is not provided, the provider will use a default.
	Project *string `pulumi:"project"`
	// The region in which the job belongs. If it
	// is not provided, the provider will use a default.
	Region *string `pulumi:"region"`
}

// The set of arguments for constructing a JobIAMPolicy resource.
type JobIAMPolicyArgs struct {
	JobId pulumi.StringInput
	// The policy data generated by a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringInput
	// The project in which the job belongs. If it
	// is not provided, the provider will use a default.
	Project pulumi.StringPtrInput
	// The region in which the job belongs. If it
	// is not provided, the provider will use a default.
	Region pulumi.StringPtrInput
}

func (JobIAMPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*jobIAMPolicyArgs)(nil)).Elem()
}

type JobIAMPolicyInput interface {
	pulumi.Input

	ToJobIAMPolicyOutput() JobIAMPolicyOutput
	ToJobIAMPolicyOutputWithContext(ctx context.Context) JobIAMPolicyOutput
}

func (*JobIAMPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((*JobIAMPolicy)(nil))
}

func (i *JobIAMPolicy) ToJobIAMPolicyOutput() JobIAMPolicyOutput {
	return i.ToJobIAMPolicyOutputWithContext(context.Background())
}

func (i *JobIAMPolicy) ToJobIAMPolicyOutputWithContext(ctx context.Context) JobIAMPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobIAMPolicyOutput)
}

func (i *JobIAMPolicy) ToJobIAMPolicyPtrOutput() JobIAMPolicyPtrOutput {
	return i.ToJobIAMPolicyPtrOutputWithContext(context.Background())
}

func (i *JobIAMPolicy) ToJobIAMPolicyPtrOutputWithContext(ctx context.Context) JobIAMPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobIAMPolicyPtrOutput)
}

type JobIAMPolicyPtrInput interface {
	pulumi.Input

	ToJobIAMPolicyPtrOutput() JobIAMPolicyPtrOutput
	ToJobIAMPolicyPtrOutputWithContext(ctx context.Context) JobIAMPolicyPtrOutput
}

type jobIAMPolicyPtrType JobIAMPolicyArgs

func (*jobIAMPolicyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**JobIAMPolicy)(nil))
}

func (i *jobIAMPolicyPtrType) ToJobIAMPolicyPtrOutput() JobIAMPolicyPtrOutput {
	return i.ToJobIAMPolicyPtrOutputWithContext(context.Background())
}

func (i *jobIAMPolicyPtrType) ToJobIAMPolicyPtrOutputWithContext(ctx context.Context) JobIAMPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobIAMPolicyPtrOutput)
}

// JobIAMPolicyArrayInput is an input type that accepts JobIAMPolicyArray and JobIAMPolicyArrayOutput values.
// You can construct a concrete instance of `JobIAMPolicyArrayInput` via:
//
//          JobIAMPolicyArray{ JobIAMPolicyArgs{...} }
type JobIAMPolicyArrayInput interface {
	pulumi.Input

	ToJobIAMPolicyArrayOutput() JobIAMPolicyArrayOutput
	ToJobIAMPolicyArrayOutputWithContext(context.Context) JobIAMPolicyArrayOutput
}

type JobIAMPolicyArray []JobIAMPolicyInput

func (JobIAMPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*JobIAMPolicy)(nil))
}

func (i JobIAMPolicyArray) ToJobIAMPolicyArrayOutput() JobIAMPolicyArrayOutput {
	return i.ToJobIAMPolicyArrayOutputWithContext(context.Background())
}

func (i JobIAMPolicyArray) ToJobIAMPolicyArrayOutputWithContext(ctx context.Context) JobIAMPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobIAMPolicyArrayOutput)
}

// JobIAMPolicyMapInput is an input type that accepts JobIAMPolicyMap and JobIAMPolicyMapOutput values.
// You can construct a concrete instance of `JobIAMPolicyMapInput` via:
//
//          JobIAMPolicyMap{ "key": JobIAMPolicyArgs{...} }
type JobIAMPolicyMapInput interface {
	pulumi.Input

	ToJobIAMPolicyMapOutput() JobIAMPolicyMapOutput
	ToJobIAMPolicyMapOutputWithContext(context.Context) JobIAMPolicyMapOutput
}

type JobIAMPolicyMap map[string]JobIAMPolicyInput

func (JobIAMPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*JobIAMPolicy)(nil))
}

func (i JobIAMPolicyMap) ToJobIAMPolicyMapOutput() JobIAMPolicyMapOutput {
	return i.ToJobIAMPolicyMapOutputWithContext(context.Background())
}

func (i JobIAMPolicyMap) ToJobIAMPolicyMapOutputWithContext(ctx context.Context) JobIAMPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobIAMPolicyMapOutput)
}

type JobIAMPolicyOutput struct {
	*pulumi.OutputState
}

func (JobIAMPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JobIAMPolicy)(nil))
}

func (o JobIAMPolicyOutput) ToJobIAMPolicyOutput() JobIAMPolicyOutput {
	return o
}

func (o JobIAMPolicyOutput) ToJobIAMPolicyOutputWithContext(ctx context.Context) JobIAMPolicyOutput {
	return o
}

func (o JobIAMPolicyOutput) ToJobIAMPolicyPtrOutput() JobIAMPolicyPtrOutput {
	return o.ToJobIAMPolicyPtrOutputWithContext(context.Background())
}

func (o JobIAMPolicyOutput) ToJobIAMPolicyPtrOutputWithContext(ctx context.Context) JobIAMPolicyPtrOutput {
	return o.ApplyT(func(v JobIAMPolicy) *JobIAMPolicy {
		return &v
	}).(JobIAMPolicyPtrOutput)
}

type JobIAMPolicyPtrOutput struct {
	*pulumi.OutputState
}

func (JobIAMPolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**JobIAMPolicy)(nil))
}

func (o JobIAMPolicyPtrOutput) ToJobIAMPolicyPtrOutput() JobIAMPolicyPtrOutput {
	return o
}

func (o JobIAMPolicyPtrOutput) ToJobIAMPolicyPtrOutputWithContext(ctx context.Context) JobIAMPolicyPtrOutput {
	return o
}

type JobIAMPolicyArrayOutput struct{ *pulumi.OutputState }

func (JobIAMPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]JobIAMPolicy)(nil))
}

func (o JobIAMPolicyArrayOutput) ToJobIAMPolicyArrayOutput() JobIAMPolicyArrayOutput {
	return o
}

func (o JobIAMPolicyArrayOutput) ToJobIAMPolicyArrayOutputWithContext(ctx context.Context) JobIAMPolicyArrayOutput {
	return o
}

func (o JobIAMPolicyArrayOutput) Index(i pulumi.IntInput) JobIAMPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) JobIAMPolicy {
		return vs[0].([]JobIAMPolicy)[vs[1].(int)]
	}).(JobIAMPolicyOutput)
}

type JobIAMPolicyMapOutput struct{ *pulumi.OutputState }

func (JobIAMPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]JobIAMPolicy)(nil))
}

func (o JobIAMPolicyMapOutput) ToJobIAMPolicyMapOutput() JobIAMPolicyMapOutput {
	return o
}

func (o JobIAMPolicyMapOutput) ToJobIAMPolicyMapOutputWithContext(ctx context.Context) JobIAMPolicyMapOutput {
	return o
}

func (o JobIAMPolicyMapOutput) MapIndex(k pulumi.StringInput) JobIAMPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) JobIAMPolicy {
		return vs[0].(map[string]JobIAMPolicy)[vs[1].(string)]
	}).(JobIAMPolicyOutput)
}

func init() {
	pulumi.RegisterOutputType(JobIAMPolicyOutput{})
	pulumi.RegisterOutputType(JobIAMPolicyPtrOutput{})
	pulumi.RegisterOutputType(JobIAMPolicyArrayOutput{})
	pulumi.RegisterOutputType(JobIAMPolicyMapOutput{})
}
