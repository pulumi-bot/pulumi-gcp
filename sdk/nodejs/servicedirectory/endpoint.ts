// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * An individual endpoint that provides a service.
 *
 * To get more information about Endpoint, see:
 *
 * * [API documentation](https://cloud.google.com/service-directory/docs/reference/rest/v1beta1/projects.locations.namespaces.services.endpoints)
 * * How-to Guides
 *     * [Configuring an endpoint](https://cloud.google.com/service-directory/docs/configuring-service-directory#configuring_an_endpoint)
 *
 * ## Example Usage - Service Directory Endpoint Basic
 *
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const exampleNamespace = new gcp.servicedirectory.Namespace("exampleNamespace", {
 *     namespaceId: "example-namespace",
 *     location: "us-central1",
 * });
 * const exampleService = new gcp.servicedirectory.Service("exampleService", {
 *     serviceId: "example-service",
 *     namespace: exampleNamespace.id,
 * });
 * const exampleEndpoint = new gcp.servicedirectory.Endpoint("exampleEndpoint", {
 *     endpointId: "example-endpoint",
 *     service: exampleService.id,
 *     metadata: {
 *         stage: "prod",
 *         region: "us-central1",
 *     },
 *     address: "1.2.3.4",
 *     port: 5353,
 * });
 * ```
 */
export class Endpoint extends pulumi.CustomResource {
    /**
     * Get an existing Endpoint resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: EndpointState, opts?: pulumi.CustomResourceOptions): Endpoint {
        return new Endpoint(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:servicedirectory/endpoint:Endpoint';

    /**
     * Returns true if the given object is an instance of Endpoint.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Endpoint {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Endpoint.__pulumiType;
    }

    /**
     * IPv4 or IPv6 address of the endpoint.
     */
    public readonly address!: pulumi.Output<string | undefined>;
    /**
     * The Resource ID must be 1-63 characters long, including digits,
     * lowercase letters or the hyphen character.
     */
    public readonly endpointId!: pulumi.Output<string>;
    /**
     * Metadata for the endpoint. This data can be consumed
     * by service clients. The entire metadata dictionary may contain
     * up to 512 characters, spread across all key-value pairs.
     * Metadata that goes beyond any these limits will be rejected.
     */
    public readonly metadata!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The resource name for the endpoint in the format 'projects/*&#47;locations/*&#47;namespaces/*&#47;services/*&#47;endpoints/*'.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Port that the endpoint is running on, must be in the
     * range of [0, 65535]. If unspecified, the default is 0.
     */
    public readonly port!: pulumi.Output<number | undefined>;
    /**
     * The resource name of the service that this endpoint provides.
     */
    public readonly service!: pulumi.Output<string>;

    /**
     * Create a Endpoint resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: EndpointArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: EndpointArgs | EndpointState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as EndpointState | undefined;
            inputs["address"] = state ? state.address : undefined;
            inputs["endpointId"] = state ? state.endpointId : undefined;
            inputs["metadata"] = state ? state.metadata : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["port"] = state ? state.port : undefined;
            inputs["service"] = state ? state.service : undefined;
        } else {
            const args = argsOrState as EndpointArgs | undefined;
            if (!args || args.endpointId === undefined) {
                throw new Error("Missing required property 'endpointId'");
            }
            if (!args || args.service === undefined) {
                throw new Error("Missing required property 'service'");
            }
            inputs["address"] = args ? args.address : undefined;
            inputs["endpointId"] = args ? args.endpointId : undefined;
            inputs["metadata"] = args ? args.metadata : undefined;
            inputs["port"] = args ? args.port : undefined;
            inputs["service"] = args ? args.service : undefined;
            inputs["name"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Endpoint.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Endpoint resources.
 */
export interface EndpointState {
    /**
     * IPv4 or IPv6 address of the endpoint.
     */
    readonly address?: pulumi.Input<string>;
    /**
     * The Resource ID must be 1-63 characters long, including digits,
     * lowercase letters or the hyphen character.
     */
    readonly endpointId?: pulumi.Input<string>;
    /**
     * Metadata for the endpoint. This data can be consumed
     * by service clients. The entire metadata dictionary may contain
     * up to 512 characters, spread across all key-value pairs.
     * Metadata that goes beyond any these limits will be rejected.
     */
    readonly metadata?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The resource name for the endpoint in the format 'projects/*&#47;locations/*&#47;namespaces/*&#47;services/*&#47;endpoints/*'.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Port that the endpoint is running on, must be in the
     * range of [0, 65535]. If unspecified, the default is 0.
     */
    readonly port?: pulumi.Input<number>;
    /**
     * The resource name of the service that this endpoint provides.
     */
    readonly service?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Endpoint resource.
 */
export interface EndpointArgs {
    /**
     * IPv4 or IPv6 address of the endpoint.
     */
    readonly address?: pulumi.Input<string>;
    /**
     * The Resource ID must be 1-63 characters long, including digits,
     * lowercase letters or the hyphen character.
     */
    readonly endpointId: pulumi.Input<string>;
    /**
     * Metadata for the endpoint. This data can be consumed
     * by service clients. The entire metadata dictionary may contain
     * up to 512 characters, spread across all key-value pairs.
     * Metadata that goes beyond any these limits will be rejected.
     */
    readonly metadata?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Port that the endpoint is running on, must be in the
     * range of [0, 65535]. If unspecified, the default is 0.
     */
    readonly port?: pulumi.Input<number>;
    /**
     * The resource name of the service that this endpoint provides.
     */
    readonly service: pulumi.Input<string>;
}
