// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.DataCatalog
{
    /// <summary>
    /// A tag template defines a tag, which can have one or more typed fields.
    /// The template is used to create and attach the tag to GCP resources.
    /// 
    /// To get more information about TagTemplate, see:
    /// 
    /// * [API documentation](https://cloud.google.com/data-catalog/docs/reference/rest/v1/projects.locations.tagTemplates)
    /// * How-to Guides
    ///     * [Official Documentation](https://cloud.google.com/data-catalog/docs)
    /// 
    /// ## Example Usage
    /// ### Data Catalog Tag Template Basic
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var basicTagTemplate = new Gcp.DataCatalog.TagTemplate("basicTagTemplate", new Gcp.DataCatalog.TagTemplateArgs
    ///         {
    ///             DisplayName = "Demo Tag Template",
    ///             Fields = 
    ///             {
    ///                 new Gcp.DataCatalog.Inputs.TagTemplateFieldArgs
    ///                 {
    ///                     DisplayName = "Source of data asset",
    ///                     FieldId = "source",
    ///                     IsRequired = true,
    ///                     Type = new Gcp.DataCatalog.Inputs.TagTemplateFieldTypeArgs
    ///                     {
    ///                         PrimitiveType = "STRING",
    ///                     },
    ///                 },
    ///                 new Gcp.DataCatalog.Inputs.TagTemplateFieldArgs
    ///                 {
    ///                     DisplayName = "Number of rows in the data asset",
    ///                     FieldId = "num_rows",
    ///                     Type = new Gcp.DataCatalog.Inputs.TagTemplateFieldTypeArgs
    ///                     {
    ///                         PrimitiveType = "DOUBLE",
    ///                     },
    ///                 },
    ///                 new Gcp.DataCatalog.Inputs.TagTemplateFieldArgs
    ///                 {
    ///                     DisplayName = "PII type",
    ///                     FieldId = "pii_type",
    ///                     Type = new Gcp.DataCatalog.Inputs.TagTemplateFieldTypeArgs
    ///                     {
    ///                         EnumType = new Gcp.DataCatalog.Inputs.TagTemplateFieldTypeEnumTypeArgs
    ///                         {
    ///                             AllowedValues = 
    ///                             {
    ///                                 new Gcp.DataCatalog.Inputs.TagTemplateFieldTypeEnumTypeAllowedValueArgs
    ///                                 {
    ///                                     DisplayName = "EMAIL",
    ///                                 },
    ///                                 new Gcp.DataCatalog.Inputs.TagTemplateFieldTypeEnumTypeAllowedValueArgs
    ///                                 {
    ///                                     DisplayName = "SOCIAL SECURITY NUMBER",
    ///                                 },
    ///                                 new Gcp.DataCatalog.Inputs.TagTemplateFieldTypeEnumTypeAllowedValueArgs
    ///                                 {
    ///                                     DisplayName = "NONE",
    ///                                 },
    ///                             },
    ///                         },
    ///                     },
    ///                 },
    ///             },
    ///             ForceDelete = false,
    ///             Region = "us-central1",
    ///             TagTemplateId = "my_template",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    public partial class TagTemplate : Pulumi.CustomResource
    {
        /// <summary>
        /// The display name for this template.
        /// </summary>
        [Output("displayName")]
        public Output<string?> DisplayName { get; private set; } = null!;

        /// <summary>
        /// Set of tag template field IDs and the settings for the field. This set is an exhaustive list of the allowed fields. This set must contain at least one field and at most 500 fields.  Structure is documented below.
        /// </summary>
        [Output("fields")]
        public Output<ImmutableArray<Outputs.TagTemplateField>> Fields { get; private set; } = null!;

        /// <summary>
        /// This confirms the deletion of any possible tags using this template. Must be set to true in order to delete the tag template.
        /// </summary>
        [Output("forceDelete")]
        public Output<bool?> ForceDelete { get; private set; } = null!;

        /// <summary>
        /// -
        /// The resource name of the tag template field in URL format. Example: projects/{project_id}/locations/{location}/tagTemplates/{tagTemplateId}/fields/{field}
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
        /// Template location region.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// The id of the tag template to create.
        /// </summary>
        [Output("tagTemplateId")]
        public Output<string> TagTemplateId { get; private set; } = null!;


        /// <summary>
        /// Create a TagTemplate resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public TagTemplate(string name, TagTemplateArgs args, CustomResourceOptions? options = null)
            : base("gcp:datacatalog/tagTemplate:TagTemplate", name, args ?? new TagTemplateArgs(), MakeResourceOptions(options, ""))
        {
        }

        private TagTemplate(string name, Input<string> id, TagTemplateState? state = null, CustomResourceOptions? options = null)
            : base("gcp:datacatalog/tagTemplate:TagTemplate", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing TagTemplate resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static TagTemplate Get(string name, Input<string> id, TagTemplateState? state = null, CustomResourceOptions? options = null)
        {
            return new TagTemplate(name, id, state, options);
        }
    }

    public sealed class TagTemplateArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The display name for this template.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        [Input("fields", required: true)]
        private InputList<Inputs.TagTemplateFieldArgs>? _fields;

        /// <summary>
        /// Set of tag template field IDs and the settings for the field. This set is an exhaustive list of the allowed fields. This set must contain at least one field and at most 500 fields.  Structure is documented below.
        /// </summary>
        public InputList<Inputs.TagTemplateFieldArgs> Fields
        {
            get => _fields ?? (_fields = new InputList<Inputs.TagTemplateFieldArgs>());
            set => _fields = value;
        }

        /// <summary>
        /// This confirms the deletion of any possible tags using this template. Must be set to true in order to delete the tag template.
        /// </summary>
        [Input("forceDelete")]
        public Input<bool>? ForceDelete { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// Template location region.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The id of the tag template to create.
        /// </summary>
        [Input("tagTemplateId", required: true)]
        public Input<string> TagTemplateId { get; set; } = null!;

        public TagTemplateArgs()
        {
        }
    }

    public sealed class TagTemplateState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The display name for this template.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        [Input("fields")]
        private InputList<Inputs.TagTemplateFieldGetArgs>? _fields;

        /// <summary>
        /// Set of tag template field IDs and the settings for the field. This set is an exhaustive list of the allowed fields. This set must contain at least one field and at most 500 fields.  Structure is documented below.
        /// </summary>
        public InputList<Inputs.TagTemplateFieldGetArgs> Fields
        {
            get => _fields ?? (_fields = new InputList<Inputs.TagTemplateFieldGetArgs>());
            set => _fields = value;
        }

        /// <summary>
        /// This confirms the deletion of any possible tags using this template. Must be set to true in order to delete the tag template.
        /// </summary>
        [Input("forceDelete")]
        public Input<bool>? ForceDelete { get; set; }

        /// <summary>
        /// -
        /// The resource name of the tag template field in URL format. Example: projects/{project_id}/locations/{location}/tagTemplates/{tagTemplateId}/fields/{field}
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
        /// Template location region.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The id of the tag template to create.
        /// </summary>
        [Input("tagTemplateId")]
        public Input<string>? TagTemplateId { get; set; }

        public TagTemplateState()
        {
        }
    }
}
