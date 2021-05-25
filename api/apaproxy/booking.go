package apaproxy

import (
	"context"
	"fmt"
	"net/http"

	sw "github.com/winnix/go-apaleo-one-boilerplate/api/clients/bookingclient"
)

const (
	apiHost         = "api.apaleo.com"
	defaultPageSize = 10
)

type BookingClient interface {
	GetReservations() ([]sw.ReservationItemModel, error)
}

type bookingClient struct {
	authCtx   context.Context
	apiClient *sw.APIClient
}

func newBookingClient(ctx context.Context, accessToken string) BookingClient {
	cfg := sw.NewConfiguration()
	cfg.Host = apiHost
	cfg.Scheme = "https"

	return &bookingClient{
		authCtx:   context.WithValue(ctx, sw.ContextAccessToken, accessToken),
		apiClient: sw.NewAPIClient(cfg),
	}
}

func (bc *bookingClient) GetReservations() ([]sw.ReservationItemModel, error) {
	req := bc.apiClient.ReservationApi.BookingReservationsGet(bc.authCtx)
	req = req.PageSize(defaultPageSize)
	items, res, err := req.Execute()
	if err != nil && res.StatusCode != http.StatusNoContent {
		return nil, fmt.Errorf("apaleo api call failed %s: %s", res.Status, err.Error())
	}

	if res.StatusCode == http.StatusNoContent {
		return []sw.ReservationItemModel{}, nil
	}

	return items.Reservations, nil
}
