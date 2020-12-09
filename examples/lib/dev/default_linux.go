package dev

import (
	"github.com/marcgeld/ble"
	"github.com/marcgeld/ble/linux"
)

// DefaultDevice ...
func DefaultDevice(opts ...ble.Option) (d ble.Device, err error) {
	return linux.NewDevice(opts...)
}
