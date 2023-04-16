package main

import (
	"context"
	"time"

	"github.com/a-clap/distillation-gui/backend/heater"
	"github.com/a-clap/distillation-gui/backend/parameters"
	"github.com/a-clap/iot/pkg/distillation"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	const DistilAddr = "http://localhost:8081"
	heater.Init(distillation.NewHeaterClient(DistilAddr, 1*time.Second))
	a := &App{
		ctx: context.Background(),
	}
	return a
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) GetHeaters() []parameters.Heater {
	return heater.Get()
}

func (a *App) EnableGlobal(id string, enable bool) string {
	if err := heater.EnableGlobal(id, enable); err != nil {
		return err.Error()
	}
	return ""

}
