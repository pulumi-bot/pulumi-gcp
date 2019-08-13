// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputApi from "../../../types/input";
import * as outputApi from "../../../types/output";
import * as utilities from "../../../utilities";

export interface JobAppEngineHttpTarget {
    appEngineRouting?: pulumi.Input<inputApi.cloudscheduler.JobAppEngineHttpTargetAppEngineRouting>;
    body?: pulumi.Input<string>;
    headers?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    httpMethod?: pulumi.Input<string>;
    relativeUri: pulumi.Input<string>;
}

export interface JobAppEngineHttpTargetAppEngineRouting {
    instance?: pulumi.Input<string>;
    service?: pulumi.Input<string>;
    version?: pulumi.Input<string>;
}

export interface JobHttpTarget {
    body?: pulumi.Input<string>;
    headers?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    httpMethod?: pulumi.Input<string>;
    uri: pulumi.Input<string>;
}

export interface JobPubsubTarget {
    attributes?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    data?: pulumi.Input<string>;
    topicName: pulumi.Input<string>;
}

export interface JobRetryConfig {
    maxBackoffDuration?: pulumi.Input<string>;
    maxDoublings?: pulumi.Input<number>;
    maxRetryDuration?: pulumi.Input<string>;
    minBackoffDuration?: pulumi.Input<string>;
    retryCount?: pulumi.Input<number>;
}
