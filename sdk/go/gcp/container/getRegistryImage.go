// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package container

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// This data source fetches the project name, and provides the appropriate URLs to use for container registry for this project.
//
// The URLs are computed entirely offline - as long as the project exists, they will be valid, but this data source does not contact Google Container Registry (GCR) at any point.
//
// {{% examples %}}
// {{% /examples %}}
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/d/google_container_registry_image.html.markdown.
func GetRegistryImage(ctx *pulumi.Context, args *GetRegistryImageArgs, opts ...pulumi.InvokeOption) (*GetRegistryImageResult, error) {
	var rv GetRegistryImageResult
	err := ctx.Invoke("gcp:container/getRegistryImage:getRegistryImage", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRegistryImage.
type GetRegistryImageArgs struct {
	Digest *string `pulumi:"digest"`
	Name string `pulumi:"name"`
	Project *string `pulumi:"project"`
	Region *string `pulumi:"region"`
	Tag *string `pulumi:"tag"`
}


// A collection of values returned by getRegistryImage.
type GetRegistryImageResult struct {
	Digest *string `pulumi:"digest"`
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	ImageUrl string `pulumi:"imageUrl"`
	Name string `pulumi:"name"`
	Project string `pulumi:"project"`
	Region *string `pulumi:"region"`
	Tag *string `pulumi:"tag"`
}

