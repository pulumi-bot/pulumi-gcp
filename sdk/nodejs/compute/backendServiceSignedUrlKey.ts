// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * A key for signing Cloud CDN signed URLs for Backend Services.
 *
 *
 * To get more information about BackendServiceSignedUrlKey, see:
 *
 * * [API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/backendServices)
 * * How-to Guides
 *     * [Using Signed URLs](https://cloud.google.com/cdn/docs/using-signed-urls/)
 *
 * > **Warning:** All arguments including `keyValue` will be stored in the raw
 * state as plain-text.
 *
 * ## Example Usage - Backend Service Signed Url Key
 *
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const webserver = new gcp.compute.InstanceTemplate("webserver", {
 *     machineType: "n1-standard-1",
 *     network_interface: [{
 *         network: "default",
 *     }],
 *     disk: [{
 *         sourceImage: "debian-cloud/debian-9",
 *         autoDelete: true,
 *         boot: true,
 *     }],
 * });
 * const webservers = new gcp.compute.InstanceGroupManager("webservers", {
 *     version: [{
 *         instanceTemplate: webserver.id,
 *         name: "primary",
 *     }],
 *     baseInstanceName: "webserver",
 *     zone: "us-central1-f",
 *     targetSize: 1,
 * });
 * const default = new gcp.compute.HttpHealthCheck("default", {
 *     requestPath: "/",
 *     checkIntervalSec: 1,
 *     timeoutSec: 1,
 * });
 * const exampleBackend = new gcp.compute.BackendService("exampleBackend", {
 *     description: "Our company website",
 *     portName: "http",
 *     protocol: "HTTP",
 *     timeoutSec: 10,
 *     enableCdn: true,
 *     backend: [{
 *         group: webservers.instanceGroup,
 *     }],
 *     healthChecks: [default.id],
 * });
 * const backendKey = new gcp.compute.BackendServiceSignedUrlKey("backendKey", {
 *     keyValue: "pPsVemX8GM46QVeezid6Rw==",
 *     backendService: exampleBackend.name,
 * });
 * ```
 */
export class BackendServiceSignedUrlKey extends pulumi.CustomResource {
    /**
     * Get an existing BackendServiceSignedUrlKey resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: BackendServiceSignedUrlKeyState, opts?: pulumi.CustomResourceOptions): BackendServiceSignedUrlKey {
        return new BackendServiceSignedUrlKey(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:compute/backendServiceSignedUrlKey:BackendServiceSignedUrlKey';

    /**
     * Returns true if the given object is an instance of BackendServiceSignedUrlKey.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is BackendServiceSignedUrlKey {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === BackendServiceSignedUrlKey.__pulumiType;
    }

    /**
     * The backend service this signed URL key belongs.
     */
    public readonly backendService!: pulumi.Output<string>;
    /**
     * 128-bit key value used for signing the URL. The key value must be a
     * valid RFC 4648 Section 5 base64url encoded string.  **Note**: This property is sensitive and will not be displayed in the plan.
     */
    public readonly keyValue!: pulumi.Output<string>;
    /**
     * Name of the signed URL key.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;

    /**
     * Create a BackendServiceSignedUrlKey resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: BackendServiceSignedUrlKeyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: BackendServiceSignedUrlKeyArgs | BackendServiceSignedUrlKeyState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as BackendServiceSignedUrlKeyState | undefined;
            inputs["backendService"] = state ? state.backendService : undefined;
            inputs["keyValue"] = state ? state.keyValue : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["project"] = state ? state.project : undefined;
        } else {
            const args = argsOrState as BackendServiceSignedUrlKeyArgs | undefined;
            if (!args || args.backendService === undefined) {
                throw new Error("Missing required property 'backendService'");
            }
            if (!args || args.keyValue === undefined) {
                throw new Error("Missing required property 'keyValue'");
            }
            inputs["backendService"] = args ? args.backendService : undefined;
            inputs["keyValue"] = args ? args.keyValue : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["project"] = args ? args.project : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(BackendServiceSignedUrlKey.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering BackendServiceSignedUrlKey resources.
 */
export interface BackendServiceSignedUrlKeyState {
    /**
     * The backend service this signed URL key belongs.
     */
    readonly backendService?: pulumi.Input<string>;
    /**
     * 128-bit key value used for signing the URL. The key value must be a
     * valid RFC 4648 Section 5 base64url encoded string.  **Note**: This property is sensitive and will not be displayed in the plan.
     */
    readonly keyValue?: pulumi.Input<string>;
    /**
     * Name of the signed URL key.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a BackendServiceSignedUrlKey resource.
 */
export interface BackendServiceSignedUrlKeyArgs {
    /**
     * The backend service this signed URL key belongs.
     */
    readonly backendService: pulumi.Input<string>;
    /**
     * 128-bit key value used for signing the URL. The key value must be a
     * valid RFC 4648 Section 5 base64url encoded string.  **Note**: This property is sensitive and will not be displayed in the plan.
     */
    readonly keyValue: pulumi.Input<string>;
    /**
     * Name of the signed URL key.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
}
