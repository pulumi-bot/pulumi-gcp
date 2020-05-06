// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Get information about a Google Cloud Function. For more information see
 * the [official documentation](https://cloud.google.com/functions/docs/)
 * and [API](https://cloud.google.com/functions/docs/apis).
 * 
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/d/datasource_cloudfunctions_function.html.markdown.
 */
export function getFunction(args: GetFunctionArgs, opts?: pulumi.InvokeOptions): Promise<GetFunctionResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("gcp:cloudfunctions/getFunction:getFunction", {
        "name": args.name,
        "project": args.project,
        "region": args.region,
    }, opts);
}

/**
 * A collection of arguments for invoking getFunction.
 */
export interface GetFunctionArgs {
    /**
     * The name of a Cloud Function.
     */
    readonly name: string;
    /**
     * The project in which the resource belongs. If it
     * is not provided, the provider project is used.
     */
    readonly project?: string;
    /**
     * The region in which the resource belongs. If it
     * is not provided, the provider region is used.
     */
    readonly region?: string;
}

/**
 * A collection of values returned by getFunction.
 */
export interface GetFunctionResult {
    /**
     * Available memory (in MB) to the function.
     */
    readonly availableMemoryMb: number;
    /**
     * Description of the function.
     */
    readonly description: string;
    /**
     * Name of a JavaScript function that will be executed when the Google Cloud Function is triggered.
     */
    readonly entryPoint: string;
    readonly environmentVariables: {[key: string]: any};
    /**
     * A source that fires events in response to a condition in another service. Structure is documented below.
     */
    readonly eventTriggers: outputs.cloudfunctions.GetFunctionEventTrigger[];
    /**
     * If function is triggered by HTTP, trigger URL is set here.
     */
    readonly httpsTriggerUrl: string;
    /**
     * Controls what traffic can reach the function.
     */
    readonly ingressSettings: string;
    /**
     * A map of labels applied to this function.
     */
    readonly labels: {[key: string]: any};
    readonly maxInstances: number;
    /**
     * The name of the Cloud Function.
     */
    readonly name: string;
    readonly project?: string;
    readonly region?: string;
    /**
     * The runtime in which the function is running.
     */
    readonly runtime: string;
    /**
     * The service account email to be assumed by the cloud function.
     */
    readonly serviceAccountEmail: string;
    /**
     * The GCS bucket containing the zip archive which contains the function.
     */
    readonly sourceArchiveBucket: string;
    /**
     * The source archive object (file) in archive bucket.
     */
    readonly sourceArchiveObject: string;
    readonly sourceRepositories: outputs.cloudfunctions.GetFunctionSourceRepository[];
    /**
     * Function execution timeout (in seconds).
     */
    readonly timeout: number;
    /**
     * If function is triggered by HTTP, this boolean is set.
     */
    readonly triggerHttp: boolean;
    /**
     * The VPC Network Connector that this cloud function can connect to. 
     */
    readonly vpcConnector: string;
    /**
     * The egress settings for the connector, controlling what traffic is diverted through it.
     */
    readonly vpcConnectorEgressSettings: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
