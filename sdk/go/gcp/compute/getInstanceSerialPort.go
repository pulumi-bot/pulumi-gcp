// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Get the serial port output from a Compute Instance. For more information see
// the official [API](https://cloud.google.com/compute/docs/instances/viewing-serial-port-output) documentation.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		serial, err := compute.LookupInstanceSerialPort(ctx, &compute.LookupInstanceSerialPortArgs{
// 			Instance: "my-instance",
// 			Zone:     "us-central1-a",
// 			Port:     1,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("serialOut", serial.Contents)
// 		return nil
// 	})
// }
// ```
//
// Using the serial port output to generate a windows password, derived from the [official guide](https://cloud.google.com/compute/docs/instances/windows/automate-pw-generation):
func GetInstanceSerialPort(ctx *pulumi.Context, args *GetInstanceSerialPortArgs, opts ...pulumi.InvokeOption) (*GetInstanceSerialPortResult, error) {
	var rv GetInstanceSerialPortResult
	err := ctx.Invoke("gcp:compute/getInstanceSerialPort:getInstanceSerialPort", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getInstanceSerialPort.
type GetInstanceSerialPortArgs struct {
	// The name of the Compute Instance to read output from.
	Instance string `pulumi:"instance"`
	// The number of the serial port to read output from. Possible values are 1-4.
	Port int `pulumi:"port"`
	// The project in which the Compute Instance exists. If it
	// is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The zone in which the Compute Instance exists.
	// If it is not provided, the provider zone is used.
	Zone *string `pulumi:"zone"`
}

// A collection of values returned by getInstanceSerialPort.
type GetInstanceSerialPortResult struct {
	// The output of the serial port. Serial port output is available only when the VM instance is running, and logs are limited to the most recent 1 MB of output per port.
	Contents string `pulumi:"contents"`
	// The provider-assigned unique ID for this managed resource.
	Id       string `pulumi:"id"`
	Instance string `pulumi:"instance"`
	Port     int    `pulumi:"port"`
	Project  string `pulumi:"project"`
	Zone     string `pulumi:"zone"`
}
