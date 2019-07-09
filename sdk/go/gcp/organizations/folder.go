// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package organizations

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type Folder struct {
	s *pulumi.ResourceState
}

// NewFolder registers a new resource with the given unique name, arguments, and options.
func NewFolder(ctx *pulumi.Context,
	name string, args *FolderArgs, opts ...pulumi.ResourceOpt) (*Folder, error) {
	if args == nil || args.DisplayName == nil {
		return nil, errors.New("missing required argument 'DisplayName'")
	}
	if args == nil || args.Parent == nil {
		return nil, errors.New("missing required argument 'Parent'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["displayName"] = nil
		inputs["parent"] = nil
	} else {
		inputs["displayName"] = args.DisplayName
		inputs["parent"] = args.Parent
	}
	inputs["createTime"] = nil
	inputs["lifecycleState"] = nil
	inputs["name"] = nil
	s, err := ctx.RegisterResource("gcp:organizations/folder:Folder", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Folder{s: s}, nil
}

// GetFolder gets an existing Folder resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFolder(ctx *pulumi.Context,
	name string, id pulumi.ID, state *FolderState, opts ...pulumi.ResourceOpt) (*Folder, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["createTime"] = state.CreateTime
		inputs["displayName"] = state.DisplayName
		inputs["lifecycleState"] = state.LifecycleState
		inputs["name"] = state.Name
		inputs["parent"] = state.Parent
	}
	s, err := ctx.ReadResource("gcp:organizations/folder:Folder", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Folder{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Folder) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Folder) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// Timestamp when the Folder was created. Assigned by the server.
// A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".
func (r *Folder) CreateTime() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["createTime"])
}

// The folder’s display name.
// A folder’s display name must be unique amongst its siblings, e.g. no two folders with the same parent can share the same display name. The display name must start and end with a letter or digit, may contain letters, digits, spaces, hyphens and underscores and can be no longer than 30 characters.
func (r *Folder) DisplayName() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["displayName"])
}

// The lifecycle state of the folder such as `ACTIVE` or `DELETE_REQUESTED`.
func (r *Folder) LifecycleState() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["lifecycleState"])
}

// The resource name of the Folder. Its format is folders/{folder_id}.
func (r *Folder) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// The resource name of the parent Folder or Organization.
// Must be of the form `folders/{folder_id}` or `organizations/{org_id}`.
func (r *Folder) Parent() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["parent"])
}

// Input properties used for looking up and filtering Folder resources.
type FolderState struct {
	// Timestamp when the Folder was created. Assigned by the server.
	// A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".
	CreateTime interface{}
	// The folder’s display name.
	// A folder’s display name must be unique amongst its siblings, e.g. no two folders with the same parent can share the same display name. The display name must start and end with a letter or digit, may contain letters, digits, spaces, hyphens and underscores and can be no longer than 30 characters.
	DisplayName interface{}
	// The lifecycle state of the folder such as `ACTIVE` or `DELETE_REQUESTED`.
	LifecycleState interface{}
	// The resource name of the Folder. Its format is folders/{folder_id}.
	Name interface{}
	// The resource name of the parent Folder or Organization.
	// Must be of the form `folders/{folder_id}` or `organizations/{org_id}`.
	Parent interface{}
}

// The set of arguments for constructing a Folder resource.
type FolderArgs struct {
	// The folder’s display name.
	// A folder’s display name must be unique amongst its siblings, e.g. no two folders with the same parent can share the same display name. The display name must start and end with a letter or digit, may contain letters, digits, spaces, hyphens and underscores and can be no longer than 30 characters.
	DisplayName interface{}
	// The resource name of the parent Folder or Organization.
	// Must be of the form `folders/{folder_id}` or `organizations/{org_id}`.
	Parent interface{}
}
