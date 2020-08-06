// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Iot
{
    /// <summary>
    /// A Google Cloud IoT Core device.
    /// 
    /// To get more information about Device, see:
    /// 
    /// * [API documentation](https://cloud.google.com/iot/docs/reference/cloudiot/rest/)
    /// * How-to Guides
    ///     * [Official Documentation](https://cloud.google.com/iot/docs/)
    /// 
    /// ## Example Usage
    /// ### Cloudiot Device Basic
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var registry = new Gcp.Iot.Registry("registry", new Gcp.Iot.RegistryArgs
    ///         {
    ///         });
    ///         var test_device = new Gcp.Iot.Device("test-device", new Gcp.Iot.DeviceArgs
    ///         {
    ///             Registry = registry.Id,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ### Cloudiot Device Full
    /// 
    /// ```csharp
    /// using System.IO;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var registry = new Gcp.Iot.Registry("registry", new Gcp.Iot.RegistryArgs
    ///         {
    ///         });
    ///         var test_device = new Gcp.Iot.Device("test-device", new Gcp.Iot.DeviceArgs
    ///         {
    ///             Registry = registry.Id,
    ///             Credentials = 
    ///             {
    ///                 new Gcp.Iot.Inputs.DeviceCredentialArgs
    ///                 {
    ///                     PublicKey = new Gcp.Iot.Inputs.DeviceCredentialPublicKeyArgs
    ///                     {
    ///                         Format = "RSA_PEM",
    ///                         Key = File.ReadAllText("test-fixtures/rsa_public.pem"),
    ///                     },
    ///                 },
    ///             },
    ///             Blocked = false,
    ///             LogLevel = "INFO",
    ///             Metadata = 
    ///             {
    ///                 { "test_key_1", "test_value_1" },
    ///             },
    ///             GatewayConfig = new Gcp.Iot.Inputs.DeviceGatewayConfigArgs
    ///             {
    ///                 GatewayType = "NON_GATEWAY",
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    public partial class Device : Pulumi.CustomResource
    {
        /// <summary>
        /// If a device is blocked, connections or requests from this device will fail.
        /// </summary>
        [Output("blocked")]
        public Output<bool?> Blocked { get; private set; } = null!;

        /// <summary>
        /// The most recent device configuration, which is eventually sent from Cloud IoT Core to the device.
        /// </summary>
        [Output("config")]
        public Output<Outputs.DeviceConfig> Config { get; private set; } = null!;

        /// <summary>
        /// The credentials used to authenticate this device.  Structure is documented below.
        /// </summary>
        [Output("credentials")]
        public Output<ImmutableArray<Outputs.DeviceCredential>> Credentials { get; private set; } = null!;

        /// <summary>
        /// Gateway-related configuration and state.  Structure is documented below.
        /// </summary>
        [Output("gatewayConfig")]
        public Output<Outputs.DeviceGatewayConfig?> GatewayConfig { get; private set; } = null!;

        /// <summary>
        /// The last time a cloud-to-device config version acknowledgment was received from the device.
        /// </summary>
        [Output("lastConfigAckTime")]
        public Output<string> LastConfigAckTime { get; private set; } = null!;

        /// <summary>
        /// The last time a cloud-to-device config version was sent to the device.
        /// </summary>
        [Output("lastConfigSendTime")]
        public Output<string> LastConfigSendTime { get; private set; } = null!;

        /// <summary>
        /// The error message of the most recent error, such as a failure to publish to Cloud Pub/Sub.
        /// </summary>
        [Output("lastErrorStatus")]
        public Output<Outputs.DeviceLastErrorStatus> LastErrorStatus { get; private set; } = null!;

        /// <summary>
        /// The time the most recent error occurred, such as a failure to publish to Cloud Pub/Sub.
        /// </summary>
        [Output("lastErrorTime")]
        public Output<string> LastErrorTime { get; private set; } = null!;

        /// <summary>
        /// The last time a telemetry event was received.
        /// </summary>
        [Output("lastEventTime")]
        public Output<string> LastEventTime { get; private set; } = null!;

        /// <summary>
        /// The last time an MQTT PINGREQ was received.
        /// </summary>
        [Output("lastHeartbeatTime")]
        public Output<string> LastHeartbeatTime { get; private set; } = null!;

        /// <summary>
        /// The last time a state event was received.
        /// </summary>
        [Output("lastStateTime")]
        public Output<string> LastStateTime { get; private set; } = null!;

        /// <summary>
        /// The logging verbosity for device activity.
        /// </summary>
        [Output("logLevel")]
        public Output<string?> LogLevel { get; private set; } = null!;

        /// <summary>
        /// The metadata key-value pairs assigned to the device.
        /// </summary>
        [Output("metadata")]
        public Output<ImmutableDictionary<string, string>?> Metadata { get; private set; } = null!;

        /// <summary>
        /// A unique name for the resource.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// A server-defined unique numeric ID for the device. This is a more compact way to identify devices, and it is globally
        /// unique.
        /// </summary>
        [Output("numId")]
        public Output<string> NumId { get; private set; } = null!;

        /// <summary>
        /// The name of the device registry where this device should be created.
        /// </summary>
        [Output("registry")]
        public Output<string> Registry { get; private set; } = null!;

        /// <summary>
        /// The state most recently received from the device.
        /// </summary>
        [Output("state")]
        public Output<Outputs.DeviceState> State { get; private set; } = null!;


        /// <summary>
        /// Create a Device resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Device(string name, DeviceArgs args, CustomResourceOptions? options = null)
            : base("gcp:iot/device:Device", name, args ?? new DeviceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Device(string name, Input<string> id, DeviceState? state = null, CustomResourceOptions? options = null)
            : base("gcp:iot/device:Device", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Device resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Device Get(string name, Input<string> id, DeviceState? state = null, CustomResourceOptions? options = null)
        {
            return new Device(name, id, state, options);
        }
    }

    public sealed class DeviceArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// If a device is blocked, connections or requests from this device will fail.
        /// </summary>
        [Input("blocked")]
        public Input<bool>? Blocked { get; set; }

        [Input("credentials")]
        private InputList<Inputs.DeviceCredentialArgs>? _credentials;

        /// <summary>
        /// The credentials used to authenticate this device.  Structure is documented below.
        /// </summary>
        public InputList<Inputs.DeviceCredentialArgs> Credentials
        {
            get => _credentials ?? (_credentials = new InputList<Inputs.DeviceCredentialArgs>());
            set => _credentials = value;
        }

        /// <summary>
        /// Gateway-related configuration and state.  Structure is documented below.
        /// </summary>
        [Input("gatewayConfig")]
        public Input<Inputs.DeviceGatewayConfigArgs>? GatewayConfig { get; set; }

        /// <summary>
        /// The logging verbosity for device activity.
        /// </summary>
        [Input("logLevel")]
        public Input<string>? LogLevel { get; set; }

        [Input("metadata")]
        private InputMap<string>? _metadata;

        /// <summary>
        /// The metadata key-value pairs assigned to the device.
        /// </summary>
        public InputMap<string> Metadata
        {
            get => _metadata ?? (_metadata = new InputMap<string>());
            set => _metadata = value;
        }

        /// <summary>
        /// A unique name for the resource.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of the device registry where this device should be created.
        /// </summary>
        [Input("registry", required: true)]
        public Input<string> Registry { get; set; } = null!;

        public DeviceArgs()
        {
        }
    }

    public sealed class DeviceState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// If a device is blocked, connections or requests from this device will fail.
        /// </summary>
        [Input("blocked")]
        public Input<bool>? Blocked { get; set; }

        /// <summary>
        /// The most recent device configuration, which is eventually sent from Cloud IoT Core to the device.
        /// </summary>
        [Input("config")]
        public Input<Inputs.DeviceConfigGetArgs>? Config { get; set; }

        [Input("credentials")]
        private InputList<Inputs.DeviceCredentialGetArgs>? _credentials;

        /// <summary>
        /// The credentials used to authenticate this device.  Structure is documented below.
        /// </summary>
        public InputList<Inputs.DeviceCredentialGetArgs> Credentials
        {
            get => _credentials ?? (_credentials = new InputList<Inputs.DeviceCredentialGetArgs>());
            set => _credentials = value;
        }

        /// <summary>
        /// Gateway-related configuration and state.  Structure is documented below.
        /// </summary>
        [Input("gatewayConfig")]
        public Input<Inputs.DeviceGatewayConfigGetArgs>? GatewayConfig { get; set; }

        /// <summary>
        /// The last time a cloud-to-device config version acknowledgment was received from the device.
        /// </summary>
        [Input("lastConfigAckTime")]
        public Input<string>? LastConfigAckTime { get; set; }

        /// <summary>
        /// The last time a cloud-to-device config version was sent to the device.
        /// </summary>
        [Input("lastConfigSendTime")]
        public Input<string>? LastConfigSendTime { get; set; }

        /// <summary>
        /// The error message of the most recent error, such as a failure to publish to Cloud Pub/Sub.
        /// </summary>
        [Input("lastErrorStatus")]
        public Input<Inputs.DeviceLastErrorStatusGetArgs>? LastErrorStatus { get; set; }

        /// <summary>
        /// The time the most recent error occurred, such as a failure to publish to Cloud Pub/Sub.
        /// </summary>
        [Input("lastErrorTime")]
        public Input<string>? LastErrorTime { get; set; }

        /// <summary>
        /// The last time a telemetry event was received.
        /// </summary>
        [Input("lastEventTime")]
        public Input<string>? LastEventTime { get; set; }

        /// <summary>
        /// The last time an MQTT PINGREQ was received.
        /// </summary>
        [Input("lastHeartbeatTime")]
        public Input<string>? LastHeartbeatTime { get; set; }

        /// <summary>
        /// The last time a state event was received.
        /// </summary>
        [Input("lastStateTime")]
        public Input<string>? LastStateTime { get; set; }

        /// <summary>
        /// The logging verbosity for device activity.
        /// </summary>
        [Input("logLevel")]
        public Input<string>? LogLevel { get; set; }

        [Input("metadata")]
        private InputMap<string>? _metadata;

        /// <summary>
        /// The metadata key-value pairs assigned to the device.
        /// </summary>
        public InputMap<string> Metadata
        {
            get => _metadata ?? (_metadata = new InputMap<string>());
            set => _metadata = value;
        }

        /// <summary>
        /// A unique name for the resource.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// A server-defined unique numeric ID for the device. This is a more compact way to identify devices, and it is globally
        /// unique.
        /// </summary>
        [Input("numId")]
        public Input<string>? NumId { get; set; }

        /// <summary>
        /// The name of the device registry where this device should be created.
        /// </summary>
        [Input("registry")]
        public Input<string>? Registry { get; set; }

        /// <summary>
        /// The state most recently received from the device.
        /// </summary>
        [Input("state")]
        public Input<Inputs.DeviceStateGetArgs>? State { get; set; }

        public DeviceState()
        {
        }
    }
}
