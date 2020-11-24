# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

# Export this package's modules as members:
from .config import *
from .config_iam_binding import *
from .config_iam_member import *
from .config_iam_policy import *
from .variable import *
from ._inputs import *
from . import outputs

def _register_module():
    import pulumi

    class Module(pulumi.runtime.ResourceModule):
        def version(self):
            return None

        def construct(self, name: str, typ: str, urn: str) -> pulumi.Resource:
            if typ == "gcp:runtimeconfig/config:Config":
                return Config(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "gcp:runtimeconfig/configIamBinding:ConfigIamBinding":
                return ConfigIamBinding(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "gcp:runtimeconfig/configIamMember:ConfigIamMember":
                return ConfigIamMember(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "gcp:runtimeconfig/configIamPolicy:ConfigIamPolicy":
                return ConfigIamPolicy(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "gcp:runtimeconfig/variable:Variable":
                return Variable(name, pulumi.ResourceOptions(urn=urn))
            else:
                raise Exception(f"unknown resource type {typ}")


    _module_instance = Module()
    pulumi.runtime.register_resource_module("gcp", "runtimeconfig/config", _module_instance)
    pulumi.runtime.register_resource_module("gcp", "runtimeconfig/configIamBinding", _module_instance)
    pulumi.runtime.register_resource_module("gcp", "runtimeconfig/configIamMember", _module_instance)
    pulumi.runtime.register_resource_module("gcp", "runtimeconfig/configIamPolicy", _module_instance)
    pulumi.runtime.register_resource_module("gcp", "runtimeconfig/variable", _module_instance)

_register_module()
