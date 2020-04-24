// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Serverless VPC Access connector resource.
 * 
 * 
 * To get more information about Connector, see:
 * 
 * * [API documentation](https://cloud.google.com/vpc/docs/reference/vpcaccess/rest/v1/projects.locations.connectors)
 * * How-to Guides
 *     * [Configuring Serverless VPC Access](https://cloud.google.com/vpc/docs/configure-serverless-vpc-access)
 * 
 * ## Example Usage - VPC Access Connector
 * 
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 * 
 * const connector = new gcp.vpcaccess.Connector("connector", {
 *     ipCidrRange: "10.8.0.0/28",
 *     network: "default",
 *     region: "us-central1",
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/vpc_access_connector.html.markdown.
 */
export class Connector extends pulumi.CustomResource {
    /**
     * Get an existing Connector resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ConnectorState, opts?: pulumi.CustomResourceOptions): Connector {
        return new Connector(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:vpcaccess/connector:Connector';

    /**
     * Returns true if the given object is an instance of Connector.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Connector {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Connector.__pulumiType;
    }

    /**
     * The range of internal addresses that follows RFC 4632 notation. Example: `10.132.0.0/28`.
     */
    public readonly ipCidrRange!: pulumi.Output<string>;
    /**
     * Maximum throughput of the connector in Mbps, must be greater than `minThroughput`. Default is 1000.
     */
    public readonly maxThroughput!: pulumi.Output<number | undefined>;
    /**
     * Minimum throughput of the connector in Mbps. Default and min is 200.
     */
    public readonly minThroughput!: pulumi.Output<number | undefined>;
    /**
     * The name of the resource (Max 25 characters).
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Name of a VPC network.
     */
    public readonly network!: pulumi.Output<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * Region where the VPC Access connector resides
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * The fully qualified name of this VPC connector
     */
    public /*out*/ readonly selfLink!: pulumi.Output<string>;
    /**
     * State of the VPC access connector.
     */
    public /*out*/ readonly state!: pulumi.Output<string>;

    /**
     * Create a Connector resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ConnectorArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ConnectorArgs | ConnectorState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ConnectorState | undefined;
            inputs["ipCidrRange"] = state ? state.ipCidrRange : undefined;
            inputs["maxThroughput"] = state ? state.maxThroughput : undefined;
            inputs["minThroughput"] = state ? state.minThroughput : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["network"] = state ? state.network : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["region"] = state ? state.region : undefined;
            inputs["selfLink"] = state ? state.selfLink : undefined;
            inputs["state"] = state ? state.state : undefined;
        } else {
            const args = argsOrState as ConnectorArgs | undefined;
            if (!args || args.ipCidrRange === undefined) {
                throw new Error("Missing required property 'ipCidrRange'");
            }
            if (!args || args.network === undefined) {
                throw new Error("Missing required property 'network'");
            }
            if (!args || args.region === undefined) {
                throw new Error("Missing required property 'region'");
            }
            inputs["ipCidrRange"] = args ? args.ipCidrRange : undefined;
            inputs["maxThroughput"] = args ? args.maxThroughput : undefined;
            inputs["minThroughput"] = args ? args.minThroughput : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["network"] = args ? args.network : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["region"] = args ? args.region : undefined;
            inputs["selfLink"] = undefined /*out*/;
            inputs["state"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Connector.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Connector resources.
 */
export interface ConnectorState {
    /**
     * The range of internal addresses that follows RFC 4632 notation. Example: `10.132.0.0/28`.
     */
    readonly ipCidrRange?: pulumi.Input<string>;
    /**
     * Maximum throughput of the connector in Mbps, must be greater than `minThroughput`. Default is 1000.
     */
    readonly maxThroughput?: pulumi.Input<number>;
    /**
     * Minimum throughput of the connector in Mbps. Default and min is 200.
     */
    readonly minThroughput?: pulumi.Input<number>;
    /**
     * The name of the resource (Max 25 characters).
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Name of a VPC network.
     */
    readonly network?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * Region where the VPC Access connector resides
     */
    readonly region?: pulumi.Input<string>;
    /**
     * The fully qualified name of this VPC connector
     */
    readonly selfLink?: pulumi.Input<string>;
    /**
     * State of the VPC access connector.
     */
    readonly state?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Connector resource.
 */
export interface ConnectorArgs {
    /**
     * The range of internal addresses that follows RFC 4632 notation. Example: `10.132.0.0/28`.
     */
    readonly ipCidrRange: pulumi.Input<string>;
    /**
     * Maximum throughput of the connector in Mbps, must be greater than `minThroughput`. Default is 1000.
     */
    readonly maxThroughput?: pulumi.Input<number>;
    /**
     * Minimum throughput of the connector in Mbps. Default and min is 200.
     */
    readonly minThroughput?: pulumi.Input<number>;
    /**
     * The name of the resource (Max 25 characters).
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Name of a VPC network.
     */
    readonly network: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * Region where the VPC Access connector resides
     */
    readonly region: pulumi.Input<string>;
}
