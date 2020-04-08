// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package kms

import (
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

//  Creates a device registry in Google's Cloud IoT Core platform. For more information see
// [the official documentation](https://cloud.google.com/iot/docs/) and
// [API](https://cloud.google.com/iot/docs/reference/cloudiot/rest/v1/projects.locations.registries).
type Registry struct {
	pulumi.CustomResourceState

	// List of public key certificates to authenticate devices. Structure is documented below.
	Credentials RegistryCredentialArrayOutput `pulumi:"credentials"`
	// List of configurations for event notification, such as
	// PubSub topics to publish device events to. Structure is documented below.
	EventNotificationConfigs RegistryEventNotificationConfigItemArrayOutput `pulumi:"eventNotificationConfigs"`
	// Activate or deactivate HTTP. Structure is documented below.
	HttpConfig RegistryHttpConfigOutput `pulumi:"httpConfig"`
	LogLevel   pulumi.StringPtrOutput   `pulumi:"logLevel"`
	// Activate or deactivate MQTT. Structure is documented below.
	MqttConfig RegistryMqttConfigOutput `pulumi:"mqttConfig"`
	// A unique name for the resource, required by device registry.
	// Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The project in which the resource belongs. If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The Region in which the created address should reside. If it is not provided, the provider region is used.
	Region pulumi.StringOutput `pulumi:"region"`
	// A PubSub topic to publish device state updates. Structure is documented below.
	StateNotificationConfig RegistryStateNotificationConfigPtrOutput `pulumi:"stateNotificationConfig"`
}

// NewRegistry registers a new resource with the given unique name, arguments, and options.
func NewRegistry(ctx *pulumi.Context,
	name string, args *RegistryArgs, opts ...pulumi.ResourceOption) (*Registry, error) {
	if args == nil {
		args = &RegistryArgs{}
	}
	var resource Registry
	err := ctx.RegisterResource("gcp:kms/registry:Registry", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRegistry gets an existing Registry resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRegistry(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RegistryState, opts ...pulumi.ResourceOption) (*Registry, error) {
	var resource Registry
	err := ctx.ReadResource("gcp:kms/registry:Registry", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Registry resources.
type registryState struct {
	// List of public key certificates to authenticate devices. Structure is documented below.
	Credentials []RegistryCredential `pulumi:"credentials"`
	// List of configurations for event notification, such as
	// PubSub topics to publish device events to. Structure is documented below.
	EventNotificationConfigs []RegistryEventNotificationConfigItem `pulumi:"eventNotificationConfigs"`
	// Activate or deactivate HTTP. Structure is documented below.
	HttpConfig *RegistryHttpConfig `pulumi:"httpConfig"`
	LogLevel   *string             `pulumi:"logLevel"`
	// Activate or deactivate MQTT. Structure is documented below.
	MqttConfig *RegistryMqttConfig `pulumi:"mqttConfig"`
	// A unique name for the resource, required by device registry.
	// Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The project in which the resource belongs. If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The Region in which the created address should reside. If it is not provided, the provider region is used.
	Region *string `pulumi:"region"`
	// A PubSub topic to publish device state updates. Structure is documented below.
	StateNotificationConfig *RegistryStateNotificationConfig `pulumi:"stateNotificationConfig"`
}

type RegistryState struct {
	// List of public key certificates to authenticate devices. Structure is documented below.
	Credentials RegistryCredentialArrayInput
	// List of configurations for event notification, such as
	// PubSub topics to publish device events to. Structure is documented below.
	EventNotificationConfigs RegistryEventNotificationConfigItemArrayInput
	// Activate or deactivate HTTP. Structure is documented below.
	HttpConfig RegistryHttpConfigPtrInput
	LogLevel   pulumi.StringPtrInput
	// Activate or deactivate MQTT. Structure is documented below.
	MqttConfig RegistryMqttConfigPtrInput
	// A unique name for the resource, required by device registry.
	// Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The project in which the resource belongs. If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The Region in which the created address should reside. If it is not provided, the provider region is used.
	Region pulumi.StringPtrInput
	// A PubSub topic to publish device state updates. Structure is documented below.
	StateNotificationConfig RegistryStateNotificationConfigPtrInput
}

func (RegistryState) ElementType() reflect.Type {
	return reflect.TypeOf((*registryState)(nil)).Elem()
}

type registryArgs struct {
	// List of public key certificates to authenticate devices. Structure is documented below.
	Credentials []RegistryCredentialArgs `pulumi:"credentials"`
	// List of configurations for event notification, such as
	// PubSub topics to publish device events to. Structure is documented below.
	EventNotificationConfigs []RegistryEventNotificationConfigItemArgs `pulumi:"eventNotificationConfigs"`
	// Activate or deactivate HTTP. Structure is documented below.
	HttpConfig *RegistryHttpConfigArgs `pulumi:"httpConfig"`
	LogLevel   *string                 `pulumi:"logLevel"`
	// Activate or deactivate MQTT. Structure is documented below.
	MqttConfig *RegistryMqttConfigArgs `pulumi:"mqttConfig"`
	// A unique name for the resource, required by device registry.
	// Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The project in which the resource belongs. If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The Region in which the created address should reside. If it is not provided, the provider region is used.
	Region *string `pulumi:"region"`
	// A PubSub topic to publish device state updates. Structure is documented below.
	StateNotificationConfig *RegistryStateNotificationConfigArgs `pulumi:"stateNotificationConfig"`
}

// The set of arguments for constructing a Registry resource.
type RegistryArgs struct {
	// List of public key certificates to authenticate devices. Structure is documented below.
	Credentials RegistryCredentialArgsArrayInput
	// List of configurations for event notification, such as
	// PubSub topics to publish device events to. Structure is documented below.
	EventNotificationConfigs RegistryEventNotificationConfigItemArgsArrayInput
	// Activate or deactivate HTTP. Structure is documented below.
	HttpConfig RegistryHttpConfigArgsPtrInput
	LogLevel   pulumi.StringPtrInput
	// Activate or deactivate MQTT. Structure is documented below.
	MqttConfig RegistryMqttConfigArgsPtrInput
	// A unique name for the resource, required by device registry.
	// Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The project in which the resource belongs. If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The Region in which the created address should reside. If it is not provided, the provider region is used.
	Region pulumi.StringPtrInput
	// A PubSub topic to publish device state updates. Structure is documented below.
	StateNotificationConfig RegistryStateNotificationConfigArgsPtrInput
}

func (RegistryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*registryArgs)(nil)).Elem()
}
