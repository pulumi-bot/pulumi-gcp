// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package healthcare

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/healthcare_hl7_v2_store.html.markdown.
type Hl7Store struct {
	s *pulumi.ResourceState
}

// NewHl7Store registers a new resource with the given unique name, arguments, and options.
func NewHl7Store(ctx *pulumi.Context,
	name string, args *Hl7StoreArgs, opts ...pulumi.ResourceOpt) (*Hl7Store, error) {
	if args == nil || args.Dataset == nil {
		return nil, errors.New("missing required argument 'Dataset'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["dataset"] = nil
		inputs["labels"] = nil
		inputs["name"] = nil
		inputs["notificationConfig"] = nil
		inputs["parserConfig"] = nil
	} else {
		inputs["dataset"] = args.Dataset
		inputs["labels"] = args.Labels
		inputs["name"] = args.Name
		inputs["notificationConfig"] = args.NotificationConfig
		inputs["parserConfig"] = args.ParserConfig
	}
	inputs["selfLink"] = nil
	s, err := ctx.RegisterResource("gcp:healthcare/hl7Store:Hl7Store", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Hl7Store{s: s}, nil
}

// GetHl7Store gets an existing Hl7Store resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetHl7Store(ctx *pulumi.Context,
	name string, id pulumi.ID, state *Hl7StoreState, opts ...pulumi.ResourceOpt) (*Hl7Store, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["dataset"] = state.Dataset
		inputs["labels"] = state.Labels
		inputs["name"] = state.Name
		inputs["notificationConfig"] = state.NotificationConfig
		inputs["parserConfig"] = state.ParserConfig
		inputs["selfLink"] = state.SelfLink
	}
	s, err := ctx.ReadResource("gcp:healthcare/hl7Store:Hl7Store", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Hl7Store{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Hl7Store) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Hl7Store) ID() *pulumi.IDOutput {
	return r.s.ID()
}

func (r *Hl7Store) Dataset() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["dataset"])
}

func (r *Hl7Store) Labels() *pulumi.MapOutput {
	return (*pulumi.MapOutput)(r.s.State["labels"])
}

func (r *Hl7Store) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

func (r *Hl7Store) NotificationConfig() *pulumi.Output {
	return r.s.State["notificationConfig"]
}

func (r *Hl7Store) ParserConfig() *pulumi.Output {
	return r.s.State["parserConfig"]
}

func (r *Hl7Store) SelfLink() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["selfLink"])
}

// Input properties used for looking up and filtering Hl7Store resources.
type Hl7StoreState struct {
	Dataset interface{}
	Labels interface{}
	Name interface{}
	NotificationConfig interface{}
	ParserConfig interface{}
	SelfLink interface{}
}

// The set of arguments for constructing a Hl7Store resource.
type Hl7StoreArgs struct {
	Dataset interface{}
	Labels interface{}
	Name interface{}
	NotificationConfig interface{}
	ParserConfig interface{}
}
