// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./getRule";
export * from "./getTestablePermissions";
export * from "./getWorkloadIdentityPool";
export * from "./getWorkloadIdentityPoolProvider";
export * from "./workloadIdentityPool";
export * from "./workloadIdentityPoolProvider";

// Import resources to register:
import { WorkloadIdentityPool } from "./workloadIdentityPool";
import { WorkloadIdentityPoolProvider } from "./workloadIdentityPoolProvider";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "gcp:iam/workloadIdentityPool:WorkloadIdentityPool":
                return new WorkloadIdentityPool(name, <any>undefined, { urn })
            case "gcp:iam/workloadIdentityPoolProvider:WorkloadIdentityPoolProvider":
                return new WorkloadIdentityPoolProvider(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("gcp", "iam/workloadIdentityPool", _module)
pulumi.runtime.registerResourceModule("gcp", "iam/workloadIdentityPoolProvider", _module)
