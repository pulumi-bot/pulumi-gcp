// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.CloudRun.Outputs
{

    [OutputType]
    public sealed class ServiceTemplateSpecContainerEnv
    {
        /// <summary>
        /// -
        /// (Optional)
        /// Name of the environment variable.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// -
        /// (Optional)
        /// Variable references $(VAR_NAME) are expanded
        /// using the previous defined environment variables in the container and
        /// any route environment variables. If a variable cannot be resolved,
        /// the reference in the input string will be unchanged. The $(VAR_NAME)
        /// syntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped
        /// references will never be expanded, regardless of whether the variable
        /// exists or not.
        /// Defaults to "".
        /// </summary>
        public readonly string? Value;

        [OutputConstructor]
        private ServiceTemplateSpecContainerEnv(
            string? name,

            string? value)
        {
            Name = name;
            Value = value;
        }
    }
}
