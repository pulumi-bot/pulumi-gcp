// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * A reservation is a mechanism used to guarantee BigQuery slots to users.
 * 
 * To get more information about Reservation, see:
 * 
 * * [API documentation](https://cloud.google.com/bigquery/docs/reference/reservations/rest/v1beta1/projects.locations.reservations/create)
 * * How-to Guides
 *     * [Introduction to Reservations](https://cloud.google.com/bigquery/docs/reservations-intro)
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/bigquery_reservation.html.markdown.
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
    public static readonly __pulumiType = 'gcp:bigquery/reservation:Reservation';

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
     * -
     * (Optional)
     * If false, any query using this reservation will use idle slots from other reservations within
     * the same admin project. If true, a query using this reservation will execute with the slot
     * capacity specified above at most.
     */
    public readonly ignoreIdleSlots!: pulumi.Output<boolean | undefined>;
    /**
     * -
     * (Optional)
     * The geographic location where the transfer config should reside.
     * Examples: US, EU, asia-northeast1. The default value is US.
     */
    public readonly location!: pulumi.Output<string | undefined>;
    /**
     * -
     * (Required)
     * The name of the reservation. This field must only contain alphanumeric characters or dash.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * -
     * (Required)
     * Minimum slots available to this reservation. A slot is a unit of computational power in BigQuery, and serves as the
     * unit of parallelism. Queries using this reservation might use more slots during runtime if ignoreIdleSlots is set to false.
     */
    public readonly slotCapacity!: pulumi.Output<number>;

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
            inputs["ignoreIdleSlots"] = state ? state.ignoreIdleSlots : undefined;
            inputs["location"] = state ? state.location : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["slotCapacity"] = state ? state.slotCapacity : undefined;
        } else {
            const args = argsOrState as ReservationArgs | undefined;
            if (!args || args.slotCapacity === undefined) {
                throw new Error("Missing required property 'slotCapacity'");
            }
            inputs["ignoreIdleSlots"] = args ? args.ignoreIdleSlots : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["slotCapacity"] = args ? args.slotCapacity : undefined;
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
     * -
     * (Optional)
     * If false, any query using this reservation will use idle slots from other reservations within
     * the same admin project. If true, a query using this reservation will execute with the slot
     * capacity specified above at most.
     */
    readonly ignoreIdleSlots?: pulumi.Input<boolean>;
    /**
     * -
     * (Optional)
     * The geographic location where the transfer config should reside.
     * Examples: US, EU, asia-northeast1. The default value is US.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * -
     * (Required)
     * The name of the reservation. This field must only contain alphanumeric characters or dash.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * -
     * (Required)
     * Minimum slots available to this reservation. A slot is a unit of computational power in BigQuery, and serves as the
     * unit of parallelism. Queries using this reservation might use more slots during runtime if ignoreIdleSlots is set to false.
     */
    readonly slotCapacity?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a Reservation resource.
 */
export interface ReservationArgs {
    /**
     * -
     * (Optional)
     * If false, any query using this reservation will use idle slots from other reservations within
     * the same admin project. If true, a query using this reservation will execute with the slot
     * capacity specified above at most.
     */
    readonly ignoreIdleSlots?: pulumi.Input<boolean>;
    /**
     * -
     * (Optional)
     * The geographic location where the transfer config should reside.
     * Examples: US, EU, asia-northeast1. The default value is US.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * -
     * (Required)
     * The name of the reservation. This field must only contain alphanumeric characters or dash.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * -
     * (Required)
     * Minimum slots available to this reservation. A slot is a unit of computational power in BigQuery, and serves as the
     * unit of parallelism. Queries using this reservation might use more slots during runtime if ignoreIdleSlots is set to false.
     */
    readonly slotCapacity: pulumi.Input<number>;
}
