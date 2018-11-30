// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package binaryauthorization

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// #     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
// #
// # ----------------------------------------------------------------------------
// #
// #     This file is automatically generated by Magic Modules and manual
// #     changes will be clobbered when the file is regenerated.
// #
// #     Please read more about how to change this file in
// #     .github/CONTRIBUTING.md.
// #
// # ----------------------------------------------------------------------------
// layout: "google"
// page_title: "Google: google_binary_authorization_attestor"
// sidebar_current: "docs-google-binary-authorization-attestor"
// description: |-
//   An attestor that attests to container image artifacts.
// ---
// 
// # google\_binary\_authorization\_attestor
// 
// An attestor that attests to container image artifacts.
// 
// ~> **Warning:** This resource is in beta, and should be used with the terraform-provider-google-beta provider.
// See [Provider Versions](https://terraform.io/docs/providers/google/provider_versions.html) for more details on beta resources.
// 
// To get more information about Attestor, see:
// 
// * [API documentation](https://cloud.google.com/binary-authorization/docs/reference/rest/)
// * How-to Guides
//     * [Official Documentation](https://cloud.google.com/binary-authorization/)
type Attestor struct {
	s *pulumi.ResourceState
}

// NewAttestor registers a new resource with the given unique name, arguments, and options.
func NewAttestor(ctx *pulumi.Context,
	name string, args *AttestorArgs, opts ...pulumi.ResourceOpt) (*Attestor, error) {
	if args == nil || args.AttestationAuthorityNote == nil {
		return nil, errors.New("missing required argument 'AttestationAuthorityNote'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["attestationAuthorityNote"] = nil
		inputs["description"] = nil
		inputs["name"] = nil
		inputs["project"] = nil
	} else {
		inputs["attestationAuthorityNote"] = args.AttestationAuthorityNote
		inputs["description"] = args.Description
		inputs["name"] = args.Name
		inputs["project"] = args.Project
	}
	s, err := ctx.RegisterResource("gcp:binaryauthorization/attestor:Attestor", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Attestor{s: s}, nil
}

// GetAttestor gets an existing Attestor resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAttestor(ctx *pulumi.Context,
	name string, id pulumi.ID, state *AttestorState, opts ...pulumi.ResourceOpt) (*Attestor, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["attestationAuthorityNote"] = state.AttestationAuthorityNote
		inputs["description"] = state.Description
		inputs["name"] = state.Name
		inputs["project"] = state.Project
	}
	s, err := ctx.ReadResource("gcp:binaryauthorization/attestor:Attestor", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Attestor{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Attestor) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Attestor) ID() *pulumi.IDOutput {
	return r.s.ID()
}

func (r *Attestor) AttestationAuthorityNote() *pulumi.Output {
	return r.s.State["attestationAuthorityNote"]
}

func (r *Attestor) Description() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["description"])
}

func (r *Attestor) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

func (r *Attestor) Project() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["project"])
}

// Input properties used for looking up and filtering Attestor resources.
type AttestorState struct {
	AttestationAuthorityNote interface{}
	Description interface{}
	Name interface{}
	Project interface{}
}

// The set of arguments for constructing a Attestor resource.
type AttestorArgs struct {
	AttestationAuthorityNote interface{}
	Description interface{}
	Name interface{}
	Project interface{}
}
