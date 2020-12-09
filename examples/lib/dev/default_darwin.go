package dev

import (
	"github.com/marcgeld/ble"
	"github.com/marcgeld/ble/darwin"
)

// DefaultDevice ...
func DefaultDevice(opts ...ble.Option) (d ble.Device, err error) {
	return darwin.NewDevice(opts...)
}
