// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This resource creates and rolls out a Cloud Endpoints service using OpenAPI or gRPC.  View the relevant docs for [OpenAPI](https://cloud.google.com/endpoints/docs/openapi/) and [gRPC](https://cloud.google.com/endpoints/docs/grpc/).
 * 
 * {{% examples %}}
 * {{% /examples %}}
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/endpoints_service.html.markdown.
 */
export class Service extends pulumi.CustomResource {
    /**
     * Get an existing Service resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
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

    public /*out*/ readonly apis!: pulumi.Output<outputs.endpoints.ServiceApi[]>;
    public /*out*/ readonly configId!: pulumi.Output<string>;
    public /*out*/ readonly dnsAddress!: pulumi.Output<string>;
    public /*out*/ readonly endpoints!: pulumi.Output<outputs.endpoints.ServiceEndpoint[]>;
    public readonly grpcConfig!: pulumi.Output<string | undefined>;
    public readonly openapiConfig!: pulumi.Output<string | undefined>;
    public readonly project!: pulumi.Output<string>;
    public readonly protocOutputBase64!: pulumi.Output<string | undefined>;
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
    readonly apis?: pulumi.Input<pulumi.Input<inputs.endpoints.ServiceApi>[]>;
    readonly configId?: pulumi.Input<string>;
    readonly dnsAddress?: pulumi.Input<string>;
    readonly endpoints?: pulumi.Input<pulumi.Input<inputs.endpoints.ServiceEndpoint>[]>;
    readonly grpcConfig?: pulumi.Input<string>;
    readonly openapiConfig?: pulumi.Input<string>;
    readonly project?: pulumi.Input<string>;
    readonly protocOutputBase64?: pulumi.Input<string>;
    readonly serviceName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Service resource.
 */
export interface ServiceArgs {
    readonly grpcConfig?: pulumi.Input<string>;
    readonly openapiConfig?: pulumi.Input<string>;
    readonly project?: pulumi.Input<string>;
    readonly protocOutputBase64?: pulumi.Input<string>;
    readonly serviceName: pulumi.Input<string>;
}
