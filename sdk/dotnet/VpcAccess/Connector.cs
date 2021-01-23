// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.VpcAccess
{
    /// <summary>
    /// Serverless VPC Access connector resource.
    /// 
    /// To get more information about Connector, see:
    /// 
    /// * [API documentation](https://cloud.google.com/vpc/docs/reference/vpcaccess/rest/v1/projects.locations.connectors)
    /// * How-to Guides
    ///     * [Configuring Serverless VPC Access](https://cloud.google.com/vpc/docs/configure-serverless-vpc-access)
    /// 
    /// ## Example Usage
    /// ### VPC Access Connector
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var connector = new Gcp.VpcAccess.Connector("connector", new Gcp.VpcAccess.ConnectorArgs
    ///         {
    ///             IpCidrRange = "10.8.0.0/28",
    ///             Network = "default",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Connector can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import gcp:vpcaccess/connector:Connector default projects/{{project}}/locations/{{region}}/connectors/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:vpcaccess/connector:Connector default {{project}}/{{region}}/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:vpcaccess/connector:Connector default {{region}}/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:vpcaccess/connector:Connector default {{name}}
    /// ```
    /// </summary>
    [GcpResourceType("gcp:vpcaccess/connector:Connector")]
    public partial class Connector : Pulumi.CustomResource
    {
        /// <summary>
        /// The range of internal addresses that follows RFC 4632 notation. Example: `10.132.0.0/28`.
        /// </summary>
        [Output("ipCidrRange")]
        public Output<string> IpCidrRange { get; private set; } = null!;

        /// <summary>
        /// Maximum throughput of the connector in Mbps, must be greater than `min_throughput`. Default is 1000.
        /// </summary>
        [Output("maxThroughput")]
        public Output<int?> MaxThroughput { get; private set; } = null!;

        /// <summary>
        /// Minimum throughput of the connector in Mbps. Default and min is 200.
        /// </summary>
        [Output("minThroughput")]
        public Output<int?> MinThroughput { get; private set; } = null!;

        /// <summary>
        /// The name of the resource (Max 25 characters).
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Name of a VPC network.
        /// </summary>
        [Output("network")]
        public Output<string> Network { get; private set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// Region where the VPC Access connector resides. If it is not provided, the provider region is used.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// The fully qualified name of this VPC connector
        /// </summary>
        [Output("selfLink")]
        public Output<string> SelfLink { get; private set; } = null!;

        /// <summary>
        /// State of the VPC access connector.
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;


        /// <summary>
        /// Create a Connector resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Connector(string name, ConnectorArgs args, CustomResourceOptions? options = null)
            : base("gcp:vpcaccess/connector:Connector", name, args ?? new ConnectorArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Connector(string name, Input<string> id, ConnectorState? state = null, CustomResourceOptions? options = null)
            : base("gcp:vpcaccess/connector:Connector", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Connector resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Connector Get(string name, Input<string> id, ConnectorState? state = null, CustomResourceOptions? options = null)
        {
            return new Connector(name, id, state, options);
        }
    }

    public sealed class ConnectorArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The range of internal addresses that follows RFC 4632 notation. Example: `10.132.0.0/28`.
        /// </summary>
        [Input("ipCidrRange", required: true)]
        public Input<string> IpCidrRange { get; set; } = null!;

        /// <summary>
        /// Maximum throughput of the connector in Mbps, must be greater than `min_throughput`. Default is 1000.
        /// </summary>
        [Input("maxThroughput")]
        public Input<int>? MaxThroughput { get; set; }

        /// <summary>
        /// Minimum throughput of the connector in Mbps. Default and min is 200.
        /// </summary>
        [Input("minThroughput")]
        public Input<int>? MinThroughput { get; set; }

        /// <summary>
        /// The name of the resource (Max 25 characters).
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Name of a VPC network.
        /// </summary>
        [Input("network", required: true)]
        public Input<string> Network { get; set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// Region where the VPC Access connector resides. If it is not provided, the provider region is used.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        public ConnectorArgs()
        {
        }
    }

    public sealed class ConnectorState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The range of internal addresses that follows RFC 4632 notation. Example: `10.132.0.0/28`.
        /// </summary>
        [Input("ipCidrRange")]
        public Input<string>? IpCidrRange { get; set; }

        /// <summary>
        /// Maximum throughput of the connector in Mbps, must be greater than `min_throughput`. Default is 1000.
        /// </summary>
        [Input("maxThroughput")]
        public Input<int>? MaxThroughput { get; set; }

        /// <summary>
        /// Minimum throughput of the connector in Mbps. Default and min is 200.
        /// </summary>
        [Input("minThroughput")]
        public Input<int>? MinThroughput { get; set; }

        /// <summary>
        /// The name of the resource (Max 25 characters).
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Name of a VPC network.
        /// </summary>
        [Input("network")]
        public Input<string>? Network { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// Region where the VPC Access connector resides. If it is not provided, the provider region is used.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The fully qualified name of this VPC connector
        /// </summary>
        [Input("selfLink")]
        public Input<string>? SelfLink { get; set; }

        /// <summary>
        /// State of the VPC access connector.
        /// </summary>
        [Input("state")]
        public Input<string>? State { get; set; }

        public ConnectorState()
        {
        }
    }
}
