package facade

import "errors"

var (
	errNilSimulatorHandler         = errors.New("nil simulator handler ")
	errNilProxyTransactionsHandler = errors.New("nil proxy transactions handler ")
)
