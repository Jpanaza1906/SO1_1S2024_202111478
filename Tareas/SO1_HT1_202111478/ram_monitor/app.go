package main

import (
	"context"
	"fmt"
	"os/exec"
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
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

// GetRAMData from de Module that is reading the ram in the file /proc/ram_module
func (a *App) GetRAMData() string {
	fmt.Println("Getting RAM data")
	fmt.Println("")

	// Get the RAM data from the module
	cmd := exec.Command("sh", "-c", "cat /proc/ram_module")
	stdout, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
		return ""
	}

	// Print the RAM data
	output := string(stdout[:])
	fmt.Println(output)

	return output
}