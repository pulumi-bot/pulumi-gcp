// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.CloudBuild.Inputs
{

    public sealed class TriggerGithubGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// -
        /// (Required)
        /// Name of the volume to mount.
        /// Volume names must be unique per build step and must be valid names for
        /// Docker volumes. Each named volume must be used by at least two build steps.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// -
        /// (Optional)
        /// Owner of the repository. For example: The owner for
        /// https://github.com/googlecloudplatform/cloud-builders is "googlecloudplatform".
        /// </summary>
        [Input("owner")]
        public Input<string>? Owner { get; set; }

        /// <summary>
        /// -
        /// (Optional)
        /// filter to match changes in pull requests.  Specify only one of pullRequest or push.  Structure is documented below.
        /// </summary>
        [Input("pullRequest")]
        public Input<Inputs.TriggerGithubPullRequestGetArgs>? PullRequest { get; set; }

        /// <summary>
        /// -
        /// (Optional)
        /// filter to match changes in refs, like branches or tags.  Specify only one of pullRequest or push.  Structure is documented below.
        /// </summary>
        [Input("push")]
        public Input<Inputs.TriggerGithubPushGetArgs>? Push { get; set; }

        public TriggerGithubGetArgs()
        {
        }
    }
}
