// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides available node types for Compute Engine sole-tenant nodes in a zone
// for a given project. For more information, see [the official documentation](https://cloud.google.com/compute/docs/nodes/#types) and [API](https://cloud.google.com/compute/docs/reference/rest/v1/nodeTypes).
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v3/go/gcp/compute"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := compute.GetNodeTypes(ctx, &compute.GetNodeTypesArgs{
// 			Zone: "us-central1-b",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = compute.NewNodeTemplate(ctx, "tmpl", &compute.NodeTemplateArgs{
// 			Region:   pulumi.String("us-central1"),
// 			NodeType: pulumi.String(data.Google_compute_node_types.Types.Names[0]),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetNodeTypes(ctx *pulumi.Context, args *GetNodeTypesArgs, opts ...pulumi.InvokeOption) (*GetNodeTypesResult, error) {
	var rv GetNodeTypesResult
	err := ctx.Invoke("gcp:compute/getNodeTypes:getNodeTypes", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getNodeTypes.
type GetNodeTypesArgs struct {
	// ID of the project to list available node types for.
	// Should match the project the nodes of this type will be deployed to.
	// Defaults to the project that the provider is authenticated with.
	Project *string `pulumi:"project"`
	// The zone to list node types for. Should be in zone of intended node groups and region of referencing node template. If `zone` is not specified, the provider-level zone must be set and is used
	// instead.
	Zone *string `pulumi:"zone"`
}

// A collection of values returned by getNodeTypes.
type GetNodeTypesResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of node types available in the given zone and project.
	Names   []string `pulumi:"names"`
	Project string   `pulumi:"project"`
	Zone    string   `pulumi:"zone"`
}
