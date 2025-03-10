// Package dao provides a Database Access Object interface to stored data.
//
// All functions require the following variables to be set in the context:
//
// * Logger: for all context-aware logging.
// * Account ID: for multi-tenancy, unless marked with UNSCOPED word.
//
// Functions marked as UNSCOPED can be safely used from contexts where there is
// exactly zero function arguments coming from an user (e.g. ID was retrieved via
// another DAO call that was scoped).
package dao

import (
	"context"

	"github.com/RHEnVision/provisioning-backend/internal/clients"
	"github.com/RHEnVision/provisioning-backend/internal/models"
)

var GetAccountDao func(ctx context.Context) AccountDao

// AccountDao represents an account (tenant)
type AccountDao interface {
	// Create is meant for integration tests, use GetOrCreateByIdentity instead.
	Create(ctx context.Context, pk *models.Account) error
	GetById(ctx context.Context, id int64) (*models.Account, error)
	GetOrCreateByIdentity(ctx context.Context, orgId string, accountNumber string) (*models.Account, error)
	GetByOrgId(ctx context.Context, orgId string) (*models.Account, error)
	List(ctx context.Context, limit, offset int64) ([]*models.Account, error)
}

var GetPubkeyDao func(ctx context.Context) PubkeyDao

// PubkeyDao represents Pubkeys (public part of ssh key pair) and corresponding Resources (uploaded pubkeys
// to specific cloud providers in specific regions).
type PubkeyDao interface {
	Create(ctx context.Context, pk *models.Pubkey) error
	Update(ctx context.Context, pk *models.Pubkey) error
	GetById(ctx context.Context, id int64) (*models.Pubkey, error)
	List(ctx context.Context, limit, offset int64) ([]*models.Pubkey, error)
	Delete(ctx context.Context, id int64) error

	UnscopedCreateResource(ctx context.Context, pkr *models.PubkeyResource) error
	UnscopedGetResourceBySourceAndRegion(ctx context.Context, pubkeyId int64, sourceId string, region string) (*models.PubkeyResource, error)
	UnscopedListResourcesByPubkeyId(ctx context.Context, pkId int64) ([]*models.PubkeyResource, error)
	UnscopedDeleteResource(ctx context.Context, id int64) error
}

var GetReservationDao func(ctx context.Context) ReservationDao

// ReservationDao represents a reservation, an abstraction of one or more background jobs with
// associated detail information different for different cloud providers (like number of vCPUs,
// instance IDs created etc).
type ReservationDao interface {
	// CreateNoop creates no operation reservation with details in a single transaction.
	CreateNoop(ctx context.Context, reservation *models.NoopReservation) error

	// CreateAWS creates AWS reservation with details in a single transaction.
	CreateAWS(ctx context.Context, reservation *models.AWSReservation) error

	// CreateAzure creates Azure reservation with details in a single transaction.
	CreateAzure(ctx context.Context, reservation *models.AzureReservation) error

	// CreateGCP creates GCP reservation with details in a single transaction.
	CreateGCP(ctx context.Context, reservation *models.GCPReservation) error

	// CreateInstance inserts instance associated to a reservation.
	CreateInstance(ctx context.Context, reservation *models.ReservationInstance) error

	// GetById returns reservation for a particular account.
	GetById(ctx context.Context, id int64) (*models.Reservation, error)

	// GetAWSById returns reservation for a particular account.
	GetAWSById(ctx context.Context, id int64) (*models.AWSReservation, error)

	// GetAzureById returns Azure reservation for a given id.
	GetAzureById(ctx context.Context, id int64) (*models.AzureReservation, error)

	// GetGCPById returns reservation for a particular account.
	GetGCPById(ctx context.Context, id int64) (*models.GCPReservation, error)

	// List returns reservation for a particular account.
	List(ctx context.Context, limit, offset int64) ([]*models.Reservation, error)

	// ListInstances returns instances associated to a reservation. UNSCOPED.
	// It currently lists all instances and not instances for a reservation, this is a TODO.
	ListInstances(ctx context.Context, reservationId int64) ([]*models.ReservationInstance, error)

	// UpdateStatus sets status field and increment step counter by addSteps. UNSCOPED.
	UpdateStatus(ctx context.Context, id int64, status string, addSteps int32) error

	// UnscopedUpdateAWSDetail updates details of the AWS reservation. UNSCOPED.
	UnscopedUpdateAWSDetail(ctx context.Context, id int64, awsDetail *models.AWSDetail) error

	// UpdateReservationIDForAWS updates AWS reservation id field. UNSCOPED.
	UpdateReservationIDForAWS(ctx context.Context, id int64, awsReservationId string) error

	// UpdateOperationNameForGCP updates GCP operation name field. UNSCOPED.
	UpdateOperationNameForGCP(ctx context.Context, id int64, gcpOperationName string) error

	// UpdateReservationInstance updates an instance with its description
	UpdateReservationInstance(ctx context.Context, reservationID int64, instance *clients.InstanceDescription) error

	// FinishWithSuccess sets Success flag. UNSCOPED.
	FinishWithSuccess(ctx context.Context, id int64) error

	// FinishWithError sets Success flag and Error flag. UNSCOPED.
	FinishWithError(ctx context.Context, id int64, errorString string) error

	// Delete deletes a reservation. Only used in tests and background cleanup job. UNSCOPED.
	Delete(ctx context.Context, id int64) error
}

var GetStatDao func(ctx context.Context) StatDao

// StatDao represents an account (tenant)
type StatDao interface {
	Get(ctx context.Context) (*models.Statistics, error)
}
