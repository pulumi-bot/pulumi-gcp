// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.CloudRun.Inputs
{

    public sealed class ServiceTemplateSpecGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// -
        /// (Optional)
        /// ContainerConcurrency specifies the maximum allowed in-flight (concurrent)
        /// requests per container of the Revision. Values are:
        /// - `0` thread-safe, the system should manage the max concurrency. This is
        /// the default value.
        /// - `1` not-thread-safe. Single concurrency
        /// - `2-N` thread-safe, max concurrency of N
        /// </summary>
        [Input("containerConcurrency")]
        public Input<int>? ContainerConcurrency { get; set; }

        [Input("containers")]
        private InputList<Inputs.ServiceTemplateSpecContainerGetArgs>? _containers;

        /// <summary>
        /// -
        /// (Required)
        /// Container defines the unit of execution for this Revision.
        /// In the context of a Revision, we disallow a number of the fields of
        /// this Container, including: name, ports, and volumeMounts.
        /// The runtime contract is documented here:
        /// https://github.com/knative/serving/blob/master/docs/runtime-contract.md  Structure is documented below.
        /// </summary>
        public InputList<Inputs.ServiceTemplateSpecContainerGetArgs> Containers
        {
            get => _containers ?? (_containers = new InputList<Inputs.ServiceTemplateSpecContainerGetArgs>());
            set => _containers = value;
        }

        /// <summary>
        /// -
        /// (Optional)
        /// Email address of the IAM service account associated with the revision of the
        /// service. The service account represents the identity of the running revision,
        /// and determines what permissions the revision has. If not provided, the revision
        /// will use the project's default service account.
        /// </summary>
        [Input("serviceAccountName")]
        public Input<string>? ServiceAccountName { get; set; }

        /// <summary>
        /// -
        /// ServingState holds a value describing the state the resources
        /// are in for this Revision.
        /// It is expected
        /// that the system will manipulate this based on routability and load.
        /// </summary>
        [Input("servingState")]
        public Input<string>? ServingState { get; set; }

        public ServiceTemplateSpecGetArgs()
        {
        }
    }
}
