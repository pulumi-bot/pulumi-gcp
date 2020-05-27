// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Represents a reservation resource. A reservation ensures that capacity is
 * held in a specific zone even if the reserved VMs are not running.
 *
 * Reservations apply only to Compute Engine, Cloud Dataproc, and Google
 * Kubernetes Engine VM usage.Reservations do not apply to `f1-micro` or
 * `g1-small` machine types, preemptible VMs, sole tenant nodes, or other
 * services not listed above
 * like Cloud SQL and Dataflow.
 *
 *
 * To get more information about Reservation, see:
 *
 * * [API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/reservations)
 * * How-to Guides
 *     * [Reserving zonal resources](https://cloud.google.com/compute/docs/instances/reserving-zonal-resources)
 *
 * ## Example Usage - Reservation Basic
 * {{% example %}}
 *
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const gceReservation = new gcp.compute.Reservation("gceReservation", {
 *     specificReservation: {
 *         count: 1,
 *         instanceProperties: {
 *             machineType: "n2-standard-2",
 *             minCpuPlatform: "Intel Cascade Lake",
 *         },
 *     },
 *     zone: "us-central1-a",
 * });
 * ```
 *
 * {{% /example %}}
 */
export class Reservation extends pulumi.CustomResource {
    /**
     * Get an existing Reservation resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ReservationState, opts?: pulumi.CustomResourceOptions): Reservation {
        return new Reservation(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:compute/reservation:Reservation';

    /**
     * Returns true if the given object is an instance of Reservation.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Reservation {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Reservation.__pulumiType;
    }

    /**
     * Full or partial URL to a parent commitment. This field displays for reservations that are tied to a commitment.
     */
    public /*out*/ readonly commitment!: pulumi.Output<string>;
    /**
     * Creation timestamp in RFC3339 text format.
     */
    public /*out*/ readonly creationTimestamp!: pulumi.Output<string>;
    /**
     * An optional description of this resource.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Name of the resource. Provided by the client when the resource is
     * created. The name must be 1-63 characters long, and comply with
     * RFC1035. Specifically, the name must be 1-63 characters long and match
     * the regular expression `a-z?` which means the
     * first character must be a lowercase letter, and all following
     * characters must be a dash, lowercase letter, or digit, except the last
     * character, which cannot be a dash.
     */
    public readonly name!: pulumi.Output<string>;
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
     * Reservation for instances with specific machine shapes.  Structure is documented below.
     */
    public readonly specificReservation!: pulumi.Output<outputs.compute.ReservationSpecificReservation>;
    /**
     * When set to true, only VMs that target this reservation by name can
     * consume this reservation. Otherwise, it can be consumed by VMs with
     * affinity for any reservation. Defaults to false.
     */
    public readonly specificReservationRequired!: pulumi.Output<boolean | undefined>;
    /**
     * The status of the reservation.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * The zone where the reservation is made.
     */
    public readonly zone!: pulumi.Output<string>;

    /**
     * Create a Reservation resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ReservationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ReservationArgs | ReservationState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ReservationState | undefined;
            inputs["commitment"] = state ? state.commitment : undefined;
            inputs["creationTimestamp"] = state ? state.creationTimestamp : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["selfLink"] = state ? state.selfLink : undefined;
            inputs["specificReservation"] = state ? state.specificReservation : undefined;
            inputs["specificReservationRequired"] = state ? state.specificReservationRequired : undefined;
            inputs["status"] = state ? state.status : undefined;
            inputs["zone"] = state ? state.zone : undefined;
        } else {
            const args = argsOrState as ReservationArgs | undefined;
            if (!args || args.specificReservation === undefined) {
                throw new Error("Missing required property 'specificReservation'");
            }
            if (!args || args.zone === undefined) {
                throw new Error("Missing required property 'zone'");
            }
            inputs["description"] = args ? args.description : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["specificReservation"] = args ? args.specificReservation : undefined;
            inputs["specificReservationRequired"] = args ? args.specificReservationRequired : undefined;
            inputs["zone"] = args ? args.zone : undefined;
            inputs["commitment"] = undefined /*out*/;
            inputs["creationTimestamp"] = undefined /*out*/;
            inputs["selfLink"] = undefined /*out*/;
            inputs["status"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Reservation.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Reservation resources.
 */
export interface ReservationState {
    /**
     * Full or partial URL to a parent commitment. This field displays for reservations that are tied to a commitment.
     */
    readonly commitment?: pulumi.Input<string>;
    /**
     * Creation timestamp in RFC3339 text format.
     */
    readonly creationTimestamp?: pulumi.Input<string>;
    /**
     * An optional description of this resource.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Name of the resource. Provided by the client when the resource is
     * created. The name must be 1-63 characters long, and comply with
     * RFC1035. Specifically, the name must be 1-63 characters long and match
     * the regular expression `a-z?` which means the
     * first character must be a lowercase letter, and all following
     * characters must be a dash, lowercase letter, or digit, except the last
     * character, which cannot be a dash.
     */
    readonly name?: pulumi.Input<string>;
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
     * Reservation for instances with specific machine shapes.  Structure is documented below.
     */
    readonly specificReservation?: pulumi.Input<inputs.compute.ReservationSpecificReservation>;
    /**
     * When set to true, only VMs that target this reservation by name can
     * consume this reservation. Otherwise, it can be consumed by VMs with
     * affinity for any reservation. Defaults to false.
     */
    readonly specificReservationRequired?: pulumi.Input<boolean>;
    /**
     * The status of the reservation.
     */
    readonly status?: pulumi.Input<string>;
    /**
     * The zone where the reservation is made.
     */
    readonly zone?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Reservation resource.
 */
export interface ReservationArgs {
    /**
     * An optional description of this resource.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Name of the resource. Provided by the client when the resource is
     * created. The name must be 1-63 characters long, and comply with
     * RFC1035. Specifically, the name must be 1-63 characters long and match
     * the regular expression `a-z?` which means the
     * first character must be a lowercase letter, and all following
     * characters must be a dash, lowercase letter, or digit, except the last
     * character, which cannot be a dash.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * Reservation for instances with specific machine shapes.  Structure is documented below.
     */
    readonly specificReservation: pulumi.Input<inputs.compute.ReservationSpecificReservation>;
    /**
     * When set to true, only VMs that target this reservation by name can
     * consume this reservation. Otherwise, it can be consumed by VMs with
     * affinity for any reservation. Defaults to false.
     */
    readonly specificReservationRequired?: pulumi.Input<boolean>;
    /**
     * The zone where the reservation is made.
     */
    readonly zone: pulumi.Input<string>;
}
