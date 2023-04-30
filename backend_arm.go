package main

import (
	"log"
	"time"

	"github.com/a-clap/distillation-gui/backend"
	"github.com/a-clap/distillation/pkg/distillation"
	"github.com/a-clap/distillation/pkg/wifi"
)

func getopts(addr string) []backend.Option {
	heaterClient, err := distillation.NewHeaterRPCCLient(addr, time.Second)
	if err != nil {
		log.Fatalln(err)
	}
	dsClient, err := distillation.NewDSRPCClient(addr, time.Second)
	if err != nil {
		log.Fatalln(err)
	}

	ptClient, err := distillation.NewPTRPCClient(addr, time.Second)
	if err != nil {
		log.Fatalln(err)
	}
	gpioClient, err := distillation.NewGPIORPCClient(addr, time.Second)
	if err != nil {
		log.Fatalln(err)
	}
	phaseClient, err := distillation.NewProcessRPCClient(addr, time.Second)
	if err != nil {
		log.Fatalln(err)
	}

	w, err := wifi.New()
	if err != nil {
		log.Fatalln(err)
	}

	return []backend.Option{
		backend.WithHeaterClient(heaterClient),
		backend.WithDSClient(dsClient),
		backend.WithPTClient(ptClient),
		backend.WithGPIOClient(gpioClient),
		backend.WithPhaseClient(phaseClient),
		backend.WithWifi(w),
	}
}
