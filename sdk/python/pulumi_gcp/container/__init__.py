# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

# Export this package's modules as members:
from .cluster import *
from .get_cluster import *
from .get_engine_versions import *
from .get_registry_image import *
from .get_registry_repository import *
from .node_pool import *
from .registry import *
from ._inputs import *
from . import outputs

def _register_module():
    import pulumi

    class Module(pulumi.runtime.ResourceModule):
        def version(self):
            return None

        def construct(self, name: str, typ: str, urn: str) -> pulumi.Resource:
            if typ == "gcp:container/cluster:Cluster":
                return Cluster(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "gcp:container/nodePool:NodePool":
                return NodePool(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "gcp:container/registry:Registry":
                return Registry(name, pulumi.ResourceOptions(urn=urn))
            else:
                raise Exception(f"unknown resource type {typ}")


    _module_instance = Module()
    pulumi.runtime.register_resource_module("gcp", "container/cluster", _module_instance)
    pulumi.runtime.register_resource_module("gcp", "container/nodePool", _module_instance)
    pulumi.runtime.register_resource_module("gcp", "container/registry", _module_instance)

_register_module()
