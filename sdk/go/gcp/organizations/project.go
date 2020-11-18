// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package organizations

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
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
type Project struct {
	pulumi.CustomResourceState

	// Create the 'default' network automatically.  Default `true`.
	// If set to `false`, the default network will be deleted.  Note that, for quota purposes, you
	// will still need to have 1 network slot available to create the project successfully, even if
	// you set `autoCreateNetwork` to `false`, since the network will exist momentarily.
	AutoCreateNetwork pulumi.BoolPtrOutput `pulumi:"autoCreateNetwork"`
	// The alphanumeric ID of the billing account this project
	// belongs to. The user or service account performing this operation with the provider
	// must have Billing Account Administrator privileges (`roles/billing.admin`) in
	// the organization. See [Google Cloud Billing API Access Control](https://cloud.google.com/billing/v1/how-tos/access-control)
	// for more details.
	BillingAccount pulumi.StringPtrOutput `pulumi:"billingAccount"`
	// The numeric ID of the folder this project should be
	// created under. Only one of `orgId` or `folderId` may be
	// specified. If the `folderId` is specified, then the project is
	// created under the specified folder. Changing this forces the
	// project to be migrated to the newly specified folder.
	FolderId pulumi.StringOutput `pulumi:"folderId"`
	// A set of key/value label pairs to assign to the project.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// The display name of the project.
	Name pulumi.StringOutput `pulumi:"name"`
	// The numeric identifier of the project.
	Number pulumi.StringOutput `pulumi:"number"`
	// The numeric ID of the organization this project belongs to.
	// Changing this forces a new project to be created.  Only one of
	// `orgId` or `folderId` may be specified. If the `orgId` is
	// specified then the project is created at the top level. Changing
	// this forces the project to be migrated to the newly specified
	// organization.
	OrgId pulumi.StringOutput `pulumi:"orgId"`
	// The project ID. Changing this forces a new project to be created.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// If true, the resource can be deleted
	// without deleting the Project via the Google API.
	SkipDelete pulumi.BoolOutput `pulumi:"skipDelete"`
}

// NewProject registers a new resource with the given unique name, arguments, and options.
func NewProject(ctx *pulumi.Context,
	name string, args *ProjectArgs, opts ...pulumi.ResourceOption) (*Project, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ProjectId == nil {
		return nil, errors.New("invalid value for required argument 'ProjectId'")
	}
	var resource Project
	err := ctx.RegisterResource("gcp:organizations/project:Project", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProject gets an existing Project resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProject(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProjectState, opts ...pulumi.ResourceOption) (*Project, error) {
	var resource Project
	err := ctx.ReadResource("gcp:organizations/project:Project", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Project resources.
type projectState struct {
	// Create the 'default' network automatically.  Default `true`.
	// If set to `false`, the default network will be deleted.  Note that, for quota purposes, you
	// will still need to have 1 network slot available to create the project successfully, even if
	// you set `autoCreateNetwork` to `false`, since the network will exist momentarily.
	AutoCreateNetwork *bool `pulumi:"autoCreateNetwork"`
	// The alphanumeric ID of the billing account this project
	// belongs to. The user or service account performing this operation with the provider
	// must have Billing Account Administrator privileges (`roles/billing.admin`) in
	// the organization. See [Google Cloud Billing API Access Control](https://cloud.google.com/billing/v1/how-tos/access-control)
	// for more details.
	BillingAccount *string `pulumi:"billingAccount"`
	// The numeric ID of the folder this project should be
	// created under. Only one of `orgId` or `folderId` may be
	// specified. If the `folderId` is specified, then the project is
	// created under the specified folder. Changing this forces the
	// project to be migrated to the newly specified folder.
	FolderId *string `pulumi:"folderId"`
	// A set of key/value label pairs to assign to the project.
	Labels map[string]string `pulumi:"labels"`
	// The display name of the project.
	Name *string `pulumi:"name"`
	// The numeric identifier of the project.
	Number *string `pulumi:"number"`
	// The numeric ID of the organization this project belongs to.
	// Changing this forces a new project to be created.  Only one of
	// `orgId` or `folderId` may be specified. If the `orgId` is
	// specified then the project is created at the top level. Changing
	// this forces the project to be migrated to the newly specified
	// organization.
	OrgId *string `pulumi:"orgId"`
	// The project ID. Changing this forces a new project to be created.
	ProjectId *string `pulumi:"projectId"`
	// If true, the resource can be deleted
	// without deleting the Project via the Google API.
	SkipDelete *bool `pulumi:"skipDelete"`
}

type ProjectState struct {
	// Create the 'default' network automatically.  Default `true`.
	// If set to `false`, the default network will be deleted.  Note that, for quota purposes, you
	// will still need to have 1 network slot available to create the project successfully, even if
	// you set `autoCreateNetwork` to `false`, since the network will exist momentarily.
	AutoCreateNetwork pulumi.BoolPtrInput
	// The alphanumeric ID of the billing account this project
	// belongs to. The user or service account performing this operation with the provider
	// must have Billing Account Administrator privileges (`roles/billing.admin`) in
	// the organization. See [Google Cloud Billing API Access Control](https://cloud.google.com/billing/v1/how-tos/access-control)
	// for more details.
	BillingAccount pulumi.StringPtrInput
	// The numeric ID of the folder this project should be
	// created under. Only one of `orgId` or `folderId` may be
	// specified. If the `folderId` is specified, then the project is
	// created under the specified folder. Changing this forces the
	// project to be migrated to the newly specified folder.
	FolderId pulumi.StringPtrInput
	// A set of key/value label pairs to assign to the project.
	Labels pulumi.StringMapInput
	// The display name of the project.
	Name pulumi.StringPtrInput
	// The numeric identifier of the project.
	Number pulumi.StringPtrInput
	// The numeric ID of the organization this project belongs to.
	// Changing this forces a new project to be created.  Only one of
	// `orgId` or `folderId` may be specified. If the `orgId` is
	// specified then the project is created at the top level. Changing
	// this forces the project to be migrated to the newly specified
	// organization.
	OrgId pulumi.StringPtrInput
	// The project ID. Changing this forces a new project to be created.
	ProjectId pulumi.StringPtrInput
	// If true, the resource can be deleted
	// without deleting the Project via the Google API.
	SkipDelete pulumi.BoolPtrInput
}

func (ProjectState) ElementType() reflect.Type {
	return reflect.TypeOf((*projectState)(nil)).Elem()
}

type projectArgs struct {
	// Create the 'default' network automatically.  Default `true`.
	// If set to `false`, the default network will be deleted.  Note that, for quota purposes, you
	// will still need to have 1 network slot available to create the project successfully, even if
	// you set `autoCreateNetwork` to `false`, since the network will exist momentarily.
	AutoCreateNetwork *bool `pulumi:"autoCreateNetwork"`
	// The alphanumeric ID of the billing account this project
	// belongs to. The user or service account performing this operation with the provider
	// must have Billing Account Administrator privileges (`roles/billing.admin`) in
	// the organization. See [Google Cloud Billing API Access Control](https://cloud.google.com/billing/v1/how-tos/access-control)
	// for more details.
	BillingAccount *string `pulumi:"billingAccount"`
	// The numeric ID of the folder this project should be
	// created under. Only one of `orgId` or `folderId` may be
	// specified. If the `folderId` is specified, then the project is
	// created under the specified folder. Changing this forces the
	// project to be migrated to the newly specified folder.
	FolderId *string `pulumi:"folderId"`
	// A set of key/value label pairs to assign to the project.
	Labels map[string]string `pulumi:"labels"`
	// The display name of the project.
	Name *string `pulumi:"name"`
	// The numeric ID of the organization this project belongs to.
	// Changing this forces a new project to be created.  Only one of
	// `orgId` or `folderId` may be specified. If the `orgId` is
	// specified then the project is created at the top level. Changing
	// this forces the project to be migrated to the newly specified
	// organization.
	OrgId *string `pulumi:"orgId"`
	// The project ID. Changing this forces a new project to be created.
	ProjectId string `pulumi:"projectId"`
	// If true, the resource can be deleted
	// without deleting the Project via the Google API.
	SkipDelete *bool `pulumi:"skipDelete"`
}

// The set of arguments for constructing a Project resource.
type ProjectArgs struct {
	// Create the 'default' network automatically.  Default `true`.
	// If set to `false`, the default network will be deleted.  Note that, for quota purposes, you
	// will still need to have 1 network slot available to create the project successfully, even if
	// you set `autoCreateNetwork` to `false`, since the network will exist momentarily.
	AutoCreateNetwork pulumi.BoolPtrInput
	// The alphanumeric ID of the billing account this project
	// belongs to. The user or service account performing this operation with the provider
	// must have Billing Account Administrator privileges (`roles/billing.admin`) in
	// the organization. See [Google Cloud Billing API Access Control](https://cloud.google.com/billing/v1/how-tos/access-control)
	// for more details.
	BillingAccount pulumi.StringPtrInput
	// The numeric ID of the folder this project should be
	// created under. Only one of `orgId` or `folderId` may be
	// specified. If the `folderId` is specified, then the project is
	// created under the specified folder. Changing this forces the
	// project to be migrated to the newly specified folder.
	FolderId pulumi.StringPtrInput
	// A set of key/value label pairs to assign to the project.
	Labels pulumi.StringMapInput
	// The display name of the project.
	Name pulumi.StringPtrInput
	// The numeric ID of the organization this project belongs to.
	// Changing this forces a new project to be created.  Only one of
	// `orgId` or `folderId` may be specified. If the `orgId` is
	// specified then the project is created at the top level. Changing
	// this forces the project to be migrated to the newly specified
	// organization.
	OrgId pulumi.StringPtrInput
	// The project ID. Changing this forces a new project to be created.
	ProjectId pulumi.StringInput
	// If true, the resource can be deleted
	// without deleting the Project via the Google API.
	SkipDelete pulumi.BoolPtrInput
}

func (ProjectArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*projectArgs)(nil)).Elem()
}

type ProjectInput interface {
	pulumi.Input

	ToProjectOutput() ProjectOutput
	ToProjectOutputWithContext(ctx context.Context) ProjectOutput
}

func (Project) ElementType() reflect.Type {
	return reflect.TypeOf((*Project)(nil)).Elem()
}

func (i Project) ToProjectOutput() ProjectOutput {
	return i.ToProjectOutputWithContext(context.Background())
}

func (i Project) ToProjectOutputWithContext(ctx context.Context) ProjectOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectOutput)
}

type ProjectOutput struct {
	*pulumi.OutputState
}

func (ProjectOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProjectOutput)(nil)).Elem()
}

func (o ProjectOutput) ToProjectOutput() ProjectOutput {
	return o
}

func (o ProjectOutput) ToProjectOutputWithContext(ctx context.Context) ProjectOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ProjectOutput{})
}
