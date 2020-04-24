// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package appengine

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Standard App Version resource to create a new version of standard GAE Application.
// Currently supporting Zip and File Containers.
// Currently does not support async operation checking.
//
//
// To get more information about StandardAppVersion, see:
//
// * [API documentation](https://cloud.google.com/appengine/docs/admin-api/reference/rest/v1/apps.services.versions)
// * How-to Guides
//     * [Official Documentation](https://cloud.google.com/appengine/docs/standard)
type StandardAppVersion struct {
	pulumi.CustomResourceState

	// If set to `true`, the service will be deleted if it is the last version.
	DeleteServiceOnDestroy pulumi.BoolPtrOutput `pulumi:"deleteServiceOnDestroy"`
	// -
	// (Optional)
	// Code and application artifacts that make up this version.  Structure is documented below.
	Deployment StandardAppVersionDeploymentPtrOutput `pulumi:"deployment"`
	// -
	// (Optional)
	// The entrypoint for the application.  Structure is documented below.
	Entrypoint StandardAppVersionEntrypointPtrOutput `pulumi:"entrypoint"`
	// -
	// (Optional)
	// Environment variables available to the application.
	EnvVariables pulumi.StringMapOutput `pulumi:"envVariables"`
	// -
	// (Optional)
	// An ordered list of URL-matching patterns that should be applied to incoming requests.
	// The first matching URL handles the request and other request handlers are not attempted.  Structure is documented below.
	Handlers StandardAppVersionHandlerArrayOutput `pulumi:"handlers"`
	// -
	// (Optional)
	// Instance class that is used to run this version. Valid values are
	// AutomaticScaling F1, F2, F4, F4_1G
	// (Only AutomaticScaling is supported at the moment)
	InstanceClass pulumi.StringPtrOutput `pulumi:"instanceClass"`
	// -
	// (Optional)
	// Configuration for third-party Python runtime libraries that are required by the application.  Structure is documented below.
	Libraries StandardAppVersionLibraryArrayOutput `pulumi:"libraries"`
	// The identifier for this object. Format specified above.
	Name pulumi.StringOutput `pulumi:"name"`
	// If set to `true`, the application version will not be deleted.
	NoopOnDestroy pulumi.BoolPtrOutput `pulumi:"noopOnDestroy"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// -
	// (Required)
	// Desired runtime. Example python27.
	Runtime pulumi.StringOutput `pulumi:"runtime"`
	// -
	// (Optional)
	// The version of the API in the given runtime environment.
	// Please see the app.yaml reference for valid values at https://cloud.google.com/appengine/docs/standard//config/appref
	RuntimeApiVersion pulumi.StringPtrOutput `pulumi:"runtimeApiVersion"`
	// -
	// (Optional)
	// AppEngine service resource
	Service pulumi.StringPtrOutput `pulumi:"service"`
	// -
	// (Optional)
	// Whether multiple requests can be dispatched to this version at once.
	Threadsafe pulumi.BoolPtrOutput `pulumi:"threadsafe"`
	// -
	// (Optional)
	// Relative name of the version within the service. For example, `v1`. Version names can contain only lowercase letters, numbers, or hyphens. Reserved names,"default", "latest", and any name with the prefix "ah-".
	VersionId pulumi.StringPtrOutput `pulumi:"versionId"`
}

// NewStandardAppVersion registers a new resource with the given unique name, arguments, and options.
func NewStandardAppVersion(ctx *pulumi.Context,
	name string, args *StandardAppVersionArgs, opts ...pulumi.ResourceOption) (*StandardAppVersion, error) {
	if args == nil || args.Runtime == nil {
		return nil, errors.New("missing required argument 'Runtime'")
	}
	if args == nil {
		args = &StandardAppVersionArgs{}
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
	// If set to `true`, the service will be deleted if it is the last version.
	DeleteServiceOnDestroy *bool `pulumi:"deleteServiceOnDestroy"`
	// -
	// (Optional)
	// Code and application artifacts that make up this version.  Structure is documented below.
	Deployment *StandardAppVersionDeployment `pulumi:"deployment"`
	// -
	// (Optional)
	// The entrypoint for the application.  Structure is documented below.
	Entrypoint *StandardAppVersionEntrypoint `pulumi:"entrypoint"`
	// -
	// (Optional)
	// Environment variables available to the application.
	EnvVariables map[string]string `pulumi:"envVariables"`
	// -
	// (Optional)
	// An ordered list of URL-matching patterns that should be applied to incoming requests.
	// The first matching URL handles the request and other request handlers are not attempted.  Structure is documented below.
	Handlers []StandardAppVersionHandler `pulumi:"handlers"`
	// -
	// (Optional)
	// Instance class that is used to run this version. Valid values are
	// AutomaticScaling F1, F2, F4, F4_1G
	// (Only AutomaticScaling is supported at the moment)
	InstanceClass *string `pulumi:"instanceClass"`
	// -
	// (Optional)
	// Configuration for third-party Python runtime libraries that are required by the application.  Structure is documented below.
	Libraries []StandardAppVersionLibrary `pulumi:"libraries"`
	// The identifier for this object. Format specified above.
	Name *string `pulumi:"name"`
	// If set to `true`, the application version will not be deleted.
	NoopOnDestroy *bool `pulumi:"noopOnDestroy"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
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
	// AppEngine service resource
	Service *string `pulumi:"service"`
	// -
	// (Optional)
	// Whether multiple requests can be dispatched to this version at once.
	Threadsafe *bool `pulumi:"threadsafe"`
	// -
	// (Optional)
	// Relative name of the version within the service. For example, `v1`. Version names can contain only lowercase letters, numbers, or hyphens. Reserved names,"default", "latest", and any name with the prefix "ah-".
	VersionId *string `pulumi:"versionId"`
}

type StandardAppVersionState struct {
	// If set to `true`, the service will be deleted if it is the last version.
	DeleteServiceOnDestroy pulumi.BoolPtrInput
	// -
	// (Optional)
	// Code and application artifacts that make up this version.  Structure is documented below.
	Deployment StandardAppVersionDeploymentPtrInput
	// -
	// (Optional)
	// The entrypoint for the application.  Structure is documented below.
	Entrypoint StandardAppVersionEntrypointPtrInput
	// -
	// (Optional)
	// Environment variables available to the application.
	EnvVariables pulumi.StringMapInput
	// -
	// (Optional)
	// An ordered list of URL-matching patterns that should be applied to incoming requests.
	// The first matching URL handles the request and other request handlers are not attempted.  Structure is documented below.
	Handlers StandardAppVersionHandlerArrayInput
	// -
	// (Optional)
	// Instance class that is used to run this version. Valid values are
	// AutomaticScaling F1, F2, F4, F4_1G
	// (Only AutomaticScaling is supported at the moment)
	InstanceClass pulumi.StringPtrInput
	// -
	// (Optional)
	// Configuration for third-party Python runtime libraries that are required by the application.  Structure is documented below.
	Libraries StandardAppVersionLibraryArrayInput
	// The identifier for this object. Format specified above.
	Name pulumi.StringPtrInput
	// If set to `true`, the application version will not be deleted.
	NoopOnDestroy pulumi.BoolPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
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
	// AppEngine service resource
	Service pulumi.StringPtrInput
	// -
	// (Optional)
	// Whether multiple requests can be dispatched to this version at once.
	Threadsafe pulumi.BoolPtrInput
	// -
	// (Optional)
	// Relative name of the version within the service. For example, `v1`. Version names can contain only lowercase letters, numbers, or hyphens. Reserved names,"default", "latest", and any name with the prefix "ah-".
	VersionId pulumi.StringPtrInput
}

func (StandardAppVersionState) ElementType() reflect.Type {
	return reflect.TypeOf((*standardAppVersionState)(nil)).Elem()
}

type standardAppVersionArgs struct {
	// If set to `true`, the service will be deleted if it is the last version.
	DeleteServiceOnDestroy *bool `pulumi:"deleteServiceOnDestroy"`
	// -
	// (Optional)
	// Code and application artifacts that make up this version.  Structure is documented below.
	Deployment *StandardAppVersionDeployment `pulumi:"deployment"`
	// -
	// (Optional)
	// The entrypoint for the application.  Structure is documented below.
	Entrypoint *StandardAppVersionEntrypoint `pulumi:"entrypoint"`
	// -
	// (Optional)
	// Environment variables available to the application.
	EnvVariables map[string]string `pulumi:"envVariables"`
	// -
	// (Optional)
	// An ordered list of URL-matching patterns that should be applied to incoming requests.
	// The first matching URL handles the request and other request handlers are not attempted.  Structure is documented below.
	Handlers []StandardAppVersionHandler `pulumi:"handlers"`
	// -
	// (Optional)
	// Instance class that is used to run this version. Valid values are
	// AutomaticScaling F1, F2, F4, F4_1G
	// (Only AutomaticScaling is supported at the moment)
	InstanceClass *string `pulumi:"instanceClass"`
	// -
	// (Optional)
	// Configuration for third-party Python runtime libraries that are required by the application.  Structure is documented below.
	Libraries []StandardAppVersionLibrary `pulumi:"libraries"`
	// If set to `true`, the application version will not be deleted.
	NoopOnDestroy *bool `pulumi:"noopOnDestroy"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
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
	// AppEngine service resource
	Service *string `pulumi:"service"`
	// -
	// (Optional)
	// Whether multiple requests can be dispatched to this version at once.
	Threadsafe *bool `pulumi:"threadsafe"`
	// -
	// (Optional)
	// Relative name of the version within the service. For example, `v1`. Version names can contain only lowercase letters, numbers, or hyphens. Reserved names,"default", "latest", and any name with the prefix "ah-".
	VersionId *string `pulumi:"versionId"`
}

// The set of arguments for constructing a StandardAppVersion resource.
type StandardAppVersionArgs struct {
	// If set to `true`, the service will be deleted if it is the last version.
	DeleteServiceOnDestroy pulumi.BoolPtrInput
	// -
	// (Optional)
	// Code and application artifacts that make up this version.  Structure is documented below.
	Deployment StandardAppVersionDeploymentPtrInput
	// -
	// (Optional)
	// The entrypoint for the application.  Structure is documented below.
	Entrypoint StandardAppVersionEntrypointPtrInput
	// -
	// (Optional)
	// Environment variables available to the application.
	EnvVariables pulumi.StringMapInput
	// -
	// (Optional)
	// An ordered list of URL-matching patterns that should be applied to incoming requests.
	// The first matching URL handles the request and other request handlers are not attempted.  Structure is documented below.
	Handlers StandardAppVersionHandlerArrayInput
	// -
	// (Optional)
	// Instance class that is used to run this version. Valid values are
	// AutomaticScaling F1, F2, F4, F4_1G
	// (Only AutomaticScaling is supported at the moment)
	InstanceClass pulumi.StringPtrInput
	// -
	// (Optional)
	// Configuration for third-party Python runtime libraries that are required by the application.  Structure is documented below.
	Libraries StandardAppVersionLibraryArrayInput
	// If set to `true`, the application version will not be deleted.
	NoopOnDestroy pulumi.BoolPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
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
	// AppEngine service resource
	Service pulumi.StringPtrInput
	// -
	// (Optional)
	// Whether multiple requests can be dispatched to this version at once.
	Threadsafe pulumi.BoolPtrInput
	// -
	// (Optional)
	// Relative name of the version within the service. For example, `v1`. Version names can contain only lowercase letters, numbers, or hyphens. Reserved names,"default", "latest", and any name with the prefix "ah-".
	VersionId pulumi.StringPtrInput
}

func (StandardAppVersionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*standardAppVersionArgs)(nil)).Elem()
}
