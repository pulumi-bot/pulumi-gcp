// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package registry

import (
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
	"https:/github.com/pulumi/pulumi-gcp/kms/RegistryCredential"
	"https:/github.com/pulumi/pulumi-gcp/kms/RegistryCredentialPublicKeyCertificate"
	"https:/github.com/pulumi/pulumi-gcp/kms/RegistryEventNotificationConfigItem"
	"https:/github.com/pulumi/pulumi-gcp/kms/RegistryHttpConfig"
	"https:/github.com/pulumi/pulumi-gcp/kms/RegistryMqttConfig"
	"https:/github.com/pulumi/pulumi-gcp/kms/RegistryStateNotificationConfig"
)

//  Creates a device registry in Google's Cloud IoT Core platform. For more information see
// [the official documentation](https://cloud.google.com/iot/docs/) and
// [API](https://cloud.google.com/iot/docs/reference/cloudiot/rest/v1/projects.locations.registries).
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/cloudiot_registry.html.markdown.
type Registry struct {
	pulumi.CustomResourceState

	// List of public key certificates to authenticate devices. Structure is documented below. 
	Credentials kmsRegistryCredential.RegistryCredentialArrayOutput `pulumi:"credentials"`
	// List of configurations for event notification, such as
	// PubSub topics to publish device events to. Structure is documented below.
	EventNotificationConfigs kmsRegistryEventNotificationConfigItem.RegistryEventNotificationConfigItemArrayOutput `pulumi:"eventNotificationConfigs"`
	// Activate or deactivate HTTP. Structure is documented below.
	HttpConfig kmsRegistryHttpConfig.RegistryHttpConfigOutput `pulumi:"httpConfig"`
	LogLevel pulumi.StringPtrOutput `pulumi:"logLevel"`
	// Activate or deactivate MQTT. Structure is documented below.
	MqttConfig kmsRegistryMqttConfig.RegistryMqttConfigOutput `pulumi:"mqttConfig"`
	// A unique name for the resource, required by device registry.
	// Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The project in which the resource belongs. If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The Region in which the created address should reside. If it is not provided, the provider region is used.
	Region pulumi.StringOutput `pulumi:"region"`
	// A PubSub topic to publish device state updates. Structure is documented below.
	StateNotificationConfig kmsRegistryStateNotificationConfig.RegistryStateNotificationConfigPtrOutput `pulumi:"stateNotificationConfig"`
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
	Credentials []kmsRegistryCredential.RegistryCredential `pulumi:"credentials"`
	// List of configurations for event notification, such as
	// PubSub topics to publish device events to. Structure is documented below.
	EventNotificationConfigs []kmsRegistryEventNotificationConfigItem.RegistryEventNotificationConfigItem `pulumi:"eventNotificationConfigs"`
	// Activate or deactivate HTTP. Structure is documented below.
	HttpConfig *kmsRegistryHttpConfig.RegistryHttpConfig `pulumi:"httpConfig"`
	LogLevel *string `pulumi:"logLevel"`
	// Activate or deactivate MQTT. Structure is documented below.
	MqttConfig *kmsRegistryMqttConfig.RegistryMqttConfig `pulumi:"mqttConfig"`
	// A unique name for the resource, required by device registry.
	// Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The project in which the resource belongs. If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The Region in which the created address should reside. If it is not provided, the provider region is used.
	Region *string `pulumi:"region"`
	// A PubSub topic to publish device state updates. Structure is documented below.
	StateNotificationConfig *kmsRegistryStateNotificationConfig.RegistryStateNotificationConfig `pulumi:"stateNotificationConfig"`
}

type RegistryState struct {
	// List of public key certificates to authenticate devices. Structure is documented below. 
	Credentials kmsRegistryCredential.RegistryCredentialArrayInput
	// List of configurations for event notification, such as
	// PubSub topics to publish device events to. Structure is documented below.
	EventNotificationConfigs kmsRegistryEventNotificationConfigItem.RegistryEventNotificationConfigItemArrayInput
	// Activate or deactivate HTTP. Structure is documented below.
	HttpConfig kmsRegistryHttpConfig.RegistryHttpConfigPtrInput
	LogLevel pulumi.StringPtrInput
	// Activate or deactivate MQTT. Structure is documented below.
	MqttConfig kmsRegistryMqttConfig.RegistryMqttConfigPtrInput
	// A unique name for the resource, required by device registry.
	// Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The project in which the resource belongs. If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The Region in which the created address should reside. If it is not provided, the provider region is used.
	Region pulumi.StringPtrInput
	// A PubSub topic to publish device state updates. Structure is documented below.
	StateNotificationConfig kmsRegistryStateNotificationConfig.RegistryStateNotificationConfigPtrInput
}

func (RegistryState) ElementType() reflect.Type {
	return reflect.TypeOf((*registryState)(nil)).Elem()
}

type registryArgs struct {
	// List of public key certificates to authenticate devices. Structure is documented below. 
	Credentials []kmsRegistryCredential.RegistryCredential `pulumi:"credentials"`
	// List of configurations for event notification, such as
	// PubSub topics to publish device events to. Structure is documented below.
	EventNotificationConfigs []kmsRegistryEventNotificationConfigItem.RegistryEventNotificationConfigItem `pulumi:"eventNotificationConfigs"`
	// Activate or deactivate HTTP. Structure is documented below.
	HttpConfig *kmsRegistryHttpConfig.RegistryHttpConfig `pulumi:"httpConfig"`
	LogLevel *string `pulumi:"logLevel"`
	// Activate or deactivate MQTT. Structure is documented below.
	MqttConfig *kmsRegistryMqttConfig.RegistryMqttConfig `pulumi:"mqttConfig"`
	// A unique name for the resource, required by device registry.
	// Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The project in which the resource belongs. If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The Region in which the created address should reside. If it is not provided, the provider region is used.
	Region *string `pulumi:"region"`
	// A PubSub topic to publish device state updates. Structure is documented below.
	StateNotificationConfig *kmsRegistryStateNotificationConfig.RegistryStateNotificationConfig `pulumi:"stateNotificationConfig"`
}

// The set of arguments for constructing a Registry resource.
type RegistryArgs struct {
	// List of public key certificates to authenticate devices. Structure is documented below. 
	Credentials kmsRegistryCredential.RegistryCredentialArrayInput
	// List of configurations for event notification, such as
	// PubSub topics to publish device events to. Structure is documented below.
	EventNotificationConfigs kmsRegistryEventNotificationConfigItem.RegistryEventNotificationConfigItemArrayInput
	// Activate or deactivate HTTP. Structure is documented below.
	HttpConfig kmsRegistryHttpConfig.RegistryHttpConfigPtrInput
	LogLevel pulumi.StringPtrInput
	// Activate or deactivate MQTT. Structure is documented below.
	MqttConfig kmsRegistryMqttConfig.RegistryMqttConfigPtrInput
	// A unique name for the resource, required by device registry.
	// Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The project in which the resource belongs. If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The Region in which the created address should reside. If it is not provided, the provider region is used.
	Region pulumi.StringPtrInput
	// A PubSub topic to publish device state updates. Structure is documented below.
	StateNotificationConfig kmsRegistryStateNotificationConfig.RegistryStateNotificationConfigPtrInput
}

func (RegistryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*registryArgs)(nil)).Elem()
}

