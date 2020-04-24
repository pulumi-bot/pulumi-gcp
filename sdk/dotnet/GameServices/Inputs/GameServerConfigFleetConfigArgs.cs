// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.GameServices.Inputs
{

    public sealed class GameServerConfigFleetConfigArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The fleet spec, which is sent to Agones to configure fleet.
        /// The spec can be passed as inline json but it is recommended to use a file reference
        /// instead. File references can contain the json or yaml format of the fleet spec. Eg:
        /// * fleet_spec = jsonencode(yamldecode(file("fleet_configs.yaml")))
        /// * fleet_spec = file("fleet_configs.json")
        /// The format of the spec can be found :
        /// `https://agones.dev/site/docs/reference/fleet/`.
        /// </summary>
        [Input("fleetSpec", required: true)]
        public Input<string> FleetSpec { get; set; } = null!;

        /// <summary>
        /// The name of the ScalingConfig
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public GameServerConfigFleetConfigArgs()
        {
        }
    }
}
