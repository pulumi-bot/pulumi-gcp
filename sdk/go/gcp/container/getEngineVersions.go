// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package container

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides access to available Google Container Engine versions in a zone or region for a given project.
// 
// > If you are using the `google_container_engine_versions` datasource with a regional cluster, ensure that you have provided a `region`
// to the datasource. A `region` can have a different set of supported versions than its corresponding `zone`s, and not all `zone`s in a 
// `region` are guaranteed to support the same version.
func LookupEngineVersions(ctx *pulumi.Context, args *GetEngineVersionsArgs) (*GetEngineVersionsResult, error) {
	inputs := make(map[string]interface{})
	if args != nil {
		inputs["project"] = args.Project
		inputs["region"] = args.Region
		inputs["versionPrefix"] = args.VersionPrefix
		inputs["zone"] = args.Zone
	}
	outputs, err := ctx.Invoke("gcp:container/getEngineVersions:getEngineVersions", inputs)
	if err != nil {
		return nil, err
	}
	return &GetEngineVersionsResult{
		DefaultClusterVersion: outputs["defaultClusterVersion"],
		LatestMasterVersion: outputs["latestMasterVersion"],
		LatestNodeVersion: outputs["latestNodeVersion"],
		Project: outputs["project"],
		Region: outputs["region"],
		ValidMasterVersions: outputs["validMasterVersions"],
		ValidNodeVersions: outputs["validNodeVersions"],
		VersionPrefix: outputs["versionPrefix"],
		Zone: outputs["zone"],
		Id: outputs["id"],
	}, nil
}

// A collection of arguments for invoking getEngineVersions.
type GetEngineVersionsArgs struct {
	// ID of the project to list available cluster versions for. Should match the project the cluster will be deployed to.
	// Defaults to the project that the provider is authenticated with.
	Project interface{}
	Region interface{}
	// If provided, Terraform will only return versions
	// that match the string prefix. For example, `1.11.` will match all `1.11` series
	// releases. Since this is just a string match, it's recommended that you append a
	// `.` after minor versions to ensure that prefixes such as `1.1` don't match
	// versions like `1.12.5-gke.10` accidentally. See [the docs on versioning schema](https://cloud.google.com/kubernetes-engine/versioning-and-upgrades#versioning_scheme)
	// for full details on how version strings are formatted.
	VersionPrefix interface{}
	// Zone to list available cluster versions for. Should match the zone the cluster will be deployed in.
	// If not specified, the provider-level zone is used. One of zone or provider-level zone is required.
	Zone interface{}
}

// A collection of values returned by getEngineVersions.
type GetEngineVersionsResult struct {
	// Version of Kubernetes the service deploys by default.
	DefaultClusterVersion interface{}
	// The latest version available in the given zone for use with master instances.
	LatestMasterVersion interface{}
	// The latest version available in the given zone for use with node instances.
	LatestNodeVersion interface{}
	Project interface{}
	Region interface{}
	// A list of versions available in the given zone for use with master instances.
	ValidMasterVersions interface{}
	// A list of versions available in the given zone for use with node instances.
	ValidNodeVersions interface{}
	VersionPrefix interface{}
	Zone interface{}
	// id is the provider-assigned unique ID for this managed resource.
	Id interface{}
}
