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
    public sealed class TriggerBuild
    {
        /// <summary>
        /// -
        /// (Optional)
        /// A list of images to be pushed upon the successful completion of all build steps.
        /// The images are pushed using the builder service account's credentials.
        /// The digests of the pushed images will be stored in the Build resource's results field.
        /// If any of the images fail to be pushed, the build status is marked FAILURE.
        /// </summary>
        public readonly ImmutableArray<string> Images;
        /// <summary>
        /// -
        /// (Required)
        /// The operations to be performed on the workspace.  Structure is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.TriggerBuildStep> Steps;
        /// <summary>
        /// -
        /// (Optional)
        /// Tags for annotation of a Build. These are not docker tags.
        /// </summary>
        public readonly ImmutableArray<string> Tags;
        /// <summary>
        /// -
        /// (Optional)
        /// Amount of time that this build should be allowed to run, to second granularity.
        /// If this amount of time elapses, work on the build will cease and the build status will be TIMEOUT.
        /// This timeout must be equal to or greater than the sum of the timeouts for build steps within the build.
        /// The expected format is the number of seconds followed by s.
        /// Default time is ten minutes (600s).
        /// </summary>
        public readonly string? Timeout;

        [OutputConstructor]
        private TriggerBuild(
            ImmutableArray<string> images,

            ImmutableArray<Outputs.TriggerBuildStep> steps,

            ImmutableArray<string> tags,

            string? timeout)
        {
            Images = images;
            Steps = steps;
            Tags = tags;
            Timeout = timeout;
        }
    }
}
