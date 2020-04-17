// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Compute.Inputs
{

    public sealed class RegionHealthCheckHttpsHealthCheckGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// -
        /// (Optional)
        /// The value of the host header in the HTTP health check request.
        /// If left empty (default value), the public IP on behalf of which this health
        /// check is performed will be used.
        /// </summary>
        [Input("host")]
        public Input<string>? Host { get; set; }

        /// <summary>
        /// -
        /// (Optional)
        /// The TCP port number for the HTTP health check request.
        /// The default value is 80.
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        /// <summary>
        /// -
        /// (Optional)
        /// Port name as defined in InstanceGroup#NamedPort#name. If both port and
        /// port_name are defined, port takes precedence.
        /// </summary>
        [Input("portName")]
        public Input<string>? PortName { get; set; }

        /// <summary>
        /// -
        /// (Optional)
        /// Specifies how port is selected for health checking, can be one of the
        /// following values:
        /// * `USE_FIXED_PORT`: The port number in `port` is used for health checking.
        /// * `USE_NAMED_PORT`: The `portName` is used for health checking.
        /// * `USE_SERVING_PORT`: For NetworkEndpointGroup, the port specified for each
        /// network endpoint is used for health checking. For other backends, the
        /// port or named port specified in the Backend Service is used for health
        /// checking.
        /// If not specified, HTTP health check follows behavior specified in `port` and
        /// `portName` fields.
        /// </summary>
        [Input("portSpecification")]
        public Input<string>? PortSpecification { get; set; }

        /// <summary>
        /// -
        /// (Optional)
        /// Specifies the type of proxy header to append before sending data to the
        /// backend, either NONE or PROXY_V1. The default is NONE.
        /// </summary>
        [Input("proxyHeader")]
        public Input<string>? ProxyHeader { get; set; }

        /// <summary>
        /// -
        /// (Optional)
        /// The request path of the HTTP health check request.
        /// The default value is /.
        /// </summary>
        [Input("requestPath")]
        public Input<string>? RequestPath { get; set; }

        /// <summary>
        /// -
        /// (Optional)
        /// The bytes to match against the beginning of the response data. If left empty
        /// (the default value), any response will indicate health. The response data
        /// can only be ASCII.
        /// </summary>
        [Input("response")]
        public Input<string>? Response { get; set; }

        public RegionHealthCheckHttpsHealthCheckGetArgs()
        {
        }
    }
}
