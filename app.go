package main

import (
	"context"
	"fmt"
	"strconv"
)

// App struct
type App struct {
	ctx context.Context
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
	return fmt.Sprintf("Pixel density: %s", name)
}

func (a *App) CalcPixDens(sensorX string, sensorY string, pixelWidth string) string {
	sX, err := strconv.ParseFloat(sensorX, 64)
	pW, err2 := strconv.ParseFloat(pixelWidth, 64)

	if (err == nil) && (err2 == nil) {

		result := sX / pW * 1000
		return fmt.Sprintf("%.2f", result)
	}

	return "Error"
}

func (a *App) CalcSensorRatio(sensorX string, sensorY string) string {
	sX, err := strconv.ParseFloat(sensorX, 64)
	sY, err2 := strconv.ParseFloat(sensorY, 64)

	if (err == nil) && (err2 == nil) {

		result := sX / sY
		return fmt.Sprintf("%.2f", result)
	}

	return "Error"
}
