// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package dataflow

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Creates a job on Dataflow, which is an implementation of Apache Beam running on Google Compute Engine. For more information see
// the official documentation for
// [Beam](https://beam.apache.org) and [Dataflow](https://cloud.google.com/dataflow/).
//
// ## Note on "destroy" / "apply"
//
// {{% examples %}}
// There are many types of Dataflow jobs.  Some Dataflow jobs run constantly, getting new data from (e.g.) a GCS bucket, and outputting data continuously.  Some jobs process a set amount of data then terminate.  All jobs can fail while running due to programming errors or other issues.  In this way, Dataflow jobs are different from most other resources.
//
// The Dataflow resource is considered 'existing' while it is in a nonterminal state.  If it reaches a terminal state (e.g. 'FAILED', 'COMPLETE', 'CANCELLED'), it will be recreated on the next 'apply'.  This is as expected for jobs which run continuously, but may surprise users who use this resource for other kinds of Dataflow jobs.
//
// A Dataflow job which is 'destroyed' may be "cancelled" or "drained".  If "cancelled", the job terminates - any data written remains where it is, but no new data will be processed.  If "drained", no new data will enter the pipeline, but any data currently in the pipeline will finish being processed.  The default is "cancelled", but if a user sets `onDelete` to `"drain"` in the configuration, you may experience a long wait for your destroy to complete.
//
// {{% /examples %}}
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/dataflow_job.html.markdown.
type Job struct {
	pulumi.CustomResourceState

	// The configuration for VM IPs.  Options are `"WORKER_IP_PUBLIC"` or `"WORKER_IP_PRIVATE"`.
	IpConfiguration pulumi.StringPtrOutput `pulumi:"ipConfiguration"`
	// The unique ID of this job.
	JobId pulumi.StringOutput `pulumi:"jobId"`
	// User labels to be specified for the job. Keys and values should follow the restrictions
	// specified in the [labeling restrictions](https://cloud.google.com/compute/docs/labeling-resources#restrictions) page.
	// **NOTE**: Google-provided Dataflow templates often provide default labels that begin with `goog-dataflow-provided`.
	// Unless explicitly set in config, these labels will be ignored to prevent diffs on re-apply.
	Labels pulumi.MapOutput `pulumi:"labels"`
	// The machine type to use for the job.
	MachineType pulumi.StringPtrOutput `pulumi:"machineType"`
	// The number of workers permitted to work on the job.  More workers may improve processing speed at additional cost.
	MaxWorkers pulumi.IntPtrOutput `pulumi:"maxWorkers"`
	// A unique name for the resource, required by Dataflow.
	Name pulumi.StringOutput `pulumi:"name"`
	// The network to which VMs will be assigned. If it is not provided, "default" will be used.
	Network pulumi.StringPtrOutput `pulumi:"network"`
	// One of "drain" or "cancel".  Specifies behavior of deletion during a destroy.  See above note.
	OnDelete pulumi.StringPtrOutput `pulumi:"onDelete"`
	// Key/Value pairs to be passed to the Dataflow job (as used in the template).
	Parameters pulumi.MapOutput `pulumi:"parameters"`
	// The project in which the resource belongs. If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	Region pulumi.StringPtrOutput `pulumi:"region"`
	// The Service Account email used to create the job.
	ServiceAccountEmail pulumi.StringPtrOutput `pulumi:"serviceAccountEmail"`
	// The current state of the resource, selected from the [JobState enum](https://cloud.google.com/dataflow/docs/reference/rest/v1b3/projects.jobs#Job.JobState)
	State pulumi.StringOutput `pulumi:"state"`
	// The subnetwork to which VMs will be assigned. Should be of the form "regions/REGION/subnetworks/SUBNETWORK".
	Subnetwork pulumi.StringPtrOutput `pulumi:"subnetwork"`
	// A writeable location on GCS for the Dataflow job to dump its temporary data.
	TempGcsLocation pulumi.StringOutput `pulumi:"tempGcsLocation"`
	// The GCS path to the Dataflow job template.
	TemplateGcsPath pulumi.StringOutput `pulumi:"templateGcsPath"`
	// The type of this job, selected from the [JobType enum](https://cloud.google.com/dataflow/docs/reference/rest/v1b3/projects.jobs#Job.JobType)
	Type pulumi.StringOutput `pulumi:"type"`
	// The zone in which the created job should run. If it is not provided, the provider zone is used.
	Zone pulumi.StringPtrOutput `pulumi:"zone"`
}

// NewJob registers a new resource with the given unique name, arguments, and options.
func NewJob(ctx *pulumi.Context,
	name string, args *JobArgs, opts ...pulumi.ResourceOption) (*Job, error) {
	if args == nil || args.TempGcsLocation == nil {
		return nil, errors.New("missing required argument 'TempGcsLocation'")
	}
	if args == nil || args.TemplateGcsPath == nil {
		return nil, errors.New("missing required argument 'TemplateGcsPath'")
	}
	if args == nil {
		args = &JobArgs{}
	}
	var resource Job
	err := ctx.RegisterResource("gcp:dataflow/job:Job", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetJob gets an existing Job resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetJob(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *JobState, opts ...pulumi.ResourceOption) (*Job, error) {
	var resource Job
	err := ctx.ReadResource("gcp:dataflow/job:Job", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Job resources.
type jobState struct {
	// The configuration for VM IPs.  Options are `"WORKER_IP_PUBLIC"` or `"WORKER_IP_PRIVATE"`.
	IpConfiguration *string `pulumi:"ipConfiguration"`
	// The unique ID of this job.
	JobId *string `pulumi:"jobId"`
	// User labels to be specified for the job. Keys and values should follow the restrictions
	// specified in the [labeling restrictions](https://cloud.google.com/compute/docs/labeling-resources#restrictions) page.
	// **NOTE**: Google-provided Dataflow templates often provide default labels that begin with `goog-dataflow-provided`.
	// Unless explicitly set in config, these labels will be ignored to prevent diffs on re-apply.
	Labels map[string]interface{} `pulumi:"labels"`
	// The machine type to use for the job.
	MachineType *string `pulumi:"machineType"`
	// The number of workers permitted to work on the job.  More workers may improve processing speed at additional cost.
	MaxWorkers *int `pulumi:"maxWorkers"`
	// A unique name for the resource, required by Dataflow.
	Name *string `pulumi:"name"`
	// The network to which VMs will be assigned. If it is not provided, "default" will be used.
	Network *string `pulumi:"network"`
	// One of "drain" or "cancel".  Specifies behavior of deletion during a destroy.  See above note.
	OnDelete *string `pulumi:"onDelete"`
	// Key/Value pairs to be passed to the Dataflow job (as used in the template).
	Parameters map[string]interface{} `pulumi:"parameters"`
	// The project in which the resource belongs. If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	Region *string `pulumi:"region"`
	// The Service Account email used to create the job.
	ServiceAccountEmail *string `pulumi:"serviceAccountEmail"`
	// The current state of the resource, selected from the [JobState enum](https://cloud.google.com/dataflow/docs/reference/rest/v1b3/projects.jobs#Job.JobState)
	State *string `pulumi:"state"`
	// The subnetwork to which VMs will be assigned. Should be of the form "regions/REGION/subnetworks/SUBNETWORK".
	Subnetwork *string `pulumi:"subnetwork"`
	// A writeable location on GCS for the Dataflow job to dump its temporary data.
	TempGcsLocation *string `pulumi:"tempGcsLocation"`
	// The GCS path to the Dataflow job template.
	TemplateGcsPath *string `pulumi:"templateGcsPath"`
	// The type of this job, selected from the [JobType enum](https://cloud.google.com/dataflow/docs/reference/rest/v1b3/projects.jobs#Job.JobType)
	Type *string `pulumi:"type"`
	// The zone in which the created job should run. If it is not provided, the provider zone is used.
	Zone *string `pulumi:"zone"`
}

type JobState struct {
	// The configuration for VM IPs.  Options are `"WORKER_IP_PUBLIC"` or `"WORKER_IP_PRIVATE"`.
	IpConfiguration pulumi.StringPtrInput
	// The unique ID of this job.
	JobId pulumi.StringPtrInput
	// User labels to be specified for the job. Keys and values should follow the restrictions
	// specified in the [labeling restrictions](https://cloud.google.com/compute/docs/labeling-resources#restrictions) page.
	// **NOTE**: Google-provided Dataflow templates often provide default labels that begin with `goog-dataflow-provided`.
	// Unless explicitly set in config, these labels will be ignored to prevent diffs on re-apply.
	Labels pulumi.MapInput
	// The machine type to use for the job.
	MachineType pulumi.StringPtrInput
	// The number of workers permitted to work on the job.  More workers may improve processing speed at additional cost.
	MaxWorkers pulumi.IntPtrInput
	// A unique name for the resource, required by Dataflow.
	Name pulumi.StringPtrInput
	// The network to which VMs will be assigned. If it is not provided, "default" will be used.
	Network pulumi.StringPtrInput
	// One of "drain" or "cancel".  Specifies behavior of deletion during a destroy.  See above note.
	OnDelete pulumi.StringPtrInput
	// Key/Value pairs to be passed to the Dataflow job (as used in the template).
	Parameters pulumi.MapInput
	// The project in which the resource belongs. If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	Region pulumi.StringPtrInput
	// The Service Account email used to create the job.
	ServiceAccountEmail pulumi.StringPtrInput
	// The current state of the resource, selected from the [JobState enum](https://cloud.google.com/dataflow/docs/reference/rest/v1b3/projects.jobs#Job.JobState)
	State pulumi.StringPtrInput
	// The subnetwork to which VMs will be assigned. Should be of the form "regions/REGION/subnetworks/SUBNETWORK".
	Subnetwork pulumi.StringPtrInput
	// A writeable location on GCS for the Dataflow job to dump its temporary data.
	TempGcsLocation pulumi.StringPtrInput
	// The GCS path to the Dataflow job template.
	TemplateGcsPath pulumi.StringPtrInput
	// The type of this job, selected from the [JobType enum](https://cloud.google.com/dataflow/docs/reference/rest/v1b3/projects.jobs#Job.JobType)
	Type pulumi.StringPtrInput
	// The zone in which the created job should run. If it is not provided, the provider zone is used.
	Zone pulumi.StringPtrInput
}

func (JobState) ElementType() reflect.Type {
	return reflect.TypeOf((*jobState)(nil)).Elem()
}

type jobArgs struct {
	// The configuration for VM IPs.  Options are `"WORKER_IP_PUBLIC"` or `"WORKER_IP_PRIVATE"`.
	IpConfiguration *string `pulumi:"ipConfiguration"`
	// User labels to be specified for the job. Keys and values should follow the restrictions
	// specified in the [labeling restrictions](https://cloud.google.com/compute/docs/labeling-resources#restrictions) page.
	// **NOTE**: Google-provided Dataflow templates often provide default labels that begin with `goog-dataflow-provided`.
	// Unless explicitly set in config, these labels will be ignored to prevent diffs on re-apply.
	Labels map[string]interface{} `pulumi:"labels"`
	// The machine type to use for the job.
	MachineType *string `pulumi:"machineType"`
	// The number of workers permitted to work on the job.  More workers may improve processing speed at additional cost.
	MaxWorkers *int `pulumi:"maxWorkers"`
	// A unique name for the resource, required by Dataflow.
	Name *string `pulumi:"name"`
	// The network to which VMs will be assigned. If it is not provided, "default" will be used.
	Network *string `pulumi:"network"`
	// One of "drain" or "cancel".  Specifies behavior of deletion during a destroy.  See above note.
	OnDelete *string `pulumi:"onDelete"`
	// Key/Value pairs to be passed to the Dataflow job (as used in the template).
	Parameters map[string]interface{} `pulumi:"parameters"`
	// The project in which the resource belongs. If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	Region *string `pulumi:"region"`
	// The Service Account email used to create the job.
	ServiceAccountEmail *string `pulumi:"serviceAccountEmail"`
	// The subnetwork to which VMs will be assigned. Should be of the form "regions/REGION/subnetworks/SUBNETWORK".
	Subnetwork *string `pulumi:"subnetwork"`
	// A writeable location on GCS for the Dataflow job to dump its temporary data.
	TempGcsLocation string `pulumi:"tempGcsLocation"`
	// The GCS path to the Dataflow job template.
	TemplateGcsPath string `pulumi:"templateGcsPath"`
	// The zone in which the created job should run. If it is not provided, the provider zone is used.
	Zone *string `pulumi:"zone"`
}

// The set of arguments for constructing a Job resource.
type JobArgs struct {
	// The configuration for VM IPs.  Options are `"WORKER_IP_PUBLIC"` or `"WORKER_IP_PRIVATE"`.
	IpConfiguration pulumi.StringPtrInput
	// User labels to be specified for the job. Keys and values should follow the restrictions
	// specified in the [labeling restrictions](https://cloud.google.com/compute/docs/labeling-resources#restrictions) page.
	// **NOTE**: Google-provided Dataflow templates often provide default labels that begin with `goog-dataflow-provided`.
	// Unless explicitly set in config, these labels will be ignored to prevent diffs on re-apply.
	Labels pulumi.MapInput
	// The machine type to use for the job.
	MachineType pulumi.StringPtrInput
	// The number of workers permitted to work on the job.  More workers may improve processing speed at additional cost.
	MaxWorkers pulumi.IntPtrInput
	// A unique name for the resource, required by Dataflow.
	Name pulumi.StringPtrInput
	// The network to which VMs will be assigned. If it is not provided, "default" will be used.
	Network pulumi.StringPtrInput
	// One of "drain" or "cancel".  Specifies behavior of deletion during a destroy.  See above note.
	OnDelete pulumi.StringPtrInput
	// Key/Value pairs to be passed to the Dataflow job (as used in the template).
	Parameters pulumi.MapInput
	// The project in which the resource belongs. If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	Region pulumi.StringPtrInput
	// The Service Account email used to create the job.
	ServiceAccountEmail pulumi.StringPtrInput
	// The subnetwork to which VMs will be assigned. Should be of the form "regions/REGION/subnetworks/SUBNETWORK".
	Subnetwork pulumi.StringPtrInput
	// A writeable location on GCS for the Dataflow job to dump its temporary data.
	TempGcsLocation pulumi.StringInput
	// The GCS path to the Dataflow job template.
	TemplateGcsPath pulumi.StringInput
	// The zone in which the created job should run. If it is not provided, the provider zone is used.
	Zone pulumi.StringPtrInput
}

func (JobArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*jobArgs)(nil)).Elem()
}

