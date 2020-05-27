// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.ContainerAnalysis
{
    /// <summary>
    /// Provides a detailed description of a Note.
    /// 
    /// 
    /// To get more information about Note, see:
    /// 
    /// * [API documentation](https://cloud.google.com/container-analysis/api/reference/rest/)
    /// * How-to Guides
    ///     * [Official Documentation](https://cloud.google.com/container-analysis/)
    /// 
    /// ## Example Usage - Container Analysis Note Basic
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
    ///         var note = new Gcp.ContainerAnalysis.Note("note", new Gcp.ContainerAnalysis.NoteArgs
    ///         {
    ///             AttestationAuthority = new Gcp.ContainerAnalysis.Inputs.NoteAttestationAuthorityArgs
    ///             {
    ///                 Hint = new Gcp.ContainerAnalysis.Inputs.NoteAttestationAuthorityHintArgs
    ///                 {
    ///                     HumanReadableName = "Attestor Note",
    ///                 },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// {{% /example %}}
    /// </summary>
    public partial class Note : Pulumi.CustomResource
    {
        /// <summary>
        /// Note kind that represents a logical attestation "role" or "authority".
        /// For example, an organization might have one AttestationAuthority for
        /// "QA" and one for "build". This Note is intended to act strictly as a
        /// grouping mechanism for the attached Occurrences (Attestations). This
        /// grouping mechanism also provides a security boundary, since IAM ACLs
        /// gate the ability for a principle to attach an Occurrence to a given
        /// Note. It also provides a single point of lookup to find all attached
        /// Attestation Occurrences, even if they don't all live in the same
        /// project.  Structure is documented below.
        /// </summary>
        [Output("attestationAuthority")]
        public Output<Outputs.NoteAttestationAuthority> AttestationAuthority { get; private set; } = null!;

        /// <summary>
        /// The name of the note.
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
        /// Create a Note resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Note(string name, NoteArgs args, CustomResourceOptions? options = null)
            : base("gcp:containeranalysis/note:Note", name, args ?? new NoteArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Note(string name, Input<string> id, NoteState? state = null, CustomResourceOptions? options = null)
            : base("gcp:containeranalysis/note:Note", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Note resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Note Get(string name, Input<string> id, NoteState? state = null, CustomResourceOptions? options = null)
        {
            return new Note(name, id, state, options);
        }
    }

    public sealed class NoteArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Note kind that represents a logical attestation "role" or "authority".
        /// For example, an organization might have one AttestationAuthority for
        /// "QA" and one for "build". This Note is intended to act strictly as a
        /// grouping mechanism for the attached Occurrences (Attestations). This
        /// grouping mechanism also provides a security boundary, since IAM ACLs
        /// gate the ability for a principle to attach an Occurrence to a given
        /// Note. It also provides a single point of lookup to find all attached
        /// Attestation Occurrences, even if they don't all live in the same
        /// project.  Structure is documented below.
        /// </summary>
        [Input("attestationAuthority", required: true)]
        public Input<Inputs.NoteAttestationAuthorityArgs> AttestationAuthority { get; set; } = null!;

        /// <summary>
        /// The name of the note.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        public NoteArgs()
        {
        }
    }

    public sealed class NoteState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Note kind that represents a logical attestation "role" or "authority".
        /// For example, an organization might have one AttestationAuthority for
        /// "QA" and one for "build". This Note is intended to act strictly as a
        /// grouping mechanism for the attached Occurrences (Attestations). This
        /// grouping mechanism also provides a security boundary, since IAM ACLs
        /// gate the ability for a principle to attach an Occurrence to a given
        /// Note. It also provides a single point of lookup to find all attached
        /// Attestation Occurrences, even if they don't all live in the same
        /// project.  Structure is documented below.
        /// </summary>
        [Input("attestationAuthority")]
        public Input<Inputs.NoteAttestationAuthorityGetArgs>? AttestationAuthority { get; set; }

        /// <summary>
        /// The name of the note.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        public NoteState()
        {
        }
    }
}
