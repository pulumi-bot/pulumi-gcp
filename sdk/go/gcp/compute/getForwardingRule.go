// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package compute

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Get a forwarding rule within GCE from its name.
//
// {{% examples %}}
// {{% /examples %}}
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/d/datasource_compute_forwarding_rule.html.markdown.
func LookupForwardingRule(ctx *pulumi.Context, args *LookupForwardingRuleArgs, opts ...pulumi.InvokeOption) (*LookupForwardingRuleResult, error) {
	var rv LookupForwardingRuleResult
	err := ctx.Invoke("gcp:compute/getForwardingRule:getForwardingRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getForwardingRule.
type LookupForwardingRuleArgs struct {
	// The name of the forwarding rule.
	Name string `pulumi:"name"`
	// The project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The region in which the resource belongs. If it
	// is not provided, the project region is used.
	Region *string `pulumi:"region"`
}


// A collection of values returned by getForwardingRule.
type LookupForwardingRuleResult struct {
	// Backend service, if this forwarding rule has one.
	BackendService string `pulumi:"backendService"`
	// Description of this forwarding rule.
	Description string `pulumi:"description"`
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// IP address of this forwarding rule.
	IpAddress string `pulumi:"ipAddress"`
	// IP protocol of this forwarding rule.
	IpProtocol string `pulumi:"ipProtocol"`
	// Type of load balancing of this forwarding rule.
	LoadBalancingScheme string `pulumi:"loadBalancingScheme"`
	Name string `pulumi:"name"`
	// Network of this forwarding rule.
	Network string `pulumi:"network"`
	// Port range, if this forwarding rule has one.
	PortRange string `pulumi:"portRange"`
	// List of ports to use for internal load balancing, if this forwarding rule has any.
	Ports []string `pulumi:"ports"`
	Project string `pulumi:"project"`
	// Region of this forwarding rule.
	Region string `pulumi:"region"`
	// The URI of the resource.
	SelfLink string `pulumi:"selfLink"`
	// Subnetwork of this forwarding rule.
	Subnetwork string `pulumi:"subnetwork"`
	// URL of the target pool, if this forwarding rule has one.
	Target string `pulumi:"target"`
}

