// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Organizations
{
    public partial class IamAuditConfig : Pulumi.CustomResource
    {
        /// <summary>
        /// The configuration for logging of each type of permission.  This can be specified multiple times.  Structure is documented below.
        /// </summary>
        [Output("auditLogConfigs")]
        public Output<ImmutableArray<Outputs.IamAuditConfigAuditLogConfigs>> AuditLogConfigs { get; private set; } = null!;

        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        /// <summary>
        /// The numeric ID of the organization in which you want to manage the audit logging config.
        /// </summary>
        [Output("orgId")]
        public Output<string> OrgId { get; private set; } = null!;

        /// <summary>
        /// Service which will be enabled for audit logging.  The special value `allServices` covers all services.  Note that if there are google\_organization\_iam\_audit\_config resources covering both `allServices` and a specific service then the union of the two AuditConfigs is used for that service: the `log_types` specified in each `audit_log_config` are enabled, and the `exempted_members` in each `audit_log_config` are exempted.
        /// </summary>
        [Output("service")]
        public Output<string> Service { get; private set; } = null!;


        /// <summary>
        /// Create a IamAuditConfig resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public IamAuditConfig(string name, IamAuditConfigArgs args, CustomResourceOptions? options = null)
            : base("gcp:organizations/iamAuditConfig:IamAuditConfig", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private IamAuditConfig(string name, Input<string> id, IamAuditConfigState? state = null, CustomResourceOptions? options = null)
            : base("gcp:organizations/iamAuditConfig:IamAuditConfig", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing IamAuditConfig resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static IamAuditConfig Get(string name, Input<string> id, IamAuditConfigState? state = null, CustomResourceOptions? options = null)
        {
            return new IamAuditConfig(name, id, state, options);
        }
    }

    public sealed class IamAuditConfigArgs : Pulumi.ResourceArgs
    {
        [Input("auditLogConfigs", required: true)]
        private InputList<Inputs.IamAuditConfigAuditLogConfigsArgs>? _auditLogConfigs;

        /// <summary>
        /// The configuration for logging of each type of permission.  This can be specified multiple times.  Structure is documented below.
        /// </summary>
        public InputList<Inputs.IamAuditConfigAuditLogConfigsArgs> AuditLogConfigs
        {
            get => _auditLogConfigs ?? (_auditLogConfigs = new InputList<Inputs.IamAuditConfigAuditLogConfigsArgs>());
            set => _auditLogConfigs = value;
        }

        /// <summary>
        /// The numeric ID of the organization in which you want to manage the audit logging config.
        /// </summary>
        [Input("orgId", required: true)]
        public Input<string> OrgId { get; set; } = null!;

        /// <summary>
        /// Service which will be enabled for audit logging.  The special value `allServices` covers all services.  Note that if there are google\_organization\_iam\_audit\_config resources covering both `allServices` and a specific service then the union of the two AuditConfigs is used for that service: the `log_types` specified in each `audit_log_config` are enabled, and the `exempted_members` in each `audit_log_config` are exempted.
        /// </summary>
        [Input("service", required: true)]
        public Input<string> Service { get; set; } = null!;

        public IamAuditConfigArgs()
        {
        }
    }

    public sealed class IamAuditConfigState : Pulumi.ResourceArgs
    {
        [Input("auditLogConfigs")]
        private InputList<Inputs.IamAuditConfigAuditLogConfigsGetArgs>? _auditLogConfigs;

        /// <summary>
        /// The configuration for logging of each type of permission.  This can be specified multiple times.  Structure is documented below.
        /// </summary>
        public InputList<Inputs.IamAuditConfigAuditLogConfigsGetArgs> AuditLogConfigs
        {
            get => _auditLogConfigs ?? (_auditLogConfigs = new InputList<Inputs.IamAuditConfigAuditLogConfigsGetArgs>());
            set => _auditLogConfigs = value;
        }

        [Input("etag")]
        public Input<string>? Etag { get; set; }

        /// <summary>
        /// The numeric ID of the organization in which you want to manage the audit logging config.
        /// </summary>
        [Input("orgId")]
        public Input<string>? OrgId { get; set; }

        /// <summary>
        /// Service which will be enabled for audit logging.  The special value `allServices` covers all services.  Note that if there are google\_organization\_iam\_audit\_config resources covering both `allServices` and a specific service then the union of the two AuditConfigs is used for that service: the `log_types` specified in each `audit_log_config` are enabled, and the `exempted_members` in each `audit_log_config` are exempted.
        /// </summary>
        [Input("service")]
        public Input<string>? Service { get; set; }

        public IamAuditConfigState()
        {
        }
    }

    namespace Inputs
    {

    public sealed class IamAuditConfigAuditLogConfigsArgs : Pulumi.ResourceArgs
    {
        [Input("exemptedMembers")]
        private InputList<string>? _exemptedMembers;

        /// <summary>
        /// Identities that do not cause logging for this type of permission.
        /// Each entry can have one of the following values:
        /// * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
        /// * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
        /// * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
        /// * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
        /// </summary>
        public InputList<string> ExemptedMembers
        {
            get => _exemptedMembers ?? (_exemptedMembers = new InputList<string>());
            set => _exemptedMembers = value;
        }

        /// <summary>
        /// Permission type for which logging is to be configured.  Must be one of `DATA_READ`, `DATA_WRITE`, or `ADMIN_READ`.
        /// </summary>
        [Input("logType", required: true)]
        public Input<string> LogType { get; set; } = null!;

        public IamAuditConfigAuditLogConfigsArgs()
        {
        }
    }

    public sealed class IamAuditConfigAuditLogConfigsGetArgs : Pulumi.ResourceArgs
    {
        [Input("exemptedMembers")]
        private InputList<string>? _exemptedMembers;

        /// <summary>
        /// Identities that do not cause logging for this type of permission.
        /// Each entry can have one of the following values:
        /// * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
        /// * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
        /// * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
        /// * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
        /// </summary>
        public InputList<string> ExemptedMembers
        {
            get => _exemptedMembers ?? (_exemptedMembers = new InputList<string>());
            set => _exemptedMembers = value;
        }

        /// <summary>
        /// Permission type for which logging is to be configured.  Must be one of `DATA_READ`, `DATA_WRITE`, or `ADMIN_READ`.
        /// </summary>
        [Input("logType", required: true)]
        public Input<string> LogType { get; set; } = null!;

        public IamAuditConfigAuditLogConfigsGetArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class IamAuditConfigAuditLogConfigs
    {
        /// <summary>
        /// Identities that do not cause logging for this type of permission.
        /// Each entry can have one of the following values:
        /// * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
        /// * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
        /// * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
        /// * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
        /// </summary>
        public readonly ImmutableArray<string> ExemptedMembers;
        /// <summary>
        /// Permission type for which logging is to be configured.  Must be one of `DATA_READ`, `DATA_WRITE`, or `ADMIN_READ`.
        /// </summary>
        public readonly string LogType;

        [OutputConstructor]
        private IamAuditConfigAuditLogConfigs(
            ImmutableArray<string> exemptedMembers,
            string logType)
        {
            ExemptedMembers = exemptedMembers;
            LogType = logType;
        }
    }
    }
}
