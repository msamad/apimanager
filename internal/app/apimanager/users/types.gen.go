// Package users provides primitives to interact the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen DO NOT EDIT.
package users

import (
	"time"
)

// ApiSummaryBean defines model for ApiSummaryBean.
type ApiSummaryBean struct {
	CreatedOn        *time.Time `json:"createdOn,omitempty"`
	Description      *string    `json:"description,omitempty"`
	Id               *string    `json:"id,omitempty"`
	Name             *string    `json:"name,omitempty"`
	OrganizationId   *string    `json:"organizationId,omitempty"`
	OrganizationName *string    `json:"organizationName,omitempty"`
}

// AuditEntryBean defines model for AuditEntryBean.
type AuditEntryBean struct {
	CreatedOn      *time.Time `json:"createdOn,omitempty"`
	Data           *string    `json:"data,omitempty"`
	EntityId       *string    `json:"entityId,omitempty"`
	EntityType     *string    `json:"entityType,omitempty"`
	EntityVersion  *string    `json:"entityVersion,omitempty"`
	Id             *int64     `json:"id,omitempty"`
	OrganizationId *string    `json:"organizationId,omitempty"`
	What           *string    `json:"what,omitempty"`
	Who            *string    `json:"who,omitempty"`
}

// ClientSummaryBean defines model for ClientSummaryBean.
type ClientSummaryBean struct {
	Description      *string `json:"description,omitempty"`
	Id               *string `json:"id,omitempty"`
	Name             *string `json:"name,omitempty"`
	NumContracts     *int32  `json:"numContracts,omitempty"`
	OrganizationId   *string `json:"organizationId,omitempty"`
	OrganizationName *string `json:"organizationName,omitempty"`
}

// CurrentUserBean defines model for CurrentUserBean.
type CurrentUserBean struct {
	Admin       *bool             `json:"admin,omitempty"`
	Email       *string           `json:"email,omitempty"`
	FullName    *string           `json:"fullName,omitempty"`
	JoinedOn    *time.Time        `json:"joinedOn,omitempty"`
	Permissions *[]PermissionBean `json:"permissions,omitempty"`
	Username    *string           `json:"username,omitempty"`
}

// OrganizationSummaryBean defines model for OrganizationSummaryBean.
type OrganizationSummaryBean struct {
	Description *string `json:"description,omitempty"`
	Id          *string `json:"id,omitempty"`
	Name        *string `json:"name,omitempty"`
	NumApis     *int32  `json:"numApis,omitempty"`
	NumClients  *int32  `json:"numClients,omitempty"`
	NumMembers  *int32  `json:"numMembers,omitempty"`
}

// PermissionBean defines model for PermissionBean.
type PermissionBean struct {
	Name           *string `json:"name,omitempty"`
	OrganizationId *string `json:"organizationId,omitempty"`
}

// SearchResultsBeanAuditEntryBean defines model for SearchResultsBeanAuditEntryBean.
type SearchResultsBeanAuditEntryBean struct {
	Beans     *[]AuditEntryBean `json:"beans,omitempty"`
	TotalSize *int32            `json:"totalSize,omitempty"`
}

// UpdateUserBean defines model for UpdateUserBean.
type UpdateUserBean struct {
	Email    *string `json:"email,omitempty"`
	FullName *string `json:"fullName,omitempty"`
}

// UserBean defines model for UserBean.
type UserBean struct {
	Admin    *bool      `json:"admin,omitempty"`
	Email    *string    `json:"email,omitempty"`
	FullName *string    `json:"fullName,omitempty"`
	JoinedOn *time.Time `json:"joinedOn,omitempty"`
	Username *string    `json:"username,omitempty"`
}

// UserPermissionsBean defines model for UserPermissionsBean.
type UserPermissionsBean struct {
	Permissions *[]PermissionBean `json:"permissions,omitempty"`
	UserId      *string           `json:"userId,omitempty"`
}

// Update3JSONBody defines parameters for Update3.
type Update3JSONBody UpdateUserBean

// GetActivity1Params defines parameters for GetActivity1.
type GetActivity1Params struct {
	Page  *int32 `json:"page,omitempty"`
	Count *int32 `json:"count,omitempty"`
}

// Update3RequestBody defines body for Update3 for application/json ContentType.
type Update3JSONRequestBody Update3JSONBody

