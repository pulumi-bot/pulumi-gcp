// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The provider type for the google-beta package. By default, resources use package-wide configuration
 * settings, however an explicit `Provider` instance may be created and passed during resource
 * construction to achieve fine-grained programmatic control over provider settings. See the
 * [documentation](https://pulumi.io/reference/programming-model.html#providers) for more information.
 */
export class Provider extends pulumi.ProviderResource {

    /**
     * Create a Provider resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ProviderArgs, opts?: pulumi.ResourceOptions) {
        let inputs: pulumi.Inputs = {};
        {
            inputs["accessToken"] = args ? args.accessToken : undefined;
            inputs["credentials"] = (args ? args.credentials : undefined) || utilities.getEnv("GOOGLE_CREDENTIALS", "GOOGLE_CLOUD_KEYFILE_JSON", "GCLOUD_KEYFILE_JSON");
            inputs["project"] = (args ? args.project : undefined) || utilities.getEnv("GOOGLE_PROJECT", "GOOGLE_CLOUD_PROJECT", "GCLOUD_PROJECT", "CLOUDSDK_CORE_PROJECT");
            inputs["region"] = (args ? args.region : undefined) || utilities.getEnv("GOOGLE_REGION", "GCLOUD_REGION", "CLOUDSDK_COMPUTE_REGION");
            inputs["scopes"] = pulumi.output(args ? args.scopes : undefined).apply(JSON.stringify);
            inputs["zone"] = (args ? args.zone : undefined) || utilities.getEnv("GOOGLE_ZONE", "GCLOUD_ZONE", "CLOUDSDK_COMPUTE_ZONE");
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super("gcp", name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Provider resource.
 */
export interface ProviderArgs {
    readonly accessToken?: pulumi.Input<string>;
    readonly credentials?: pulumi.Input<string>;
    readonly project?: pulumi.Input<string>;
    readonly region?: pulumi.Input<string>;
    readonly scopes?: pulumi.Input<pulumi.Input<string>[]>;
    readonly zone?: pulumi.Input<string>;
}
