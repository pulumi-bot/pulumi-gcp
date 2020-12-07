// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./accountIamBinding";
export * from "./accountIamMember";
export * from "./accountIamPolicy";
export * from "./budget";

// Import resources to register:
import { AccountIamBinding } from "./accountIamBinding";
import { AccountIamMember } from "./accountIamMember";
import { AccountIamPolicy } from "./accountIamPolicy";
import { Budget } from "./budget";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "gcp:billing/accountIamBinding:AccountIamBinding":
                return new AccountIamBinding(name, <any>undefined, { urn })
            case "gcp:billing/accountIamMember:AccountIamMember":
                return new AccountIamMember(name, <any>undefined, { urn })
            case "gcp:billing/accountIamPolicy:AccountIamPolicy":
                return new AccountIamPolicy(name, <any>undefined, { urn })
            case "gcp:billing/budget:Budget":
                return new Budget(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("gcp", "billing/accountIamBinding", _module)
pulumi.runtime.registerResourceModule("gcp", "billing/accountIamMember", _module)
pulumi.runtime.registerResourceModule("gcp", "billing/accountIamPolicy", _module)
pulumi.runtime.registerResourceModule("gcp", "billing/budget", _module)
