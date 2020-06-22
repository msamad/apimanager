package gateways

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"sync"
)

type Gateway struct {
	GatewayBean map[int64]GatewayBean
	NextId int64
	Lock sync.Mutex
}

func NewGateway() *Gateway {
	return &Gateway{
		GatewayBean: make(map[int64]GatewayBean),
		NextId: 1000,
	}
}

func (g *Gateway) ListGateways(ctx echo.Context) error {
	// We're always asynchronous, so lock unsafe operations below
	g.Lock.Lock()
	defer g.Lock.Unlock()

	// Action performed
	return ctx.NoContent(http.StatusNoContent)
}

func (g *Gateway) PostGateways(ctx echo.Context) error {
	panic("implement me")
}

func (g *Gateway) PutGateways(ctx echo.Context) error {
	panic("implement me")
}

func (g *Gateway) DeleteGatewayById(ctx echo.Context, gatewayId string) error {
	panic("implement me")
}

func (g *Gateway) GetGatewayById(ctx echo.Context, gatewayId string) error {
	panic("implement me")
}

func (g *Gateway) PutGatewayById(ctx echo.Context, gatewayId string) error {
	panic("implement me")
}

func (g *Gateway) GetGatewayEndpoint1(ctx echo.Context, gatewayId string) error {
	panic("implement me")
}

