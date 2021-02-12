// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export class AuthorityIamPolicy extends pulumi.CustomResource {
    /**
     * Get an existing AuthorityIamPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AuthorityIamPolicyState, opts?: pulumi.CustomResourceOptions): AuthorityIamPolicy {
        return new AuthorityIamPolicy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:certificateauthority/authorityIamPolicy:AuthorityIamPolicy';

    /**
     * Returns true if the given object is an instance of AuthorityIamPolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AuthorityIamPolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AuthorityIamPolicy.__pulumiType;
    }

    public readonly certificateAuthority!: pulumi.Output<string>;
    public /*out*/ readonly etag!: pulumi.Output<string>;
    public readonly policyData!: pulumi.Output<string>;

    /**
     * Create a AuthorityIamPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AuthorityIamPolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AuthorityIamPolicyArgs | AuthorityIamPolicyState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AuthorityIamPolicyState | undefined;
            inputs["certificateAuthority"] = state ? state.certificateAuthority : undefined;
            inputs["etag"] = state ? state.etag : undefined;
            inputs["policyData"] = state ? state.policyData : undefined;
        } else {
            const args = argsOrState as AuthorityIamPolicyArgs | undefined;
            if ((!args || args.certificateAuthority === undefined) && !opts.urn) {
                throw new Error("Missing required property 'certificateAuthority'");
            }
            if ((!args || args.policyData === undefined) && !opts.urn) {
                throw new Error("Missing required property 'policyData'");
            }
            inputs["certificateAuthority"] = args ? args.certificateAuthority : undefined;
            inputs["policyData"] = args ? args.policyData : undefined;
            inputs["etag"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(AuthorityIamPolicy.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AuthorityIamPolicy resources.
 */
export interface AuthorityIamPolicyState {
    readonly certificateAuthority?: pulumi.Input<string>;
    readonly etag?: pulumi.Input<string>;
    readonly policyData?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AuthorityIamPolicy resource.
 */
export interface AuthorityIamPolicyArgs {
    readonly certificateAuthority: pulumi.Input<string>;
    readonly policyData: pulumi.Input<string>;
}
