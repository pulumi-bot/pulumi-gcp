// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.CloudBuild.Outputs
{

    [OutputType]
    public sealed class TriggerBuildStepVolume
    {
        /// <summary>
        /// Name of the volume to mount.
        /// Volume names must be unique per build step and must be valid names for
        /// Docker volumes. Each named volume must be used by at least two build steps.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Path at which to mount the volume.
        /// Paths must be absolute and cannot conflict with other volume paths on
        /// the same build step or with certain reserved volume paths.
        /// </summary>
        public readonly string Path;

        [OutputConstructor]
        private TriggerBuildStepVolume(
            string name,

            string path)
        {
            Name = name;
            Path = path;
        }
    }
}
