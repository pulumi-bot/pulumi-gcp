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
    public sealed class TriggerGithubPullRequest
    {
        /// <summary>
        /// -
        /// (Required)
        /// Regex of branches to match.
        /// </summary>
        public readonly string Branch;
        /// <summary>
        /// -
        /// (Optional)
        /// Whether to block builds on a "/gcbrun" comment from a repository owner or collaborator.
        /// </summary>
        public readonly string? CommentControl;

        [OutputConstructor]
        private TriggerGithubPullRequest(
            string branch,

            string? commentControl)
        {
            Branch = branch;
            CommentControl = commentControl;
        }
    }
}
