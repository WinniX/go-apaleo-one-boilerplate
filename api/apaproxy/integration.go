package apaproxy

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	sw "github.com/winnix/go-apaleo-one-boilerplate/api/clients/integrationclient"
)

const (
	uiIntegrationLabel      = "My cool app"
	uiIntegrationSourceType = "Public"
	uiIntegrationTarget     = "AccountMenuApps"

	integrationHost = "integration.apaleo.com"
)

type IntegrationClient interface {
	CreateIntegration(homeUrl string) (string, error)
}

type integrationClient struct {
	authCtx   context.Context
	apiClient *sw.APIClient
}

func newIntegrationClient(ctx context.Context, accessToken string) IntegrationClient {
	cfg := sw.NewConfiguration()
	cfg.Host = integrationHost
	cfg.Scheme = "https"

	return &integrationClient{
		authCtx:   context.WithValue(ctx, sw.ContextAccessToken, accessToken),
		apiClient: sw.NewAPIClient(cfg),
	}
}

func addIntegration(c *sw.APIClient, authCtx context.Context, homeUrl string) (string, error) {
	req := c.UiIntegrationsApi.IntegrationUiIntegrationsByTargetPost(authCtx, uiIntegrationTarget)
	model := sw.NewCreateUiIntegrationModel(uiIntegrationLabel, homeUrl, uiIntegrationSourceType)
	req = req.Body(*model)
	result, _, err := req.Execute()
	if err != nil {
		return "", fmt.Errorf("add integration failed: %s", err.Error())
	}

	return result.Id, nil
}

func (ic *integrationClient) CreateIntegration(homeUrl string) (string, error) {
	items, res, err := ic.apiClient.UiIntegrationsApi.
		IntegrationUiIntegrationsByTargetGet(ic.authCtx, uiIntegrationTarget).
		Execute()
	if err != nil && res.StatusCode != http.StatusNoContent {
		return "", fmt.Errorf("apaleo api call failed %s: %s", res.Status, err.Error())
	}

	if items.Count == 0 {
		id, err := addIntegration(ic.apiClient, ic.authCtx, homeUrl)
		if err != nil {
			return "", fmt.Errorf("apaleo api call failed: %s", err.Error())
		}

		return id, nil
	}

	for _, item := range items.UiIntegrations {
		if strings.Compare(item.SourceUrl, homeUrl) == 0 {
			return item.Id, nil
		}
	}

	id, err := addIntegration(ic.apiClient, ic.authCtx, homeUrl)
	if err != nil {
		return "", fmt.Errorf("apaleo api call failed: %s", err.Error())
	}

	return id, nil
}
