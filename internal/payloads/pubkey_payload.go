package payloads

import (
	"net/http"

	"github.com/RHEnVision/provisioning-backend/internal/models"

	"github.com/go-chi/render"
)

// See models.Pubkey
type PubkeyRequest struct {
	Name string `json:"name" yaml:"name"`
	Body string `json:"body" yaml:"body"`
}

// See models.Pubkey
type PubkeyResponse struct {
	ID                int64  `json:"id" yaml:"id"`
	AccountID         int64  `json:"-" yaml:"-"`
	Name              string `json:"name" yaml:"name"`
	Body              string `json:"body" yaml:"body"`
	Type              string `json:"type,omitempty" yaml:"type,omitempty"`
	Fingerprint       string `json:"fingerprint,omitempty" yaml:"fingerprint,omitempty"`
	FingerprintLegacy string `json:"fingerprint_legacy,omitempty" yaml:"fingerprint_legacy,omitempty"`
}

func (p *PubkeyRequest) Bind(_ *http.Request) error {
	return nil
}

func (p *PubkeyResponse) Render(_ http.ResponseWriter, _ *http.Request) error {
	return nil
}

func (p *PubkeyRequest) NewModel() *models.Pubkey {
	return &models.Pubkey{
		Name: p.Name,
		Body: p.Body,
	}
}

func NewPubkeyResponse(pubkey *models.Pubkey) render.Renderer {
	return &PubkeyResponse{
		ID:                pubkey.ID,
		AccountID:         pubkey.AccountID,
		Name:              pubkey.Name,
		Body:              pubkey.Body,
		Type:              pubkey.Type,
		Fingerprint:       pubkey.Fingerprint,
		FingerprintLegacy: pubkey.FingerprintLegacy,
	}
}

func NewPubkeyListResponse(pubkeys []*models.Pubkey) []render.Renderer {
	list := make([]render.Renderer, len(pubkeys))
	for i, pubkey := range pubkeys {
		list[i] = NewPubkeyResponse(pubkey)
	}
	return list
}
