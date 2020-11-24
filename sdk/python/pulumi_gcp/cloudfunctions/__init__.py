# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

# Export this package's modules as members:
from .function import *
from .function_iam_binding import *
from .function_iam_member import *
from .function_iam_policy import *
from .get_function import *
from ._inputs import *
from . import outputs

def _register_module():
    import pulumi

    class Module(pulumi.runtime.ResourceModule):
        def version(self):
            return None

        def construct(self, name: str, typ: str, urn: str) -> pulumi.Resource:
            if typ == "gcp:cloudfunctions/function:Function":
                return Function(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "gcp:cloudfunctions/functionIamBinding:FunctionIamBinding":
                return FunctionIamBinding(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "gcp:cloudfunctions/functionIamMember:FunctionIamMember":
                return FunctionIamMember(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "gcp:cloudfunctions/functionIamPolicy:FunctionIamPolicy":
                return FunctionIamPolicy(name, pulumi.ResourceOptions(urn=urn))
            else:
                raise Exception(f"unknown resource type {typ}")


    _module_instance = Module()
    pulumi.runtime.register_resource_module("gcp", "cloudfunctions/function", _module_instance)
    pulumi.runtime.register_resource_module("gcp", "cloudfunctions/functionIamBinding", _module_instance)
    pulumi.runtime.register_resource_module("gcp", "cloudfunctions/functionIamMember", _module_instance)
    pulumi.runtime.register_resource_module("gcp", "cloudfunctions/functionIamPolicy", _module_instance)

_register_module()
