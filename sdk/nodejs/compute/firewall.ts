// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Each network has its own firewall controlling access to and from the
 * instances.
 * 
 * All traffic to instances, even from other instances, is blocked by the
 * firewall unless firewall rules are created to allow it.
 * 
 * The default network has automatically created firewall rules that are
 * shown in default firewall rules. No manually created network has
 * automatically created firewall rules except for a default "allow" rule for
 * outgoing traffic and a default "deny" for incoming traffic. For all
 * networks except the default network, you must create any firewall rules
 * you need.
 * 
 * 
 * To get more information about Firewall, see:
 * 
 * * [API documentation](https://cloud.google.com/compute/docs/reference/v1/firewalls)
 * * How-to Guides
 *     * [Official Documentation](https://cloud.google.com/vpc/docs/firewalls)
 * 
 * ## Example Usage - Firewall Basic
 * 
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 * 
 * const defaultNetwork = new gcp.compute.Network("defaultNetwork", {});
 * const defaultFirewall = new gcp.compute.Firewall("defaultFirewall", {
 *     network: defaultNetwork.name,
 *     allow: [
 *         {
 *             protocol: "icmp",
 *         },
 *         {
 *             protocol: "tcp",
 *             ports: [
 *                 "80",
 *                 "8080",
 *                 "1000-2000",
 *             ],
 *         },
 *     ],
 *     sourceTags: ["web"],
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/compute_firewall.html.markdown.
 */
export class Firewall extends pulumi.CustomResource {
    /**
     * Get an existing Firewall resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FirewallState, opts?: pulumi.CustomResourceOptions): Firewall {
        return new Firewall(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:compute/firewall:Firewall';

    /**
     * Returns true if the given object is an instance of Firewall.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Firewall {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Firewall.__pulumiType;
    }

    /**
     * The list of ALLOW rules specified by this firewall. Each rule
     * specifies a protocol and port-range tuple that describes a permitted
     * connection.  Structure is documented below.
     */
    public readonly allows!: pulumi.Output<outputs.compute.FirewallAllow[] | undefined>;
    /**
     * Creation timestamp in RFC3339 text format.
     */
    public /*out*/ readonly creationTimestamp!: pulumi.Output<string>;
    /**
     * The list of DENY rules specified by this firewall. Each rule specifies
     * a protocol and port-range tuple that describes a denied connection.  Structure is documented below.
     */
    public readonly denies!: pulumi.Output<outputs.compute.FirewallDeny[] | undefined>;
    /**
     * An optional description of this resource. Provide this property when
     * you create the resource.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * If destination ranges are specified, the firewall will apply only to
     * traffic that has destination IP address in these ranges. These ranges
     * must be expressed in CIDR format. Only IPv4 is supported.
     */
    public readonly destinationRanges!: pulumi.Output<string[]>;
    /**
     * Direction of traffic to which this firewall applies; default is
     * INGRESS. Note: For INGRESS traffic, it is NOT supported to specify
     * destinationRanges; For EGRESS traffic, it is NOT supported to specify
     * sourceRanges OR sourceTags.
     */
    public readonly direction!: pulumi.Output<string>;
    /**
     * Denotes whether the firewall rule is disabled, i.e not applied to the
     * network it is associated with. When set to true, the firewall rule is
     * not enforced and the network behaves as if it did not exist. If this
     * is unspecified, the firewall rule will be enabled.
     */
    public readonly disabled!: pulumi.Output<boolean | undefined>;
    /**
     * This field denotes whether to enable logging for a particular
     * firewall rule. If logging is enabled, logs will be exported to
     * Stackdriver.
     */
    public readonly enableLogging!: pulumi.Output<boolean | undefined>;
    /**
     * Name of the resource. Provided by the client when the resource is
     * created. The name must be 1-63 characters long, and comply with
     * RFC1035. Specifically, the name must be 1-63 characters long and match
     * the regular expression `a-z?` which means the
     * first character must be a lowercase letter, and all following
     * characters must be a dash, lowercase letter, or digit, except the last
     * character, which cannot be a dash.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The name or selfLink of the network to attach this firewall to.
     */
    public readonly network!: pulumi.Output<string>;
    /**
     * Priority for this rule. This is an integer between 0 and 65535, both
     * inclusive. When not specified, the value assumed is 1000. Relative
     * priorities determine precedence of conflicting rules. Lower value of
     * priority implies higher precedence (eg, a rule with priority 0 has
     * higher precedence than a rule with priority 1). DENY rules take
     * precedence over ALLOW rules having equal priority.
     */
    public readonly priority!: pulumi.Output<number | undefined>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * The URI of the created resource.
     */
    public /*out*/ readonly selfLink!: pulumi.Output<string>;
    /**
     * If source ranges are specified, the firewall will apply only to
     * traffic that has source IP address in these ranges. These ranges must
     * be expressed in CIDR format. One or both of sourceRanges and
     * sourceTags may be set. If both properties are set, the firewall will
     * apply to traffic that has source IP address within sourceRanges OR the
     * source IP that belongs to a tag listed in the sourceTags property. The
     * connection does not need to match both properties for the firewall to
     * apply. Only IPv4 is supported.
     */
    public readonly sourceRanges!: pulumi.Output<string[]>;
    /**
     * If source service accounts are specified, the firewall will apply only
     * to traffic originating from an instance with a service account in this
     * list. Source service accounts cannot be used to control traffic to an
     * instance's external IP address because service accounts are associated
     * with an instance, not an IP address. sourceRanges can be set at the
     * same time as sourceServiceAccounts. If both are set, the firewall will
     * apply to traffic that has source IP address within sourceRanges OR the
     * source IP belongs to an instance with service account listed in
     * sourceServiceAccount. The connection does not need to match both
     * properties for the firewall to apply. sourceServiceAccounts cannot be
     * used at the same time as sourceTags or targetTags.
     */
    public readonly sourceServiceAccounts!: pulumi.Output<string[] | undefined>;
    /**
     * If source tags are specified, the firewall will apply only to traffic
     * with source IP that belongs to a tag listed in source tags. Source
     * tags cannot be used to control traffic to an instance's external IP
     * address. Because tags are associated with an instance, not an IP
     * address. One or both of sourceRanges and sourceTags may be set. If
     * both properties are set, the firewall will apply to traffic that has
     * source IP address within sourceRanges OR the source IP that belongs to
     * a tag listed in the sourceTags property. The connection does not need
     * to match both properties for the firewall to apply.
     */
    public readonly sourceTags!: pulumi.Output<string[] | undefined>;
    /**
     * A list of service accounts indicating sets of instances located in the
     * network that may make network connections as specified in allowed[].
     * targetServiceAccounts cannot be used at the same time as targetTags or
     * sourceTags. If neither targetServiceAccounts nor targetTags are
     * specified, the firewall rule applies to all instances on the specified
     * network.
     */
    public readonly targetServiceAccounts!: pulumi.Output<string[] | undefined>;
    /**
     * A list of instance tags indicating sets of instances located in the
     * network that may make network connections as specified in allowed[].
     * If no targetTags are specified, the firewall rule applies to all
     * instances on the specified network.
     */
    public readonly targetTags!: pulumi.Output<string[] | undefined>;

    /**
     * Create a Firewall resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FirewallArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FirewallArgs | FirewallState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as FirewallState | undefined;
            inputs["allows"] = state ? state.allows : undefined;
            inputs["creationTimestamp"] = state ? state.creationTimestamp : undefined;
            inputs["denies"] = state ? state.denies : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["destinationRanges"] = state ? state.destinationRanges : undefined;
            inputs["direction"] = state ? state.direction : undefined;
            inputs["disabled"] = state ? state.disabled : undefined;
            inputs["enableLogging"] = state ? state.enableLogging : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["network"] = state ? state.network : undefined;
            inputs["priority"] = state ? state.priority : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["selfLink"] = state ? state.selfLink : undefined;
            inputs["sourceRanges"] = state ? state.sourceRanges : undefined;
            inputs["sourceServiceAccounts"] = state ? state.sourceServiceAccounts : undefined;
            inputs["sourceTags"] = state ? state.sourceTags : undefined;
            inputs["targetServiceAccounts"] = state ? state.targetServiceAccounts : undefined;
            inputs["targetTags"] = state ? state.targetTags : undefined;
        } else {
            const args = argsOrState as FirewallArgs | undefined;
            if (!args || args.network === undefined) {
                throw new Error("Missing required property 'network'");
            }
            inputs["allows"] = args ? args.allows : undefined;
            inputs["denies"] = args ? args.denies : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["destinationRanges"] = args ? args.destinationRanges : undefined;
            inputs["direction"] = args ? args.direction : undefined;
            inputs["disabled"] = args ? args.disabled : undefined;
            inputs["enableLogging"] = args ? args.enableLogging : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["network"] = args ? args.network : undefined;
            inputs["priority"] = args ? args.priority : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["sourceRanges"] = args ? args.sourceRanges : undefined;
            inputs["sourceServiceAccounts"] = args ? args.sourceServiceAccounts : undefined;
            inputs["sourceTags"] = args ? args.sourceTags : undefined;
            inputs["targetServiceAccounts"] = args ? args.targetServiceAccounts : undefined;
            inputs["targetTags"] = args ? args.targetTags : undefined;
            inputs["creationTimestamp"] = undefined /*out*/;
            inputs["selfLink"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Firewall.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Firewall resources.
 */
export interface FirewallState {
    /**
     * The list of ALLOW rules specified by this firewall. Each rule
     * specifies a protocol and port-range tuple that describes a permitted
     * connection.  Structure is documented below.
     */
    readonly allows?: pulumi.Input<pulumi.Input<inputs.compute.FirewallAllow>[]>;
    /**
     * Creation timestamp in RFC3339 text format.
     */
    readonly creationTimestamp?: pulumi.Input<string>;
    /**
     * The list of DENY rules specified by this firewall. Each rule specifies
     * a protocol and port-range tuple that describes a denied connection.  Structure is documented below.
     */
    readonly denies?: pulumi.Input<pulumi.Input<inputs.compute.FirewallDeny>[]>;
    /**
     * An optional description of this resource. Provide this property when
     * you create the resource.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * If destination ranges are specified, the firewall will apply only to
     * traffic that has destination IP address in these ranges. These ranges
     * must be expressed in CIDR format. Only IPv4 is supported.
     */
    readonly destinationRanges?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Direction of traffic to which this firewall applies; default is
     * INGRESS. Note: For INGRESS traffic, it is NOT supported to specify
     * destinationRanges; For EGRESS traffic, it is NOT supported to specify
     * sourceRanges OR sourceTags.
     */
    readonly direction?: pulumi.Input<string>;
    /**
     * Denotes whether the firewall rule is disabled, i.e not applied to the
     * network it is associated with. When set to true, the firewall rule is
     * not enforced and the network behaves as if it did not exist. If this
     * is unspecified, the firewall rule will be enabled.
     */
    readonly disabled?: pulumi.Input<boolean>;
    /**
     * This field denotes whether to enable logging for a particular
     * firewall rule. If logging is enabled, logs will be exported to
     * Stackdriver.
     */
    readonly enableLogging?: pulumi.Input<boolean>;
    /**
     * Name of the resource. Provided by the client when the resource is
     * created. The name must be 1-63 characters long, and comply with
     * RFC1035. Specifically, the name must be 1-63 characters long and match
     * the regular expression `a-z?` which means the
     * first character must be a lowercase letter, and all following
     * characters must be a dash, lowercase letter, or digit, except the last
     * character, which cannot be a dash.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The name or selfLink of the network to attach this firewall to.
     */
    readonly network?: pulumi.Input<string>;
    /**
     * Priority for this rule. This is an integer between 0 and 65535, both
     * inclusive. When not specified, the value assumed is 1000. Relative
     * priorities determine precedence of conflicting rules. Lower value of
     * priority implies higher precedence (eg, a rule with priority 0 has
     * higher precedence than a rule with priority 1). DENY rules take
     * precedence over ALLOW rules having equal priority.
     */
    readonly priority?: pulumi.Input<number>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * The URI of the created resource.
     */
    readonly selfLink?: pulumi.Input<string>;
    /**
     * If source ranges are specified, the firewall will apply only to
     * traffic that has source IP address in these ranges. These ranges must
     * be expressed in CIDR format. One or both of sourceRanges and
     * sourceTags may be set. If both properties are set, the firewall will
     * apply to traffic that has source IP address within sourceRanges OR the
     * source IP that belongs to a tag listed in the sourceTags property. The
     * connection does not need to match both properties for the firewall to
     * apply. Only IPv4 is supported.
     */
    readonly sourceRanges?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * If source service accounts are specified, the firewall will apply only
     * to traffic originating from an instance with a service account in this
     * list. Source service accounts cannot be used to control traffic to an
     * instance's external IP address because service accounts are associated
     * with an instance, not an IP address. sourceRanges can be set at the
     * same time as sourceServiceAccounts. If both are set, the firewall will
     * apply to traffic that has source IP address within sourceRanges OR the
     * source IP belongs to an instance with service account listed in
     * sourceServiceAccount. The connection does not need to match both
     * properties for the firewall to apply. sourceServiceAccounts cannot be
     * used at the same time as sourceTags or targetTags.
     */
    readonly sourceServiceAccounts?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * If source tags are specified, the firewall will apply only to traffic
     * with source IP that belongs to a tag listed in source tags. Source
     * tags cannot be used to control traffic to an instance's external IP
     * address. Because tags are associated with an instance, not an IP
     * address. One or both of sourceRanges and sourceTags may be set. If
     * both properties are set, the firewall will apply to traffic that has
     * source IP address within sourceRanges OR the source IP that belongs to
     * a tag listed in the sourceTags property. The connection does not need
     * to match both properties for the firewall to apply.
     */
    readonly sourceTags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A list of service accounts indicating sets of instances located in the
     * network that may make network connections as specified in allowed[].
     * targetServiceAccounts cannot be used at the same time as targetTags or
     * sourceTags. If neither targetServiceAccounts nor targetTags are
     * specified, the firewall rule applies to all instances on the specified
     * network.
     */
    readonly targetServiceAccounts?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A list of instance tags indicating sets of instances located in the
     * network that may make network connections as specified in allowed[].
     * If no targetTags are specified, the firewall rule applies to all
     * instances on the specified network.
     */
    readonly targetTags?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a Firewall resource.
 */
export interface FirewallArgs {
    /**
     * The list of ALLOW rules specified by this firewall. Each rule
     * specifies a protocol and port-range tuple that describes a permitted
     * connection.  Structure is documented below.
     */
    readonly allows?: pulumi.Input<pulumi.Input<inputs.compute.FirewallAllow>[]>;
    /**
     * The list of DENY rules specified by this firewall. Each rule specifies
     * a protocol and port-range tuple that describes a denied connection.  Structure is documented below.
     */
    readonly denies?: pulumi.Input<pulumi.Input<inputs.compute.FirewallDeny>[]>;
    /**
     * An optional description of this resource. Provide this property when
     * you create the resource.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * If destination ranges are specified, the firewall will apply only to
     * traffic that has destination IP address in these ranges. These ranges
     * must be expressed in CIDR format. Only IPv4 is supported.
     */
    readonly destinationRanges?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Direction of traffic to which this firewall applies; default is
     * INGRESS. Note: For INGRESS traffic, it is NOT supported to specify
     * destinationRanges; For EGRESS traffic, it is NOT supported to specify
     * sourceRanges OR sourceTags.
     */
    readonly direction?: pulumi.Input<string>;
    /**
     * Denotes whether the firewall rule is disabled, i.e not applied to the
     * network it is associated with. When set to true, the firewall rule is
     * not enforced and the network behaves as if it did not exist. If this
     * is unspecified, the firewall rule will be enabled.
     */
    readonly disabled?: pulumi.Input<boolean>;
    /**
     * This field denotes whether to enable logging for a particular
     * firewall rule. If logging is enabled, logs will be exported to
     * Stackdriver.
     */
    readonly enableLogging?: pulumi.Input<boolean>;
    /**
     * Name of the resource. Provided by the client when the resource is
     * created. The name must be 1-63 characters long, and comply with
     * RFC1035. Specifically, the name must be 1-63 characters long and match
     * the regular expression `a-z?` which means the
     * first character must be a lowercase letter, and all following
     * characters must be a dash, lowercase letter, or digit, except the last
     * character, which cannot be a dash.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The name or selfLink of the network to attach this firewall to.
     */
    readonly network: pulumi.Input<string>;
    /**
     * Priority for this rule. This is an integer between 0 and 65535, both
     * inclusive. When not specified, the value assumed is 1000. Relative
     * priorities determine precedence of conflicting rules. Lower value of
     * priority implies higher precedence (eg, a rule with priority 0 has
     * higher precedence than a rule with priority 1). DENY rules take
     * precedence over ALLOW rules having equal priority.
     */
    readonly priority?: pulumi.Input<number>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * If source ranges are specified, the firewall will apply only to
     * traffic that has source IP address in these ranges. These ranges must
     * be expressed in CIDR format. One or both of sourceRanges and
     * sourceTags may be set. If both properties are set, the firewall will
     * apply to traffic that has source IP address within sourceRanges OR the
     * source IP that belongs to a tag listed in the sourceTags property. The
     * connection does not need to match both properties for the firewall to
     * apply. Only IPv4 is supported.
     */
    readonly sourceRanges?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * If source service accounts are specified, the firewall will apply only
     * to traffic originating from an instance with a service account in this
     * list. Source service accounts cannot be used to control traffic to an
     * instance's external IP address because service accounts are associated
     * with an instance, not an IP address. sourceRanges can be set at the
     * same time as sourceServiceAccounts. If both are set, the firewall will
     * apply to traffic that has source IP address within sourceRanges OR the
     * source IP belongs to an instance with service account listed in
     * sourceServiceAccount. The connection does not need to match both
     * properties for the firewall to apply. sourceServiceAccounts cannot be
     * used at the same time as sourceTags or targetTags.
     */
    readonly sourceServiceAccounts?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * If source tags are specified, the firewall will apply only to traffic
     * with source IP that belongs to a tag listed in source tags. Source
     * tags cannot be used to control traffic to an instance's external IP
     * address. Because tags are associated with an instance, not an IP
     * address. One or both of sourceRanges and sourceTags may be set. If
     * both properties are set, the firewall will apply to traffic that has
     * source IP address within sourceRanges OR the source IP that belongs to
     * a tag listed in the sourceTags property. The connection does not need
     * to match both properties for the firewall to apply.
     */
    readonly sourceTags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A list of service accounts indicating sets of instances located in the
     * network that may make network connections as specified in allowed[].
     * targetServiceAccounts cannot be used at the same time as targetTags or
     * sourceTags. If neither targetServiceAccounts nor targetTags are
     * specified, the firewall rule applies to all instances on the specified
     * network.
     */
    readonly targetServiceAccounts?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A list of instance tags indicating sets of instances located in the
     * network that may make network connections as specified in allowed[].
     * If no targetTags are specified, the firewall rule applies to all
     * instances on the specified network.
     */
    readonly targetTags?: pulumi.Input<pulumi.Input<string>[]>;
}
