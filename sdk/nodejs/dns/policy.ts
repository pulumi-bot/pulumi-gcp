// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/dns_policy.html.markdown.
 */
export class Policy extends pulumi.CustomResource {
    /**
     * Get an existing Policy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PolicyState, opts?: pulumi.CustomResourceOptions): Policy {
        return new Policy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:dns/policy:Policy';

    /**
     * Returns true if the given object is an instance of Policy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Policy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Policy.__pulumiType;
    }

    public readonly alternativeNameServerConfig!: pulumi.Output<{ targetNameServers?: { ipv4Address?: string }[] } | undefined>;
    public readonly description!: pulumi.Output<string | undefined>;
    public readonly enableInboundForwarding!: pulumi.Output<boolean | undefined>;
    public readonly enableLogging!: pulumi.Output<boolean | undefined>;
    public readonly name!: pulumi.Output<string>;
    public readonly networks!: pulumi.Output<{ networkUrl?: string }[] | undefined>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;

    /**
     * Create a Policy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: PolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PolicyArgs | PolicyState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as PolicyState | undefined;
            inputs["alternativeNameServerConfig"] = state ? state.alternativeNameServerConfig : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["enableInboundForwarding"] = state ? state.enableInboundForwarding : undefined;
            inputs["enableLogging"] = state ? state.enableLogging : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["networks"] = state ? state.networks : undefined;
            inputs["project"] = state ? state.project : undefined;
        } else {
            const args = argsOrState as PolicyArgs | undefined;
            inputs["alternativeNameServerConfig"] = args ? args.alternativeNameServerConfig : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["enableInboundForwarding"] = args ? args.enableInboundForwarding : undefined;
            inputs["enableLogging"] = args ? args.enableLogging : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["networks"] = args ? args.networks : undefined;
            inputs["project"] = args ? args.project : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Policy.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Policy resources.
 */
export interface PolicyState {
    readonly alternativeNameServerConfig?: pulumi.Input<{ targetNameServers?: pulumi.Input<pulumi.Input<{ ipv4Address?: pulumi.Input<string> }>[]> }>;
    readonly description?: pulumi.Input<string>;
    readonly enableInboundForwarding?: pulumi.Input<boolean>;
    readonly enableLogging?: pulumi.Input<boolean>;
    readonly name?: pulumi.Input<string>;
    readonly networks?: pulumi.Input<pulumi.Input<{ networkUrl?: pulumi.Input<string> }>[]>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Policy resource.
 */
export interface PolicyArgs {
    readonly alternativeNameServerConfig?: pulumi.Input<{ targetNameServers?: pulumi.Input<pulumi.Input<{ ipv4Address?: pulumi.Input<string> }>[]> }>;
    readonly description?: pulumi.Input<string>;
    readonly enableInboundForwarding?: pulumi.Input<boolean>;
    readonly enableLogging?: pulumi.Input<boolean>;
    readonly name?: pulumi.Input<string>;
    readonly networks?: pulumi.Input<pulumi.Input<{ networkUrl?: pulumi.Input<string> }>[]>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
}
