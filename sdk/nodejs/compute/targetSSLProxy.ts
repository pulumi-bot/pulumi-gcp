// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export class TargetSSLProxy extends pulumi.CustomResource {
    /**
     * Get an existing TargetSSLProxy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: TargetSSLProxyState, opts?: pulumi.CustomResourceOptions): TargetSSLProxy {
        return new TargetSSLProxy(name, <any>state, { ...opts, id: id });
    }

    public readonly backendService: pulumi.Output<string>;
    public /*out*/ readonly creationTimestamp: pulumi.Output<string>;
    public readonly description: pulumi.Output<string | undefined>;
    public readonly name: pulumi.Output<string>;
    public readonly project: pulumi.Output<string>;
    public readonly proxyHeader: pulumi.Output<string | undefined>;
    public /*out*/ readonly proxyId: pulumi.Output<number>;
    public /*out*/ readonly selfLink: pulumi.Output<string>;
    public readonly sslCertificates: pulumi.Output<string>;
    public readonly sslPolicy: pulumi.Output<string | undefined>;

    /**
     * Create a TargetSSLProxy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TargetSSLProxyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: TargetSSLProxyArgs | TargetSSLProxyState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: TargetSSLProxyState = argsOrState as TargetSSLProxyState | undefined;
            inputs["backendService"] = state ? state.backendService : undefined;
            inputs["creationTimestamp"] = state ? state.creationTimestamp : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["proxyHeader"] = state ? state.proxyHeader : undefined;
            inputs["proxyId"] = state ? state.proxyId : undefined;
            inputs["selfLink"] = state ? state.selfLink : undefined;
            inputs["sslCertificates"] = state ? state.sslCertificates : undefined;
            inputs["sslPolicy"] = state ? state.sslPolicy : undefined;
        } else {
            const args = argsOrState as TargetSSLProxyArgs | undefined;
            if (!args || args.backendService === undefined) {
                throw new Error("Missing required property 'backendService'");
            }
            if (!args || args.sslCertificates === undefined) {
                throw new Error("Missing required property 'sslCertificates'");
            }
            inputs["backendService"] = args ? args.backendService : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["proxyHeader"] = args ? args.proxyHeader : undefined;
            inputs["sslCertificates"] = args ? args.sslCertificates : undefined;
            inputs["sslPolicy"] = args ? args.sslPolicy : undefined;
            inputs["creationTimestamp"] = undefined /*out*/;
            inputs["proxyId"] = undefined /*out*/;
            inputs["selfLink"] = undefined /*out*/;
        }
        super("gcp:compute/targetSSLProxy:TargetSSLProxy", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering TargetSSLProxy resources.
 */
export interface TargetSSLProxyState {
    readonly backendService?: pulumi.Input<string>;
    readonly creationTimestamp?: pulumi.Input<string>;
    readonly description?: pulumi.Input<string>;
    readonly name?: pulumi.Input<string>;
    readonly project?: pulumi.Input<string>;
    readonly proxyHeader?: pulumi.Input<string>;
    readonly proxyId?: pulumi.Input<number>;
    readonly selfLink?: pulumi.Input<string>;
    readonly sslCertificates?: pulumi.Input<string>;
    readonly sslPolicy?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a TargetSSLProxy resource.
 */
export interface TargetSSLProxyArgs {
    readonly backendService: pulumi.Input<string>;
    readonly description?: pulumi.Input<string>;
    readonly name?: pulumi.Input<string>;
    readonly project?: pulumi.Input<string>;
    readonly proxyHeader?: pulumi.Input<string>;
    readonly sslCertificates: pulumi.Input<string>;
    readonly sslPolicy?: pulumi.Input<string>;
}
