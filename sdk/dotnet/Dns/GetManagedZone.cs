// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Dns
{
    public static partial class Invokes
    {
        /// <summary>
        /// Provides access to a zone's attributes within Google Cloud DNS.
        /// For more information see
        /// [the official documentation](https://cloud.google.com/dns/zones/)
        /// and
        /// [API](https://cloud.google.com/dns/api/v1/managedZones).
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/d/dns_managed_zone.html.markdown.
        /// </summary>
        public static Task<GetManagedZoneResult> GetManagedZone(GetManagedZoneArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetManagedZoneResult>("gcp:dns/getManagedZone:getManagedZone", args ?? ResourceArgs.Empty, options.WithVersion());
    }

    public sealed class GetManagedZoneArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A unique name for the resource.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The ID of the project for the Google Cloud DNS zone.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        public GetManagedZoneArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetManagedZoneResult
    {
        /// <summary>
        /// A textual description field.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The fully qualified DNS name of this zone, e.g. `example.com.`.
        /// </summary>
        public readonly string DnsName;
        public readonly string Name;
        /// <summary>
        /// The list of nameservers that will be authoritative for this
        /// domain. Use NS records to redirect from your DNS provider to these names,
        /// thus making Google Cloud DNS authoritative for this zone.
        /// </summary>
        public readonly ImmutableArray<string> NameServers;
        public readonly string? Project;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetManagedZoneResult(
            string description,
            string dnsName,
            string name,
            ImmutableArray<string> nameServers,
            string? project,
            string id)
        {
            Description = description;
            DnsName = dnsName;
            Name = name;
            NameServers = nameServers;
            Project = project;
            Id = id;
        }
    }
}
