// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cloudrun

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Service acts as a top-level container that manages a set of Routes and
// Configurations which implement a network service. Service exists to provide a
// singular abstraction which can be access controlled, reasoned about, and
// which encapsulates software lifecycle decisions such as rollout policy and
// team resource ownership. Service acts only as an orchestrator of the
// underlying Routes and Configurations (much as a kubernetes Deployment
// orchestrates ReplicaSets).
//
// The Service's controller will track the statuses of its owned Configuration
// and Route, reflecting their statuses and conditions as its own.
//
// See also:
// https://github.com/knative/serving/blob/master/docs/spec/overview.md#service
//
// To get more information about Service, see:
//
// * [API documentation](https://cloud.google.com/run/docs/reference/rest/v1/projects.locations.services)
// * How-to Guides
//     * [Official Documentation](https://cloud.google.com/run/docs/)
//
// ## Example Usage
//
// ## Import
//
// Service can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import gcp:cloudrun/service:Service default locations/{{location}}/namespaces/{{project}}/services/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:cloudrun/service:Service default {{location}}/{{project}}/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:cloudrun/service:Service default {{location}}/{{name}}
// ```
type Service struct {
	pulumi.CustomResourceState

	// If set to `true`, the revision name (template.metadata.name) will be omitted and
	// autogenerated by Cloud Run. This cannot be set to `true` while `template.metadata.name`
	// is also set.
	// (For legacy support, if `template.metadata.name` is unset in state while
	// this field is set to false, the revision name will still autogenerate.)
	AutogenerateRevisionName pulumi.BoolPtrOutput `pulumi:"autogenerateRevisionName"`
	// The location of the cloud run instance. eg us-central1
	Location pulumi.StringOutput `pulumi:"location"`
	// Metadata associated with this Service, including name, namespace, labels,
	// and annotations.
	// Structure is documented below.
	Metadata ServiceMetadataOutput `pulumi:"metadata"`
	// Name of the port.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The current status of the Service.
	Statuses ServiceStatusArrayOutput `pulumi:"statuses"`
	// template holds the latest specification for the Revision to
	// be stamped out. The template references the container image, and may also
	// include labels and annotations that should be attached to the Revision.
	// To correlate a Revision, and/or to force a Revision to be created when the
	// spec doesn't otherwise change, a nonce label may be provided in the
	// template metadata. For more details, see:
	// https://github.com/knative/serving/blob/master/docs/client-conventions.md#associate-modifications-with-revisions
	// Cloud Run does not currently support referencing a build that is
	// responsible for materializing the container image from source.
	// Structure is documented below.
	Template ServiceTemplatePtrOutput `pulumi:"template"`
	// Traffic specifies how to distribute traffic over a collection of Knative Revisions
	// and Configurations
	// Structure is documented below.
	Traffics ServiceTrafficArrayOutput `pulumi:"traffics"`
}

// NewService registers a new resource with the given unique name, arguments, and options.
func NewService(ctx *pulumi.Context,
	name string, args *ServiceArgs, opts ...pulumi.ResourceOption) (*Service, error) {
	if args == nil || args.Location == nil {
		return nil, errors.New("missing required argument 'Location'")
	}
	if args == nil {
		args = &ServiceArgs{}
	}
	var resource Service
	err := ctx.RegisterResource("gcp:cloudrun/service:Service", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetService gets an existing Service resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetService(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServiceState, opts ...pulumi.ResourceOption) (*Service, error) {
	var resource Service
	err := ctx.ReadResource("gcp:cloudrun/service:Service", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Service resources.
type serviceState struct {
	// If set to `true`, the revision name (template.metadata.name) will be omitted and
	// autogenerated by Cloud Run. This cannot be set to `true` while `template.metadata.name`
	// is also set.
	// (For legacy support, if `template.metadata.name` is unset in state while
	// this field is set to false, the revision name will still autogenerate.)
	AutogenerateRevisionName *bool `pulumi:"autogenerateRevisionName"`
	// The location of the cloud run instance. eg us-central1
	Location *string `pulumi:"location"`
	// Metadata associated with this Service, including name, namespace, labels,
	// and annotations.
	// Structure is documented below.
	Metadata *ServiceMetadata `pulumi:"metadata"`
	// Name of the port.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The current status of the Service.
	Statuses []ServiceStatus `pulumi:"statuses"`
	// template holds the latest specification for the Revision to
	// be stamped out. The template references the container image, and may also
	// include labels and annotations that should be attached to the Revision.
	// To correlate a Revision, and/or to force a Revision to be created when the
	// spec doesn't otherwise change, a nonce label may be provided in the
	// template metadata. For more details, see:
	// https://github.com/knative/serving/blob/master/docs/client-conventions.md#associate-modifications-with-revisions
	// Cloud Run does not currently support referencing a build that is
	// responsible for materializing the container image from source.
	// Structure is documented below.
	Template *ServiceTemplate `pulumi:"template"`
	// Traffic specifies how to distribute traffic over a collection of Knative Revisions
	// and Configurations
	// Structure is documented below.
	Traffics []ServiceTraffic `pulumi:"traffics"`
}

type ServiceState struct {
	// If set to `true`, the revision name (template.metadata.name) will be omitted and
	// autogenerated by Cloud Run. This cannot be set to `true` while `template.metadata.name`
	// is also set.
	// (For legacy support, if `template.metadata.name` is unset in state while
	// this field is set to false, the revision name will still autogenerate.)
	AutogenerateRevisionName pulumi.BoolPtrInput
	// The location of the cloud run instance. eg us-central1
	Location pulumi.StringPtrInput
	// Metadata associated with this Service, including name, namespace, labels,
	// and annotations.
	// Structure is documented below.
	Metadata ServiceMetadataPtrInput
	// Name of the port.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The current status of the Service.
	Statuses ServiceStatusArrayInput
	// template holds the latest specification for the Revision to
	// be stamped out. The template references the container image, and may also
	// include labels and annotations that should be attached to the Revision.
	// To correlate a Revision, and/or to force a Revision to be created when the
	// spec doesn't otherwise change, a nonce label may be provided in the
	// template metadata. For more details, see:
	// https://github.com/knative/serving/blob/master/docs/client-conventions.md#associate-modifications-with-revisions
	// Cloud Run does not currently support referencing a build that is
	// responsible for materializing the container image from source.
	// Structure is documented below.
	Template ServiceTemplatePtrInput
	// Traffic specifies how to distribute traffic over a collection of Knative Revisions
	// and Configurations
	// Structure is documented below.
	Traffics ServiceTrafficArrayInput
}

func (ServiceState) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceState)(nil)).Elem()
}

type serviceArgs struct {
	// If set to `true`, the revision name (template.metadata.name) will be omitted and
	// autogenerated by Cloud Run. This cannot be set to `true` while `template.metadata.name`
	// is also set.
	// (For legacy support, if `template.metadata.name` is unset in state while
	// this field is set to false, the revision name will still autogenerate.)
	AutogenerateRevisionName *bool `pulumi:"autogenerateRevisionName"`
	// The location of the cloud run instance. eg us-central1
	Location string `pulumi:"location"`
	// Metadata associated with this Service, including name, namespace, labels,
	// and annotations.
	// Structure is documented below.
	Metadata *ServiceMetadata `pulumi:"metadata"`
	// Name of the port.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// template holds the latest specification for the Revision to
	// be stamped out. The template references the container image, and may also
	// include labels and annotations that should be attached to the Revision.
	// To correlate a Revision, and/or to force a Revision to be created when the
	// spec doesn't otherwise change, a nonce label may be provided in the
	// template metadata. For more details, see:
	// https://github.com/knative/serving/blob/master/docs/client-conventions.md#associate-modifications-with-revisions
	// Cloud Run does not currently support referencing a build that is
	// responsible for materializing the container image from source.
	// Structure is documented below.
	Template *ServiceTemplate `pulumi:"template"`
	// Traffic specifies how to distribute traffic over a collection of Knative Revisions
	// and Configurations
	// Structure is documented below.
	Traffics []ServiceTraffic `pulumi:"traffics"`
}

// The set of arguments for constructing a Service resource.
type ServiceArgs struct {
	// If set to `true`, the revision name (template.metadata.name) will be omitted and
	// autogenerated by Cloud Run. This cannot be set to `true` while `template.metadata.name`
	// is also set.
	// (For legacy support, if `template.metadata.name` is unset in state while
	// this field is set to false, the revision name will still autogenerate.)
	AutogenerateRevisionName pulumi.BoolPtrInput
	// The location of the cloud run instance. eg us-central1
	Location pulumi.StringInput
	// Metadata associated with this Service, including name, namespace, labels,
	// and annotations.
	// Structure is documented below.
	Metadata ServiceMetadataPtrInput
	// Name of the port.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// template holds the latest specification for the Revision to
	// be stamped out. The template references the container image, and may also
	// include labels and annotations that should be attached to the Revision.
	// To correlate a Revision, and/or to force a Revision to be created when the
	// spec doesn't otherwise change, a nonce label may be provided in the
	// template metadata. For more details, see:
	// https://github.com/knative/serving/blob/master/docs/client-conventions.md#associate-modifications-with-revisions
	// Cloud Run does not currently support referencing a build that is
	// responsible for materializing the container image from source.
	// Structure is documented below.
	Template ServiceTemplatePtrInput
	// Traffic specifies how to distribute traffic over a collection of Knative Revisions
	// and Configurations
	// Structure is documented below.
	Traffics ServiceTrafficArrayInput
}

func (ServiceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceArgs)(nil)).Elem()
}
