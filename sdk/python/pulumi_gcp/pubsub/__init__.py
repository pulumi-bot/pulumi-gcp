# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

# Export this package's modules as members:
from .get_topic import *
from .subscription import *
from .subscription_iam_binding import *
from .subscription_iam_member import *
from .subscription_iam_policy import *
from .topic import *
from .topic_iam_binding import *
from .topic_iam_member import *
from .topic_iam_policy import *
from ._inputs import *
from . import outputs

def _register_module():
    import pulumi

    class Module(pulumi.runtime.ResourceModule):
        def version(self):
            return None

        def construct(self, name: str, typ: str, urn: str) -> pulumi.Resource:
            if typ == "gcp:pubsub/subscription:Subscription":
                return Subscription(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "gcp:pubsub/subscriptionIAMBinding:SubscriptionIAMBinding":
                return SubscriptionIAMBinding(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "gcp:pubsub/subscriptionIAMMember:SubscriptionIAMMember":
                return SubscriptionIAMMember(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "gcp:pubsub/subscriptionIAMPolicy:SubscriptionIAMPolicy":
                return SubscriptionIAMPolicy(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "gcp:pubsub/topic:Topic":
                return Topic(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "gcp:pubsub/topicIAMBinding:TopicIAMBinding":
                return TopicIAMBinding(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "gcp:pubsub/topicIAMMember:TopicIAMMember":
                return TopicIAMMember(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "gcp:pubsub/topicIAMPolicy:TopicIAMPolicy":
                return TopicIAMPolicy(name, pulumi.ResourceOptions(urn=urn))
            else:
                raise Exception(f"unknown resource type {typ}")


    _module_instance = Module()
    pulumi.runtime.register_resource_module("gcp", "pubsub/subscription", _module_instance)
    pulumi.runtime.register_resource_module("gcp", "pubsub/subscriptionIAMBinding", _module_instance)
    pulumi.runtime.register_resource_module("gcp", "pubsub/subscriptionIAMMember", _module_instance)
    pulumi.runtime.register_resource_module("gcp", "pubsub/subscriptionIAMPolicy", _module_instance)
    pulumi.runtime.register_resource_module("gcp", "pubsub/topic", _module_instance)
    pulumi.runtime.register_resource_module("gcp", "pubsub/topicIAMBinding", _module_instance)
    pulumi.runtime.register_resource_module("gcp", "pubsub/topicIAMMember", _module_instance)
    pulumi.runtime.register_resource_module("gcp", "pubsub/topicIAMPolicy", _module_instance)

_register_module()
