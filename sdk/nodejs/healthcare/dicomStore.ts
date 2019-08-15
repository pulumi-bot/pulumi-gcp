// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputApi from "../types/input";
import * as outputApi from "../types/output";
import * as utilities from "../utilities";

/**
 * A DicomStore is a datastore inside a Healthcare dataset that conforms to the DICOM
 * (https://www.dicomstandard.org/about/) standard for Healthcare information exchange
 * 
 * To get more information about DicomStore, see:
 * 
 * * [API documentation](https://cloud.google.com/healthcare/docs/reference/rest/v1beta1/projects.locations.datasets.dicomStores)
 * * How-to Guides
 *     * [Creating a DICOM store](https://cloud.google.com/healthcare/docs/how-tos/dicom)
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/healthcare_dicom_store.html.markdown.
 */
export class DicomStore extends pulumi.CustomResource {
    /**
     * Get an existing DicomStore resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DicomStoreState, opts?: pulumi.CustomResourceOptions): DicomStore {
        return new DicomStore(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:healthcare/dicomStore:DicomStore';

    /**
     * Returns true if the given object is an instance of DicomStore.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DicomStore {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DicomStore.__pulumiType;
    }

    public readonly dataset!: pulumi.Output<string>;
    public readonly labels!: pulumi.Output<{[key: string]: string} | undefined>;
    public readonly name!: pulumi.Output<string>;
    public readonly notificationConfig!: pulumi.Output<outputApi.healthcare.DicomStoreNotificationConfig | undefined>;
    public /*out*/ readonly selfLink!: pulumi.Output<string>;

    /**
     * Create a DicomStore resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DicomStoreArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DicomStoreArgs | DicomStoreState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as DicomStoreState | undefined;
            inputs["dataset"] = state ? state.dataset : undefined;
            inputs["labels"] = state ? state.labels : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["notificationConfig"] = state ? state.notificationConfig : undefined;
            inputs["selfLink"] = state ? state.selfLink : undefined;
        } else {
            const args = argsOrState as DicomStoreArgs | undefined;
            if (!args || args.dataset === undefined) {
                throw new Error("Missing required property 'dataset'");
            }
            inputs["dataset"] = args ? args.dataset : undefined;
            inputs["labels"] = args ? args.labels : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["notificationConfig"] = args ? args.notificationConfig : undefined;
            inputs["selfLink"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(DicomStore.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DicomStore resources.
 */
export interface DicomStoreState {
    readonly dataset?: pulumi.Input<string>;
    readonly labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    readonly name?: pulumi.Input<string>;
    readonly notificationConfig?: pulumi.Input<inputApi.healthcare.DicomStoreNotificationConfig>;
    readonly selfLink?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a DicomStore resource.
 */
export interface DicomStoreArgs {
    readonly dataset: pulumi.Input<string>;
    readonly labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    readonly name?: pulumi.Input<string>;
    readonly notificationConfig?: pulumi.Input<inputApi.healthcare.DicomStoreNotificationConfig>;
}
