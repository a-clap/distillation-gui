package main

import (
	"context"
	"fmt"
)

// App struct
type App struct {
	ctx context.Context
}

type Value struct {
	Name string  `json:"name"`
	V    int     `json:"value"`
	F    float64 `json:"floatV"`
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) GetValues() []Value {
	vs := make([]Value, 3)
	for i := range vs {
		vs[i].Name = fmt.Sprint("name ", i)
		vs[i].V = i
	}
	return vs
}
