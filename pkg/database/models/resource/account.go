package resource

import (
	"github.com/gofrs/uuid"
	"github.com/sageflow/sageflow/pkg/database/models"
)

// Account represents credentials needed for an app to authorize access.
type Account struct {
	models.Base
	UserID            uuid.UUID
	AccessTokenCredID uuid.UUID `gorm:"unique; type:uuid"`
	XApp              []*App    `gorm:"many2many:apps_x_accounts"`
}