// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

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
    public static readonly __pulumiType = 'gcp:folder/iamAuditConfig:IamAuditConfig';

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
     * The configuration for logging of each type of permission. This can be specified multiple times.
     */
    public readonly auditLogConfigs!: pulumi.Output<outputs.folder.IamAuditConfigAuditLogConfig[]>;
    /**
     * The etag of iam policy
     */
    public /*out*/ readonly etag!: pulumi.Output<string>;
    public readonly folder!: pulumi.Output<string>;
    /**
     * Service which will be enabled for audit logging. The special value allServices covers all services.
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
            inputs["folder"] = state ? state.folder : undefined;
            inputs["service"] = state ? state.service : undefined;
        } else {
            const args = argsOrState as IamAuditConfigArgs | undefined;
            if ((!args || args.auditLogConfigs === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'auditLogConfigs'");
            }
            if ((!args || args.folder === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'folder'");
            }
            if ((!args || args.service === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'service'");
            }
            inputs["auditLogConfigs"] = args ? args.auditLogConfigs : undefined;
            inputs["folder"] = args ? args.folder : undefined;
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
     * The configuration for logging of each type of permission. This can be specified multiple times.
     */
    readonly auditLogConfigs?: pulumi.Input<pulumi.Input<inputs.folder.IamAuditConfigAuditLogConfig>[]>;
    /**
     * The etag of iam policy
     */
    readonly etag?: pulumi.Input<string>;
    readonly folder?: pulumi.Input<string>;
    /**
     * Service which will be enabled for audit logging. The special value allServices covers all services.
     */
    readonly service?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a IamAuditConfig resource.
 */
export interface IamAuditConfigArgs {
    /**
     * The configuration for logging of each type of permission. This can be specified multiple times.
     */
    readonly auditLogConfigs: pulumi.Input<pulumi.Input<inputs.folder.IamAuditConfigAuditLogConfig>[]>;
    readonly folder: pulumi.Input<string>;
    /**
     * Service which will be enabled for audit logging. The special value allServices covers all services.
     */
    readonly service: pulumi.Input<string>;
}
