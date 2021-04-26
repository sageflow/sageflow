package resource

import (
	"github.com/gofrs/uuid"
	"github.com/gigamono/gigamono/pkg/database/models"
)

// Log represents a log information.
// Log does not use foreign key constraints because it is supposed to exist even after associated keys are removed.
type Log struct {
	models.Base
	UserID             uuid.UUID
	EngineID           uuid.UUID
	WorkflowID         uuid.UUID
	WorkflowInstanceID uuid.UUID
	Message            string
	Level              string
}
