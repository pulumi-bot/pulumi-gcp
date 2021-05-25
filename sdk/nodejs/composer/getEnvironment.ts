// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Provides access to Cloud Composer environment configuration in a region for a given project.
 */
export function getEnvironment(args: GetEnvironmentArgs, opts?: pulumi.InvokeOptions): Promise<GetEnvironmentResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("gcp:composer/getEnvironment:getEnvironment", {
        "name": args.name,
        "project": args.project,
        "region": args.region,
    }, opts);
}

/**
 * A collection of arguments for invoking getEnvironment.
 */
export interface GetEnvironmentArgs {
    /**
     * Name of the environment.
     */
    name: string;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: string;
    /**
     * The location or Compute Engine region of the environment.
     */
    region?: string;
}

/**
 * A collection of values returned by getEnvironment.
 */
export interface GetEnvironmentResult {
    /**
     * Configuration parameters for the environment.
     */
    readonly configs: outputs.composer.GetEnvironmentConfig[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly labels: {[key: string]: string};
    readonly name: string;
    readonly project?: string;
    readonly region?: string;
}
