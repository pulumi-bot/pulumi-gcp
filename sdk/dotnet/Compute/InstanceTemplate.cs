// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Compute
{
    /// <summary>
    /// Manages a VM instance template resource within GCE. For more information see
    /// [the official documentation](https://cloud.google.com/compute/docs/instance-templates)
    /// and
    /// [API](https://cloud.google.com/compute/docs/reference/latest/instanceTemplates).
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var myImage = Output.Create(Gcp.Compute.GetImage.InvokeAsync(new Gcp.Compute.GetImageArgs
    ///         {
    ///             Family = "debian-9",
    ///             Project = "debian-cloud",
    ///         }));
    ///         var foobar = new Gcp.Compute.Disk("foobar", new Gcp.Compute.DiskArgs
    ///         {
    ///             Image = myImage.Apply(myImage =&gt; myImage.SelfLink),
    ///             Size = 10,
    ///             Type = "pd-ssd",
    ///             Zone = "us-central1-a",
    ///         });
    ///         var @default = new Gcp.Compute.InstanceTemplate("default", new Gcp.Compute.InstanceTemplateArgs
    ///         {
    ///             Description = "This template is used to create app server instances.",
    ///             Tags = 
    ///             {
    ///                 "foo",
    ///                 "bar",
    ///             },
    ///             Labels = 
    ///             {
    ///                 { "environment", "dev" },
    ///             },
    ///             InstanceDescription = "description assigned to instances",
    ///             MachineType = "e2-medium",
    ///             CanIpForward = false,
    ///             Scheduling = new Gcp.Compute.Inputs.InstanceTemplateSchedulingArgs
    ///             {
    ///                 AutomaticRestart = true,
    ///                 OnHostMaintenance = "MIGRATE",
    ///             },
    ///             Disks = 
    ///             {
    ///                 new Gcp.Compute.Inputs.InstanceTemplateDiskArgs
    ///                 {
    ///                     SourceImage = "debian-cloud/debian-9",
    ///                     AutoDelete = true,
    ///                     Boot = true,
    ///                 },
    ///                 new Gcp.Compute.Inputs.InstanceTemplateDiskArgs
    ///                 {
    ///                     Source = foobar.Name,
    ///                     AutoDelete = false,
    ///                     Boot = false,
    ///                 },
    ///             },
    ///             NetworkInterfaces = 
    ///             {
    ///                 new Gcp.Compute.Inputs.InstanceTemplateNetworkInterfaceArgs
    ///                 {
    ///                     Network = "default",
    ///                 },
    ///             },
    ///             Metadata = 
    ///             {
    ///                 { "foo", "bar" },
    ///             },
    ///             ServiceAccount = new Gcp.Compute.Inputs.InstanceTemplateServiceAccountArgs
    ///             {
    ///                 Scopes = 
    ///                 {
    ///                     "userinfo-email",
    ///                     "compute-ro",
    ///                     "storage-ro",
    ///                 },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ## Deploying the Latest Image
    /// 
    /// A common way to use instance templates and managed instance groups is to deploy the
    /// latest image in a family, usually the latest build of your application. There are two
    /// ways to do this in the provider, and they have their pros and cons. The difference ends
    /// up being in how "latest" is interpreted. You can either deploy the latest image available
    /// when the provider runs, or you can have each instance check what the latest image is when
    /// it's being created, either as part of a scaling event or being rebuilt by the instance
    /// group manager.
    /// 
    /// If you're not sure, we recommend deploying the latest image available when the provider runs,
    /// because this means all the instances in your group will be based on the same image, always,
    /// and means that no upgrades or changes to your instances happen outside of a `pulumi up`.
    /// You can achieve this by using the `gcp.compute.Image`
    /// data source, which will retrieve the latest image on every `pulumi apply`, and will update
    /// the template to use that specific image:
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var myImage = Output.Create(Gcp.Compute.GetImage.InvokeAsync(new Gcp.Compute.GetImageArgs
    ///         {
    ///             Family = "debian-9",
    ///             Project = "debian-cloud",
    ///         }));
    ///         var instanceTemplate = new Gcp.Compute.InstanceTemplate("instanceTemplate", new Gcp.Compute.InstanceTemplateArgs
    ///         {
    ///             NamePrefix = "instance-template-",
    ///             MachineType = "e2-medium",
    ///             Region = "us-central1",
    ///             Disks = 
    ///             {
    ///                 new Gcp.Compute.Inputs.InstanceTemplateDiskArgs
    ///                 {
    ///                     SourceImage = google_compute_image.My_image.Self_link,
    ///                 },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// To have instances update to the latest on every scaling event or instance re-creation,
    /// use the family as the image for the disk, and it will use GCP's default behavior, setting
    /// the image for the template to the family:
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var instanceTemplate = new Gcp.Compute.InstanceTemplate("instanceTemplate", new Gcp.Compute.InstanceTemplateArgs
    ///         {
    ///             Disks = 
    ///             {
    ///                 new Gcp.Compute.Inputs.InstanceTemplateDiskArgs
    ///                 {
    ///                     SourceImage = "debian-cloud/debian-9",
    ///                 },
    ///             },
    ///             MachineType = "e2-medium",
    ///             NamePrefix = "instance-template-",
    ///             Region = "us-central1",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Instance templates can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import gcp:compute/instanceTemplate:InstanceTemplate default projects/{{project}}/global/instanceTemplates/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:compute/instanceTemplate:InstanceTemplate default {{project}}/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:compute/instanceTemplate:InstanceTemplate default {{name}}
    /// ```
    /// 
    ///  [custom-vm-types]https://cloud.google.com/dataproc/docs/concepts/compute/custom-machine-types [network-tier]https://cloud.google.com/network-tiers/docs/overview
    /// </summary>
    [GcpResourceType("gcp:compute/instanceTemplate:InstanceTemplate")]
    public partial class InstanceTemplate : Pulumi.CustomResource
    {
        /// <summary>
        /// Whether to allow sending and receiving of
        /// packets with non-matching source or destination IPs. This defaults to false.
        /// </summary>
        [Output("canIpForward")]
        public Output<bool?> CanIpForward { get; private set; } = null!;

        /// <summary>
        /// ) - Enable [Confidential Mode](https://cloud.google.com/compute/confidential-vm/docs/about-cvm) on this VM.
        /// </summary>
        [Output("confidentialInstanceConfig")]
        public Output<Outputs.InstanceTemplateConfidentialInstanceConfig> ConfidentialInstanceConfig { get; private set; } = null!;

        /// <summary>
        /// A brief description of this resource.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Disks to attach to instances created from this template.
        /// This can be specified multiple times for multiple disks. Structure is
        /// documented below.
        /// </summary>
        [Output("disks")]
        public Output<ImmutableArray<Outputs.InstanceTemplateDisk>> Disks { get; private set; } = null!;

        /// <summary>
        /// Enable [Virtual Displays](https://cloud.google.com/compute/docs/instances/enable-instance-virtual-display#verify_display_driver) on this instance.
        /// **Note**: `allow_stopping_for_update` must be set to true in order to update this field.
        /// </summary>
        [Output("enableDisplay")]
        public Output<bool?> EnableDisplay { get; private set; } = null!;

        /// <summary>
        /// List of the type and count of accelerator cards attached to the instance. Structure documented below.
        /// </summary>
        [Output("guestAccelerators")]
        public Output<ImmutableArray<Outputs.InstanceTemplateGuestAccelerator>> GuestAccelerators { get; private set; } = null!;

        /// <summary>
        /// A brief description to use for instances
        /// created from this template.
        /// </summary>
        [Output("instanceDescription")]
        public Output<string?> InstanceDescription { get; private set; } = null!;

        /// <summary>
        /// A set of key/value label pairs to assign to instances
        /// created from this template,
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>?> Labels { get; private set; } = null!;

        /// <summary>
        /// The machine type to create.
        /// </summary>
        [Output("machineType")]
        public Output<string> MachineType { get; private set; } = null!;

        /// <summary>
        /// Metadata key/value pairs to make available from
        /// within instances created from this template.
        /// </summary>
        [Output("metadata")]
        public Output<ImmutableDictionary<string, object>?> Metadata { get; private set; } = null!;

        /// <summary>
        /// The unique fingerprint of the metadata.
        /// </summary>
        [Output("metadataFingerprint")]
        public Output<string> MetadataFingerprint { get; private set; } = null!;

        /// <summary>
        /// An alternative to using the
        /// startup-script metadata key, mostly to match the compute_instance resource.
        /// This replaces the startup-script metadata key on the created instance and
        /// thus the two mechanisms are not allowed to be used simultaneously.
        /// </summary>
        [Output("metadataStartupScript")]
        public Output<string?> MetadataStartupScript { get; private set; } = null!;

        /// <summary>
        /// Specifies a minimum CPU platform. Applicable values are the friendly names of CPU platforms, such as
        /// `Intel Haswell` or `Intel Skylake`. See the complete list [here](https://cloud.google.com/compute/docs/instances/specify-min-cpu-platform).
        /// </summary>
        [Output("minCpuPlatform")]
        public Output<string?> MinCpuPlatform { get; private set; } = null!;

        /// <summary>
        /// The name of the instance template. If you leave
        /// this blank, the provider will auto-generate a unique name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Creates a unique name beginning with the specified
        /// prefix. Conflicts with `name`.
        /// </summary>
        [Output("namePrefix")]
        public Output<string> NamePrefix { get; private set; } = null!;

        /// <summary>
        /// Networks to attach to instances created from
        /// this template. This can be specified multiple times for multiple networks.
        /// Structure is documented below.
        /// </summary>
        [Output("networkInterfaces")]
        public Output<ImmutableArray<Outputs.InstanceTemplateNetworkInterface>> NetworkInterfaces { get; private set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs. If it
        /// is not provided, the provider project is used.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// An instance template is a global resource that is not
        /// bound to a zone or a region. However, you can still specify some regional
        /// resources in an instance template, which restricts the template to the
        /// region where that resource resides. For example, a custom `subnetwork`
        /// resource is tied to a specific region. Defaults to the region of the
        /// Provider if no value is given.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// The scheduling strategy to use. More details about
        /// this configuration option are detailed below.
        /// </summary>
        [Output("scheduling")]
        public Output<Outputs.InstanceTemplateScheduling> Scheduling { get; private set; } = null!;

        /// <summary>
        /// The URI of the created resource.
        /// </summary>
        [Output("selfLink")]
        public Output<string> SelfLink { get; private set; } = null!;

        /// <summary>
        /// Service account to attach to the instance. Structure is documented below.
        /// </summary>
        [Output("serviceAccount")]
        public Output<Outputs.InstanceTemplateServiceAccount?> ServiceAccount { get; private set; } = null!;

        /// <summary>
        /// Enable [Shielded VM](https://cloud.google.com/security/shielded-cloud/shielded-vm) on this instance. Shielded VM provides verifiable integrity to prevent against malware and rootkits. Defaults to disabled. Structure is documented below.
        /// **Note**: `shielded_instance_config` can only be used with boot images with shielded vm support. See the complete list [here](https://cloud.google.com/compute/docs/images#shielded-images).
        /// </summary>
        [Output("shieldedInstanceConfig")]
        public Output<Outputs.InstanceTemplateShieldedInstanceConfig> ShieldedInstanceConfig { get; private set; } = null!;

        /// <summary>
        /// Tags to attach to the instance.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<string>> Tags { get; private set; } = null!;

        /// <summary>
        /// The unique fingerprint of the tags.
        /// </summary>
        [Output("tagsFingerprint")]
        public Output<string> TagsFingerprint { get; private set; } = null!;


        /// <summary>
        /// Create a InstanceTemplate resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public InstanceTemplate(string name, InstanceTemplateArgs args, CustomResourceOptions? options = null)
            : base("gcp:compute/instanceTemplate:InstanceTemplate", name, args ?? new InstanceTemplateArgs(), MakeResourceOptions(options, ""))
        {
        }

        private InstanceTemplate(string name, Input<string> id, InstanceTemplateState? state = null, CustomResourceOptions? options = null)
            : base("gcp:compute/instanceTemplate:InstanceTemplate", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing InstanceTemplate resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static InstanceTemplate Get(string name, Input<string> id, InstanceTemplateState? state = null, CustomResourceOptions? options = null)
        {
            return new InstanceTemplate(name, id, state, options);
        }
    }

    public sealed class InstanceTemplateArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether to allow sending and receiving of
        /// packets with non-matching source or destination IPs. This defaults to false.
        /// </summary>
        [Input("canIpForward")]
        public Input<bool>? CanIpForward { get; set; }

        /// <summary>
        /// ) - Enable [Confidential Mode](https://cloud.google.com/compute/confidential-vm/docs/about-cvm) on this VM.
        /// </summary>
        [Input("confidentialInstanceConfig")]
        public Input<Inputs.InstanceTemplateConfidentialInstanceConfigArgs>? ConfidentialInstanceConfig { get; set; }

        /// <summary>
        /// A brief description of this resource.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("disks", required: true)]
        private InputList<Inputs.InstanceTemplateDiskArgs>? _disks;

        /// <summary>
        /// Disks to attach to instances created from this template.
        /// This can be specified multiple times for multiple disks. Structure is
        /// documented below.
        /// </summary>
        public InputList<Inputs.InstanceTemplateDiskArgs> Disks
        {
            get => _disks ?? (_disks = new InputList<Inputs.InstanceTemplateDiskArgs>());
            set => _disks = value;
        }

        /// <summary>
        /// Enable [Virtual Displays](https://cloud.google.com/compute/docs/instances/enable-instance-virtual-display#verify_display_driver) on this instance.
        /// **Note**: `allow_stopping_for_update` must be set to true in order to update this field.
        /// </summary>
        [Input("enableDisplay")]
        public Input<bool>? EnableDisplay { get; set; }

        [Input("guestAccelerators")]
        private InputList<Inputs.InstanceTemplateGuestAcceleratorArgs>? _guestAccelerators;

        /// <summary>
        /// List of the type and count of accelerator cards attached to the instance. Structure documented below.
        /// </summary>
        public InputList<Inputs.InstanceTemplateGuestAcceleratorArgs> GuestAccelerators
        {
            get => _guestAccelerators ?? (_guestAccelerators = new InputList<Inputs.InstanceTemplateGuestAcceleratorArgs>());
            set => _guestAccelerators = value;
        }

        /// <summary>
        /// A brief description to use for instances
        /// created from this template.
        /// </summary>
        [Input("instanceDescription")]
        public Input<string>? InstanceDescription { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// A set of key/value label pairs to assign to instances
        /// created from this template,
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// The machine type to create.
        /// </summary>
        [Input("machineType", required: true)]
        public Input<string> MachineType { get; set; } = null!;

        [Input("metadata")]
        private InputMap<object>? _metadata;

        /// <summary>
        /// Metadata key/value pairs to make available from
        /// within instances created from this template.
        /// </summary>
        public InputMap<object> Metadata
        {
            get => _metadata ?? (_metadata = new InputMap<object>());
            set => _metadata = value;
        }

        /// <summary>
        /// An alternative to using the
        /// startup-script metadata key, mostly to match the compute_instance resource.
        /// This replaces the startup-script metadata key on the created instance and
        /// thus the two mechanisms are not allowed to be used simultaneously.
        /// </summary>
        [Input("metadataStartupScript")]
        public Input<string>? MetadataStartupScript { get; set; }

        /// <summary>
        /// Specifies a minimum CPU platform. Applicable values are the friendly names of CPU platforms, such as
        /// `Intel Haswell` or `Intel Skylake`. See the complete list [here](https://cloud.google.com/compute/docs/instances/specify-min-cpu-platform).
        /// </summary>
        [Input("minCpuPlatform")]
        public Input<string>? MinCpuPlatform { get; set; }

        /// <summary>
        /// The name of the instance template. If you leave
        /// this blank, the provider will auto-generate a unique name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Creates a unique name beginning with the specified
        /// prefix. Conflicts with `name`.
        /// </summary>
        [Input("namePrefix")]
        public Input<string>? NamePrefix { get; set; }

        [Input("networkInterfaces")]
        private InputList<Inputs.InstanceTemplateNetworkInterfaceArgs>? _networkInterfaces;

        /// <summary>
        /// Networks to attach to instances created from
        /// this template. This can be specified multiple times for multiple networks.
        /// Structure is documented below.
        /// </summary>
        public InputList<Inputs.InstanceTemplateNetworkInterfaceArgs> NetworkInterfaces
        {
            get => _networkInterfaces ?? (_networkInterfaces = new InputList<Inputs.InstanceTemplateNetworkInterfaceArgs>());
            set => _networkInterfaces = value;
        }

        /// <summary>
        /// The ID of the project in which the resource belongs. If it
        /// is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// An instance template is a global resource that is not
        /// bound to a zone or a region. However, you can still specify some regional
        /// resources in an instance template, which restricts the template to the
        /// region where that resource resides. For example, a custom `subnetwork`
        /// resource is tied to a specific region. Defaults to the region of the
        /// Provider if no value is given.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The scheduling strategy to use. More details about
        /// this configuration option are detailed below.
        /// </summary>
        [Input("scheduling")]
        public Input<Inputs.InstanceTemplateSchedulingArgs>? Scheduling { get; set; }

        /// <summary>
        /// Service account to attach to the instance. Structure is documented below.
        /// </summary>
        [Input("serviceAccount")]
        public Input<Inputs.InstanceTemplateServiceAccountArgs>? ServiceAccount { get; set; }

        /// <summary>
        /// Enable [Shielded VM](https://cloud.google.com/security/shielded-cloud/shielded-vm) on this instance. Shielded VM provides verifiable integrity to prevent against malware and rootkits. Defaults to disabled. Structure is documented below.
        /// **Note**: `shielded_instance_config` can only be used with boot images with shielded vm support. See the complete list [here](https://cloud.google.com/compute/docs/images#shielded-images).
        /// </summary>
        [Input("shieldedInstanceConfig")]
        public Input<Inputs.InstanceTemplateShieldedInstanceConfigArgs>? ShieldedInstanceConfig { get; set; }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// Tags to attach to the instance.
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        public InstanceTemplateArgs()
        {
        }
    }

    public sealed class InstanceTemplateState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether to allow sending and receiving of
        /// packets with non-matching source or destination IPs. This defaults to false.
        /// </summary>
        [Input("canIpForward")]
        public Input<bool>? CanIpForward { get; set; }

        /// <summary>
        /// ) - Enable [Confidential Mode](https://cloud.google.com/compute/confidential-vm/docs/about-cvm) on this VM.
        /// </summary>
        [Input("confidentialInstanceConfig")]
        public Input<Inputs.InstanceTemplateConfidentialInstanceConfigGetArgs>? ConfidentialInstanceConfig { get; set; }

        /// <summary>
        /// A brief description of this resource.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("disks")]
        private InputList<Inputs.InstanceTemplateDiskGetArgs>? _disks;

        /// <summary>
        /// Disks to attach to instances created from this template.
        /// This can be specified multiple times for multiple disks. Structure is
        /// documented below.
        /// </summary>
        public InputList<Inputs.InstanceTemplateDiskGetArgs> Disks
        {
            get => _disks ?? (_disks = new InputList<Inputs.InstanceTemplateDiskGetArgs>());
            set => _disks = value;
        }

        /// <summary>
        /// Enable [Virtual Displays](https://cloud.google.com/compute/docs/instances/enable-instance-virtual-display#verify_display_driver) on this instance.
        /// **Note**: `allow_stopping_for_update` must be set to true in order to update this field.
        /// </summary>
        [Input("enableDisplay")]
        public Input<bool>? EnableDisplay { get; set; }

        [Input("guestAccelerators")]
        private InputList<Inputs.InstanceTemplateGuestAcceleratorGetArgs>? _guestAccelerators;

        /// <summary>
        /// List of the type and count of accelerator cards attached to the instance. Structure documented below.
        /// </summary>
        public InputList<Inputs.InstanceTemplateGuestAcceleratorGetArgs> GuestAccelerators
        {
            get => _guestAccelerators ?? (_guestAccelerators = new InputList<Inputs.InstanceTemplateGuestAcceleratorGetArgs>());
            set => _guestAccelerators = value;
        }

        /// <summary>
        /// A brief description to use for instances
        /// created from this template.
        /// </summary>
        [Input("instanceDescription")]
        public Input<string>? InstanceDescription { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// A set of key/value label pairs to assign to instances
        /// created from this template,
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// The machine type to create.
        /// </summary>
        [Input("machineType")]
        public Input<string>? MachineType { get; set; }

        [Input("metadata")]
        private InputMap<object>? _metadata;

        /// <summary>
        /// Metadata key/value pairs to make available from
        /// within instances created from this template.
        /// </summary>
        public InputMap<object> Metadata
        {
            get => _metadata ?? (_metadata = new InputMap<object>());
            set => _metadata = value;
        }

        /// <summary>
        /// The unique fingerprint of the metadata.
        /// </summary>
        [Input("metadataFingerprint")]
        public Input<string>? MetadataFingerprint { get; set; }

        /// <summary>
        /// An alternative to using the
        /// startup-script metadata key, mostly to match the compute_instance resource.
        /// This replaces the startup-script metadata key on the created instance and
        /// thus the two mechanisms are not allowed to be used simultaneously.
        /// </summary>
        [Input("metadataStartupScript")]
        public Input<string>? MetadataStartupScript { get; set; }

        /// <summary>
        /// Specifies a minimum CPU platform. Applicable values are the friendly names of CPU platforms, such as
        /// `Intel Haswell` or `Intel Skylake`. See the complete list [here](https://cloud.google.com/compute/docs/instances/specify-min-cpu-platform).
        /// </summary>
        [Input("minCpuPlatform")]
        public Input<string>? MinCpuPlatform { get; set; }

        /// <summary>
        /// The name of the instance template. If you leave
        /// this blank, the provider will auto-generate a unique name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Creates a unique name beginning with the specified
        /// prefix. Conflicts with `name`.
        /// </summary>
        [Input("namePrefix")]
        public Input<string>? NamePrefix { get; set; }

        [Input("networkInterfaces")]
        private InputList<Inputs.InstanceTemplateNetworkInterfaceGetArgs>? _networkInterfaces;

        /// <summary>
        /// Networks to attach to instances created from
        /// this template. This can be specified multiple times for multiple networks.
        /// Structure is documented below.
        /// </summary>
        public InputList<Inputs.InstanceTemplateNetworkInterfaceGetArgs> NetworkInterfaces
        {
            get => _networkInterfaces ?? (_networkInterfaces = new InputList<Inputs.InstanceTemplateNetworkInterfaceGetArgs>());
            set => _networkInterfaces = value;
        }

        /// <summary>
        /// The ID of the project in which the resource belongs. If it
        /// is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// An instance template is a global resource that is not
        /// bound to a zone or a region. However, you can still specify some regional
        /// resources in an instance template, which restricts the template to the
        /// region where that resource resides. For example, a custom `subnetwork`
        /// resource is tied to a specific region. Defaults to the region of the
        /// Provider if no value is given.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The scheduling strategy to use. More details about
        /// this configuration option are detailed below.
        /// </summary>
        [Input("scheduling")]
        public Input<Inputs.InstanceTemplateSchedulingGetArgs>? Scheduling { get; set; }

        /// <summary>
        /// The URI of the created resource.
        /// </summary>
        [Input("selfLink")]
        public Input<string>? SelfLink { get; set; }

        /// <summary>
        /// Service account to attach to the instance. Structure is documented below.
        /// </summary>
        [Input("serviceAccount")]
        public Input<Inputs.InstanceTemplateServiceAccountGetArgs>? ServiceAccount { get; set; }

        /// <summary>
        /// Enable [Shielded VM](https://cloud.google.com/security/shielded-cloud/shielded-vm) on this instance. Shielded VM provides verifiable integrity to prevent against malware and rootkits. Defaults to disabled. Structure is documented below.
        /// **Note**: `shielded_instance_config` can only be used with boot images with shielded vm support. See the complete list [here](https://cloud.google.com/compute/docs/images#shielded-images).
        /// </summary>
        [Input("shieldedInstanceConfig")]
        public Input<Inputs.InstanceTemplateShieldedInstanceConfigGetArgs>? ShieldedInstanceConfig { get; set; }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// Tags to attach to the instance.
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The unique fingerprint of the tags.
        /// </summary>
        [Input("tagsFingerprint")]
        public Input<string>? TagsFingerprint { get; set; }

        public InstanceTemplateState()
        {
        }
    }
}
