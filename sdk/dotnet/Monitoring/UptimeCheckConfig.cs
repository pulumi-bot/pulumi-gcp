// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Monitoring
{
    /// <summary>
    /// This message configures which resources and services to monitor for availability.
    /// 
    /// 
    /// To get more information about UptimeCheckConfig, see:
    /// 
    /// * [API documentation](https://cloud.google.com/monitoring/api/ref_v3/rest/v3/projects.uptimeCheckConfigs)
    /// * How-to Guides
    ///     * [Official Documentation](https://cloud.google.com/monitoring/uptime-checks/)
    /// 
    /// &gt; **Warning:** All arguments including `http_check.auth_info.password` will be stored in the raw
    /// state as plain-text. [Read more about secrets in state](https://www.pulumi.com/docs/intro/concepts/programming-model/#secrets).
    /// 
    /// ## Example Usage - Uptime Check Config Http
    /// {{% example %}}
    /// 
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var http = new Gcp.Monitoring.UptimeCheckConfig("http", new Gcp.Monitoring.UptimeCheckConfigArgs
    ///         {
    ///             ContentMatchers = 
    ///             {
    ///                 new Gcp.Monitoring.Inputs.UptimeCheckConfigContentMatcherArgs
    ///                 {
    ///                     Content = "example",
    ///                 },
    ///             },
    ///             DisplayName = "http-uptime-check",
    ///             HttpCheck = new Gcp.Monitoring.Inputs.UptimeCheckConfigHttpCheckArgs
    ///             {
    ///                 Path = "/some-path",
    ///                 Port = "8010",
    ///             },
    ///             MonitoredResource = new Gcp.Monitoring.Inputs.UptimeCheckConfigMonitoredResourceArgs
    ///             {
    ///                 Labels = 
    ///                 {
    ///                     { "host", "192.168.1.1" },
    ///                     { "project_id", "my-project-name" },
    ///                 },
    ///                 Type = "uptime_url",
    ///             },
    ///             Timeout = "60s",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// {{% /example %}}
    /// ## Example Usage - Uptime Check Config Https
    /// {{% example %}}
    /// 
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var https = new Gcp.Monitoring.UptimeCheckConfig("https", new Gcp.Monitoring.UptimeCheckConfigArgs
    ///         {
    ///             ContentMatchers = 
    ///             {
    ///                 new Gcp.Monitoring.Inputs.UptimeCheckConfigContentMatcherArgs
    ///                 {
    ///                     Content = "example",
    ///                 },
    ///             },
    ///             DisplayName = "https-uptime-check",
    ///             HttpCheck = new Gcp.Monitoring.Inputs.UptimeCheckConfigHttpCheckArgs
    ///             {
    ///                 Path = "/some-path",
    ///                 Port = "443",
    ///                 UseSsl = true,
    ///                 ValidateSsl = true,
    ///             },
    ///             MonitoredResource = new Gcp.Monitoring.Inputs.UptimeCheckConfigMonitoredResourceArgs
    ///             {
    ///                 Labels = 
    ///                 {
    ///                     { "host", "192.168.1.1" },
    ///                     { "project_id", "my-project-name" },
    ///                 },
    ///                 Type = "uptime_url",
    ///             },
    ///             Timeout = "60s",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// {{% /example %}}
    /// </summary>
    public partial class UptimeCheckConfig : Pulumi.CustomResource
    {
        /// <summary>
        /// The expected content on the page the check is run against. Currently, only the first entry in the list is supported, and other entries will be ignored. The server will look for an exact match of the string in the page response's content. This field is optional and should only be specified if a content match is required.  Structure is documented below.
        /// </summary>
        [Output("contentMatchers")]
        public Output<ImmutableArray<Outputs.UptimeCheckConfigContentMatcher>> ContentMatchers { get; private set; } = null!;

        /// <summary>
        /// A human-friendly name for the uptime check configuration. The display name should be unique within a Stackdriver Workspace in order to make it easier to identify; however, uniqueness is not enforced.
        /// </summary>
        [Output("displayName")]
        public Output<string> DisplayName { get; private set; } = null!;

        /// <summary>
        /// Contains information needed to make an HTTP or HTTPS check.  Structure is documented below.
        /// </summary>
        [Output("httpCheck")]
        public Output<Outputs.UptimeCheckConfigHttpCheck?> HttpCheck { get; private set; } = null!;

        /// <summary>
        /// The monitored resource (https://cloud.google.com/monitoring/api/resources) associated with the configuration. The following monitored resource types are supported for uptime checks:  uptime_url  gce_instance  gae_app  aws_ec2_instance  aws_elb_load_balancer  Structure is documented below.
        /// </summary>
        [Output("monitoredResource")]
        public Output<Outputs.UptimeCheckConfigMonitoredResource?> MonitoredResource { get; private set; } = null!;

        /// <summary>
        /// A unique resource name for this UptimeCheckConfig. The format is
        /// projects/[PROJECT_ID]/uptimeCheckConfigs/[UPTIME_CHECK_ID].
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// How often, in seconds, the uptime check is performed. Currently, the only supported values are 60s (1 minute), 300s (5 minutes), 600s (10 minutes), and 900s (15 minutes). Optional, defaults to 300s.
        /// </summary>
        [Output("period")]
        public Output<string?> Period { get; private set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// The group resource associated with the configuration.  Structure is documented below.
        /// </summary>
        [Output("resourceGroup")]
        public Output<Outputs.UptimeCheckConfigResourceGroup?> ResourceGroup { get; private set; } = null!;

        /// <summary>
        /// The list of regions from which the check will be run. Some regions contain one location, and others contain more than one. If this field is specified, enough regions to include a minimum of 3 locations must be provided, or an error message is returned. Not specifying this field will result in uptime checks running from all regions.
        /// </summary>
        [Output("selectedRegions")]
        public Output<ImmutableArray<string>> SelectedRegions { get; private set; } = null!;

        /// <summary>
        /// Contains information needed to make a TCP check.  Structure is documented below.
        /// </summary>
        [Output("tcpCheck")]
        public Output<Outputs.UptimeCheckConfigTcpCheck?> TcpCheck { get; private set; } = null!;

        /// <summary>
        /// The maximum amount of time to wait for the request to complete (must be between 1 and 60 seconds). Accepted formats https://developers.google.com/protocol-buffers/docs/reference/google.protobuf#google.protobuf.Duration
        /// </summary>
        [Output("timeout")]
        public Output<string> Timeout { get; private set; } = null!;

        /// <summary>
        /// The id of the uptime check
        /// </summary>
        [Output("uptimeCheckId")]
        public Output<string> UptimeCheckId { get; private set; } = null!;


        /// <summary>
        /// Create a UptimeCheckConfig resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public UptimeCheckConfig(string name, UptimeCheckConfigArgs args, CustomResourceOptions? options = null)
            : base("gcp:monitoring/uptimeCheckConfig:UptimeCheckConfig", name, args ?? new UptimeCheckConfigArgs(), MakeResourceOptions(options, ""))
        {
        }

        private UptimeCheckConfig(string name, Input<string> id, UptimeCheckConfigState? state = null, CustomResourceOptions? options = null)
            : base("gcp:monitoring/uptimeCheckConfig:UptimeCheckConfig", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing UptimeCheckConfig resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static UptimeCheckConfig Get(string name, Input<string> id, UptimeCheckConfigState? state = null, CustomResourceOptions? options = null)
        {
            return new UptimeCheckConfig(name, id, state, options);
        }
    }

    public sealed class UptimeCheckConfigArgs : Pulumi.ResourceArgs
    {
        [Input("contentMatchers")]
        private InputList<Inputs.UptimeCheckConfigContentMatcherArgs>? _contentMatchers;

        /// <summary>
        /// The expected content on the page the check is run against. Currently, only the first entry in the list is supported, and other entries will be ignored. The server will look for an exact match of the string in the page response's content. This field is optional and should only be specified if a content match is required.  Structure is documented below.
        /// </summary>
        public InputList<Inputs.UptimeCheckConfigContentMatcherArgs> ContentMatchers
        {
            get => _contentMatchers ?? (_contentMatchers = new InputList<Inputs.UptimeCheckConfigContentMatcherArgs>());
            set => _contentMatchers = value;
        }

        /// <summary>
        /// A human-friendly name for the uptime check configuration. The display name should be unique within a Stackdriver Workspace in order to make it easier to identify; however, uniqueness is not enforced.
        /// </summary>
        [Input("displayName", required: true)]
        public Input<string> DisplayName { get; set; } = null!;

        /// <summary>
        /// Contains information needed to make an HTTP or HTTPS check.  Structure is documented below.
        /// </summary>
        [Input("httpCheck")]
        public Input<Inputs.UptimeCheckConfigHttpCheckArgs>? HttpCheck { get; set; }

        /// <summary>
        /// The monitored resource (https://cloud.google.com/monitoring/api/resources) associated with the configuration. The following monitored resource types are supported for uptime checks:  uptime_url  gce_instance  gae_app  aws_ec2_instance  aws_elb_load_balancer  Structure is documented below.
        /// </summary>
        [Input("monitoredResource")]
        public Input<Inputs.UptimeCheckConfigMonitoredResourceArgs>? MonitoredResource { get; set; }

        /// <summary>
        /// How often, in seconds, the uptime check is performed. Currently, the only supported values are 60s (1 minute), 300s (5 minutes), 600s (10 minutes), and 900s (15 minutes). Optional, defaults to 300s.
        /// </summary>
        [Input("period")]
        public Input<string>? Period { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The group resource associated with the configuration.  Structure is documented below.
        /// </summary>
        [Input("resourceGroup")]
        public Input<Inputs.UptimeCheckConfigResourceGroupArgs>? ResourceGroup { get; set; }

        [Input("selectedRegions")]
        private InputList<string>? _selectedRegions;

        /// <summary>
        /// The list of regions from which the check will be run. Some regions contain one location, and others contain more than one. If this field is specified, enough regions to include a minimum of 3 locations must be provided, or an error message is returned. Not specifying this field will result in uptime checks running from all regions.
        /// </summary>
        public InputList<string> SelectedRegions
        {
            get => _selectedRegions ?? (_selectedRegions = new InputList<string>());
            set => _selectedRegions = value;
        }

        /// <summary>
        /// Contains information needed to make a TCP check.  Structure is documented below.
        /// </summary>
        [Input("tcpCheck")]
        public Input<Inputs.UptimeCheckConfigTcpCheckArgs>? TcpCheck { get; set; }

        /// <summary>
        /// The maximum amount of time to wait for the request to complete (must be between 1 and 60 seconds). Accepted formats https://developers.google.com/protocol-buffers/docs/reference/google.protobuf#google.protobuf.Duration
        /// </summary>
        [Input("timeout", required: true)]
        public Input<string> Timeout { get; set; } = null!;

        public UptimeCheckConfigArgs()
        {
        }
    }

    public sealed class UptimeCheckConfigState : Pulumi.ResourceArgs
    {
        [Input("contentMatchers")]
        private InputList<Inputs.UptimeCheckConfigContentMatcherGetArgs>? _contentMatchers;

        /// <summary>
        /// The expected content on the page the check is run against. Currently, only the first entry in the list is supported, and other entries will be ignored. The server will look for an exact match of the string in the page response's content. This field is optional and should only be specified if a content match is required.  Structure is documented below.
        /// </summary>
        public InputList<Inputs.UptimeCheckConfigContentMatcherGetArgs> ContentMatchers
        {
            get => _contentMatchers ?? (_contentMatchers = new InputList<Inputs.UptimeCheckConfigContentMatcherGetArgs>());
            set => _contentMatchers = value;
        }

        /// <summary>
        /// A human-friendly name for the uptime check configuration. The display name should be unique within a Stackdriver Workspace in order to make it easier to identify; however, uniqueness is not enforced.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// Contains information needed to make an HTTP or HTTPS check.  Structure is documented below.
        /// </summary>
        [Input("httpCheck")]
        public Input<Inputs.UptimeCheckConfigHttpCheckGetArgs>? HttpCheck { get; set; }

        /// <summary>
        /// The monitored resource (https://cloud.google.com/monitoring/api/resources) associated with the configuration. The following monitored resource types are supported for uptime checks:  uptime_url  gce_instance  gae_app  aws_ec2_instance  aws_elb_load_balancer  Structure is documented below.
        /// </summary>
        [Input("monitoredResource")]
        public Input<Inputs.UptimeCheckConfigMonitoredResourceGetArgs>? MonitoredResource { get; set; }

        /// <summary>
        /// A unique resource name for this UptimeCheckConfig. The format is
        /// projects/[PROJECT_ID]/uptimeCheckConfigs/[UPTIME_CHECK_ID].
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// How often, in seconds, the uptime check is performed. Currently, the only supported values are 60s (1 minute), 300s (5 minutes), 600s (10 minutes), and 900s (15 minutes). Optional, defaults to 300s.
        /// </summary>
        [Input("period")]
        public Input<string>? Period { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The group resource associated with the configuration.  Structure is documented below.
        /// </summary>
        [Input("resourceGroup")]
        public Input<Inputs.UptimeCheckConfigResourceGroupGetArgs>? ResourceGroup { get; set; }

        [Input("selectedRegions")]
        private InputList<string>? _selectedRegions;

        /// <summary>
        /// The list of regions from which the check will be run. Some regions contain one location, and others contain more than one. If this field is specified, enough regions to include a minimum of 3 locations must be provided, or an error message is returned. Not specifying this field will result in uptime checks running from all regions.
        /// </summary>
        public InputList<string> SelectedRegions
        {
            get => _selectedRegions ?? (_selectedRegions = new InputList<string>());
            set => _selectedRegions = value;
        }

        /// <summary>
        /// Contains information needed to make a TCP check.  Structure is documented below.
        /// </summary>
        [Input("tcpCheck")]
        public Input<Inputs.UptimeCheckConfigTcpCheckGetArgs>? TcpCheck { get; set; }

        /// <summary>
        /// The maximum amount of time to wait for the request to complete (must be between 1 and 60 seconds). Accepted formats https://developers.google.com/protocol-buffers/docs/reference/google.protobuf#google.protobuf.Duration
        /// </summary>
        [Input("timeout")]
        public Input<string>? Timeout { get; set; }

        /// <summary>
        /// The id of the uptime check
        /// </summary>
        [Input("uptimeCheckId")]
        public Input<string>? UptimeCheckId { get; set; }

        public UptimeCheckConfigState()
        {
        }
    }
}
