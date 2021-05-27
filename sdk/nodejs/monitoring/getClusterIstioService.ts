// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * A Monitoring Service is the root resource under which operational aspects of a
 * generic service are accessible. A service is some discrete, autonomous, and
 * network-accessible unit, designed to solve an individual concern
 *
 * An Cluster Istio monitoring service is automatically created by GCP to monitor
 * Cluster Istio services.
 *
 * To get more information about Service, see:
 *
 * * [API documentation](https://cloud.google.com/monitoring/api/ref_v3/rest/v3/services)
 * * How-to Guides
 *     * [Service Monitoring](https://cloud.google.com/monitoring/service-monitoring)
 *     * [Monitoring API Documentation](https://cloud.google.com/monitoring/api/v3/)
 *
 * ## Example Usage
 * ### Monitoring Cluster Istio Service
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * // Monitors the default ClusterIstio service
 * const defaultClusterIstioService = pulumi.output(gcp.monitoring.getClusterIstioService({
 *     clusterName: "west",
 *     location: "us-west2-a",
 *     serviceName: "istio-policy",
 *     serviceNamespace: "istio-system",
 * }));
 * ```
 */
export function getClusterIstioService(args: GetClusterIstioServiceArgs, opts?: pulumi.InvokeOptions): Promise<GetClusterIstioServiceResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("gcp:monitoring/getClusterIstioService:getClusterIstioService", {
        "clusterName": args.clusterName,
        "location": args.location,
        "project": args.project,
        "serviceName": args.serviceName,
        "serviceNamespace": args.serviceNamespace,
    }, opts);
}

/**
 * A collection of arguments for invoking getClusterIstioService.
 */
export interface GetClusterIstioServiceArgs {
    /**
     * The name of the Kubernetes cluster in which this Istio service 
     * is defined. Corresponds to the clusterName resource label in k8sCluster resources.
     */
    clusterName: string;
    /**
     * The location of the Kubernetes cluster in which this Istio service 
     * is defined. Corresponds to the location resource label in k8sCluster resources.
     */
    location: string;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: string;
    /**
     * The name of the Istio service underlying this service.
     * Corresponds to the destinationServiceName metric label in Istio metrics.
     */
    serviceName: string;
    /**
     * The namespace of the Istio service underlying this service.
     * Corresponds to the destinationServiceNamespace metric label in Istio metrics.
     */
    serviceNamespace: string;
}

/**
 * A collection of values returned by getClusterIstioService.
 */
export interface GetClusterIstioServiceResult {
    readonly clusterName: string;
    readonly displayName: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly location: string;
    readonly name: string;
    readonly project?: string;
    readonly serviceId: string;
    readonly serviceName: string;
    readonly serviceNamespace: string;
    readonly telemetries: outputs.monitoring.GetClusterIstioServiceTelemetry[];
}

export function getClusterIstioServiceApply(args: GetClusterIstioServiceApplyArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetClusterIstioServiceResult> {
    return pulumi.output(args).apply(a => getClusterIstioService(a, opts))
}

/**
 * A collection of arguments for invoking getClusterIstioService.
 */
export interface GetClusterIstioServiceApplyArgs {
    /**
     * The name of the Kubernetes cluster in which this Istio service 
     * is defined. Corresponds to the clusterName resource label in k8sCluster resources.
     */
    clusterName: pulumi.Input<string>;
    /**
     * The location of the Kubernetes cluster in which this Istio service 
     * is defined. Corresponds to the location resource label in k8sCluster resources.
     */
    location: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * The name of the Istio service underlying this service.
     * Corresponds to the destinationServiceName metric label in Istio metrics.
     */
    serviceName: pulumi.Input<string>;
    /**
     * The namespace of the Istio service underlying this service.
     * Corresponds to the destinationServiceNamespace metric label in Istio metrics.
     */
    serviceNamespace: pulumi.Input<string>;
}
