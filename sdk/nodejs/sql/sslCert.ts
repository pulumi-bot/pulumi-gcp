// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Creates a new Google SQL SSL Cert on a Google SQL Instance. For more information, see the [official documentation](https://cloud.google.com/sql/), or the [JSON API](https://cloud.google.com/sql/docs/mysql/admin-api/v1beta4/sslCerts).
 *
 * ## Example Usage
 *
 *
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 * import * as random from "@pulumi/random";
 *
 * const dbNameSuffix = new random.RandomId("dbNameSuffix", {byteLength: 4});
 * const master = new gcp.sql.DatabaseInstance("master", {settings: {
 *     tier: "db-f1-micro",
 * }});
 * const clientCert = new gcp.sql.SslCert("clientCert", {
 *     commonName: "client-name",
 *     instance: master.name,
 * });
 * ```
 */
export class SslCert extends pulumi.CustomResource {
    /**
     * Get an existing SslCert resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SslCertState, opts?: pulumi.CustomResourceOptions): SslCert {
        return new SslCert(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:sql/sslCert:SslCert';

    /**
     * Returns true if the given object is an instance of SslCert.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SslCert {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SslCert.__pulumiType;
    }

    /**
     * The actual certificate data for this client certificate.
     */
    public /*out*/ readonly cert!: pulumi.Output<string>;
    /**
     * The serial number extracted from the certificate data.
     */
    public /*out*/ readonly certSerialNumber!: pulumi.Output<string>;
    /**
     * The common name to be used in the certificate to identify the
     * client. Constrained to [a-zA-Z.-_ ]+. Changing this forces a new resource to be created.
     */
    public readonly commonName!: pulumi.Output<string>;
    /**
     * The time when the certificate was created in RFC 3339 format,
     * for example 2012-11-15T16:19:00.094Z.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * The time when the certificate expires in RFC 3339 format,
     * for example 2012-11-15T16:19:00.094Z.
     */
    public /*out*/ readonly expirationTime!: pulumi.Output<string>;
    /**
     * The name of the Cloud SQL instance. Changing this
     * forces a new resource to be created.
     */
    public readonly instance!: pulumi.Output<string>;
    /**
     * The private key associated with the client certificate.
     */
    public /*out*/ readonly privateKey!: pulumi.Output<string>;
    /**
     * The ID of the project in which the resource belongs. If it
     * is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * The CA cert of the server this client cert was generated from.
     */
    public /*out*/ readonly serverCaCert!: pulumi.Output<string>;
    /**
     * The SHA1 Fingerprint of the certificate.
     */
    public /*out*/ readonly sha1Fingerprint!: pulumi.Output<string>;

    /**
     * Create a SslCert resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SslCertArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SslCertArgs | SslCertState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as SslCertState | undefined;
            inputs["cert"] = state ? state.cert : undefined;
            inputs["certSerialNumber"] = state ? state.certSerialNumber : undefined;
            inputs["commonName"] = state ? state.commonName : undefined;
            inputs["createTime"] = state ? state.createTime : undefined;
            inputs["expirationTime"] = state ? state.expirationTime : undefined;
            inputs["instance"] = state ? state.instance : undefined;
            inputs["privateKey"] = state ? state.privateKey : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["serverCaCert"] = state ? state.serverCaCert : undefined;
            inputs["sha1Fingerprint"] = state ? state.sha1Fingerprint : undefined;
        } else {
            const args = argsOrState as SslCertArgs | undefined;
            if (!args || args.commonName === undefined) {
                throw new Error("Missing required property 'commonName'");
            }
            if (!args || args.instance === undefined) {
                throw new Error("Missing required property 'instance'");
            }
            inputs["commonName"] = args ? args.commonName : undefined;
            inputs["instance"] = args ? args.instance : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["cert"] = undefined /*out*/;
            inputs["certSerialNumber"] = undefined /*out*/;
            inputs["createTime"] = undefined /*out*/;
            inputs["expirationTime"] = undefined /*out*/;
            inputs["privateKey"] = undefined /*out*/;
            inputs["serverCaCert"] = undefined /*out*/;
            inputs["sha1Fingerprint"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(SslCert.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SslCert resources.
 */
export interface SslCertState {
    /**
     * The actual certificate data for this client certificate.
     */
    readonly cert?: pulumi.Input<string>;
    /**
     * The serial number extracted from the certificate data.
     */
    readonly certSerialNumber?: pulumi.Input<string>;
    /**
     * The common name to be used in the certificate to identify the
     * client. Constrained to [a-zA-Z.-_ ]+. Changing this forces a new resource to be created.
     */
    readonly commonName?: pulumi.Input<string>;
    /**
     * The time when the certificate was created in RFC 3339 format,
     * for example 2012-11-15T16:19:00.094Z.
     */
    readonly createTime?: pulumi.Input<string>;
    /**
     * The time when the certificate expires in RFC 3339 format,
     * for example 2012-11-15T16:19:00.094Z.
     */
    readonly expirationTime?: pulumi.Input<string>;
    /**
     * The name of the Cloud SQL instance. Changing this
     * forces a new resource to be created.
     */
    readonly instance?: pulumi.Input<string>;
    /**
     * The private key associated with the client certificate.
     */
    readonly privateKey?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs. If it
     * is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * The CA cert of the server this client cert was generated from.
     */
    readonly serverCaCert?: pulumi.Input<string>;
    /**
     * The SHA1 Fingerprint of the certificate.
     */
    readonly sha1Fingerprint?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SslCert resource.
 */
export interface SslCertArgs {
    /**
     * The common name to be used in the certificate to identify the
     * client. Constrained to [a-zA-Z.-_ ]+. Changing this forces a new resource to be created.
     */
    readonly commonName: pulumi.Input<string>;
    /**
     * The name of the Cloud SQL instance. Changing this
     * forces a new resource to be created.
     */
    readonly instance: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs. If it
     * is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
}
