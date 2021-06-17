// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Compute.Inputs
{

    public sealed class InstanceGroupManagerStatusStatefulArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A bit indicating whether the managed instance group has stateful configuration, that is, if you have configured any items in a stateful policy or in per-instance configs. The group might report that it has no stateful config even when there is still some preserved state on a managed instance, for example, if you have deleted all PICs but not yet applied those deletions.
        /// </summary>
        [Input("hasStatefulConfig")]
        public Input<bool>? HasStatefulConfig { get; set; }

        [Input("perInstanceConfigs")]
        private InputList<Inputs.InstanceGroupManagerStatusStatefulPerInstanceConfigArgs>? _perInstanceConfigs;

        /// <summary>
        /// Status of per-instance configs on the instance.
        /// </summary>
        public InputList<Inputs.InstanceGroupManagerStatusStatefulPerInstanceConfigArgs> PerInstanceConfigs
        {
            get => _perInstanceConfigs ?? (_perInstanceConfigs = new InputList<Inputs.InstanceGroupManagerStatusStatefulPerInstanceConfigArgs>());
            set => _perInstanceConfigs = value;
        }

        public InstanceGroupManagerStatusStatefulArgs()
        {
        }
    }
}
