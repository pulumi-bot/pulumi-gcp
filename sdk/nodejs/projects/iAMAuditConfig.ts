// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export class IAMAuditConfig extends pulumi.CustomResource {
    /**
     * Get an existing IAMAuditConfig resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: IAMAuditConfigState, opts?: pulumi.CustomResourceOptions): IAMAuditConfig {
        return new IAMAuditConfig(name, <any>state, { ...opts, id: id });
    }

    public readonly auditLogConfigs!: pulumi.Output<{ exemptedMembers?: string[], logType: string }[]>;
    public /*out*/ readonly etag!: pulumi.Output<string>;
    public readonly project!: pulumi.Output<string | undefined>;
    public readonly service!: pulumi.Output<string>;

    /**
     * Create a IAMAuditConfig resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: IAMAuditConfigArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: IAMAuditConfigArgs | IAMAuditConfigState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as IAMAuditConfigState | undefined;
            inputs["auditLogConfigs"] = state ? state.auditLogConfigs : undefined;
            inputs["etag"] = state ? state.etag : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["service"] = state ? state.service : undefined;
        } else {
            const args = argsOrState as IAMAuditConfigArgs | undefined;
            if (!args || args.auditLogConfigs === undefined) {
                throw new Error("Missing required property 'auditLogConfigs'");
            }
            if (!args || args.service === undefined) {
                throw new Error("Missing required property 'service'");
            }
            inputs["auditLogConfigs"] = args ? args.auditLogConfigs : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["service"] = args ? args.service : undefined;
            inputs["etag"] = undefined /*out*/;
        }
        super("gcp:projects/iAMAuditConfig:IAMAuditConfig", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering IAMAuditConfig resources.
 */
export interface IAMAuditConfigState {
    readonly auditLogConfigs?: pulumi.Input<pulumi.Input<{ exemptedMembers?: pulumi.Input<pulumi.Input<string>[]>, logType: pulumi.Input<string> }>[]>;
    readonly etag?: pulumi.Input<string>;
    readonly project?: pulumi.Input<string>;
    readonly service?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a IAMAuditConfig resource.
 */
export interface IAMAuditConfigArgs {
    readonly auditLogConfigs: pulumi.Input<pulumi.Input<{ exemptedMembers?: pulumi.Input<pulumi.Input<string>[]>, logType: pulumi.Input<string> }>[]>;
    readonly project?: pulumi.Input<string>;
    readonly service: pulumi.Input<string>;
}
