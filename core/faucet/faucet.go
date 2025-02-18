package faucet

import (
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"

	"github.com/iotaledger/inx-app/pkg/httpserver"
	"github.com/iotaledger/inx-faucet/pkg/faucet"
)

func getFaucetInfo(_ echo.Context) (*faucet.InfoResponse, error) {
	return deps.Faucet.Info()
}

func addFaucetOutputToQueue(c echo.Context) (*faucet.EnqueueResponse, error) {

	request := &faucetEnqueueRequest{}
	if err := c.Bind(request); err != nil {
		return nil, errors.WithMessagef(httpserver.ErrInvalidParameter, "Invalid Request! Error: %s", err)
	}

	response, err := deps.Faucet.Enqueue(request.Address)
	if err != nil {
		return nil, err
	}

	return response, nil
}
