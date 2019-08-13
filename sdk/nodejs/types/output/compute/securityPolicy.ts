// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputApi from "../../../types/input";
import * as outputApi from "../../../types/output";
import * as utilities from "../../../utilities";

export interface SecurityPolicyRule {
    action: string;
    /**
     * An optional description of this security policy. Max size is 2048.
     */
    description?: string;
    match: outputApi.compute.SecurityPolicyRuleMatch;
    preview?: boolean;
    priority: number;
}

export interface SecurityPolicyRuleMatch {
    config: outputApi.compute.SecurityPolicyRuleMatchConfig;
    versionedExpr: string;
}

export interface SecurityPolicyRuleMatchConfig {
    srcIpRanges: string[];
}
