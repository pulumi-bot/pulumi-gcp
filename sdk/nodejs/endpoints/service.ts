// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * This resource creates and rolls out a Cloud Endpoints service using OpenAPI or gRPC.  View the relevant docs for [OpenAPI](https://cloud.google.com/endpoints/docs/openapi/) and [gRPC](https://cloud.google.com/endpoints/docs/grpc/).
 */
export class Service extends pulumi.CustomResource {
    /**
     * Get an existing Service resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ServiceState, opts?: pulumi.CustomResourceOptions): Service {
        return new Service(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:endpoints/service:Service';

    /**
     * Returns true if the given object is an instance of Service.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Service {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Service.__pulumiType;
    }

    /**
     * A list of API objects.
     */
    public /*out*/ readonly apis!: pulumi.Output<outputs.endpoints.ServiceApi[]>;
    /**
     * The autogenerated ID for the configuration that is rolled out as part of the creation of this resource. Must be provided
     * to compute engine instances as a tag.
     */
    public /*out*/ readonly configId!: pulumi.Output<string>;
    /**
     * The address at which the service can be found - usually the same as the service name.
     */
    public /*out*/ readonly dnsAddress!: pulumi.Output<string>;
    /**
     * A list of Endpoint objects.
     */
    public /*out*/ readonly endpoints!: pulumi.Output<outputs.endpoints.ServiceEndpoint[]>;
    /**
     * The full text of the Service Config YAML file (Example located here). If provided, must also provide
     * protoc_output_base64. open_api config must not be provided.
     */
    public readonly grpcConfig!: pulumi.Output<string | undefined>;
    /**
     * The full text of the OpenAPI YAML configuration as described here. Either this, or both of grpc_config and
     * protoc_output_base64 must be specified.
     */
    public readonly openapiConfig!: pulumi.Output<string | undefined>;
    /**
     * The project ID that the service belongs to. If not provided, provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * The full contents of the Service Descriptor File generated by protoc. This should be a compiled .pb file,
     * base64-encoded.
     */
    public readonly protocOutputBase64!: pulumi.Output<string | undefined>;
    /**
     * The name of the service. Usually of the form $apiname.endpoints.$projectid.cloud.goog.
     */
    public readonly serviceName!: pulumi.Output<string>;

    /**
     * Create a Service resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ServiceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ServiceArgs | ServiceState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ServiceState | undefined;
            inputs["apis"] = state ? state.apis : undefined;
            inputs["configId"] = state ? state.configId : undefined;
            inputs["dnsAddress"] = state ? state.dnsAddress : undefined;
            inputs["endpoints"] = state ? state.endpoints : undefined;
            inputs["grpcConfig"] = state ? state.grpcConfig : undefined;
            inputs["openapiConfig"] = state ? state.openapiConfig : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["protocOutputBase64"] = state ? state.protocOutputBase64 : undefined;
            inputs["serviceName"] = state ? state.serviceName : undefined;
        } else {
            const args = argsOrState as ServiceArgs | undefined;
            if (!args || args.serviceName === undefined) {
                throw new Error("Missing required property 'serviceName'");
            }
            inputs["grpcConfig"] = args ? args.grpcConfig : undefined;
            inputs["openapiConfig"] = args ? args.openapiConfig : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["protocOutputBase64"] = args ? args.protocOutputBase64 : undefined;
            inputs["serviceName"] = args ? args.serviceName : undefined;
            inputs["apis"] = undefined /*out*/;
            inputs["configId"] = undefined /*out*/;
            inputs["dnsAddress"] = undefined /*out*/;
            inputs["endpoints"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Service.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Service resources.
 */
export interface ServiceState {
    /**
     * A list of API objects.
     */
    readonly apis?: pulumi.Input<pulumi.Input<inputs.endpoints.ServiceApi>[]>;
    /**
     * The autogenerated ID for the configuration that is rolled out as part of the creation of this resource. Must be provided
     * to compute engine instances as a tag.
     */
    readonly configId?: pulumi.Input<string>;
    /**
     * The address at which the service can be found - usually the same as the service name.
     */
    readonly dnsAddress?: pulumi.Input<string>;
    /**
     * A list of Endpoint objects.
     */
    readonly endpoints?: pulumi.Input<pulumi.Input<inputs.endpoints.ServiceEndpoint>[]>;
    /**
     * The full text of the Service Config YAML file (Example located here). If provided, must also provide
     * protoc_output_base64. open_api config must not be provided.
     */
    readonly grpcConfig?: pulumi.Input<string>;
    /**
     * The full text of the OpenAPI YAML configuration as described here. Either this, or both of grpc_config and
     * protoc_output_base64 must be specified.
     */
    readonly openapiConfig?: pulumi.Input<string>;
    /**
     * The project ID that the service belongs to. If not provided, provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * The full contents of the Service Descriptor File generated by protoc. This should be a compiled .pb file,
     * base64-encoded.
     */
    readonly protocOutputBase64?: pulumi.Input<string>;
    /**
     * The name of the service. Usually of the form $apiname.endpoints.$projectid.cloud.goog.
     */
    readonly serviceName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Service resource.
 */
export interface ServiceArgs {
    /**
     * The full text of the Service Config YAML file (Example located here). If provided, must also provide
     * protoc_output_base64. open_api config must not be provided.
     */
    readonly grpcConfig?: pulumi.Input<string>;
    /**
     * The full text of the OpenAPI YAML configuration as described here. Either this, or both of grpc_config and
     * protoc_output_base64 must be specified.
     */
    readonly openapiConfig?: pulumi.Input<string>;
    /**
     * The project ID that the service belongs to. If not provided, provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * The full contents of the Service Descriptor File generated by protoc. This should be a compiled .pb file,
     * base64-encoded.
     */
    readonly protocOutputBase64?: pulumi.Input<string>;
    /**
     * The name of the service. Usually of the form $apiname.endpoints.$projectid.cloud.goog.
     */
    readonly serviceName: pulumi.Input<string>;
}
