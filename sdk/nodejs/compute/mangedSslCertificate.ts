// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputApi from "../types/input";
import * as outputApi from "../types/output";
import * as utilities from "../utilities";

/**
 * An SslCertificate resource, used for HTTPS load balancing.  This resource
 * represents a certificate for which the certificate secrets are created and
 * managed by Google.
 * 
 * For a resource where you provide the key, see the
 * SSL Certificate resource.
 * 
 * To get more information about ManagedSslCertificate, see:
 * 
 * * [API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/sslCertificates)
 * * How-to Guides
 *     * [Official Documentation](https://cloud.google.com/load-balancing/docs/ssl-certificates)
 * 
 * > **Warning:** This resource should be used with extreme caution!  Provisioning an SSL
 * certificate is complex.  Ensure that you understand the lifecycle of a
 * certificate before attempting complex tasks like cert rotation automatically.
 * This resource will "return" as soon as the certificate object is created,
 * but post-creation the certificate object will go through a "provisioning"
 * process.  The provisioning process can complete only when the domain name
 * for which the certificate is created points to a target pool which, itself,
 * points at the certificate.  Depending on your DNS provider, this may take
 * some time, and migrating from self-managed certificates to Google-managed
 * certificates may entail some downtime while the certificate provisions.
 * 
 * In conclusion: Be extremely cautious.
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/compute_managed_ssl_certificate.html.markdown.
 */
export class MangedSslCertificate extends pulumi.CustomResource {
    /**
     * Get an existing MangedSslCertificate resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: MangedSslCertificateState, opts?: pulumi.CustomResourceOptions): MangedSslCertificate {
        return new MangedSslCertificate(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:compute/mangedSslCertificate:MangedSslCertificate';

    /**
     * Returns true if the given object is an instance of MangedSslCertificate.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is MangedSslCertificate {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === MangedSslCertificate.__pulumiType;
    }

    public readonly certificateId!: pulumi.Output<number>;
    public /*out*/ readonly creationTimestamp!: pulumi.Output<string>;
    public readonly description!: pulumi.Output<string | undefined>;
    public /*out*/ readonly expireTime!: pulumi.Output<string>;
    public readonly managed!: pulumi.Output<outputApi.compute.MangedSslCertificateManaged | undefined>;
    public readonly name!: pulumi.Output<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * The URI of the created resource.
     */
    public /*out*/ readonly selfLink!: pulumi.Output<string>;
    public /*out*/ readonly subjectAlternativeNames!: pulumi.Output<string[]>;
    public readonly type!: pulumi.Output<string | undefined>;

    /**
     * Create a MangedSslCertificate resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: MangedSslCertificateArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: MangedSslCertificateArgs | MangedSslCertificateState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as MangedSslCertificateState | undefined;
            inputs["certificateId"] = state ? state.certificateId : undefined;
            inputs["creationTimestamp"] = state ? state.creationTimestamp : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["expireTime"] = state ? state.expireTime : undefined;
            inputs["managed"] = state ? state.managed : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["selfLink"] = state ? state.selfLink : undefined;
            inputs["subjectAlternativeNames"] = state ? state.subjectAlternativeNames : undefined;
            inputs["type"] = state ? state.type : undefined;
        } else {
            const args = argsOrState as MangedSslCertificateArgs | undefined;
            inputs["certificateId"] = args ? args.certificateId : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["managed"] = args ? args.managed : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["type"] = args ? args.type : undefined;
            inputs["creationTimestamp"] = undefined /*out*/;
            inputs["expireTime"] = undefined /*out*/;
            inputs["selfLink"] = undefined /*out*/;
            inputs["subjectAlternativeNames"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(MangedSslCertificate.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering MangedSslCertificate resources.
 */
export interface MangedSslCertificateState {
    readonly certificateId?: pulumi.Input<number>;
    readonly creationTimestamp?: pulumi.Input<string>;
    readonly description?: pulumi.Input<string>;
    readonly expireTime?: pulumi.Input<string>;
    readonly managed?: pulumi.Input<inputApi.compute.MangedSslCertificateManaged>;
    readonly name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * The URI of the created resource.
     */
    readonly selfLink?: pulumi.Input<string>;
    readonly subjectAlternativeNames?: pulumi.Input<pulumi.Input<string>[]>;
    readonly type?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a MangedSslCertificate resource.
 */
export interface MangedSslCertificateArgs {
    readonly certificateId?: pulumi.Input<number>;
    readonly description?: pulumi.Input<string>;
    readonly managed?: pulumi.Input<inputApi.compute.MangedSslCertificateManaged>;
    readonly name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    readonly type?: pulumi.Input<string>;
}
