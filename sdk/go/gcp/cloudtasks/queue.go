// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cloudtasks

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// A named resource to which messages are sent by publishers.
//
// ## Example Usage
type Queue struct {
	pulumi.CustomResourceState

	// Overrides for task-level appEngineRouting. These settings apply only
	// to App Engine tasks in this queue
	// Structure is documented below.
	AppEngineRoutingOverride QueueAppEngineRoutingOverridePtrOutput `pulumi:"appEngineRoutingOverride"`
	// The location of the queue
	Location pulumi.StringOutput `pulumi:"location"`
	// The queue name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// Rate limits for task dispatches.
	// The queue's actual dispatch rate is the result of:
	// * Number of tasks in the queue
	// * User-specified throttling: rateLimits, retryConfig, and the queue's state.
	// * System throttling due to 429 (Too Many Requests) or 503 (Service
	//   Unavailable) responses from the worker, high error rates, or to
	//   smooth sudden large traffic spikes.
	//   Structure is documented below.
	RateLimits QueueRateLimitsOutput `pulumi:"rateLimits"`
	// Settings that determine the retry behavior.
	// Structure is documented below.
	RetryConfig QueueRetryConfigOutput `pulumi:"retryConfig"`
	// Configuration options for writing logs to Stackdriver Logging.
	// Structure is documented below.
	StackdriverLoggingConfig QueueStackdriverLoggingConfigPtrOutput `pulumi:"stackdriverLoggingConfig"`
}

// NewQueue registers a new resource with the given unique name, arguments, and options.
func NewQueue(ctx *pulumi.Context,
	name string, args *QueueArgs, opts ...pulumi.ResourceOption) (*Queue, error) {
	if args == nil || args.Location == nil {
		return nil, errors.New("missing required argument 'Location'")
	}
	if args == nil {
		args = &QueueArgs{}
	}
	var resource Queue
	err := ctx.RegisterResource("gcp:cloudtasks/queue:Queue", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetQueue gets an existing Queue resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetQueue(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *QueueState, opts ...pulumi.ResourceOption) (*Queue, error) {
	var resource Queue
	err := ctx.ReadResource("gcp:cloudtasks/queue:Queue", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Queue resources.
type queueState struct {
	// Overrides for task-level appEngineRouting. These settings apply only
	// to App Engine tasks in this queue
	// Structure is documented below.
	AppEngineRoutingOverride *QueueAppEngineRoutingOverride `pulumi:"appEngineRoutingOverride"`
	// The location of the queue
	Location *string `pulumi:"location"`
	// The queue name.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// Rate limits for task dispatches.
	// The queue's actual dispatch rate is the result of:
	// * Number of tasks in the queue
	// * User-specified throttling: rateLimits, retryConfig, and the queue's state.
	// * System throttling due to 429 (Too Many Requests) or 503 (Service
	//   Unavailable) responses from the worker, high error rates, or to
	//   smooth sudden large traffic spikes.
	//   Structure is documented below.
	RateLimits *QueueRateLimits `pulumi:"rateLimits"`
	// Settings that determine the retry behavior.
	// Structure is documented below.
	RetryConfig *QueueRetryConfig `pulumi:"retryConfig"`
	// Configuration options for writing logs to Stackdriver Logging.
	// Structure is documented below.
	StackdriverLoggingConfig *QueueStackdriverLoggingConfig `pulumi:"stackdriverLoggingConfig"`
}

type QueueState struct {
	// Overrides for task-level appEngineRouting. These settings apply only
	// to App Engine tasks in this queue
	// Structure is documented below.
	AppEngineRoutingOverride QueueAppEngineRoutingOverridePtrInput
	// The location of the queue
	Location pulumi.StringPtrInput
	// The queue name.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// Rate limits for task dispatches.
	// The queue's actual dispatch rate is the result of:
	// * Number of tasks in the queue
	// * User-specified throttling: rateLimits, retryConfig, and the queue's state.
	// * System throttling due to 429 (Too Many Requests) or 503 (Service
	//   Unavailable) responses from the worker, high error rates, or to
	//   smooth sudden large traffic spikes.
	//   Structure is documented below.
	RateLimits QueueRateLimitsPtrInput
	// Settings that determine the retry behavior.
	// Structure is documented below.
	RetryConfig QueueRetryConfigPtrInput
	// Configuration options for writing logs to Stackdriver Logging.
	// Structure is documented below.
	StackdriverLoggingConfig QueueStackdriverLoggingConfigPtrInput
}

func (QueueState) ElementType() reflect.Type {
	return reflect.TypeOf((*queueState)(nil)).Elem()
}

type queueArgs struct {
	// Overrides for task-level appEngineRouting. These settings apply only
	// to App Engine tasks in this queue
	// Structure is documented below.
	AppEngineRoutingOverride *QueueAppEngineRoutingOverride `pulumi:"appEngineRoutingOverride"`
	// The location of the queue
	Location string `pulumi:"location"`
	// The queue name.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// Rate limits for task dispatches.
	// The queue's actual dispatch rate is the result of:
	// * Number of tasks in the queue
	// * User-specified throttling: rateLimits, retryConfig, and the queue's state.
	// * System throttling due to 429 (Too Many Requests) or 503 (Service
	//   Unavailable) responses from the worker, high error rates, or to
	//   smooth sudden large traffic spikes.
	//   Structure is documented below.
	RateLimits *QueueRateLimits `pulumi:"rateLimits"`
	// Settings that determine the retry behavior.
	// Structure is documented below.
	RetryConfig *QueueRetryConfig `pulumi:"retryConfig"`
	// Configuration options for writing logs to Stackdriver Logging.
	// Structure is documented below.
	StackdriverLoggingConfig *QueueStackdriverLoggingConfig `pulumi:"stackdriverLoggingConfig"`
}

// The set of arguments for constructing a Queue resource.
type QueueArgs struct {
	// Overrides for task-level appEngineRouting. These settings apply only
	// to App Engine tasks in this queue
	// Structure is documented below.
	AppEngineRoutingOverride QueueAppEngineRoutingOverridePtrInput
	// The location of the queue
	Location pulumi.StringInput
	// The queue name.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// Rate limits for task dispatches.
	// The queue's actual dispatch rate is the result of:
	// * Number of tasks in the queue
	// * User-specified throttling: rateLimits, retryConfig, and the queue's state.
	// * System throttling due to 429 (Too Many Requests) or 503 (Service
	//   Unavailable) responses from the worker, high error rates, or to
	//   smooth sudden large traffic spikes.
	//   Structure is documented below.
	RateLimits QueueRateLimitsPtrInput
	// Settings that determine the retry behavior.
	// Structure is documented below.
	RetryConfig QueueRetryConfigPtrInput
	// Configuration options for writing logs to Stackdriver Logging.
	// Structure is documented below.
	StackdriverLoggingConfig QueueStackdriverLoggingConfigPtrInput
}

func (QueueArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*queueArgs)(nil)).Elem()
}

type QueueInput interface {
	pulumi.Input

	ToQueueOutput() QueueOutput
	ToQueueOutputWithContext(ctx context.Context) QueueOutput
}

func (Queue) ElementType() reflect.Type {
	return reflect.TypeOf((*Queue)(nil)).Elem()
}

func (i Queue) ToQueueOutput() QueueOutput {
	return i.ToQueueOutputWithContext(context.Background())
}

func (i Queue) ToQueueOutputWithContext(ctx context.Context) QueueOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueueOutput)
}

type QueueOutput struct {
	*pulumi.OutputState
}

func (QueueOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*QueueOutput)(nil)).Elem()
}

func (o QueueOutput) ToQueueOutput() QueueOutput {
	return o
}

func (o QueueOutput) ToQueueOutputWithContext(ctx context.Context) QueueOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(QueueOutput{})
}
