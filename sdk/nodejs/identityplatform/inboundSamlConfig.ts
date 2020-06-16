// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Inbound SAML configuration for a Identity Toolkit project.
 *
 * You must enable the
 * [Google Identity Platform](https://console.cloud.google.com/marketplace/details/google-cloud-platform/customer-identity) in
 * the marketplace prior to using this resource.
 *
 * {{% examples %}}
 * ## Example Usage
 * {{% example %}}
 * ### Identity Platform Inbound Saml Config Basic
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 * import * from "fs";
 *
 * const samlConfig = new gcp.identityplatform.InboundSamlConfig("samlConfig", {
 *     displayName: "Display Name",
 *     idp_config: {
 *         idpEntityId: "tf-idp",
 *         signRequest: true,
 *         ssoUrl: "https://example.com",
 *         idp_certificates: [{
 *             x509Certificate: fs.readFileSync("test-fixtures/rsa_cert.pem"),
 *         }],
 *     },
 *     sp_config: {
 *         spEntityId: "tf-sp",
 *         callbackUri: "https://example.com",
 *     },
 * });
 * ```
 * {{% /example %}}
 * {{% /examples %}}
 */
export class InboundSamlConfig extends pulumi.CustomResource {
    /**
     * Get an existing InboundSamlConfig resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: InboundSamlConfigState, opts?: pulumi.CustomResourceOptions): InboundSamlConfig {
        return new InboundSamlConfig(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:identityplatform/inboundSamlConfig:InboundSamlConfig';

    /**
     * Returns true if the given object is an instance of InboundSamlConfig.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is InboundSamlConfig {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === InboundSamlConfig.__pulumiType;
    }

    /**
     * Human friendly display name.
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * If this config allows users to sign in with the provider.
     */
    public readonly enabled!: pulumi.Output<boolean | undefined>;
    /**
     * SAML IdP configuration when the project acts as the relying party  Structure is documented below.
     */
    public readonly idpConfig!: pulumi.Output<outputs.identityplatform.InboundSamlConfigIdpConfig>;
    /**
     * The name of the InboundSamlConfig resource. Must start with 'saml.' and can only have alphanumeric characters,
     * hyphens, underscores or periods. The part after 'saml.' must also start with a lowercase letter, end with an
     * alphanumeric character, and have at least 2 characters.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * SAML SP (Service Provider) configuration when the project acts as the relying party to receive
     * and accept an authentication assertion issued by a SAML identity provider.  Structure is documented below.
     */
    public readonly spConfig!: pulumi.Output<outputs.identityplatform.InboundSamlConfigSpConfig>;

    /**
     * Create a InboundSamlConfig resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: InboundSamlConfigArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: InboundSamlConfigArgs | InboundSamlConfigState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as InboundSamlConfigState | undefined;
            inputs["displayName"] = state ? state.displayName : undefined;
            inputs["enabled"] = state ? state.enabled : undefined;
            inputs["idpConfig"] = state ? state.idpConfig : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["spConfig"] = state ? state.spConfig : undefined;
        } else {
            const args = argsOrState as InboundSamlConfigArgs | undefined;
            if (!args || args.displayName === undefined) {
                throw new Error("Missing required property 'displayName'");
            }
            if (!args || args.idpConfig === undefined) {
                throw new Error("Missing required property 'idpConfig'");
            }
            if (!args || args.spConfig === undefined) {
                throw new Error("Missing required property 'spConfig'");
            }
            inputs["displayName"] = args ? args.displayName : undefined;
            inputs["enabled"] = args ? args.enabled : undefined;
            inputs["idpConfig"] = args ? args.idpConfig : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["spConfig"] = args ? args.spConfig : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(InboundSamlConfig.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering InboundSamlConfig resources.
 */
export interface InboundSamlConfigState {
    /**
     * Human friendly display name.
     */
    readonly displayName?: pulumi.Input<string>;
    /**
     * If this config allows users to sign in with the provider.
     */
    readonly enabled?: pulumi.Input<boolean>;
    /**
     * SAML IdP configuration when the project acts as the relying party  Structure is documented below.
     */
    readonly idpConfig?: pulumi.Input<inputs.identityplatform.InboundSamlConfigIdpConfig>;
    /**
     * The name of the InboundSamlConfig resource. Must start with 'saml.' and can only have alphanumeric characters,
     * hyphens, underscores or periods. The part after 'saml.' must also start with a lowercase letter, end with an
     * alphanumeric character, and have at least 2 characters.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * SAML SP (Service Provider) configuration when the project acts as the relying party to receive
     * and accept an authentication assertion issued by a SAML identity provider.  Structure is documented below.
     */
    readonly spConfig?: pulumi.Input<inputs.identityplatform.InboundSamlConfigSpConfig>;
}

/**
 * The set of arguments for constructing a InboundSamlConfig resource.
 */
export interface InboundSamlConfigArgs {
    /**
     * Human friendly display name.
     */
    readonly displayName: pulumi.Input<string>;
    /**
     * If this config allows users to sign in with the provider.
     */
    readonly enabled?: pulumi.Input<boolean>;
    /**
     * SAML IdP configuration when the project acts as the relying party  Structure is documented below.
     */
    readonly idpConfig: pulumi.Input<inputs.identityplatform.InboundSamlConfigIdpConfig>;
    /**
     * The name of the InboundSamlConfig resource. Must start with 'saml.' and can only have alphanumeric characters,
     * hyphens, underscores or periods. The part after 'saml.' must also start with a lowercase letter, end with an
     * alphanumeric character, and have at least 2 characters.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * SAML SP (Service Provider) configuration when the project acts as the relying party to receive
     * and accept an authentication assertion issued by a SAML identity provider.  Structure is documented below.
     */
    readonly spConfig: pulumi.Input<inputs.identityplatform.InboundSamlConfigSpConfig>;
}
