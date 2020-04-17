// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Compute.Outputs
{

    [OutputType]
    public sealed class HealthCheckHttp2HealthCheck
    {
        /// <summary>
        /// -
        /// (Optional)
        /// The value of the host header in the HTTP health check request.
        /// If left empty (default value), the public IP on behalf of which this health
        /// check is performed will be used.
        /// </summary>
        public readonly string? Host;
        /// <summary>
        /// -
        /// (Optional)
        /// The TCP port number for the HTTP health check request.
        /// The default value is 80.
        /// </summary>
        public readonly int? Port;
        /// <summary>
        /// -
        /// (Optional)
        /// Port name as defined in InstanceGroup#NamedPort#name. If both port and
        /// port_name are defined, port takes precedence.
        /// </summary>
        public readonly string? PortName;
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
        public readonly string? PortSpecification;
        /// <summary>
        /// -
        /// (Optional)
        /// Specifies the type of proxy header to append before sending data to the
        /// backend, either NONE or PROXY_V1. The default is NONE.
        /// </summary>
        public readonly string? ProxyHeader;
        /// <summary>
        /// -
        /// (Optional)
        /// The request path of the HTTP health check request.
        /// The default value is /.
        /// </summary>
        public readonly string? RequestPath;
        /// <summary>
        /// -
        /// (Optional)
        /// The bytes to match against the beginning of the response data. If left empty
        /// (the default value), any response will indicate health. The response data
        /// can only be ASCII.
        /// </summary>
        public readonly string? Response;

        [OutputConstructor]
        private HealthCheckHttp2HealthCheck(
            string? host,

            int? port,

            string? portName,

            string? portSpecification,

            string? proxyHeader,

            string? requestPath,

            string? response)
        {
            Host = host;
            Port = port;
            PortName = portName;
            PortSpecification = portSpecification;
            ProxyHeader = proxyHeader;
            RequestPath = requestPath;
            Response = response;
        }
    }
}
