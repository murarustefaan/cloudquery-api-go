// Package cloudquery_api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.13.3 DO NOT EDIT.
package cloudquery_api

import (
	"time"
)

const (
	BearerAuthScopes = "bearerAuth.Scopes"
)

// Defines values for PluginCategory.
const (
	CloudInfrastructure  PluginCategory = "cloud-infrastructure"
	Databases            PluginCategory = "databases"
	EngineeringAnalytics PluginCategory = "engineering-analytics"
	Other                PluginCategory = "other"
	SalesMarketing       PluginCategory = "sales-marketing"
)

// Defines values for PluginTier.
const (
	Free PluginTier = "free"
	Paid PluginTier = "paid"
)

// Defines values for PluginVersionPackageType.
const (
	PluginVersionPackageTypeDocker PluginVersionPackageType = "docker"
	PluginVersionPackageTypeNative PluginVersionPackageType = "native"
)

// Defines values for PluginSortBy.
const (
	PluginSortByCreatedAt PluginSortBy = "created_at"
	PluginSortByDownloads PluginSortBy = "downloads"
	PluginSortByName      PluginSortBy = "name"
	PluginSortByUpdatedAt PluginSortBy = "updated_at"
)

// Defines values for VersionSortBy.
const (
	VersionSortByCreatedAt VersionSortBy = "created_at"
)

// Defines values for ListPluginsParamsSortBy.
const (
	ListPluginsParamsSortByCreatedAt ListPluginsParamsSortBy = "created_at"
	ListPluginsParamsSortByDownloads ListPluginsParamsSortBy = "downloads"
	ListPluginsParamsSortByName      ListPluginsParamsSortBy = "name"
	ListPluginsParamsSortByUpdatedAt ListPluginsParamsSortBy = "updated_at"
)

// Defines values for ListPluginVersionsParamsSortBy.
const (
	ListPluginVersionsParamsSortByCreatedAt ListPluginVersionsParamsSortBy = "created_at"
)

// Defines values for CreatePluginVersionJSONBodyPackageType.
const (
	CreatePluginVersionJSONBodyPackageTypeDocker CreatePluginVersionJSONBodyPackageType = "docker"
	CreatePluginVersionJSONBodyPackageTypeNative CreatePluginVersionJSONBodyPackageType = "native"
)

// ApiKey API Key to interact with CloudQuery Cloud under specific team
type ApiKey struct {
	CreatedAt *time.Time `json:"created_at,omitempty"`
	CreatedBy *Email     `json:"created_by,omitempty"`

	// Key API key. Will be shown only in the response when creating the key.
	Key  *string `json:"key,omitempty"`
	Name *string `json:"name,omitempty"`
}

// BasicError Basic Error
type BasicError struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

// Email defines model for Email.
type Email = string

// ImageURL defines model for ImageURL.
type ImageURL struct {
	DownloadUrl string `json:"download_url"`
	UploadUrl   string `json:"upload_url"`
}

// Invitation defines model for Invitation.
type Invitation struct {
	CreatedAt time.Time `json:"created_at"`
	Email     Email     `json:"email"`
	Role      string    `json:"role"`

	// TeamName The unique name for the team.
	TeamName TeamName `json:"team_name"`
}

// ListMetadata defines model for ListMetadata.
type ListMetadata struct {
	LastPage   *int `json:"last_page,omitempty"`
	TotalCount *int `json:"total_count,omitempty"`
}

// Membership defines model for Membership.
type Membership struct {
	Role string `json:"role"`

	// Team CloudQuery Team
	Team *Team `json:"team,omitempty"`

	// User CloudQuery User
	User *User `json:"user,omitempty"`
}

// Plugin CloudQuery Plugin
type Plugin struct {
	// Category Supported categories for plugins
	Category    PluginCategory `json:"category"`
	CreatedAt   time.Time      `json:"created_at"`
	Destination bool           `json:"destination"`

	// DisplayName The plugin's display name
	DisplayName string  `json:"display_name"`
	Homepage    *string `json:"homepage,omitempty"`
	Logo        string  `json:"logo"`

	// Name The unique name for the plugin.
	Name             PluginName `json:"name"`
	Official         bool       `json:"official"`
	Repository       *string    `json:"repository,omitempty"`
	ShortDescription string     `json:"short_description"`
	Source           bool       `json:"source"`

	// TeamName The unique name for the team.
	TeamName TeamName `json:"team_name"`

	// Tier Supported tiers for plugins
	Tier PluginTier `json:"tier"`
}

// PluginCategory Supported categories for plugins
type PluginCategory string

// PluginCreate defines model for PluginCreate.
type PluginCreate struct {
	// Category Supported categories for plugins
	Category    PluginCategory `json:"category"`
	Destination bool           `json:"destination"`

	// DisplayName The plugin's display name, as shown in the CloudQuery Hub.
	DisplayName string  `json:"display_name"`
	Homepage    *string `json:"homepage,omitempty"`

	// Logo URL to the plugin's logo. This will be shown in the CloudQuery Hub. This must point to https://images.cloudquery.io/...
	Logo string `json:"logo"`

	// Name The unique name for the plugin.
	Name       PluginName `json:"name"`
	Repository *string    `json:"repository,omitempty"`

	// ShortDescription Short description of the plugin. This will be shown in the CloudQuery Hub.
	ShortDescription string `json:"short_description"`
	Source           bool   `json:"source"`

	// TeamName The unique name for the team.
	TeamName TeamName `json:"team_name"`

	// Tier Supported tiers for plugins
	Tier PluginTier `json:"tier"`
}

// PluginDocsPage CloudQuery Plugin Documentation Page
type PluginDocsPage struct {
	// Content The content of the documentation page. Supports markdown.
	Content string `json:"content"`

	// Name The unique name for the plugin documentation page.
	Name PluginDocsPageName `json:"name"`

	// Title The title of the documentation page
	Title string `json:"title"`
}

// PluginDocsPageName The unique name for the plugin documentation page.
type PluginDocsPageName = string

// PluginName The unique name for the plugin.
type PluginName = string

// PluginTable CloudQuery Plugin Table
type PluginTable struct {
	// Description Description of the table
	Description *string `json:"description,omitempty"`

	// IsIncremental Whether the table is incremental
	IsIncremental *bool `json:"is_incremental,omitempty"`

	// Name Name of the table
	Name *PluginTableName `json:"name,omitempty"`

	// Parent Name of the parent table, if any
	Parent *string `json:"parent,omitempty"`

	// Relations Names of the tables that depend on this table
	Relations *[]string `json:"relations,omitempty"`

	// Title Title of the table
	Title *string `json:"title,omitempty"`
}

// PluginTableColumn CloudQuery Plugin Column
type PluginTableColumn struct {
	// Description Description of the column
	Description *string `json:"description,omitempty"`

	// IncrementalKey Whether the column is used as an incremental key
	IncrementalKey *bool `json:"incremental_key,omitempty"`

	// IsUnique Whether the column has a unique constraint
	IsUnique *bool `json:"is_unique,omitempty"`

	// Name Name of the column
	Name *string `json:"name,omitempty"`

	// NotNull Whether the column is nullable
	NotNull *bool `json:"not_null,omitempty"`

	// PrimaryKey Whether the column is part of the primary key
	PrimaryKey *bool `json:"primary_key,omitempty"`

	// Type Arrow Type of the column
	Type *string `json:"type,omitempty"`
}

// PluginTableDetails defines model for PluginTableDetails.
type PluginTableDetails struct {
	// Columns List of columns
	Columns *[]PluginTableColumn `json:"columns,omitempty"`

	// Description Description of the table
	Description *string `json:"description,omitempty"`

	// IsIncremental Whether the table is incremental
	IsIncremental *bool `json:"is_incremental,omitempty"`

	// Name Name of the table
	Name *string `json:"name,omitempty"`

	// Parent Name of the parent table, if any
	Parent *string `json:"parent,omitempty"`

	// Relations Names of the tables that depend on this table
	Relations *[]string `json:"relations,omitempty"`

	// Title Title of the table
	Title *string `json:"title,omitempty"`
}

// PluginTableName Name of the table
type PluginTableName = string

// PluginTier Supported tiers for plugins
type PluginTier string

// PluginUpdate defines model for PluginUpdate.
type PluginUpdate struct {
	// Category Supported categories for plugins
	Category    PluginCategory `json:"category"`
	Destination bool           `json:"destination"`

	// DisplayName The plugin's display name, as shown in the CloudQuery Hub.
	DisplayName string  `json:"display_name"`
	Homepage    *string `json:"homepage,omitempty"`

	// Listed If plugin is not listed, it won't be visible or accessible in the registry.
	Listed *bool `json:"listed,omitempty"`

	// Logo URL to the plugin's logo. This will be shown in the CloudQuery Hub. This must point to https://images.cloudquery.io/...
	Logo       string  `json:"logo"`
	Repository *string `json:"repository,omitempty"`

	// ShortDescription Short description of the plugin. This will be shown in the CloudQuery Hub.
	ShortDescription string `json:"short_description"`
	Source           bool   `json:"source"`

	// Tier Supported tiers for plugins
	Tier PluginTier `json:"tier"`
}

// PluginVersion CloudQuery Plugin Version
type PluginVersion struct {
	// Checksums The checksums of the plugin assets
	Checksums []string `json:"checksums"`

	// CreatedAt The date and time the plugin version was created.
	CreatedAt time.Time `json:"created_at"`

	// Draft If a plugin version is in draft, it will not show to members outside the team or be counted as the latest version.
	Draft bool `json:"draft"`

	// Message Description of what's new or changed in this version (supports markdown)
	Message string `json:"message"`

	// Name The version in semantic version format.
	Name VersionName `json:"name"`

	// PackageType The package type of the plugin assets
	PackageType PluginVersionPackageType `json:"package_type"`

	// Protocols The CloudQuery protocols supported by this plugin version
	Protocols []int `json:"protocols"`

	// Retracted If a plugin version is retracted, assets will not be available and it will not be counted as the latest version.
	Retracted bool `json:"retracted"`

	// SupportedTargets The targets supported by this plugin version, formatted as <os>_<arch>
	SupportedTargets []string `json:"supported_targets"`
}

// PluginVersionPackageType The package type of the plugin assets
type PluginVersionPackageType string

// PluginVersionUpdate defines model for PluginVersionUpdate.
type PluginVersionUpdate struct {
	// Checksums The SHA-256 checksums of the plugin binaries, one per supported target.
	Checksums *[]string `json:"checksums,omitempty"`

	// Draft If a plugin version is in draft, it will not show to members outside the team or be counted as the latest version. Once draft is set to false, only certain fields can be updated.
	Draft *bool `json:"draft,omitempty"`

	// Message Description of what's new or changed in this version (supports markdown)
	Message *string `json:"message,omitempty"`

	// PackageType The package type of the plugin binaries
	PackageType *string `json:"package_type,omitempty"`

	// Protocols The supported CloudQuery protocols by this plugin version
	Protocols *[]int `json:"protocols,omitempty"`

	// Retracted If a plugin version is retracted, assets will not be available and it will not be counted as the latest version.
	Retracted        *bool     `json:"retracted,omitempty"`
	SupportedTargets *[]string `json:"supported_targets,omitempty"`
}

// ReleaseURL defines model for ReleaseURL.
type ReleaseURL struct {
	Url *string `json:"url,omitempty"`
}

// Team CloudQuery Team
type Team struct {
	CreatedAt *time.Time `json:"created_at,omitempty"`

	// DisplayName The team's display name
	DisplayName string `json:"display_name"`

	// Name The unique name for the team.
	Name TeamName `json:"name"`
}

// TeamName The unique name for the team.
type TeamName = string

// User CloudQuery User
type User struct {
	CreatedAt *time.Time `json:"created_at,omitempty"`
	Email     Email      `json:"email"`

	// Name The unique name for the user.
	Name      *UserName  `json:"name,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

// UserName The unique name for the user.
type UserName = string

// ValidationError defines model for ValidationError.
type ValidationError struct {
	Errors      *[]string          `json:"errors,omitempty"`
	FieldErrors *map[string]string `json:"field_errors,omitempty"`
	Message     string             `json:"message"`
	Status      int                `json:"status"`
}

// VersionName The version in semantic version format.
type VersionName = string

// ApiKeyName defines model for apikey_name.
type ApiKeyName = string

// Page defines model for page.
type Page = int64

// PerPage defines model for per_page.
type PerPage = int64

// PluginSortBy defines model for plugin_sort_by.
type PluginSortBy string

// TargetName defines model for target_name.
type TargetName = string

// VersionSortBy defines model for version_sort_by.
type VersionSortBy string

// Forbidden Basic Error
type Forbidden = BasicError

// InternalError Basic Error
type InternalError = BasicError

// NotFound Basic Error
type NotFound = BasicError

// RequiresAuthentication Basic Error
type RequiresAuthentication = BasicError

// UnprocessableEntity defines model for UnprocessableEntity.
type UnprocessableEntity = ValidationError

// ListPluginsParams defines parameters for ListPlugins.
type ListPluginsParams struct {
	// SortBy The field to sort by
	SortBy *ListPluginsParamsSortBy `form:"sort_by,omitempty" json:"sort_by,omitempty"`

	// Page Page number of the results to fetch
	Page *Page `form:"page,omitempty" json:"page,omitempty"`

	// PerPage The number of results per page (max 1000).
	PerPage *PerPage `form:"per_page,omitempty" json:"per_page,omitempty"`
}

// ListPluginsParamsSortBy defines parameters for ListPlugins.
type ListPluginsParamsSortBy string

// ListPluginVersionsParams defines parameters for ListPluginVersions.
type ListPluginVersionsParams struct {
	// SortBy The field to sort by
	SortBy *ListPluginVersionsParamsSortBy `form:"sort_by,omitempty" json:"sort_by,omitempty"`

	// Page Page number of the results to fetch
	Page *Page `form:"page,omitempty" json:"page,omitempty"`

	// PerPage The number of results per page (max 1000).
	PerPage *PerPage `form:"per_page,omitempty" json:"per_page,omitempty"`
}

// ListPluginVersionsParamsSortBy defines parameters for ListPluginVersions.
type ListPluginVersionsParamsSortBy string

// CreatePluginVersionJSONBody defines parameters for CreatePluginVersion.
type CreatePluginVersionJSONBody struct {
	// Checksums List of SHA-256 checksums for this plugin version, one for each supported target.
	Checksums []string `json:"checksums"`

	// Message A message describing what's new or changed in this version.
	// This message will be displayed to users in the plugin's changelog.
	// Supports limited markdown syntax.
	Message string `json:"message"`

	// Name The version in semantic version format.
	Name VersionName `json:"name"`

	// PackageType The package type of the plugin assets
	PackageType CreatePluginVersionJSONBodyPackageType `json:"package_type"`

	// Protocols List of protocols supported by this plugin version
	Protocols []int `json:"protocols"`

	// SupportedTargets The targets supported by this plugin version, formatted as <os>_<arch>
	SupportedTargets []string `json:"supported_targets"`
}

// CreatePluginVersionJSONBodyPackageType defines parameters for CreatePluginVersion.
type CreatePluginVersionJSONBodyPackageType string

// DeletePluginVersionDocsJSONBody defines parameters for DeletePluginVersionDocs.
type DeletePluginVersionDocsJSONBody struct {
	Names []PluginDocsPageName `json:"names"`
}

// ListPluginVersionDocsParams defines parameters for ListPluginVersionDocs.
type ListPluginVersionDocsParams struct {
	// Page Page number of the results to fetch
	Page *Page `form:"page,omitempty" json:"page,omitempty"`

	// PerPage The number of results per page (max 1000).
	PerPage *PerPage `form:"per_page,omitempty" json:"per_page,omitempty"`
}

// CreatePluginVersionDocsJSONBody defines parameters for CreatePluginVersionDocs.
type CreatePluginVersionDocsJSONBody struct {
	Pages []PluginDocsPage `json:"pages"`
}

// DeletePluginVersionTablesJSONBody defines parameters for DeletePluginVersionTables.
type DeletePluginVersionTablesJSONBody struct {
	Names []PluginTableName `json:"names"`
}

// ListPluginVersionTablesParams defines parameters for ListPluginVersionTables.
type ListPluginVersionTablesParams struct {
	// Page Page number of the results to fetch
	Page *Page `form:"page,omitempty" json:"page,omitempty"`

	// PerPage The number of results per page (max 1000).
	PerPage *PerPage `form:"per_page,omitempty" json:"per_page,omitempty"`
}

// CreatePluginVersionTablesJSONBody defines parameters for CreatePluginVersionTables.
type CreatePluginVersionTablesJSONBody struct {
	Tables []PluginTable `json:"tables"`
}

// ListTeamsParams defines parameters for ListTeams.
type ListTeamsParams struct {
	// PerPage The number of results per page (max 1000).
	PerPage *PerPage `form:"per_page,omitempty" json:"per_page,omitempty"`

	// Page Page number of the results to fetch
	Page *Page `form:"page,omitempty" json:"page,omitempty"`
}

// CreateTeamJSONBody defines parameters for CreateTeam.
type CreateTeamJSONBody struct {
	// DisplayName The team's display name
	DisplayName string `json:"display_name"`

	// Name The unique name for the team.
	Name TeamName `json:"name"`
}

// UpdateTeamJSONBody defines parameters for UpdateTeam.
type UpdateTeamJSONBody struct {
	// DisplayName The team's display name
	DisplayName *string `json:"display_name,omitempty"`
}

// ListTeamApiKeysParams defines parameters for ListTeamApiKeys.
type ListTeamApiKeysParams struct {
	// PerPage The number of results per page (max 1000).
	PerPage *PerPage `form:"per_page,omitempty" json:"per_page,omitempty"`

	// Page Page number of the results to fetch
	Page *Page `form:"page,omitempty" json:"page,omitempty"`
}

// ListTeamInvitationsParams defines parameters for ListTeamInvitations.
type ListTeamInvitationsParams struct {
	// Page Page number of the results to fetch
	Page *Page `form:"page,omitempty" json:"page,omitempty"`

	// PerPage The number of results per page (max 1000).
	PerPage *PerPage `form:"per_page,omitempty" json:"per_page,omitempty"`
}

// EmailTeamInvitationJSONBody defines parameters for EmailTeamInvitation.
type EmailTeamInvitationJSONBody struct {
	Email string `json:"email"`
	Role  string `json:"role"`
}

// GetTeamMembershipsParams defines parameters for GetTeamMemberships.
type GetTeamMembershipsParams struct {
	// Page Page number of the results to fetch
	Page *Page `form:"page,omitempty" json:"page,omitempty"`

	// PerPage The number of results per page (max 1000).
	PerPage *PerPage `form:"per_page,omitempty" json:"per_page,omitempty"`
}

// ListPluginsByTeamParams defines parameters for ListPluginsByTeam.
type ListPluginsByTeamParams struct {
	// Page Page number of the results to fetch
	Page *Page `form:"page,omitempty" json:"page,omitempty"`

	// PerPage The number of results per page (max 1000).
	PerPage *PerPage `form:"per_page,omitempty" json:"per_page,omitempty"`
}

// ListUsersByTeamParams defines parameters for ListUsersByTeam.
type ListUsersByTeamParams struct {
	// PerPage The number of results per page (max 1000).
	PerPage *PerPage `form:"per_page,omitempty" json:"per_page,omitempty"`

	// Page Page number of the results to fetch
	Page *Page `form:"page,omitempty" json:"page,omitempty"`
}

// ListCurrentUserInvitationsParams defines parameters for ListCurrentUserInvitations.
type ListCurrentUserInvitationsParams struct {
	// Page Page number of the results to fetch
	Page *Page `form:"page,omitempty" json:"page,omitempty"`

	// PerPage The number of results per page (max 1000).
	PerPage *PerPage `form:"per_page,omitempty" json:"per_page,omitempty"`
}

// GetCurrentUserMembershipsParams defines parameters for GetCurrentUserMemberships.
type GetCurrentUserMembershipsParams struct {
	// Page Page number of the results to fetch
	Page *Page `form:"page,omitempty" json:"page,omitempty"`

	// PerPage The number of results per page (max 1000).
	PerPage *PerPage `form:"per_page,omitempty" json:"per_page,omitempty"`
}

// CreatePluginJSONRequestBody defines body for CreatePlugin for application/json ContentType.
type CreatePluginJSONRequestBody = PluginCreate

// UpdatePluginJSONRequestBody defines body for UpdatePlugin for application/json ContentType.
type UpdatePluginJSONRequestBody = PluginUpdate

// CreatePluginVersionJSONRequestBody defines body for CreatePluginVersion for application/json ContentType.
type CreatePluginVersionJSONRequestBody CreatePluginVersionJSONBody

// UpdatePluginVersionJSONRequestBody defines body for UpdatePluginVersion for application/json ContentType.
type UpdatePluginVersionJSONRequestBody = PluginVersionUpdate

// DeletePluginVersionDocsJSONRequestBody defines body for DeletePluginVersionDocs for application/json ContentType.
type DeletePluginVersionDocsJSONRequestBody DeletePluginVersionDocsJSONBody

// CreatePluginVersionDocsJSONRequestBody defines body for CreatePluginVersionDocs for application/json ContentType.
type CreatePluginVersionDocsJSONRequestBody CreatePluginVersionDocsJSONBody

// DeletePluginVersionTablesJSONRequestBody defines body for DeletePluginVersionTables for application/json ContentType.
type DeletePluginVersionTablesJSONRequestBody DeletePluginVersionTablesJSONBody

// CreatePluginVersionTablesJSONRequestBody defines body for CreatePluginVersionTables for application/json ContentType.
type CreatePluginVersionTablesJSONRequestBody CreatePluginVersionTablesJSONBody

// CreateTeamJSONRequestBody defines body for CreateTeam for application/json ContentType.
type CreateTeamJSONRequestBody CreateTeamJSONBody

// UpdateTeamJSONRequestBody defines body for UpdateTeam for application/json ContentType.
type UpdateTeamJSONRequestBody UpdateTeamJSONBody

// EmailTeamInvitationJSONRequestBody defines body for EmailTeamInvitation for application/json ContentType.
type EmailTeamInvitationJSONRequestBody EmailTeamInvitationJSONBody
