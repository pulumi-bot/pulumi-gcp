// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package diagflow

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// A Dialogflow agent is a virtual agent that handles conversations with your end-users. It is a natural language
// understanding module that understands the nuances of human language. Dialogflow translates end-user text or audio
// during a conversation to structured data that your apps and services can understand. You design and build a Dialogflow
// agent to handle the types of conversations required for your system.
//
// To get more information about Agent, see:
//
// * [API documentation](https://cloud.google.com/dialogflow/docs/reference/rest/v2/projects/agent)
// * How-to Guides
//     * [Official Documentation](https://cloud.google.com/dialogflow/docs/)
//
// ## Example Usage
// ### Dialogflow Agent Full
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/diagflow"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := diagflow.NewAgent(ctx, "fullAgent", &diagflow.AgentArgs{
// 			ApiVersion:              pulumi.String("API_VERSION_V2_BETA_1"),
// 			AvatarUri:               pulumi.String("https://cloud.google.com/_static/images/cloud/icons/favicons/onecloud/super_cloud.png"),
// 			ClassificationThreshold: pulumi.Float64(0.3),
// 			DefaultLanguageCode:     pulumi.String("en"),
// 			Description:             pulumi.String("Example description."),
// 			DisplayName:             pulumi.String("dialogflow-agent"),
// 			EnableLogging:           pulumi.Bool(true),
// 			MatchMode:               pulumi.String("MATCH_MODE_ML_ONLY"),
// 			SupportedLanguageCodes: pulumi.StringArray{
// 				pulumi.String("fr"),
// 				pulumi.String("de"),
// 				pulumi.String("es"),
// 			},
// 			Tier:     pulumi.String("TIER_STANDARD"),
// 			TimeZone: pulumi.String("America/New_York"),
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
// Agent can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import gcp:diagflow/agent:Agent default {{project}}
// ```
type Agent struct {
	pulumi.CustomResourceState

	// API version displayed in Dialogflow console. If not specified, V2 API is assumed. Clients are free to query
	// different service endpoints for different API versions. However, bots connectors and webhook calls will follow
	// the specified API version.
	// * API_VERSION_V1: Legacy V1 API.
	// * API_VERSION_V2: V2 API.
	// * API_VERSION_V2_BETA_1: V2beta1 API.
	//   Possible values are `API_VERSION_V1`, `API_VERSION_V2`, and `API_VERSION_V2_BETA_1`.
	ApiVersion pulumi.StringOutput `pulumi:"apiVersion"`
	// The URI of the agent's avatar, which are used throughout the Dialogflow console. When an image URL is entered
	// into this field, the Dialogflow will save the image in the backend. The address of the backend image returned
	// from the API will be shown in the [avatarUriBackend] field.
	AvatarUri pulumi.StringPtrOutput `pulumi:"avatarUri"`
	// The URI of the agent's avatar as returned from the API. Output only. To provide an image URL for the agent avatar, the
	// [avatarUri] field can be used.
	AvatarUriBackend pulumi.StringOutput `pulumi:"avatarUriBackend"`
	// To filter out false positive results and still get variety in matched natural language inputs for your agent,
	// you can tune the machine learning classification threshold. If the returned score value is less than the threshold
	// value, then a fallback intent will be triggered or, if there are no fallback intents defined, no intent will be
	// triggered. The score values range from 0.0 (completely uncertain) to 1.0 (completely certain). If set to 0.0, the
	// default of 0.3 is used.
	ClassificationThreshold pulumi.Float64PtrOutput `pulumi:"classificationThreshold"`
	// The default language of the agent as a language tag. [See Language Support](https://cloud.google.com/dialogflow/docs/reference/language)
	// for a list of the currently supported language codes. This field cannot be updated after creation.
	DefaultLanguageCode pulumi.StringOutput `pulumi:"defaultLanguageCode"`
	// The description of this agent. The maximum length is 500 characters. If exceeded, the request is rejected.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The name of this agent.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// Determines whether this agent should log conversation queries.
	EnableLogging pulumi.BoolPtrOutput `pulumi:"enableLogging"`
	// Determines how intents are detected from user queries.
	// * MATCH_MODE_HYBRID: Best for agents with a small number of examples in intents and/or wide use of templates
	//   syntax and composite entities.
	// * MATCH_MODE_ML_ONLY: Can be used for agents with a large number of examples in intents, especially the ones
	//   using @sys.any or very large developer entities.
	//   Possible values are `MATCH_MODE_HYBRID` and `MATCH_MODE_ML_ONLY`.
	MatchMode pulumi.StringOutput `pulumi:"matchMode"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The list of all languages supported by this agent (except for the defaultLanguageCode).
	SupportedLanguageCodes pulumi.StringArrayOutput `pulumi:"supportedLanguageCodes"`
	// The agent tier. If not specified, TIER_STANDARD is assumed.
	// * TIER_STANDARD: Standard tier.
	// * TIER_ENTERPRISE: Enterprise tier (Essentials).
	// * TIER_ENTERPRISE_PLUS: Enterprise tier (Plus).
	//   NOTE: Due to consistency issues, the provider will not read this field from the API. Drift is possible between
	//   the the provider state and Dialogflow if the agent tier is changed outside of the provider.
	Tier pulumi.StringPtrOutput `pulumi:"tier"`
	// The time zone of this agent from the [time zone database](https://www.iana.org/time-zones), e.g., America/New_York,
	// Europe/Paris.
	TimeZone pulumi.StringOutput `pulumi:"timeZone"`
}

// NewAgent registers a new resource with the given unique name, arguments, and options.
func NewAgent(ctx *pulumi.Context,
	name string, args *AgentArgs, opts ...pulumi.ResourceOption) (*Agent, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DefaultLanguageCode == nil {
		return nil, errors.New("invalid value for required argument 'DefaultLanguageCode'")
	}
	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	if args.TimeZone == nil {
		return nil, errors.New("invalid value for required argument 'TimeZone'")
	}
	var resource Agent
	err := ctx.RegisterResource("gcp:diagflow/agent:Agent", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAgent gets an existing Agent resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAgent(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AgentState, opts ...pulumi.ResourceOption) (*Agent, error) {
	var resource Agent
	err := ctx.ReadResource("gcp:diagflow/agent:Agent", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Agent resources.
type agentState struct {
	// API version displayed in Dialogflow console. If not specified, V2 API is assumed. Clients are free to query
	// different service endpoints for different API versions. However, bots connectors and webhook calls will follow
	// the specified API version.
	// * API_VERSION_V1: Legacy V1 API.
	// * API_VERSION_V2: V2 API.
	// * API_VERSION_V2_BETA_1: V2beta1 API.
	//   Possible values are `API_VERSION_V1`, `API_VERSION_V2`, and `API_VERSION_V2_BETA_1`.
	ApiVersion *string `pulumi:"apiVersion"`
	// The URI of the agent's avatar, which are used throughout the Dialogflow console. When an image URL is entered
	// into this field, the Dialogflow will save the image in the backend. The address of the backend image returned
	// from the API will be shown in the [avatarUriBackend] field.
	AvatarUri *string `pulumi:"avatarUri"`
	// The URI of the agent's avatar as returned from the API. Output only. To provide an image URL for the agent avatar, the
	// [avatarUri] field can be used.
	AvatarUriBackend *string `pulumi:"avatarUriBackend"`
	// To filter out false positive results and still get variety in matched natural language inputs for your agent,
	// you can tune the machine learning classification threshold. If the returned score value is less than the threshold
	// value, then a fallback intent will be triggered or, if there are no fallback intents defined, no intent will be
	// triggered. The score values range from 0.0 (completely uncertain) to 1.0 (completely certain). If set to 0.0, the
	// default of 0.3 is used.
	ClassificationThreshold *float64 `pulumi:"classificationThreshold"`
	// The default language of the agent as a language tag. [See Language Support](https://cloud.google.com/dialogflow/docs/reference/language)
	// for a list of the currently supported language codes. This field cannot be updated after creation.
	DefaultLanguageCode *string `pulumi:"defaultLanguageCode"`
	// The description of this agent. The maximum length is 500 characters. If exceeded, the request is rejected.
	Description *string `pulumi:"description"`
	// The name of this agent.
	DisplayName *string `pulumi:"displayName"`
	// Determines whether this agent should log conversation queries.
	EnableLogging *bool `pulumi:"enableLogging"`
	// Determines how intents are detected from user queries.
	// * MATCH_MODE_HYBRID: Best for agents with a small number of examples in intents and/or wide use of templates
	//   syntax and composite entities.
	// * MATCH_MODE_ML_ONLY: Can be used for agents with a large number of examples in intents, especially the ones
	//   using @sys.any or very large developer entities.
	//   Possible values are `MATCH_MODE_HYBRID` and `MATCH_MODE_ML_ONLY`.
	MatchMode *string `pulumi:"matchMode"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The list of all languages supported by this agent (except for the defaultLanguageCode).
	SupportedLanguageCodes []string `pulumi:"supportedLanguageCodes"`
	// The agent tier. If not specified, TIER_STANDARD is assumed.
	// * TIER_STANDARD: Standard tier.
	// * TIER_ENTERPRISE: Enterprise tier (Essentials).
	// * TIER_ENTERPRISE_PLUS: Enterprise tier (Plus).
	//   NOTE: Due to consistency issues, the provider will not read this field from the API. Drift is possible between
	//   the the provider state and Dialogflow if the agent tier is changed outside of the provider.
	Tier *string `pulumi:"tier"`
	// The time zone of this agent from the [time zone database](https://www.iana.org/time-zones), e.g., America/New_York,
	// Europe/Paris.
	TimeZone *string `pulumi:"timeZone"`
}

type AgentState struct {
	// API version displayed in Dialogflow console. If not specified, V2 API is assumed. Clients are free to query
	// different service endpoints for different API versions. However, bots connectors and webhook calls will follow
	// the specified API version.
	// * API_VERSION_V1: Legacy V1 API.
	// * API_VERSION_V2: V2 API.
	// * API_VERSION_V2_BETA_1: V2beta1 API.
	//   Possible values are `API_VERSION_V1`, `API_VERSION_V2`, and `API_VERSION_V2_BETA_1`.
	ApiVersion pulumi.StringPtrInput
	// The URI of the agent's avatar, which are used throughout the Dialogflow console. When an image URL is entered
	// into this field, the Dialogflow will save the image in the backend. The address of the backend image returned
	// from the API will be shown in the [avatarUriBackend] field.
	AvatarUri pulumi.StringPtrInput
	// The URI of the agent's avatar as returned from the API. Output only. To provide an image URL for the agent avatar, the
	// [avatarUri] field can be used.
	AvatarUriBackend pulumi.StringPtrInput
	// To filter out false positive results and still get variety in matched natural language inputs for your agent,
	// you can tune the machine learning classification threshold. If the returned score value is less than the threshold
	// value, then a fallback intent will be triggered or, if there are no fallback intents defined, no intent will be
	// triggered. The score values range from 0.0 (completely uncertain) to 1.0 (completely certain). If set to 0.0, the
	// default of 0.3 is used.
	ClassificationThreshold pulumi.Float64PtrInput
	// The default language of the agent as a language tag. [See Language Support](https://cloud.google.com/dialogflow/docs/reference/language)
	// for a list of the currently supported language codes. This field cannot be updated after creation.
	DefaultLanguageCode pulumi.StringPtrInput
	// The description of this agent. The maximum length is 500 characters. If exceeded, the request is rejected.
	Description pulumi.StringPtrInput
	// The name of this agent.
	DisplayName pulumi.StringPtrInput
	// Determines whether this agent should log conversation queries.
	EnableLogging pulumi.BoolPtrInput
	// Determines how intents are detected from user queries.
	// * MATCH_MODE_HYBRID: Best for agents with a small number of examples in intents and/or wide use of templates
	//   syntax and composite entities.
	// * MATCH_MODE_ML_ONLY: Can be used for agents with a large number of examples in intents, especially the ones
	//   using @sys.any or very large developer entities.
	//   Possible values are `MATCH_MODE_HYBRID` and `MATCH_MODE_ML_ONLY`.
	MatchMode pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The list of all languages supported by this agent (except for the defaultLanguageCode).
	SupportedLanguageCodes pulumi.StringArrayInput
	// The agent tier. If not specified, TIER_STANDARD is assumed.
	// * TIER_STANDARD: Standard tier.
	// * TIER_ENTERPRISE: Enterprise tier (Essentials).
	// * TIER_ENTERPRISE_PLUS: Enterprise tier (Plus).
	//   NOTE: Due to consistency issues, the provider will not read this field from the API. Drift is possible between
	//   the the provider state and Dialogflow if the agent tier is changed outside of the provider.
	Tier pulumi.StringPtrInput
	// The time zone of this agent from the [time zone database](https://www.iana.org/time-zones), e.g., America/New_York,
	// Europe/Paris.
	TimeZone pulumi.StringPtrInput
}

func (AgentState) ElementType() reflect.Type {
	return reflect.TypeOf((*agentState)(nil)).Elem()
}

type agentArgs struct {
	// API version displayed in Dialogflow console. If not specified, V2 API is assumed. Clients are free to query
	// different service endpoints for different API versions. However, bots connectors and webhook calls will follow
	// the specified API version.
	// * API_VERSION_V1: Legacy V1 API.
	// * API_VERSION_V2: V2 API.
	// * API_VERSION_V2_BETA_1: V2beta1 API.
	//   Possible values are `API_VERSION_V1`, `API_VERSION_V2`, and `API_VERSION_V2_BETA_1`.
	ApiVersion *string `pulumi:"apiVersion"`
	// The URI of the agent's avatar, which are used throughout the Dialogflow console. When an image URL is entered
	// into this field, the Dialogflow will save the image in the backend. The address of the backend image returned
	// from the API will be shown in the [avatarUriBackend] field.
	AvatarUri *string `pulumi:"avatarUri"`
	// To filter out false positive results and still get variety in matched natural language inputs for your agent,
	// you can tune the machine learning classification threshold. If the returned score value is less than the threshold
	// value, then a fallback intent will be triggered or, if there are no fallback intents defined, no intent will be
	// triggered. The score values range from 0.0 (completely uncertain) to 1.0 (completely certain). If set to 0.0, the
	// default of 0.3 is used.
	ClassificationThreshold *float64 `pulumi:"classificationThreshold"`
	// The default language of the agent as a language tag. [See Language Support](https://cloud.google.com/dialogflow/docs/reference/language)
	// for a list of the currently supported language codes. This field cannot be updated after creation.
	DefaultLanguageCode string `pulumi:"defaultLanguageCode"`
	// The description of this agent. The maximum length is 500 characters. If exceeded, the request is rejected.
	Description *string `pulumi:"description"`
	// The name of this agent.
	DisplayName string `pulumi:"displayName"`
	// Determines whether this agent should log conversation queries.
	EnableLogging *bool `pulumi:"enableLogging"`
	// Determines how intents are detected from user queries.
	// * MATCH_MODE_HYBRID: Best for agents with a small number of examples in intents and/or wide use of templates
	//   syntax and composite entities.
	// * MATCH_MODE_ML_ONLY: Can be used for agents with a large number of examples in intents, especially the ones
	//   using @sys.any or very large developer entities.
	//   Possible values are `MATCH_MODE_HYBRID` and `MATCH_MODE_ML_ONLY`.
	MatchMode *string `pulumi:"matchMode"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The list of all languages supported by this agent (except for the defaultLanguageCode).
	SupportedLanguageCodes []string `pulumi:"supportedLanguageCodes"`
	// The agent tier. If not specified, TIER_STANDARD is assumed.
	// * TIER_STANDARD: Standard tier.
	// * TIER_ENTERPRISE: Enterprise tier (Essentials).
	// * TIER_ENTERPRISE_PLUS: Enterprise tier (Plus).
	//   NOTE: Due to consistency issues, the provider will not read this field from the API. Drift is possible between
	//   the the provider state and Dialogflow if the agent tier is changed outside of the provider.
	Tier *string `pulumi:"tier"`
	// The time zone of this agent from the [time zone database](https://www.iana.org/time-zones), e.g., America/New_York,
	// Europe/Paris.
	TimeZone string `pulumi:"timeZone"`
}

// The set of arguments for constructing a Agent resource.
type AgentArgs struct {
	// API version displayed in Dialogflow console. If not specified, V2 API is assumed. Clients are free to query
	// different service endpoints for different API versions. However, bots connectors and webhook calls will follow
	// the specified API version.
	// * API_VERSION_V1: Legacy V1 API.
	// * API_VERSION_V2: V2 API.
	// * API_VERSION_V2_BETA_1: V2beta1 API.
	//   Possible values are `API_VERSION_V1`, `API_VERSION_V2`, and `API_VERSION_V2_BETA_1`.
	ApiVersion pulumi.StringPtrInput
	// The URI of the agent's avatar, which are used throughout the Dialogflow console. When an image URL is entered
	// into this field, the Dialogflow will save the image in the backend. The address of the backend image returned
	// from the API will be shown in the [avatarUriBackend] field.
	AvatarUri pulumi.StringPtrInput
	// To filter out false positive results and still get variety in matched natural language inputs for your agent,
	// you can tune the machine learning classification threshold. If the returned score value is less than the threshold
	// value, then a fallback intent will be triggered or, if there are no fallback intents defined, no intent will be
	// triggered. The score values range from 0.0 (completely uncertain) to 1.0 (completely certain). If set to 0.0, the
	// default of 0.3 is used.
	ClassificationThreshold pulumi.Float64PtrInput
	// The default language of the agent as a language tag. [See Language Support](https://cloud.google.com/dialogflow/docs/reference/language)
	// for a list of the currently supported language codes. This field cannot be updated after creation.
	DefaultLanguageCode pulumi.StringInput
	// The description of this agent. The maximum length is 500 characters. If exceeded, the request is rejected.
	Description pulumi.StringPtrInput
	// The name of this agent.
	DisplayName pulumi.StringInput
	// Determines whether this agent should log conversation queries.
	EnableLogging pulumi.BoolPtrInput
	// Determines how intents are detected from user queries.
	// * MATCH_MODE_HYBRID: Best for agents with a small number of examples in intents and/or wide use of templates
	//   syntax and composite entities.
	// * MATCH_MODE_ML_ONLY: Can be used for agents with a large number of examples in intents, especially the ones
	//   using @sys.any or very large developer entities.
	//   Possible values are `MATCH_MODE_HYBRID` and `MATCH_MODE_ML_ONLY`.
	MatchMode pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The list of all languages supported by this agent (except for the defaultLanguageCode).
	SupportedLanguageCodes pulumi.StringArrayInput
	// The agent tier. If not specified, TIER_STANDARD is assumed.
	// * TIER_STANDARD: Standard tier.
	// * TIER_ENTERPRISE: Enterprise tier (Essentials).
	// * TIER_ENTERPRISE_PLUS: Enterprise tier (Plus).
	//   NOTE: Due to consistency issues, the provider will not read this field from the API. Drift is possible between
	//   the the provider state and Dialogflow if the agent tier is changed outside of the provider.
	Tier pulumi.StringPtrInput
	// The time zone of this agent from the [time zone database](https://www.iana.org/time-zones), e.g., America/New_York,
	// Europe/Paris.
	TimeZone pulumi.StringInput
}

func (AgentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*agentArgs)(nil)).Elem()
}

type AgentInput interface {
	pulumi.Input

	ToAgentOutput() AgentOutput
	ToAgentOutputWithContext(ctx context.Context) AgentOutput
}

func (*Agent) ElementType() reflect.Type {
	return reflect.TypeOf((*Agent)(nil))
}

func (i *Agent) ToAgentOutput() AgentOutput {
	return i.ToAgentOutputWithContext(context.Background())
}

func (i *Agent) ToAgentOutputWithContext(ctx context.Context) AgentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AgentOutput)
}

func (i *Agent) ToAgentPtrOutput() AgentPtrOutput {
	return i.ToAgentPtrOutputWithContext(context.Background())
}

func (i *Agent) ToAgentPtrOutputWithContext(ctx context.Context) AgentPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AgentPtrOutput)
}

type AgentPtrInput interface {
	pulumi.Input

	ToAgentPtrOutput() AgentPtrOutput
	ToAgentPtrOutputWithContext(ctx context.Context) AgentPtrOutput
}

type agentPtrType AgentArgs

func (*agentPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Agent)(nil))
}

func (i *agentPtrType) ToAgentPtrOutput() AgentPtrOutput {
	return i.ToAgentPtrOutputWithContext(context.Background())
}

func (i *agentPtrType) ToAgentPtrOutputWithContext(ctx context.Context) AgentPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AgentOutput).ToAgentPtrOutput()
}

// AgentArrayInput is an input type that accepts AgentArray and AgentArrayOutput values.
// You can construct a concrete instance of `AgentArrayInput` via:
//
//          AgentArray{ AgentArgs{...} }
type AgentArrayInput interface {
	pulumi.Input

	ToAgentArrayOutput() AgentArrayOutput
	ToAgentArrayOutputWithContext(context.Context) AgentArrayOutput
}

type AgentArray []AgentInput

func (AgentArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Agent)(nil))
}

func (i AgentArray) ToAgentArrayOutput() AgentArrayOutput {
	return i.ToAgentArrayOutputWithContext(context.Background())
}

func (i AgentArray) ToAgentArrayOutputWithContext(ctx context.Context) AgentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AgentArrayOutput)
}

// AgentMapInput is an input type that accepts AgentMap and AgentMapOutput values.
// You can construct a concrete instance of `AgentMapInput` via:
//
//          AgentMap{ "key": AgentArgs{...} }
type AgentMapInput interface {
	pulumi.Input

	ToAgentMapOutput() AgentMapOutput
	ToAgentMapOutputWithContext(context.Context) AgentMapOutput
}

type AgentMap map[string]AgentInput

func (AgentMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Agent)(nil))
}

func (i AgentMap) ToAgentMapOutput() AgentMapOutput {
	return i.ToAgentMapOutputWithContext(context.Background())
}

func (i AgentMap) ToAgentMapOutputWithContext(ctx context.Context) AgentMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AgentMapOutput)
}

type AgentOutput struct {
	*pulumi.OutputState
}

func (AgentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Agent)(nil))
}

func (o AgentOutput) ToAgentOutput() AgentOutput {
	return o
}

func (o AgentOutput) ToAgentOutputWithContext(ctx context.Context) AgentOutput {
	return o
}

func (o AgentOutput) ToAgentPtrOutput() AgentPtrOutput {
	return o.ToAgentPtrOutputWithContext(context.Background())
}

func (o AgentOutput) ToAgentPtrOutputWithContext(ctx context.Context) AgentPtrOutput {
	return o.ApplyT(func(v Agent) *Agent {
		return &v
	}).(AgentPtrOutput)
}

type AgentPtrOutput struct {
	*pulumi.OutputState
}

func (AgentPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Agent)(nil))
}

func (o AgentPtrOutput) ToAgentPtrOutput() AgentPtrOutput {
	return o
}

func (o AgentPtrOutput) ToAgentPtrOutputWithContext(ctx context.Context) AgentPtrOutput {
	return o
}

type AgentArrayOutput struct{ *pulumi.OutputState }

func (AgentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Agent)(nil))
}

func (o AgentArrayOutput) ToAgentArrayOutput() AgentArrayOutput {
	return o
}

func (o AgentArrayOutput) ToAgentArrayOutputWithContext(ctx context.Context) AgentArrayOutput {
	return o
}

func (o AgentArrayOutput) Index(i pulumi.IntInput) AgentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Agent {
		return vs[0].([]Agent)[vs[1].(int)]
	}).(AgentOutput)
}

type AgentMapOutput struct{ *pulumi.OutputState }

func (AgentMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Agent)(nil))
}

func (o AgentMapOutput) ToAgentMapOutput() AgentMapOutput {
	return o
}

func (o AgentMapOutput) ToAgentMapOutputWithContext(ctx context.Context) AgentMapOutput {
	return o
}

func (o AgentMapOutput) MapIndex(k pulumi.StringInput) AgentOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Agent {
		return vs[0].(map[string]Agent)[vs[1].(string)]
	}).(AgentOutput)
}

func init() {
	pulumi.RegisterOutputType(AgentOutput{})
	pulumi.RegisterOutputType(AgentPtrOutput{})
	pulumi.RegisterOutputType(AgentArrayOutput{})
	pulumi.RegisterOutputType(AgentMapOutput{})
}
