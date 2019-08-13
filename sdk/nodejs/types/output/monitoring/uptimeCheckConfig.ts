// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputApi from "../../../types/input";
import * as outputApi from "../../../types/output";
import * as utilities from "../../../utilities";

export interface UptimeCheckConfigContentMatcher {
    content?: string;
}

export interface UptimeCheckConfigHttpCheck {
    authInfo?: outputApi.monitoring.UptimeCheckConfigHttpCheckAuthInfo;
    headers?: {[key: string]: string};
    maskHeaders?: boolean;
    path?: string;
    port: number;
    useSsl?: boolean;
}

export interface UptimeCheckConfigHttpCheckAuthInfo {
    password?: string;
    username?: string;
}

export interface UptimeCheckConfigInternalChecker {
    displayName?: string;
    gcpZone?: string;
    name?: string;
    network?: string;
    peerProjectId?: string;
}

export interface UptimeCheckConfigMonitoredResource {
    labels: {[key: string]: string};
    type: string;
}

export interface UptimeCheckConfigResourceGroup {
    groupId?: string;
    resourceType?: string;
}

export interface UptimeCheckConfigTcpCheck {
    port: number;
}
