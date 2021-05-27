// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Compute
{
    public static class GetInstanceTemplate
    {
        /// <summary>
        /// Get information about a VM instance template resource within GCE. For more information see
        /// [the official documentation](https://cloud.google.com/compute/docs/instance-templates)
        /// and
        /// [API](https://cloud.google.com/compute/docs/reference/rest/v1/instanceTemplates).
        /// </summary>
        public static Task<GetInstanceTemplateResult> InvokeAsync(GetInstanceTemplateArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetInstanceTemplateResult>("gcp:compute/getInstanceTemplate:getInstanceTemplate", args ?? new GetInstanceTemplateArgs(), options.WithVersion());

        public static Output<GetInstanceTemplateResult> Invoke(GetInstanceTemplateOutputArgs args, InvokeOptions? options = null)
        {
            return Pulumi.Output.All(
                args.Filter.Box(),
                args.MostRecent.Box(),
                args.Name.Box(),
                args.Project.Box()
            ).Apply(a => {
                    var args = new GetInstanceTemplateArgs();
                    a[0].Set(args, nameof(args.Filter));
                    a[1].Set(args, nameof(args.MostRecent));
                    a[2].Set(args, nameof(args.Name));
                    a[3].Set(args, nameof(args.Project));
                    return InvokeAsync(args, options);
            });
        }
    }


    public sealed class GetInstanceTemplateArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// A filter to retrieve the instance templates.
        /// See [gcloud topic filters](https://cloud.google.com/sdk/gcloud/reference/topic/filters) for reference.
        /// If multiple instance templates match, either adjust the filter or specify `most_recent`. One of `name` or `filter` must be provided.
        /// </summary>
        [Input("filter")]
        public string? Filter { get; set; }

        /// <summary>
        /// If `filter` is provided, ensures the most recent template is returned when multiple instance templates match. One of `name` or `filter` must be provided.
        /// </summary>
        [Input("mostRecent")]
        public bool? MostRecent { get; set; }

        /// <summary>
        /// The name of the instance template. One of `name` or `filter` must be provided.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If `project` is not provideded, the provider project is used.
        /// </summary>
        [Input("project", required: true)]
        public string Project { get; set; } = null!;

        public GetInstanceTemplateArgs()
        {
        }
    }

    public sealed class GetInstanceTemplateOutputArgs
    {
        /// <summary>
        /// A filter to retrieve the instance templates.
        /// See [gcloud topic filters](https://cloud.google.com/sdk/gcloud/reference/topic/filters) for reference.
        /// If multiple instance templates match, either adjust the filter or specify `most_recent`. One of `name` or `filter` must be provided.
        /// </summary>
        [Input("filter")]
        public Input<string>? Filter { get; set; }

        /// <summary>
        /// If `filter` is provided, ensures the most recent template is returned when multiple instance templates match. One of `name` or `filter` must be provided.
        /// </summary>
        [Input("mostRecent")]
        public Input<bool>? MostRecent { get; set; }

        /// <summary>
        /// The name of the instance template. One of `name` or `filter` must be provided.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If `project` is not provideded, the provider project is used.
        /// </summary>
        [Input("project", required: true)]
        public Input<string> Project { get; set; } = null!;

        public GetInstanceTemplateOutputArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetInstanceTemplateResult
    {
        /// <summary>
        /// Whether to allow sending and receiving of
        /// packets with non-matching source or destination IPs. This defaults to false.
        /// </summary>
        public readonly bool CanIpForward;
        /// <summary>
        /// Enable [Confidential Mode](https://cloud.google.com/compute/confidential-vm/docs/about-cvm) on this VM.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetInstanceTemplateConfidentialInstanceConfigResult> ConfidentialInstanceConfigs;
        /// <summary>
        /// A brief description of this resource.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// Disks to attach to instances created from this template.
        /// This can be specified multiple times for multiple disks. Structure is
        /// documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetInstanceTemplateDiskResult> Disks;
        /// <summary>
        /// Enable [Virtual Displays](https://cloud.google.com/compute/docs/instances/enable-instance-virtual-display#verify_display_driver) on this instance.
        /// **Note**: `allow_stopping_for_update` must be set to true in order to update this field.
        /// </summary>
        public readonly bool EnableDisplay;
        public readonly string? Filter;
        /// <summary>
        /// List of the type and count of accelerator cards attached to the instance. Structure documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetInstanceTemplateGuestAcceleratorResult> GuestAccelerators;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A brief description to use for instances
        /// created from this template.
        /// </summary>
        public readonly string InstanceDescription;
        /// <summary>
        /// (Optional) A set of ket/value label pairs to assign to disk created from
        /// this template
        /// </summary>
        public readonly ImmutableDictionary<string, string> Labels;
        /// <summary>
        /// The machine type to create.
        /// </summary>
        public readonly string MachineType;
        /// <summary>
        /// Metadata key/value pairs to make available from
        /// within instances created from this template.
        /// </summary>
        public readonly ImmutableDictionary<string, object> Metadata;
        /// <summary>
        /// The unique fingerprint of the metadata.
        /// </summary>
        public readonly string MetadataFingerprint;
        /// <summary>
        /// An alternative to using the
        /// startup-script metadata key, mostly to match the compute_instance resource.
        /// This replaces the startup-script metadata key on the created instance and
        /// thus the two mechanisms are not allowed to be used simultaneously.
        /// </summary>
        public readonly string MetadataStartupScript;
        /// <summary>
        /// Specifies a minimum CPU platform. Applicable values are the friendly names of CPU platforms, such as
        /// `Intel Haswell` or `Intel Skylake`. See the complete list [here](https://cloud.google.com/compute/docs/instances/specify-min-cpu-platform).
        /// </summary>
        public readonly string MinCpuPlatform;
        public readonly bool? MostRecent;
        /// <summary>
        /// The name of the instance template. If you leave
        /// this blank, the provider will auto-generate a unique name.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// Creates a unique name beginning with the specified
        /// prefix. Conflicts with `name`.
        /// </summary>
        public readonly string NamePrefix;
        /// <summary>
        /// Networks to attach to instances created from
        /// this template. This can be specified multiple times for multiple networks.
        /// Structure is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetInstanceTemplateNetworkInterfaceResult> NetworkInterfaces;
        /// <summary>
        /// The ID of the project in which the resource belongs. If it
        /// is not provided, the provider project is used.
        /// </summary>
        public readonly string Project;
        /// <summary>
        /// An instance template is a global resource that is not
        /// bound to a zone or a region. However, you can still specify some regional
        /// resources in an instance template, which restricts the template to the
        /// region where that resource resides. For example, a custom `subnetwork`
        /// resource is tied to a specific region. Defaults to the region of the
        /// Provider if no value is given.
        /// </summary>
        public readonly string Region;
        /// <summary>
        /// The scheduling strategy to use. More details about
        /// this configuration option are detailed below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetInstanceTemplateSchedulingResult> Schedulings;
        /// <summary>
        /// The URI of the created resource.
        /// </summary>
        public readonly string SelfLink;
        /// <summary>
        /// Service account to attach to the instance. Structure is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetInstanceTemplateServiceAccountResult> ServiceAccounts;
        /// <summary>
        /// Enable [Shielded VM](https://cloud.google.com/security/shielded-cloud/shielded-vm) on this instance. Shielded VM provides verifiable integrity to prevent against malware and rootkits. Defaults to disabled. Structure is documented below.
        /// **Note**: `shielded_instance_config` can only be used with boot images with shielded vm support. See the complete list [here](https://cloud.google.com/compute/docs/images#shielded-images).
        /// </summary>
        public readonly ImmutableArray<Outputs.GetInstanceTemplateShieldedInstanceConfigResult> ShieldedInstanceConfigs;
        /// <summary>
        /// Tags to attach to the instance.
        /// </summary>
        public readonly ImmutableArray<string> Tags;
        /// <summary>
        /// The unique fingerprint of the tags.
        /// </summary>
        public readonly string TagsFingerprint;

        [OutputConstructor]
        private GetInstanceTemplateResult(
            bool canIpForward,

            ImmutableArray<Outputs.GetInstanceTemplateConfidentialInstanceConfigResult> confidentialInstanceConfigs,

            string description,

            ImmutableArray<Outputs.GetInstanceTemplateDiskResult> disks,

            bool enableDisplay,

            string? filter,

            ImmutableArray<Outputs.GetInstanceTemplateGuestAcceleratorResult> guestAccelerators,

            string id,

            string instanceDescription,

            ImmutableDictionary<string, string> labels,

            string machineType,

            ImmutableDictionary<string, object> metadata,

            string metadataFingerprint,

            string metadataStartupScript,

            string minCpuPlatform,

            bool? mostRecent,

            string? name,

            string namePrefix,

            ImmutableArray<Outputs.GetInstanceTemplateNetworkInterfaceResult> networkInterfaces,

            string project,

            string region,

            ImmutableArray<Outputs.GetInstanceTemplateSchedulingResult> schedulings,

            string selfLink,

            ImmutableArray<Outputs.GetInstanceTemplateServiceAccountResult> serviceAccounts,

            ImmutableArray<Outputs.GetInstanceTemplateShieldedInstanceConfigResult> shieldedInstanceConfigs,

            ImmutableArray<string> tags,

            string tagsFingerprint)
        {
            CanIpForward = canIpForward;
            ConfidentialInstanceConfigs = confidentialInstanceConfigs;
            Description = description;
            Disks = disks;
            EnableDisplay = enableDisplay;
            Filter = filter;
            GuestAccelerators = guestAccelerators;
            Id = id;
            InstanceDescription = instanceDescription;
            Labels = labels;
            MachineType = machineType;
            Metadata = metadata;
            MetadataFingerprint = metadataFingerprint;
            MetadataStartupScript = metadataStartupScript;
            MinCpuPlatform = minCpuPlatform;
            MostRecent = mostRecent;
            Name = name;
            NamePrefix = namePrefix;
            NetworkInterfaces = networkInterfaces;
            Project = project;
            Region = region;
            Schedulings = schedulings;
            SelfLink = selfLink;
            ServiceAccounts = serviceAccounts;
            ShieldedInstanceConfigs = shieldedInstanceConfigs;
            Tags = tags;
            TagsFingerprint = tagsFingerprint;
        }
    }
}
