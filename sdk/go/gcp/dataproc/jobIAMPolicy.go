// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package dataproc

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Three different resources help you manage IAM policies on dataproc jobs. Each of these resources serves a different use case:
// 
// * `dataproc.JobIAMPolicy`: Authoritative. Sets the IAM policy for the job and replaces any existing policy already attached.
// * `dataproc.JobIAMBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the job are preserved.
// * `dataproc.JobIAMMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the job are preserved.
// 
// > **Note:** `dataproc.JobIAMPolicy` **cannot** be used in conjunction with `dataproc.JobIAMBinding` and `dataproc.JobIAMMember` or they will fight over what your policy should be. In addition, be careful not to accidentaly unset ownership of the job as `dataproc.JobIAMPolicy` replaces the entire policy.
// 
// > **Note:** `dataproc.JobIAMBinding` resources **can be** used in conjunction with `dataproc.JobIAMMember` resources **only if** they do not grant privilege to the same role.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/dataproc_job_iam_policy.html.markdown.
type JobIAMPolicy struct {
	s *pulumi.ResourceState
}

// NewJobIAMPolicy registers a new resource with the given unique name, arguments, and options.
func NewJobIAMPolicy(ctx *pulumi.Context,
	name string, args *JobIAMPolicyArgs, opts ...pulumi.ResourceOpt) (*JobIAMPolicy, error) {
	if args == nil || args.JobId == nil {
		return nil, errors.New("missing required argument 'JobId'")
	}
	if args == nil || args.PolicyData == nil {
		return nil, errors.New("missing required argument 'PolicyData'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["jobId"] = nil
		inputs["policyData"] = nil
		inputs["project"] = nil
		inputs["region"] = nil
	} else {
		inputs["jobId"] = args.JobId
		inputs["policyData"] = args.PolicyData
		inputs["project"] = args.Project
		inputs["region"] = args.Region
	}
	inputs["etag"] = nil
	s, err := ctx.RegisterResource("gcp:dataproc/jobIAMPolicy:JobIAMPolicy", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &JobIAMPolicy{s: s}, nil
}

// GetJobIAMPolicy gets an existing JobIAMPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetJobIAMPolicy(ctx *pulumi.Context,
	name string, id pulumi.ID, state *JobIAMPolicyState, opts ...pulumi.ResourceOpt) (*JobIAMPolicy, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["etag"] = state.Etag
		inputs["jobId"] = state.JobId
		inputs["policyData"] = state.PolicyData
		inputs["project"] = state.Project
		inputs["region"] = state.Region
	}
	s, err := ctx.ReadResource("gcp:dataproc/jobIAMPolicy:JobIAMPolicy", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &JobIAMPolicy{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *JobIAMPolicy) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *JobIAMPolicy) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// (Computed) The etag of the jobs's IAM policy.
func (r *JobIAMPolicy) Etag() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["etag"])
}

func (r *JobIAMPolicy) JobId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["jobId"])
}

// The policy data generated by a `organizations.getIAMPolicy` data source.
func (r *JobIAMPolicy) PolicyData() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["policyData"])
}

// The project in which the job belongs. If it
// is not provided, this provider will use the provider default.
func (r *JobIAMPolicy) Project() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["project"])
}

// The region in which the job belongs. If it
// is not provided, this provider will use the provider default.
func (r *JobIAMPolicy) Region() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["region"])
}

// Input properties used for looking up and filtering JobIAMPolicy resources.
type JobIAMPolicyState struct {
	// (Computed) The etag of the jobs's IAM policy.
	Etag interface{}
	JobId interface{}
	// The policy data generated by a `organizations.getIAMPolicy` data source.
	PolicyData interface{}
	// The project in which the job belongs. If it
	// is not provided, this provider will use the provider default.
	Project interface{}
	// The region in which the job belongs. If it
	// is not provided, this provider will use the provider default.
	Region interface{}
}

// The set of arguments for constructing a JobIAMPolicy resource.
type JobIAMPolicyArgs struct {
	JobId interface{}
	// The policy data generated by a `organizations.getIAMPolicy` data source.
	PolicyData interface{}
	// The project in which the job belongs. If it
	// is not provided, this provider will use the provider default.
	Project interface{}
	// The region in which the job belongs. If it
	// is not provided, this provider will use the provider default.
	Region interface{}
}
