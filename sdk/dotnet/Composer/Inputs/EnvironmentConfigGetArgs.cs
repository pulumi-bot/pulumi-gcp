// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Composer.Inputs
{

    public sealed class EnvironmentConfigGetArgs : Pulumi.ResourceArgs
    {
        [Input("airflowUri")]
        public Input<string>? AirflowUri { get; set; }

        [Input("dagGcsPrefix")]
        public Input<string>? DagGcsPrefix { get; set; }

        [Input("gkeCluster")]
        public Input<string>? GkeCluster { get; set; }

        /// <summary>
        /// The configuration used for the Kubernetes Engine cluster.  Structure is documented below.
        /// </summary>
        [Input("nodeConfig")]
        public Input<Inputs.EnvironmentConfigNodeConfigGetArgs>? NodeConfig { get; set; }

        /// <summary>
        /// The number of nodes in the Kubernetes Engine cluster that
        /// will be used to run this environment.
        /// </summary>
        [Input("nodeCount")]
        public Input<int>? NodeCount { get; set; }

        /// <summary>
        /// The configuration used for the Private IP Cloud Composer environment. Structure is documented below.
        /// </summary>
        [Input("privateEnvironmentConfig")]
        public Input<Inputs.EnvironmentConfigPrivateEnvironmentConfigGetArgs>? PrivateEnvironmentConfig { get; set; }

        /// <summary>
        /// The configuration settings for software inside the environment.  Structure is documented below.
        /// </summary>
        [Input("softwareConfig")]
        public Input<Inputs.EnvironmentConfigSoftwareConfigGetArgs>? SoftwareConfig { get; set; }

        public EnvironmentConfigGetArgs()
        {
        }
    }
}
