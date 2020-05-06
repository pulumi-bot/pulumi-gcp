// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * A Monitoring Service is the root resource under which operational aspects of a
 * generic service are accessible. A service is some discrete, autonomous, and
 * network-accessible unit, designed to solve an individual concern
 * 
 * An App Engine monitoring service is automatically created by GCP to monitor
 * App Engine services.
 * 
 * 
 * To get more information about Service, see:
 * 
 * * [API documentation](https://cloud.google.com/monitoring/api/ref_v3/rest/v3/services)
 * * How-to Guides
 *     * [Service Monitoring](https://cloud.google.com/monitoring/service-monitoring)
 *     * [Monitoring API Documentation](https://cloud.google.com/monitoring/api/v3/)
 * 
 * ## Example Usage - Monitoring App Engine Service
 * 
 * 
 * {{ % example typescript % }}
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 * 
 * const bucket = new gcp.storage.Bucket("bucket", {});
 * const object = new gcp.storage.BucketObject("object", {
 *     bucket: bucket.name,
 *     source: new pulumi.asset.FileAsset("./test-fixtures/appengine/hello-world.zip"),
 * });
 * const myapp = new gcp.appengine.StandardAppVersion("myapp", {
 *     versionId: "v1",
 *     service: "myapp",
 *     runtime: "nodejs10",
 *     entrypoint: {
 *         shell: "node ./app.js",
 *     },
 *     deployment: {
 *         zip: {
 *             sourceUrl: pulumi.all([bucket.name, object.name]).apply(([bucketName, objectName]) => `https://storage.googleapis.com/${bucketName}/${objectName}`),
 *         },
 *     },
 *     envVariables: {
 *         port: "8080",
 *     },
 *     deleteServiceOnDestroy: false,
 * });
 * const srv = myapp.service.apply(service => gcp.monitoring.getAppEngineService({
 *     moduleId: service,
 * }));
 * ```
 * {{ % /example % }}
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/d/datasource_monitoring_app_engine_service.html.markdown.
 */
export function getAppEngineService(args: GetAppEngineServiceArgs, opts?: pulumi.InvokeOptions): Promise<GetAppEngineServiceResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("gcp:monitoring/getAppEngineService:getAppEngineService", {
        "moduleId": args.moduleId,
        "project": args.project,
    }, opts);
}

/**
 * A collection of arguments for invoking getAppEngineService.
 */
export interface GetAppEngineServiceArgs {
    /**
     * The ID of the App Engine module underlying this
     * service. Corresponds to the moduleId resource label in the [gaeApp](https://cloud.google.com/monitoring/api/resources#tag_gae_app) monitored resource, or the service/module name.
     */
    readonly moduleId: string;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: string;
}

/**
 * A collection of values returned by getAppEngineService.
 */
export interface GetAppEngineServiceResult {
    readonly displayName: string;
    readonly moduleId: string;
    readonly name: string;
    readonly project?: string;
    readonly serviceId: string;
    readonly telemetries: outputs.monitoring.GetAppEngineServiceTelemetry[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
