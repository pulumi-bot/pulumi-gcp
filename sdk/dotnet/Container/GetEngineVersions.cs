// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Container
{
    public static class GetEngineVersions
    {
        /// <summary>
        /// Provides access to available Google Kubernetes Engine versions in a zone or region for a given project.
        /// 
        /// &gt; If you are using the `gcp.container.getEngineVersions` datasource with a
        /// regional cluster, ensure that you have provided a region as the `location` to
        /// the datasource. A region can have a different set of supported versions than
        /// its component zones, and not all zones in a region are guaranteed to
        /// support the same version.
        /// 
        /// {{% examples %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetEngineVersionsResult> InvokeAsync(GetEngineVersionsArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetEngineVersionsResult>("gcp:container/getEngineVersions:getEngineVersions", args ?? new GetEngineVersionsArgs(), options.WithVersion());
    }


    public sealed class GetEngineVersionsArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The location (region or zone) to list versions for.
        /// Must exactly match the location the cluster will be deployed in, or listed
        /// versions may not be available. If `location`, `region`, and `zone` are not
        /// specified, the provider-level zone must be set and is used instead.
        /// </summary>
        [Input("location")]
        public string? Location { get; set; }

        /// <summary>
        /// ID of the project to list available cluster versions for. Should match the project the cluster will be deployed to.
        /// Defaults to the project that the provider is authenticated with.
        /// </summary>
        [Input("project")]
        public string? Project { get; set; }

        /// <summary>
        /// If provided, the provider will only return versions
        /// that match the string prefix. For example, `1.11.` will match all `1.11` series
        /// releases. Since this is just a string match, it's recommended that you append a
        /// `.` after minor versions to ensure that prefixes such as `1.1` don't match
        /// versions like `1.12.5-gke.10` accidentally. See [the docs on versioning schema](https://cloud.google.com/kubernetes-engine/versioning-and-upgrades#versioning_scheme)
        /// for full details on how version strings are formatted.
        /// </summary>
        [Input("versionPrefix")]
        public string? VersionPrefix { get; set; }

        public GetEngineVersionsArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetEngineVersionsResult
    {
        /// <summary>
        /// Version of Kubernetes the service deploys by default.
        /// </summary>
        public readonly string DefaultClusterVersion;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The latest version available in the given zone for use with master instances.
        /// </summary>
        public readonly string LatestMasterVersion;
        /// <summary>
        /// The latest version available in the given zone for use with node instances.
        /// </summary>
        public readonly string LatestNodeVersion;
        public readonly string? Location;
        public readonly string? Project;
        /// <summary>
        /// A list of versions available in the given zone for use with master instances.
        /// </summary>
        public readonly ImmutableArray<string> ValidMasterVersions;
        /// <summary>
        /// A list of versions available in the given zone for use with node instances.
        /// </summary>
        public readonly ImmutableArray<string> ValidNodeVersions;
        public readonly string? VersionPrefix;

        [OutputConstructor]
        private GetEngineVersionsResult(
            string defaultClusterVersion,

            string id,

            string latestMasterVersion,

            string latestNodeVersion,

            string? location,

            string? project,

            ImmutableArray<string> validMasterVersions,

            ImmutableArray<string> validNodeVersions,

            string? versionPrefix)
        {
            DefaultClusterVersion = defaultClusterVersion;
            Id = id;
            LatestMasterVersion = latestMasterVersion;
            LatestNodeVersion = latestNodeVersion;
            Location = location;
            Project = project;
            ValidMasterVersions = validMasterVersions;
            ValidNodeVersions = validNodeVersions;
            VersionPrefix = versionPrefix;
        }
    }
}
