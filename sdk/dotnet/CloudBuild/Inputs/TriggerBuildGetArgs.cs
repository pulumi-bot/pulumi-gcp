// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.CloudBuild.Inputs
{

    public sealed class TriggerBuildGetArgs : Pulumi.ResourceArgs
    {
        [Input("images")]
        private InputList<string>? _images;

        /// <summary>
        /// -
        /// (Optional)
        /// A list of images to be pushed upon the successful completion of all build steps.
        /// The images are pushed using the builder service account's credentials.
        /// The digests of the pushed images will be stored in the Build resource's results field.
        /// If any of the images fail to be pushed, the build status is marked FAILURE.
        /// </summary>
        public InputList<string> Images
        {
            get => _images ?? (_images = new InputList<string>());
            set => _images = value;
        }

        [Input("steps", required: true)]
        private InputList<Inputs.TriggerBuildStepGetArgs>? _steps;

        /// <summary>
        /// -
        /// (Required)
        /// The operations to be performed on the workspace.  Structure is documented below.
        /// </summary>
        public InputList<Inputs.TriggerBuildStepGetArgs> Steps
        {
            get => _steps ?? (_steps = new InputList<Inputs.TriggerBuildStepGetArgs>());
            set => _steps = value;
        }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// -
        /// (Optional)
        /// Tags for annotation of a Build. These are not docker tags.
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        /// <summary>
        /// -
        /// (Optional)
        /// Time limit for executing this build step. If not defined,
        /// the step has no
        /// time limit and will be allowed to continue to run until either it
        /// completes or the build itself times out.
        /// </summary>
        [Input("timeout")]
        public Input<string>? Timeout { get; set; }

        public TriggerBuildGetArgs()
        {
        }
    }
}
