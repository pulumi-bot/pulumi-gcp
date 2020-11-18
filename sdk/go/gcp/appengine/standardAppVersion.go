// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package appengine

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Standard App Version resource to create a new version of standard GAE Application.
// Learn about the differences between the standard environment and the flexible environment
// at https://cloud.google.com/appengine/docs/the-appengine-environments.
// Currently supporting Zip and File Containers.
//
// To get more information about StandardAppVersion, see:
//
// * [API documentation](https://cloud.google.com/appengine/docs/admin-api/reference/rest/v1/apps.services.versions)
// * How-to Guides
//     * [Official Documentation](https://cloud.google.com/appengine/docs/standard)
//
// ## Example Usage
type StandardAppVersion struct {
	pulumi.CustomResourceState

	// Automatic scaling is based on request rate, response latencies, and other application metrics.
	// Structure is documented below.
	AutomaticScaling StandardAppVersionAutomaticScalingPtrOutput `pulumi:"automaticScaling"`
	// Basic scaling creates instances when your application receives requests. Each instance will be shut down when the application becomes idle. Basic scaling is ideal for work that is intermittent or driven by user activity.
	// Structure is documented below.
	BasicScaling StandardAppVersionBasicScalingPtrOutput `pulumi:"basicScaling"`
	// If set to `true`, the service will be deleted if it is the last version.
	DeleteServiceOnDestroy pulumi.BoolPtrOutput `pulumi:"deleteServiceOnDestroy"`
	// Code and application artifacts that make up this version.
	// Structure is documented below.
	Deployment StandardAppVersionDeploymentOutput `pulumi:"deployment"`
	// The entrypoint for the application.
	// Structure is documented below.
	Entrypoint StandardAppVersionEntrypointPtrOutput `pulumi:"entrypoint"`
	// Environment variables available to the application.
	EnvVariables pulumi.StringMapOutput `pulumi:"envVariables"`
	// An ordered list of URL-matching patterns that should be applied to incoming requests.
	// The first matching URL handles the request and other request handlers are not attempted.
	// Structure is documented below.
	Handlers StandardAppVersionHandlerArrayOutput `pulumi:"handlers"`
	// A list of the types of messages that this application is able to receive.
	// Each value may be one of `INBOUND_SERVICE_MAIL`, `INBOUND_SERVICE_MAIL_BOUNCE`, `INBOUND_SERVICE_XMPP_ERROR`, `INBOUND_SERVICE_XMPP_MESSAGE`, `INBOUND_SERVICE_XMPP_SUBSCRIBE`, `INBOUND_SERVICE_XMPP_PRESENCE`, `INBOUND_SERVICE_CHANNEL_PRESENCE`, and `INBOUND_SERVICE_WARMUP`.
	InboundServices pulumi.StringArrayOutput `pulumi:"inboundServices"`
	// Instance class that is used to run this version. Valid values are
	// AutomaticScaling: F1, F2, F4, F4_1G
	// BasicScaling or ManualScaling: B1, B2, B4, B4_1G, B8
	// Defaults to F1 for AutomaticScaling and B2 for ManualScaling and BasicScaling. If no scaling is specified, AutomaticScaling is chosen.
	InstanceClass pulumi.StringOutput `pulumi:"instanceClass"`
	// Configuration for third-party Python runtime libraries that are required by the application.
	// Structure is documented below.
	Libraries StandardAppVersionLibraryArrayOutput `pulumi:"libraries"`
	// A service with manual scaling runs continuously, allowing you to perform complex initialization and rely on the state of its memory over time.
	// Structure is documented below.
	ManualScaling StandardAppVersionManualScalingPtrOutput `pulumi:"manualScaling"`
	// Full Serverless VPC Access Connector name e.g. /projects/my-project/locations/us-central1/connectors/c1.
	Name pulumi.StringOutput `pulumi:"name"`
	// If set to `true`, the application version will not be deleted.
	NoopOnDestroy pulumi.BoolPtrOutput `pulumi:"noopOnDestroy"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// Desired runtime. Example python27.
	Runtime pulumi.StringOutput `pulumi:"runtime"`
	// The version of the API in the given runtime environment.
	// Please see the app.yaml reference for valid values at https://cloud.google.com/appengine/docs/standard//config/appref
	RuntimeApiVersion pulumi.StringPtrOutput `pulumi:"runtimeApiVersion"`
	// AppEngine service resource
	Service pulumi.StringOutput `pulumi:"service"`
	// Whether multiple requests can be dispatched to this version at once.
	Threadsafe pulumi.BoolPtrOutput `pulumi:"threadsafe"`
	// Relative name of the version within the service. For example, `v1`. Version names can contain only lowercase letters, numbers, or hyphens. Reserved names,"default", "latest", and any name with the prefix "ah-".
	VersionId pulumi.StringPtrOutput `pulumi:"versionId"`
	// Enables VPC connectivity for standard apps.
	// Structure is documented below.
	VpcAccessConnector StandardAppVersionVpcAccessConnectorPtrOutput `pulumi:"vpcAccessConnector"`
}

// NewStandardAppVersion registers a new resource with the given unique name, arguments, and options.
func NewStandardAppVersion(ctx *pulumi.Context,
	name string, args *StandardAppVersionArgs, opts ...pulumi.ResourceOption) (*StandardAppVersion, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Deployment == nil {
		return nil, errors.New("invalid value for required argument 'Deployment'")
	}
	if args.Runtime == nil {
		return nil, errors.New("invalid value for required argument 'Runtime'")
	}
	if args.Service == nil {
		return nil, errors.New("invalid value for required argument 'Service'")
	}
	var resource StandardAppVersion
	err := ctx.RegisterResource("gcp:appengine/standardAppVersion:StandardAppVersion", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetStandardAppVersion gets an existing StandardAppVersion resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetStandardAppVersion(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StandardAppVersionState, opts ...pulumi.ResourceOption) (*StandardAppVersion, error) {
	var resource StandardAppVersion
	err := ctx.ReadResource("gcp:appengine/standardAppVersion:StandardAppVersion", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering StandardAppVersion resources.
type standardAppVersionState struct {
	// Automatic scaling is based on request rate, response latencies, and other application metrics.
	// Structure is documented below.
	AutomaticScaling *StandardAppVersionAutomaticScaling `pulumi:"automaticScaling"`
	// Basic scaling creates instances when your application receives requests. Each instance will be shut down when the application becomes idle. Basic scaling is ideal for work that is intermittent or driven by user activity.
	// Structure is documented below.
	BasicScaling *StandardAppVersionBasicScaling `pulumi:"basicScaling"`
	// If set to `true`, the service will be deleted if it is the last version.
	DeleteServiceOnDestroy *bool `pulumi:"deleteServiceOnDestroy"`
	// Code and application artifacts that make up this version.
	// Structure is documented below.
	Deployment *StandardAppVersionDeployment `pulumi:"deployment"`
	// The entrypoint for the application.
	// Structure is documented below.
	Entrypoint *StandardAppVersionEntrypoint `pulumi:"entrypoint"`
	// Environment variables available to the application.
	EnvVariables map[string]string `pulumi:"envVariables"`
	// An ordered list of URL-matching patterns that should be applied to incoming requests.
	// The first matching URL handles the request and other request handlers are not attempted.
	// Structure is documented below.
	Handlers []StandardAppVersionHandler `pulumi:"handlers"`
	// A list of the types of messages that this application is able to receive.
	// Each value may be one of `INBOUND_SERVICE_MAIL`, `INBOUND_SERVICE_MAIL_BOUNCE`, `INBOUND_SERVICE_XMPP_ERROR`, `INBOUND_SERVICE_XMPP_MESSAGE`, `INBOUND_SERVICE_XMPP_SUBSCRIBE`, `INBOUND_SERVICE_XMPP_PRESENCE`, `INBOUND_SERVICE_CHANNEL_PRESENCE`, and `INBOUND_SERVICE_WARMUP`.
	InboundServices []string `pulumi:"inboundServices"`
	// Instance class that is used to run this version. Valid values are
	// AutomaticScaling: F1, F2, F4, F4_1G
	// BasicScaling or ManualScaling: B1, B2, B4, B4_1G, B8
	// Defaults to F1 for AutomaticScaling and B2 for ManualScaling and BasicScaling. If no scaling is specified, AutomaticScaling is chosen.
	InstanceClass *string `pulumi:"instanceClass"`
	// Configuration for third-party Python runtime libraries that are required by the application.
	// Structure is documented below.
	Libraries []StandardAppVersionLibrary `pulumi:"libraries"`
	// A service with manual scaling runs continuously, allowing you to perform complex initialization and rely on the state of its memory over time.
	// Structure is documented below.
	ManualScaling *StandardAppVersionManualScaling `pulumi:"manualScaling"`
	// Full Serverless VPC Access Connector name e.g. /projects/my-project/locations/us-central1/connectors/c1.
	Name *string `pulumi:"name"`
	// If set to `true`, the application version will not be deleted.
	NoopOnDestroy *bool `pulumi:"noopOnDestroy"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// Desired runtime. Example python27.
	Runtime *string `pulumi:"runtime"`
	// The version of the API in the given runtime environment.
	// Please see the app.yaml reference for valid values at https://cloud.google.com/appengine/docs/standard//config/appref
	RuntimeApiVersion *string `pulumi:"runtimeApiVersion"`
	// AppEngine service resource
	Service *string `pulumi:"service"`
	// Whether multiple requests can be dispatched to this version at once.
	Threadsafe *bool `pulumi:"threadsafe"`
	// Relative name of the version within the service. For example, `v1`. Version names can contain only lowercase letters, numbers, or hyphens. Reserved names,"default", "latest", and any name with the prefix "ah-".
	VersionId *string `pulumi:"versionId"`
	// Enables VPC connectivity for standard apps.
	// Structure is documented below.
	VpcAccessConnector *StandardAppVersionVpcAccessConnector `pulumi:"vpcAccessConnector"`
}

type StandardAppVersionState struct {
	// Automatic scaling is based on request rate, response latencies, and other application metrics.
	// Structure is documented below.
	AutomaticScaling StandardAppVersionAutomaticScalingPtrInput
	// Basic scaling creates instances when your application receives requests. Each instance will be shut down when the application becomes idle. Basic scaling is ideal for work that is intermittent or driven by user activity.
	// Structure is documented below.
	BasicScaling StandardAppVersionBasicScalingPtrInput
	// If set to `true`, the service will be deleted if it is the last version.
	DeleteServiceOnDestroy pulumi.BoolPtrInput
	// Code and application artifacts that make up this version.
	// Structure is documented below.
	Deployment StandardAppVersionDeploymentPtrInput
	// The entrypoint for the application.
	// Structure is documented below.
	Entrypoint StandardAppVersionEntrypointPtrInput
	// Environment variables available to the application.
	EnvVariables pulumi.StringMapInput
	// An ordered list of URL-matching patterns that should be applied to incoming requests.
	// The first matching URL handles the request and other request handlers are not attempted.
	// Structure is documented below.
	Handlers StandardAppVersionHandlerArrayInput
	// A list of the types of messages that this application is able to receive.
	// Each value may be one of `INBOUND_SERVICE_MAIL`, `INBOUND_SERVICE_MAIL_BOUNCE`, `INBOUND_SERVICE_XMPP_ERROR`, `INBOUND_SERVICE_XMPP_MESSAGE`, `INBOUND_SERVICE_XMPP_SUBSCRIBE`, `INBOUND_SERVICE_XMPP_PRESENCE`, `INBOUND_SERVICE_CHANNEL_PRESENCE`, and `INBOUND_SERVICE_WARMUP`.
	InboundServices pulumi.StringArrayInput
	// Instance class that is used to run this version. Valid values are
	// AutomaticScaling: F1, F2, F4, F4_1G
	// BasicScaling or ManualScaling: B1, B2, B4, B4_1G, B8
	// Defaults to F1 for AutomaticScaling and B2 for ManualScaling and BasicScaling. If no scaling is specified, AutomaticScaling is chosen.
	InstanceClass pulumi.StringPtrInput
	// Configuration for third-party Python runtime libraries that are required by the application.
	// Structure is documented below.
	Libraries StandardAppVersionLibraryArrayInput
	// A service with manual scaling runs continuously, allowing you to perform complex initialization and rely on the state of its memory over time.
	// Structure is documented below.
	ManualScaling StandardAppVersionManualScalingPtrInput
	// Full Serverless VPC Access Connector name e.g. /projects/my-project/locations/us-central1/connectors/c1.
	Name pulumi.StringPtrInput
	// If set to `true`, the application version will not be deleted.
	NoopOnDestroy pulumi.BoolPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// Desired runtime. Example python27.
	Runtime pulumi.StringPtrInput
	// The version of the API in the given runtime environment.
	// Please see the app.yaml reference for valid values at https://cloud.google.com/appengine/docs/standard//config/appref
	RuntimeApiVersion pulumi.StringPtrInput
	// AppEngine service resource
	Service pulumi.StringPtrInput
	// Whether multiple requests can be dispatched to this version at once.
	Threadsafe pulumi.BoolPtrInput
	// Relative name of the version within the service. For example, `v1`. Version names can contain only lowercase letters, numbers, or hyphens. Reserved names,"default", "latest", and any name with the prefix "ah-".
	VersionId pulumi.StringPtrInput
	// Enables VPC connectivity for standard apps.
	// Structure is documented below.
	VpcAccessConnector StandardAppVersionVpcAccessConnectorPtrInput
}

func (StandardAppVersionState) ElementType() reflect.Type {
	return reflect.TypeOf((*standardAppVersionState)(nil)).Elem()
}

type standardAppVersionArgs struct {
	// Automatic scaling is based on request rate, response latencies, and other application metrics.
	// Structure is documented below.
	AutomaticScaling *StandardAppVersionAutomaticScaling `pulumi:"automaticScaling"`
	// Basic scaling creates instances when your application receives requests. Each instance will be shut down when the application becomes idle. Basic scaling is ideal for work that is intermittent or driven by user activity.
	// Structure is documented below.
	BasicScaling *StandardAppVersionBasicScaling `pulumi:"basicScaling"`
	// If set to `true`, the service will be deleted if it is the last version.
	DeleteServiceOnDestroy *bool `pulumi:"deleteServiceOnDestroy"`
	// Code and application artifacts that make up this version.
	// Structure is documented below.
	Deployment StandardAppVersionDeployment `pulumi:"deployment"`
	// The entrypoint for the application.
	// Structure is documented below.
	Entrypoint *StandardAppVersionEntrypoint `pulumi:"entrypoint"`
	// Environment variables available to the application.
	EnvVariables map[string]string `pulumi:"envVariables"`
	// An ordered list of URL-matching patterns that should be applied to incoming requests.
	// The first matching URL handles the request and other request handlers are not attempted.
	// Structure is documented below.
	Handlers []StandardAppVersionHandler `pulumi:"handlers"`
	// A list of the types of messages that this application is able to receive.
	// Each value may be one of `INBOUND_SERVICE_MAIL`, `INBOUND_SERVICE_MAIL_BOUNCE`, `INBOUND_SERVICE_XMPP_ERROR`, `INBOUND_SERVICE_XMPP_MESSAGE`, `INBOUND_SERVICE_XMPP_SUBSCRIBE`, `INBOUND_SERVICE_XMPP_PRESENCE`, `INBOUND_SERVICE_CHANNEL_PRESENCE`, and `INBOUND_SERVICE_WARMUP`.
	InboundServices []string `pulumi:"inboundServices"`
	// Instance class that is used to run this version. Valid values are
	// AutomaticScaling: F1, F2, F4, F4_1G
	// BasicScaling or ManualScaling: B1, B2, B4, B4_1G, B8
	// Defaults to F1 for AutomaticScaling and B2 for ManualScaling and BasicScaling. If no scaling is specified, AutomaticScaling is chosen.
	InstanceClass *string `pulumi:"instanceClass"`
	// Configuration for third-party Python runtime libraries that are required by the application.
	// Structure is documented below.
	Libraries []StandardAppVersionLibrary `pulumi:"libraries"`
	// A service with manual scaling runs continuously, allowing you to perform complex initialization and rely on the state of its memory over time.
	// Structure is documented below.
	ManualScaling *StandardAppVersionManualScaling `pulumi:"manualScaling"`
	// If set to `true`, the application version will not be deleted.
	NoopOnDestroy *bool `pulumi:"noopOnDestroy"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// Desired runtime. Example python27.
	Runtime string `pulumi:"runtime"`
	// The version of the API in the given runtime environment.
	// Please see the app.yaml reference for valid values at https://cloud.google.com/appengine/docs/standard//config/appref
	RuntimeApiVersion *string `pulumi:"runtimeApiVersion"`
	// AppEngine service resource
	Service string `pulumi:"service"`
	// Whether multiple requests can be dispatched to this version at once.
	Threadsafe *bool `pulumi:"threadsafe"`
	// Relative name of the version within the service. For example, `v1`. Version names can contain only lowercase letters, numbers, or hyphens. Reserved names,"default", "latest", and any name with the prefix "ah-".
	VersionId *string `pulumi:"versionId"`
	// Enables VPC connectivity for standard apps.
	// Structure is documented below.
	VpcAccessConnector *StandardAppVersionVpcAccessConnector `pulumi:"vpcAccessConnector"`
}

// The set of arguments for constructing a StandardAppVersion resource.
type StandardAppVersionArgs struct {
	// Automatic scaling is based on request rate, response latencies, and other application metrics.
	// Structure is documented below.
	AutomaticScaling StandardAppVersionAutomaticScalingPtrInput
	// Basic scaling creates instances when your application receives requests. Each instance will be shut down when the application becomes idle. Basic scaling is ideal for work that is intermittent or driven by user activity.
	// Structure is documented below.
	BasicScaling StandardAppVersionBasicScalingPtrInput
	// If set to `true`, the service will be deleted if it is the last version.
	DeleteServiceOnDestroy pulumi.BoolPtrInput
	// Code and application artifacts that make up this version.
	// Structure is documented below.
	Deployment StandardAppVersionDeploymentInput
	// The entrypoint for the application.
	// Structure is documented below.
	Entrypoint StandardAppVersionEntrypointPtrInput
	// Environment variables available to the application.
	EnvVariables pulumi.StringMapInput
	// An ordered list of URL-matching patterns that should be applied to incoming requests.
	// The first matching URL handles the request and other request handlers are not attempted.
	// Structure is documented below.
	Handlers StandardAppVersionHandlerArrayInput
	// A list of the types of messages that this application is able to receive.
	// Each value may be one of `INBOUND_SERVICE_MAIL`, `INBOUND_SERVICE_MAIL_BOUNCE`, `INBOUND_SERVICE_XMPP_ERROR`, `INBOUND_SERVICE_XMPP_MESSAGE`, `INBOUND_SERVICE_XMPP_SUBSCRIBE`, `INBOUND_SERVICE_XMPP_PRESENCE`, `INBOUND_SERVICE_CHANNEL_PRESENCE`, and `INBOUND_SERVICE_WARMUP`.
	InboundServices pulumi.StringArrayInput
	// Instance class that is used to run this version. Valid values are
	// AutomaticScaling: F1, F2, F4, F4_1G
	// BasicScaling or ManualScaling: B1, B2, B4, B4_1G, B8
	// Defaults to F1 for AutomaticScaling and B2 for ManualScaling and BasicScaling. If no scaling is specified, AutomaticScaling is chosen.
	InstanceClass pulumi.StringPtrInput
	// Configuration for third-party Python runtime libraries that are required by the application.
	// Structure is documented below.
	Libraries StandardAppVersionLibraryArrayInput
	// A service with manual scaling runs continuously, allowing you to perform complex initialization and rely on the state of its memory over time.
	// Structure is documented below.
	ManualScaling StandardAppVersionManualScalingPtrInput
	// If set to `true`, the application version will not be deleted.
	NoopOnDestroy pulumi.BoolPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// Desired runtime. Example python27.
	Runtime pulumi.StringInput
	// The version of the API in the given runtime environment.
	// Please see the app.yaml reference for valid values at https://cloud.google.com/appengine/docs/standard//config/appref
	RuntimeApiVersion pulumi.StringPtrInput
	// AppEngine service resource
	Service pulumi.StringInput
	// Whether multiple requests can be dispatched to this version at once.
	Threadsafe pulumi.BoolPtrInput
	// Relative name of the version within the service. For example, `v1`. Version names can contain only lowercase letters, numbers, or hyphens. Reserved names,"default", "latest", and any name with the prefix "ah-".
	VersionId pulumi.StringPtrInput
	// Enables VPC connectivity for standard apps.
	// Structure is documented below.
	VpcAccessConnector StandardAppVersionVpcAccessConnectorPtrInput
}

func (StandardAppVersionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*standardAppVersionArgs)(nil)).Elem()
}

type StandardAppVersionInput interface {
	pulumi.Input

	ToStandardAppVersionOutput() StandardAppVersionOutput
	ToStandardAppVersionOutputWithContext(ctx context.Context) StandardAppVersionOutput
}

func (StandardAppVersion) ElementType() reflect.Type {
	return reflect.TypeOf((*StandardAppVersion)(nil)).Elem()
}

func (i StandardAppVersion) ToStandardAppVersionOutput() StandardAppVersionOutput {
	return i.ToStandardAppVersionOutputWithContext(context.Background())
}

func (i StandardAppVersion) ToStandardAppVersionOutputWithContext(ctx context.Context) StandardAppVersionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StandardAppVersionOutput)
}

type StandardAppVersionOutput struct {
	*pulumi.OutputState
}

func (StandardAppVersionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StandardAppVersionOutput)(nil)).Elem()
}

func (o StandardAppVersionOutput) ToStandardAppVersionOutput() StandardAppVersionOutput {
	return o
}

func (o StandardAppVersionOutput) ToStandardAppVersionOutputWithContext(ctx context.Context) StandardAppVersionOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(StandardAppVersionOutput{})
}
