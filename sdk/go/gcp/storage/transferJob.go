// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package storage

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Creates a new Transfer Job in Google Cloud Storage Transfer.
// 
// To get more information about Google Cloud Storage Transfer, see:
// 
// * [Overview](https://cloud.google.com/storage-transfer/docs/overview)
// * [API documentation](https://cloud.google.com/storage-transfer/docs/reference/rest/v1/transferJobs#TransferJob)
// * How-to Guides
//     * [Configuring Access to Data Sources and Sinks](https://cloud.google.com/storage-transfer/docs/configure-access)
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/storage_transfer_job.html.markdown.
type TransferJob struct {
	s *pulumi.ResourceState
}

// NewTransferJob registers a new resource with the given unique name, arguments, and options.
func NewTransferJob(ctx *pulumi.Context,
	name string, args *TransferJobArgs, opts ...pulumi.ResourceOpt) (*TransferJob, error) {
	if args == nil || args.Description == nil {
		return nil, errors.New("missing required argument 'Description'")
	}
	if args == nil || args.Schedule == nil {
		return nil, errors.New("missing required argument 'Schedule'")
	}
	if args == nil || args.TransferSpec == nil {
		return nil, errors.New("missing required argument 'TransferSpec'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["description"] = nil
		inputs["project"] = nil
		inputs["schedule"] = nil
		inputs["status"] = nil
		inputs["transferSpec"] = nil
	} else {
		inputs["description"] = args.Description
		inputs["project"] = args.Project
		inputs["schedule"] = args.Schedule
		inputs["status"] = args.Status
		inputs["transferSpec"] = args.TransferSpec
	}
	inputs["creationTime"] = nil
	inputs["deletionTime"] = nil
	inputs["lastModificationTime"] = nil
	inputs["name"] = nil
	s, err := ctx.RegisterResource("gcp:storage/transferJob:TransferJob", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &TransferJob{s: s}, nil
}

// GetTransferJob gets an existing TransferJob resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTransferJob(ctx *pulumi.Context,
	name string, id pulumi.ID, state *TransferJobState, opts ...pulumi.ResourceOpt) (*TransferJob, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["creationTime"] = state.CreationTime
		inputs["deletionTime"] = state.DeletionTime
		inputs["description"] = state.Description
		inputs["lastModificationTime"] = state.LastModificationTime
		inputs["name"] = state.Name
		inputs["project"] = state.Project
		inputs["schedule"] = state.Schedule
		inputs["status"] = state.Status
		inputs["transferSpec"] = state.TransferSpec
	}
	s, err := ctx.ReadResource("gcp:storage/transferJob:TransferJob", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &TransferJob{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *TransferJob) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *TransferJob) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// When the Transfer Job was created.
func (r *TransferJob) CreationTime() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["creationTime"])
}

// When the Transfer Job was deleted.
func (r *TransferJob) DeletionTime() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["deletionTime"])
}

// Unique description to identify the Transfer Job.
func (r *TransferJob) Description() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["description"])
}

// When the Transfer Job was last modified.
func (r *TransferJob) LastModificationTime() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["lastModificationTime"])
}

// The name of the Transfer Job.
func (r *TransferJob) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// The project in which the resource belongs. If it
// is not provided, the provider project is used.
func (r *TransferJob) Project() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["project"])
}

// Schedule specification defining when the Transfer Job should be scheduled to start, end and and what time to run. Structure documented below.
func (r *TransferJob) Schedule() *pulumi.Output {
	return r.s.State["schedule"]
}

// Status of the job. Default: `ENABLED`. **NOTE: The effect of the new job status takes place during a subsequent job run. For example, if you change the job status from ENABLED to DISABLED, and an operation spawned by the transfer is running, the status change would not affect the current operation.**
func (r *TransferJob) Status() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["status"])
}

// Transfer specification. Structure documented below.
func (r *TransferJob) TransferSpec() *pulumi.Output {
	return r.s.State["transferSpec"]
}

// Input properties used for looking up and filtering TransferJob resources.
type TransferJobState struct {
	// When the Transfer Job was created.
	CreationTime interface{}
	// When the Transfer Job was deleted.
	DeletionTime interface{}
	// Unique description to identify the Transfer Job.
	Description interface{}
	// When the Transfer Job was last modified.
	LastModificationTime interface{}
	// The name of the Transfer Job.
	Name interface{}
	// The project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project interface{}
	// Schedule specification defining when the Transfer Job should be scheduled to start, end and and what time to run. Structure documented below.
	Schedule interface{}
	// Status of the job. Default: `ENABLED`. **NOTE: The effect of the new job status takes place during a subsequent job run. For example, if you change the job status from ENABLED to DISABLED, and an operation spawned by the transfer is running, the status change would not affect the current operation.**
	Status interface{}
	// Transfer specification. Structure documented below.
	TransferSpec interface{}
}

// The set of arguments for constructing a TransferJob resource.
type TransferJobArgs struct {
	// Unique description to identify the Transfer Job.
	Description interface{}
	// The project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project interface{}
	// Schedule specification defining when the Transfer Job should be scheduled to start, end and and what time to run. Structure documented below.
	Schedule interface{}
	// Status of the job. Default: `ENABLED`. **NOTE: The effect of the new job status takes place during a subsequent job run. For example, if you change the job status from ENABLED to DISABLED, and an operation spawned by the transfer is running, the status change would not affect the current operation.**
	Status interface{}
	// Transfer specification. Structure documented below.
	TransferSpec interface{}
}
