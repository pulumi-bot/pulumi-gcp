// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package iot

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// A Google Cloud IoT Core device.
//
// To get more information about Device, see:
//
// * [API documentation](https://cloud.google.com/iot/docs/reference/cloudiot/rest/)
// * How-to Guides
//     * [Official Documentation](https://cloud.google.com/iot/docs/)
//
// ## Example Usage
type Device struct {
	pulumi.CustomResourceState

	// If a device is blocked, connections or requests from this device will fail.
	Blocked pulumi.BoolPtrOutput `pulumi:"blocked"`
	// The most recent device configuration, which is eventually sent from Cloud IoT Core to the device.
	Configs DeviceConfigArrayOutput `pulumi:"configs"`
	// The credentials used to authenticate this device.
	// Structure is documented below.
	Credentials DeviceCredentialArrayOutput `pulumi:"credentials"`
	// Gateway-related configuration and state.
	// Structure is documented below.
	GatewayConfig DeviceGatewayConfigPtrOutput `pulumi:"gatewayConfig"`
	// The last time a cloud-to-device config version acknowledgment was received from the device.
	LastConfigAckTime pulumi.StringOutput `pulumi:"lastConfigAckTime"`
	// The last time a cloud-to-device config version was sent to the device.
	LastConfigSendTime pulumi.StringOutput `pulumi:"lastConfigSendTime"`
	// The error message of the most recent error, such as a failure to publish to Cloud Pub/Sub.
	LastErrorStatuses DeviceLastErrorStatusArrayOutput `pulumi:"lastErrorStatuses"`
	// The time the most recent error occurred, such as a failure to publish to Cloud Pub/Sub.
	LastErrorTime pulumi.StringOutput `pulumi:"lastErrorTime"`
	// The last time a telemetry event was received.
	LastEventTime pulumi.StringOutput `pulumi:"lastEventTime"`
	// The last time an MQTT PINGREQ was received.
	LastHeartbeatTime pulumi.StringOutput `pulumi:"lastHeartbeatTime"`
	// The last time a state event was received.
	LastStateTime pulumi.StringOutput `pulumi:"lastStateTime"`
	// The logging verbosity for device activity.
	// Possible values are `NONE`, `ERROR`, `INFO`, and `DEBUG`.
	LogLevel pulumi.StringPtrOutput `pulumi:"logLevel"`
	// The metadata key-value pairs assigned to the device.
	Metadata pulumi.StringMapOutput `pulumi:"metadata"`
	// A unique name for the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// A server-defined unique numeric ID for the device. This is a more compact way to identify devices, and it is globally
	// unique.
	NumId pulumi.StringOutput `pulumi:"numId"`
	// The name of the device registry where this device should be created.
	Registry pulumi.StringOutput `pulumi:"registry"`
	// The state most recently received from the device.
	States DeviceStateTypeArrayOutput `pulumi:"states"`
}

// NewDevice registers a new resource with the given unique name, arguments, and options.
func NewDevice(ctx *pulumi.Context,
	name string, args *DeviceArgs, opts ...pulumi.ResourceOption) (*Device, error) {
	if args == nil || args.Registry == nil {
		return nil, errors.New("missing required argument 'Registry'")
	}
	if args == nil {
		args = &DeviceArgs{}
	}
	var resource Device
	err := ctx.RegisterResource("gcp:iot/device:Device", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDevice gets an existing Device resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDevice(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DeviceState, opts ...pulumi.ResourceOption) (*Device, error) {
	var resource Device
	err := ctx.ReadResource("gcp:iot/device:Device", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Device resources.
type deviceState struct {
	// If a device is blocked, connections or requests from this device will fail.
	Blocked *bool `pulumi:"blocked"`
	// The most recent device configuration, which is eventually sent from Cloud IoT Core to the device.
	Configs []DeviceConfig `pulumi:"configs"`
	// The credentials used to authenticate this device.
	// Structure is documented below.
	Credentials []DeviceCredential `pulumi:"credentials"`
	// Gateway-related configuration and state.
	// Structure is documented below.
	GatewayConfig *DeviceGatewayConfig `pulumi:"gatewayConfig"`
	// The last time a cloud-to-device config version acknowledgment was received from the device.
	LastConfigAckTime *string `pulumi:"lastConfigAckTime"`
	// The last time a cloud-to-device config version was sent to the device.
	LastConfigSendTime *string `pulumi:"lastConfigSendTime"`
	// The error message of the most recent error, such as a failure to publish to Cloud Pub/Sub.
	LastErrorStatuses []DeviceLastErrorStatus `pulumi:"lastErrorStatuses"`
	// The time the most recent error occurred, such as a failure to publish to Cloud Pub/Sub.
	LastErrorTime *string `pulumi:"lastErrorTime"`
	// The last time a telemetry event was received.
	LastEventTime *string `pulumi:"lastEventTime"`
	// The last time an MQTT PINGREQ was received.
	LastHeartbeatTime *string `pulumi:"lastHeartbeatTime"`
	// The last time a state event was received.
	LastStateTime *string `pulumi:"lastStateTime"`
	// The logging verbosity for device activity.
	// Possible values are `NONE`, `ERROR`, `INFO`, and `DEBUG`.
	LogLevel *string `pulumi:"logLevel"`
	// The metadata key-value pairs assigned to the device.
	Metadata map[string]string `pulumi:"metadata"`
	// A unique name for the resource.
	Name *string `pulumi:"name"`
	// A server-defined unique numeric ID for the device. This is a more compact way to identify devices, and it is globally
	// unique.
	NumId *string `pulumi:"numId"`
	// The name of the device registry where this device should be created.
	Registry *string `pulumi:"registry"`
	// The state most recently received from the device.
	States []DeviceStateType `pulumi:"states"`
}

type DeviceState struct {
	// If a device is blocked, connections or requests from this device will fail.
	Blocked pulumi.BoolPtrInput
	// The most recent device configuration, which is eventually sent from Cloud IoT Core to the device.
	Configs DeviceConfigArrayInput
	// The credentials used to authenticate this device.
	// Structure is documented below.
	Credentials DeviceCredentialArrayInput
	// Gateway-related configuration and state.
	// Structure is documented below.
	GatewayConfig DeviceGatewayConfigPtrInput
	// The last time a cloud-to-device config version acknowledgment was received from the device.
	LastConfigAckTime pulumi.StringPtrInput
	// The last time a cloud-to-device config version was sent to the device.
	LastConfigSendTime pulumi.StringPtrInput
	// The error message of the most recent error, such as a failure to publish to Cloud Pub/Sub.
	LastErrorStatuses DeviceLastErrorStatusArrayInput
	// The time the most recent error occurred, such as a failure to publish to Cloud Pub/Sub.
	LastErrorTime pulumi.StringPtrInput
	// The last time a telemetry event was received.
	LastEventTime pulumi.StringPtrInput
	// The last time an MQTT PINGREQ was received.
	LastHeartbeatTime pulumi.StringPtrInput
	// The last time a state event was received.
	LastStateTime pulumi.StringPtrInput
	// The logging verbosity for device activity.
	// Possible values are `NONE`, `ERROR`, `INFO`, and `DEBUG`.
	LogLevel pulumi.StringPtrInput
	// The metadata key-value pairs assigned to the device.
	Metadata pulumi.StringMapInput
	// A unique name for the resource.
	Name pulumi.StringPtrInput
	// A server-defined unique numeric ID for the device. This is a more compact way to identify devices, and it is globally
	// unique.
	NumId pulumi.StringPtrInput
	// The name of the device registry where this device should be created.
	Registry pulumi.StringPtrInput
	// The state most recently received from the device.
	States DeviceStateTypeArrayInput
}

func (DeviceState) ElementType() reflect.Type {
	return reflect.TypeOf((*deviceState)(nil)).Elem()
}

type deviceArgs struct {
	// If a device is blocked, connections or requests from this device will fail.
	Blocked *bool `pulumi:"blocked"`
	// The credentials used to authenticate this device.
	// Structure is documented below.
	Credentials []DeviceCredential `pulumi:"credentials"`
	// Gateway-related configuration and state.
	// Structure is documented below.
	GatewayConfig *DeviceGatewayConfig `pulumi:"gatewayConfig"`
	// The logging verbosity for device activity.
	// Possible values are `NONE`, `ERROR`, `INFO`, and `DEBUG`.
	LogLevel *string `pulumi:"logLevel"`
	// The metadata key-value pairs assigned to the device.
	Metadata map[string]string `pulumi:"metadata"`
	// A unique name for the resource.
	Name *string `pulumi:"name"`
	// The name of the device registry where this device should be created.
	Registry string `pulumi:"registry"`
}

// The set of arguments for constructing a Device resource.
type DeviceArgs struct {
	// If a device is blocked, connections or requests from this device will fail.
	Blocked pulumi.BoolPtrInput
	// The credentials used to authenticate this device.
	// Structure is documented below.
	Credentials DeviceCredentialArrayInput
	// Gateway-related configuration and state.
	// Structure is documented below.
	GatewayConfig DeviceGatewayConfigPtrInput
	// The logging verbosity for device activity.
	// Possible values are `NONE`, `ERROR`, `INFO`, and `DEBUG`.
	LogLevel pulumi.StringPtrInput
	// The metadata key-value pairs assigned to the device.
	Metadata pulumi.StringMapInput
	// A unique name for the resource.
	Name pulumi.StringPtrInput
	// The name of the device registry where this device should be created.
	Registry pulumi.StringInput
}

func (DeviceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*deviceArgs)(nil)).Elem()
}

type DeviceInput interface {
	pulumi.Input

	ToDeviceOutput() DeviceOutput
	ToDeviceOutputWithContext(ctx context.Context) DeviceOutput
}

func (Device) ElementType() reflect.Type {
	return reflect.TypeOf((*Device)(nil)).Elem()
}

func (i Device) ToDeviceOutput() DeviceOutput {
	return i.ToDeviceOutputWithContext(context.Background())
}

func (i Device) ToDeviceOutputWithContext(ctx context.Context) DeviceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeviceOutput)
}

type DeviceOutput struct {
	*pulumi.OutputState
}

func (DeviceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DeviceOutput)(nil)).Elem()
}

func (o DeviceOutput) ToDeviceOutput() DeviceOutput {
	return o
}

func (o DeviceOutput) ToDeviceOutputWithContext(ctx context.Context) DeviceOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(DeviceOutput{})
}
