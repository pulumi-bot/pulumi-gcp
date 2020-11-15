// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.RuntimeConfig
{
    /// <summary>
    /// Manages a RuntimeConfig resource in Google Cloud. For more information, see the
    /// [official documentation](https://cloud.google.com/deployment-manager/runtime-configurator/),
    /// or the
    /// [JSON API](https://cloud.google.com/deployment-manager/runtime-configurator/reference/rest/).
    /// 
    /// ## Import
    /// 
    /// Runtime Configs can be imported using the `name` or full config name, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import gcp:runtimeconfig/config:Config myconfig myconfig
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:runtimeconfig/config:Config myconfig projects/my-gcp-project/configs/myconfig
    /// ```
    /// 
    ///  When importing using only the name, the provider project must be set.
    /// </summary>
    public partial class Config : Pulumi.CustomResource
    {
        /// <summary>
        /// The description to associate with the runtime
        /// config.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The name of the runtime config.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs. If it
        /// is not provided, the provider project is used.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;


        /// <summary>
        /// Create a Config resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Config(string name, ConfigArgs? args = null, CustomResourceOptions? options = null)
            : base("gcp:runtimeconfig/config:Config", name, args ?? new ConfigArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Config(string name, Input<string> id, ConfigState? state = null, CustomResourceOptions? options = null)
            : base("gcp:runtimeconfig/config:Config", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Config resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Config Get(string name, Input<string> id, ConfigState? state = null, CustomResourceOptions? options = null)
        {
            return new Config(name, id, state, options);
        }
    }

    public sealed class ConfigArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description to associate with the runtime
        /// config.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The name of the runtime config.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs. If it
        /// is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        public ConfigArgs()
        {
        }
    }

    public sealed class ConfigState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description to associate with the runtime
        /// config.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The name of the runtime config.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs. If it
        /// is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        public ConfigState()
        {
        }
    }
}
