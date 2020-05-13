// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * A ScanConfig resource contains the configurations to launch a scan.
 * 
 * To get more information about ScanConfig, see:
 * 
 * * [API documentation](https://cloud.google.com/security-scanner/docs/reference/rest/v1beta/projects.scanConfigs)
 * * How-to Guides
 *     * [Using Cloud Security Scanner](https://cloud.google.com/security-scanner/docs/scanning)
 * 
 * > **Warning:** All arguments including `authentication.google_account.password` and `authentication.custom_account.password` will be stored in the raw
 * state as plain-text.[Read more about secrets in state](https://www.pulumi.com/docs/intro/concepts/programming-model/#secrets)
 * 
 * ## Example Usage - Scan Config Basic
 * 
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 * 
 * const scannerStaticIp = new gcp.compute.Address("scannerStaticIp", {});
 * const scan-config = new gcp.compute.SecurityScanConfig("scan-config", {
 *     displayName: "scan-config",
 *     startingUrls: [pulumi.interpolate`http://${scannerStaticIp.address}`],
 *     targetPlatforms: ["COMPUTE"],
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/security_scanner_scan_config.html.markdown.
 */
export class SecurityScanConfig extends pulumi.CustomResource {
    /**
     * Get an existing SecurityScanConfig resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SecurityScanConfigState, opts?: pulumi.CustomResourceOptions): SecurityScanConfig {
        return new SecurityScanConfig(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:compute/securityScanConfig:SecurityScanConfig';

    /**
     * Returns true if the given object is an instance of SecurityScanConfig.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SecurityScanConfig {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SecurityScanConfig.__pulumiType;
    }

    /**
     * The authentication configuration.
     * If specified, service will use the authentication configuration during scanning.  Structure is documented below.
     */
    public readonly authentication!: pulumi.Output<outputs.compute.SecurityScanConfigAuthentication | undefined>;
    /**
     * The blacklist URL patterns as described in
     * https://cloud.google.com/security-scanner/docs/excluded-urls
     */
    public readonly blacklistPatterns!: pulumi.Output<string[] | undefined>;
    /**
     * The user provider display name of the ScanConfig.
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * Controls export of scan configurations and results to Cloud Security Command Center.
     */
    public readonly exportToSecurityCommandCenter!: pulumi.Output<string | undefined>;
    /**
     * The maximum QPS during scanning. A valid value ranges from 5 to 20 inclusively.
     * Defaults to 15.
     */
    public readonly maxQps!: pulumi.Output<number | undefined>;
    /**
     * A server defined name for this index. Format: 'projects/{{project}}/scanConfigs/{{server_generated_id}}'
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * The schedule of the ScanConfig  Structure is documented below.
     */
    public readonly schedule!: pulumi.Output<outputs.compute.SecurityScanConfigSchedule | undefined>;
    /**
     * The starting URLs from which the scanner finds site pages.
     */
    public readonly startingUrls!: pulumi.Output<string[]>;
    /**
     * Set of Cloud Platforms targeted by the scan. If empty, APP_ENGINE will be used as a default.
     */
    public readonly targetPlatforms!: pulumi.Output<string[] | undefined>;
    /**
     * Type of the user agents used for scanning
     */
    public readonly userAgent!: pulumi.Output<string | undefined>;

    /**
     * Create a SecurityScanConfig resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SecurityScanConfigArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SecurityScanConfigArgs | SecurityScanConfigState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as SecurityScanConfigState | undefined;
            inputs["authentication"] = state ? state.authentication : undefined;
            inputs["blacklistPatterns"] = state ? state.blacklistPatterns : undefined;
            inputs["displayName"] = state ? state.displayName : undefined;
            inputs["exportToSecurityCommandCenter"] = state ? state.exportToSecurityCommandCenter : undefined;
            inputs["maxQps"] = state ? state.maxQps : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["schedule"] = state ? state.schedule : undefined;
            inputs["startingUrls"] = state ? state.startingUrls : undefined;
            inputs["targetPlatforms"] = state ? state.targetPlatforms : undefined;
            inputs["userAgent"] = state ? state.userAgent : undefined;
        } else {
            const args = argsOrState as SecurityScanConfigArgs | undefined;
            if (!args || args.displayName === undefined) {
                throw new Error("Missing required property 'displayName'");
            }
            if (!args || args.startingUrls === undefined) {
                throw new Error("Missing required property 'startingUrls'");
            }
            inputs["authentication"] = args ? args.authentication : undefined;
            inputs["blacklistPatterns"] = args ? args.blacklistPatterns : undefined;
            inputs["displayName"] = args ? args.displayName : undefined;
            inputs["exportToSecurityCommandCenter"] = args ? args.exportToSecurityCommandCenter : undefined;
            inputs["maxQps"] = args ? args.maxQps : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["schedule"] = args ? args.schedule : undefined;
            inputs["startingUrls"] = args ? args.startingUrls : undefined;
            inputs["targetPlatforms"] = args ? args.targetPlatforms : undefined;
            inputs["userAgent"] = args ? args.userAgent : undefined;
            inputs["name"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(SecurityScanConfig.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SecurityScanConfig resources.
 */
export interface SecurityScanConfigState {
    /**
     * The authentication configuration.
     * If specified, service will use the authentication configuration during scanning.  Structure is documented below.
     */
    readonly authentication?: pulumi.Input<inputs.compute.SecurityScanConfigAuthentication>;
    /**
     * The blacklist URL patterns as described in
     * https://cloud.google.com/security-scanner/docs/excluded-urls
     */
    readonly blacklistPatterns?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The user provider display name of the ScanConfig.
     */
    readonly displayName?: pulumi.Input<string>;
    /**
     * Controls export of scan configurations and results to Cloud Security Command Center.
     */
    readonly exportToSecurityCommandCenter?: pulumi.Input<string>;
    /**
     * The maximum QPS during scanning. A valid value ranges from 5 to 20 inclusively.
     * Defaults to 15.
     */
    readonly maxQps?: pulumi.Input<number>;
    /**
     * A server defined name for this index. Format: 'projects/{{project}}/scanConfigs/{{server_generated_id}}'
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * The schedule of the ScanConfig  Structure is documented below.
     */
    readonly schedule?: pulumi.Input<inputs.compute.SecurityScanConfigSchedule>;
    /**
     * The starting URLs from which the scanner finds site pages.
     */
    readonly startingUrls?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Set of Cloud Platforms targeted by the scan. If empty, APP_ENGINE will be used as a default.
     */
    readonly targetPlatforms?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Type of the user agents used for scanning
     */
    readonly userAgent?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SecurityScanConfig resource.
 */
export interface SecurityScanConfigArgs {
    /**
     * The authentication configuration.
     * If specified, service will use the authentication configuration during scanning.  Structure is documented below.
     */
    readonly authentication?: pulumi.Input<inputs.compute.SecurityScanConfigAuthentication>;
    /**
     * The blacklist URL patterns as described in
     * https://cloud.google.com/security-scanner/docs/excluded-urls
     */
    readonly blacklistPatterns?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The user provider display name of the ScanConfig.
     */
    readonly displayName: pulumi.Input<string>;
    /**
     * Controls export of scan configurations and results to Cloud Security Command Center.
     */
    readonly exportToSecurityCommandCenter?: pulumi.Input<string>;
    /**
     * The maximum QPS during scanning. A valid value ranges from 5 to 20 inclusively.
     * Defaults to 15.
     */
    readonly maxQps?: pulumi.Input<number>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * The schedule of the ScanConfig  Structure is documented below.
     */
    readonly schedule?: pulumi.Input<inputs.compute.SecurityScanConfigSchedule>;
    /**
     * The starting URLs from which the scanner finds site pages.
     */
    readonly startingUrls: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Set of Cloud Platforms targeted by the scan. If empty, APP_ENGINE will be used as a default.
     */
    readonly targetPlatforms?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Type of the user agents used for scanning
     */
    readonly userAgent?: pulumi.Input<string>;
}
