// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package organizations

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Allows creation and management of a Google Cloud Platform project.
// 
// Projects created with this resource must be associated with an Organization.
// See the [Organization documentation](https://cloud.google.com/resource-manager/docs/quickstarts) for more details.
// 
// The service account used to run this provider when creating a `organizations.Project`
// resource must have `roles/resourcemanager.projectCreator`. See the
// [Access Control for Organizations Using IAM](https://cloud.google.com/resource-manager/docs/access-control-org)
// doc for more information.
// 
// Note that prior to 0.8.5, `organizations.Project` functioned like a data source,
// meaning any project referenced by it had to be created and managed outside
// this provider. As of 0.8.5, `organizations.Project` functions like any other
// resource, with this provider creating and managing the project. To replicate the old
// behavior, either:
// 
// * Use the project ID directly in whatever is referencing the project, using the
//   [projects.IAMPolicy](https://www.terraform.io/docs/providers/google/r/google_project_iam.html)
//   to replace the old `policyData` property.
// * Use the [import](https://www.terraform.io/docs/import/usage.html) functionality
//   to import your pre-existing project into this provider, where it can be referenced and
//   used just like always, keeping in mind that this provider will attempt to undo any changes
//   made outside this provider.
// 
// > It's important to note that any project resources that were added to your config
// prior to 0.8.5 will continue to function as they always have, and will not be managed by
// this provider. Only newly added projects are affected.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/project.html.markdown.
type Project struct {
	s *pulumi.ResourceState
}

// NewProject registers a new resource with the given unique name, arguments, and options.
func NewProject(ctx *pulumi.Context,
	name string, args *ProjectArgs, opts ...pulumi.ResourceOpt) (*Project, error) {
	if args == nil || args.ProjectId == nil {
		return nil, errors.New("missing required argument 'ProjectId'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["autoCreateNetwork"] = nil
		inputs["billingAccount"] = nil
		inputs["folderId"] = nil
		inputs["labels"] = nil
		inputs["name"] = nil
		inputs["orgId"] = nil
		inputs["projectId"] = nil
		inputs["skipDelete"] = nil
	} else {
		inputs["autoCreateNetwork"] = args.AutoCreateNetwork
		inputs["billingAccount"] = args.BillingAccount
		inputs["folderId"] = args.FolderId
		inputs["labels"] = args.Labels
		inputs["name"] = args.Name
		inputs["orgId"] = args.OrgId
		inputs["projectId"] = args.ProjectId
		inputs["skipDelete"] = args.SkipDelete
	}
	inputs["number"] = nil
	s, err := ctx.RegisterResource("gcp:organizations/project:Project", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Project{s: s}, nil
}

// GetProject gets an existing Project resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProject(ctx *pulumi.Context,
	name string, id pulumi.ID, state *ProjectState, opts ...pulumi.ResourceOpt) (*Project, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["autoCreateNetwork"] = state.AutoCreateNetwork
		inputs["billingAccount"] = state.BillingAccount
		inputs["folderId"] = state.FolderId
		inputs["labels"] = state.Labels
		inputs["name"] = state.Name
		inputs["number"] = state.Number
		inputs["orgId"] = state.OrgId
		inputs["projectId"] = state.ProjectId
		inputs["skipDelete"] = state.SkipDelete
	}
	s, err := ctx.ReadResource("gcp:organizations/project:Project", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Project{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Project) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Project) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// Create the 'default' network automatically.  Default `true`.
// If set to `false`, the default network will be deleted.  Note that, for quota purposes, you
// will still need to have 1 network slot available to create the project succesfully, even if
// you set `autoCreateNetwork` to `false`, since the network will exist momentarily.
func (r *Project) AutoCreateNetwork() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["autoCreateNetwork"])
}

// The alphanumeric ID of the billing account this project
// belongs to. The user or service account performing this operation with this provider
// must have Billing Account Administrator privileges (`roles/billing.admin`) in
// the organization. See [Google Cloud Billing API Access Control](https://cloud.google.com/billing/v1/how-tos/access-control)
// for more details.
func (r *Project) BillingAccount() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["billingAccount"])
}

// The numeric ID of the folder this project should be
// created under. Only one of `orgId` or `folderId` may be
// specified. If the `folderId` is specified, then the project is
// created under the specified folder. Changing this forces the
// project to be migrated to the newly specified folder.
func (r *Project) FolderId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["folderId"])
}

// A set of key/value label pairs to assign to the project.
func (r *Project) Labels() *pulumi.MapOutput {
	return (*pulumi.MapOutput)(r.s.State["labels"])
}

// The display name of the project.
func (r *Project) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// The numeric identifier of the project.
func (r *Project) Number() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["number"])
}

// The numeric ID of the organization this project belongs to.
// Changing this forces a new project to be created.  Only one of
// `orgId` or `folderId` may be specified. If the `orgId` is
// specified then the project is created at the top level. Changing
// this forces the project to be migrated to the newly specified
// organization.
func (r *Project) OrgId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["orgId"])
}

// The project ID. Changing this forces a new project to be created.
func (r *Project) ProjectId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["projectId"])
}

// If true, the resource can be deleted
// without deleting the Project via the Google API.
func (r *Project) SkipDelete() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["skipDelete"])
}

// Input properties used for looking up and filtering Project resources.
type ProjectState struct {
	// Create the 'default' network automatically.  Default `true`.
	// If set to `false`, the default network will be deleted.  Note that, for quota purposes, you
	// will still need to have 1 network slot available to create the project succesfully, even if
	// you set `autoCreateNetwork` to `false`, since the network will exist momentarily.
	AutoCreateNetwork interface{}
	// The alphanumeric ID of the billing account this project
	// belongs to. The user or service account performing this operation with this provider
	// must have Billing Account Administrator privileges (`roles/billing.admin`) in
	// the organization. See [Google Cloud Billing API Access Control](https://cloud.google.com/billing/v1/how-tos/access-control)
	// for more details.
	BillingAccount interface{}
	// The numeric ID of the folder this project should be
	// created under. Only one of `orgId` or `folderId` may be
	// specified. If the `folderId` is specified, then the project is
	// created under the specified folder. Changing this forces the
	// project to be migrated to the newly specified folder.
	FolderId interface{}
	// A set of key/value label pairs to assign to the project.
	Labels interface{}
	// The display name of the project.
	Name interface{}
	// The numeric identifier of the project.
	Number interface{}
	// The numeric ID of the organization this project belongs to.
	// Changing this forces a new project to be created.  Only one of
	// `orgId` or `folderId` may be specified. If the `orgId` is
	// specified then the project is created at the top level. Changing
	// this forces the project to be migrated to the newly specified
	// organization.
	OrgId interface{}
	// The project ID. Changing this forces a new project to be created.
	ProjectId interface{}
	// If true, the resource can be deleted
	// without deleting the Project via the Google API.
	SkipDelete interface{}
}

// The set of arguments for constructing a Project resource.
type ProjectArgs struct {
	// Create the 'default' network automatically.  Default `true`.
	// If set to `false`, the default network will be deleted.  Note that, for quota purposes, you
	// will still need to have 1 network slot available to create the project succesfully, even if
	// you set `autoCreateNetwork` to `false`, since the network will exist momentarily.
	AutoCreateNetwork interface{}
	// The alphanumeric ID of the billing account this project
	// belongs to. The user or service account performing this operation with this provider
	// must have Billing Account Administrator privileges (`roles/billing.admin`) in
	// the organization. See [Google Cloud Billing API Access Control](https://cloud.google.com/billing/v1/how-tos/access-control)
	// for more details.
	BillingAccount interface{}
	// The numeric ID of the folder this project should be
	// created under. Only one of `orgId` or `folderId` may be
	// specified. If the `folderId` is specified, then the project is
	// created under the specified folder. Changing this forces the
	// project to be migrated to the newly specified folder.
	FolderId interface{}
	// A set of key/value label pairs to assign to the project.
	Labels interface{}
	// The display name of the project.
	Name interface{}
	// The numeric ID of the organization this project belongs to.
	// Changing this forces a new project to be created.  Only one of
	// `orgId` or `folderId` may be specified. If the `orgId` is
	// specified then the project is created at the top level. Changing
	// this forces the project to be migrated to the newly specified
	// organization.
	OrgId interface{}
	// The project ID. Changing this forces a new project to be created.
	ProjectId interface{}
	// If true, the resource can be deleted
	// without deleting the Project via the Google API.
	SkipDelete interface{}
}
