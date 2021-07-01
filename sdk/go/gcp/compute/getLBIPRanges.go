// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to access IP ranges in your firewall rules.
//
// https://cloud.google.com/compute/docs/load-balancing/health-checks#health_check_source_ips_and_firewall_rules
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/compute"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		ranges, err := compute.GetLBIPRanges(ctx, nil, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = compute.NewFirewall(ctx, "lb", &compute.FirewallArgs{
// 			Network: pulumi.Any(google_compute_network.Main.Name),
// 			Allows: compute.FirewallAllowArray{
// 				&compute.FirewallAllowArgs{
// 					Protocol: pulumi.String("tcp"),
// 					Ports: pulumi.StringArray{
// 						pulumi.String("80"),
// 					},
// 				},
// 			},
// 			SourceRanges: interface{}(ranges.Networks),
// 			TargetTags: pulumi.StringArray{
// 				pulumi.String("InstanceBehindLoadBalancer"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetLBIPRanges(ctx *pulumi.Context, opts ...pulumi.InvokeOption) (*GetLBIPRangesResult, error) {
	var rv GetLBIPRangesResult
	err := ctx.Invoke("gcp:compute/getLBIPRanges:getLBIPRanges", nil, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of values returned by getLBIPRanges.
type GetLBIPRangesResult struct {
	// The IP ranges used for health checks when **HTTP(S), SSL proxy, TCP proxy, and Internal load balancing** is used
	HttpSslTcpInternals []string `pulumi:"httpSslTcpInternals"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The IP ranges used for health checks when **Network load balancing** is used
	Networks []string `pulumi:"networks"`
}
