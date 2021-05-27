// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.ServiceAccount
{
    public static class GetAccountAccessToken
    {
        /// <summary>
        /// This data source provides a google `oauth2` `access_token` for a different service account than the one initially running the script.
        /// 
        /// For more information see
        /// [the official documentation](https://cloud.google.com/iam/docs/creating-short-lived-service-account-credentials) as well as [iamcredentials.generateAccessToken()](https://cloud.google.com/iam/credentials/reference/rest/v1/projects.serviceAccounts/generateAccessToken)
        /// </summary>
        public static Task<GetAccountAccessTokenResult> InvokeAsync(GetAccountAccessTokenArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetAccountAccessTokenResult>("gcp:serviceAccount/getAccountAccessToken:getAccountAccessToken", args ?? new GetAccountAccessTokenArgs(), options.WithVersion());

        public static Output<GetAccountAccessTokenResult> Invoke(GetAccountAccessTokenOutputArgs args, InvokeOptions? options = null)
        {
            return Pulumi.Output.All(
                args.Delegates.ToList().Box(),
                args.Lifetime.Box(),
                args.Scopes.ToList().Box(),
                args.TargetServiceAccount.Box()
            ).Apply(a => {
                    var args = new GetAccountAccessTokenArgs();
                    a[0].Set(args, nameof(args.Delegates));
                    a[1].Set(args, nameof(args.Lifetime));
                    a[2].Set(args, nameof(args.Scopes));
                    a[3].Set(args, nameof(args.TargetServiceAccount));
                    return InvokeAsync(args, options);
            });
        }
    }


    public sealed class GetAccountAccessTokenArgs : Pulumi.InvokeArgs
    {
        [Input("delegates")]
        private List<string>? _delegates;

        /// <summary>
        /// Delegate chain of approvals needed to perform full impersonation. Specify the fully qualified service account name.  (e.g. `["projects/-/serviceAccounts/delegate-svc-account@project-id.iam.gserviceaccount.com"]`)
        /// </summary>
        public List<string> Delegates
        {
            get => _delegates ?? (_delegates = new List<string>());
            set => _delegates = value;
        }

        /// <summary>
        /// Lifetime of the impersonated token (defaults to its max: `3600s`).
        /// </summary>
        [Input("lifetime")]
        public string? Lifetime { get; set; }

        [Input("scopes", required: true)]
        private List<string>? _scopes;

        /// <summary>
        /// The scopes the new credential should have (e.g. `["cloud-platform"]`)
        /// </summary>
        public List<string> Scopes
        {
            get => _scopes ?? (_scopes = new List<string>());
            set => _scopes = value;
        }

        /// <summary>
        /// The service account _to_ impersonate (e.g. `service_B@your-project-id.iam.gserviceaccount.com`)
        /// </summary>
        [Input("targetServiceAccount", required: true)]
        public string TargetServiceAccount { get; set; } = null!;

        public GetAccountAccessTokenArgs()
        {
        }
    }

    public sealed class GetAccountAccessTokenOutputArgs
    {
        [Input("delegates")]
        private InputList<string>? _delegates;

        /// <summary>
        /// Delegate chain of approvals needed to perform full impersonation. Specify the fully qualified service account name.  (e.g. `["projects/-/serviceAccounts/delegate-svc-account@project-id.iam.gserviceaccount.com"]`)
        /// </summary>
        public InputList<string> Delegates
        {
            get => _delegates ?? (_delegates = new InputList<string>());
            set => _delegates = value;
        }

        /// <summary>
        /// Lifetime of the impersonated token (defaults to its max: `3600s`).
        /// </summary>
        [Input("lifetime")]
        public Input<string>? Lifetime { get; set; }

        [Input("scopes", required: true)]
        private InputList<string>? _scopes;

        /// <summary>
        /// The scopes the new credential should have (e.g. `["cloud-platform"]`)
        /// </summary>
        public InputList<string> Scopes
        {
            get => _scopes ?? (_scopes = new InputList<string>());
            set => _scopes = value;
        }

        /// <summary>
        /// The service account _to_ impersonate (e.g. `service_B@your-project-id.iam.gserviceaccount.com`)
        /// </summary>
        [Input("targetServiceAccount", required: true)]
        public Input<string> TargetServiceAccount { get; set; } = null!;

        public GetAccountAccessTokenOutputArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetAccountAccessTokenResult
    {
        /// <summary>
        /// The `access_token` representing the new generated identity.
        /// </summary>
        public readonly string AccessToken;
        public readonly ImmutableArray<string> Delegates;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? Lifetime;
        public readonly ImmutableArray<string> Scopes;
        public readonly string TargetServiceAccount;

        [OutputConstructor]
        private GetAccountAccessTokenResult(
            string accessToken,

            ImmutableArray<string> delegates,

            string id,

            string? lifetime,

            ImmutableArray<string> scopes,

            string targetServiceAccount)
        {
            AccessToken = accessToken;
            Delegates = delegates;
            Id = id;
            Lifetime = lifetime;
            Scopes = scopes;
            TargetServiceAccount = targetServiceAccount;
        }
    }
}
