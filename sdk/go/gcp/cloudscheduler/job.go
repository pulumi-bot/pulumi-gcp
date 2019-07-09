// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cloudscheduler

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// A scheduled job that can publish a pubsub message or a http request
// every X interval of time, using crontab format string.
// 
// To use Cloud Scheduler your project must contain an App Engine app
// that is located in one of the supported regions. If your project
// does not have an App Engine app, you must create one.
// 
// 
// To get more information about Job, see:
// 
// * [API documentation](https://cloud.google.com/scheduler/docs/reference/rest/)
// * How-to Guides
//     * [Official Documentation](https://cloud.google.com/scheduler/)
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/cloud_scheduler_job.html.markdown.
type Job struct {
	s *pulumi.ResourceState
}

// NewJob registers a new resource with the given unique name, arguments, and options.
func NewJob(ctx *pulumi.Context,
	name string, args *JobArgs, opts ...pulumi.ResourceOpt) (*Job, error) {
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["appEngineHttpTarget"] = nil
		inputs["description"] = nil
		inputs["httpTarget"] = nil
		inputs["name"] = nil
		inputs["project"] = nil
		inputs["pubsubTarget"] = nil
		inputs["region"] = nil
		inputs["retryConfig"] = nil
		inputs["schedule"] = nil
		inputs["timeZone"] = nil
	} else {
		inputs["appEngineHttpTarget"] = args.AppEngineHttpTarget
		inputs["description"] = args.Description
		inputs["httpTarget"] = args.HttpTarget
		inputs["name"] = args.Name
		inputs["project"] = args.Project
		inputs["pubsubTarget"] = args.PubsubTarget
		inputs["region"] = args.Region
		inputs["retryConfig"] = args.RetryConfig
		inputs["schedule"] = args.Schedule
		inputs["timeZone"] = args.TimeZone
	}
	s, err := ctx.RegisterResource("gcp:cloudscheduler/job:Job", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Job{s: s}, nil
}

// GetJob gets an existing Job resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetJob(ctx *pulumi.Context,
	name string, id pulumi.ID, state *JobState, opts ...pulumi.ResourceOpt) (*Job, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["appEngineHttpTarget"] = state.AppEngineHttpTarget
		inputs["description"] = state.Description
		inputs["httpTarget"] = state.HttpTarget
		inputs["name"] = state.Name
		inputs["project"] = state.Project
		inputs["pubsubTarget"] = state.PubsubTarget
		inputs["region"] = state.Region
		inputs["retryConfig"] = state.RetryConfig
		inputs["schedule"] = state.Schedule
		inputs["timeZone"] = state.TimeZone
	}
	s, err := ctx.ReadResource("gcp:cloudscheduler/job:Job", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Job{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Job) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Job) ID() *pulumi.IDOutput {
	return r.s.ID()
}

func (r *Job) AppEngineHttpTarget() *pulumi.Output {
	return r.s.State["appEngineHttpTarget"]
}

func (r *Job) Description() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["description"])
}

func (r *Job) HttpTarget() *pulumi.Output {
	return r.s.State["httpTarget"]
}

func (r *Job) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// The ID of the project in which the resource belongs.
// If it is not provided, the provider project is used.
func (r *Job) Project() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["project"])
}

func (r *Job) PubsubTarget() *pulumi.Output {
	return r.s.State["pubsubTarget"]
}

func (r *Job) Region() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["region"])
}

func (r *Job) RetryConfig() *pulumi.Output {
	return r.s.State["retryConfig"]
}

func (r *Job) Schedule() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["schedule"])
}

func (r *Job) TimeZone() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["timeZone"])
}

// Input properties used for looking up and filtering Job resources.
type JobState struct {
	AppEngineHttpTarget interface{}
	Description interface{}
	HttpTarget interface{}
	Name interface{}
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project interface{}
	PubsubTarget interface{}
	Region interface{}
	RetryConfig interface{}
	Schedule interface{}
	TimeZone interface{}
}

// The set of arguments for constructing a Job resource.
type JobArgs struct {
	AppEngineHttpTarget interface{}
	Description interface{}
	HttpTarget interface{}
	Name interface{}
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project interface{}
	PubsubTarget interface{}
	Region interface{}
	RetryConfig interface{}
	Schedule interface{}
	TimeZone interface{}
}
