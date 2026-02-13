package main

import (
	"github.com/hilmoo/go-backend-example/internal/cmd"
)

// @title Go Backend Example API
// @version 0.0.1
// @description This is a sample server for a Go backend example application.

// @BasePath /api/
func main() {
	cmd.RootCmd()
}