// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.CloudTasks.Inputs
{

    public sealed class QueueRateLimitsGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// -
        /// The max burst size.
        /// Max burst size limits how fast tasks in queue are processed when many tasks are
        /// in the queue and the rate is high. This field allows the queue to have a high
        /// rate so processing starts shortly after a task is enqueued, but still limits
        /// resource usage when many tasks are enqueued in a short period of time.
        /// </summary>
        [Input("maxBurstSize")]
        public Input<int>? MaxBurstSize { get; set; }

        /// <summary>
        /// The maximum number of concurrent tasks that Cloud Tasks allows to
        /// be dispatched for this queue. After this threshold has been
        /// reached, Cloud Tasks stops dispatching tasks until the number of
        /// concurrent requests decreases.
        /// </summary>
        [Input("maxConcurrentDispatches")]
        public Input<int>? MaxConcurrentDispatches { get; set; }

        /// <summary>
        /// The maximum rate at which tasks are dispatched from this queue.
        /// If unspecified when the queue is created, Cloud Tasks will pick the default.
        /// </summary>
        [Input("maxDispatchesPerSecond")]
        public Input<double>? MaxDispatchesPerSecond { get; set; }

        public QueueRateLimitsGetArgs()
        {
        }
    }
}
