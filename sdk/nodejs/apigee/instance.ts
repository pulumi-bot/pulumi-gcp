// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * An `Instance` is the runtime dataplane in Apigee.
 *
 * To get more information about Instance, see:
 *
 * * [API documentation](https://cloud.google.com/apigee/docs/reference/apis/apigee/rest/v1/organizations.instances/create)
 * * How-to Guides
 *     * [Creating a runtime instance](https://cloud.google.com/apigee/docs/api-platform/get-started/create-instance)
 *
 * ## Example Usage
 *
 * ## Import
 *
 * Instance can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import gcp:apigee/instance:Instance default {{org_id}}/instances/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:apigee/instance:Instance default {{org_id}}/{{name}}
 * ```
 */
export class Instance extends pulumi.CustomResource {
    /**
     * Get an existing Instance resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: InstanceState, opts?: pulumi.CustomResourceOptions): Instance {
        return new Instance(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:apigee/instance:Instance';

    /**
     * Returns true if the given object is an instance of Instance.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Instance {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Instance.__pulumiType;
    }

    /**
     * Description of the instance.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Customer Managed Encryption Key (CMEK) used for disk and volume encryption. Required for Apigee paid subscriptions only.
     * Use the following format: `projects/([^/]+)/locations/([^/]+)/keyRings/([^/]+)/cryptoKeys/([^/]+)`
     */
    public readonly diskEncryptionKeyName!: pulumi.Output<string | undefined>;
    /**
     * Display name of the instance.
     */
    public readonly displayName!: pulumi.Output<string | undefined>;
    /**
     * Output only. Hostname or IP address of the exposed Apigee endpoint used by clients to connect to the service.
     */
    public /*out*/ readonly host!: pulumi.Output<string>;
    /**
     * Compute Engine location where the instance resides. For trial organization
     * subscriptions, the location must be a Compute Engine zone. For paid organization
     * subscriptions, it should correspond to a Compute Engine region.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * Resource ID of the instance.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The Apigee Organization associated with the Apigee instance,
     * in the format `organizations/{{org_name}}`.
     */
    public readonly orgId!: pulumi.Output<string>;
    /**
     * The size of the CIDR block range that will be reserved by the instance.
     * Default value is `SLASH_16`.
     * Possible values are `SLASH_16` and `SLASH_20`.
     */
    public readonly peeringCidrRange!: pulumi.Output<string | undefined>;
    /**
     * Output only. Port number of the exposed Apigee endpoint.
     */
    public /*out*/ readonly port!: pulumi.Output<string>;

    /**
     * Create a Instance resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: InstanceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: InstanceArgs | InstanceState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as InstanceState | undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["diskEncryptionKeyName"] = state ? state.diskEncryptionKeyName : undefined;
            inputs["displayName"] = state ? state.displayName : undefined;
            inputs["host"] = state ? state.host : undefined;
            inputs["location"] = state ? state.location : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["orgId"] = state ? state.orgId : undefined;
            inputs["peeringCidrRange"] = state ? state.peeringCidrRange : undefined;
            inputs["port"] = state ? state.port : undefined;
        } else {
            const args = argsOrState as InstanceArgs | undefined;
            if ((!args || args.location === undefined) && !opts.urn) {
                throw new Error("Missing required property 'location'");
            }
            if ((!args || args.orgId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'orgId'");
            }
            inputs["description"] = args ? args.description : undefined;
            inputs["diskEncryptionKeyName"] = args ? args.diskEncryptionKeyName : undefined;
            inputs["displayName"] = args ? args.displayName : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["orgId"] = args ? args.orgId : undefined;
            inputs["peeringCidrRange"] = args ? args.peeringCidrRange : undefined;
            inputs["host"] = undefined /*out*/;
            inputs["port"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Instance.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Instance resources.
 */
export interface InstanceState {
    /**
     * Description of the instance.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Customer Managed Encryption Key (CMEK) used for disk and volume encryption. Required for Apigee paid subscriptions only.
     * Use the following format: `projects/([^/]+)/locations/([^/]+)/keyRings/([^/]+)/cryptoKeys/([^/]+)`
     */
    readonly diskEncryptionKeyName?: pulumi.Input<string>;
    /**
     * Display name of the instance.
     */
    readonly displayName?: pulumi.Input<string>;
    /**
     * Output only. Hostname or IP address of the exposed Apigee endpoint used by clients to connect to the service.
     */
    readonly host?: pulumi.Input<string>;
    /**
     * Compute Engine location where the instance resides. For trial organization
     * subscriptions, the location must be a Compute Engine zone. For paid organization
     * subscriptions, it should correspond to a Compute Engine region.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * Resource ID of the instance.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The Apigee Organization associated with the Apigee instance,
     * in the format `organizations/{{org_name}}`.
     */
    readonly orgId?: pulumi.Input<string>;
    /**
     * The size of the CIDR block range that will be reserved by the instance.
     * Default value is `SLASH_16`.
     * Possible values are `SLASH_16` and `SLASH_20`.
     */
    readonly peeringCidrRange?: pulumi.Input<string>;
    /**
     * Output only. Port number of the exposed Apigee endpoint.
     */
    readonly port?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Instance resource.
 */
export interface InstanceArgs {
    /**
     * Description of the instance.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Customer Managed Encryption Key (CMEK) used for disk and volume encryption. Required for Apigee paid subscriptions only.
     * Use the following format: `projects/([^/]+)/locations/([^/]+)/keyRings/([^/]+)/cryptoKeys/([^/]+)`
     */
    readonly diskEncryptionKeyName?: pulumi.Input<string>;
    /**
     * Display name of the instance.
     */
    readonly displayName?: pulumi.Input<string>;
    /**
     * Compute Engine location where the instance resides. For trial organization
     * subscriptions, the location must be a Compute Engine zone. For paid organization
     * subscriptions, it should correspond to a Compute Engine region.
     */
    readonly location: pulumi.Input<string>;
    /**
     * Resource ID of the instance.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The Apigee Organization associated with the Apigee instance,
     * in the format `organizations/{{org_name}}`.
     */
    readonly orgId: pulumi.Input<string>;
    /**
     * The size of the CIDR block range that will be reserved by the instance.
     * Default value is `SLASH_16`.
     * Possible values are `SLASH_16` and `SLASH_20`.
     */
    readonly peeringCidrRange?: pulumi.Input<string>;
}