// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Represents an Autoscaler resource.
//
// Autoscalers allow you to automatically scale virtual machine instances in
// managed instance groups according to an autoscaling policy that you
// define.
//
//
// To get more information about Autoscaler, see:
//
// * [API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/autoscalers)
// * How-to Guides
//     * [Autoscaling Groups of Instances](https://cloud.google.com/compute/docs/autoscaler/)
type Autoscalar struct {
	pulumi.CustomResourceState

	// -
	// (Required)
	// The configuration parameters for the autoscaling algorithm. You can
	// define one or more of the policies for an autoscaler: cpuUtilization,
	// customMetricUtilizations, and loadBalancingUtilization.
	// If none of these are specified, the default will be to autoscale based
	// on cpuUtilization to 0.6 or 60%.  Structure is documented below.
	AutoscalingPolicy AutoscalarAutoscalingPolicyOutput `pulumi:"autoscalingPolicy"`
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringOutput `pulumi:"creationTimestamp"`
	// -
	// (Optional)
	// An optional description of this resource.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// -
	// (Required)
	// The identifier (type) of the Stackdriver Monitoring metric.
	// The metric cannot have negative values.
	// The metric must have a value type of INT64 or DOUBLE.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The URI of the created resource.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
	// -
	// (Required)
	// Fraction of backend capacity utilization (set in HTTP(s) load
	// balancing configuration) that autoscaler should maintain. Must
	// be a positive float value. If not defined, the default is 0.8.
	Target pulumi.StringOutput `pulumi:"target"`
	// -
	// (Optional)
	// URL of the zone where the instance group resides.
	Zone pulumi.StringOutput `pulumi:"zone"`
}

// NewAutoscalar registers a new resource with the given unique name, arguments, and options.
func NewAutoscalar(ctx *pulumi.Context,
	name string, args *AutoscalarArgs, opts ...pulumi.ResourceOption) (*Autoscalar, error) {
	if args == nil || args.AutoscalingPolicy == nil {
		return nil, errors.New("missing required argument 'AutoscalingPolicy'")
	}
	if args == nil || args.Target == nil {
		return nil, errors.New("missing required argument 'Target'")
	}
	if args == nil {
		args = &AutoscalarArgs{}
	}
	var resource Autoscalar
	err := ctx.RegisterResource("gcp:compute/autoscalar:Autoscalar", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAutoscalar gets an existing Autoscalar resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAutoscalar(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AutoscalarState, opts ...pulumi.ResourceOption) (*Autoscalar, error) {
	var resource Autoscalar
	err := ctx.ReadResource("gcp:compute/autoscalar:Autoscalar", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Autoscalar resources.
type autoscalarState struct {
	// -
	// (Required)
	// The configuration parameters for the autoscaling algorithm. You can
	// define one or more of the policies for an autoscaler: cpuUtilization,
	// customMetricUtilizations, and loadBalancingUtilization.
	// If none of these are specified, the default will be to autoscale based
	// on cpuUtilization to 0.6 or 60%.  Structure is documented below.
	AutoscalingPolicy *AutoscalarAutoscalingPolicy `pulumi:"autoscalingPolicy"`
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp *string `pulumi:"creationTimestamp"`
	// -
	// (Optional)
	// An optional description of this resource.
	Description *string `pulumi:"description"`
	// -
	// (Required)
	// The identifier (type) of the Stackdriver Monitoring metric.
	// The metric cannot have negative values.
	// The metric must have a value type of INT64 or DOUBLE.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The URI of the created resource.
	SelfLink *string `pulumi:"selfLink"`
	// -
	// (Required)
	// Fraction of backend capacity utilization (set in HTTP(s) load
	// balancing configuration) that autoscaler should maintain. Must
	// be a positive float value. If not defined, the default is 0.8.
	Target *string `pulumi:"target"`
	// -
	// (Optional)
	// URL of the zone where the instance group resides.
	Zone *string `pulumi:"zone"`
}

type AutoscalarState struct {
	// -
	// (Required)
	// The configuration parameters for the autoscaling algorithm. You can
	// define one or more of the policies for an autoscaler: cpuUtilization,
	// customMetricUtilizations, and loadBalancingUtilization.
	// If none of these are specified, the default will be to autoscale based
	// on cpuUtilization to 0.6 or 60%.  Structure is documented below.
	AutoscalingPolicy AutoscalarAutoscalingPolicyPtrInput
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringPtrInput
	// -
	// (Optional)
	// An optional description of this resource.
	Description pulumi.StringPtrInput
	// -
	// (Required)
	// The identifier (type) of the Stackdriver Monitoring metric.
	// The metric cannot have negative values.
	// The metric must have a value type of INT64 or DOUBLE.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The URI of the created resource.
	SelfLink pulumi.StringPtrInput
	// -
	// (Required)
	// Fraction of backend capacity utilization (set in HTTP(s) load
	// balancing configuration) that autoscaler should maintain. Must
	// be a positive float value. If not defined, the default is 0.8.
	Target pulumi.StringPtrInput
	// -
	// (Optional)
	// URL of the zone where the instance group resides.
	Zone pulumi.StringPtrInput
}

func (AutoscalarState) ElementType() reflect.Type {
	return reflect.TypeOf((*autoscalarState)(nil)).Elem()
}

type autoscalarArgs struct {
	// -
	// (Required)
	// The configuration parameters for the autoscaling algorithm. You can
	// define one or more of the policies for an autoscaler: cpuUtilization,
	// customMetricUtilizations, and loadBalancingUtilization.
	// If none of these are specified, the default will be to autoscale based
	// on cpuUtilization to 0.6 or 60%.  Structure is documented below.
	AutoscalingPolicy AutoscalarAutoscalingPolicy `pulumi:"autoscalingPolicy"`
	// -
	// (Optional)
	// An optional description of this resource.
	Description *string `pulumi:"description"`
	// -
	// (Required)
	// The identifier (type) of the Stackdriver Monitoring metric.
	// The metric cannot have negative values.
	// The metric must have a value type of INT64 or DOUBLE.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// -
	// (Required)
	// Fraction of backend capacity utilization (set in HTTP(s) load
	// balancing configuration) that autoscaler should maintain. Must
	// be a positive float value. If not defined, the default is 0.8.
	Target string `pulumi:"target"`
	// -
	// (Optional)
	// URL of the zone where the instance group resides.
	Zone *string `pulumi:"zone"`
}

// The set of arguments for constructing a Autoscalar resource.
type AutoscalarArgs struct {
	// -
	// (Required)
	// The configuration parameters for the autoscaling algorithm. You can
	// define one or more of the policies for an autoscaler: cpuUtilization,
	// customMetricUtilizations, and loadBalancingUtilization.
	// If none of these are specified, the default will be to autoscale based
	// on cpuUtilization to 0.6 or 60%.  Structure is documented below.
	AutoscalingPolicy AutoscalarAutoscalingPolicyInput
	// -
	// (Optional)
	// An optional description of this resource.
	Description pulumi.StringPtrInput
	// -
	// (Required)
	// The identifier (type) of the Stackdriver Monitoring metric.
	// The metric cannot have negative values.
	// The metric must have a value type of INT64 or DOUBLE.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// -
	// (Required)
	// Fraction of backend capacity utilization (set in HTTP(s) load
	// balancing configuration) that autoscaler should maintain. Must
	// be a positive float value. If not defined, the default is 0.8.
	Target pulumi.StringInput
	// -
	// (Optional)
	// URL of the zone where the instance group resides.
	Zone pulumi.StringPtrInput
}

func (AutoscalarArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*autoscalarArgs)(nil)).Elem()
}
