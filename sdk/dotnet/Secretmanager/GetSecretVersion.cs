// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.SecretManager
{
    public static partial class Invokes
    {
        [Obsolete("Use GetSecretVersion.InvokeAsync() instead")]
        public static Task<GetSecretVersionResult> GetSecretVersion(GetSecretVersionArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetSecretVersionResult>("gcp:secretmanager/getSecretVersion:getSecretVersion", args ?? InvokeArgs.Empty, options.WithVersion());
    }
    public static class GetSecretVersion
    {
        public static Task<GetSecretVersionResult> InvokeAsync(GetSecretVersionArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetSecretVersionResult>("gcp:secretmanager/getSecretVersion:getSecretVersion", args ?? InvokeArgs.Empty, options.WithVersion());
    }

    public sealed class GetSecretVersionArgs : Pulumi.InvokeArgs
    {
        [Input("project")]
        public string? Project { get; set; }

        [Input("secret", required: true)]
        public string Secret { get; set; } = null!;

        [Input("version")]
        public string? Version { get; set; }

        public GetSecretVersionArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetSecretVersionResult
    {
        public readonly string CreateTime;
        public readonly string DestroyTime;
        public readonly bool Enabled;
        public readonly string Name;
        public readonly string Project;
        public readonly string Secret;
        public readonly string SecretData;
        public readonly string Version;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetSecretVersionResult(
            string createTime,
            string destroyTime,
            bool enabled,
            string name,
            string project,
            string secret,
            string secretData,
            string version,
            string id)
        {
            CreateTime = createTime;
            DestroyTime = destroyTime;
            Enabled = enabled;
            Name = name;
            Project = project;
            Secret = secret;
            SecretData = secretData;
            Version = version;
            Id = id;
        }
    }
}
