// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package dataflow

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Creates a [Flex Template](https://cloud.google.com/dataflow/docs/guides/templates/using-flex-templates)
// job on Dataflow, which is an implementation of Apache Beam running on Google
// Compute Engine. For more information see the official documentation for [Beam](https://beam.apache.org)
// and [Dataflow](https://cloud.google.com/dataflow/).
//
// ## Note on "destroy" / "apply"
//
// There are many types of Dataflow jobs.  Some Dataflow jobs run constantly,
// getting new data from (e.g.) a GCS bucket, and outputting data continuously.
// Some jobs process a set amount of data then terminate. All jobs can fail while
// running due to programming errors or other issues. In this way, Dataflow jobs
// are different from most other provider / Google resources.
//
// The Dataflow resource is considered 'existing' while it is in a nonterminal
// state.  If it reaches a terminal state (e.g. 'FAILED', 'COMPLETE',
// 'CANCELLED'), it will be recreated on the next 'apply'.  This is as expected for
// jobs which run continuously, but may surprise users who use this resource for
// other kinds of Dataflow jobs.
//
// A Dataflow job which is 'destroyed' may be "cancelled" or "drained".  If
// "cancelled", the job terminates - any data written remains where it is, but no
// new data will be processed.  If "drained", no new data will enter the pipeline,
// but any data currently in the pipeline will finish being processed.  The default
// is "cancelled", but if a user sets `onDelete` to `"drain"` in the
// configuration, you may experience a long wait for your `pulumi destroy` to
// complete.
//
// ## Import
//
// This resource does not support import.
type FlexTemplateJob struct {
	pulumi.CustomResourceState

	// The GCS path to the Dataflow job Flex
	// Template.
	ContainerSpecGcsPath pulumi.StringOutput `pulumi:"containerSpecGcsPath"`
	// The unique ID of this job.
	JobId pulumi.StringOutput `pulumi:"jobId"`
	// Deprecated: Deprecated until the API supports this field
	Labels pulumi.MapOutput `pulumi:"labels"`
	// A unique name for the resource, required by Dataflow.
	Name pulumi.StringOutput `pulumi:"name"`
	// One of "drain" or "cancel". Specifies behavior of
	// deletion during `pulumi destroy`.  See above note.
	OnDelete pulumi.StringPtrOutput `pulumi:"onDelete"`
	// Key/Value pairs to be passed to the Dataflow job (as
	// used in the template).
	Parameters pulumi.MapOutput `pulumi:"parameters"`
	// The project in which the resource belongs. If it is not
	// provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The region in which the created job should run.
	Region pulumi.StringOutput `pulumi:"region"`
	// The current state of the resource, selected from the [JobState enum](https://cloud.google.com/dataflow/docs/reference/rest/v1b3/projects.jobs#Job.JobState)
	State pulumi.StringOutput `pulumi:"state"`
}

// NewFlexTemplateJob registers a new resource with the given unique name, arguments, and options.
func NewFlexTemplateJob(ctx *pulumi.Context,
	name string, args *FlexTemplateJobArgs, opts ...pulumi.ResourceOption) (*FlexTemplateJob, error) {
	if args == nil || args.ContainerSpecGcsPath == nil {
		return nil, errors.New("missing required argument 'ContainerSpecGcsPath'")
	}
	if args == nil {
		args = &FlexTemplateJobArgs{}
	}
	var resource FlexTemplateJob
	err := ctx.RegisterResource("gcp:dataflow/flexTemplateJob:FlexTemplateJob", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFlexTemplateJob gets an existing FlexTemplateJob resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFlexTemplateJob(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FlexTemplateJobState, opts ...pulumi.ResourceOption) (*FlexTemplateJob, error) {
	var resource FlexTemplateJob
	err := ctx.ReadResource("gcp:dataflow/flexTemplateJob:FlexTemplateJob", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FlexTemplateJob resources.
type flexTemplateJobState struct {
	// The GCS path to the Dataflow job Flex
	// Template.
	ContainerSpecGcsPath *string `pulumi:"containerSpecGcsPath"`
	// The unique ID of this job.
	JobId *string `pulumi:"jobId"`
	// Deprecated: Deprecated until the API supports this field
	Labels map[string]interface{} `pulumi:"labels"`
	// A unique name for the resource, required by Dataflow.
	Name *string `pulumi:"name"`
	// One of "drain" or "cancel". Specifies behavior of
	// deletion during `pulumi destroy`.  See above note.
	OnDelete *string `pulumi:"onDelete"`
	// Key/Value pairs to be passed to the Dataflow job (as
	// used in the template).
	Parameters map[string]interface{} `pulumi:"parameters"`
	// The project in which the resource belongs. If it is not
	// provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The region in which the created job should run.
	Region *string `pulumi:"region"`
	// The current state of the resource, selected from the [JobState enum](https://cloud.google.com/dataflow/docs/reference/rest/v1b3/projects.jobs#Job.JobState)
	State *string `pulumi:"state"`
}

type FlexTemplateJobState struct {
	// The GCS path to the Dataflow job Flex
	// Template.
	ContainerSpecGcsPath pulumi.StringPtrInput
	// The unique ID of this job.
	JobId pulumi.StringPtrInput
	// Deprecated: Deprecated until the API supports this field
	Labels pulumi.MapInput
	// A unique name for the resource, required by Dataflow.
	Name pulumi.StringPtrInput
	// One of "drain" or "cancel". Specifies behavior of
	// deletion during `pulumi destroy`.  See above note.
	OnDelete pulumi.StringPtrInput
	// Key/Value pairs to be passed to the Dataflow job (as
	// used in the template).
	Parameters pulumi.MapInput
	// The project in which the resource belongs. If it is not
	// provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The region in which the created job should run.
	Region pulumi.StringPtrInput
	// The current state of the resource, selected from the [JobState enum](https://cloud.google.com/dataflow/docs/reference/rest/v1b3/projects.jobs#Job.JobState)
	State pulumi.StringPtrInput
}

func (FlexTemplateJobState) ElementType() reflect.Type {
	return reflect.TypeOf((*flexTemplateJobState)(nil)).Elem()
}

type flexTemplateJobArgs struct {
	// The GCS path to the Dataflow job Flex
	// Template.
	ContainerSpecGcsPath string `pulumi:"containerSpecGcsPath"`
	// Deprecated: Deprecated until the API supports this field
	Labels map[string]interface{} `pulumi:"labels"`
	// A unique name for the resource, required by Dataflow.
	Name *string `pulumi:"name"`
	// One of "drain" or "cancel". Specifies behavior of
	// deletion during `pulumi destroy`.  See above note.
	OnDelete *string `pulumi:"onDelete"`
	// Key/Value pairs to be passed to the Dataflow job (as
	// used in the template).
	Parameters map[string]interface{} `pulumi:"parameters"`
	// The project in which the resource belongs. If it is not
	// provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The region in which the created job should run.
	Region *string `pulumi:"region"`
}

// The set of arguments for constructing a FlexTemplateJob resource.
type FlexTemplateJobArgs struct {
	// The GCS path to the Dataflow job Flex
	// Template.
	ContainerSpecGcsPath pulumi.StringInput
	// Deprecated: Deprecated until the API supports this field
	Labels pulumi.MapInput
	// A unique name for the resource, required by Dataflow.
	Name pulumi.StringPtrInput
	// One of "drain" or "cancel". Specifies behavior of
	// deletion during `pulumi destroy`.  See above note.
	OnDelete pulumi.StringPtrInput
	// Key/Value pairs to be passed to the Dataflow job (as
	// used in the template).
	Parameters pulumi.MapInput
	// The project in which the resource belongs. If it is not
	// provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The region in which the created job should run.
	Region pulumi.StringPtrInput
}

func (FlexTemplateJobArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*flexTemplateJobArgs)(nil)).Elem()
}
