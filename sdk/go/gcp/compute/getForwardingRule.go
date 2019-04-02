// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Get a forwarding rule within GCE from its name.
func LookupForwardingRule(ctx *pulumi.Context, args *GetForwardingRuleArgs) (*GetForwardingRuleResult, error) {
	inputs := make(map[string]interface{})
	if args != nil {
		inputs["name"] = args.Name
		inputs["project"] = args.Project
		inputs["region"] = args.Region
	}
	outputs, err := ctx.Invoke("gcp:compute/getForwardingRule:getForwardingRule", inputs)
	if err != nil {
		return nil, err
	}
	return &GetForwardingRuleResult{
		BackendService: outputs["backendService"],
		Description: outputs["description"],
		IpAddress: outputs["ipAddress"],
		IpProtocol: outputs["ipProtocol"],
		LoadBalancingScheme: outputs["loadBalancingScheme"],
		Name: outputs["name"],
		Network: outputs["network"],
		PortRange: outputs["portRange"],
		Ports: outputs["ports"],
		Project: outputs["project"],
		Region: outputs["region"],
		SelfLink: outputs["selfLink"],
		Subnetwork: outputs["subnetwork"],
		Target: outputs["target"],
		Id: outputs["id"],
	}, nil
}

// A collection of arguments for invoking getForwardingRule.
type GetForwardingRuleArgs struct {
	// The name of the forwarding rule.
	Name interface{}
	// The project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project interface{}
	// The region in which the resource belongs. If it
	// is not provided, the project region is used.
	Region interface{}
}

// A collection of values returned by getForwardingRule.
type GetForwardingRuleResult struct {
	// Backend service, if this forwarding rule has one.
	BackendService interface{}
	// Description of this forwarding rule.
	Description interface{}
	// IP address of this forwarding rule.
	IpAddress interface{}
	// IP protocol of this forwarding rule.
	IpProtocol interface{}
	// Type of load balancing of this forwarding rule.
	LoadBalancingScheme interface{}
	Name interface{}
	// Network of this forwarding rule.
	Network interface{}
	// Port range, if this forwarding rule has one.
	PortRange interface{}
	// List of ports to use for internal load balancing, if this forwarding rule has any.
	Ports interface{}
	Project interface{}
	// Region of this forwarding rule.
	Region interface{}
	// The URI of the resource.
	SelfLink interface{}
	// Subnetwork of this forwarding rule.
	Subnetwork interface{}
	// URL of the target pool, if this forwarding rule has one.
	Target interface{}
	// id is the provider-assigned unique ID for this managed resource.
	Id interface{}
}
