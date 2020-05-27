// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * A Hl7V2Store is a datastore inside a Healthcare dataset that conforms to the FHIR (https://www.hl7.org/hl7V2/STU3/)
 * standard for Healthcare information exchange
 *
 * To get more information about Hl7V2Store, see:
 *
 * * [API documentation](https://cloud.google.com/healthcare/docs/reference/rest/v1/projects.locations.datasets.hl7V2Stores)
 * * How-to Guides
 *     * [Creating a HL7v2 Store](https://cloud.google.com/healthcare/docs/how-tos/hl7v2)
 *
 * ## Example Usage - Healthcare Hl7 V2 Store Basic
 * {{% example %}}
 *
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const topic = new gcp.pubsub.Topic("topic", {});
 * const dataset = new gcp.healthcare.Dataset("dataset", {location: "us-central1"});
 * const default = new gcp.healthcare.Hl7Store("default", {
 *     dataset: dataset.id,
 *     notification_configs: [{
 *         pubsubTopic: topic.id,
 *     }],
 *     labels: {
 *         label1: "labelvalue1",
 *     },
 * });
 * ```
 * {{% /example %}}
 * ## Example Usage - Healthcare Hl7 V2 Store Parser Config
 * {{% example %}}
 *
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const dataset = new gcp.healthcare.Dataset("dataset", {location: "us-central1"});
 * const default = new gcp.healthcare.Hl7Store("default", {
 *     dataset: dataset.id,
 *     parser_config: {
 *         allowNullHeader: false,
 *         segmentTerminator: "Jw==",
 *         schema: `{
 *   "schemas": [{
 *     "messageSchemaConfigs": {
 *       "ADT_A01": {
 *         "name": "ADT_A01",
 *         "minOccurs": 1,
 *         "maxOccurs": 1,
 *         "members": [{
 *             "segment": {
 *               "type": "MSH",
 *               "minOccurs": 1,
 *               "maxOccurs": 1
 *             }
 *           },
 *           {
 *             "segment": {
 *               "type": "EVN",
 *               "minOccurs": 1,
 *               "maxOccurs": 1
 *             }
 *           },
 *           {
 *             "segment": {
 *               "type": "PID",
 *               "minOccurs": 1,
 *               "maxOccurs": 1
 *             }
 *           },
 *           {
 *             "segment": {
 *               "type": "ZPD",
 *               "minOccurs": 1,
 *               "maxOccurs": 1
 *             }
 *           },
 *           {
 *             "segment": {
 *               "type": "OBX"
 *             }
 *           },
 *           {
 *             "group": {
 *               "name": "PROCEDURE",
 *               "members": [{
 *                   "segment": {
 *                     "type": "PR1",
 *                     "minOccurs": 1,
 *                     "maxOccurs": 1
 *                   }
 *                 },
 *                 {
 *                   "segment": {
 *                     "type": "ROL"
 *                   }
 *                 }
 *               ]
 *             }
 *           },
 *           {
 *             "segment": {
 *               "type": "PDA",
 *               "maxOccurs": 1
 *             }
 *           }
 *         ]
 *       }
 *     }
 *   }],
 *   "types": [{
 *     "type": [{
 *         "name": "ZPD",
 *         "primitive": "VARIES"
 *       }
 *
 *     ]
 *   }],
 *   "ignoreMinOccurs": true
 * }
 * `,
 *     },
 * });
 * ```
 *
 * {{% /example %}}
 */
export class Hl7Store extends pulumi.CustomResource {
    /**
     * Get an existing Hl7Store resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: Hl7StoreState, opts?: pulumi.CustomResourceOptions): Hl7Store {
        return new Hl7Store(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:healthcare/hl7Store:Hl7Store';

    /**
     * Returns true if the given object is an instance of Hl7Store.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Hl7Store {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Hl7Store.__pulumiType;
    }

    /**
     * Identifies the dataset addressed by this request. Must be in the format
     * 'projects/{project}/locations/{location}/datasets/{dataset}'
     */
    public readonly dataset!: pulumi.Output<string>;
    /**
     * User-supplied key-value pairs used to organize HL7v2 stores.
     * Label keys must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must
     * conform to the following PCRE regular expression: [\p{Ll}\p{Lo}][\p{Ll}\p{Lo}\p{N}_-]{0,62}
     * Label values are optional, must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128
     * bytes, and must conform to the following PCRE regular expression: [\p{Ll}\p{Lo}\p{N}_-]{0,63}
     * No more than 64 labels can be associated with a given store.
     * An object containing a list of "key": value pairs.
     * Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The resource name for the Hl7V2Store.
     * ** Changing this property may recreate the Hl7v2 store (removing all data) **
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * -
     * (Optional, Deprecated)
     * A nested object resource  Structure is documented below.
     */
    public readonly notificationConfig!: pulumi.Output<outputs.healthcare.Hl7StoreNotificationConfig | undefined>;
    /**
     * A list of notification configs. Each configuration uses a filter to determine whether to publish a
     * message (both Ingest & Create) on the corresponding notification destination. Only the message name
     * is sent as part of the notification. Supplied by the client.  Structure is documented below.
     */
    public readonly notificationConfigs!: pulumi.Output<outputs.healthcare.Hl7StoreNotificationConfigs[] | undefined>;
    /**
     * A nested object resource  Structure is documented below.
     */
    public readonly parserConfig!: pulumi.Output<outputs.healthcare.Hl7StoreParserConfig | undefined>;
    /**
     * The fully qualified name of this dataset
     */
    public /*out*/ readonly selfLink!: pulumi.Output<string>;

    /**
     * Create a Hl7Store resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: Hl7StoreArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: Hl7StoreArgs | Hl7StoreState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as Hl7StoreState | undefined;
            inputs["dataset"] = state ? state.dataset : undefined;
            inputs["labels"] = state ? state.labels : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["notificationConfig"] = state ? state.notificationConfig : undefined;
            inputs["notificationConfigs"] = state ? state.notificationConfigs : undefined;
            inputs["parserConfig"] = state ? state.parserConfig : undefined;
            inputs["selfLink"] = state ? state.selfLink : undefined;
        } else {
            const args = argsOrState as Hl7StoreArgs | undefined;
            if (!args || args.dataset === undefined) {
                throw new Error("Missing required property 'dataset'");
            }
            inputs["dataset"] = args ? args.dataset : undefined;
            inputs["labels"] = args ? args.labels : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["notificationConfig"] = args ? args.notificationConfig : undefined;
            inputs["notificationConfigs"] = args ? args.notificationConfigs : undefined;
            inputs["parserConfig"] = args ? args.parserConfig : undefined;
            inputs["selfLink"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Hl7Store.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Hl7Store resources.
 */
export interface Hl7StoreState {
    /**
     * Identifies the dataset addressed by this request. Must be in the format
     * 'projects/{project}/locations/{location}/datasets/{dataset}'
     */
    readonly dataset?: pulumi.Input<string>;
    /**
     * User-supplied key-value pairs used to organize HL7v2 stores.
     * Label keys must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must
     * conform to the following PCRE regular expression: [\p{Ll}\p{Lo}][\p{Ll}\p{Lo}\p{N}_-]{0,62}
     * Label values are optional, must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128
     * bytes, and must conform to the following PCRE regular expression: [\p{Ll}\p{Lo}\p{N}_-]{0,63}
     * No more than 64 labels can be associated with a given store.
     * An object containing a list of "key": value pairs.
     * Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
     */
    readonly labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The resource name for the Hl7V2Store.
     * ** Changing this property may recreate the Hl7v2 store (removing all data) **
     */
    readonly name?: pulumi.Input<string>;
    /**
     * -
     * (Optional, Deprecated)
     * A nested object resource  Structure is documented below.
     * @deprecated This field has been replaced by notificationConfigs
     */
    readonly notificationConfig?: pulumi.Input<inputs.healthcare.Hl7StoreNotificationConfig>;
    /**
     * A list of notification configs. Each configuration uses a filter to determine whether to publish a
     * message (both Ingest & Create) on the corresponding notification destination. Only the message name
     * is sent as part of the notification. Supplied by the client.  Structure is documented below.
     */
    readonly notificationConfigs?: pulumi.Input<pulumi.Input<inputs.healthcare.Hl7StoreNotificationConfigs>[]>;
    /**
     * A nested object resource  Structure is documented below.
     */
    readonly parserConfig?: pulumi.Input<inputs.healthcare.Hl7StoreParserConfig>;
    /**
     * The fully qualified name of this dataset
     */
    readonly selfLink?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Hl7Store resource.
 */
export interface Hl7StoreArgs {
    /**
     * Identifies the dataset addressed by this request. Must be in the format
     * 'projects/{project}/locations/{location}/datasets/{dataset}'
     */
    readonly dataset: pulumi.Input<string>;
    /**
     * User-supplied key-value pairs used to organize HL7v2 stores.
     * Label keys must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must
     * conform to the following PCRE regular expression: [\p{Ll}\p{Lo}][\p{Ll}\p{Lo}\p{N}_-]{0,62}
     * Label values are optional, must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128
     * bytes, and must conform to the following PCRE regular expression: [\p{Ll}\p{Lo}\p{N}_-]{0,63}
     * No more than 64 labels can be associated with a given store.
     * An object containing a list of "key": value pairs.
     * Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
     */
    readonly labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The resource name for the Hl7V2Store.
     * ** Changing this property may recreate the Hl7v2 store (removing all data) **
     */
    readonly name?: pulumi.Input<string>;
    /**
     * -
     * (Optional, Deprecated)
     * A nested object resource  Structure is documented below.
     * @deprecated This field has been replaced by notificationConfigs
     */
    readonly notificationConfig?: pulumi.Input<inputs.healthcare.Hl7StoreNotificationConfig>;
    /**
     * A list of notification configs. Each configuration uses a filter to determine whether to publish a
     * message (both Ingest & Create) on the corresponding notification destination. Only the message name
     * is sent as part of the notification. Supplied by the client.  Structure is documented below.
     */
    readonly notificationConfigs?: pulumi.Input<pulumi.Input<inputs.healthcare.Hl7StoreNotificationConfigs>[]>;
    /**
     * A nested object resource  Structure is documented below.
     */
    readonly parserConfig?: pulumi.Input<inputs.healthcare.Hl7StoreParserConfig>;
}
