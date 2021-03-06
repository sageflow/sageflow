package resource

import (
	"github.com/gigamono/gigamono/pkg/database/models"
	"github.com/gofrs/uuid"
)

// Workspace represents a workspace.
type Workspace struct {
	models.Base
	Name                            string        `json:"name"`
	AvatarURL                       string        `json:"avatar_url"`
	CreatorID                       uuid.UUID     `pg:"type:uuid" json:"creator_id"`
	Spaces                          []Space       `pg:"rel:has-many" json:"spaces"`
	XUserWorkspaceMemberships       []User        `pg:"many2many:x_user_workspace_memberships" json:"-"`
	XWorkspaceInstalledIntegrations []Integration `pg:"many2many:x_workspace_installed_integrations" json:"-"`
}
