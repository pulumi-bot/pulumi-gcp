// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.AccessContextManager.Inputs
{

    public sealed class ServicePerimeterStatusVpcAccessibleServicesGetArgs : Pulumi.ResourceArgs
    {
        [Input("allowedServices")]
        private InputList<string>? _allowedServices;

        /// <summary>
        /// The list of APIs usable within the Service Perimeter.
        /// Must be empty unless `enableRestriction` is True.
        /// </summary>
        public InputList<string> AllowedServices
        {
            get => _allowedServices ?? (_allowedServices = new InputList<string>());
            set => _allowedServices = value;
        }

        /// <summary>
        /// Whether to restrict API calls within the Service Perimeter to the
        /// list of APIs specified in 'allowedServices'.
        /// </summary>
        [Input("enableRestriction")]
        public Input<bool>? EnableRestriction { get; set; }

        public ServicePerimeterStatusVpcAccessibleServicesGetArgs()
        {
        }
    }
}
