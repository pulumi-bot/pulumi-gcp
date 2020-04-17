// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Composer.Outputs
{

    [OutputType]
    public sealed class EnvironmentConfigSoftwareConfig
    {
        public readonly ImmutableDictionary<string, string>? AirflowConfigOverrides;
        public readonly ImmutableDictionary<string, string>? EnvVariables;
        public readonly string? ImageVersion;
        public readonly ImmutableDictionary<string, string>? PypiPackages;
        public readonly string? PythonVersion;

        [OutputConstructor]
        private EnvironmentConfigSoftwareConfig(
            ImmutableDictionary<string, string>? airflowConfigOverrides,

            ImmutableDictionary<string, string>? envVariables,

            string? imageVersion,

            ImmutableDictionary<string, string>? pypiPackages,

            string? pythonVersion)
        {
            AirflowConfigOverrides = airflowConfigOverrides;
            EnvVariables = envVariables;
            ImageVersion = imageVersion;
            PypiPackages = pypiPackages;
            PythonVersion = pythonVersion;
        }
    }
}
