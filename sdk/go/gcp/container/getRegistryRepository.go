// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package container

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// This data source fetches the project name, and provides the appropriate URLs to use for container registry for this project.
// 
// The URLs are computed entirely offline - as long as the project exists, they will be valid, but this data source does not contact Google Container Registry (GCR) at any point.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/d/container_registry_repository.html.markdown.
func LookupRegistryRepository(ctx *pulumi.Context, args *GetRegistryRepositoryArgs) (*GetRegistryRepositoryResult, error) {
	inputs := make(map[string]interface{})
	if args != nil {
		inputs["project"] = args.Project
		inputs["region"] = args.Region
	}
	outputs, err := ctx.Invoke("gcp:container/getRegistryRepository:getRegistryRepository", inputs)
	if err != nil {
		return nil, err
	}
	return &GetRegistryRepositoryResult{
		Project: outputs["project"],
		Region: outputs["region"],
		RepositoryUrl: outputs["repositoryUrl"],
		Id: outputs["id"],
	}, nil
}

// A collection of arguments for invoking getRegistryRepository.
type GetRegistryRepositoryArgs struct {
	Project interface{}
	Region interface{}
}

// A collection of values returned by getRegistryRepository.
type GetRegistryRepositoryResult struct {
	Project interface{}
	Region interface{}
	RepositoryUrl interface{}
	// id is the provider-assigned unique ID for this managed resource.
	Id interface{}
}
