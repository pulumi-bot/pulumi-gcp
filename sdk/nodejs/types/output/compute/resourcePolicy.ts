// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputApi from "../../../types/input";
import * as outputApi from "../../../types/output";
import * as utilities from "../../../utilities";

export interface ResourcePolicySnapshotSchedulePolicy {
    retentionPolicy?: outputApi.compute.ResourcePolicySnapshotSchedulePolicyRetentionPolicy;
    schedule: outputApi.compute.ResourcePolicySnapshotSchedulePolicySchedule;
    snapshotProperties?: outputApi.compute.ResourcePolicySnapshotSchedulePolicySnapshotProperties;
}

export interface ResourcePolicySnapshotSchedulePolicyRetentionPolicy {
    maxRetentionDays: number;
    onSourceDiskDelete?: string;
}

export interface ResourcePolicySnapshotSchedulePolicySchedule {
    dailySchedule?: outputApi.compute.ResourcePolicySnapshotSchedulePolicyScheduleDailySchedule;
    hourlySchedule?: outputApi.compute.ResourcePolicySnapshotSchedulePolicyScheduleHourlySchedule;
    weeklySchedule?: outputApi.compute.ResourcePolicySnapshotSchedulePolicyScheduleWeeklySchedule;
}

export interface ResourcePolicySnapshotSchedulePolicyScheduleDailySchedule {
    daysInCycle: number;
    startTime: string;
}

export interface ResourcePolicySnapshotSchedulePolicyScheduleHourlySchedule {
    hoursInCycle: number;
    startTime: string;
}

export interface ResourcePolicySnapshotSchedulePolicyScheduleWeeklySchedule {
    dayOfWeeks: outputApi.compute.ResourcePolicySnapshotSchedulePolicyScheduleWeeklyScheduleDayOfWeek[];
}

export interface ResourcePolicySnapshotSchedulePolicyScheduleWeeklyScheduleDayOfWeek {
    day: string;
    startTime: string;
}

export interface ResourcePolicySnapshotSchedulePolicySnapshotProperties {
    guestFlush?: boolean;
    labels?: {[key: string]: string};
    storageLocations?: string;
}
