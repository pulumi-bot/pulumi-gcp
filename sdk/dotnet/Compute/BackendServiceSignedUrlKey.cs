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
    /// A key for signing Cloud CDN signed URLs for Backend Services.
    /// 
    /// To get more information about BackendServiceSignedUrlKey, see:
    /// 
    /// * [API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/backendServices)
    /// * How-to Guides
    ///     * [Using Signed URLs](https://cloud.google.com/cdn/docs/using-signed-urls/)
    /// 
    /// &gt; **Warning:** All arguments including `key_value` will be stored in the raw
    /// state as plain-text.
    /// 
    /// ## Example Usage
    /// ### Backend Service Signed Url Key
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// using Random = Pulumi.Random;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var urlSignature = new Random.RandomId("urlSignature", new Random.RandomIdArgs
    ///         {
    ///             ByteLength = 16,
    ///         });
    ///         var webserver = new Gcp.Compute.InstanceTemplate("webserver", new Gcp.Compute.InstanceTemplateArgs
    ///         {
    ///             MachineType = "e2-medium",
    ///             NetworkInterfaces = 
    ///             {
    ///                 new Gcp.Compute.Inputs.InstanceTemplateNetworkInterfaceArgs
    ///                 {
    ///                     Network = "default",
    ///                 },
    ///             },
    ///             Disks = 
    ///             {
    ///                 new Gcp.Compute.Inputs.InstanceTemplateDiskArgs
    ///                 {
    ///                     SourceImage = "debian-cloud/debian-9",
    ///                     AutoDelete = true,
    ///                     Boot = true,
    ///                 },
    ///             },
    ///         });
    ///         var webservers = new Gcp.Compute.InstanceGroupManager("webservers", new Gcp.Compute.InstanceGroupManagerArgs
    ///         {
    ///             Versions = 
    ///             {
    ///                 new Gcp.Compute.Inputs.InstanceGroupManagerVersionArgs
    ///                 {
    ///                     InstanceTemplate = webserver.Id,
    ///                     Name = "primary",
    ///                 },
    ///             },
    ///             BaseInstanceName = "webserver",
    ///             Zone = "us-central1-f",
    ///             TargetSize = 1,
    ///         });
    ///         var @default = new Gcp.Compute.HttpHealthCheck("default", new Gcp.Compute.HttpHealthCheckArgs
    ///         {
    ///             RequestPath = "/",
    ///             CheckIntervalSec = 1,
    ///             TimeoutSec = 1,
    ///         });
    ///         var exampleBackend = new Gcp.Compute.BackendService("exampleBackend", new Gcp.Compute.BackendServiceArgs
    ///         {
    ///             Description = "Our company website",
    ///             PortName = "http",
    ///             Protocol = "HTTP",
    ///             TimeoutSec = 10,
    ///             EnableCdn = true,
    ///             Backends = 
    ///             {
    ///                 new Gcp.Compute.Inputs.BackendServiceBackendArgs
    ///                 {
    ///                     Group = webservers.InstanceGroup,
    ///                 },
    ///             },
    ///             HealthChecks = 
    ///             {
    ///                 @default.Id,
    ///             },
    ///         });
    ///         var backendKey = new Gcp.Compute.BackendServiceSignedUrlKey("backendKey", new Gcp.Compute.BackendServiceSignedUrlKeyArgs
    ///         {
    ///             KeyValue = urlSignature.B64Url,
    ///             BackendService = exampleBackend.Name,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// This resource does not support import.
    /// </summary>
    [GcpResourceType("gcp:compute/backendServiceSignedUrlKey:BackendServiceSignedUrlKey")]
    public partial class BackendServiceSignedUrlKey : Pulumi.CustomResource
    {
        /// <summary>
        /// The backend service this signed URL key belongs.
        /// </summary>
        [Output("backendService")]
        public Output<string> BackendService { get; private set; } = null!;

        /// <summary>
        /// 128-bit key value used for signing the URL. The key value must be a
        /// valid RFC 4648 Section 5 base64url encoded string.
        /// **Note**: This property is sensitive and will not be displayed in the plan.
        /// </summary>
        [Output("keyValue")]
        public Output<string> KeyValue { get; private set; } = null!;

        /// <summary>
        /// Name of the signed URL key.
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
        /// Create a BackendServiceSignedUrlKey resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public BackendServiceSignedUrlKey(string name, BackendServiceSignedUrlKeyArgs args, CustomResourceOptions? options = null)
            : base("gcp:compute/backendServiceSignedUrlKey:BackendServiceSignedUrlKey", name, args ?? new BackendServiceSignedUrlKeyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private BackendServiceSignedUrlKey(string name, Input<string> id, BackendServiceSignedUrlKeyState? state = null, CustomResourceOptions? options = null)
            : base("gcp:compute/backendServiceSignedUrlKey:BackendServiceSignedUrlKey", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing BackendServiceSignedUrlKey resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static BackendServiceSignedUrlKey Get(string name, Input<string> id, BackendServiceSignedUrlKeyState? state = null, CustomResourceOptions? options = null)
        {
            return new BackendServiceSignedUrlKey(name, id, state, options);
        }
    }

    public sealed class BackendServiceSignedUrlKeyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The backend service this signed URL key belongs.
        /// </summary>
        [Input("backendService", required: true)]
        public Input<string> BackendService { get; set; } = null!;

        /// <summary>
        /// 128-bit key value used for signing the URL. The key value must be a
        /// valid RFC 4648 Section 5 base64url encoded string.
        /// **Note**: This property is sensitive and will not be displayed in the plan.
        /// </summary>
        [Input("keyValue", required: true)]
        public Input<string> KeyValue { get; set; } = null!;

        /// <summary>
        /// Name of the signed URL key.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        public BackendServiceSignedUrlKeyArgs()
        {
        }
    }

    public sealed class BackendServiceSignedUrlKeyState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The backend service this signed URL key belongs.
        /// </summary>
        [Input("backendService")]
        public Input<string>? BackendService { get; set; }

        /// <summary>
        /// 128-bit key value used for signing the URL. The key value must be a
        /// valid RFC 4648 Section 5 base64url encoded string.
        /// **Note**: This property is sensitive and will not be displayed in the plan.
        /// </summary>
        [Input("keyValue")]
        public Input<string>? KeyValue { get; set; }

        /// <summary>
        /// Name of the signed URL key.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        public BackendServiceSignedUrlKeyState()
        {
        }
    }
}
