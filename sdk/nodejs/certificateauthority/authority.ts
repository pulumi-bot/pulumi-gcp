// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * ## Import
 *
 * CertificateAuthority can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import gcp:certificateauthority/authority:Authority default projects/{{project}}/locations/{{location}}/certificateAuthorities/{{certificate_authority_id}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:certificateauthority/authority:Authority default {{project}}/{{location}}/{{certificate_authority_id}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:certificateauthority/authority:Authority default {{location}}/{{certificate_authority_id}}
 * ```
 */
export class Authority extends pulumi.CustomResource {
    /**
     * Get an existing Authority resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AuthorityState, opts?: pulumi.CustomResourceOptions): Authority {
        return new Authority(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:certificateauthority/authority:Authority';

    /**
     * Returns true if the given object is an instance of Authority.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Authority {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Authority.__pulumiType;
    }

    /**
     * URLs for accessing content published by this CA, such as the CA certificate and CRLs.
     */
    public /*out*/ readonly accessUrls!: pulumi.Output<outputs.certificateauthority.AuthorityAccessUrl[]>;
    /**
     * The user provided Resource ID for this Certificate Authority.
     */
    public readonly certificateAuthorityId!: pulumi.Output<string>;
    /**
     * The config used to create a self-signed X.509 certificate or CSR.
     * Structure is documented below.
     */
    public readonly config!: pulumi.Output<outputs.certificateauthority.AuthorityConfig>;
    /**
     * The time at which this CertificateAuthority was created. A timestamp in RFC3339 UTC "Zulu" format, with nanosecond
     * resolution and up to nine fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * If set to `true`, the Certificate Authority will be disabled
     * on delete. If the Certitificate Authorities is not disabled,
     * it cannot be deleted. Use with care. Defaults to `false`.
     */
    public readonly disableOnDelete!: pulumi.Output<boolean | undefined>;
    /**
     * The name of a Cloud Storage bucket where this CertificateAuthority will publish content,
     * such as the CA certificate and CRLs. This must be a bucket name, without any prefixes
     * (such as `gs://`) or suffixes (such as `.googleapis.com`). For example, to use a bucket named
     * my-bucket, you would simply specify `my-bucket`. If not specified, a managed bucket will be
     * created.
     */
    public readonly gcsBucket!: pulumi.Output<string | undefined>;
    /**
     * Options that affect all certificates issued by a CertificateAuthority.
     * Structure is documented below.
     */
    public readonly issuingOptions!: pulumi.Output<outputs.certificateauthority.AuthorityIssuingOptions | undefined>;
    /**
     * Used when issuing certificates for this CertificateAuthority. If this CertificateAuthority
     * is a self-signed CertificateAuthority, this key is also used to sign the self-signed CA
     * certificate. Otherwise, it is used to sign a CSR.
     * Structure is documented below.
     */
    public readonly keySpec!: pulumi.Output<outputs.certificateauthority.AuthorityKeySpec>;
    /**
     * Labels with user-defined metadata.
     * An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass":
     * "1.3kg", "count": "3" }.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The desired lifetime of the CA certificate. Used to create the "notBeforeTime" and
     * "notAfterTime" fields inside an X.509 certificate. A duration in seconds with up to nine
     * fractional digits, terminated by 's'. Example: "3.5s".
     */
    public readonly lifetime!: pulumi.Output<string | undefined>;
    /**
     * Location of the CertificateAuthority. A full list of valid locations can be found by
     * running `gcloud beta privateca locations list`.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * The resource name for this CertificateAuthority in the format projects/*&#47;locations/*&#47;certificateAuthorities/*.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * This CertificateAuthority's certificate chain, including the current CertificateAuthority's certificate. Ordered such
     * that the root issuer is the final element (consistent with RFC 5246). For a self-signed CA, this will only list the
     * current CertificateAuthority's certificate.
     */
    public /*out*/ readonly pemCaCertificates!: pulumi.Output<string[]>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * The State for this CertificateAuthority.
     */
    public /*out*/ readonly state!: pulumi.Output<string>;
    /**
     * The Tier of this CertificateAuthority. `ENTERPRISE` Certificate Authorities track
     * server side certificates issued, and support certificate revocation. For more details,
     * please check the [associated documentation](https://cloud.google.com/certificate-authority-service/docs/tiers).
     * Default value is `ENTERPRISE`.
     * Possible values are `ENTERPRISE` and `DEVOPS`.
     */
    public readonly tier!: pulumi.Output<string | undefined>;
    /**
     * The Type of this CertificateAuthority.
     * > **Note:** For `SUBORDINATE` Certificate Authorities, they need to
     * be manually activated (via Cloud Console of `gcloud`) before they can
     * issue certificates.
     * Default value is `SELF_SIGNED`.
     * Possible values are `SELF_SIGNED` and `SUBORDINATE`.
     */
    public readonly type!: pulumi.Output<string | undefined>;
    /**
     * The time at which this CertificateAuthority was updated. A timestamp in RFC3339 UTC "Zulu" format, with nanosecond
     * resolution and up to nine fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
     */
    public /*out*/ readonly updateTime!: pulumi.Output<string>;

    /**
     * Create a Authority resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AuthorityArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AuthorityArgs | AuthorityState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AuthorityState | undefined;
            inputs["accessUrls"] = state ? state.accessUrls : undefined;
            inputs["certificateAuthorityId"] = state ? state.certificateAuthorityId : undefined;
            inputs["config"] = state ? state.config : undefined;
            inputs["createTime"] = state ? state.createTime : undefined;
            inputs["disableOnDelete"] = state ? state.disableOnDelete : undefined;
            inputs["gcsBucket"] = state ? state.gcsBucket : undefined;
            inputs["issuingOptions"] = state ? state.issuingOptions : undefined;
            inputs["keySpec"] = state ? state.keySpec : undefined;
            inputs["labels"] = state ? state.labels : undefined;
            inputs["lifetime"] = state ? state.lifetime : undefined;
            inputs["location"] = state ? state.location : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["pemCaCertificates"] = state ? state.pemCaCertificates : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["state"] = state ? state.state : undefined;
            inputs["tier"] = state ? state.tier : undefined;
            inputs["type"] = state ? state.type : undefined;
            inputs["updateTime"] = state ? state.updateTime : undefined;
        } else {
            const args = argsOrState as AuthorityArgs | undefined;
            if ((!args || args.certificateAuthorityId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'certificateAuthorityId'");
            }
            if ((!args || args.config === undefined) && !opts.urn) {
                throw new Error("Missing required property 'config'");
            }
            if ((!args || args.keySpec === undefined) && !opts.urn) {
                throw new Error("Missing required property 'keySpec'");
            }
            if ((!args || args.location === undefined) && !opts.urn) {
                throw new Error("Missing required property 'location'");
            }
            inputs["certificateAuthorityId"] = args ? args.certificateAuthorityId : undefined;
            inputs["config"] = args ? args.config : undefined;
            inputs["disableOnDelete"] = args ? args.disableOnDelete : undefined;
            inputs["gcsBucket"] = args ? args.gcsBucket : undefined;
            inputs["issuingOptions"] = args ? args.issuingOptions : undefined;
            inputs["keySpec"] = args ? args.keySpec : undefined;
            inputs["labels"] = args ? args.labels : undefined;
            inputs["lifetime"] = args ? args.lifetime : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["tier"] = args ? args.tier : undefined;
            inputs["type"] = args ? args.type : undefined;
            inputs["accessUrls"] = undefined /*out*/;
            inputs["createTime"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["pemCaCertificates"] = undefined /*out*/;
            inputs["state"] = undefined /*out*/;
            inputs["updateTime"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Authority.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Authority resources.
 */
export interface AuthorityState {
    /**
     * URLs for accessing content published by this CA, such as the CA certificate and CRLs.
     */
    readonly accessUrls?: pulumi.Input<pulumi.Input<inputs.certificateauthority.AuthorityAccessUrl>[]>;
    /**
     * The user provided Resource ID for this Certificate Authority.
     */
    readonly certificateAuthorityId?: pulumi.Input<string>;
    /**
     * The config used to create a self-signed X.509 certificate or CSR.
     * Structure is documented below.
     */
    readonly config?: pulumi.Input<inputs.certificateauthority.AuthorityConfig>;
    /**
     * The time at which this CertificateAuthority was created. A timestamp in RFC3339 UTC "Zulu" format, with nanosecond
     * resolution and up to nine fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
     */
    readonly createTime?: pulumi.Input<string>;
    /**
     * If set to `true`, the Certificate Authority will be disabled
     * on delete. If the Certitificate Authorities is not disabled,
     * it cannot be deleted. Use with care. Defaults to `false`.
     */
    readonly disableOnDelete?: pulumi.Input<boolean>;
    /**
     * The name of a Cloud Storage bucket where this CertificateAuthority will publish content,
     * such as the CA certificate and CRLs. This must be a bucket name, without any prefixes
     * (such as `gs://`) or suffixes (such as `.googleapis.com`). For example, to use a bucket named
     * my-bucket, you would simply specify `my-bucket`. If not specified, a managed bucket will be
     * created.
     */
    readonly gcsBucket?: pulumi.Input<string>;
    /**
     * Options that affect all certificates issued by a CertificateAuthority.
     * Structure is documented below.
     */
    readonly issuingOptions?: pulumi.Input<inputs.certificateauthority.AuthorityIssuingOptions>;
    /**
     * Used when issuing certificates for this CertificateAuthority. If this CertificateAuthority
     * is a self-signed CertificateAuthority, this key is also used to sign the self-signed CA
     * certificate. Otherwise, it is used to sign a CSR.
     * Structure is documented below.
     */
    readonly keySpec?: pulumi.Input<inputs.certificateauthority.AuthorityKeySpec>;
    /**
     * Labels with user-defined metadata.
     * An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass":
     * "1.3kg", "count": "3" }.
     */
    readonly labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The desired lifetime of the CA certificate. Used to create the "notBeforeTime" and
     * "notAfterTime" fields inside an X.509 certificate. A duration in seconds with up to nine
     * fractional digits, terminated by 's'. Example: "3.5s".
     */
    readonly lifetime?: pulumi.Input<string>;
    /**
     * Location of the CertificateAuthority. A full list of valid locations can be found by
     * running `gcloud beta privateca locations list`.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The resource name for this CertificateAuthority in the format projects/*&#47;locations/*&#47;certificateAuthorities/*.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * This CertificateAuthority's certificate chain, including the current CertificateAuthority's certificate. Ordered such
     * that the root issuer is the final element (consistent with RFC 5246). For a self-signed CA, this will only list the
     * current CertificateAuthority's certificate.
     */
    readonly pemCaCertificates?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * The State for this CertificateAuthority.
     */
    readonly state?: pulumi.Input<string>;
    /**
     * The Tier of this CertificateAuthority. `ENTERPRISE` Certificate Authorities track
     * server side certificates issued, and support certificate revocation. For more details,
     * please check the [associated documentation](https://cloud.google.com/certificate-authority-service/docs/tiers).
     * Default value is `ENTERPRISE`.
     * Possible values are `ENTERPRISE` and `DEVOPS`.
     */
    readonly tier?: pulumi.Input<string>;
    /**
     * The Type of this CertificateAuthority.
     * > **Note:** For `SUBORDINATE` Certificate Authorities, they need to
     * be manually activated (via Cloud Console of `gcloud`) before they can
     * issue certificates.
     * Default value is `SELF_SIGNED`.
     * Possible values are `SELF_SIGNED` and `SUBORDINATE`.
     */
    readonly type?: pulumi.Input<string>;
    /**
     * The time at which this CertificateAuthority was updated. A timestamp in RFC3339 UTC "Zulu" format, with nanosecond
     * resolution and up to nine fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
     */
    readonly updateTime?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Authority resource.
 */
export interface AuthorityArgs {
    /**
     * The user provided Resource ID for this Certificate Authority.
     */
    readonly certificateAuthorityId: pulumi.Input<string>;
    /**
     * The config used to create a self-signed X.509 certificate or CSR.
     * Structure is documented below.
     */
    readonly config: pulumi.Input<inputs.certificateauthority.AuthorityConfig>;
    /**
     * If set to `true`, the Certificate Authority will be disabled
     * on delete. If the Certitificate Authorities is not disabled,
     * it cannot be deleted. Use with care. Defaults to `false`.
     */
    readonly disableOnDelete?: pulumi.Input<boolean>;
    /**
     * The name of a Cloud Storage bucket where this CertificateAuthority will publish content,
     * such as the CA certificate and CRLs. This must be a bucket name, without any prefixes
     * (such as `gs://`) or suffixes (such as `.googleapis.com`). For example, to use a bucket named
     * my-bucket, you would simply specify `my-bucket`. If not specified, a managed bucket will be
     * created.
     */
    readonly gcsBucket?: pulumi.Input<string>;
    /**
     * Options that affect all certificates issued by a CertificateAuthority.
     * Structure is documented below.
     */
    readonly issuingOptions?: pulumi.Input<inputs.certificateauthority.AuthorityIssuingOptions>;
    /**
     * Used when issuing certificates for this CertificateAuthority. If this CertificateAuthority
     * is a self-signed CertificateAuthority, this key is also used to sign the self-signed CA
     * certificate. Otherwise, it is used to sign a CSR.
     * Structure is documented below.
     */
    readonly keySpec: pulumi.Input<inputs.certificateauthority.AuthorityKeySpec>;
    /**
     * Labels with user-defined metadata.
     * An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass":
     * "1.3kg", "count": "3" }.
     */
    readonly labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The desired lifetime of the CA certificate. Used to create the "notBeforeTime" and
     * "notAfterTime" fields inside an X.509 certificate. A duration in seconds with up to nine
     * fractional digits, terminated by 's'. Example: "3.5s".
     */
    readonly lifetime?: pulumi.Input<string>;
    /**
     * Location of the CertificateAuthority. A full list of valid locations can be found by
     * running `gcloud beta privateca locations list`.
     */
    readonly location: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * The Tier of this CertificateAuthority. `ENTERPRISE` Certificate Authorities track
     * server side certificates issued, and support certificate revocation. For more details,
     * please check the [associated documentation](https://cloud.google.com/certificate-authority-service/docs/tiers).
     * Default value is `ENTERPRISE`.
     * Possible values are `ENTERPRISE` and `DEVOPS`.
     */
    readonly tier?: pulumi.Input<string>;
    /**
     * The Type of this CertificateAuthority.
     * > **Note:** For `SUBORDINATE` Certificate Authorities, they need to
     * be manually activated (via Cloud Console of `gcloud`) before they can
     * issue certificates.
     * Default value is `SELF_SIGNED`.
     * Possible values are `SELF_SIGNED` and `SUBORDINATE`.
     */
    readonly type?: pulumi.Input<string>;
}
