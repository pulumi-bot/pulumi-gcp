// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputApi from "../../../types/input";
import * as outputApi from "../../../types/output";
import * as utilities from "../../../utilities";

export interface NodeNetworkEndpoint {
    ipAddress?: pulumi.Input<string>;
    port?: pulumi.Input<number>;
}

export interface NodeSchedulingConfig {
    preemptible?: pulumi.Input<boolean>;
}
