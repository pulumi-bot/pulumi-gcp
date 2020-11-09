// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * A scheduled job that can publish a pubsub message or a http request
 * every X interval of time, using crontab format string.
 *
 * To use Cloud Scheduler your project must contain an App Engine app
 * that is located in one of the supported regions. If your project
 * does not have an App Engine app, you must create one.
 *
 * To get more information about Job, see:
 *
 * * [API documentation](https://cloud.google.com/scheduler/docs/reference/rest/)
 * * How-to Guides
 *     * [Official Documentation](https://cloud.google.com/scheduler/)
 *
 * ## Example Usage
 *
 * ## Import
 *
 * Job can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import gcp:cloudscheduler/job:Job default projects/{{project}}/locations/{{region}}/jobs/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:cloudscheduler/job:Job default {{project}}/{{region}}/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:cloudscheduler/job:Job default {{region}}/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:cloudscheduler/job:Job default {{name}}
 * ```
 */
export class Job extends pulumi.CustomResource {
    /**
     * Get an existing Job resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: JobState, opts?: pulumi.CustomResourceOptions): Job {
        return new Job(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:cloudscheduler/job:Job';

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
     * App Engine HTTP target.
     * If the job providers a App Engine HTTP target the cron will
     * send a request to the service instance
     * Structure is documented below.
     */
    public readonly appEngineHttpTarget!: pulumi.Output<outputs.cloudscheduler.JobAppEngineHttpTarget | undefined>;
    /**
     * The deadline for job attempts. If the request handler does not respond by this deadline then the request is
     * cancelled and the attempt is marked as a DEADLINE_EXCEEDED failure. The failed attempt can be viewed in
     * execution logs. Cloud Scheduler will retry the job according to the RetryConfig.
     * The allowed duration for this deadline is:
     * * For HTTP targets, between 15 seconds and 30 minutes.
     * * For App Engine HTTP targets, between 15 seconds and 24 hours.
     * * **Note**: For PubSub targets, this field is ignored - setting it will introduce an unresolvable diff.
     * A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s"
     */
    public readonly attemptDeadline!: pulumi.Output<string | undefined>;
    /**
     * A human-readable description for the job.
     * This string must not contain more than 500 characters.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * HTTP target.
     * If the job providers a httpTarget the cron will
     * send a request to the targeted url
     * Structure is documented below.
     */
    public readonly httpTarget!: pulumi.Output<outputs.cloudscheduler.JobHttpTarget | undefined>;
    /**
     * The name of the job.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * Pub/Sub target
     * If the job providers a Pub/Sub target the cron will publish
     * a message to the provided topic
     * Structure is documented below.
     */
    public readonly pubsubTarget!: pulumi.Output<outputs.cloudscheduler.JobPubsubTarget | undefined>;
    /**
     * Region where the scheduler job resides. If it is not provided, this provider will use the provider default.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * By default, if a job does not complete successfully,
     * meaning that an acknowledgement is not received from the handler,
     * then it will be retried with exponential backoff according to the settings
     * Structure is documented below.
     */
    public readonly retryConfig!: pulumi.Output<outputs.cloudscheduler.JobRetryConfig | undefined>;
    /**
     * Describes the schedule on which the job will be executed.
     */
    public readonly schedule!: pulumi.Output<string | undefined>;
    /**
     * Specifies the time zone to be used in interpreting schedule.
     * The value of this field must be a time zone name from the tz database.
     */
    public readonly timeZone!: pulumi.Output<string | undefined>;

    /**
     * Create a Job resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: JobArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: JobArgs | JobState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as JobState | undefined;
            inputs["appEngineHttpTarget"] = state ? state.appEngineHttpTarget : undefined;
            inputs["attemptDeadline"] = state ? state.attemptDeadline : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["httpTarget"] = state ? state.httpTarget : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["pubsubTarget"] = state ? state.pubsubTarget : undefined;
            inputs["region"] = state ? state.region : undefined;
            inputs["retryConfig"] = state ? state.retryConfig : undefined;
            inputs["schedule"] = state ? state.schedule : undefined;
            inputs["timeZone"] = state ? state.timeZone : undefined;
        } else {
            const args = argsOrState as JobArgs | undefined;
            inputs["appEngineHttpTarget"] = args ? args.appEngineHttpTarget : undefined;
            inputs["attemptDeadline"] = args ? args.attemptDeadline : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["httpTarget"] = args ? args.httpTarget : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["pubsubTarget"] = args ? args.pubsubTarget : undefined;
            inputs["region"] = args ? args.region : undefined;
            inputs["retryConfig"] = args ? args.retryConfig : undefined;
            inputs["schedule"] = args ? args.schedule : undefined;
            inputs["timeZone"] = args ? args.timeZone : undefined;
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
     * App Engine HTTP target.
     * If the job providers a App Engine HTTP target the cron will
     * send a request to the service instance
     * Structure is documented below.
     */
    readonly appEngineHttpTarget?: pulumi.Input<inputs.cloudscheduler.JobAppEngineHttpTarget>;
    /**
     * The deadline for job attempts. If the request handler does not respond by this deadline then the request is
     * cancelled and the attempt is marked as a DEADLINE_EXCEEDED failure. The failed attempt can be viewed in
     * execution logs. Cloud Scheduler will retry the job according to the RetryConfig.
     * The allowed duration for this deadline is:
     * * For HTTP targets, between 15 seconds and 30 minutes.
     * * For App Engine HTTP targets, between 15 seconds and 24 hours.
     * * **Note**: For PubSub targets, this field is ignored - setting it will introduce an unresolvable diff.
     * A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s"
     */
    readonly attemptDeadline?: pulumi.Input<string>;
    /**
     * A human-readable description for the job.
     * This string must not contain more than 500 characters.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * HTTP target.
     * If the job providers a httpTarget the cron will
     * send a request to the targeted url
     * Structure is documented below.
     */
    readonly httpTarget?: pulumi.Input<inputs.cloudscheduler.JobHttpTarget>;
    /**
     * The name of the job.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * Pub/Sub target
     * If the job providers a Pub/Sub target the cron will publish
     * a message to the provided topic
     * Structure is documented below.
     */
    readonly pubsubTarget?: pulumi.Input<inputs.cloudscheduler.JobPubsubTarget>;
    /**
     * Region where the scheduler job resides. If it is not provided, this provider will use the provider default.
     */
    readonly region?: pulumi.Input<string>;
    /**
     * By default, if a job does not complete successfully,
     * meaning that an acknowledgement is not received from the handler,
     * then it will be retried with exponential backoff according to the settings
     * Structure is documented below.
     */
    readonly retryConfig?: pulumi.Input<inputs.cloudscheduler.JobRetryConfig>;
    /**
     * Describes the schedule on which the job will be executed.
     */
    readonly schedule?: pulumi.Input<string>;
    /**
     * Specifies the time zone to be used in interpreting schedule.
     * The value of this field must be a time zone name from the tz database.
     */
    readonly timeZone?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Job resource.
 */
export interface JobArgs {
    /**
     * App Engine HTTP target.
     * If the job providers a App Engine HTTP target the cron will
     * send a request to the service instance
     * Structure is documented below.
     */
    readonly appEngineHttpTarget?: pulumi.Input<inputs.cloudscheduler.JobAppEngineHttpTarget>;
    /**
     * The deadline for job attempts. If the request handler does not respond by this deadline then the request is
     * cancelled and the attempt is marked as a DEADLINE_EXCEEDED failure. The failed attempt can be viewed in
     * execution logs. Cloud Scheduler will retry the job according to the RetryConfig.
     * The allowed duration for this deadline is:
     * * For HTTP targets, between 15 seconds and 30 minutes.
     * * For App Engine HTTP targets, between 15 seconds and 24 hours.
     * * **Note**: For PubSub targets, this field is ignored - setting it will introduce an unresolvable diff.
     * A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s"
     */
    readonly attemptDeadline?: pulumi.Input<string>;
    /**
     * A human-readable description for the job.
     * This string must not contain more than 500 characters.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * HTTP target.
     * If the job providers a httpTarget the cron will
     * send a request to the targeted url
     * Structure is documented below.
     */
    readonly httpTarget?: pulumi.Input<inputs.cloudscheduler.JobHttpTarget>;
    /**
     * The name of the job.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * Pub/Sub target
     * If the job providers a Pub/Sub target the cron will publish
     * a message to the provided topic
     * Structure is documented below.
     */
    readonly pubsubTarget?: pulumi.Input<inputs.cloudscheduler.JobPubsubTarget>;
    /**
     * Region where the scheduler job resides. If it is not provided, this provider will use the provider default.
     */
    readonly region?: pulumi.Input<string>;
    /**
     * By default, if a job does not complete successfully,
     * meaning that an acknowledgement is not received from the handler,
     * then it will be retried with exponential backoff according to the settings
     * Structure is documented below.
     */
    readonly retryConfig?: pulumi.Input<inputs.cloudscheduler.JobRetryConfig>;
    /**
     * Describes the schedule on which the job will be executed.
     */
    readonly schedule?: pulumi.Input<string>;
    /**
     * Specifies the time zone to be used in interpreting schedule.
     * The value of this field must be a time zone name from the tz database.
     */
    readonly timeZone?: pulumi.Input<string>;
}
