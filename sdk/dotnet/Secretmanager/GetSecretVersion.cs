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
        /// <summary>
        /// Get a Secret Manager secret's version. For more information see the [official documentation](https://cloud.google.com/secret-manager/docs/) and [API](https://cloud.google.com/secret-manager/docs/reference/rest/v1beta1/projects.secrets.versions).
        /// 
        /// 
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/d/datasource_google_secret_manager_secret_version.html.markdown.
        /// </summary>
        [Obsolete("Use GetSecretVersion.InvokeAsync() instead")]
        public static Task<GetSecretVersionResult> GetSecretVersion(GetSecretVersionArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetSecretVersionResult>("gcp:secretmanager/getSecretVersion:getSecretVersion", args ?? InvokeArgs.Empty, options.WithVersion());
    }
    public static class GetSecretVersion
    {
        /// <summary>
        /// Get a Secret Manager secret's version. For more information see the [official documentation](https://cloud.google.com/secret-manager/docs/) and [API](https://cloud.google.com/secret-manager/docs/reference/rest/v1beta1/projects.secrets.versions).
        /// 
        /// 
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/d/datasource_google_secret_manager_secret_version.html.markdown.
        /// </summary>
        public static Task<GetSecretVersionResult> InvokeAsync(GetSecretVersionArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetSecretVersionResult>("gcp:secretmanager/getSecretVersion:getSecretVersion", args ?? InvokeArgs.Empty, options.WithVersion());
    }

    public sealed class GetSecretVersionArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The project to get the secret version for. If it
        /// is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public string? Project { get; set; }

        /// <summary>
        /// The secret to get the secret version for.
        /// </summary>
        [Input("secret", required: true)]
        public string Secret { get; set; } = null!;

        /// <summary>
        /// The version of the secret to get. If it
        /// is not provided, the latest version is retrieved.
        /// </summary>
        [Input("version")]
        public string? Version { get; set; }

        public GetSecretVersionArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetSecretVersionResult
    {
        /// <summary>
        /// The time at which the Secret was created.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// The time at which the Secret was destroyed. Only present if state is DESTROYED.
        /// </summary>
        public readonly string DestroyTime;
        /// <summary>
        /// True if the current state of the SecretVersion is enabled.
        /// </summary>
        public readonly bool Enabled;
        /// <summary>
        /// The resource name of the SecretVersion. Format:
        /// `projects/{{project}}/secrets/{{secret_id}}/versions/{{version}}`
        /// </summary>
        public readonly string Name;
        public readonly string Project;
        public readonly string Secret;
        /// <summary>
        /// The secret data. No larger than 64KiB.
        /// </summary>
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
