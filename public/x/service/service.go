package service

import (
	"context"

	"github.com/Jeffail/benthos/v3/lib/service"
)

// RunCLI executes Benthos as a CLI, allowing users to specify a configuration
// file path(s) and execute subcommands for linting configs, testing configs,
// etc. This is how a standard distribution of Benthos operates.
//
// This call blocks until either:
// - The service shuts down gracefully due to the inputs closing
// - A termination signal is received
// - The provided context is cancelled
//
// This function must only be called once during the entire lifecycle of your
// program, as it interacts with singleton state. In order to manage multiple
// Benthos stream lifecycles in a program use the StreamBuilder API instead.
func RunCLI(ctx context.Context) {
	service.RunWithOpts(service.OptUseContext(ctx))
}
