package payloads

import (
	"net/http"

	"github.com/RHEnVision/provisioning-backend/internal/clients"
	"github.com/go-chi/render"
)

type InstanceTypeResponse clients.InstanceType

func (s *InstanceTypeResponse) Bind(_ *http.Request) error {
	return nil
}

func (s *InstanceTypeResponse) Render(_ http.ResponseWriter, _ *http.Request) error {
	return nil
}

func NewListInstanceTypeResponse(sl []*clients.InstanceType) []render.Renderer {
	list := make([]render.Renderer, len(sl))
	for i, it := range sl {
		list[i] = &InstanceTypeResponse{
			Name:               it.Name,
			VCPUs:              it.VCPUs,
			Cores:              it.Cores,
			MemoryMiB:          it.MemoryMiB,
			EphemeralStorageGB: it.EphemeralStorageGB,
			Supported:          it.Supported,
			Architecture:       it.Architecture,
			AzureDetail:        it.AzureDetail,
		}
	}
	return list
}
