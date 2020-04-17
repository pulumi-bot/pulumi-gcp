// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.CloudBuild.Inputs
{

    public sealed class TriggerBuildStepArgs : Pulumi.ResourceArgs
    {
        [Input("args")]
        private InputList<string>? _args;

        /// <summary>
        /// -
        /// (Optional)
        /// A list of arguments that will be presented to the step when it is started.
        /// If the image used to run the step's container has an entrypoint, the args
        /// are used as arguments to that entrypoint. If the image does not define an
        /// entrypoint, the first element in args is used as the entrypoint, and the
        /// remainder will be used as arguments.
        /// </summary>
        public InputList<string> Args
        {
            get => _args ?? (_args = new InputList<string>());
            set => _args = value;
        }

        /// <summary>
        /// -
        /// (Optional)
        /// Directory, relative to the source root, in which to run the build.
        /// This must be a relative path. If a step's dir is specified and
        /// is an absolute path, this value is ignored for that step's
        /// execution.
        /// </summary>
        [Input("dir")]
        public Input<string>? Dir { get; set; }

        /// <summary>
        /// -
        /// (Optional)
        /// Entrypoint to be used instead of the build step image's
        /// default entrypoint.
        /// If unset, the image's default entrypoint is used
        /// </summary>
        [Input("entrypoint")]
        public Input<string>? Entrypoint { get; set; }

        [Input("envs")]
        private InputList<string>? _envs;

        /// <summary>
        /// -
        /// (Optional)
        /// A list of environment variable definitions to be used when
        /// running a step.
        /// The elements are of the form "KEY=VALUE" for the environment variable
        /// "KEY" being given the value "VALUE".
        /// </summary>
        public InputList<string> Envs
        {
            get => _envs ?? (_envs = new InputList<string>());
            set => _envs = value;
        }

        /// <summary>
        /// -
        /// (Optional)
        /// Unique identifier for this build step, used in `wait_for` to
        /// reference this build step as a dependency.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// -
        /// (Optional)
        /// Name of the trigger. Must be unique within the project.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("secretEnvs")]
        private InputList<string>? _secretEnvs;

        /// <summary>
        /// -
        /// (Optional)
        /// A list of environment variables which are encrypted using
        /// a Cloud Key
        /// Management Service crypto key. These values must be specified in
        /// the build's `Secret`.
        /// </summary>
        public InputList<string> SecretEnvs
        {
            get => _secretEnvs ?? (_secretEnvs = new InputList<string>());
            set => _secretEnvs = value;
        }

        /// <summary>
        /// -
        /// (Optional)
        /// Amount of time that this build should be allowed to run, to second granularity.
        /// If this amount of time elapses, work on the build will cease and the build status will be TIMEOUT.
        /// This timeout must be equal to or greater than the sum of the timeouts for build steps within the build.
        /// The expected format is the number of seconds followed by s.
        /// Default time is ten minutes (600s).
        /// </summary>
        [Input("timeout")]
        public Input<string>? Timeout { get; set; }

        /// <summary>
        /// -
        /// (Optional)
        /// Output only. Stores timing information for executing this
        /// build step.
        /// </summary>
        [Input("timing")]
        public Input<string>? Timing { get; set; }

        [Input("volumes")]
        private InputList<Inputs.TriggerBuildStepVolumeArgs>? _volumes;

        /// <summary>
        /// -
        /// (Optional)
        /// List of volumes to mount into the build step.
        /// Each volume is created as an empty volume prior to execution of the
        /// build step. Upon completion of the build, volumes and their contents
        /// are discarded.
        /// Using a named volume in only one step is not valid as it is
        /// indicative of a build request with an incorrect configuration.  Structure is documented below.
        /// </summary>
        public InputList<Inputs.TriggerBuildStepVolumeArgs> Volumes
        {
            get => _volumes ?? (_volumes = new InputList<Inputs.TriggerBuildStepVolumeArgs>());
            set => _volumes = value;
        }

        [Input("waitFors")]
        private InputList<string>? _waitFors;

        /// <summary>
        /// -
        /// (Optional)
        /// The ID(s) of the step(s) that this build step depends on.
        /// This build step will not start until all the build steps in `wait_for`
        /// have completed successfully. If `wait_for` is empty, this build step
        /// will start when all previous build steps in the `Build.Steps` list
        /// have completed successfully.
        /// </summary>
        public InputList<string> WaitFors
        {
            get => _waitFors ?? (_waitFors = new InputList<string>());
            set => _waitFors = value;
        }

        public TriggerBuildStepArgs()
        {
        }
    }
}
