// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Creates a job on Dataflow, which is an implementation of Apache Beam running on Google Compute Engine. For more information see
 * the official documentation for
 * [Beam](https://beam.apache.org) and [Dataflow](https://cloud.google.com/dataflow/).
 * 
 * 
 * ## Note on "destroy" / "apply"
 * 
 * There are many types of Dataflow jobs.  Some Dataflow jobs run constantly, getting new data from (e.g.) a GCS bucket, and outputting data continuously.  Some jobs process a set amount of data then terminate.  All jobs can fail while running due to programming errors or other issues.  In this way, Dataflow jobs are different from most other Google resources.
 * 
 * The Dataflow resource is considered 'existing' while it is in a nonterminal state.  If it reaches a terminal state (e.g. 'FAILED', 'COMPLETE', 'CANCELLED'), it will be recreated on the next 'apply'.  This is as expected for jobs which run continuously, but may surprise users who use this resource for other kinds of Dataflow jobs.
 * 
 * A Dataflow job which is 'destroyed' may be "cancelled" or "drained".  If "cancelled", the job terminates - any data written remains where it is, but no new data will be processed.  If "drained", no new data will enter the pipeline, but any data currently in the pipeline will finish being processed.  The default is "cancelled", but if a user sets `onDelete` to `"drain"` in the configuration, you may experience a long wait for your `pulumi destroy` to complete.
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/dataflow_job.html.markdown.
 */
export class Job extends pulumi.CustomResource {
    /**
     * Get an existing Job resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: JobState, opts?: pulumi.CustomResourceOptions): Job {
        return new Job(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:dataflow/job:Job';

    /**
     * Returns true if the given object is an instance of Job.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Job {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Job.__pulumiType;
    }

    /**
     * List of experiments that should be used by the job. An example value is `["enableStackdriverAgentMetrics"]`.
     */
    public readonly additionalExperiments!: pulumi.Output<string[] | undefined>;
    /**
     * The configuration for VM IPs.  Options are `"WORKER_IP_PUBLIC"` or `"WORKER_IP_PRIVATE"`.
     */
    public readonly ipConfiguration!: pulumi.Output<string | undefined>;
    /**
     * The unique ID of this job.
     */
    public /*out*/ readonly jobId!: pulumi.Output<string>;
    /**
     * User labels to be specified for the job. Keys and values should follow the restrictions
     * specified in the [labeling restrictions](https://cloud.google.com/compute/docs/labeling-resources#restrictions) page.
     * **NOTE**: Google-provided Dataflow templates often provide default labels that begin with `goog-dataflow-provided`.
     * Unless explicitly set in config, these labels will be ignored to prevent diffs on re-apply.
     */
    public readonly labels!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * The machine type to use for the job.
     */
    public readonly machineType!: pulumi.Output<string | undefined>;
    /**
     * The number of workers permitted to work on the job.  More workers may improve processing speed at additional cost.
     */
    public readonly maxWorkers!: pulumi.Output<number | undefined>;
    /**
     * A unique name for the resource, required by Dataflow.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The network to which VMs will be assigned. If it is not provided, "default" will be used.
     */
    public readonly network!: pulumi.Output<string | undefined>;
    /**
     * One of "drain" or "cancel".  Specifies behavior of deletion during `pulumi destroy`.  See above note.
     */
    public readonly onDelete!: pulumi.Output<string | undefined>;
    /**
     * Key/Value pairs to be passed to the Dataflow job (as used in the template).
     */
    public readonly parameters!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * The project in which the resource belongs. If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    public readonly region!: pulumi.Output<string | undefined>;
    /**
     * The Service Account email used to create the job.
     */
    public readonly serviceAccountEmail!: pulumi.Output<string | undefined>;
    /**
     * The current state of the resource, selected from the [JobState enum](https://cloud.google.com/dataflow/docs/reference/rest/v1b3/projects.jobs#Job.JobState)
     */
    public /*out*/ readonly state!: pulumi.Output<string>;
    /**
     * The subnetwork to which VMs will be assigned. Should be of the form "regions/REGION/subnetworks/SUBNETWORK".
     */
    public readonly subnetwork!: pulumi.Output<string | undefined>;
    /**
     * A writeable location on GCS for the Dataflow job to dump its temporary data.
     */
    public readonly tempGcsLocation!: pulumi.Output<string>;
    /**
     * The GCS path to the Dataflow job template.
     */
    public readonly templateGcsPath!: pulumi.Output<string>;
    /**
     * The type of this job, selected from the [JobType enum](https://cloud.google.com/dataflow/docs/reference/rest/v1b3/projects.jobs#Job.JobType)
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * The zone in which the created job should run. If it is not provided, the provider zone is used.
     */
    public readonly zone!: pulumi.Output<string | undefined>;

    /**
     * Create a Job resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: JobArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: JobArgs | JobState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as JobState | undefined;
            inputs["additionalExperiments"] = state ? state.additionalExperiments : undefined;
            inputs["ipConfiguration"] = state ? state.ipConfiguration : undefined;
            inputs["jobId"] = state ? state.jobId : undefined;
            inputs["labels"] = state ? state.labels : undefined;
            inputs["machineType"] = state ? state.machineType : undefined;
            inputs["maxWorkers"] = state ? state.maxWorkers : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["network"] = state ? state.network : undefined;
            inputs["onDelete"] = state ? state.onDelete : undefined;
            inputs["parameters"] = state ? state.parameters : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["region"] = state ? state.region : undefined;
            inputs["serviceAccountEmail"] = state ? state.serviceAccountEmail : undefined;
            inputs["state"] = state ? state.state : undefined;
            inputs["subnetwork"] = state ? state.subnetwork : undefined;
            inputs["tempGcsLocation"] = state ? state.tempGcsLocation : undefined;
            inputs["templateGcsPath"] = state ? state.templateGcsPath : undefined;
            inputs["type"] = state ? state.type : undefined;
            inputs["zone"] = state ? state.zone : undefined;
        } else {
            const args = argsOrState as JobArgs | undefined;
            if (!args || args.tempGcsLocation === undefined) {
                throw new Error("Missing required property 'tempGcsLocation'");
            }
            if (!args || args.templateGcsPath === undefined) {
                throw new Error("Missing required property 'templateGcsPath'");
            }
            inputs["additionalExperiments"] = args ? args.additionalExperiments : undefined;
            inputs["ipConfiguration"] = args ? args.ipConfiguration : undefined;
            inputs["labels"] = args ? args.labels : undefined;
            inputs["machineType"] = args ? args.machineType : undefined;
            inputs["maxWorkers"] = args ? args.maxWorkers : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["network"] = args ? args.network : undefined;
            inputs["onDelete"] = args ? args.onDelete : undefined;
            inputs["parameters"] = args ? args.parameters : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["region"] = args ? args.region : undefined;
            inputs["serviceAccountEmail"] = args ? args.serviceAccountEmail : undefined;
            inputs["subnetwork"] = args ? args.subnetwork : undefined;
            inputs["tempGcsLocation"] = args ? args.tempGcsLocation : undefined;
            inputs["templateGcsPath"] = args ? args.templateGcsPath : undefined;
            inputs["zone"] = args ? args.zone : undefined;
            inputs["jobId"] = undefined /*out*/;
            inputs["state"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Job.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Job resources.
 */
export interface JobState {
    /**
     * List of experiments that should be used by the job. An example value is `["enableStackdriverAgentMetrics"]`.
     */
    readonly additionalExperiments?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The configuration for VM IPs.  Options are `"WORKER_IP_PUBLIC"` or `"WORKER_IP_PRIVATE"`.
     */
    readonly ipConfiguration?: pulumi.Input<string>;
    /**
     * The unique ID of this job.
     */
    readonly jobId?: pulumi.Input<string>;
    /**
     * User labels to be specified for the job. Keys and values should follow the restrictions
     * specified in the [labeling restrictions](https://cloud.google.com/compute/docs/labeling-resources#restrictions) page.
     * **NOTE**: Google-provided Dataflow templates often provide default labels that begin with `goog-dataflow-provided`.
     * Unless explicitly set in config, these labels will be ignored to prevent diffs on re-apply.
     */
    readonly labels?: pulumi.Input<{[key: string]: any}>;
    /**
     * The machine type to use for the job.
     */
    readonly machineType?: pulumi.Input<string>;
    /**
     * The number of workers permitted to work on the job.  More workers may improve processing speed at additional cost.
     */
    readonly maxWorkers?: pulumi.Input<number>;
    /**
     * A unique name for the resource, required by Dataflow.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The network to which VMs will be assigned. If it is not provided, "default" will be used.
     */
    readonly network?: pulumi.Input<string>;
    /**
     * One of "drain" or "cancel".  Specifies behavior of deletion during `pulumi destroy`.  See above note.
     */
    readonly onDelete?: pulumi.Input<string>;
    /**
     * Key/Value pairs to be passed to the Dataflow job (as used in the template).
     */
    readonly parameters?: pulumi.Input<{[key: string]: any}>;
    /**
     * The project in which the resource belongs. If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    readonly region?: pulumi.Input<string>;
    /**
     * The Service Account email used to create the job.
     */
    readonly serviceAccountEmail?: pulumi.Input<string>;
    /**
     * The current state of the resource, selected from the [JobState enum](https://cloud.google.com/dataflow/docs/reference/rest/v1b3/projects.jobs#Job.JobState)
     */
    readonly state?: pulumi.Input<string>;
    /**
     * The subnetwork to which VMs will be assigned. Should be of the form "regions/REGION/subnetworks/SUBNETWORK".
     */
    readonly subnetwork?: pulumi.Input<string>;
    /**
     * A writeable location on GCS for the Dataflow job to dump its temporary data.
     */
    readonly tempGcsLocation?: pulumi.Input<string>;
    /**
     * The GCS path to the Dataflow job template.
     */
    readonly templateGcsPath?: pulumi.Input<string>;
    /**
     * The type of this job, selected from the [JobType enum](https://cloud.google.com/dataflow/docs/reference/rest/v1b3/projects.jobs#Job.JobType)
     */
    readonly type?: pulumi.Input<string>;
    /**
     * The zone in which the created job should run. If it is not provided, the provider zone is used.
     */
    readonly zone?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Job resource.
 */
export interface JobArgs {
    /**
     * List of experiments that should be used by the job. An example value is `["enableStackdriverAgentMetrics"]`.
     */
    readonly additionalExperiments?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The configuration for VM IPs.  Options are `"WORKER_IP_PUBLIC"` or `"WORKER_IP_PRIVATE"`.
     */
    readonly ipConfiguration?: pulumi.Input<string>;
    /**
     * User labels to be specified for the job. Keys and values should follow the restrictions
     * specified in the [labeling restrictions](https://cloud.google.com/compute/docs/labeling-resources#restrictions) page.
     * **NOTE**: Google-provided Dataflow templates often provide default labels that begin with `goog-dataflow-provided`.
     * Unless explicitly set in config, these labels will be ignored to prevent diffs on re-apply.
     */
    readonly labels?: pulumi.Input<{[key: string]: any}>;
    /**
     * The machine type to use for the job.
     */
    readonly machineType?: pulumi.Input<string>;
    /**
     * The number of workers permitted to work on the job.  More workers may improve processing speed at additional cost.
     */
    readonly maxWorkers?: pulumi.Input<number>;
    /**
     * A unique name for the resource, required by Dataflow.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The network to which VMs will be assigned. If it is not provided, "default" will be used.
     */
    readonly network?: pulumi.Input<string>;
    /**
     * One of "drain" or "cancel".  Specifies behavior of deletion during `pulumi destroy`.  See above note.
     */
    readonly onDelete?: pulumi.Input<string>;
    /**
     * Key/Value pairs to be passed to the Dataflow job (as used in the template).
     */
    readonly parameters?: pulumi.Input<{[key: string]: any}>;
    /**
     * The project in which the resource belongs. If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    readonly region?: pulumi.Input<string>;
    /**
     * The Service Account email used to create the job.
     */
    readonly serviceAccountEmail?: pulumi.Input<string>;
    /**
     * The subnetwork to which VMs will be assigned. Should be of the form "regions/REGION/subnetworks/SUBNETWORK".
     */
    readonly subnetwork?: pulumi.Input<string>;
    /**
     * A writeable location on GCS for the Dataflow job to dump its temporary data.
     */
    readonly tempGcsLocation: pulumi.Input<string>;
    /**
     * The GCS path to the Dataflow job template.
     */
    readonly templateGcsPath: pulumi.Input<string>;
    /**
     * The zone in which the created job should run. If it is not provided, the provider zone is used.
     */
    readonly zone?: pulumi.Input<string>;
}
