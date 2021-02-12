// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Adds a trust between Active Directory domains
 *
 * To get more information about DomainTrust, see:
 *
 * * [API documentation](https://cloud.google.com/managed-microsoft-ad/reference/rest/v1/projects.locations.global.domains/attachTrust)
 * * How-to Guides
 *     * [Active Directory Trust](https://cloud.google.com/managed-microsoft-ad/docs/create-one-way-trust)
 *
 * > **Warning:** All arguments including `trustHandshakeSecret` will be stored in the raw
 * state as plain-text. [Read more about sensitive data in state](https://www.terraform.io/docs/state/sensitive-data.html).
 *
 * ## Example Usage
 * ### Active Directory Domain Trust Basic
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const ad_domain_trust = new gcp.activedirectory.DomainTrust("ad-domain-trust", {
 *     domain: "test-managed-ad.com",
 *     targetDnsIpAddresses: ["10.1.0.100"],
 *     targetDomainName: "example-gcp.com",
 *     trustDirection: "OUTBOUND",
 *     trustHandshakeSecret: "Testing1!",
 *     trustType: "FOREST",
 * });
 * ```
 *
 * ## Import
 *
 * DomainTrust can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import gcp:activedirectory/domainTrust:DomainTrust default projects/{{project}}/locations/global/domains/{{domain}}/{{target_domain_name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:activedirectory/domainTrust:DomainTrust default {{project}}/{{domain}}/{{target_domain_name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:activedirectory/domainTrust:DomainTrust default {{domain}}/{{target_domain_name}}
 * ```
 */
export class DomainTrust extends pulumi.CustomResource {
    /**
     * Get an existing DomainTrust resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DomainTrustState, opts?: pulumi.CustomResourceOptions): DomainTrust {
        return new DomainTrust(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:activedirectory/domainTrust:DomainTrust';

    /**
     * Returns true if the given object is an instance of DomainTrust.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DomainTrust {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DomainTrust.__pulumiType;
    }

    /**
     * The fully qualified domain name. e.g. mydomain.myorganization.com, with the restrictions,
     * https://cloud.google.com/managed-microsoft-ad/reference/rest/v1/projects.locations.global.domains.
     */
    public readonly domain!: pulumi.Output<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * Whether the trusted side has forest/domain wide access or selective access to an approved set of resources.
     */
    public readonly selectiveAuthentication!: pulumi.Output<boolean | undefined>;
    /**
     * The target DNS server IP addresses which can resolve the remote domain involved in the trust.
     */
    public readonly targetDnsIpAddresses!: pulumi.Output<string[]>;
    /**
     * The fully qualified target domain name which will be in trust with the current domain.
     */
    public readonly targetDomainName!: pulumi.Output<string>;
    /**
     * The trust direction, which decides if the current domain is trusted, trusting, or both.
     * Possible values are `INBOUND`, `OUTBOUND`, and `BIDIRECTIONAL`.
     */
    public readonly trustDirection!: pulumi.Output<string>;
    /**
     * The trust secret used for the handshake with the target domain. This will not be stored.
     * **Note**: This property is sensitive and will not be displayed in the plan.
     */
    public readonly trustHandshakeSecret!: pulumi.Output<string>;
    /**
     * The type of trust represented by the trust resource.
     * Possible values are `FOREST` and `EXTERNAL`.
     */
    public readonly trustType!: pulumi.Output<string>;

    /**
     * Create a DomainTrust resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DomainTrustArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DomainTrustArgs | DomainTrustState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DomainTrustState | undefined;
            inputs["domain"] = state ? state.domain : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["selectiveAuthentication"] = state ? state.selectiveAuthentication : undefined;
            inputs["targetDnsIpAddresses"] = state ? state.targetDnsIpAddresses : undefined;
            inputs["targetDomainName"] = state ? state.targetDomainName : undefined;
            inputs["trustDirection"] = state ? state.trustDirection : undefined;
            inputs["trustHandshakeSecret"] = state ? state.trustHandshakeSecret : undefined;
            inputs["trustType"] = state ? state.trustType : undefined;
        } else {
            const args = argsOrState as DomainTrustArgs | undefined;
            if ((!args || args.domain === undefined) && !opts.urn) {
                throw new Error("Missing required property 'domain'");
            }
            if ((!args || args.targetDnsIpAddresses === undefined) && !opts.urn) {
                throw new Error("Missing required property 'targetDnsIpAddresses'");
            }
            if ((!args || args.targetDomainName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'targetDomainName'");
            }
            if ((!args || args.trustDirection === undefined) && !opts.urn) {
                throw new Error("Missing required property 'trustDirection'");
            }
            if ((!args || args.trustHandshakeSecret === undefined) && !opts.urn) {
                throw new Error("Missing required property 'trustHandshakeSecret'");
            }
            if ((!args || args.trustType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'trustType'");
            }
            inputs["domain"] = args ? args.domain : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["selectiveAuthentication"] = args ? args.selectiveAuthentication : undefined;
            inputs["targetDnsIpAddresses"] = args ? args.targetDnsIpAddresses : undefined;
            inputs["targetDomainName"] = args ? args.targetDomainName : undefined;
            inputs["trustDirection"] = args ? args.trustDirection : undefined;
            inputs["trustHandshakeSecret"] = args ? args.trustHandshakeSecret : undefined;
            inputs["trustType"] = args ? args.trustType : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(DomainTrust.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DomainTrust resources.
 */
export interface DomainTrustState {
    /**
     * The fully qualified domain name. e.g. mydomain.myorganization.com, with the restrictions,
     * https://cloud.google.com/managed-microsoft-ad/reference/rest/v1/projects.locations.global.domains.
     */
    readonly domain?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * Whether the trusted side has forest/domain wide access or selective access to an approved set of resources.
     */
    readonly selectiveAuthentication?: pulumi.Input<boolean>;
    /**
     * The target DNS server IP addresses which can resolve the remote domain involved in the trust.
     */
    readonly targetDnsIpAddresses?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The fully qualified target domain name which will be in trust with the current domain.
     */
    readonly targetDomainName?: pulumi.Input<string>;
    /**
     * The trust direction, which decides if the current domain is trusted, trusting, or both.
     * Possible values are `INBOUND`, `OUTBOUND`, and `BIDIRECTIONAL`.
     */
    readonly trustDirection?: pulumi.Input<string>;
    /**
     * The trust secret used for the handshake with the target domain. This will not be stored.
     * **Note**: This property is sensitive and will not be displayed in the plan.
     */
    readonly trustHandshakeSecret?: pulumi.Input<string>;
    /**
     * The type of trust represented by the trust resource.
     * Possible values are `FOREST` and `EXTERNAL`.
     */
    readonly trustType?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a DomainTrust resource.
 */
export interface DomainTrustArgs {
    /**
     * The fully qualified domain name. e.g. mydomain.myorganization.com, with the restrictions,
     * https://cloud.google.com/managed-microsoft-ad/reference/rest/v1/projects.locations.global.domains.
     */
    readonly domain: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * Whether the trusted side has forest/domain wide access or selective access to an approved set of resources.
     */
    readonly selectiveAuthentication?: pulumi.Input<boolean>;
    /**
     * The target DNS server IP addresses which can resolve the remote domain involved in the trust.
     */
    readonly targetDnsIpAddresses: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The fully qualified target domain name which will be in trust with the current domain.
     */
    readonly targetDomainName: pulumi.Input<string>;
    /**
     * The trust direction, which decides if the current domain is trusted, trusting, or both.
     * Possible values are `INBOUND`, `OUTBOUND`, and `BIDIRECTIONAL`.
     */
    readonly trustDirection: pulumi.Input<string>;
    /**
     * The trust secret used for the handshake with the target domain. This will not be stored.
     * **Note**: This property is sensitive and will not be displayed in the plan.
     */
    readonly trustHandshakeSecret: pulumi.Input<string>;
    /**
     * The type of trust represented by the trust resource.
     * Possible values are `FOREST` and `EXTERNAL`.
     */
    readonly trustType: pulumi.Input<string>;
}
