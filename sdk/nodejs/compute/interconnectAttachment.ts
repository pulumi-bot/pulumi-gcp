// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Represents an InterconnectAttachment (VLAN attachment) resource. For more
 * information, see Creating VLAN Attachments.
 *
 * ## Example Usage
 *
 * ## Import
 *
 * InterconnectAttachment can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import gcp:compute/interconnectAttachment:InterconnectAttachment default projects/{{project}}/regions/{{region}}/interconnectAttachments/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:compute/interconnectAttachment:InterconnectAttachment default {{project}}/{{region}}/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:compute/interconnectAttachment:InterconnectAttachment default {{region}}/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:compute/interconnectAttachment:InterconnectAttachment default {{name}}
 * ```
 */
export class InterconnectAttachment extends pulumi.CustomResource {
    /**
     * Get an existing InterconnectAttachment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: InterconnectAttachmentState, opts?: pulumi.CustomResourceOptions): InterconnectAttachment {
        return new InterconnectAttachment(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:compute/interconnectAttachment:InterconnectAttachment';

    /**
     * Returns true if the given object is an instance of InterconnectAttachment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is InterconnectAttachment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === InterconnectAttachment.__pulumiType;
    }

    /**
     * Whether the VLAN attachment is enabled or disabled.  When using
     * PARTNER type this will Pre-Activate the interconnect attachment
     */
    public readonly adminEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * Provisioned bandwidth capacity for the interconnect attachment.
     * For attachments of type DEDICATED, the user can set the bandwidth.
     * For attachments of type PARTNER, the Google Partner that is operating the interconnect must set the bandwidth.
     * Output only for PARTNER type, mutable for PARTNER_PROVIDER and DEDICATED,
     * Defaults to BPS_10G
     * Possible values are `BPS_50M`, `BPS_100M`, `BPS_200M`, `BPS_300M`, `BPS_400M`, `BPS_500M`, `BPS_1G`, `BPS_2G`, `BPS_5G`, `BPS_10G`, `BPS_20G`, and `BPS_50G`.
     */
    public readonly bandwidth!: pulumi.Output<string>;
    /**
     * Up to 16 candidate prefixes that can be used to restrict the allocation
     * of cloudRouterIpAddress and customerRouterIpAddress for this attachment.
     * All prefixes must be within link-local address space (169.254.0.0/16)
     * and must be /29 or shorter (/28, /27, etc). Google will attempt to select
     * an unused /29 from the supplied candidate prefix(es). The request will
     * fail if all possible /29s are in use on Google's edge. If not supplied,
     * Google will randomly select an unused /29 from all of link-local space.
     */
    public readonly candidateSubnets!: pulumi.Output<string[] | undefined>;
    /**
     * IPv4 address + prefix length to be configured on Cloud Router Interface for this interconnect attachment.
     */
    public /*out*/ readonly cloudRouterIpAddress!: pulumi.Output<string>;
    /**
     * Creation timestamp in RFC3339 text format.
     */
    public /*out*/ readonly creationTimestamp!: pulumi.Output<string>;
    /**
     * IPv4 address + prefix length to be configured on the customer router subinterface for this interconnect attachment.
     */
    public /*out*/ readonly customerRouterIpAddress!: pulumi.Output<string>;
    /**
     * An optional description of this resource.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Desired availability domain for the attachment. Only available for type
     * PARTNER, at creation time. For improved reliability, customers should
     * configure a pair of attachments with one per availability domain. The
     * selected availability domain will be provided to the Partner via the
     * pairing key so that the provisioned circuit will lie in the specified
     * domain. If not specified, the value will default to AVAILABILITY_DOMAIN_ANY.
     */
    public readonly edgeAvailabilityDomain!: pulumi.Output<string>;
    /**
     * Google reference ID, to be used when raising support tickets with Google or otherwise to debug backend connectivity
     * issues.
     */
    public /*out*/ readonly googleReferenceId!: pulumi.Output<string>;
    /**
     * URL of the underlying Interconnect object that this attachment's
     * traffic will traverse through. Required if type is DEDICATED, must not
     * be set if type is PARTNER.
     */
    public readonly interconnect!: pulumi.Output<string | undefined>;
    /**
     * Name of the resource. Provided by the client when the resource is created. The
     * name must be 1-63 characters long, and comply with RFC1035. Specifically, the
     * name must be 1-63 characters long and match the regular expression
     * `a-z?` which means the first character must be a
     * lowercase letter, and all following characters must be a dash, lowercase
     * letter, or digit, except the last character, which cannot be a dash.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * [Output only for type PARTNER. Not present for DEDICATED]. The opaque identifier of an PARTNER attachment used to
     * initiate provisioning with a selected partner. Of the form "XXXXX/region/domain"
     */
    public /*out*/ readonly pairingKey!: pulumi.Output<string>;
    /**
     * [Output only for type PARTNER. Not present for DEDICATED]. Optional BGP ASN for the router that should be supplied by a
     * layer 3 Partner if they configured BGP on behalf of the customer.
     */
    public /*out*/ readonly partnerAsn!: pulumi.Output<string>;
    /**
     * Information specific to an InterconnectAttachment. This property is populated if the interconnect that this is attached
     * to is of type DEDICATED.
     */
    public /*out*/ readonly privateInterconnectInfos!: pulumi.Output<outputs.compute.InterconnectAttachmentPrivateInterconnectInfo[]>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * Region where the regional interconnect attachment resides.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * URL of the cloud router to be used for dynamic routing. This router must be in
     * the same region as this InterconnectAttachment. The InterconnectAttachment will
     * automatically connect the Interconnect to the network & region within which the
     * Cloud Router is configured.
     */
    public readonly router!: pulumi.Output<string>;
    /**
     * The URI of the created resource.
     */
    public /*out*/ readonly selfLink!: pulumi.Output<string>;
    /**
     * [Output Only] The current state of this attachment's functionality.
     */
    public /*out*/ readonly state!: pulumi.Output<string>;
    /**
     * The type of InterconnectAttachment you wish to create. Defaults to
     * DEDICATED.
     * Possible values are `DEDICATED`, `PARTNER`, and `PARTNER_PROVIDER`.
     */
    public readonly type!: pulumi.Output<string>;
    /**
     * The IEEE 802.1Q VLAN tag for this attachment, in the range 2-4094. When
     * using PARTNER type this will be managed upstream.
     */
    public readonly vlanTag8021q!: pulumi.Output<number>;

    /**
     * Create a InterconnectAttachment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: InterconnectAttachmentArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: InterconnectAttachmentArgs | InterconnectAttachmentState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as InterconnectAttachmentState | undefined;
            inputs["adminEnabled"] = state ? state.adminEnabled : undefined;
            inputs["bandwidth"] = state ? state.bandwidth : undefined;
            inputs["candidateSubnets"] = state ? state.candidateSubnets : undefined;
            inputs["cloudRouterIpAddress"] = state ? state.cloudRouterIpAddress : undefined;
            inputs["creationTimestamp"] = state ? state.creationTimestamp : undefined;
            inputs["customerRouterIpAddress"] = state ? state.customerRouterIpAddress : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["edgeAvailabilityDomain"] = state ? state.edgeAvailabilityDomain : undefined;
            inputs["googleReferenceId"] = state ? state.googleReferenceId : undefined;
            inputs["interconnect"] = state ? state.interconnect : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["pairingKey"] = state ? state.pairingKey : undefined;
            inputs["partnerAsn"] = state ? state.partnerAsn : undefined;
            inputs["privateInterconnectInfos"] = state ? state.privateInterconnectInfos : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["region"] = state ? state.region : undefined;
            inputs["router"] = state ? state.router : undefined;
            inputs["selfLink"] = state ? state.selfLink : undefined;
            inputs["state"] = state ? state.state : undefined;
            inputs["type"] = state ? state.type : undefined;
            inputs["vlanTag8021q"] = state ? state.vlanTag8021q : undefined;
        } else {
            const args = argsOrState as InterconnectAttachmentArgs | undefined;
            if (!args || args.router === undefined) {
                throw new Error("Missing required property 'router'");
            }
            inputs["adminEnabled"] = args ? args.adminEnabled : undefined;
            inputs["bandwidth"] = args ? args.bandwidth : undefined;
            inputs["candidateSubnets"] = args ? args.candidateSubnets : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["edgeAvailabilityDomain"] = args ? args.edgeAvailabilityDomain : undefined;
            inputs["interconnect"] = args ? args.interconnect : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["region"] = args ? args.region : undefined;
            inputs["router"] = args ? args.router : undefined;
            inputs["type"] = args ? args.type : undefined;
            inputs["vlanTag8021q"] = args ? args.vlanTag8021q : undefined;
            inputs["cloudRouterIpAddress"] = undefined /*out*/;
            inputs["creationTimestamp"] = undefined /*out*/;
            inputs["customerRouterIpAddress"] = undefined /*out*/;
            inputs["googleReferenceId"] = undefined /*out*/;
            inputs["pairingKey"] = undefined /*out*/;
            inputs["partnerAsn"] = undefined /*out*/;
            inputs["privateInterconnectInfos"] = undefined /*out*/;
            inputs["selfLink"] = undefined /*out*/;
            inputs["state"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(InterconnectAttachment.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering InterconnectAttachment resources.
 */
export interface InterconnectAttachmentState {
    /**
     * Whether the VLAN attachment is enabled or disabled.  When using
     * PARTNER type this will Pre-Activate the interconnect attachment
     */
    readonly adminEnabled?: pulumi.Input<boolean>;
    /**
     * Provisioned bandwidth capacity for the interconnect attachment.
     * For attachments of type DEDICATED, the user can set the bandwidth.
     * For attachments of type PARTNER, the Google Partner that is operating the interconnect must set the bandwidth.
     * Output only for PARTNER type, mutable for PARTNER_PROVIDER and DEDICATED,
     * Defaults to BPS_10G
     * Possible values are `BPS_50M`, `BPS_100M`, `BPS_200M`, `BPS_300M`, `BPS_400M`, `BPS_500M`, `BPS_1G`, `BPS_2G`, `BPS_5G`, `BPS_10G`, `BPS_20G`, and `BPS_50G`.
     */
    readonly bandwidth?: pulumi.Input<string>;
    /**
     * Up to 16 candidate prefixes that can be used to restrict the allocation
     * of cloudRouterIpAddress and customerRouterIpAddress for this attachment.
     * All prefixes must be within link-local address space (169.254.0.0/16)
     * and must be /29 or shorter (/28, /27, etc). Google will attempt to select
     * an unused /29 from the supplied candidate prefix(es). The request will
     * fail if all possible /29s are in use on Google's edge. If not supplied,
     * Google will randomly select an unused /29 from all of link-local space.
     */
    readonly candidateSubnets?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * IPv4 address + prefix length to be configured on Cloud Router Interface for this interconnect attachment.
     */
    readonly cloudRouterIpAddress?: pulumi.Input<string>;
    /**
     * Creation timestamp in RFC3339 text format.
     */
    readonly creationTimestamp?: pulumi.Input<string>;
    /**
     * IPv4 address + prefix length to be configured on the customer router subinterface for this interconnect attachment.
     */
    readonly customerRouterIpAddress?: pulumi.Input<string>;
    /**
     * An optional description of this resource.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Desired availability domain for the attachment. Only available for type
     * PARTNER, at creation time. For improved reliability, customers should
     * configure a pair of attachments with one per availability domain. The
     * selected availability domain will be provided to the Partner via the
     * pairing key so that the provisioned circuit will lie in the specified
     * domain. If not specified, the value will default to AVAILABILITY_DOMAIN_ANY.
     */
    readonly edgeAvailabilityDomain?: pulumi.Input<string>;
    /**
     * Google reference ID, to be used when raising support tickets with Google or otherwise to debug backend connectivity
     * issues.
     */
    readonly googleReferenceId?: pulumi.Input<string>;
    /**
     * URL of the underlying Interconnect object that this attachment's
     * traffic will traverse through. Required if type is DEDICATED, must not
     * be set if type is PARTNER.
     */
    readonly interconnect?: pulumi.Input<string>;
    /**
     * Name of the resource. Provided by the client when the resource is created. The
     * name must be 1-63 characters long, and comply with RFC1035. Specifically, the
     * name must be 1-63 characters long and match the regular expression
     * `a-z?` which means the first character must be a
     * lowercase letter, and all following characters must be a dash, lowercase
     * letter, or digit, except the last character, which cannot be a dash.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * [Output only for type PARTNER. Not present for DEDICATED]. The opaque identifier of an PARTNER attachment used to
     * initiate provisioning with a selected partner. Of the form "XXXXX/region/domain"
     */
    readonly pairingKey?: pulumi.Input<string>;
    /**
     * [Output only for type PARTNER. Not present for DEDICATED]. Optional BGP ASN for the router that should be supplied by a
     * layer 3 Partner if they configured BGP on behalf of the customer.
     */
    readonly partnerAsn?: pulumi.Input<string>;
    /**
     * Information specific to an InterconnectAttachment. This property is populated if the interconnect that this is attached
     * to is of type DEDICATED.
     */
    readonly privateInterconnectInfos?: pulumi.Input<pulumi.Input<inputs.compute.InterconnectAttachmentPrivateInterconnectInfo>[]>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * Region where the regional interconnect attachment resides.
     */
    readonly region?: pulumi.Input<string>;
    /**
     * URL of the cloud router to be used for dynamic routing. This router must be in
     * the same region as this InterconnectAttachment. The InterconnectAttachment will
     * automatically connect the Interconnect to the network & region within which the
     * Cloud Router is configured.
     */
    readonly router?: pulumi.Input<string>;
    /**
     * The URI of the created resource.
     */
    readonly selfLink?: pulumi.Input<string>;
    /**
     * [Output Only] The current state of this attachment's functionality.
     */
    readonly state?: pulumi.Input<string>;
    /**
     * The type of InterconnectAttachment you wish to create. Defaults to
     * DEDICATED.
     * Possible values are `DEDICATED`, `PARTNER`, and `PARTNER_PROVIDER`.
     */
    readonly type?: pulumi.Input<string>;
    /**
     * The IEEE 802.1Q VLAN tag for this attachment, in the range 2-4094. When
     * using PARTNER type this will be managed upstream.
     */
    readonly vlanTag8021q?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a InterconnectAttachment resource.
 */
export interface InterconnectAttachmentArgs {
    /**
     * Whether the VLAN attachment is enabled or disabled.  When using
     * PARTNER type this will Pre-Activate the interconnect attachment
     */
    readonly adminEnabled?: pulumi.Input<boolean>;
    /**
     * Provisioned bandwidth capacity for the interconnect attachment.
     * For attachments of type DEDICATED, the user can set the bandwidth.
     * For attachments of type PARTNER, the Google Partner that is operating the interconnect must set the bandwidth.
     * Output only for PARTNER type, mutable for PARTNER_PROVIDER and DEDICATED,
     * Defaults to BPS_10G
     * Possible values are `BPS_50M`, `BPS_100M`, `BPS_200M`, `BPS_300M`, `BPS_400M`, `BPS_500M`, `BPS_1G`, `BPS_2G`, `BPS_5G`, `BPS_10G`, `BPS_20G`, and `BPS_50G`.
     */
    readonly bandwidth?: pulumi.Input<string>;
    /**
     * Up to 16 candidate prefixes that can be used to restrict the allocation
     * of cloudRouterIpAddress and customerRouterIpAddress for this attachment.
     * All prefixes must be within link-local address space (169.254.0.0/16)
     * and must be /29 or shorter (/28, /27, etc). Google will attempt to select
     * an unused /29 from the supplied candidate prefix(es). The request will
     * fail if all possible /29s are in use on Google's edge. If not supplied,
     * Google will randomly select an unused /29 from all of link-local space.
     */
    readonly candidateSubnets?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * An optional description of this resource.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Desired availability domain for the attachment. Only available for type
     * PARTNER, at creation time. For improved reliability, customers should
     * configure a pair of attachments with one per availability domain. The
     * selected availability domain will be provided to the Partner via the
     * pairing key so that the provisioned circuit will lie in the specified
     * domain. If not specified, the value will default to AVAILABILITY_DOMAIN_ANY.
     */
    readonly edgeAvailabilityDomain?: pulumi.Input<string>;
    /**
     * URL of the underlying Interconnect object that this attachment's
     * traffic will traverse through. Required if type is DEDICATED, must not
     * be set if type is PARTNER.
     */
    readonly interconnect?: pulumi.Input<string>;
    /**
     * Name of the resource. Provided by the client when the resource is created. The
     * name must be 1-63 characters long, and comply with RFC1035. Specifically, the
     * name must be 1-63 characters long and match the regular expression
     * `a-z?` which means the first character must be a
     * lowercase letter, and all following characters must be a dash, lowercase
     * letter, or digit, except the last character, which cannot be a dash.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * Region where the regional interconnect attachment resides.
     */
    readonly region?: pulumi.Input<string>;
    /**
     * URL of the cloud router to be used for dynamic routing. This router must be in
     * the same region as this InterconnectAttachment. The InterconnectAttachment will
     * automatically connect the Interconnect to the network & region within which the
     * Cloud Router is configured.
     */
    readonly router: pulumi.Input<string>;
    /**
     * The type of InterconnectAttachment you wish to create. Defaults to
     * DEDICATED.
     * Possible values are `DEDICATED`, `PARTNER`, and `PARTNER_PROVIDER`.
     */
    readonly type?: pulumi.Input<string>;
    /**
     * The IEEE 802.1Q VLAN tag for this attachment, in the range 2-4094. When
     * using PARTNER type this will be managed upstream.
     */
    readonly vlanTag8021q?: pulumi.Input<number>;
}
