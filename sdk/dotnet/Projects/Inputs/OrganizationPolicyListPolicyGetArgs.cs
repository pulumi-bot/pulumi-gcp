// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Projects.Inputs
{

    public sealed class OrganizationPolicyListPolicyGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// or `deny` - (Optional) One or the other must be set.
        /// </summary>
        [Input("allow")]
        public Input<Inputs.OrganizationPolicyListPolicyAllowGetArgs>? Allow { get; set; }

        [Input("deny")]
        public Input<Inputs.OrganizationPolicyListPolicyDenyGetArgs>? Deny { get; set; }

        /// <summary>
        /// If set to true, the values from the effective Policy of the parent resource
        /// are inherited, meaning the values set in this Policy are added to the values inherited up the hierarchy.
        /// </summary>
        [Input("inheritFromParent")]
        public Input<bool>? InheritFromParent { get; set; }

        /// <summary>
        /// The Google Cloud Console will try to default to a configuration that matches the value specified in this field.
        /// </summary>
        [Input("suggestedValue")]
        public Input<string>? SuggestedValue { get; set; }

        public OrganizationPolicyListPolicyGetArgs()
        {
        }
    }
}
