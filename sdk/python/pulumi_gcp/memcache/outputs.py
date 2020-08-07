# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = [
    'InstanceMemcacheParameters',
    'InstanceNodeConfig',
]

@pulumi.output_type
class InstanceMemcacheParameters(dict):
    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        """
        -
        This is a unique ID associated with this set of parameters.
        """
        ...

    @property
    @pulumi.getter
    def params(self) -> Optional[Mapping[str, str]]:
        """
        User-defined set of parameters to use in the memcache process.
        """
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class InstanceNodeConfig(dict):
    @property
    @pulumi.getter(name="cpuCount")
    def cpu_count(self) -> float:
        """
        Number of CPUs per node.
        """
        ...

    @property
    @pulumi.getter(name="memorySizeMb")
    def memory_size_mb(self) -> float:
        """
        Memory size in Mebibytes for each memcache node.
        """
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


