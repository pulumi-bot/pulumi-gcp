// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Configuration for an automated build in response to source repository changes.
 *
 * To get more information about Trigger, see:
 *
 * * [API documentation](https://cloud.google.com/cloud-build/docs/api/reference/rest/)
 * * How-to Guides
 *     * [Automating builds using build triggers](https://cloud.google.com/cloud-build/docs/running-builds/automate-builds)
 *
 * ## Example Usage
 */
export class Trigger extends pulumi.CustomResource {
    /**
     * Get an existing Trigger resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: TriggerState, opts?: pulumi.CustomResourceOptions): Trigger {
        return new Trigger(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:cloudbuild/trigger:Trigger';

    /**
     * Returns true if the given object is an instance of Trigger.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Trigger {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Trigger.__pulumiType;
    }

    /**
     * Contents of the build template. Either a filename or build template must be provided.
     * Structure is documented below.
     */
    public readonly build!: pulumi.Output<outputs.cloudbuild.TriggerBuild | undefined>;
    /**
     * Time when the trigger was created.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * Human-readable description of the trigger.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Whether the trigger is disabled or not. If true, the trigger will never result in a build.
     */
    public readonly disabled!: pulumi.Output<boolean | undefined>;
    /**
     * Path, from the source root, to a file whose contents is used for the template. Either a filename or build template must be provided.
     */
    public readonly filename!: pulumi.Output<string | undefined>;
    /**
     * Describes the configuration of a trigger that creates a build whenever a GitHub event is received.
     * One of `triggerTemplate` or `github` must be provided.
     * Structure is documented below.
     */
    public readonly github!: pulumi.Output<outputs.cloudbuild.TriggerGithub | undefined>;
    /**
     * ignoredFiles and includedFiles are file glob matches using https://golang.org/pkg/path/filepath/#Match
     * extended with support for `**`.
     * If ignoredFiles and changed files are both empty, then they are not
     * used to determine whether or not to trigger a build.
     * If ignoredFiles is not empty, then we ignore any files that match any
     * of the ignoredFile globs. If the change has no files that are outside
     * of the ignoredFiles globs, then we do not trigger a build.
     */
    public readonly ignoredFiles!: pulumi.Output<string[] | undefined>;
    /**
     * ignoredFiles and includedFiles are file glob matches using https://golang.org/pkg/path/filepath/#Match
     * extended with support for `**`.
     * If any of the files altered in the commit pass the ignoredFiles filter
     * and includedFiles is empty, then as far as this filter is concerned, we
     * should trigger the build.
     * If any of the files altered in the commit pass the ignoredFiles filter
     * and includedFiles is not empty, then we make sure that at least one of
     * those files matches a includedFiles glob. If not, then we do not trigger
     * a build.
     */
    public readonly includedFiles!: pulumi.Output<string[] | undefined>;
    /**
     * Name of the volume to mount.
     * Volume names must be unique per build step and must be valid names for Docker volumes.
     * Each named volume must be used by at least two build steps.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * Substitutions to use in a triggered build. Should only be used with triggers.run
     */
    public readonly substitutions!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Tags for annotation of a Build. These are not docker tags.
     */
    public readonly tags!: pulumi.Output<string[] | undefined>;
    /**
     * The unique identifier for the trigger.
     */
    public /*out*/ readonly triggerId!: pulumi.Output<string>;
    /**
     * Template describing the types of source changes to trigger a build.
     * Branch and tag names in trigger templates are interpreted as regular
     * expressions. Any branch or tag change that matches that regular
     * expression will trigger a build.
     * One of `triggerTemplate` or `github` must be provided.
     * Structure is documented below.
     */
    public readonly triggerTemplate!: pulumi.Output<outputs.cloudbuild.TriggerTriggerTemplate | undefined>;

    /**
     * Create a Trigger resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: TriggerArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: TriggerArgs | TriggerState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as TriggerState | undefined;
            inputs["build"] = state ? state.build : undefined;
            inputs["createTime"] = state ? state.createTime : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["disabled"] = state ? state.disabled : undefined;
            inputs["filename"] = state ? state.filename : undefined;
            inputs["github"] = state ? state.github : undefined;
            inputs["ignoredFiles"] = state ? state.ignoredFiles : undefined;
            inputs["includedFiles"] = state ? state.includedFiles : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["substitutions"] = state ? state.substitutions : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["triggerId"] = state ? state.triggerId : undefined;
            inputs["triggerTemplate"] = state ? state.triggerTemplate : undefined;
        } else {
            const args = argsOrState as TriggerArgs | undefined;
            inputs["build"] = args ? args.build : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["disabled"] = args ? args.disabled : undefined;
            inputs["filename"] = args ? args.filename : undefined;
            inputs["github"] = args ? args.github : undefined;
            inputs["ignoredFiles"] = args ? args.ignoredFiles : undefined;
            inputs["includedFiles"] = args ? args.includedFiles : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["substitutions"] = args ? args.substitutions : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["triggerTemplate"] = args ? args.triggerTemplate : undefined;
            inputs["createTime"] = undefined /*out*/;
            inputs["triggerId"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Trigger.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Trigger resources.
 */
export interface TriggerState {
    /**
     * Contents of the build template. Either a filename or build template must be provided.
     * Structure is documented below.
     */
    readonly build?: pulumi.Input<inputs.cloudbuild.TriggerBuild>;
    /**
     * Time when the trigger was created.
     */
    readonly createTime?: pulumi.Input<string>;
    /**
     * Human-readable description of the trigger.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Whether the trigger is disabled or not. If true, the trigger will never result in a build.
     */
    readonly disabled?: pulumi.Input<boolean>;
    /**
     * Path, from the source root, to a file whose contents is used for the template. Either a filename or build template must be provided.
     */
    readonly filename?: pulumi.Input<string>;
    /**
     * Describes the configuration of a trigger that creates a build whenever a GitHub event is received.
     * One of `triggerTemplate` or `github` must be provided.
     * Structure is documented below.
     */
    readonly github?: pulumi.Input<inputs.cloudbuild.TriggerGithub>;
    /**
     * ignoredFiles and includedFiles are file glob matches using https://golang.org/pkg/path/filepath/#Match
     * extended with support for `**`.
     * If ignoredFiles and changed files are both empty, then they are not
     * used to determine whether or not to trigger a build.
     * If ignoredFiles is not empty, then we ignore any files that match any
     * of the ignoredFile globs. If the change has no files that are outside
     * of the ignoredFiles globs, then we do not trigger a build.
     */
    readonly ignoredFiles?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * ignoredFiles and includedFiles are file glob matches using https://golang.org/pkg/path/filepath/#Match
     * extended with support for `**`.
     * If any of the files altered in the commit pass the ignoredFiles filter
     * and includedFiles is empty, then as far as this filter is concerned, we
     * should trigger the build.
     * If any of the files altered in the commit pass the ignoredFiles filter
     * and includedFiles is not empty, then we make sure that at least one of
     * those files matches a includedFiles glob. If not, then we do not trigger
     * a build.
     */
    readonly includedFiles?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Name of the volume to mount.
     * Volume names must be unique per build step and must be valid names for Docker volumes.
     * Each named volume must be used by at least two build steps.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * Substitutions to use in a triggered build. Should only be used with triggers.run
     */
    readonly substitutions?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Tags for annotation of a Build. These are not docker tags.
     */
    readonly tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The unique identifier for the trigger.
     */
    readonly triggerId?: pulumi.Input<string>;
    /**
     * Template describing the types of source changes to trigger a build.
     * Branch and tag names in trigger templates are interpreted as regular
     * expressions. Any branch or tag change that matches that regular
     * expression will trigger a build.
     * One of `triggerTemplate` or `github` must be provided.
     * Structure is documented below.
     */
    readonly triggerTemplate?: pulumi.Input<inputs.cloudbuild.TriggerTriggerTemplate>;
}

/**
 * The set of arguments for constructing a Trigger resource.
 */
export interface TriggerArgs {
    /**
     * Contents of the build template. Either a filename or build template must be provided.
     * Structure is documented below.
     */
    readonly build?: pulumi.Input<inputs.cloudbuild.TriggerBuild>;
    /**
     * Human-readable description of the trigger.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Whether the trigger is disabled or not. If true, the trigger will never result in a build.
     */
    readonly disabled?: pulumi.Input<boolean>;
    /**
     * Path, from the source root, to a file whose contents is used for the template. Either a filename or build template must be provided.
     */
    readonly filename?: pulumi.Input<string>;
    /**
     * Describes the configuration of a trigger that creates a build whenever a GitHub event is received.
     * One of `triggerTemplate` or `github` must be provided.
     * Structure is documented below.
     */
    readonly github?: pulumi.Input<inputs.cloudbuild.TriggerGithub>;
    /**
     * ignoredFiles and includedFiles are file glob matches using https://golang.org/pkg/path/filepath/#Match
     * extended with support for `**`.
     * If ignoredFiles and changed files are both empty, then they are not
     * used to determine whether or not to trigger a build.
     * If ignoredFiles is not empty, then we ignore any files that match any
     * of the ignoredFile globs. If the change has no files that are outside
     * of the ignoredFiles globs, then we do not trigger a build.
     */
    readonly ignoredFiles?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * ignoredFiles and includedFiles are file glob matches using https://golang.org/pkg/path/filepath/#Match
     * extended with support for `**`.
     * If any of the files altered in the commit pass the ignoredFiles filter
     * and includedFiles is empty, then as far as this filter is concerned, we
     * should trigger the build.
     * If any of the files altered in the commit pass the ignoredFiles filter
     * and includedFiles is not empty, then we make sure that at least one of
     * those files matches a includedFiles glob. If not, then we do not trigger
     * a build.
     */
    readonly includedFiles?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Name of the volume to mount.
     * Volume names must be unique per build step and must be valid names for Docker volumes.
     * Each named volume must be used by at least two build steps.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * Substitutions to use in a triggered build. Should only be used with triggers.run
     */
    readonly substitutions?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Tags for annotation of a Build. These are not docker tags.
     */
    readonly tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Template describing the types of source changes to trigger a build.
     * Branch and tag names in trigger templates are interpreted as regular
     * expressions. Any branch or tag change that matches that regular
     * expression will trigger a build.
     * One of `triggerTemplate` or `github` must be provided.
     * Structure is documented below.
     */
    readonly triggerTemplate?: pulumi.Input<inputs.cloudbuild.TriggerTriggerTemplate>;
}
