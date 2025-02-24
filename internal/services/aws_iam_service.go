package services

import (
	"net/http"

	"github.com/RHEnVision/provisioning-backend/internal/clients"
	"github.com/RHEnVision/provisioning-backend/internal/config"
	"github.com/RHEnVision/provisioning-backend/internal/models"
	"github.com/RHEnVision/provisioning-backend/internal/payloads"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/rs/zerolog"
)

func ValidatePermissions(w http.ResponseWriter, r *http.Request) {
	logger := zerolog.Ctx(r.Context())
	sourceId := chi.URLParam(r, "ID")
	region := r.URL.Query().Get("region")

	if region == "" {
		region = config.AWS.DefaultRegion
	}

	// Get Sources client
	sourcesClient, err := clients.GetSourcesClient(r.Context())
	if err != nil {
		renderError(w, r, payloads.NewClientError(r.Context(), err))
		return
	}

	// Fetch arn from Sources
	authentication, err := sourcesClient.GetAuthentication(r.Context(), sourceId)
	if err != nil {
		renderError(w, r, payloads.NewClientError(r.Context(), err))
		return
	}

	if !authentication.Is(models.ProviderTypeAWS) {
		if err = render.Render(w, r, payloads.NewPermissionsResponse(nil)); err != nil {
			renderError(w, r, payloads.NewRenderError(r.Context(), "unable to render missing permissions", err))
			return
		}
		return
	}

	ec2Client, err := clients.GetEC2Client(r.Context(), authentication, region)
	if err != nil {
		renderError(w, r, payloads.NewAWSError(r.Context(), "unable to get AWS EC2 client", err))
		return
	}

	logger.Info().Msgf("Listing permissions.")
	permissions, err := ec2Client.CheckPermission(r.Context(), authentication)
	if err != nil {
		renderError(w, r, payloads.NewAWSError(r.Context(), "unable to check aws permissions", err))
		return
	}

	if err := render.Render(w, r, payloads.NewPermissionsResponse(permissions)); err != nil {
		renderError(w, r, payloads.NewRenderError(r.Context(), "unable to render missing permissions", err))
		return
	}
}
