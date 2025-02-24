package kafka

import (
	"context"

	"github.com/RHEnVision/provisioning-backend/internal/identity"
)

type StatusType string

const (
	StatusUnavailable StatusType = "unavailable"
	StatusAvaliable   StatusType = "available"
)

type SourceResult struct {
	ResourceID string `json:"resource_id"`

	// Resource type of the source
	ResourceType string `json:"resource_type"`

	Status StatusType `json:"status"`

	Err error `json:"error"`

	Identity identity.Principal `json:"-"`
}

func (sr SourceResult) GenericMessage(ctx context.Context) (GenericMessage, error) {
	return genericMessage(ctx, sr, sr.ResourceID, SourcesStatusTopic)
}

func (st StatusType) String() string {
	return string(st)
}
