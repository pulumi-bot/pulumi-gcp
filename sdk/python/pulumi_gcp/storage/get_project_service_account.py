# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GetProjectServiceAccountResult',
    'AwaitableGetProjectServiceAccountResult',
    'get_project_service_account',
]

@pulumi.output_type
class GetProjectServiceAccountResult:
    """
    A collection of values returned by getProjectServiceAccount.
    """
    def __init__(__self__, email_address=None, id=None, project=None, user_project=None):
        if email_address and not isinstance(email_address, str):
            raise TypeError("Expected argument 'email_address' to be a str")
        pulumi.set(__self__, "email_address", email_address)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if project and not isinstance(project, str):
            raise TypeError("Expected argument 'project' to be a str")
        pulumi.set(__self__, "project", project)
        if user_project and not isinstance(user_project, str):
            raise TypeError("Expected argument 'user_project' to be a str")
        pulumi.set(__self__, "user_project", user_project)

    @property
    @pulumi.getter(name="emailAddress")
    def email_address(self) -> str:
        """
        The email address of the service account. This value is often used to refer to the service account
        in order to grant IAM permissions.
        """
        return pulumi.get(self, "email_address")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def project(self) -> str:
        return pulumi.get(self, "project")

    @property
    @pulumi.getter(name="userProject")
    def user_project(self) -> Optional[str]:
        return pulumi.get(self, "user_project")


class AwaitableGetProjectServiceAccountResult(GetProjectServiceAccountResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetProjectServiceAccountResult(
            email_address=self.email_address,
            id=self.id,
            project=self.project,
            user_project=self.user_project)


def get_project_service_account(project: Optional[str] = None,
                                user_project: Optional[str] = None,
                                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetProjectServiceAccountResult:
    """
    Get the email address of a project's unique [automatic Google Cloud Storage service account](https://cloud.google.com/storage/docs/projects#service-accounts).

    For each Google Cloud project, Google maintains a unique service account which
    is used as the identity for various Google Cloud Storage operations, including
    operations involving
    [customer-managed encryption keys](https://cloud.google.com/storage/docs/encryption/customer-managed-keys)
    and those involving
    [storage notifications to pub/sub](https://cloud.google.com/storage/docs/gsutil/commands/notification).
    This automatic Google service account requires access to the relevant Cloud KMS keys or pub/sub topics, respectively, in order for Cloud Storage to use
    these customer-managed resources.

    The service account has a well-known, documented naming format which is parameterised on the numeric Google project ID.
    However, as noted in [the docs](https://cloud.google.com/storage/docs/projects#service-accounts), it is only created when certain relevant actions occur which
    presuppose its existence.
    These actions include calling a [Cloud Storage API endpoint](https://cloud.google.com/storage/docs/json_api/v1/projects/serviceAccount/get) to yield the
    service account's identity, or performing some operations in the UI which must use the service account's identity, such as attempting to list Cloud KMS keys
    on the bucket creation page.

    Use of this data source calls the relevant API endpoint to obtain the service account's identity and thus ensures it exists prior to any API operations
    which demand its existence, such as specifying it in Cloud IAM policy.
    Always prefer to use this data source over interpolating the project ID into the well-known format for this service account, as the latter approach may cause
    provider update errors in cases where the service account does not yet exist.

    >  When you write provider code which uses features depending on this service account *and* your provider code adds the service account in IAM policy on other resources,
       you must take care for race conditions between the establishment of the IAM policy and creation of the relevant Cloud Storage resource.
       Cloud Storage APIs will require permissions on resources such as pub/sub topics or Cloud KMS keys to exist *before* the attempt to utilise them in a
       bucket configuration, otherwise the API calls will fail.
       You may need to use `depends_on` to create an explicit dependency between the IAM policy resource and the Cloud Storage resource which depends on it.
       See the examples here and in the `storage.Notification` resource.

    For more information see
    [the API reference](https://cloud.google.com/storage/docs/json_api/v1/projects/serviceAccount).

    ## Example Usage
    ### Pub/Sub Notifications

    ```python
    import pulumi
    import pulumi_gcp as gcp

    gcs_account = gcp.storage.get_project_service_account()
    binding = gcp.pubsub.TopicIAMBinding("binding",
        topic=google_pubsub_topic["topic"]["name"],
        role="roles/pubsub.publisher",
        members=[f"serviceAccount:{gcs_account.email_address}"])
    ```
    ### Cloud KMS Keys

    ```python
    import pulumi
    import pulumi_gcp as gcp

    gcs_account = gcp.storage.get_project_service_account()
    binding = gcp.kms.CryptoKeyIAMBinding("binding",
        crypto_key_id="your-crypto-key-id",
        role="roles/cloudkms.cryptoKeyEncrypterDecrypter",
        members=[f"serviceAccount:{gcs_account.email_address}"])
    bucket = gcp.storage.Bucket("bucket", encryption=gcp.storage.BucketEncryptionArgs(
        default_kms_key_name="your-crypto-key-id",
    ),
    opts=pulumi.ResourceOptions(depends_on=[binding]))
    ```


    :param str project: The project the unique service account was created for. If it is not provided, the provider project is used.
    :param str user_project: The project the lookup originates from. This field is used if you are making the request
           from a different account than the one you are finding the service account for.
    """
    __args__ = dict()
    __args__['project'] = project
    __args__['userProject'] = user_project
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('gcp:storage/getProjectServiceAccount:getProjectServiceAccount', __args__, opts=opts, typ=GetProjectServiceAccountResult).value

    return AwaitableGetProjectServiceAccountResult(
        email_address=__ret__.email_address,
        id=__ret__.id,
        project=__ret__.project,
        user_project=__ret__.user_project)
