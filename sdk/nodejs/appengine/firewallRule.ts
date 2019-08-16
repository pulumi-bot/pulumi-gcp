// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * A single firewall rule that is evaluated against incoming traffic
 * and provides an action to take on matched requests.
 * 
 * 
 * To get more information about FirewallRule, see:
 * 
 * * [API documentation](https://cloud.google.com/appengine/docs/admin-api/reference/rest/v1/apps.firewall.ingressRules)
 * * How-to Guides
 *     * [Official Documentation](https://cloud.google.com/appengine/docs/standard/python/creating-firewalls#creating_firewall_rules)
 * 
 * ## Example Usage - App Engine Firewall Rule Basic
 * 
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 * 
 * const myProject = new gcp.organizations.Project("myProject", {
 *     orgId: "123456789",
 *     projectId: "test-project",
 * });
 * const app = new gcp.appengine.Application("app", {
 *     locationId: "us-central",
 *     project: myProject.projectId,
 * });
 * const rule = new gcp.appengine.FirewallRule("rule", {
 *     action: "ALLOW",
 *     priority: 1000,
 *     project: app.project,
 *     sourceRange: "*",
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/app_engine_firewall_rule.html.markdown.
 */
export class FirewallRule extends pulumi.CustomResource {
    /**
     * Get an existing FirewallRule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FirewallRuleState, opts?: pulumi.CustomResourceOptions): FirewallRule {
        return new FirewallRule(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:appengine/firewallRule:FirewallRule';

    /**
     * Returns true if the given object is an instance of FirewallRule.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FirewallRule {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === FirewallRule.__pulumiType;
    }

    public readonly action!: pulumi.Output<string>;
    public readonly description!: pulumi.Output<string | undefined>;
    public readonly priority!: pulumi.Output<number | undefined>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    public readonly sourceRange!: pulumi.Output<string>;

    /**
     * Create a FirewallRule resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FirewallRuleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FirewallRuleArgs | FirewallRuleState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as FirewallRuleState | undefined;
            inputs["action"] = state ? state.action : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["priority"] = state ? state.priority : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["sourceRange"] = state ? state.sourceRange : undefined;
        } else {
            const args = argsOrState as FirewallRuleArgs | undefined;
            if (!args || args.action === undefined) {
                throw new Error("Missing required property 'action'");
            }
            if (!args || args.sourceRange === undefined) {
                throw new Error("Missing required property 'sourceRange'");
            }
            inputs["action"] = args ? args.action : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["priority"] = args ? args.priority : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["sourceRange"] = args ? args.sourceRange : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(FirewallRule.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering FirewallRule resources.
 */
export interface FirewallRuleState {
    readonly action?: pulumi.Input<string>;
    readonly description?: pulumi.Input<string>;
    readonly priority?: pulumi.Input<number>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    readonly sourceRange?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a FirewallRule resource.
 */
export interface FirewallRuleArgs {
    readonly action: pulumi.Input<string>;
    readonly description?: pulumi.Input<string>;
    readonly priority?: pulumi.Input<number>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    readonly sourceRange: pulumi.Input<string>;
}
