// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Kms
{
    /// <summary>
    ///  Creates a device registry in Google's Cloud IoT Core platform. For more information see
    /// [the official documentation](https://cloud.google.com/iot/docs/) and
    /// [API](https://cloud.google.com/iot/docs/reference/cloudiot/rest/v1/projects.locations.registries).
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/cloudiot_registry.html.markdown.
    /// </summary>
    public partial class Registry : Pulumi.CustomResource
    {
        /// <summary>
        /// List of public key certificates to authenticate devices. Structure is documented below. 
        /// </summary>
        [Output("credentials")]
        public Output<ImmutableArray<Outputs.RegistryCredentials>> Credentials { get; private set; } = null!;

        /// <summary>
        /// List of configurations for event notification, such as
        /// PubSub topics to publish device events to. Structure is documented below.
        /// </summary>
        [Output("eventNotificationConfigs")]
        public Output<ImmutableArray<Outputs.RegistryEventNotificationConfigItem>> EventNotificationConfigs { get; private set; } = null!;

        /// <summary>
        /// Activate or deactivate HTTP. Structure is documented below.
        /// </summary>
        [Output("httpConfig")]
        public Output<Outputs.RegistryHttpConfig> HttpConfig { get; private set; } = null!;

        [Output("logLevel")]
        public Output<string?> LogLevel { get; private set; } = null!;

        /// <summary>
        /// Activate or deactivate MQTT. Structure is documented below.
        /// </summary>
        [Output("mqttConfig")]
        public Output<Outputs.RegistryMqttConfig> MqttConfig { get; private set; } = null!;

        /// <summary>
        /// A unique name for the resource, required by device registry.
        /// Changing this forces a new resource to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The project in which the resource belongs. If it is not provided, the provider project is used.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// The Region in which the created address should reside. If it is not provided, the provider region is used.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// A PubSub topic to publish device state updates. Structure is documented below.
        /// </summary>
        [Output("stateNotificationConfig")]
        public Output<Outputs.RegistryStateNotificationConfig?> StateNotificationConfig { get; private set; } = null!;


        /// <summary>
        /// Create a Registry resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Registry(string name, RegistryArgs? args = null, CustomResourceOptions? options = null)
            : base("gcp:kms/registry:Registry", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private Registry(string name, Input<string> id, RegistryState? state = null, CustomResourceOptions? options = null)
            : base("gcp:kms/registry:Registry", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Registry resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Registry Get(string name, Input<string> id, RegistryState? state = null, CustomResourceOptions? options = null)
        {
            return new Registry(name, id, state, options);
        }
    }

    public sealed class RegistryArgs : Pulumi.ResourceArgs
    {
        [Input("credentials")]
        private InputList<Inputs.RegistryCredentialsArgs>? _credentials;

        /// <summary>
        /// List of public key certificates to authenticate devices. Structure is documented below. 
        /// </summary>
        public InputList<Inputs.RegistryCredentialsArgs> Credentials
        {
            get => _credentials ?? (_credentials = new InputList<Inputs.RegistryCredentialsArgs>());
            set => _credentials = value;
        }

        [Input("eventNotificationConfigs")]
        private InputList<Inputs.RegistryEventNotificationConfigItemArgs>? _eventNotificationConfigs;

        /// <summary>
        /// List of configurations for event notification, such as
        /// PubSub topics to publish device events to. Structure is documented below.
        /// </summary>
        public InputList<Inputs.RegistryEventNotificationConfigItemArgs> EventNotificationConfigs
        {
            get => _eventNotificationConfigs ?? (_eventNotificationConfigs = new InputList<Inputs.RegistryEventNotificationConfigItemArgs>());
            set => _eventNotificationConfigs = value;
        }

        /// <summary>
        /// Activate or deactivate HTTP. Structure is documented below.
        /// </summary>
        [Input("httpConfig")]
        public Input<Inputs.RegistryHttpConfigArgs>? HttpConfig { get; set; }

        [Input("logLevel")]
        public Input<string>? LogLevel { get; set; }

        /// <summary>
        /// Activate or deactivate MQTT. Structure is documented below.
        /// </summary>
        [Input("mqttConfig")]
        public Input<Inputs.RegistryMqttConfigArgs>? MqttConfig { get; set; }

        /// <summary>
        /// A unique name for the resource, required by device registry.
        /// Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The project in which the resource belongs. If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The Region in which the created address should reside. If it is not provided, the provider region is used.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// A PubSub topic to publish device state updates. Structure is documented below.
        /// </summary>
        [Input("stateNotificationConfig")]
        public Input<Inputs.RegistryStateNotificationConfigArgs>? StateNotificationConfig { get; set; }

        public RegistryArgs()
        {
        }
    }

    public sealed class RegistryState : Pulumi.ResourceArgs
    {
        [Input("credentials")]
        private InputList<Inputs.RegistryCredentialsGetArgs>? _credentials;

        /// <summary>
        /// List of public key certificates to authenticate devices. Structure is documented below. 
        /// </summary>
        public InputList<Inputs.RegistryCredentialsGetArgs> Credentials
        {
            get => _credentials ?? (_credentials = new InputList<Inputs.RegistryCredentialsGetArgs>());
            set => _credentials = value;
        }

        [Input("eventNotificationConfigs")]
        private InputList<Inputs.RegistryEventNotificationConfigItemGetArgs>? _eventNotificationConfigs;

        /// <summary>
        /// List of configurations for event notification, such as
        /// PubSub topics to publish device events to. Structure is documented below.
        /// </summary>
        public InputList<Inputs.RegistryEventNotificationConfigItemGetArgs> EventNotificationConfigs
        {
            get => _eventNotificationConfigs ?? (_eventNotificationConfigs = new InputList<Inputs.RegistryEventNotificationConfigItemGetArgs>());
            set => _eventNotificationConfigs = value;
        }

        /// <summary>
        /// Activate or deactivate HTTP. Structure is documented below.
        /// </summary>
        [Input("httpConfig")]
        public Input<Inputs.RegistryHttpConfigGetArgs>? HttpConfig { get; set; }

        [Input("logLevel")]
        public Input<string>? LogLevel { get; set; }

        /// <summary>
        /// Activate or deactivate MQTT. Structure is documented below.
        /// </summary>
        [Input("mqttConfig")]
        public Input<Inputs.RegistryMqttConfigGetArgs>? MqttConfig { get; set; }

        /// <summary>
        /// A unique name for the resource, required by device registry.
        /// Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The project in which the resource belongs. If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The Region in which the created address should reside. If it is not provided, the provider region is used.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// A PubSub topic to publish device state updates. Structure is documented below.
        /// </summary>
        [Input("stateNotificationConfig")]
        public Input<Inputs.RegistryStateNotificationConfigGetArgs>? StateNotificationConfig { get; set; }

        public RegistryState()
        {
        }
    }

    namespace Inputs
    {

    public sealed class RegistryCredentialsArgs : Pulumi.ResourceArgs
    {
        [Input("publicKeyCertificate", required: true)]
        public Input<RegistryCredentialsPublicKeyCertificateArgs> PublicKeyCertificate { get; set; } = null!;

        public RegistryCredentialsArgs()
        {
        }
    }

    public sealed class RegistryCredentialsGetArgs : Pulumi.ResourceArgs
    {
        [Input("publicKeyCertificate", required: true)]
        public Input<RegistryCredentialsPublicKeyCertificateGetArgs> PublicKeyCertificate { get; set; } = null!;

        public RegistryCredentialsGetArgs()
        {
        }
    }

    public sealed class RegistryCredentialsPublicKeyCertificateArgs : Pulumi.ResourceArgs
    {
        [Input("certificate", required: true)]
        public Input<string> Certificate { get; set; } = null!;

        [Input("format", required: true)]
        public Input<string> Format { get; set; } = null!;

        public RegistryCredentialsPublicKeyCertificateArgs()
        {
        }
    }

    public sealed class RegistryCredentialsPublicKeyCertificateGetArgs : Pulumi.ResourceArgs
    {
        [Input("certificate", required: true)]
        public Input<string> Certificate { get; set; } = null!;

        [Input("format", required: true)]
        public Input<string> Format { get; set; } = null!;

        public RegistryCredentialsPublicKeyCertificateGetArgs()
        {
        }
    }

    public sealed class RegistryEventNotificationConfigItemArgs : Pulumi.ResourceArgs
    {
        [Input("pubsubTopicName", required: true)]
        public Input<string> PubsubTopicName { get; set; } = null!;

        [Input("subfolderMatches")]
        public Input<string>? SubfolderMatches { get; set; }

        public RegistryEventNotificationConfigItemArgs()
        {
        }
    }

    public sealed class RegistryEventNotificationConfigItemGetArgs : Pulumi.ResourceArgs
    {
        [Input("pubsubTopicName", required: true)]
        public Input<string> PubsubTopicName { get; set; } = null!;

        [Input("subfolderMatches")]
        public Input<string>? SubfolderMatches { get; set; }

        public RegistryEventNotificationConfigItemGetArgs()
        {
        }
    }

    public sealed class RegistryHttpConfigArgs : Pulumi.ResourceArgs
    {
        [Input("httpEnabledState", required: true)]
        public Input<string> HttpEnabledState { get; set; } = null!;

        public RegistryHttpConfigArgs()
        {
        }
    }

    public sealed class RegistryHttpConfigGetArgs : Pulumi.ResourceArgs
    {
        [Input("httpEnabledState", required: true)]
        public Input<string> HttpEnabledState { get; set; } = null!;

        public RegistryHttpConfigGetArgs()
        {
        }
    }

    public sealed class RegistryMqttConfigArgs : Pulumi.ResourceArgs
    {
        [Input("mqttEnabledState", required: true)]
        public Input<string> MqttEnabledState { get; set; } = null!;

        public RegistryMqttConfigArgs()
        {
        }
    }

    public sealed class RegistryMqttConfigGetArgs : Pulumi.ResourceArgs
    {
        [Input("mqttEnabledState", required: true)]
        public Input<string> MqttEnabledState { get; set; } = null!;

        public RegistryMqttConfigGetArgs()
        {
        }
    }

    public sealed class RegistryStateNotificationConfigArgs : Pulumi.ResourceArgs
    {
        [Input("pubsubTopicName", required: true)]
        public Input<string> PubsubTopicName { get; set; } = null!;

        public RegistryStateNotificationConfigArgs()
        {
        }
    }

    public sealed class RegistryStateNotificationConfigGetArgs : Pulumi.ResourceArgs
    {
        [Input("pubsubTopicName", required: true)]
        public Input<string> PubsubTopicName { get; set; } = null!;

        public RegistryStateNotificationConfigGetArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class RegistryCredentials
    {
        public readonly RegistryCredentialsPublicKeyCertificate PublicKeyCertificate;

        [OutputConstructor]
        private RegistryCredentials(RegistryCredentialsPublicKeyCertificate publicKeyCertificate)
        {
            PublicKeyCertificate = publicKeyCertificate;
        }
    }

    [OutputType]
    public sealed class RegistryCredentialsPublicKeyCertificate
    {
        public readonly string Certificate;
        public readonly string Format;

        [OutputConstructor]
        private RegistryCredentialsPublicKeyCertificate(
            string certificate,
            string format)
        {
            Certificate = certificate;
            Format = format;
        }
    }

    [OutputType]
    public sealed class RegistryEventNotificationConfigItem
    {
        public readonly string PubsubTopicName;
        public readonly string? SubfolderMatches;

        [OutputConstructor]
        private RegistryEventNotificationConfigItem(
            string pubsubTopicName,
            string? subfolderMatches)
        {
            PubsubTopicName = pubsubTopicName;
            SubfolderMatches = subfolderMatches;
        }
    }

    [OutputType]
    public sealed class RegistryHttpConfig
    {
        public readonly string HttpEnabledState;

        [OutputConstructor]
        private RegistryHttpConfig(string httpEnabledState)
        {
            HttpEnabledState = httpEnabledState;
        }
    }

    [OutputType]
    public sealed class RegistryMqttConfig
    {
        public readonly string MqttEnabledState;

        [OutputConstructor]
        private RegistryMqttConfig(string mqttEnabledState)
        {
            MqttEnabledState = mqttEnabledState;
        }
    }

    [OutputType]
    public sealed class RegistryStateNotificationConfig
    {
        public readonly string PubsubTopicName;

        [OutputConstructor]
        private RegistryStateNotificationConfig(string pubsubTopicName)
        {
            PubsubTopicName = pubsubTopicName;
        }
    }
    }
}
