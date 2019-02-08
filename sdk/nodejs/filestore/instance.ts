// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * A Google Cloud Filestore instance.
 * 
 * > **Warning:** This resource is in beta, and should be used with the terraform-provider-google-beta provider.
 * See [Provider Versions](https://terraform.io/docs/providers/google/provider_versions.html) for more details on beta resources.
 * 
 * To get more information about Instance, see:
 * 
 * * [API documentation](https://cloud.google.com/filestore/docs/reference/rest/v1beta1/projects.locations.instances/create)
 * * How-to Guides
 *     * [Official Documentation](https://cloud.google.com/filestore/docs/creating-instances)
 *     * [Use with Kubernetes](https://cloud.google.com/filestore/docs/accessing-fileshares)
 *     * [Copying Data In/Out](https://cloud.google.com/filestore/docs/copying-data)
 * 
 * <div class = "oics-button" style="float: right; margin: 0 0 -15px">
 *   <a href="https://console.cloud.google.com/cloudshell/open?cloudshell_git_repo=https%3A%2F%2Fgithub.com%2Fterraform-google-modules%2Fdocs-examples.git&cloudshell_working_dir=filestore_instance_basic&cloudshell_image=gcr.io%2Fgraphite-cloud-shell-images%2Fterraform%3Alatest&open_in_editor=main.tf&cloudshell_print=.%2Fmotd&cloudshell_tutorial=.%2Ftutorial.md" target="_blank">
 *     <img alt="Open in Cloud Shell" src="//gstatic.com/cloudssh/images/open-btn.svg" style="max-height: 44px; margin: 32px auto; max-width: 100%;">
 *   </a>
 * </div>
 * ## Example Usage - Filestore Instance Basic
 * 
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 * 
 * const instance = new gcp.filestore.Instance("instance", {
 *     fileShares: {
 *         capacityGb: 2660,
 *         name: "share1",
 *     },
 *     networks: [{
 *         modes: ["MODE_IPV4"],
 *         network: "default",
 *     }],
 *     tier: "PREMIUM",
 *     zone: "us-central1-b",
 * });
 * ```
 */
export class Instance extends pulumi.CustomResource {
    /**
     * Get an existing Instance resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: InstanceState, opts?: pulumi.CustomResourceOptions): Instance {
        return new Instance(name, <any>state, { ...opts, id: id });
    }

    public /*out*/ readonly createTime: pulumi.Output<string>;
    public readonly description: pulumi.Output<string | undefined>;
    public /*out*/ readonly etag: pulumi.Output<string>;
    public readonly fileShares: pulumi.Output<{ capacityGb: number, name: string }>;
    public readonly labels: pulumi.Output<{[key: string]: string} | undefined>;
    public readonly name: pulumi.Output<string>;
    public readonly networks: pulumi.Output<{ ipAddresses: string[], modes: string[], network: string, reservedIpRange: string }[]>;
    public readonly project: pulumi.Output<string>;
    public readonly tier: pulumi.Output<string>;
    public readonly zone: pulumi.Output<string>;

    /**
     * Create a Instance resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: InstanceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: InstanceArgs | InstanceState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: InstanceState = argsOrState as InstanceState | undefined;
            inputs["createTime"] = state ? state.createTime : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["etag"] = state ? state.etag : undefined;
            inputs["fileShares"] = state ? state.fileShares : undefined;
            inputs["labels"] = state ? state.labels : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["networks"] = state ? state.networks : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["tier"] = state ? state.tier : undefined;
            inputs["zone"] = state ? state.zone : undefined;
        } else {
            const args = argsOrState as InstanceArgs | undefined;
            if (!args || args.fileShares === undefined) {
                throw new Error("Missing required property 'fileShares'");
            }
            if (!args || args.networks === undefined) {
                throw new Error("Missing required property 'networks'");
            }
            if (!args || args.tier === undefined) {
                throw new Error("Missing required property 'tier'");
            }
            if (!args || args.zone === undefined) {
                throw new Error("Missing required property 'zone'");
            }
            inputs["description"] = args ? args.description : undefined;
            inputs["fileShares"] = args ? args.fileShares : undefined;
            inputs["labels"] = args ? args.labels : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["networks"] = args ? args.networks : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["tier"] = args ? args.tier : undefined;
            inputs["zone"] = args ? args.zone : undefined;
            inputs["createTime"] = undefined /*out*/;
            inputs["etag"] = undefined /*out*/;
        }
        super("gcp:filestore/instance:Instance", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Instance resources.
 */
export interface InstanceState {
    readonly createTime?: pulumi.Input<string>;
    readonly description?: pulumi.Input<string>;
    readonly etag?: pulumi.Input<string>;
    readonly fileShares?: pulumi.Input<{ capacityGb: pulumi.Input<number>, name: pulumi.Input<string> }>;
    readonly labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    readonly name?: pulumi.Input<string>;
    readonly networks?: pulumi.Input<pulumi.Input<{ ipAddresses?: pulumi.Input<pulumi.Input<string>[]>, modes: pulumi.Input<pulumi.Input<string>[]>, network: pulumi.Input<string>, reservedIpRange?: pulumi.Input<string> }>[]>;
    readonly project?: pulumi.Input<string>;
    readonly tier?: pulumi.Input<string>;
    readonly zone?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Instance resource.
 */
export interface InstanceArgs {
    readonly description?: pulumi.Input<string>;
    readonly fileShares: pulumi.Input<{ capacityGb: pulumi.Input<number>, name: pulumi.Input<string> }>;
    readonly labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    readonly name?: pulumi.Input<string>;
    readonly networks: pulumi.Input<pulumi.Input<{ ipAddresses?: pulumi.Input<pulumi.Input<string>[]>, modes: pulumi.Input<pulumi.Input<string>[]>, network: pulumi.Input<string>, reservedIpRange?: pulumi.Input<string> }>[]>;
    readonly project?: pulumi.Input<string>;
    readonly tier: pulumi.Input<string>;
    readonly zone: pulumi.Input<string>;
}
