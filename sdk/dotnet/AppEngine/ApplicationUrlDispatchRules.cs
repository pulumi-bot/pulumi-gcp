// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.AppEngine
{
    /// <summary>
    /// Rules to match an HTTP request and dispatch that request to a service.
    /// 
    /// To get more information about ApplicationUrlDispatchRules, see:
    /// 
    /// * [API documentation](https://cloud.google.com/appengine/docs/admin-api/reference/rest/v1/apps#UrlDispatchRule)
    /// 
    /// ## Example Usage
    /// </summary>
    public partial class ApplicationUrlDispatchRules : Pulumi.CustomResource
    {
        /// <summary>
        /// Rules to match an HTTP request and dispatch that request to a service.  Structure is documented below.
        /// </summary>
        [Output("dispatchRules")]
        public Output<ImmutableArray<Outputs.ApplicationUrlDispatchRulesDispatchRule>> DispatchRules { get; private set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;


        /// <summary>
        /// Create a ApplicationUrlDispatchRules resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ApplicationUrlDispatchRules(string name, ApplicationUrlDispatchRulesArgs args, CustomResourceOptions? options = null)
            : base("gcp:appengine/applicationUrlDispatchRules:ApplicationUrlDispatchRules", name, args ?? new ApplicationUrlDispatchRulesArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ApplicationUrlDispatchRules(string name, Input<string> id, ApplicationUrlDispatchRulesState? state = null, CustomResourceOptions? options = null)
            : base("gcp:appengine/applicationUrlDispatchRules:ApplicationUrlDispatchRules", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ApplicationUrlDispatchRules resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ApplicationUrlDispatchRules Get(string name, Input<string> id, ApplicationUrlDispatchRulesState? state = null, CustomResourceOptions? options = null)
        {
            return new ApplicationUrlDispatchRules(name, id, state, options);
        }
    }

    public sealed class ApplicationUrlDispatchRulesArgs : Pulumi.ResourceArgs
    {
        [Input("dispatchRules", required: true)]
        private InputList<Inputs.ApplicationUrlDispatchRulesDispatchRuleArgs>? _dispatchRules;

        /// <summary>
        /// Rules to match an HTTP request and dispatch that request to a service.  Structure is documented below.
        /// </summary>
        public InputList<Inputs.ApplicationUrlDispatchRulesDispatchRuleArgs> DispatchRules
        {
            get => _dispatchRules ?? (_dispatchRules = new InputList<Inputs.ApplicationUrlDispatchRulesDispatchRuleArgs>());
            set => _dispatchRules = value;
        }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        public ApplicationUrlDispatchRulesArgs()
        {
        }
    }

    public sealed class ApplicationUrlDispatchRulesState : Pulumi.ResourceArgs
    {
        [Input("dispatchRules")]
        private InputList<Inputs.ApplicationUrlDispatchRulesDispatchRuleGetArgs>? _dispatchRules;

        /// <summary>
        /// Rules to match an HTTP request and dispatch that request to a service.  Structure is documented below.
        /// </summary>
        public InputList<Inputs.ApplicationUrlDispatchRulesDispatchRuleGetArgs> DispatchRules
        {
            get => _dispatchRules ?? (_dispatchRules = new InputList<Inputs.ApplicationUrlDispatchRulesDispatchRuleGetArgs>());
            set => _dispatchRules = value;
        }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        public ApplicationUrlDispatchRulesState()
        {
        }
    }
}
