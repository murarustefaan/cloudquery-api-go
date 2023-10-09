// Package cloudquery_api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.13.3 DO NOT EDIT.
package cloudquery_api

import (
	"time"

	openapi_types "github.com/deepmap/oapi-codegen/pkg/types"
)

const (
	BearerAuthScopes = "bearerAuth.Scopes"
)

// Defines values for APIKeyScope.
const (
	ReadAndWrite APIKeyScope = "read-and-write"
)

// Defines values for PluginCategory.
const (
	CloudInfrastructure  PluginCategory = "cloud-infrastructure"
	Databases            PluginCategory = "databases"
	EngineeringAnalytics PluginCategory = "engineering-analytics"
	Other                PluginCategory = "other"
	SalesMarketing       PluginCategory = "sales-marketing"
)

// Defines values for PluginKind.
const (
	Destination PluginKind = "destination"
	Source      PluginKind = "source"
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

// Defines values for EmailTeamInvitationJSONBodyRole.
const (
	Admin  EmailTeamInvitationJSONBodyRole = "admin"
	Member EmailTeamInvitationJSONBodyRole = "member"
)

// APIKey API Key to interact with CloudQuery Cloud under specific team
type APIKey struct {
	CreatedAt *time.Time `json:"created_at,omitempty"`
	CreatedBy *Email     `json:"created_by,omitempty"`

	// ExpiresAt Timestamp at which API key will stop working
	ExpiresAt time.Time `json:"expires_at"`

	// Key API key. Will be shown only in the response when creating the key.
	Key *string `json:"key,omitempty"`

	// Name Name of the API key
	Name APIKeyName `json:"name"`

	// Scope Scope of permissions for the API key. API keys are used for creating new plugin versions and downloading existing plugins
	Scope APIKeyScope `json:"scope"`
}

// APIKeyName Name of the API key
type APIKeyName = string

// APIKeyScope Scope of permissions for the API key. API keys are used for creating new plugin versions and downloading existing plugins
type APIKeyScope string

// BasicError Basic Error
type BasicError struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

// Email defines model for Email.
type Email = openapi_types.Email

// FieldError defines model for FieldError.
type FieldError struct {
	Errors      *[]string          `json:"errors,omitempty"`
	FieldErrors *map[string]string `json:"field_errors,omitempty"`
	Message     string             `json:"message"`
	Status      int                `json:"status"`
}

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

// InvitationWithToken defines model for InvitationWithToken.
type InvitationWithToken struct {
	CreatedAt time.Time `json:"created_at"`
	Email     Email     `json:"email"`
	Role      string    `json:"role"`

	// TeamName The unique name for the team.
	TeamName TeamName `json:"team_name"`

	// Token The token used to accept the invitation
	Token openapi_types.UUID `json:"token"`
}

// ListMetadata defines model for ListMetadata.
type ListMetadata struct {
	LastPage   *int `json:"last_page,omitempty"`
	TotalCount *int `json:"total_count,omitempty"`
}

// MembershipWithTeam defines model for MembershipWithTeam.
type MembershipWithTeam struct {
	Role string `json:"role"`

	// Team CloudQuery Team
	Team *Team `json:"team,omitempty"`
}

// MembershipWithUser defines model for MembershipWithUser.
type MembershipWithUser struct {
	Role string `json:"role"`

	// User CloudQuery User
	User *User `json:"user,omitempty"`
}

// MonthlyLimit A configurable monthly limit for plugin usage.
type MonthlyLimit struct {
	// CreatedAt The date and time the plugin limit was created.
	CreatedAt time.Time `json:"created_at"`

	// PluginKind The kind of plugin, ie. source or destination.
	PluginKind PluginKind `json:"plugin_kind"`

	// PluginName The unique name for the plugin.
	PluginName PluginName `json:"plugin_name"`

	// PluginTeam The unique name for the team.
	PluginTeam TeamName `json:"plugin_team"`

	// UpdatedAt The date and time the plugin limit was last updated.
	UpdatedAt time.Time `json:"updated_at"`

	// Usd The maximum USD amount the plugin is allowed to use within a calendar month.
	USD int `json:"usd"`
}

// MonthlyLimitCreate A configurable monthly limit for plugin usage.
type MonthlyLimitCreate struct {
	// PluginKind The kind of plugin, ie. source or destination.
	PluginKind PluginKind `json:"plugin_kind"`

	// PluginName The unique name for the plugin.
	PluginName PluginName `json:"plugin_name"`

	// PluginTeam The unique name for the team.
	PluginTeam TeamName `json:"plugin_team"`

	// Usd The maximum USD amount the plugin is allowed to use within a calendar month.
	USD int `json:"usd"`
}

// MonthlyLimitUpdate A configurable monthly limit for plugin usage.
type MonthlyLimitUpdate struct {
	// Usd The maximum USD amount the plugin is allowed to use within a calendar month.
	USD int `json:"usd"`
}

// Plugin CloudQuery Plugin
type Plugin struct {
	// Category Supported categories for plugins
	Category  PluginCategory `json:"category"`
	CreatedAt time.Time      `json:"created_at"`

	// DisplayName The plugin's display name
	DisplayName string `json:"display_name"`

	// FreeRowsPerMonth The number of rows that can be synced for free each month.
	FreeRowsPerMonth int64   `json:"free_rows_per_month"`
	Homepage         *string `json:"homepage,omitempty"`

	// Kind The kind of plugin, ie. source or destination.
	Kind PluginKind `json:"kind"`
	Logo string     `json:"logo"`

	// Name The unique name for the plugin.
	Name PluginName `json:"name"`

	// Official True if the plugin is maintained by CloudQuery, false otherwise
	Official bool `json:"official"`

	// Public Whether the plugin is listed in the CloudQuery Hub. If false, the plugin will not be shown in the CloudQuery Hub and will only be visible to members of the plugin's team.
	Public           *bool   `json:"public,omitempty"`
	Repository       *string `json:"repository,omitempty"`
	ShortDescription string  `json:"short_description"`

	// TeamName The unique name for the team.
	TeamName TeamName `json:"team_name"`

	// Tier Supported tiers for plugins
	Tier PluginTier `json:"tier"`

	// UsdPerRow The price per row in USD. This is used to calculate the price of a sync.
	USDPerRow string `json:"usd_per_row"`
}

// PluginCategory Supported categories for plugins
type PluginCategory string

// PluginCreate defines model for PluginCreate.
type PluginCreate struct {
	// Category Supported categories for plugins
	Category PluginCategory `json:"category"`

	// DisplayName The plugin's display name, as shown in the CloudQuery Hub.
	DisplayName string `json:"display_name"`

	// FreeRowsPerMonth The number of rows that can be synced for free each month.
	FreeRowsPerMonth *int64  `json:"free_rows_per_month,omitempty"`
	Homepage         *string `json:"homepage,omitempty"`

	// Kind The kind of plugin, ie. source or destination.
	Kind PluginKind `json:"kind"`

	// Logo URL to the plugin's logo. This will be shown in the CloudQuery Hub. This must point to https://images.cloudquery.io/...
	Logo string `json:"logo"`

	// Name The unique name for the plugin.
	Name PluginName `json:"name"`

	// Public Whether the plugin is listed in the CloudQuery Hub. If false, the plugin will not be shown in the CloudQuery Hub and will only be visible to members of the team.
	Public     bool    `json:"public"`
	Repository *string `json:"repository,omitempty"`

	// ShortDescription Short description of the plugin. This will be shown in the CloudQuery Hub.
	ShortDescription string `json:"short_description"`

	// TeamName The unique name for the team.
	TeamName TeamName `json:"team_name"`

	// Tier Supported tiers for plugins
	Tier PluginTier `json:"tier"`

	// UsdPerRow The price per row in USD. This is used to calculate the price of a sync.
	USDPerRow *string `json:"usd_per_row,omitempty"`
}

// PluginDocsPage CloudQuery Plugin Documentation Page
type PluginDocsPage struct {
	// Content The content of the documentation page. Supports markdown.
	Content string `json:"content"`

	// Name The unique name for the plugin documentation page.
	Name PluginDocsPageName `json:"name"`

	// OrdinalPosition The position of the page in the documentation
	OrdinalPosition *int `json:"ordinal_position,omitempty"`

	// Title The title of the documentation page
	Title string `json:"title"`
}

// PluginDocsPageCreate CloudQuery Plugin Documentation Page
type PluginDocsPageCreate struct {
	// Content The content of the documentation page. Supports markdown.
	Content string `json:"content"`

	// Name The unique name for the plugin documentation page.
	Name PluginDocsPageName `json:"name"`

	// OrdinalPosition The position of the page in the documentation
	OrdinalPosition *int `json:"ordinal_position,omitempty"`

	// Title The title of the documentation page
	Title string `json:"title"`
}

// PluginDocsPageName The unique name for the plugin documentation page.
type PluginDocsPageName = string

// PluginKind The kind of plugin, ie. source or destination.
type PluginKind string

// PluginName The unique name for the plugin.
type PluginName = string

// PluginTable CloudQuery Plugin Table
type PluginTable struct {
	// Description Description of the table
	Description string `json:"description"`

	// IsIncremental Whether the table is incremental
	IsIncremental bool `json:"is_incremental"`

	// Name Name of the table
	Name PluginTableName `json:"name"`

	// Parent Name of the parent table, if any
	Parent *string `json:"parent,omitempty"`

	// Relations Names of the tables that depend on this table
	Relations []string `json:"relations"`

	// Title Title of the table
	Title string `json:"title"`
}

// PluginTableColumn CloudQuery Plugin Column
type PluginTableColumn struct {
	// Description Description of the column
	Description string `json:"description"`

	// IncrementalKey Whether the column is used as an incremental key
	IncrementalKey bool `json:"incremental_key"`

	// Name Name of the column
	Name string `json:"name"`

	// NotNull Whether the column is nullable
	NotNull bool `json:"not_null"`

	// PrimaryKey Whether the column is part of the primary key
	PrimaryKey bool `json:"primary_key"`

	// Type Arrow Type of the column
	Type string `json:"type"`

	// Unique Whether the column has a unique constraint
	Unique bool `json:"unique"`
}

// PluginTableCreate CloudQuery Plugin Table
type PluginTableCreate struct {
	Columns *[]PluginTableColumn `json:"columns,omitempty"`

	// Description Description of the table
	Description *string `json:"description,omitempty"`

	// IsIncremental Whether the table is incremental
	IsIncremental *bool `json:"is_incremental,omitempty"`

	// Name Name of the table
	Name PluginTableName `json:"name"`

	// Parent Name of the parent table, if any
	Parent *string `json:"parent,omitempty"`

	// Relations Names of the tables that depend on this table
	Relations *[]string `json:"relations,omitempty"`

	// Title Title of the table
	Title *string `json:"title,omitempty"`
}

// PluginTableDetails defines model for PluginTableDetails.
type PluginTableDetails struct {
	// Columns List of columns
	Columns []PluginTableColumn `json:"columns"`

	// Description Description of the table
	Description string `json:"description"`

	// IsIncremental Whether the table is incremental
	IsIncremental bool `json:"is_incremental"`

	// Name Name of the table
	Name string `json:"name"`

	// Parent Name of the parent table, if any
	Parent *string `json:"parent,omitempty"`

	// Relations Names of the tables that depend on this table
	Relations []string `json:"relations"`

	// Title Title of the table
	Title string `json:"title"`
}

// PluginTableName Name of the table
type PluginTableName = string

// PluginTier Supported tiers for plugins
type PluginTier string

// PluginUpdate defines model for PluginUpdate.
type PluginUpdate struct {
	// Category Supported categories for plugins
	Category *PluginCategory `json:"category,omitempty"`

	// DisplayName The plugin's display name, as shown in the CloudQuery Hub.
	DisplayName *string `json:"display_name,omitempty"`

	// FreeRowsPerMonth The number of rows that can be synced for free each month.
	FreeRowsPerMonth *int64  `json:"free_rows_per_month,omitempty"`
	Homepage         *string `json:"homepage,omitempty"`

	// Logo URL to the plugin's logo. This will be shown in the CloudQuery Hub. This must point to https://images.cloudquery.io/...
	Logo *string `json:"logo,omitempty"`

	// Public If plugin is not public, it won't be visible to other teams in the CloudQuery Hub.
	Public     *bool   `json:"public,omitempty"`
	Repository *string `json:"repository,omitempty"`

	// ShortDescription Short description of the plugin. This will be shown in the CloudQuery Hub.
	ShortDescription *string `json:"short_description,omitempty"`

	// Tier Supported tiers for plugins
	Tier *PluginTier `json:"tier,omitempty"`

	// UsdPerRow The price per row in USD. This is used to calculate the price of a sync.
	USDPerRow *string `json:"usd_per_row,omitempty"`
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

	// PublishedAt The date and time the plugin version was set to non-draft (published).
	PublishedAt *time.Time `json:"published_at,omitempty"`

	// Retracted If a plugin version is retracted, assets will still be available for download, but the version will be marked as retracted to discourage use.
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

	// Retracted If a plugin version is retracted, assets will still be available for download, but the version will be marked as retracted to discourage use.
	Retracted        *bool     `json:"retracted,omitempty"`
	SupportedTargets *[]string `json:"supported_targets,omitempty"`
}

// ReleaseURL defines model for ReleaseURL.
type ReleaseURL struct {
	Url string `json:"url"`
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

// UsageCurrent The usage of a plugin within the current calendar month.
type UsageCurrent struct {
	// PluginKind The kind of plugin, ie. source or destination.
	PluginKind PluginKind `json:"plugin_kind"`

	// PluginName The unique name for the plugin.
	PluginName PluginName `json:"plugin_name"`

	// PluginTeam The unique name for the team.
	PluginTeam TeamName `json:"plugin_team"`

	// RemainingRows The estimated number of rows remaining in the plugin's quota for the calendar month at the current price per row. This includes both free and paid rows up to the monthly limit defined for the plugin.
	RemainingRows *int64 `json:"remaining_rows,omitempty"`

	// RemainingUsd The remaining USD amount in the plugin's quota for the calendar month.
	RemainingUSD *string `json:"remaining_usd,omitempty"`

	// Rows The number of rows used by the plugin in the calendar month.
	Rows int64 `json:"rows"`

	// Usd The USD amount used by the plugin in the calendar month, rounded to two decimal places.
	USD string `json:"usd"`
}

// UsageIncrease Increase the usage of a plugin. This can incur billing costs and should be used only by plugins.
type UsageIncrease struct {
	// PluginKind The kind of plugin, ie. source or destination.
	PluginKind PluginKind `json:"plugin_kind"`

	// PluginName The unique name for the plugin.
	PluginName PluginName `json:"plugin_name"`

	// PluginTeam The unique name for the team.
	PluginTeam TeamName `json:"plugin_team"`

	// RequestId A unique ID associated with the usage increase.
	RequestId openapi_types.UUID `json:"request_id"`

	// Rows The additional rows used by the plugin.
	Rows int `json:"rows"`
}

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

// VersionName The version in semantic version format.
type VersionName = string

// APIKeyPathName defines model for apikey_name.
type APIKeyPathName = string

// IncludeDrafts defines model for include_drafts.
type IncludeDrafts = bool

// IncludePrivate defines model for include_private.
type IncludePrivate = bool

// Page defines model for page.
type Page = int64

// PerPage defines model for per_page.
type PerPage = int64

// PluginSortBy defines model for plugin_sort_by.
type PluginSortBy string

// PluginTeam The unique name for the team.
type PluginTeam = TeamName

// TargetName defines model for target_name.
type TargetName = string

// VersionSortBy defines model for version_sort_by.
type VersionSortBy string

// BadRequest defines model for BadRequest.
type BadRequest = FieldError

// Forbidden defines model for Forbidden.
type Forbidden = FieldError

// InternalError Basic Error
type InternalError = BasicError

// MethodNotAllowed Basic Error
type MethodNotAllowed = BasicError

// NotFound Basic Error
type NotFound = BasicError

// RequiresAuthentication Basic Error
type RequiresAuthentication = BasicError

// TooManyRequests Basic Error
type TooManyRequests = BasicError

// UnprocessableEntity defines model for UnprocessableEntity.
type UnprocessableEntity = FieldError

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

	// IncludeDrafts Whether to include draft plugins
	IncludeDrafts *IncludeDrafts `form:"include_drafts,omitempty" json:"include_drafts,omitempty"`
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
	Pages []PluginDocsPageCreate `json:"pages"`
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
	Tables []PluginTableCreate `json:"tables"`
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

// ListTeamAPIKeysParams defines parameters for ListTeamAPIKeys.
type ListTeamAPIKeysParams struct {
	// PerPage The number of results per page (max 1000).
	PerPage *PerPage `form:"per_page,omitempty" json:"per_page,omitempty"`

	// Page Page number of the results to fetch
	Page *Page `form:"page,omitempty" json:"page,omitempty"`
}

// CreateTeamAPIKeyJSONBody defines parameters for CreateTeamAPIKey.
type CreateTeamAPIKeyJSONBody struct {
	ExpiresAt time.Time `json:"expires_at"`
	Name      string    `json:"name"`

	// Scope Scope of permissions for the API key. API keys are used for creating new plugin versions and downloading existing plugins
	Scope *APIKeyScope `json:"scope,omitempty"`
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
	Email openapi_types.Email             `json:"email"`
	Role  EmailTeamInvitationJSONBodyRole `json:"role"`
}

// EmailTeamInvitationJSONBodyRole defines parameters for EmailTeamInvitation.
type EmailTeamInvitationJSONBodyRole string

// AcceptTeamInvitationJSONBody defines parameters for AcceptTeamInvitation.
type AcceptTeamInvitationJSONBody struct {
	Token openapi_types.UUID `json:"token"`
}

// GetTeamMembershipsParams defines parameters for GetTeamMemberships.
type GetTeamMembershipsParams struct {
	// Page Page number of the results to fetch
	Page *Page `form:"page,omitempty" json:"page,omitempty"`

	// PerPage The number of results per page (max 1000).
	PerPage *PerPage `form:"per_page,omitempty" json:"per_page,omitempty"`
}

// ListMonthlyLimitsByTeamParams defines parameters for ListMonthlyLimitsByTeam.
type ListMonthlyLimitsByTeamParams struct {
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

	// IncludePrivate Whether to include private plugins
	IncludePrivate *IncludePrivate `form:"include_private,omitempty" json:"include_private,omitempty"`
}

// ListUsersByTeamParams defines parameters for ListUsersByTeam.
type ListUsersByTeamParams struct {
	// PerPage The number of results per page (max 1000).
	PerPage *PerPage `form:"per_page,omitempty" json:"per_page,omitempty"`

	// Page Page number of the results to fetch
	Page *Page `form:"page,omitempty" json:"page,omitempty"`
}

// UpdateCurrentUserJSONBody defines parameters for UpdateCurrentUser.
type UpdateCurrentUserJSONBody struct {
	// Name The user's name
	Name *string `json:"name,omitempty"`
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

// UpdatePluginVersionJSONRequestBody defines body for UpdatePluginVersion for application/json ContentType.
type UpdatePluginVersionJSONRequestBody = PluginVersionUpdate

// CreatePluginVersionJSONRequestBody defines body for CreatePluginVersion for application/json ContentType.
type CreatePluginVersionJSONRequestBody CreatePluginVersionJSONBody

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

// CreateTeamAPIKeyJSONRequestBody defines body for CreateTeamAPIKey for application/json ContentType.
type CreateTeamAPIKeyJSONRequestBody CreateTeamAPIKeyJSONBody

// EmailTeamInvitationJSONRequestBody defines body for EmailTeamInvitation for application/json ContentType.
type EmailTeamInvitationJSONRequestBody EmailTeamInvitationJSONBody

// AcceptTeamInvitationJSONRequestBody defines body for AcceptTeamInvitation for application/json ContentType.
type AcceptTeamInvitationJSONRequestBody AcceptTeamInvitationJSONBody

// CreateMonthlyLimitJSONRequestBody defines body for CreateMonthlyLimit for application/json ContentType.
type CreateMonthlyLimitJSONRequestBody = MonthlyLimitCreate

// UpdateMonthlyLimitJSONRequestBody defines body for UpdateMonthlyLimit for application/json ContentType.
type UpdateMonthlyLimitJSONRequestBody = MonthlyLimitUpdate

// IncreaseTeamPluginUsageJSONRequestBody defines body for IncreaseTeamPluginUsage for application/json ContentType.
type IncreaseTeamPluginUsageJSONRequestBody = UsageIncrease

// UpdateCurrentUserJSONRequestBody defines body for UpdateCurrentUser for application/json ContentType.
type UpdateCurrentUserJSONRequestBody UpdateCurrentUserJSONBody
