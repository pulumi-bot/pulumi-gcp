// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package pubsub

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// A named resource representing the stream of messages from a single,
// specific topic, to be delivered to the subscribing application.
//
// To get more information about Subscription, see:
//
// * [API documentation](https://cloud.google.com/pubsub/docs/reference/rest/v1/projects.subscriptions)
// * How-to Guides
//     * [Managing Subscriptions](https://cloud.google.com/pubsub/docs/admin#managing_subscriptions)
//
// ## Example Usage
// ### Pubsub Lite Subscription Basic
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/organizations"
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/pubsub"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		project, err := organizations.LookupProject(ctx, nil, nil)
// 		if err != nil {
// 			return err
// 		}
// 		exampleLiteTopic, err := pubsub.NewLiteTopic(ctx, "exampleLiteTopic", &pubsub.LiteTopicArgs{
// 			Project: pulumi.String(project.Number),
// 			PartitionConfig: &pubsub.LiteTopicPartitionConfigArgs{
// 				Count: pulumi.Int(1),
// 				Capacity: &pubsub.LiteTopicPartitionConfigCapacityArgs{
// 					PublishMibPerSec:   pulumi.Int(4),
// 					SubscribeMibPerSec: pulumi.Int(8),
// 				},
// 			},
// 			RetentionConfig: &pubsub.LiteTopicRetentionConfigArgs{
// 				PerPartitionBytes: pulumi.String("32212254720"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = pubsub.NewLiteSubscription(ctx, "exampleLiteSubscription", &pubsub.LiteSubscriptionArgs{
// 			Topic: exampleLiteTopic.Name,
// 			DeliveryConfig: &pubsub.LiteSubscriptionDeliveryConfigArgs{
// 				DeliveryRequirement: pulumi.String("DELIVER_AFTER_STORED"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// Subscription can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import gcp:pubsub/liteSubscription:LiteSubscription default projects/{{project}}/locations/{{zone}}/subscriptions/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:pubsub/liteSubscription:LiteSubscription default {{project}}/{{zone}}/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:pubsub/liteSubscription:LiteSubscription default {{zone}}/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:pubsub/liteSubscription:LiteSubscription default {{name}}
// ```
type LiteSubscription struct {
	pulumi.CustomResourceState

	// The settings for this subscription's message delivery.
	// Structure is documented below.
	DeliveryConfig LiteSubscriptionDeliveryConfigPtrOutput `pulumi:"deliveryConfig"`
	// Name of the subscription.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The region of the pubsub lite topic.
	Region pulumi.StringPtrOutput `pulumi:"region"`
	// A reference to a Topic resource.
	Topic pulumi.StringOutput `pulumi:"topic"`
	// The zone of the pubsub lite topic.
	Zone pulumi.StringPtrOutput `pulumi:"zone"`
}

// NewLiteSubscription registers a new resource with the given unique name, arguments, and options.
func NewLiteSubscription(ctx *pulumi.Context,
	name string, args *LiteSubscriptionArgs, opts ...pulumi.ResourceOption) (*LiteSubscription, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Topic == nil {
		return nil, errors.New("invalid value for required argument 'Topic'")
	}
	var resource LiteSubscription
	err := ctx.RegisterResource("gcp:pubsub/liteSubscription:LiteSubscription", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLiteSubscription gets an existing LiteSubscription resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLiteSubscription(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LiteSubscriptionState, opts ...pulumi.ResourceOption) (*LiteSubscription, error) {
	var resource LiteSubscription
	err := ctx.ReadResource("gcp:pubsub/liteSubscription:LiteSubscription", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LiteSubscription resources.
type liteSubscriptionState struct {
	// The settings for this subscription's message delivery.
	// Structure is documented below.
	DeliveryConfig *LiteSubscriptionDeliveryConfig `pulumi:"deliveryConfig"`
	// Name of the subscription.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The region of the pubsub lite topic.
	Region *string `pulumi:"region"`
	// A reference to a Topic resource.
	Topic *string `pulumi:"topic"`
	// The zone of the pubsub lite topic.
	Zone *string `pulumi:"zone"`
}

type LiteSubscriptionState struct {
	// The settings for this subscription's message delivery.
	// Structure is documented below.
	DeliveryConfig LiteSubscriptionDeliveryConfigPtrInput
	// Name of the subscription.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The region of the pubsub lite topic.
	Region pulumi.StringPtrInput
	// A reference to a Topic resource.
	Topic pulumi.StringPtrInput
	// The zone of the pubsub lite topic.
	Zone pulumi.StringPtrInput
}

func (LiteSubscriptionState) ElementType() reflect.Type {
	return reflect.TypeOf((*liteSubscriptionState)(nil)).Elem()
}

type liteSubscriptionArgs struct {
	// The settings for this subscription's message delivery.
	// Structure is documented below.
	DeliveryConfig *LiteSubscriptionDeliveryConfig `pulumi:"deliveryConfig"`
	// Name of the subscription.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The region of the pubsub lite topic.
	Region *string `pulumi:"region"`
	// A reference to a Topic resource.
	Topic string `pulumi:"topic"`
	// The zone of the pubsub lite topic.
	Zone *string `pulumi:"zone"`
}

// The set of arguments for constructing a LiteSubscription resource.
type LiteSubscriptionArgs struct {
	// The settings for this subscription's message delivery.
	// Structure is documented below.
	DeliveryConfig LiteSubscriptionDeliveryConfigPtrInput
	// Name of the subscription.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The region of the pubsub lite topic.
	Region pulumi.StringPtrInput
	// A reference to a Topic resource.
	Topic pulumi.StringInput
	// The zone of the pubsub lite topic.
	Zone pulumi.StringPtrInput
}

func (LiteSubscriptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*liteSubscriptionArgs)(nil)).Elem()
}

type LiteSubscriptionInput interface {
	pulumi.Input

	ToLiteSubscriptionOutput() LiteSubscriptionOutput
	ToLiteSubscriptionOutputWithContext(ctx context.Context) LiteSubscriptionOutput
}

func (*LiteSubscription) ElementType() reflect.Type {
	return reflect.TypeOf((*LiteSubscription)(nil))
}

func (i *LiteSubscription) ToLiteSubscriptionOutput() LiteSubscriptionOutput {
	return i.ToLiteSubscriptionOutputWithContext(context.Background())
}

func (i *LiteSubscription) ToLiteSubscriptionOutputWithContext(ctx context.Context) LiteSubscriptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LiteSubscriptionOutput)
}

type LiteSubscriptionOutput struct {
	*pulumi.OutputState
}

func (LiteSubscriptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LiteSubscription)(nil))
}

func (o LiteSubscriptionOutput) ToLiteSubscriptionOutput() LiteSubscriptionOutput {
	return o
}

func (o LiteSubscriptionOutput) ToLiteSubscriptionOutputWithContext(ctx context.Context) LiteSubscriptionOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(LiteSubscriptionOutput{})
}
