// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputApi from "../../../types/input";
import * as outputApi from "../../../types/output";
import * as utilities from "../../../utilities";

export interface SecurityScanConfigAuthentication {
    customAccount?: outputApi.compute.SecurityScanConfigAuthenticationCustomAccount;
    googleAccount?: outputApi.compute.SecurityScanConfigAuthenticationGoogleAccount;
}

export interface SecurityScanConfigAuthenticationCustomAccount {
    loginUrl: string;
    password: string;
    username: string;
}

export interface SecurityScanConfigAuthenticationGoogleAccount {
    password: string;
    username: string;
}

export interface SecurityScanConfigSchedule {
    intervalDurationDays: number;
    scheduleTime?: string;
}
