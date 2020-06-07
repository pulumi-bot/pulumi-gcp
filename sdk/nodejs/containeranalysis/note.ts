// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides a detailed description of a Note.
 *
 *
 * To get more information about Note, see:
 *
 * * [API documentation](https://cloud.google.com/container-analysis/api/reference/rest/)
 * * How-to Guides
 *     * [Official Documentation](https://cloud.google.com/container-analysis/)
 *
 * ## Example Usage - Container Analysis Note Basic
 *
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const note = new gcp.containeranalysis.Note("note", {
 *     attestationAuthority: {
 *         hint: {
 *             humanReadableName: "Attestor Note",
 *         },
 *     },
 * });
 * ```
 */
export class Note extends pulumi.CustomResource {
    /**
     * Get an existing Note resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: NoteState, opts?: pulumi.CustomResourceOptions): Note {
        return new Note(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:containeranalysis/note:Note';

    /**
     * Returns true if the given object is an instance of Note.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Note {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Note.__pulumiType;
    }

    /**
     * Note kind that represents a logical attestation "role" or "authority".
     * For example, an organization might have one AttestationAuthority for
     * "QA" and one for "build". This Note is intended to act strictly as a
     * grouping mechanism for the attached Occurrences (Attestations). This
     * grouping mechanism also provides a security boundary, since IAM ACLs
     * gate the ability for a principle to attach an Occurrence to a given
     * Note. It also provides a single point of lookup to find all attached
     * Attestation Occurrences, even if they don't all live in the same
     * project.  Structure is documented below.
     */
    public readonly attestationAuthority!: pulumi.Output<outputs.containeranalysis.NoteAttestationAuthority>;
    /**
     * The name of the note.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;

    /**
     * Create a Note resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NoteArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: NoteArgs | NoteState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as NoteState | undefined;
            inputs["attestationAuthority"] = state ? state.attestationAuthority : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["project"] = state ? state.project : undefined;
        } else {
            const args = argsOrState as NoteArgs | undefined;
            if (!args || args.attestationAuthority === undefined) {
                throw new Error("Missing required property 'attestationAuthority'");
            }
            inputs["attestationAuthority"] = args ? args.attestationAuthority : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["project"] = args ? args.project : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Note.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Note resources.
 */
export interface NoteState {
    /**
     * Note kind that represents a logical attestation "role" or "authority".
     * For example, an organization might have one AttestationAuthority for
     * "QA" and one for "build". This Note is intended to act strictly as a
     * grouping mechanism for the attached Occurrences (Attestations). This
     * grouping mechanism also provides a security boundary, since IAM ACLs
     * gate the ability for a principle to attach an Occurrence to a given
     * Note. It also provides a single point of lookup to find all attached
     * Attestation Occurrences, even if they don't all live in the same
     * project.  Structure is documented below.
     */
    readonly attestationAuthority?: pulumi.Input<inputs.containeranalysis.NoteAttestationAuthority>;
    /**
     * The name of the note.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Note resource.
 */
export interface NoteArgs {
    /**
     * Note kind that represents a logical attestation "role" or "authority".
     * For example, an organization might have one AttestationAuthority for
     * "QA" and one for "build". This Note is intended to act strictly as a
     * grouping mechanism for the attached Occurrences (Attestations). This
     * grouping mechanism also provides a security boundary, since IAM ACLs
     * gate the ability for a principle to attach an Occurrence to a given
     * Note. It also provides a single point of lookup to find all attached
     * Attestation Occurrences, even if they don't all live in the same
     * project.  Structure is documented below.
     */
    readonly attestationAuthority: pulumi.Input<inputs.containeranalysis.NoteAttestationAuthority>;
    /**
     * The name of the note.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
}
