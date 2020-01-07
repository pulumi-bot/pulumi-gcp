// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export function getRouter(args: GetRouterArgs, opts?: pulumi.InvokeOptions): Promise<GetRouterResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("gcp:compute/getRouter:getRouter", {
        "name": args.name,
        "network": args.network,
        "project": args.project,
        "region": args.region,
    }, opts);
}

/**
 * A collection of arguments for invoking getRouter.
 */
export interface GetRouterArgs {
    readonly name: string;
    readonly network: string;
    readonly project?: string;
    readonly region?: string;
}

/**
 * A collection of values returned by getRouter.
 */
export interface GetRouterResult {
    readonly bgps: outputs.compute.GetRouterBgp[];
    readonly creationTimestamp: string;
    readonly description: string;
    readonly name: string;
    readonly network: string;
    readonly project?: string;
    readonly region?: string;
    readonly selfLink: string;
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
