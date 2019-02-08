// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * An SslCertificate resource, used for HTTPS load balancing. This resource
 * provides a mechanism to upload an SSL key and certificate to
 * the load balancer to serve secure connections from the user.
 * 
 * 
 * To get more information about SslCertificate, see:
 * 
 * * [API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/sslCertificates)
 * * How-to Guides
 *     * [Official Documentation](https://cloud.google.com/load-balancing/docs/ssl-certificates)
 * 
 * <div class = "oics-button" style="float: right; margin: 0 0 -15px">
 *   <a href="https://console.cloud.google.com/cloudshell/open?cloudshell_git_repo=https%3A%2F%2Fgithub.com%2Fterraform-google-modules%2Fdocs-examples.git&cloudshell_working_dir=ssl_certificate_basic&cloudshell_image=gcr.io%2Fgraphite-cloud-shell-images%2Fterraform%3Alatest&open_in_editor=main.tf&cloudshell_print=.%2Fmotd&cloudshell_tutorial=.%2Ftutorial.md" target="_blank">
 *     <img alt="Open in Cloud Shell" src="//gstatic.com/cloudssh/images/open-btn.svg" style="max-height: 44px; margin: 32px auto; max-width: 100%;">
 *   </a>
 * </div>
 */
export class SSLCertificate extends pulumi.CustomResource {
    /**
     * Get an existing SSLCertificate resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SSLCertificateState, opts?: pulumi.CustomResourceOptions): SSLCertificate {
        return new SSLCertificate(name, <any>state, { ...opts, id: id });
    }

    public readonly certificate: pulumi.Output<string>;
    public /*out*/ readonly certificateId: pulumi.Output<number>;
    public /*out*/ readonly creationTimestamp: pulumi.Output<string>;
    public readonly description: pulumi.Output<string | undefined>;
    public readonly name: pulumi.Output<string>;
    /**
     * Creates a unique name beginning with the
     * specified prefix. Conflicts with `name`.
     */
    public readonly namePrefix: pulumi.Output<string>;
    public readonly privateKey: pulumi.Output<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project: pulumi.Output<string>;
    /**
     * The URI of the created resource.
     */
    public /*out*/ readonly selfLink: pulumi.Output<string>;

    /**
     * Create a SSLCertificate resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SSLCertificateArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SSLCertificateArgs | SSLCertificateState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: SSLCertificateState = argsOrState as SSLCertificateState | undefined;
            inputs["certificate"] = state ? state.certificate : undefined;
            inputs["certificateId"] = state ? state.certificateId : undefined;
            inputs["creationTimestamp"] = state ? state.creationTimestamp : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["namePrefix"] = state ? state.namePrefix : undefined;
            inputs["privateKey"] = state ? state.privateKey : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["selfLink"] = state ? state.selfLink : undefined;
        } else {
            const args = argsOrState as SSLCertificateArgs | undefined;
            if (!args || args.certificate === undefined) {
                throw new Error("Missing required property 'certificate'");
            }
            if (!args || args.privateKey === undefined) {
                throw new Error("Missing required property 'privateKey'");
            }
            inputs["certificate"] = args ? args.certificate : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["namePrefix"] = args ? args.namePrefix : undefined;
            inputs["privateKey"] = args ? args.privateKey : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["certificateId"] = undefined /*out*/;
            inputs["creationTimestamp"] = undefined /*out*/;
            inputs["selfLink"] = undefined /*out*/;
        }
        super("gcp:compute/sSLCertificate:SSLCertificate", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SSLCertificate resources.
 */
export interface SSLCertificateState {
    readonly certificate?: pulumi.Input<string>;
    readonly certificateId?: pulumi.Input<number>;
    readonly creationTimestamp?: pulumi.Input<string>;
    readonly description?: pulumi.Input<string>;
    readonly name?: pulumi.Input<string>;
    /**
     * Creates a unique name beginning with the
     * specified prefix. Conflicts with `name`.
     */
    readonly namePrefix?: pulumi.Input<string>;
    readonly privateKey?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * The URI of the created resource.
     */
    readonly selfLink?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SSLCertificate resource.
 */
export interface SSLCertificateArgs {
    readonly certificate: pulumi.Input<string>;
    readonly description?: pulumi.Input<string>;
    readonly name?: pulumi.Input<string>;
    /**
     * Creates a unique name beginning with the
     * specified prefix. Conflicts with `name`.
     */
    readonly namePrefix?: pulumi.Input<string>;
    readonly privateKey: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
}
