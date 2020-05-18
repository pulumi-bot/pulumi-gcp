// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * A Google Cloud Firebase web application configuration
 *
 * To get more information about WebApp, see:
 *
 * * [API documentation](https://firebase.google.com/docs/projects/api/reference/rest/v1beta1/projects.webApps)
 * * How-to Guides
 *     * [Official Documentation](https://firebase.google.com/)
 */
export function getWebAppConfig(args: GetWebAppConfigArgs, opts?: pulumi.InvokeOptions): Promise<GetWebAppConfigResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("gcp:firebase/getWebAppConfig:getWebAppConfig", {
        "project": args.project,
        "webAppId": args.webAppId,
    }, opts);
}

/**
 * A collection of arguments for invoking getWebAppConfig.
 */
export interface GetWebAppConfigArgs {
    /**
     * The ID of the project in which the resource belongs. If it
     * is not provided, the provider project is used.
     */
    readonly project?: string;
    /**
     * the id of the firebase web app
     */
    readonly webAppId: string;
}

/**
 * A collection of values returned by getWebAppConfig.
 */
export interface GetWebAppConfigResult {
    readonly apiKey: string;
    readonly authDomain: string;
    readonly databaseUrl: string;
    readonly locationId: string;
    readonly measurementId: string;
    readonly messagingSenderId: string;
    readonly project?: string;
    readonly storageBucket: string;
    readonly webAppId: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
