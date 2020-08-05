# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import _utilities, _tables


class GetIAMPolicyResult:
    """
    A collection of values returned by getIAMPolicy.
    """
    def __init__(__self__, audit_configs=None, bindings=None, id=None, policy_data=None):
        if audit_configs and not isinstance(audit_configs, list):
            raise TypeError("Expected argument 'audit_configs' to be a list")
        __self__.audit_configs = audit_configs
        if bindings and not isinstance(bindings, list):
            raise TypeError("Expected argument 'bindings' to be a list")
        __self__.bindings = bindings
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        The provider-assigned unique ID for this managed resource.
        """
        if policy_data and not isinstance(policy_data, str):
            raise TypeError("Expected argument 'policy_data' to be a str")
        __self__.policy_data = policy_data
        """
        The above bindings serialized in a format suitable for
        referencing from a resource that supports IAM.
        """


class AwaitableGetIAMPolicyResult(GetIAMPolicyResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetIAMPolicyResult(
            audit_configs=self.audit_configs,
            bindings=self.bindings,
            id=self.id,
            policy_data=self.policy_data)


def get_iam_policy(audit_configs=None, bindings=None, opts=None):
    """
    Generates an IAM policy document that may be referenced by and applied to
    other Google Cloud Platform resources, such as the `organizations.Project` resource.

    **Note:** Several restrictions apply when setting IAM policies through this API.
    See the [setIamPolicy docs](https://cloud.google.com/resource-manager/reference/rest/v1/projects/setIamPolicy)
    for a list of these restrictions.

    ```python
    import pulumi
    import pulumi_gcp as gcp

    admin = gcp.organizations.get_iam_policy(audit_configs=[{
            "auditLogConfigs": [
                {
                    "exemptedMembers": ["user:you@domain.com"],
                    "logType": "DATA_READ",
                },
                {
                    "logType": "DATA_WRITE",
                },
                {
                    "logType": "ADMIN_READ",
                },
            ],
            "service": "cloudkms.googleapis.com",
        }],
        bindings=[
            {
                "members": ["serviceAccount:your-custom-sa@your-project.iam.gserviceaccount.com"],
                "role": "roles/compute.instanceAdmin",
            },
            {
                "members": ["user:alice@gmail.com"],
                "role": "roles/storage.objectViewer",
            },
        ])
    ```

    This data source is used to define IAM policies to apply to other resources.
    Currently, defining a policy through a datasource and referencing that policy
    from another resource is the only way to apply an IAM policy to a resource.


    :param list audit_configs: A nested configuration block that defines logging additional configuration for your project.
    :param list bindings: A nested configuration block (described below)
           defining a binding to be included in the policy document. Multiple
           `binding` arguments are supported.

    The **audit_configs** object supports the following:

      * `audit_log_configs` (`list`) - A nested block that defines the operations you'd like to log.
        * `exemptedMembers` (`list`) - Specifies the identities that are exempt from these types of logging operations. Follows the same format of the `members` array for `binding`.
        * `logType` (`str`) - Defines the logging level. `DATA_READ`, `DATA_WRITE` and `ADMIN_READ` capture different types of events. See [the audit configuration documentation](https://cloud.google.com/resource-manager/reference/rest/Shared.Types/AuditConfig) for more details.

      * `service` (`str`) - Defines a service that will be enabled for audit logging. For example, `storage.googleapis.com`, `cloudsql.googleapis.com`. `allServices` is a special value that covers all services.

    The **bindings** object supports the following:

      * `condition` (`dict`) - An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding. Structure is documented below.
        * `description` (`str`) - An optional description of the expression. This is a longer text which describes the expression, e.g. when hovered over it in a UI.
        * `expression` (`str`) - Textual representation of an expression in Common Expression Language syntax.
        * `title` (`str`) - A title for the expression, i.e. a short string describing its purpose.

      * `members` (`list`) - An array of identities that will be granted the privilege in the `role`. For more details on format and restrictions see https://cloud.google.com/billing/reference/rest/v1/Policy#Binding
        Each entry can have one of the following values:
        * **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account. It **can't** be used with the `organizations.Project` resource.
        * **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account. It **can't** be used with the `organizations.Project` resource.
        * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com.
        * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
        * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
        * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
      * `role` (`str`) - The role/permission that will be granted to the members.
        See the [IAM Roles](https://cloud.google.com/compute/docs/access/iam) documentation for a complete list of roles.
        Note that custom roles must be of the format `[projects|organizations]/{parent-name}/roles/{role-name}`.
    """
    __args__ = dict()
    __args__['auditConfigs'] = audit_configs
    __args__['bindings'] = bindings
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('gcp:organizations/getIAMPolicy:getIAMPolicy', __args__, opts=opts).value

    return AwaitableGetIAMPolicyResult(
        audit_configs=__ret__.get('auditConfigs'),
        bindings=__ret__.get('bindings'),
        id=__ret__.get('id'),
        policy_data=__ret__.get('policyData'))
