// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Compute
{
    /// <summary>
    /// Health Checks determine whether instances are responsive and able to do work.
    /// They are an important part of a comprehensive load balancing configuration,
    /// as they enable monitoring instances behind load balancers.
    /// 
    /// Health Checks poll instances at a specified interval. Instances that
    /// do not respond successfully to some number of probes in a row are marked
    /// as unhealthy. No new connections are sent to unhealthy instances,
    /// though existing connections will continue. The health check will
    /// continue to poll unhealthy instances. If an instance later responds
    /// successfully to some number of consecutive probes, it is marked
    /// healthy again and can receive new connections.
    /// 
    /// 
    /// To get more information about RegionHealthCheck, see:
    /// 
    /// * [API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/regionHealthChecks)
    /// * How-to Guides
    ///     * [Official Documentation](https://cloud.google.com/load-balancing/docs/health-checks)
    /// 
    /// ## Example Usage - Region Health Check Tcp
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
    ///         var tcp_region_health_check = new Gcp.Compute.RegionHealthCheck("tcp-region-health-check", new Gcp.Compute.RegionHealthCheckArgs
    ///         {
    ///             CheckIntervalSec = 1,
    ///             TcpHealthCheck = new Gcp.Compute.Inputs.RegionHealthCheckTcpHealthCheckArgs
    ///             {
    ///                 Port = 80,
    ///             },
    ///             TimeoutSec = 1,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ## Example Usage - Region Health Check Tcp Full
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
    ///         var tcp_region_health_check = new Gcp.Compute.RegionHealthCheck("tcp-region-health-check", new Gcp.Compute.RegionHealthCheckArgs
    ///         {
    ///             CheckIntervalSec = 1,
    ///             Description = "Health check via tcp",
    ///             HealthyThreshold = 4,
    ///             TcpHealthCheck = new Gcp.Compute.Inputs.RegionHealthCheckTcpHealthCheckArgs
    ///             {
    ///                 PortName = "health-check-port",
    ///                 PortSpecification = "USE_NAMED_PORT",
    ///                 ProxyHeader = "NONE",
    ///                 Request = "ARE YOU HEALTHY?",
    ///                 Response = "I AM HEALTHY",
    ///             },
    ///             TimeoutSec = 1,
    ///             UnhealthyThreshold = 5,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ## Example Usage - Region Health Check Ssl
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
    ///         var ssl_region_health_check = new Gcp.Compute.RegionHealthCheck("ssl-region-health-check", new Gcp.Compute.RegionHealthCheckArgs
    ///         {
    ///             CheckIntervalSec = 1,
    ///             SslHealthCheck = new Gcp.Compute.Inputs.RegionHealthCheckSslHealthCheckArgs
    ///             {
    ///                 Port = 443,
    ///             },
    ///             TimeoutSec = 1,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ## Example Usage - Region Health Check Ssl Full
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
    ///         var ssl_region_health_check = new Gcp.Compute.RegionHealthCheck("ssl-region-health-check", new Gcp.Compute.RegionHealthCheckArgs
    ///         {
    ///             CheckIntervalSec = 1,
    ///             Description = "Health check via ssl",
    ///             HealthyThreshold = 4,
    ///             SslHealthCheck = new Gcp.Compute.Inputs.RegionHealthCheckSslHealthCheckArgs
    ///             {
    ///                 PortName = "health-check-port",
    ///                 PortSpecification = "USE_NAMED_PORT",
    ///                 ProxyHeader = "NONE",
    ///                 Request = "ARE YOU HEALTHY?",
    ///                 Response = "I AM HEALTHY",
    ///             },
    ///             TimeoutSec = 1,
    ///             UnhealthyThreshold = 5,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ## Example Usage - Region Health Check Http
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
    ///         var http_region_health_check = new Gcp.Compute.RegionHealthCheck("http-region-health-check", new Gcp.Compute.RegionHealthCheckArgs
    ///         {
    ///             CheckIntervalSec = 1,
    ///             HttpHealthCheck = new Gcp.Compute.Inputs.RegionHealthCheckHttpHealthCheckArgs
    ///             {
    ///                 Port = 80,
    ///             },
    ///             TimeoutSec = 1,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ## Example Usage - Region Health Check Http Full
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
    ///         var http_region_health_check = new Gcp.Compute.RegionHealthCheck("http-region-health-check", new Gcp.Compute.RegionHealthCheckArgs
    ///         {
    ///             CheckIntervalSec = 1,
    ///             Description = "Health check via http",
    ///             HealthyThreshold = 4,
    ///             HttpHealthCheck = new Gcp.Compute.Inputs.RegionHealthCheckHttpHealthCheckArgs
    ///             {
    ///                 Host = "1.2.3.4",
    ///                 PortName = "health-check-port",
    ///                 PortSpecification = "USE_NAMED_PORT",
    ///                 ProxyHeader = "NONE",
    ///                 RequestPath = "/mypath",
    ///                 Response = "I AM HEALTHY",
    ///             },
    ///             TimeoutSec = 1,
    ///             UnhealthyThreshold = 5,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ## Example Usage - Region Health Check Https
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
    ///         var https_region_health_check = new Gcp.Compute.RegionHealthCheck("https-region-health-check", new Gcp.Compute.RegionHealthCheckArgs
    ///         {
    ///             CheckIntervalSec = 1,
    ///             HttpsHealthCheck = new Gcp.Compute.Inputs.RegionHealthCheckHttpsHealthCheckArgs
    ///             {
    ///                 Port = 443,
    ///             },
    ///             TimeoutSec = 1,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ## Example Usage - Region Health Check Https Full
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
    ///         var https_region_health_check = new Gcp.Compute.RegionHealthCheck("https-region-health-check", new Gcp.Compute.RegionHealthCheckArgs
    ///         {
    ///             CheckIntervalSec = 1,
    ///             Description = "Health check via https",
    ///             HealthyThreshold = 4,
    ///             HttpsHealthCheck = new Gcp.Compute.Inputs.RegionHealthCheckHttpsHealthCheckArgs
    ///             {
    ///                 Host = "1.2.3.4",
    ///                 PortName = "health-check-port",
    ///                 PortSpecification = "USE_NAMED_PORT",
    ///                 ProxyHeader = "NONE",
    ///                 RequestPath = "/mypath",
    ///                 Response = "I AM HEALTHY",
    ///             },
    ///             TimeoutSec = 1,
    ///             UnhealthyThreshold = 5,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ## Example Usage - Region Health Check Http2
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
    ///         var http2_region_health_check = new Gcp.Compute.RegionHealthCheck("http2-region-health-check", new Gcp.Compute.RegionHealthCheckArgs
    ///         {
    ///             CheckIntervalSec = 1,
    ///             Http2HealthCheck = new Gcp.Compute.Inputs.RegionHealthCheckHttp2HealthCheckArgs
    ///             {
    ///                 Port = 443,
    ///             },
    ///             TimeoutSec = 1,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ## Example Usage - Region Health Check Http2 Full
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
    ///         var http2_region_health_check = new Gcp.Compute.RegionHealthCheck("http2-region-health-check", new Gcp.Compute.RegionHealthCheckArgs
    ///         {
    ///             CheckIntervalSec = 1,
    ///             Description = "Health check via http2",
    ///             HealthyThreshold = 4,
    ///             Http2HealthCheck = new Gcp.Compute.Inputs.RegionHealthCheckHttp2HealthCheckArgs
    ///             {
    ///                 Host = "1.2.3.4",
    ///                 PortName = "health-check-port",
    ///                 PortSpecification = "USE_NAMED_PORT",
    ///                 ProxyHeader = "NONE",
    ///                 RequestPath = "/mypath",
    ///                 Response = "I AM HEALTHY",
    ///             },
    ///             TimeoutSec = 1,
    ///             UnhealthyThreshold = 5,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    public partial class RegionHealthCheck : Pulumi.CustomResource
    {
        /// <summary>
        /// How often (in seconds) to send a health check. The default value is 5
        /// seconds.
        /// </summary>
        [Output("checkIntervalSec")]
        public Output<int?> CheckIntervalSec { get; private set; } = null!;

        /// <summary>
        /// Creation timestamp in RFC3339 text format.
        /// </summary>
        [Output("creationTimestamp")]
        public Output<string> CreationTimestamp { get; private set; } = null!;

        /// <summary>
        /// An optional description of this resource. Provide this property when
        /// you create the resource.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// A so-far unhealthy instance will be marked healthy after this many
        /// consecutive successes. The default value is 2.
        /// </summary>
        [Output("healthyThreshold")]
        public Output<int?> HealthyThreshold { get; private set; } = null!;

        /// <summary>
        /// A nested object resource  Structure is documented below.
        /// </summary>
        [Output("http2HealthCheck")]
        public Output<Outputs.RegionHealthCheckHttp2HealthCheck?> Http2HealthCheck { get; private set; } = null!;

        /// <summary>
        /// A nested object resource  Structure is documented below.
        /// </summary>
        [Output("httpHealthCheck")]
        public Output<Outputs.RegionHealthCheckHttpHealthCheck?> HttpHealthCheck { get; private set; } = null!;

        /// <summary>
        /// A nested object resource  Structure is documented below.
        /// </summary>
        [Output("httpsHealthCheck")]
        public Output<Outputs.RegionHealthCheckHttpsHealthCheck?> HttpsHealthCheck { get; private set; } = null!;

        /// <summary>
        /// Configure logging on this health check.  Structure is documented below.
        /// </summary>
        [Output("logConfig")]
        public Output<Outputs.RegionHealthCheckLogConfig?> LogConfig { get; private set; } = null!;

        /// <summary>
        /// Name of the resource. Provided by the client when the resource is
        /// created. The name must be 1-63 characters long, and comply with
        /// RFC1035.  Specifically, the name must be 1-63 characters long and
        /// match the regular expression `a-z?` which means
        /// the first character must be a lowercase letter, and all following
        /// characters must be a dash, lowercase letter, or digit, except the
        /// last character, which cannot be a dash.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// The Region in which the created health check should reside.
        /// If it is not provided, the provider region is used.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// The URI of the created resource.
        /// </summary>
        [Output("selfLink")]
        public Output<string> SelfLink { get; private set; } = null!;

        /// <summary>
        /// A nested object resource  Structure is documented below.
        /// </summary>
        [Output("sslHealthCheck")]
        public Output<Outputs.RegionHealthCheckSslHealthCheck?> SslHealthCheck { get; private set; } = null!;

        /// <summary>
        /// A nested object resource  Structure is documented below.
        /// </summary>
        [Output("tcpHealthCheck")]
        public Output<Outputs.RegionHealthCheckTcpHealthCheck?> TcpHealthCheck { get; private set; } = null!;

        /// <summary>
        /// How long (in seconds) to wait before claiming failure.
        /// The default value is 5 seconds.  It is invalid for timeoutSec to have
        /// greater value than checkIntervalSec.
        /// </summary>
        [Output("timeoutSec")]
        public Output<int?> TimeoutSec { get; private set; } = null!;

        /// <summary>
        /// The type of the health check. One of HTTP, HTTP2, HTTPS, TCP, or SSL.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// A so-far healthy instance will be marked unhealthy after this many
        /// consecutive failures. The default value is 2.
        /// </summary>
        [Output("unhealthyThreshold")]
        public Output<int?> UnhealthyThreshold { get; private set; } = null!;


        /// <summary>
        /// Create a RegionHealthCheck resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RegionHealthCheck(string name, RegionHealthCheckArgs? args = null, CustomResourceOptions? options = null)
            : base("gcp:compute/regionHealthCheck:RegionHealthCheck", name, args ?? new RegionHealthCheckArgs(), MakeResourceOptions(options, ""))
        {
        }

        private RegionHealthCheck(string name, Input<string> id, RegionHealthCheckState? state = null, CustomResourceOptions? options = null)
            : base("gcp:compute/regionHealthCheck:RegionHealthCheck", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing RegionHealthCheck resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RegionHealthCheck Get(string name, Input<string> id, RegionHealthCheckState? state = null, CustomResourceOptions? options = null)
        {
            return new RegionHealthCheck(name, id, state, options);
        }
    }

    public sealed class RegionHealthCheckArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// How often (in seconds) to send a health check. The default value is 5
        /// seconds.
        /// </summary>
        [Input("checkIntervalSec")]
        public Input<int>? CheckIntervalSec { get; set; }

        /// <summary>
        /// An optional description of this resource. Provide this property when
        /// you create the resource.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// A so-far unhealthy instance will be marked healthy after this many
        /// consecutive successes. The default value is 2.
        /// </summary>
        [Input("healthyThreshold")]
        public Input<int>? HealthyThreshold { get; set; }

        /// <summary>
        /// A nested object resource  Structure is documented below.
        /// </summary>
        [Input("http2HealthCheck")]
        public Input<Inputs.RegionHealthCheckHttp2HealthCheckArgs>? Http2HealthCheck { get; set; }

        /// <summary>
        /// A nested object resource  Structure is documented below.
        /// </summary>
        [Input("httpHealthCheck")]
        public Input<Inputs.RegionHealthCheckHttpHealthCheckArgs>? HttpHealthCheck { get; set; }

        /// <summary>
        /// A nested object resource  Structure is documented below.
        /// </summary>
        [Input("httpsHealthCheck")]
        public Input<Inputs.RegionHealthCheckHttpsHealthCheckArgs>? HttpsHealthCheck { get; set; }

        /// <summary>
        /// Configure logging on this health check.  Structure is documented below.
        /// </summary>
        [Input("logConfig")]
        public Input<Inputs.RegionHealthCheckLogConfigArgs>? LogConfig { get; set; }

        /// <summary>
        /// Name of the resource. Provided by the client when the resource is
        /// created. The name must be 1-63 characters long, and comply with
        /// RFC1035.  Specifically, the name must be 1-63 characters long and
        /// match the regular expression `a-z?` which means
        /// the first character must be a lowercase letter, and all following
        /// characters must be a dash, lowercase letter, or digit, except the
        /// last character, which cannot be a dash.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The Region in which the created health check should reside.
        /// If it is not provided, the provider region is used.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// A nested object resource  Structure is documented below.
        /// </summary>
        [Input("sslHealthCheck")]
        public Input<Inputs.RegionHealthCheckSslHealthCheckArgs>? SslHealthCheck { get; set; }

        /// <summary>
        /// A nested object resource  Structure is documented below.
        /// </summary>
        [Input("tcpHealthCheck")]
        public Input<Inputs.RegionHealthCheckTcpHealthCheckArgs>? TcpHealthCheck { get; set; }

        /// <summary>
        /// How long (in seconds) to wait before claiming failure.
        /// The default value is 5 seconds.  It is invalid for timeoutSec to have
        /// greater value than checkIntervalSec.
        /// </summary>
        [Input("timeoutSec")]
        public Input<int>? TimeoutSec { get; set; }

        /// <summary>
        /// A so-far healthy instance will be marked unhealthy after this many
        /// consecutive failures. The default value is 2.
        /// </summary>
        [Input("unhealthyThreshold")]
        public Input<int>? UnhealthyThreshold { get; set; }

        public RegionHealthCheckArgs()
        {
        }
    }

    public sealed class RegionHealthCheckState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// How often (in seconds) to send a health check. The default value is 5
        /// seconds.
        /// </summary>
        [Input("checkIntervalSec")]
        public Input<int>? CheckIntervalSec { get; set; }

        /// <summary>
        /// Creation timestamp in RFC3339 text format.
        /// </summary>
        [Input("creationTimestamp")]
        public Input<string>? CreationTimestamp { get; set; }

        /// <summary>
        /// An optional description of this resource. Provide this property when
        /// you create the resource.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// A so-far unhealthy instance will be marked healthy after this many
        /// consecutive successes. The default value is 2.
        /// </summary>
        [Input("healthyThreshold")]
        public Input<int>? HealthyThreshold { get; set; }

        /// <summary>
        /// A nested object resource  Structure is documented below.
        /// </summary>
        [Input("http2HealthCheck")]
        public Input<Inputs.RegionHealthCheckHttp2HealthCheckGetArgs>? Http2HealthCheck { get; set; }

        /// <summary>
        /// A nested object resource  Structure is documented below.
        /// </summary>
        [Input("httpHealthCheck")]
        public Input<Inputs.RegionHealthCheckHttpHealthCheckGetArgs>? HttpHealthCheck { get; set; }

        /// <summary>
        /// A nested object resource  Structure is documented below.
        /// </summary>
        [Input("httpsHealthCheck")]
        public Input<Inputs.RegionHealthCheckHttpsHealthCheckGetArgs>? HttpsHealthCheck { get; set; }

        /// <summary>
        /// Configure logging on this health check.  Structure is documented below.
        /// </summary>
        [Input("logConfig")]
        public Input<Inputs.RegionHealthCheckLogConfigGetArgs>? LogConfig { get; set; }

        /// <summary>
        /// Name of the resource. Provided by the client when the resource is
        /// created. The name must be 1-63 characters long, and comply with
        /// RFC1035.  Specifically, the name must be 1-63 characters long and
        /// match the regular expression `a-z?` which means
        /// the first character must be a lowercase letter, and all following
        /// characters must be a dash, lowercase letter, or digit, except the
        /// last character, which cannot be a dash.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The Region in which the created health check should reside.
        /// If it is not provided, the provider region is used.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The URI of the created resource.
        /// </summary>
        [Input("selfLink")]
        public Input<string>? SelfLink { get; set; }

        /// <summary>
        /// A nested object resource  Structure is documented below.
        /// </summary>
        [Input("sslHealthCheck")]
        public Input<Inputs.RegionHealthCheckSslHealthCheckGetArgs>? SslHealthCheck { get; set; }

        /// <summary>
        /// A nested object resource  Structure is documented below.
        /// </summary>
        [Input("tcpHealthCheck")]
        public Input<Inputs.RegionHealthCheckTcpHealthCheckGetArgs>? TcpHealthCheck { get; set; }

        /// <summary>
        /// How long (in seconds) to wait before claiming failure.
        /// The default value is 5 seconds.  It is invalid for timeoutSec to have
        /// greater value than checkIntervalSec.
        /// </summary>
        [Input("timeoutSec")]
        public Input<int>? TimeoutSec { get; set; }

        /// <summary>
        /// The type of the health check. One of HTTP, HTTP2, HTTPS, TCP, or SSL.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// A so-far healthy instance will be marked unhealthy after this many
        /// consecutive failures. The default value is 2.
        /// </summary>
        [Input("unhealthyThreshold")]
        public Input<int>? UnhealthyThreshold { get; set; }

        public RegionHealthCheckState()
        {
        }
    }
}
