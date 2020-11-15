// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * A policy for container image binary authorization.
 *
 * To get more information about Policy, see:
 *
 * * [API documentation](https://cloud.google.com/binary-authorization/docs/reference/rest/)
 * * How-to Guides
 *     * [Official Documentation](https://cloud.google.com/binary-authorization/)
 *
 * ## Example Usage
 *
 * ## Import
 *
 * Policy can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import gcp:binaryauthorization/policy:Policy default projects/{{project}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:binaryauthorization/policy:Policy default {{project}}
 * ```
 */
export class Policy extends pulumi.CustomResource {
    /**
     * Get an existing Policy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PolicyState, opts?: pulumi.CustomResourceOptions): Policy {
        return new Policy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:binaryauthorization/policy:Policy';

    /**
     * Returns true if the given object is an instance of Policy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Policy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Policy.__pulumiType;
    }

    /**
     * A whitelist of image patterns to exclude from admission rules. If an
     * image's name matches a whitelist pattern, the image's admission
     * requests will always be permitted regardless of your admission rules.
     * Structure is documented below.
     */
    public readonly admissionWhitelistPatterns!: pulumi.Output<outputs.binaryauthorization.PolicyAdmissionWhitelistPattern[] | undefined>;
    /**
     * Per-cluster admission rules. An admission rule specifies either that
     * all container images used in a pod creation request must be attested
     * to by one or more attestors, that all pod creations will be allowed,
     * or that all pod creations will be denied. There can be at most one
     * admission rule per cluster spec.
     */
    public readonly clusterAdmissionRules!: pulumi.Output<outputs.binaryauthorization.PolicyClusterAdmissionRule[] | undefined>;
    /**
     * Default admission rule for a cluster without a per-cluster admission
     * rule.
     * Structure is documented below.
     */
    public readonly defaultAdmissionRule!: pulumi.Output<outputs.binaryauthorization.PolicyDefaultAdmissionRule>;
    /**
     * A descriptive comment.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Controls the evaluation of a Google-maintained global admission policy
     * for common system-level images. Images not covered by the global
     * policy will be subject to the project admission policy.
     * Possible values are `ENABLE` and `DISABLE`.
     */
    public readonly globalPolicyEvaluationMode!: pulumi.Output<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;

    /**
     * Create a Policy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PolicyArgs | PolicyState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as PolicyState | undefined;
            inputs["admissionWhitelistPatterns"] = state ? state.admissionWhitelistPatterns : undefined;
            inputs["clusterAdmissionRules"] = state ? state.clusterAdmissionRules : undefined;
            inputs["defaultAdmissionRule"] = state ? state.defaultAdmissionRule : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["globalPolicyEvaluationMode"] = state ? state.globalPolicyEvaluationMode : undefined;
            inputs["project"] = state ? state.project : undefined;
        } else {
            const args = argsOrState as PolicyArgs | undefined;
            if (!args || args.defaultAdmissionRule === undefined) {
                throw new Error("Missing required property 'defaultAdmissionRule'");
            }
            inputs["admissionWhitelistPatterns"] = args ? args.admissionWhitelistPatterns : undefined;
            inputs["clusterAdmissionRules"] = args ? args.clusterAdmissionRules : undefined;
            inputs["defaultAdmissionRule"] = args ? args.defaultAdmissionRule : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["globalPolicyEvaluationMode"] = args ? args.globalPolicyEvaluationMode : undefined;
            inputs["project"] = args ? args.project : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Policy.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Policy resources.
 */
export interface PolicyState {
    /**
     * A whitelist of image patterns to exclude from admission rules. If an
     * image's name matches a whitelist pattern, the image's admission
     * requests will always be permitted regardless of your admission rules.
     * Structure is documented below.
     */
    readonly admissionWhitelistPatterns?: pulumi.Input<pulumi.Input<inputs.binaryauthorization.PolicyAdmissionWhitelistPattern>[]>;
    /**
     * Per-cluster admission rules. An admission rule specifies either that
     * all container images used in a pod creation request must be attested
     * to by one or more attestors, that all pod creations will be allowed,
     * or that all pod creations will be denied. There can be at most one
     * admission rule per cluster spec.
     */
    readonly clusterAdmissionRules?: pulumi.Input<pulumi.Input<inputs.binaryauthorization.PolicyClusterAdmissionRule>[]>;
    /**
     * Default admission rule for a cluster without a per-cluster admission
     * rule.
     * Structure is documented below.
     */
    readonly defaultAdmissionRule?: pulumi.Input<inputs.binaryauthorization.PolicyDefaultAdmissionRule>;
    /**
     * A descriptive comment.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Controls the evaluation of a Google-maintained global admission policy
     * for common system-level images. Images not covered by the global
     * policy will be subject to the project admission policy.
     * Possible values are `ENABLE` and `DISABLE`.
     */
    readonly globalPolicyEvaluationMode?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Policy resource.
 */
export interface PolicyArgs {
    /**
     * A whitelist of image patterns to exclude from admission rules. If an
     * image's name matches a whitelist pattern, the image's admission
     * requests will always be permitted regardless of your admission rules.
     * Structure is documented below.
     */
    readonly admissionWhitelistPatterns?: pulumi.Input<pulumi.Input<inputs.binaryauthorization.PolicyAdmissionWhitelistPattern>[]>;
    /**
     * Per-cluster admission rules. An admission rule specifies either that
     * all container images used in a pod creation request must be attested
     * to by one or more attestors, that all pod creations will be allowed,
     * or that all pod creations will be denied. There can be at most one
     * admission rule per cluster spec.
     */
    readonly clusterAdmissionRules?: pulumi.Input<pulumi.Input<inputs.binaryauthorization.PolicyClusterAdmissionRule>[]>;
    /**
     * Default admission rule for a cluster without a per-cluster admission
     * rule.
     * Structure is documented below.
     */
    readonly defaultAdmissionRule: pulumi.Input<inputs.binaryauthorization.PolicyDefaultAdmissionRule>;
    /**
     * A descriptive comment.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Controls the evaluation of a Google-maintained global admission policy
     * for common system-level images. Images not covered by the global
     * policy will be subject to the project admission policy.
     * Possible values are `ENABLE` and `DISABLE`.
     */
    readonly globalPolicyEvaluationMode?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
}
