// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export class VPNTunnel extends pulumi.CustomResource {
    /**
     * Get an existing VPNTunnel resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: VPNTunnelState, opts?: pulumi.CustomResourceOptions): VPNTunnel {
        return new VPNTunnel(name, <any>state, { ...opts, id: id });
    }

    public /*out*/ readonly creationTimestamp: pulumi.Output<string>;
    public readonly description: pulumi.Output<string | undefined>;
    public /*out*/ readonly detailedStatus: pulumi.Output<string>;
    public readonly ikeVersion: pulumi.Output<number | undefined>;
    public /*out*/ readonly labelFingerprint: pulumi.Output<string>;
    public readonly labels: pulumi.Output<{[key: string]: string} | undefined>;
    public readonly localTrafficSelectors: pulumi.Output<string[]>;
    public readonly name: pulumi.Output<string>;
    public readonly peerIp: pulumi.Output<string>;
    public readonly project: pulumi.Output<string>;
    public readonly region: pulumi.Output<string>;
    public readonly remoteTrafficSelectors: pulumi.Output<string[]>;
    public readonly router: pulumi.Output<string | undefined>;
    public /*out*/ readonly selfLink: pulumi.Output<string>;
    public readonly sharedSecret: pulumi.Output<string>;
    public /*out*/ readonly sharedSecretHash: pulumi.Output<string>;
    public readonly targetVpnGateway: pulumi.Output<string>;

    /**
     * Create a VPNTunnel resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VPNTunnelArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: VPNTunnelArgs | VPNTunnelState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: VPNTunnelState = argsOrState as VPNTunnelState | undefined;
            inputs["creationTimestamp"] = state ? state.creationTimestamp : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["detailedStatus"] = state ? state.detailedStatus : undefined;
            inputs["ikeVersion"] = state ? state.ikeVersion : undefined;
            inputs["labelFingerprint"] = state ? state.labelFingerprint : undefined;
            inputs["labels"] = state ? state.labels : undefined;
            inputs["localTrafficSelectors"] = state ? state.localTrafficSelectors : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["peerIp"] = state ? state.peerIp : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["region"] = state ? state.region : undefined;
            inputs["remoteTrafficSelectors"] = state ? state.remoteTrafficSelectors : undefined;
            inputs["router"] = state ? state.router : undefined;
            inputs["selfLink"] = state ? state.selfLink : undefined;
            inputs["sharedSecret"] = state ? state.sharedSecret : undefined;
            inputs["sharedSecretHash"] = state ? state.sharedSecretHash : undefined;
            inputs["targetVpnGateway"] = state ? state.targetVpnGateway : undefined;
        } else {
            const args = argsOrState as VPNTunnelArgs | undefined;
            if (!args || args.peerIp === undefined) {
                throw new Error("Missing required property 'peerIp'");
            }
            if (!args || args.sharedSecret === undefined) {
                throw new Error("Missing required property 'sharedSecret'");
            }
            if (!args || args.targetVpnGateway === undefined) {
                throw new Error("Missing required property 'targetVpnGateway'");
            }
            inputs["description"] = args ? args.description : undefined;
            inputs["ikeVersion"] = args ? args.ikeVersion : undefined;
            inputs["labels"] = args ? args.labels : undefined;
            inputs["localTrafficSelectors"] = args ? args.localTrafficSelectors : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["peerIp"] = args ? args.peerIp : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["region"] = args ? args.region : undefined;
            inputs["remoteTrafficSelectors"] = args ? args.remoteTrafficSelectors : undefined;
            inputs["router"] = args ? args.router : undefined;
            inputs["sharedSecret"] = args ? args.sharedSecret : undefined;
            inputs["targetVpnGateway"] = args ? args.targetVpnGateway : undefined;
            inputs["creationTimestamp"] = undefined /*out*/;
            inputs["detailedStatus"] = undefined /*out*/;
            inputs["labelFingerprint"] = undefined /*out*/;
            inputs["selfLink"] = undefined /*out*/;
            inputs["sharedSecretHash"] = undefined /*out*/;
        }
        super("gcp:compute/vPNTunnel:VPNTunnel", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering VPNTunnel resources.
 */
export interface VPNTunnelState {
    readonly creationTimestamp?: pulumi.Input<string>;
    readonly description?: pulumi.Input<string>;
    readonly detailedStatus?: pulumi.Input<string>;
    readonly ikeVersion?: pulumi.Input<number>;
    readonly labelFingerprint?: pulumi.Input<string>;
    readonly labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    readonly localTrafficSelectors?: pulumi.Input<pulumi.Input<string>[]>;
    readonly name?: pulumi.Input<string>;
    readonly peerIp?: pulumi.Input<string>;
    readonly project?: pulumi.Input<string>;
    readonly region?: pulumi.Input<string>;
    readonly remoteTrafficSelectors?: pulumi.Input<pulumi.Input<string>[]>;
    readonly router?: pulumi.Input<string>;
    readonly selfLink?: pulumi.Input<string>;
    readonly sharedSecret?: pulumi.Input<string>;
    readonly sharedSecretHash?: pulumi.Input<string>;
    readonly targetVpnGateway?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a VPNTunnel resource.
 */
export interface VPNTunnelArgs {
    readonly description?: pulumi.Input<string>;
    readonly ikeVersion?: pulumi.Input<number>;
    readonly labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    readonly localTrafficSelectors?: pulumi.Input<pulumi.Input<string>[]>;
    readonly name?: pulumi.Input<string>;
    readonly peerIp: pulumi.Input<string>;
    readonly project?: pulumi.Input<string>;
    readonly region?: pulumi.Input<string>;
    readonly remoteTrafficSelectors?: pulumi.Input<pulumi.Input<string>[]>;
    readonly router?: pulumi.Input<string>;
    readonly sharedSecret: pulumi.Input<string>;
    readonly targetVpnGateway: pulumi.Input<string>;
}
