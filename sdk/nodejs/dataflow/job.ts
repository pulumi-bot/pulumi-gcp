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
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 * 
 * const google_dataflow_job_big_data_job = new gcp.dataflow.Job("big_data_job", {
 *     name: "dataflow-job",
 *     parameters: {
 *         baz: "qux",
 *         foo: "bar",
 *     },
 *     tempGcsLocation: "gs://my-bucket/tmp_dir",
 *     templateGcsPath: "gs://my-bucket/templates/template_file",
 * });
 * ```
 * 
 * ## Note on "destroy" / "apply"
 * 
 * There are many types of Dataflow jobs.  Some Dataflow jobs run constantly, getting new data from (e.g.) a GCS bucket, and outputting data continuously.  Some jobs process a set amount of data then terminate.  All jobs can fail while running due to programming errors or other issues.  In this way, Dataflow jobs are different from most other Terraform / Google resources.
 * 
 * The Dataflow resource is considered 'existing' while it is in a nonterminal state.  If it reaches a terminal state (e.g. 'FAILED', 'COMPLETE', 'CANCELLED'), it will be recreated on the next 'apply'.  This is as expected for jobs which run continously, but may surprise users who use this resource for other kinds of Dataflow jobs.
 * 
 * A Dataflow job which is 'destroyed' may be "cancelled" or "drained".  If "cancelled", the job terminates - any data written remains where it is, but no new data will be processed.  If "drained", no new data will enter the pipeline, but any data currently in the pipeline will finish being processed.  The default is "cancelled", but if a user sets `on_delete` to `"drain"` in the configuration, you may experience a long wait for your `terraform destroy` to complete.
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

    /**
     * The number of workers permitted to work on the job.  More workers may improve processing speed at additional cost.
     */
    public readonly maxWorkers: pulumi.Output<number | undefined>;
    /**
     * A unique name for the resource, required by Dataflow.
     */
    public readonly name: pulumi.Output<string>;
    /**
     * One of "drain" or "cancel".  Specifies behavior of deletion during `terraform destroy`.  See above note.
     */
    public readonly onDelete: pulumi.Output<string | undefined>;
    /**
     * Key/Value pairs to be passed to the Dataflow job (as used in the template).
     */
    public readonly parameters: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * The project in which the resource belongs. If it is not provided, the provider project is used.
     */
    public readonly project: pulumi.Output<string | undefined>;
    public readonly region: pulumi.Output<string | undefined>;
    /**
     * The current state of the resource, selected from the [JobState enum](https://cloud.google.com/dataflow/docs/reference/rest/v1b3/projects.jobs#Job.JobState)
     */
    public /*out*/ readonly state: pulumi.Output<string>;
    /**
     * A writeable location on GCS for the Dataflow job to dump its temporary data.
     */
    public readonly tempGcsLocation: pulumi.Output<string>;
    /**
     * The GCS path to the Dataflow job template.
     */
    public readonly templateGcsPath: pulumi.Output<string>;
    /**
     * The zone in which the created job should run. If it is not provided, the provider zone is used.
     */
    public readonly zone: pulumi.Output<string | undefined>;

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
            const state: JobState = argsOrState as JobState | undefined;
            inputs["maxWorkers"] = state ? state.maxWorkers : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["onDelete"] = state ? state.onDelete : undefined;
            inputs["parameters"] = state ? state.parameters : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["region"] = state ? state.region : undefined;
            inputs["state"] = state ? state.state : undefined;
            inputs["tempGcsLocation"] = state ? state.tempGcsLocation : undefined;
            inputs["templateGcsPath"] = state ? state.templateGcsPath : undefined;
            inputs["zone"] = state ? state.zone : undefined;
        } else {
            const args = argsOrState as JobArgs | undefined;
            if (!args || args.tempGcsLocation === undefined) {
                throw new Error("Missing required property 'tempGcsLocation'");
            }
            if (!args || args.templateGcsPath === undefined) {
                throw new Error("Missing required property 'templateGcsPath'");
            }
            inputs["maxWorkers"] = args ? args.maxWorkers : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["onDelete"] = args ? args.onDelete : undefined;
            inputs["parameters"] = args ? args.parameters : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["region"] = args ? args.region : undefined;
            inputs["tempGcsLocation"] = args ? args.tempGcsLocation : undefined;
            inputs["templateGcsPath"] = args ? args.templateGcsPath : undefined;
            inputs["zone"] = args ? args.zone : undefined;
            inputs["state"] = undefined /*out*/;
        }
        super("gcp:dataflow/job:Job", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Job resources.
 */
export interface JobState {
    /**
     * The number of workers permitted to work on the job.  More workers may improve processing speed at additional cost.
     */
    readonly maxWorkers?: pulumi.Input<number>;
    /**
     * A unique name for the resource, required by Dataflow.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * One of "drain" or "cancel".  Specifies behavior of deletion during `terraform destroy`.  See above note.
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
     * The current state of the resource, selected from the [JobState enum](https://cloud.google.com/dataflow/docs/reference/rest/v1b3/projects.jobs#Job.JobState)
     */
    readonly state?: pulumi.Input<string>;
    /**
     * A writeable location on GCS for the Dataflow job to dump its temporary data.
     */
    readonly tempGcsLocation?: pulumi.Input<string>;
    /**
     * The GCS path to the Dataflow job template.
     */
    readonly templateGcsPath?: pulumi.Input<string>;
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
     * The number of workers permitted to work on the job.  More workers may improve processing speed at additional cost.
     */
    readonly maxWorkers?: pulumi.Input<number>;
    /**
     * A unique name for the resource, required by Dataflow.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * One of "drain" or "cancel".  Specifies behavior of deletion during `terraform destroy`.  See above note.
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
