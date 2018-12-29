// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export class Variavble extends pulumi.CustomResource {
    /**
     * Get an existing Variavble resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: VariavbleState, opts?: pulumi.CustomResourceOptions): Variavble {
        return new Variavble(name, <any>state, { ...opts, id: id });
    }

    public readonly name: pulumi.Output<string>;
    public readonly parent: pulumi.Output<string>;
    public readonly project: pulumi.Output<string>;
    public readonly text: pulumi.Output<string | undefined>;
    public /*out*/ readonly updateTime: pulumi.Output<string>;
    public readonly value: pulumi.Output<string | undefined>;

    /**
     * Create a Variavble resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VariavbleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: VariavbleArgs | VariavbleState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: VariavbleState = argsOrState as VariavbleState | undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["parent"] = state ? state.parent : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["text"] = state ? state.text : undefined;
            inputs["updateTime"] = state ? state.updateTime : undefined;
            inputs["value"] = state ? state.value : undefined;
        } else {
            const args = argsOrState as VariavbleArgs | undefined;
            if (!args || args.parent === undefined) {
                throw new Error("Missing required property 'parent'");
            }
            inputs["name"] = args ? args.name : undefined;
            inputs["parent"] = args ? args.parent : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["text"] = args ? args.text : undefined;
            inputs["value"] = args ? args.value : undefined;
            inputs["updateTime"] = undefined /*out*/;
        }
        super("gcp:runtimeconfig/variavble:Variavble", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Variavble resources.
 */
export interface VariavbleState {
    readonly name?: pulumi.Input<string>;
    readonly parent?: pulumi.Input<string>;
    readonly project?: pulumi.Input<string>;
    readonly text?: pulumi.Input<string>;
    readonly updateTime?: pulumi.Input<string>;
    readonly value?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Variavble resource.
 */
export interface VariavbleArgs {
    readonly name?: pulumi.Input<string>;
    readonly parent: pulumi.Input<string>;
    readonly project?: pulumi.Input<string>;
    readonly text?: pulumi.Input<string>;
    readonly value?: pulumi.Input<string>;
}
