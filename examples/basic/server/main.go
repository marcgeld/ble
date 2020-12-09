package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/marcgeld/ble/darwin"
	"log"
	"time"

	"github.com/marcgeld/ble"
	"github.com/marcgeld/ble/examples/lib"
	"github.com/marcgeld/ble/examples/lib/dev"
	"github.com/pkg/errors"
)

var (
	device = flag.String("device", "default", "implementation of ble")
	du     = flag.Duration("du", 500*time.Second, "advertising duration, 0 for indefinitely")
)

func main() {
	flag.Parse()

	darwin.Logger.SetLevel(1000)

	d, err := dev.NewDevice(*device)
	if err != nil {
		log.Fatalf("can't new device : %s", err)
	}
	ble.SetDefaultDevice(d)

	testSvc := ble.NewService(lib.TestSvcUUID)
	testSvc.AddCharacteristic(lib.NewCountChar())
	testSvc.AddCharacteristic(lib.NewEchoChar())

	if err := ble.AddService(testSvc); err != nil {
		log.Fatalf("can't add service: %s", err)
	}
	ble.RemoveAllServices()
	// Advertise for specified durantion, or until interrupted by user.
	fmt.Printf("Advertising for %s...\n", *du)
	ctx := ble.WithSigHandler(context.WithTimeout(context.Background(), *du))
	chkErr(ble.AdvertiseNameAndServices(ctx, "Gopher", testSvc.UUID))
}

func chkErr(err error) {
	switch errors.Cause(err) {
	case nil:
	case context.DeadlineExceeded:
		fmt.Printf("done\n")
	case context.Canceled:
		fmt.Printf("canceled\n")
	default:
		log.Fatalf(err.Error())
	}
}
