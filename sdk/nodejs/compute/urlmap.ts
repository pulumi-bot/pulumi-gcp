// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * UrlMaps are used to route requests to a backend service based on rules
 * that you define for the host and path of an incoming URL.
 *
 * To get more information about UrlMap, see:
 *
 * * [API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/urlMaps)
 *
 * ## Example Usage
 *
 * ## Import
 *
 * UrlMap can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import gcp:compute/uRLMap:URLMap default projects/{{project}}/global/urlMaps/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:compute/uRLMap:URLMap default {{project}}/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:compute/uRLMap:URLMap default {{name}}
 * ```
 */
export class URLMap extends pulumi.CustomResource {
    /**
     * Get an existing URLMap resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: URLMapState, opts?: pulumi.CustomResourceOptions): URLMap {
        return new URLMap(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:compute/uRLMap:URLMap';

    /**
     * Returns true if the given object is an instance of URLMap.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is URLMap {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === URLMap.__pulumiType;
    }

    /**
     * Creation timestamp in RFC3339 text format.
     */
    public /*out*/ readonly creationTimestamp!: pulumi.Output<string>;
    /**
     * defaultRouteAction takes effect when none of the pathRules or routeRules match. The load balancer performs
     * advanced routing actions like URL rewrites, header transformations, etc. prior to forwarding the request
     * to the selected backend. If defaultRouteAction specifies any weightedBackendServices, defaultService must not be set.
     * Conversely if defaultService is set, defaultRouteAction cannot contain any weightedBackendServices.
     * Only one of defaultRouteAction or defaultUrlRedirect must be set.
     * Structure is documented below.
     */
    public readonly defaultRouteAction!: pulumi.Output<outputs.compute.URLMapDefaultRouteAction | undefined>;
    /**
     * The backend service or backend bucket to use when none of the given paths match.
     */
    public readonly defaultService!: pulumi.Output<string | undefined>;
    /**
     * When none of the specified hostRules match, the request is redirected to a URL specified
     * by defaultUrlRedirect. If defaultUrlRedirect is specified, defaultService or
     * defaultRouteAction must not be set.
     * Structure is documented below.
     */
    public readonly defaultUrlRedirect!: pulumi.Output<outputs.compute.URLMapDefaultUrlRedirect | undefined>;
    /**
     * Description of this test case.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking.
     */
    public /*out*/ readonly fingerprint!: pulumi.Output<string>;
    /**
     * Specifies changes to request and response headers that need to take effect for
     * the selected backendService.
     * headerAction specified here take effect before headerAction in the enclosing
     * HttpRouteRule, PathMatcher and UrlMap.
     * Structure is documented below.
     */
    public readonly headerAction!: pulumi.Output<outputs.compute.URLMapHeaderAction | undefined>;
    /**
     * The list of HostRules to use against the URL.
     * Structure is documented below.
     */
    public readonly hostRules!: pulumi.Output<outputs.compute.URLMapHostRule[] | undefined>;
    /**
     * The unique identifier for the resource.
     */
    public /*out*/ readonly mapId!: pulumi.Output<number>;
    /**
     * The name of the query parameter to match. The query parameter must exist in the
     * request, in the absence of which the request match fails.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The name of the PathMatcher to use to match the path portion of the URL if the
     * hostRule matches the URL's host portion.
     */
    public readonly pathMatchers!: pulumi.Output<outputs.compute.URLMapPathMatcher[] | undefined>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * The URI of the created resource.
     */
    public /*out*/ readonly selfLink!: pulumi.Output<string>;
    /**
     * The list of expected URL mapping tests. Request to update this UrlMap will
     * succeed only if all of the test cases pass. You can specify a maximum of 100
     * tests per UrlMap.
     * Structure is documented below.
     */
    public readonly tests!: pulumi.Output<outputs.compute.URLMapTest[] | undefined>;

    /**
     * Create a URLMap resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: URLMapArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: URLMapArgs | URLMapState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as URLMapState | undefined;
            inputs["creationTimestamp"] = state ? state.creationTimestamp : undefined;
            inputs["defaultRouteAction"] = state ? state.defaultRouteAction : undefined;
            inputs["defaultService"] = state ? state.defaultService : undefined;
            inputs["defaultUrlRedirect"] = state ? state.defaultUrlRedirect : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["fingerprint"] = state ? state.fingerprint : undefined;
            inputs["headerAction"] = state ? state.headerAction : undefined;
            inputs["hostRules"] = state ? state.hostRules : undefined;
            inputs["mapId"] = state ? state.mapId : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["pathMatchers"] = state ? state.pathMatchers : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["selfLink"] = state ? state.selfLink : undefined;
            inputs["tests"] = state ? state.tests : undefined;
        } else {
            const args = argsOrState as URLMapArgs | undefined;
            inputs["defaultRouteAction"] = args ? args.defaultRouteAction : undefined;
            inputs["defaultService"] = args ? args.defaultService : undefined;
            inputs["defaultUrlRedirect"] = args ? args.defaultUrlRedirect : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["headerAction"] = args ? args.headerAction : undefined;
            inputs["hostRules"] = args ? args.hostRules : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["pathMatchers"] = args ? args.pathMatchers : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["tests"] = args ? args.tests : undefined;
            inputs["creationTimestamp"] = undefined /*out*/;
            inputs["fingerprint"] = undefined /*out*/;
            inputs["mapId"] = undefined /*out*/;
            inputs["selfLink"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(URLMap.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering URLMap resources.
 */
export interface URLMapState {
    /**
     * Creation timestamp in RFC3339 text format.
     */
    readonly creationTimestamp?: pulumi.Input<string>;
    /**
     * defaultRouteAction takes effect when none of the pathRules or routeRules match. The load balancer performs
     * advanced routing actions like URL rewrites, header transformations, etc. prior to forwarding the request
     * to the selected backend. If defaultRouteAction specifies any weightedBackendServices, defaultService must not be set.
     * Conversely if defaultService is set, defaultRouteAction cannot contain any weightedBackendServices.
     * Only one of defaultRouteAction or defaultUrlRedirect must be set.
     * Structure is documented below.
     */
    readonly defaultRouteAction?: pulumi.Input<inputs.compute.URLMapDefaultRouteAction>;
    /**
     * The backend service or backend bucket to use when none of the given paths match.
     */
    readonly defaultService?: pulumi.Input<string>;
    /**
     * When none of the specified hostRules match, the request is redirected to a URL specified
     * by defaultUrlRedirect. If defaultUrlRedirect is specified, defaultService or
     * defaultRouteAction must not be set.
     * Structure is documented below.
     */
    readonly defaultUrlRedirect?: pulumi.Input<inputs.compute.URLMapDefaultUrlRedirect>;
    /**
     * Description of this test case.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking.
     */
    readonly fingerprint?: pulumi.Input<string>;
    /**
     * Specifies changes to request and response headers that need to take effect for
     * the selected backendService.
     * headerAction specified here take effect before headerAction in the enclosing
     * HttpRouteRule, PathMatcher and UrlMap.
     * Structure is documented below.
     */
    readonly headerAction?: pulumi.Input<inputs.compute.URLMapHeaderAction>;
    /**
     * The list of HostRules to use against the URL.
     * Structure is documented below.
     */
    readonly hostRules?: pulumi.Input<pulumi.Input<inputs.compute.URLMapHostRule>[]>;
    /**
     * The unique identifier for the resource.
     */
    readonly mapId?: pulumi.Input<number>;
    /**
     * The name of the query parameter to match. The query parameter must exist in the
     * request, in the absence of which the request match fails.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The name of the PathMatcher to use to match the path portion of the URL if the
     * hostRule matches the URL's host portion.
     */
    readonly pathMatchers?: pulumi.Input<pulumi.Input<inputs.compute.URLMapPathMatcher>[]>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * The URI of the created resource.
     */
    readonly selfLink?: pulumi.Input<string>;
    /**
     * The list of expected URL mapping tests. Request to update this UrlMap will
     * succeed only if all of the test cases pass. You can specify a maximum of 100
     * tests per UrlMap.
     * Structure is documented below.
     */
    readonly tests?: pulumi.Input<pulumi.Input<inputs.compute.URLMapTest>[]>;
}

/**
 * The set of arguments for constructing a URLMap resource.
 */
export interface URLMapArgs {
    /**
     * defaultRouteAction takes effect when none of the pathRules or routeRules match. The load balancer performs
     * advanced routing actions like URL rewrites, header transformations, etc. prior to forwarding the request
     * to the selected backend. If defaultRouteAction specifies any weightedBackendServices, defaultService must not be set.
     * Conversely if defaultService is set, defaultRouteAction cannot contain any weightedBackendServices.
     * Only one of defaultRouteAction or defaultUrlRedirect must be set.
     * Structure is documented below.
     */
    readonly defaultRouteAction?: pulumi.Input<inputs.compute.URLMapDefaultRouteAction>;
    /**
     * The backend service or backend bucket to use when none of the given paths match.
     */
    readonly defaultService?: pulumi.Input<string>;
    /**
     * When none of the specified hostRules match, the request is redirected to a URL specified
     * by defaultUrlRedirect. If defaultUrlRedirect is specified, defaultService or
     * defaultRouteAction must not be set.
     * Structure is documented below.
     */
    readonly defaultUrlRedirect?: pulumi.Input<inputs.compute.URLMapDefaultUrlRedirect>;
    /**
     * Description of this test case.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Specifies changes to request and response headers that need to take effect for
     * the selected backendService.
     * headerAction specified here take effect before headerAction in the enclosing
     * HttpRouteRule, PathMatcher and UrlMap.
     * Structure is documented below.
     */
    readonly headerAction?: pulumi.Input<inputs.compute.URLMapHeaderAction>;
    /**
     * The list of HostRules to use against the URL.
     * Structure is documented below.
     */
    readonly hostRules?: pulumi.Input<pulumi.Input<inputs.compute.URLMapHostRule>[]>;
    /**
     * The name of the query parameter to match. The query parameter must exist in the
     * request, in the absence of which the request match fails.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The name of the PathMatcher to use to match the path portion of the URL if the
     * hostRule matches the URL's host portion.
     */
    readonly pathMatchers?: pulumi.Input<pulumi.Input<inputs.compute.URLMapPathMatcher>[]>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * The list of expected URL mapping tests. Request to update this UrlMap will
     * succeed only if all of the test cases pass. You can specify a maximum of 100
     * tests per UrlMap.
     * Structure is documented below.
     */
    readonly tests?: pulumi.Input<pulumi.Input<inputs.compute.URLMapTest>[]>;
}
