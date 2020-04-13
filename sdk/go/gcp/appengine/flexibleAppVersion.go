// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package appengine

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Flexible App Version resource to create a new version of flexible GAE Application. Based on Google Compute Engine,
// the App Engine flexible environment automatically scales your app up and down while also balancing the load.
// Learn about the differences between the standard environment and the flexible environment
// at https://cloud.google.com/appengine/docs/the-appengine-environments.
//
// > **Note:** The App Engine flexible environment service account uses the member ID `service-[YOUR_PROJECT_NUMBER]@gae-api-prod.google.com.iam.gserviceaccount.com`
// It should have the App Engine Flexible Environment Service Agent role, which will be applied when the `appengineflex.googleapis.com` service is enabled.
//
//
// To get more information about FlexibleAppVersion, see:
//
// * [API documentation](https://cloud.google.com/appengine/docs/admin-api/reference/rest/v1/apps.services.versions)
// * How-to Guides
//     * [Official Documentation](https://cloud.google.com/appengine/docs/flexible)
type FlexibleAppVersion struct {
	pulumi.CustomResourceState

	// -
	// (Optional)
	// Serving configuration for Google Cloud Endpoints.  Structure is documented below.
	ApiConfig FlexibleAppVersionApiConfigPtrOutput `pulumi:"apiConfig"`
	// -
	// (Optional)
	// Automatic scaling is based on request rate, response latencies, and other application metrics.  Structure is documented below.
	AutomaticScaling FlexibleAppVersionAutomaticScalingPtrOutput `pulumi:"automaticScaling"`
	// -
	// (Optional)
	// Metadata settings that are supplied to this version to enable beta runtime features.
	BetaSettings pulumi.StringMapOutput `pulumi:"betaSettings"`
	// -
	// (Optional)
	// Duration that static files should be cached by web proxies and browsers.
	// Only applicable if the corresponding StaticFilesHandler does not specify its own expiration time.
	DefaultExpiration pulumi.StringPtrOutput `pulumi:"defaultExpiration"`
	// If set to `true`, the service will be deleted if it is the last version.
	DeleteServiceOnDestroy pulumi.BoolPtrOutput `pulumi:"deleteServiceOnDestroy"`
	// -
	// (Optional)
	// Code and application artifacts that make up this version.  Structure is documented below.
	Deployment FlexibleAppVersionDeploymentPtrOutput `pulumi:"deployment"`
	// -
	// (Optional)
	// Code and application artifacts that make up this version.  Structure is documented below.
	EndpointsApiService FlexibleAppVersionEndpointsApiServicePtrOutput `pulumi:"endpointsApiService"`
	// -
	// (Optional)
	// The entrypoint for the application.  Structure is documented below.
	Entrypoint FlexibleAppVersionEntrypointPtrOutput `pulumi:"entrypoint"`
	// Environment variables available to the application. As these are not returned in the API request, Terraform will not
	// detect any changes made outside of the Terraform config.
	EnvVariables pulumi.StringMapOutput `pulumi:"envVariables"`
	// -
	// (Optional)
	// Before an application can receive email or XMPP messages, the application must be configured to enable the service.
	InboundServices pulumi.StringArrayOutput `pulumi:"inboundServices"`
	// -
	// (Optional)
	// Instance class that is used to run this version. Valid values are
	// AutomaticScaling: F1, F2, F4, F4_1G
	// ManualScaling: B1, B2, B4, B8, B4_1G
	// Defaults to F1 for AutomaticScaling and B1 for ManualScaling.
	InstanceClass pulumi.StringPtrOutput `pulumi:"instanceClass"`
	// -
	// (Required)
	// Health checking configuration for VM instances. Unhealthy instances are killed and replaced with new instances.  Structure is documented below.
	LivenessCheck FlexibleAppVersionLivenessCheckOutput `pulumi:"livenessCheck"`
	// -
	// (Optional)
	// A service with manual scaling runs continuously, allowing you to perform complex initialization and rely on the state of its memory over time.  Structure is documented below.
	ManualScaling FlexibleAppVersionManualScalingPtrOutput `pulumi:"manualScaling"`
	// -
	// (Required)
	// Full Serverless VPC Access Connector name e.g. /projects/my-project/locations/us-central1/connectors/c1.
	Name pulumi.StringOutput `pulumi:"name"`
	// -
	// (Optional)
	// Extra network settings  Structure is documented below.
	Network FlexibleAppVersionNetworkPtrOutput `pulumi:"network"`
	// -
	// (Optional)
	// Files that match this pattern will not be built into this version. Only applicable for Go runtimes.
	NobuildFilesRegex pulumi.StringPtrOutput `pulumi:"nobuildFilesRegex"`
	// If set to `true`, the application version will not be deleted.
	NoopOnDestroy pulumi.BoolPtrOutput `pulumi:"noopOnDestroy"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// -
	// (Required)
	// Configures readiness health checking for instances. Unhealthy instances are not put into the backend traffic rotation.  Structure is documented below.
	ReadinessCheck FlexibleAppVersionReadinessCheckOutput `pulumi:"readinessCheck"`
	// -
	// (Optional)
	// Machine resources for a version.  Structure is documented below.
	Resources FlexibleAppVersionResourcesPtrOutput `pulumi:"resources"`
	// -
	// (Required)
	// Desired runtime. Example python27.
	Runtime pulumi.StringOutput `pulumi:"runtime"`
	// -
	// (Optional)
	// The version of the API in the given runtime environment.
	// Please see the app.yaml reference for valid values at https://cloud.google.com/appengine/docs/standard//config/appref
	RuntimeApiVersion pulumi.StringOutput `pulumi:"runtimeApiVersion"`
	// -
	// (Optional)
	// The channel of the runtime to use. Only available for some runtimes.
	RuntimeChannel pulumi.StringPtrOutput `pulumi:"runtimeChannel"`
	// -
	// (Optional)
	// The path or name of the app's main executable.
	RuntimeMainExecutablePath pulumi.StringPtrOutput `pulumi:"runtimeMainExecutablePath"`
	// -
	// (Optional)
	// AppEngine service resource
	Service pulumi.StringPtrOutput `pulumi:"service"`
	// -
	// (Optional)
	// Current serving status of this version. Only the versions with a SERVING status create instances and can be billed.
	// Defaults to SERVING.
	ServingStatus pulumi.StringPtrOutput `pulumi:"servingStatus"`
	// -
	// (Optional)
	// Relative name of the version within the service. For example, `v1`. Version names can contain only lowercase letters, numbers, or hyphens.
	// Reserved names,"default", "latest", and any name with the prefix "ah-".
	VersionId pulumi.StringPtrOutput `pulumi:"versionId"`
	// -
	// (Optional)
	// Enables VPC connectivity for standard apps.  Structure is documented below.
	VpcAccessConnector FlexibleAppVersionVpcAccessConnectorPtrOutput `pulumi:"vpcAccessConnector"`
}

// NewFlexibleAppVersion registers a new resource with the given unique name, arguments, and options.
func NewFlexibleAppVersion(ctx *pulumi.Context,
	name string, args *FlexibleAppVersionArgs, opts ...pulumi.ResourceOption) (*FlexibleAppVersion, error) {
	if args == nil || args.LivenessCheck == nil {
		return nil, errors.New("missing required argument 'LivenessCheck'")
	}
	if args == nil || args.ReadinessCheck == nil {
		return nil, errors.New("missing required argument 'ReadinessCheck'")
	}
	if args == nil || args.Runtime == nil {
		return nil, errors.New("missing required argument 'Runtime'")
	}
	if args == nil {
		args = &FlexibleAppVersionArgs{}
	}
	var resource FlexibleAppVersion
	err := ctx.RegisterResource("gcp:appengine/flexibleAppVersion:FlexibleAppVersion", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFlexibleAppVersion gets an existing FlexibleAppVersion resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFlexibleAppVersion(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FlexibleAppVersionState, opts ...pulumi.ResourceOption) (*FlexibleAppVersion, error) {
	var resource FlexibleAppVersion
	err := ctx.ReadResource("gcp:appengine/flexibleAppVersion:FlexibleAppVersion", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FlexibleAppVersion resources.
type flexibleAppVersionState struct {
	// -
	// (Optional)
	// Serving configuration for Google Cloud Endpoints.  Structure is documented below.
	ApiConfig *FlexibleAppVersionApiConfig `pulumi:"apiConfig"`
	// -
	// (Optional)
	// Automatic scaling is based on request rate, response latencies, and other application metrics.  Structure is documented below.
	AutomaticScaling *FlexibleAppVersionAutomaticScaling `pulumi:"automaticScaling"`
	// -
	// (Optional)
	// Metadata settings that are supplied to this version to enable beta runtime features.
	BetaSettings map[string]string `pulumi:"betaSettings"`
	// -
	// (Optional)
	// Duration that static files should be cached by web proxies and browsers.
	// Only applicable if the corresponding StaticFilesHandler does not specify its own expiration time.
	DefaultExpiration *string `pulumi:"defaultExpiration"`
	// If set to `true`, the service will be deleted if it is the last version.
	DeleteServiceOnDestroy *bool `pulumi:"deleteServiceOnDestroy"`
	// -
	// (Optional)
	// Code and application artifacts that make up this version.  Structure is documented below.
	Deployment *FlexibleAppVersionDeployment `pulumi:"deployment"`
	// -
	// (Optional)
	// Code and application artifacts that make up this version.  Structure is documented below.
	EndpointsApiService *FlexibleAppVersionEndpointsApiService `pulumi:"endpointsApiService"`
	// -
	// (Optional)
	// The entrypoint for the application.  Structure is documented below.
	Entrypoint *FlexibleAppVersionEntrypoint `pulumi:"entrypoint"`
	// Environment variables available to the application. As these are not returned in the API request, Terraform will not
	// detect any changes made outside of the Terraform config.
	EnvVariables map[string]string `pulumi:"envVariables"`
	// -
	// (Optional)
	// Before an application can receive email or XMPP messages, the application must be configured to enable the service.
	InboundServices []string `pulumi:"inboundServices"`
	// -
	// (Optional)
	// Instance class that is used to run this version. Valid values are
	// AutomaticScaling: F1, F2, F4, F4_1G
	// ManualScaling: B1, B2, B4, B8, B4_1G
	// Defaults to F1 for AutomaticScaling and B1 for ManualScaling.
	InstanceClass *string `pulumi:"instanceClass"`
	// -
	// (Required)
	// Health checking configuration for VM instances. Unhealthy instances are killed and replaced with new instances.  Structure is documented below.
	LivenessCheck *FlexibleAppVersionLivenessCheck `pulumi:"livenessCheck"`
	// -
	// (Optional)
	// A service with manual scaling runs continuously, allowing you to perform complex initialization and rely on the state of its memory over time.  Structure is documented below.
	ManualScaling *FlexibleAppVersionManualScaling `pulumi:"manualScaling"`
	// -
	// (Required)
	// Full Serverless VPC Access Connector name e.g. /projects/my-project/locations/us-central1/connectors/c1.
	Name *string `pulumi:"name"`
	// -
	// (Optional)
	// Extra network settings  Structure is documented below.
	Network *FlexibleAppVersionNetwork `pulumi:"network"`
	// -
	// (Optional)
	// Files that match this pattern will not be built into this version. Only applicable for Go runtimes.
	NobuildFilesRegex *string `pulumi:"nobuildFilesRegex"`
	// If set to `true`, the application version will not be deleted.
	NoopOnDestroy *bool `pulumi:"noopOnDestroy"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// -
	// (Required)
	// Configures readiness health checking for instances. Unhealthy instances are not put into the backend traffic rotation.  Structure is documented below.
	ReadinessCheck *FlexibleAppVersionReadinessCheck `pulumi:"readinessCheck"`
	// -
	// (Optional)
	// Machine resources for a version.  Structure is documented below.
	Resources *FlexibleAppVersionResources `pulumi:"resources"`
	// -
	// (Required)
	// Desired runtime. Example python27.
	Runtime *string `pulumi:"runtime"`
	// -
	// (Optional)
	// The version of the API in the given runtime environment.
	// Please see the app.yaml reference for valid values at https://cloud.google.com/appengine/docs/standard//config/appref
	RuntimeApiVersion *string `pulumi:"runtimeApiVersion"`
	// -
	// (Optional)
	// The channel of the runtime to use. Only available for some runtimes.
	RuntimeChannel *string `pulumi:"runtimeChannel"`
	// -
	// (Optional)
	// The path or name of the app's main executable.
	RuntimeMainExecutablePath *string `pulumi:"runtimeMainExecutablePath"`
	// -
	// (Optional)
	// AppEngine service resource
	Service *string `pulumi:"service"`
	// -
	// (Optional)
	// Current serving status of this version. Only the versions with a SERVING status create instances and can be billed.
	// Defaults to SERVING.
	ServingStatus *string `pulumi:"servingStatus"`
	// -
	// (Optional)
	// Relative name of the version within the service. For example, `v1`. Version names can contain only lowercase letters, numbers, or hyphens.
	// Reserved names,"default", "latest", and any name with the prefix "ah-".
	VersionId *string `pulumi:"versionId"`
	// -
	// (Optional)
	// Enables VPC connectivity for standard apps.  Structure is documented below.
	VpcAccessConnector *FlexibleAppVersionVpcAccessConnector `pulumi:"vpcAccessConnector"`
}

type FlexibleAppVersionState struct {
	// -
	// (Optional)
	// Serving configuration for Google Cloud Endpoints.  Structure is documented below.
	ApiConfig FlexibleAppVersionApiConfigPtrInput
	// -
	// (Optional)
	// Automatic scaling is based on request rate, response latencies, and other application metrics.  Structure is documented below.
	AutomaticScaling FlexibleAppVersionAutomaticScalingPtrInput
	// -
	// (Optional)
	// Metadata settings that are supplied to this version to enable beta runtime features.
	BetaSettings pulumi.StringMapInput
	// -
	// (Optional)
	// Duration that static files should be cached by web proxies and browsers.
	// Only applicable if the corresponding StaticFilesHandler does not specify its own expiration time.
	DefaultExpiration pulumi.StringPtrInput
	// If set to `true`, the service will be deleted if it is the last version.
	DeleteServiceOnDestroy pulumi.BoolPtrInput
	// -
	// (Optional)
	// Code and application artifacts that make up this version.  Structure is documented below.
	Deployment FlexibleAppVersionDeploymentPtrInput
	// -
	// (Optional)
	// Code and application artifacts that make up this version.  Structure is documented below.
	EndpointsApiService FlexibleAppVersionEndpointsApiServicePtrInput
	// -
	// (Optional)
	// The entrypoint for the application.  Structure is documented below.
	Entrypoint FlexibleAppVersionEntrypointPtrInput
	// Environment variables available to the application. As these are not returned in the API request, Terraform will not
	// detect any changes made outside of the Terraform config.
	EnvVariables pulumi.StringMapInput
	// -
	// (Optional)
	// Before an application can receive email or XMPP messages, the application must be configured to enable the service.
	InboundServices pulumi.StringArrayInput
	// -
	// (Optional)
	// Instance class that is used to run this version. Valid values are
	// AutomaticScaling: F1, F2, F4, F4_1G
	// ManualScaling: B1, B2, B4, B8, B4_1G
	// Defaults to F1 for AutomaticScaling and B1 for ManualScaling.
	InstanceClass pulumi.StringPtrInput
	// -
	// (Required)
	// Health checking configuration for VM instances. Unhealthy instances are killed and replaced with new instances.  Structure is documented below.
	LivenessCheck FlexibleAppVersionLivenessCheckPtrInput
	// -
	// (Optional)
	// A service with manual scaling runs continuously, allowing you to perform complex initialization and rely on the state of its memory over time.  Structure is documented below.
	ManualScaling FlexibleAppVersionManualScalingPtrInput
	// -
	// (Required)
	// Full Serverless VPC Access Connector name e.g. /projects/my-project/locations/us-central1/connectors/c1.
	Name pulumi.StringPtrInput
	// -
	// (Optional)
	// Extra network settings  Structure is documented below.
	Network FlexibleAppVersionNetworkPtrInput
	// -
	// (Optional)
	// Files that match this pattern will not be built into this version. Only applicable for Go runtimes.
	NobuildFilesRegex pulumi.StringPtrInput
	// If set to `true`, the application version will not be deleted.
	NoopOnDestroy pulumi.BoolPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// -
	// (Required)
	// Configures readiness health checking for instances. Unhealthy instances are not put into the backend traffic rotation.  Structure is documented below.
	ReadinessCheck FlexibleAppVersionReadinessCheckPtrInput
	// -
	// (Optional)
	// Machine resources for a version.  Structure is documented below.
	Resources FlexibleAppVersionResourcesPtrInput
	// -
	// (Required)
	// Desired runtime. Example python27.
	Runtime pulumi.StringPtrInput
	// -
	// (Optional)
	// The version of the API in the given runtime environment.
	// Please see the app.yaml reference for valid values at https://cloud.google.com/appengine/docs/standard//config/appref
	RuntimeApiVersion pulumi.StringPtrInput
	// -
	// (Optional)
	// The channel of the runtime to use. Only available for some runtimes.
	RuntimeChannel pulumi.StringPtrInput
	// -
	// (Optional)
	// The path or name of the app's main executable.
	RuntimeMainExecutablePath pulumi.StringPtrInput
	// -
	// (Optional)
	// AppEngine service resource
	Service pulumi.StringPtrInput
	// -
	// (Optional)
	// Current serving status of this version. Only the versions with a SERVING status create instances and can be billed.
	// Defaults to SERVING.
	ServingStatus pulumi.StringPtrInput
	// -
	// (Optional)
	// Relative name of the version within the service. For example, `v1`. Version names can contain only lowercase letters, numbers, or hyphens.
	// Reserved names,"default", "latest", and any name with the prefix "ah-".
	VersionId pulumi.StringPtrInput
	// -
	// (Optional)
	// Enables VPC connectivity for standard apps.  Structure is documented below.
	VpcAccessConnector FlexibleAppVersionVpcAccessConnectorPtrInput
}

func (FlexibleAppVersionState) ElementType() reflect.Type {
	return reflect.TypeOf((*flexibleAppVersionState)(nil)).Elem()
}

type flexibleAppVersionArgs struct {
	// -
	// (Optional)
	// Serving configuration for Google Cloud Endpoints.  Structure is documented below.
	ApiConfig *FlexibleAppVersionApiConfig `pulumi:"apiConfig"`
	// -
	// (Optional)
	// Automatic scaling is based on request rate, response latencies, and other application metrics.  Structure is documented below.
	AutomaticScaling *FlexibleAppVersionAutomaticScaling `pulumi:"automaticScaling"`
	// -
	// (Optional)
	// Metadata settings that are supplied to this version to enable beta runtime features.
	BetaSettings map[string]string `pulumi:"betaSettings"`
	// -
	// (Optional)
	// Duration that static files should be cached by web proxies and browsers.
	// Only applicable if the corresponding StaticFilesHandler does not specify its own expiration time.
	DefaultExpiration *string `pulumi:"defaultExpiration"`
	// If set to `true`, the service will be deleted if it is the last version.
	DeleteServiceOnDestroy *bool `pulumi:"deleteServiceOnDestroy"`
	// -
	// (Optional)
	// Code and application artifacts that make up this version.  Structure is documented below.
	Deployment *FlexibleAppVersionDeployment `pulumi:"deployment"`
	// -
	// (Optional)
	// Code and application artifacts that make up this version.  Structure is documented below.
	EndpointsApiService *FlexibleAppVersionEndpointsApiService `pulumi:"endpointsApiService"`
	// -
	// (Optional)
	// The entrypoint for the application.  Structure is documented below.
	Entrypoint *FlexibleAppVersionEntrypoint `pulumi:"entrypoint"`
	// Environment variables available to the application. As these are not returned in the API request, Terraform will not
	// detect any changes made outside of the Terraform config.
	EnvVariables map[string]string `pulumi:"envVariables"`
	// -
	// (Optional)
	// Before an application can receive email or XMPP messages, the application must be configured to enable the service.
	InboundServices []string `pulumi:"inboundServices"`
	// -
	// (Optional)
	// Instance class that is used to run this version. Valid values are
	// AutomaticScaling: F1, F2, F4, F4_1G
	// ManualScaling: B1, B2, B4, B8, B4_1G
	// Defaults to F1 for AutomaticScaling and B1 for ManualScaling.
	InstanceClass *string `pulumi:"instanceClass"`
	// -
	// (Required)
	// Health checking configuration for VM instances. Unhealthy instances are killed and replaced with new instances.  Structure is documented below.
	LivenessCheck FlexibleAppVersionLivenessCheck `pulumi:"livenessCheck"`
	// -
	// (Optional)
	// A service with manual scaling runs continuously, allowing you to perform complex initialization and rely on the state of its memory over time.  Structure is documented below.
	ManualScaling *FlexibleAppVersionManualScaling `pulumi:"manualScaling"`
	// -
	// (Optional)
	// Extra network settings  Structure is documented below.
	Network *FlexibleAppVersionNetwork `pulumi:"network"`
	// -
	// (Optional)
	// Files that match this pattern will not be built into this version. Only applicable for Go runtimes.
	NobuildFilesRegex *string `pulumi:"nobuildFilesRegex"`
	// If set to `true`, the application version will not be deleted.
	NoopOnDestroy *bool `pulumi:"noopOnDestroy"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// -
	// (Required)
	// Configures readiness health checking for instances. Unhealthy instances are not put into the backend traffic rotation.  Structure is documented below.
	ReadinessCheck FlexibleAppVersionReadinessCheck `pulumi:"readinessCheck"`
	// -
	// (Optional)
	// Machine resources for a version.  Structure is documented below.
	Resources *FlexibleAppVersionResources `pulumi:"resources"`
	// -
	// (Required)
	// Desired runtime. Example python27.
	Runtime string `pulumi:"runtime"`
	// -
	// (Optional)
	// The version of the API in the given runtime environment.
	// Please see the app.yaml reference for valid values at https://cloud.google.com/appengine/docs/standard//config/appref
	RuntimeApiVersion *string `pulumi:"runtimeApiVersion"`
	// -
	// (Optional)
	// The channel of the runtime to use. Only available for some runtimes.
	RuntimeChannel *string `pulumi:"runtimeChannel"`
	// -
	// (Optional)
	// The path or name of the app's main executable.
	RuntimeMainExecutablePath *string `pulumi:"runtimeMainExecutablePath"`
	// -
	// (Optional)
	// AppEngine service resource
	Service *string `pulumi:"service"`
	// -
	// (Optional)
	// Current serving status of this version. Only the versions with a SERVING status create instances and can be billed.
	// Defaults to SERVING.
	ServingStatus *string `pulumi:"servingStatus"`
	// -
	// (Optional)
	// Relative name of the version within the service. For example, `v1`. Version names can contain only lowercase letters, numbers, or hyphens.
	// Reserved names,"default", "latest", and any name with the prefix "ah-".
	VersionId *string `pulumi:"versionId"`
	// -
	// (Optional)
	// Enables VPC connectivity for standard apps.  Structure is documented below.
	VpcAccessConnector *FlexibleAppVersionVpcAccessConnector `pulumi:"vpcAccessConnector"`
}

// The set of arguments for constructing a FlexibleAppVersion resource.
type FlexibleAppVersionArgs struct {
	// -
	// (Optional)
	// Serving configuration for Google Cloud Endpoints.  Structure is documented below.
	ApiConfig FlexibleAppVersionApiConfigPtrInput
	// -
	// (Optional)
	// Automatic scaling is based on request rate, response latencies, and other application metrics.  Structure is documented below.
	AutomaticScaling FlexibleAppVersionAutomaticScalingPtrInput
	// -
	// (Optional)
	// Metadata settings that are supplied to this version to enable beta runtime features.
	BetaSettings pulumi.StringMapInput
	// -
	// (Optional)
	// Duration that static files should be cached by web proxies and browsers.
	// Only applicable if the corresponding StaticFilesHandler does not specify its own expiration time.
	DefaultExpiration pulumi.StringPtrInput
	// If set to `true`, the service will be deleted if it is the last version.
	DeleteServiceOnDestroy pulumi.BoolPtrInput
	// -
	// (Optional)
	// Code and application artifacts that make up this version.  Structure is documented below.
	Deployment FlexibleAppVersionDeploymentPtrInput
	// -
	// (Optional)
	// Code and application artifacts that make up this version.  Structure is documented below.
	EndpointsApiService FlexibleAppVersionEndpointsApiServicePtrInput
	// -
	// (Optional)
	// The entrypoint for the application.  Structure is documented below.
	Entrypoint FlexibleAppVersionEntrypointPtrInput
	// Environment variables available to the application. As these are not returned in the API request, Terraform will not
	// detect any changes made outside of the Terraform config.
	EnvVariables pulumi.StringMapInput
	// -
	// (Optional)
	// Before an application can receive email or XMPP messages, the application must be configured to enable the service.
	InboundServices pulumi.StringArrayInput
	// -
	// (Optional)
	// Instance class that is used to run this version. Valid values are
	// AutomaticScaling: F1, F2, F4, F4_1G
	// ManualScaling: B1, B2, B4, B8, B4_1G
	// Defaults to F1 for AutomaticScaling and B1 for ManualScaling.
	InstanceClass pulumi.StringPtrInput
	// -
	// (Required)
	// Health checking configuration for VM instances. Unhealthy instances are killed and replaced with new instances.  Structure is documented below.
	LivenessCheck FlexibleAppVersionLivenessCheckInput
	// -
	// (Optional)
	// A service with manual scaling runs continuously, allowing you to perform complex initialization and rely on the state of its memory over time.  Structure is documented below.
	ManualScaling FlexibleAppVersionManualScalingPtrInput
	// -
	// (Optional)
	// Extra network settings  Structure is documented below.
	Network FlexibleAppVersionNetworkPtrInput
	// -
	// (Optional)
	// Files that match this pattern will not be built into this version. Only applicable for Go runtimes.
	NobuildFilesRegex pulumi.StringPtrInput
	// If set to `true`, the application version will not be deleted.
	NoopOnDestroy pulumi.BoolPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// -
	// (Required)
	// Configures readiness health checking for instances. Unhealthy instances are not put into the backend traffic rotation.  Structure is documented below.
	ReadinessCheck FlexibleAppVersionReadinessCheckInput
	// -
	// (Optional)
	// Machine resources for a version.  Structure is documented below.
	Resources FlexibleAppVersionResourcesPtrInput
	// -
	// (Required)
	// Desired runtime. Example python27.
	Runtime pulumi.StringInput
	// -
	// (Optional)
	// The version of the API in the given runtime environment.
	// Please see the app.yaml reference for valid values at https://cloud.google.com/appengine/docs/standard//config/appref
	RuntimeApiVersion pulumi.StringPtrInput
	// -
	// (Optional)
	// The channel of the runtime to use. Only available for some runtimes.
	RuntimeChannel pulumi.StringPtrInput
	// -
	// (Optional)
	// The path or name of the app's main executable.
	RuntimeMainExecutablePath pulumi.StringPtrInput
	// -
	// (Optional)
	// AppEngine service resource
	Service pulumi.StringPtrInput
	// -
	// (Optional)
	// Current serving status of this version. Only the versions with a SERVING status create instances and can be billed.
	// Defaults to SERVING.
	ServingStatus pulumi.StringPtrInput
	// -
	// (Optional)
	// Relative name of the version within the service. For example, `v1`. Version names can contain only lowercase letters, numbers, or hyphens.
	// Reserved names,"default", "latest", and any name with the prefix "ah-".
	VersionId pulumi.StringPtrInput
	// -
	// (Optional)
	// Enables VPC connectivity for standard apps.  Structure is documented below.
	VpcAccessConnector FlexibleAppVersionVpcAccessConnectorPtrInput
}

func (FlexibleAppVersionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*flexibleAppVersionArgs)(nil)).Elem()
}
