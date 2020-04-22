// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Get the serial port output from a Compute Instance. For more information see
 * the official [API](https://cloud.google.com/compute/docs/instances/viewing-serial-port-output) documentation.
 * 
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/d/datasource_compute_instance_serial_port.html.markdown.
 */
export function getInstanceSerialPort(args: GetInstanceSerialPortArgs, opts?: pulumi.InvokeOptions): Promise<GetInstanceSerialPortResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("gcp:compute/getInstanceSerialPort:getInstanceSerialPort", {
        "instance": args.instance,
        "port": args.port,
        "project": args.project,
        "zone": args.zone,
    }, opts);
}

/**
 * A collection of arguments for invoking getInstanceSerialPort.
 */
export interface GetInstanceSerialPortArgs {
    /**
     * The name of the Compute Instance to read output from.
     */
    readonly instance: string;
    /**
     * The number of the serial port to read output from. Possible values are 1-4.
     */
    readonly port: number;
    /**
     * The project in which the Compute Instance exists. If it
     * is not provided, the provider project is used.
     */
    readonly project?: string;
    /**
     * The zone in which the Compute Instance exists.
     * If it is not provided, the provider zone is used.
     */
    readonly zone?: string;
}

/**
 * A collection of values returned by getInstanceSerialPort.
 */
export interface GetInstanceSerialPortResult {
    /**
     * The output of the serial port. Serial port output is available only when the VM instance is running, and logs are limited to the most recent 1 MB of output per port.
     */
    readonly contents: string;
    readonly instance: string;
    readonly port: number;
    readonly project: string;
    readonly zone: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
