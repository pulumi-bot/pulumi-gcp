// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.GameServices
{
    /// <summary>
    /// A Realm resource.
    /// 
    /// To get more information about Realm, see:
    /// 
    /// * [API documentation](https://cloud.google.com/game-servers/docs/reference/rest/v1beta/projects.locations.realms)
    /// * How-to Guides
    ///     * [Official Documentation](https://cloud.google.com/game-servers/docs)
    /// 
    /// ## Example Usage
    /// ### Game Service Realm Basic
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var @default = new Gcp.GameServices.Realm("default", new Gcp.GameServices.RealmArgs
    ///         {
    ///             RealmId = "tf-test-realm",
    ///             TimeZone = "EST",
    ///             Location = "global",
    ///             Description = "one of the nine",
    ///         }, new CustomResourceOptions {
    ///             Provider = google_beta,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    public partial class Realm : Pulumi.CustomResource
    {
        /// <summary>
        /// Human readable description of the realm.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// ETag of the resource.
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        /// <summary>
        /// The labels associated with this realm. Each label is a key-value pair.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>?> Labels { get; private set; } = null!;

        /// <summary>
        /// Location of the Realm.
        /// </summary>
        [Output("location")]
        public Output<string?> Location { get; private set; } = null!;

        /// <summary>
        /// The resource id of the realm, of the form: 'projects/{project_id}/locations/{location}/realms/{realm_id}'. For example,
        /// 'projects/my-project/locations/{location}/realms/my-realm'.
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
        /// GCP region of the Realm.
        /// </summary>
        [Output("realmId")]
        public Output<string> RealmId { get; private set; } = null!;

        /// <summary>
        /// Required. Time zone where all realm-specific policies are evaluated. The value of
        /// this field must be from the IANA time zone database:
        /// https://www.iana.org/time-zones.
        /// </summary>
        [Output("timeZone")]
        public Output<string> TimeZone { get; private set; } = null!;


        /// <summary>
        /// Create a Realm resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Realm(string name, RealmArgs args, CustomResourceOptions? options = null)
            : base("gcp:gameservices/realm:Realm", name, args ?? new RealmArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Realm(string name, Input<string> id, RealmState? state = null, CustomResourceOptions? options = null)
            : base("gcp:gameservices/realm:Realm", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Realm resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Realm Get(string name, Input<string> id, RealmState? state = null, CustomResourceOptions? options = null)
        {
            return new Realm(name, id, state, options);
        }
    }

    public sealed class RealmArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Human readable description of the realm.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// The labels associated with this realm. Each label is a key-value pair.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// Location of the Realm.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// GCP region of the Realm.
        /// </summary>
        [Input("realmId", required: true)]
        public Input<string> RealmId { get; set; } = null!;

        /// <summary>
        /// Required. Time zone where all realm-specific policies are evaluated. The value of
        /// this field must be from the IANA time zone database:
        /// https://www.iana.org/time-zones.
        /// </summary>
        [Input("timeZone", required: true)]
        public Input<string> TimeZone { get; set; } = null!;

        public RealmArgs()
        {
        }
    }

    public sealed class RealmState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Human readable description of the realm.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// ETag of the resource.
        /// </summary>
        [Input("etag")]
        public Input<string>? Etag { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// The labels associated with this realm. Each label is a key-value pair.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// Location of the Realm.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The resource id of the realm, of the form: 'projects/{project_id}/locations/{location}/realms/{realm_id}'. For example,
        /// 'projects/my-project/locations/{location}/realms/my-realm'.
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
        /// GCP region of the Realm.
        /// </summary>
        [Input("realmId")]
        public Input<string>? RealmId { get; set; }

        /// <summary>
        /// Required. Time zone where all realm-specific policies are evaluated. The value of
        /// this field must be from the IANA time zone database:
        /// https://www.iana.org/time-zones.
        /// </summary>
        [Input("timeZone")]
        public Input<string>? TimeZone { get; set; }

        public RealmState()
        {
        }
    }
}
