# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

# Export this package's modules as members:
from .billing_account_bucket_config import *
from .billing_account_exclusion import *
from .billing_account_sink import *
from .folder_bucket_config import *
from .folder_exclusion import *
from .folder_sink import *
from .metric import *
from .organization_bucket_config import *
from .organization_exclusion import *
from .organization_sink import *
from .project_bucket_config import *
from .project_exclusion import *
from .project_sink import *
from ._inputs import *
from . import outputs

def _register_module():
    import pulumi

    class Module(pulumi.runtime.ResourceModule):
        def version(self):
            return None

        def construct(self, name: str, typ: str, urn: str) -> pulumi.Resource:
            if typ == "gcp:logging/billingAccountBucketConfig:BillingAccountBucketConfig":
                return BillingAccountBucketConfig(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "gcp:logging/billingAccountExclusion:BillingAccountExclusion":
                return BillingAccountExclusion(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "gcp:logging/billingAccountSink:BillingAccountSink":
                return BillingAccountSink(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "gcp:logging/folderBucketConfig:FolderBucketConfig":
                return FolderBucketConfig(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "gcp:logging/folderExclusion:FolderExclusion":
                return FolderExclusion(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "gcp:logging/folderSink:FolderSink":
                return FolderSink(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "gcp:logging/metric:Metric":
                return Metric(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "gcp:logging/organizationBucketConfig:OrganizationBucketConfig":
                return OrganizationBucketConfig(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "gcp:logging/organizationExclusion:OrganizationExclusion":
                return OrganizationExclusion(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "gcp:logging/organizationSink:OrganizationSink":
                return OrganizationSink(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "gcp:logging/projectBucketConfig:ProjectBucketConfig":
                return ProjectBucketConfig(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "gcp:logging/projectExclusion:ProjectExclusion":
                return ProjectExclusion(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "gcp:logging/projectSink:ProjectSink":
                return ProjectSink(name, pulumi.ResourceOptions(urn=urn))
            else:
                raise Exception(f"unknown resource type {typ}")


    _module_instance = Module()
    pulumi.runtime.register_resource_module("gcp", "logging/billingAccountBucketConfig", _module_instance)
    pulumi.runtime.register_resource_module("gcp", "logging/billingAccountExclusion", _module_instance)
    pulumi.runtime.register_resource_module("gcp", "logging/billingAccountSink", _module_instance)
    pulumi.runtime.register_resource_module("gcp", "logging/folderBucketConfig", _module_instance)
    pulumi.runtime.register_resource_module("gcp", "logging/folderExclusion", _module_instance)
    pulumi.runtime.register_resource_module("gcp", "logging/folderSink", _module_instance)
    pulumi.runtime.register_resource_module("gcp", "logging/metric", _module_instance)
    pulumi.runtime.register_resource_module("gcp", "logging/organizationBucketConfig", _module_instance)
    pulumi.runtime.register_resource_module("gcp", "logging/organizationExclusion", _module_instance)
    pulumi.runtime.register_resource_module("gcp", "logging/organizationSink", _module_instance)
    pulumi.runtime.register_resource_module("gcp", "logging/projectBucketConfig", _module_instance)
    pulumi.runtime.register_resource_module("gcp", "logging/projectExclusion", _module_instance)
    pulumi.runtime.register_resource_module("gcp", "logging/projectSink", _module_instance)

_register_module()
