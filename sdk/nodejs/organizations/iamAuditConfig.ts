// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Allows management of audit logging config for a given service for a Google Cloud Platform Organization.
 *
 * {{% examples %}}
 * ## Example Usage
 * {{% example %}}
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const config = new gcp.organizations.IamAuditConfig("config", {
 *     auditLogConfigs: [{
 *         exemptedMembers: ["user:joebloggs@hashicorp.com"],
 *         logType: "DATA_READ",
 *     }],
 *     orgId: "your-organization-id",
 *     service: "allServices",
 * });
 * ```
 * {{% /example %}}
 * {{% /examples %}}
 */
export class IamAuditConfig extends pulumi.CustomResource {
    /**
     * Get an existing IamAuditConfig resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: IamAuditConfigState, opts?: pulumi.CustomResourceOptions): IamAuditConfig {
        return new IamAuditConfig(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:organizations/iamAuditConfig:IamAuditConfig';

    /**
     * Returns true if the given object is an instance of IamAuditConfig.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is IamAuditConfig {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === IamAuditConfig.__pulumiType;
    }

    /**
     * The configuration for logging of each type of permission.  This can be specified multiple times.  Structure is documented below.
     */
    public readonly auditLogConfigs!: pulumi.Output<outputs.organizations.IamAuditConfigAuditLogConfig[]>;
    public /*out*/ readonly etag!: pulumi.Output<string>;
    /**
     * The numeric ID of the organization in which you want to manage the audit logging config.
     */
    public readonly orgId!: pulumi.Output<string>;
    /**
     * Service which will be enabled for audit logging.  The special value `allServices` covers all services.  Note that if there are google\_organization\_iam\_audit\_config resources covering both `allServices` and a specific service then the union of the two AuditConfigs is used for that service: the `logTypes` specified in each `auditLogConfig` are enabled, and the `exemptedMembers` in each `auditLogConfig` are exempted.
     */
    public readonly service!: pulumi.Output<string>;

    /**
     * Create a IamAuditConfig resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: IamAuditConfigArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: IamAuditConfigArgs | IamAuditConfigState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as IamAuditConfigState | undefined;
            inputs["auditLogConfigs"] = state ? state.auditLogConfigs : undefined;
            inputs["etag"] = state ? state.etag : undefined;
            inputs["orgId"] = state ? state.orgId : undefined;
            inputs["service"] = state ? state.service : undefined;
        } else {
            const args = argsOrState as IamAuditConfigArgs | undefined;
            if (!args || args.auditLogConfigs === undefined) {
                throw new Error("Missing required property 'auditLogConfigs'");
            }
            if (!args || args.orgId === undefined) {
                throw new Error("Missing required property 'orgId'");
            }
            if (!args || args.service === undefined) {
                throw new Error("Missing required property 'service'");
            }
            inputs["auditLogConfigs"] = args ? args.auditLogConfigs : undefined;
            inputs["orgId"] = args ? args.orgId : undefined;
            inputs["service"] = args ? args.service : undefined;
            inputs["etag"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(IamAuditConfig.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering IamAuditConfig resources.
 */
export interface IamAuditConfigState {
    /**
     * The configuration for logging of each type of permission.  This can be specified multiple times.  Structure is documented below.
     */
    readonly auditLogConfigs?: pulumi.Input<pulumi.Input<inputs.organizations.IamAuditConfigAuditLogConfig>[]>;
    readonly etag?: pulumi.Input<string>;
    /**
     * The numeric ID of the organization in which you want to manage the audit logging config.
     */
    readonly orgId?: pulumi.Input<string>;
    /**
     * Service which will be enabled for audit logging.  The special value `allServices` covers all services.  Note that if there are google\_organization\_iam\_audit\_config resources covering both `allServices` and a specific service then the union of the two AuditConfigs is used for that service: the `logTypes` specified in each `auditLogConfig` are enabled, and the `exemptedMembers` in each `auditLogConfig` are exempted.
     */
    readonly service?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a IamAuditConfig resource.
 */
export interface IamAuditConfigArgs {
    /**
     * The configuration for logging of each type of permission.  This can be specified multiple times.  Structure is documented below.
     */
    readonly auditLogConfigs: pulumi.Input<pulumi.Input<inputs.organizations.IamAuditConfigAuditLogConfig>[]>;
    /**
     * The numeric ID of the organization in which you want to manage the audit logging config.
     */
    readonly orgId: pulumi.Input<string>;
    /**
     * Service which will be enabled for audit logging.  The special value `allServices` covers all services.  Note that if there are google\_organization\_iam\_audit\_config resources covering both `allServices` and a specific service then the union of the two AuditConfigs is used for that service: the `logTypes` specified in each `auditLogConfig` are enabled, and the `exemptedMembers` in each `auditLogConfig` are exempted.
     */
    readonly service: pulumi.Input<string>;
}
