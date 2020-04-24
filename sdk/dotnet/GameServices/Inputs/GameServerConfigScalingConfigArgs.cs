// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.GameServices.Inputs
{

    public sealed class GameServerConfigScalingConfigArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Fleet autoscaler spec, which is sent to Agones.
        /// Example spec can be found :
        /// https://agones.dev/site/docs/reference/fleetautoscaler/
        /// </summary>
        [Input("fleetAutoscalerSpec", required: true)]
        public Input<string> FleetAutoscalerSpec { get; set; } = null!;

        /// <summary>
        /// The name of the ScalingConfig
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("schedules")]
        private InputList<Inputs.GameServerConfigScalingConfigScheduleArgs>? _schedules;

        /// <summary>
        /// The schedules to which this scaling config applies.  Structure is documented below.
        /// </summary>
        public InputList<Inputs.GameServerConfigScalingConfigScheduleArgs> Schedules
        {
            get => _schedules ?? (_schedules = new InputList<Inputs.GameServerConfigScalingConfigScheduleArgs>());
            set => _schedules = value;
        }

        [Input("selectors")]
        private InputList<Inputs.GameServerConfigScalingConfigSelectorArgs>? _selectors;

        /// <summary>
        /// Labels used to identify the clusters to which this scaling config
        /// applies. A cluster is subject to this scaling config if its labels match
        /// any of the selector entries.  Structure is documented below.
        /// </summary>
        public InputList<Inputs.GameServerConfigScalingConfigSelectorArgs> Selectors
        {
            get => _selectors ?? (_selectors = new InputList<Inputs.GameServerConfigScalingConfigSelectorArgs>());
            set => _selectors = value;
        }

        public GameServerConfigScalingConfigArgs()
        {
        }
    }
}
