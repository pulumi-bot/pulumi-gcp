// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package function

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
	"https:/github.com/pulumi/pulumi-gcp/cloudfunctions/FunctionEventTrigger"
	"https:/github.com/pulumi/pulumi-gcp/cloudfunctions/FunctionEventTriggerFailurePolicy"
	"https:/github.com/pulumi/pulumi-gcp/cloudfunctions/FunctionSourceRepository"
)

// Creates a new Cloud Function. For more information see
// [the official documentation](https://cloud.google.com/functions/docs/)
// and
// [API](https://cloud.google.com/functions/docs/apis).
// 
// > **Warning:** As of November 1, 2019, newly created Functions are
// private-by-default and will require [appropriate IAM permissions](https://cloud.google.com/functions/docs/reference/iam/roles)
// to be invoked. See below examples for how to set up the appropriate permissions,
// or view the [Cloud Functions IAM resources](https://www.terraform.io/docs/providers/google/r/cloudfunctions_cloud_function_iam.html)
// for Cloud Functions.
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/cloudfunctions_function.html.markdown.
type Function struct {
	pulumi.CustomResourceState

	// Memory (in MB), available to the function. Default value is 256MB. Allowed values are: 128MB, 256MB, 512MB, 1024MB, and 2048MB.
	AvailableMemoryMb pulumi.IntPtrOutput `pulumi:"availableMemoryMb"`
	// Description of the function.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Name of the function that will be executed when the Google Cloud Function is triggered.
	EntryPoint pulumi.StringPtrOutput `pulumi:"entryPoint"`
	// A set of key/value environment variable pairs to assign to the function.
	EnvironmentVariables pulumi.MapOutput `pulumi:"environmentVariables"`
	// A source that fires events in response to a condition in another service. Structure is documented below. Cannot be used with `triggerHttp`.
	EventTrigger cloudfunctionsFunctionEventTrigger.FunctionEventTriggerOutput `pulumi:"eventTrigger"`
	// URL which triggers function execution. Returned only if `triggerHttp` is used.
	HttpsTriggerUrl pulumi.StringOutput `pulumi:"httpsTriggerUrl"`
	// A set of key/value label pairs to assign to the function.
	Labels pulumi.MapOutput `pulumi:"labels"`
	// The limit on the maximum number of function instances that may coexist at a given time.
	MaxInstances pulumi.IntPtrOutput `pulumi:"maxInstances"`
	// A user-defined name of the function. Function names must be unique globally.
	Name pulumi.StringOutput `pulumi:"name"`
	// Project of the function. If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// Region of function. Currently can be only "us-central1". If it is not provided, the provider region is used.
	Region pulumi.StringOutput `pulumi:"region"`
	// The runtime in which the function is going to run.
	// Eg. `"nodejs8"`, `"nodejs10"`, `"python37"`, `"go111"`.
	Runtime pulumi.StringOutput `pulumi:"runtime"`
	// If provided, the self-provided service account to run the function with.
	ServiceAccountEmail pulumi.StringOutput `pulumi:"serviceAccountEmail"`
	// The GCS bucket containing the zip archive which contains the function.
	SourceArchiveBucket pulumi.StringPtrOutput `pulumi:"sourceArchiveBucket"`
	// The source archive object (file) in archive bucket.
	SourceArchiveObject pulumi.StringPtrOutput `pulumi:"sourceArchiveObject"`
	// Represents parameters related to source repository where a function is hosted.
	// Cannot be set alongside `sourceArchiveBucket` or `sourceArchiveObject`. Structure is documented below.
	SourceRepository cloudfunctionsFunctionSourceRepository.FunctionSourceRepositoryPtrOutput `pulumi:"sourceRepository"`
	// Timeout (in seconds) for the function. Default value is 60 seconds. Cannot be more than 540 seconds.
	Timeout pulumi.IntPtrOutput `pulumi:"timeout"`
	// Boolean variable. Any HTTP request (of a supported type) to the endpoint will trigger function execution. Supported HTTP request types are: POST, PUT, GET, DELETE, and OPTIONS. Endpoint is returned as `httpsTriggerUrl`. Cannot be used with `triggerBucket` and `triggerTopic`.
	TriggerHttp pulumi.BoolPtrOutput `pulumi:"triggerHttp"`
	// The VPC Network Connector that this cloud function can connect to. It can be either the fully-qualified URI, or the short name of the network connector resource. The format of this field is `projects/*/locations/*/connectors/*`.
	VpcConnector pulumi.StringPtrOutput `pulumi:"vpcConnector"`
}

// NewFunction registers a new resource with the given unique name, arguments, and options.
func NewFunction(ctx *pulumi.Context,
	name string, args *FunctionArgs, opts ...pulumi.ResourceOption) (*Function, error) {
	if args == nil || args.Runtime == nil {
		return nil, errors.New("missing required argument 'Runtime'")
	}
	if args == nil {
		args = &FunctionArgs{}
	}
	var resource Function
	err := ctx.RegisterResource("gcp:cloudfunctions/function:Function", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFunction gets an existing Function resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFunction(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FunctionState, opts ...pulumi.ResourceOption) (*Function, error) {
	var resource Function
	err := ctx.ReadResource("gcp:cloudfunctions/function:Function", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Function resources.
type functionState struct {
	// Memory (in MB), available to the function. Default value is 256MB. Allowed values are: 128MB, 256MB, 512MB, 1024MB, and 2048MB.
	AvailableMemoryMb *int `pulumi:"availableMemoryMb"`
	// Description of the function.
	Description *string `pulumi:"description"`
	// Name of the function that will be executed when the Google Cloud Function is triggered.
	EntryPoint *string `pulumi:"entryPoint"`
	// A set of key/value environment variable pairs to assign to the function.
	EnvironmentVariables map[string]interface{} `pulumi:"environmentVariables"`
	// A source that fires events in response to a condition in another service. Structure is documented below. Cannot be used with `triggerHttp`.
	EventTrigger *cloudfunctionsFunctionEventTrigger.FunctionEventTrigger `pulumi:"eventTrigger"`
	// URL which triggers function execution. Returned only if `triggerHttp` is used.
	HttpsTriggerUrl *string `pulumi:"httpsTriggerUrl"`
	// A set of key/value label pairs to assign to the function.
	Labels map[string]interface{} `pulumi:"labels"`
	// The limit on the maximum number of function instances that may coexist at a given time.
	MaxInstances *int `pulumi:"maxInstances"`
	// A user-defined name of the function. Function names must be unique globally.
	Name *string `pulumi:"name"`
	// Project of the function. If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// Region of function. Currently can be only "us-central1". If it is not provided, the provider region is used.
	Region *string `pulumi:"region"`
	// The runtime in which the function is going to run.
	// Eg. `"nodejs8"`, `"nodejs10"`, `"python37"`, `"go111"`.
	Runtime *string `pulumi:"runtime"`
	// If provided, the self-provided service account to run the function with.
	ServiceAccountEmail *string `pulumi:"serviceAccountEmail"`
	// The GCS bucket containing the zip archive which contains the function.
	SourceArchiveBucket *string `pulumi:"sourceArchiveBucket"`
	// The source archive object (file) in archive bucket.
	SourceArchiveObject *string `pulumi:"sourceArchiveObject"`
	// Represents parameters related to source repository where a function is hosted.
	// Cannot be set alongside `sourceArchiveBucket` or `sourceArchiveObject`. Structure is documented below.
	SourceRepository *cloudfunctionsFunctionSourceRepository.FunctionSourceRepository `pulumi:"sourceRepository"`
	// Timeout (in seconds) for the function. Default value is 60 seconds. Cannot be more than 540 seconds.
	Timeout *int `pulumi:"timeout"`
	// Boolean variable. Any HTTP request (of a supported type) to the endpoint will trigger function execution. Supported HTTP request types are: POST, PUT, GET, DELETE, and OPTIONS. Endpoint is returned as `httpsTriggerUrl`. Cannot be used with `triggerBucket` and `triggerTopic`.
	TriggerHttp *bool `pulumi:"triggerHttp"`
	// The VPC Network Connector that this cloud function can connect to. It can be either the fully-qualified URI, or the short name of the network connector resource. The format of this field is `projects/*/locations/*/connectors/*`.
	VpcConnector *string `pulumi:"vpcConnector"`
}

type FunctionState struct {
	// Memory (in MB), available to the function. Default value is 256MB. Allowed values are: 128MB, 256MB, 512MB, 1024MB, and 2048MB.
	AvailableMemoryMb pulumi.IntPtrInput
	// Description of the function.
	Description pulumi.StringPtrInput
	// Name of the function that will be executed when the Google Cloud Function is triggered.
	EntryPoint pulumi.StringPtrInput
	// A set of key/value environment variable pairs to assign to the function.
	EnvironmentVariables pulumi.MapInput
	// A source that fires events in response to a condition in another service. Structure is documented below. Cannot be used with `triggerHttp`.
	EventTrigger cloudfunctionsFunctionEventTrigger.FunctionEventTriggerPtrInput
	// URL which triggers function execution. Returned only if `triggerHttp` is used.
	HttpsTriggerUrl pulumi.StringPtrInput
	// A set of key/value label pairs to assign to the function.
	Labels pulumi.MapInput
	// The limit on the maximum number of function instances that may coexist at a given time.
	MaxInstances pulumi.IntPtrInput
	// A user-defined name of the function. Function names must be unique globally.
	Name pulumi.StringPtrInput
	// Project of the function. If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// Region of function. Currently can be only "us-central1". If it is not provided, the provider region is used.
	Region pulumi.StringPtrInput
	// The runtime in which the function is going to run.
	// Eg. `"nodejs8"`, `"nodejs10"`, `"python37"`, `"go111"`.
	Runtime pulumi.StringPtrInput
	// If provided, the self-provided service account to run the function with.
	ServiceAccountEmail pulumi.StringPtrInput
	// The GCS bucket containing the zip archive which contains the function.
	SourceArchiveBucket pulumi.StringPtrInput
	// The source archive object (file) in archive bucket.
	SourceArchiveObject pulumi.StringPtrInput
	// Represents parameters related to source repository where a function is hosted.
	// Cannot be set alongside `sourceArchiveBucket` or `sourceArchiveObject`. Structure is documented below.
	SourceRepository cloudfunctionsFunctionSourceRepository.FunctionSourceRepositoryPtrInput
	// Timeout (in seconds) for the function. Default value is 60 seconds. Cannot be more than 540 seconds.
	Timeout pulumi.IntPtrInput
	// Boolean variable. Any HTTP request (of a supported type) to the endpoint will trigger function execution. Supported HTTP request types are: POST, PUT, GET, DELETE, and OPTIONS. Endpoint is returned as `httpsTriggerUrl`. Cannot be used with `triggerBucket` and `triggerTopic`.
	TriggerHttp pulumi.BoolPtrInput
	// The VPC Network Connector that this cloud function can connect to. It can be either the fully-qualified URI, or the short name of the network connector resource. The format of this field is `projects/*/locations/*/connectors/*`.
	VpcConnector pulumi.StringPtrInput
}

func (FunctionState) ElementType() reflect.Type {
	return reflect.TypeOf((*functionState)(nil)).Elem()
}

type functionArgs struct {
	// Memory (in MB), available to the function. Default value is 256MB. Allowed values are: 128MB, 256MB, 512MB, 1024MB, and 2048MB.
	AvailableMemoryMb *int `pulumi:"availableMemoryMb"`
	// Description of the function.
	Description *string `pulumi:"description"`
	// Name of the function that will be executed when the Google Cloud Function is triggered.
	EntryPoint *string `pulumi:"entryPoint"`
	// A set of key/value environment variable pairs to assign to the function.
	EnvironmentVariables map[string]interface{} `pulumi:"environmentVariables"`
	// A source that fires events in response to a condition in another service. Structure is documented below. Cannot be used with `triggerHttp`.
	EventTrigger *cloudfunctionsFunctionEventTrigger.FunctionEventTrigger `pulumi:"eventTrigger"`
	// URL which triggers function execution. Returned only if `triggerHttp` is used.
	HttpsTriggerUrl *string `pulumi:"httpsTriggerUrl"`
	// A set of key/value label pairs to assign to the function.
	Labels map[string]interface{} `pulumi:"labels"`
	// The limit on the maximum number of function instances that may coexist at a given time.
	MaxInstances *int `pulumi:"maxInstances"`
	// A user-defined name of the function. Function names must be unique globally.
	Name *string `pulumi:"name"`
	// Project of the function. If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// Region of function. Currently can be only "us-central1". If it is not provided, the provider region is used.
	Region *string `pulumi:"region"`
	// The runtime in which the function is going to run.
	// Eg. `"nodejs8"`, `"nodejs10"`, `"python37"`, `"go111"`.
	Runtime string `pulumi:"runtime"`
	// If provided, the self-provided service account to run the function with.
	ServiceAccountEmail *string `pulumi:"serviceAccountEmail"`
	// The GCS bucket containing the zip archive which contains the function.
	SourceArchiveBucket *string `pulumi:"sourceArchiveBucket"`
	// The source archive object (file) in archive bucket.
	SourceArchiveObject *string `pulumi:"sourceArchiveObject"`
	// Represents parameters related to source repository where a function is hosted.
	// Cannot be set alongside `sourceArchiveBucket` or `sourceArchiveObject`. Structure is documented below.
	SourceRepository *cloudfunctionsFunctionSourceRepository.FunctionSourceRepository `pulumi:"sourceRepository"`
	// Timeout (in seconds) for the function. Default value is 60 seconds. Cannot be more than 540 seconds.
	Timeout *int `pulumi:"timeout"`
	// Boolean variable. Any HTTP request (of a supported type) to the endpoint will trigger function execution. Supported HTTP request types are: POST, PUT, GET, DELETE, and OPTIONS. Endpoint is returned as `httpsTriggerUrl`. Cannot be used with `triggerBucket` and `triggerTopic`.
	TriggerHttp *bool `pulumi:"triggerHttp"`
	// The VPC Network Connector that this cloud function can connect to. It can be either the fully-qualified URI, or the short name of the network connector resource. The format of this field is `projects/*/locations/*/connectors/*`.
	VpcConnector *string `pulumi:"vpcConnector"`
}

// The set of arguments for constructing a Function resource.
type FunctionArgs struct {
	// Memory (in MB), available to the function. Default value is 256MB. Allowed values are: 128MB, 256MB, 512MB, 1024MB, and 2048MB.
	AvailableMemoryMb pulumi.IntPtrInput
	// Description of the function.
	Description pulumi.StringPtrInput
	// Name of the function that will be executed when the Google Cloud Function is triggered.
	EntryPoint pulumi.StringPtrInput
	// A set of key/value environment variable pairs to assign to the function.
	EnvironmentVariables pulumi.MapInput
	// A source that fires events in response to a condition in another service. Structure is documented below. Cannot be used with `triggerHttp`.
	EventTrigger cloudfunctionsFunctionEventTrigger.FunctionEventTriggerPtrInput
	// URL which triggers function execution. Returned only if `triggerHttp` is used.
	HttpsTriggerUrl pulumi.StringPtrInput
	// A set of key/value label pairs to assign to the function.
	Labels pulumi.MapInput
	// The limit on the maximum number of function instances that may coexist at a given time.
	MaxInstances pulumi.IntPtrInput
	// A user-defined name of the function. Function names must be unique globally.
	Name pulumi.StringPtrInput
	// Project of the function. If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// Region of function. Currently can be only "us-central1". If it is not provided, the provider region is used.
	Region pulumi.StringPtrInput
	// The runtime in which the function is going to run.
	// Eg. `"nodejs8"`, `"nodejs10"`, `"python37"`, `"go111"`.
	Runtime pulumi.StringInput
	// If provided, the self-provided service account to run the function with.
	ServiceAccountEmail pulumi.StringPtrInput
	// The GCS bucket containing the zip archive which contains the function.
	SourceArchiveBucket pulumi.StringPtrInput
	// The source archive object (file) in archive bucket.
	SourceArchiveObject pulumi.StringPtrInput
	// Represents parameters related to source repository where a function is hosted.
	// Cannot be set alongside `sourceArchiveBucket` or `sourceArchiveObject`. Structure is documented below.
	SourceRepository cloudfunctionsFunctionSourceRepository.FunctionSourceRepositoryPtrInput
	// Timeout (in seconds) for the function. Default value is 60 seconds. Cannot be more than 540 seconds.
	Timeout pulumi.IntPtrInput
	// Boolean variable. Any HTTP request (of a supported type) to the endpoint will trigger function execution. Supported HTTP request types are: POST, PUT, GET, DELETE, and OPTIONS. Endpoint is returned as `httpsTriggerUrl`. Cannot be used with `triggerBucket` and `triggerTopic`.
	TriggerHttp pulumi.BoolPtrInput
	// The VPC Network Connector that this cloud function can connect to. It can be either the fully-qualified URI, or the short name of the network connector resource. The format of this field is `projects/*/locations/*/connectors/*`.
	VpcConnector pulumi.StringPtrInput
}

func (FunctionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*functionArgs)(nil)).Elem()
}

