// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package sourcerepo

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// For more information, see [the official
// documentation](https://cloud.google.com/source-repositories/) and
// [API](https://cloud.google.com/source-repositories/docs/reference/rest/v1/projects.repos)
type Repository struct {
	s *pulumi.ResourceState
}

// NewRepository registers a new resource with the given unique name, arguments, and options.
func NewRepository(ctx *pulumi.Context,
	name string, args *RepositoryArgs, opts ...pulumi.ResourceOpt) (*Repository, error) {
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["name"] = nil
		inputs["project"] = nil
	} else {
		inputs["name"] = args.Name
		inputs["project"] = args.Project
	}
	inputs["size"] = nil
	inputs["url"] = nil
	s, err := ctx.RegisterResource("gcp:sourcerepo/repository:Repository", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Repository{s: s}, nil
}

// GetRepository gets an existing Repository resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRepository(ctx *pulumi.Context,
	name string, id pulumi.ID, state *RepositoryState, opts ...pulumi.ResourceOpt) (*Repository, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["name"] = state.Name
		inputs["project"] = state.Project
		inputs["size"] = state.Size
		inputs["url"] = state.Url
	}
	s, err := ctx.ReadResource("gcp:sourcerepo/repository:Repository", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Repository{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Repository) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Repository) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// The name of the repository that will be created.
func (r *Repository) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// The ID of the project in which the resource belongs. If it
// is not provided, the provider project is used.
func (r *Repository) Project() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["project"])
}

// The size of the repository.
func (r *Repository) Size() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["size"])
}

// The url to clone the repository.
func (r *Repository) Url() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["url"])
}

// Input properties used for looking up and filtering Repository resources.
type RepositoryState struct {
	// The name of the repository that will be created.
	Name interface{}
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project interface{}
	// The size of the repository.
	Size interface{}
	// The url to clone the repository.
	Url interface{}
}

// The set of arguments for constructing a Repository resource.
type RepositoryArgs struct {
	// The name of the repository that will be created.
	Name interface{}
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project interface{}
}
