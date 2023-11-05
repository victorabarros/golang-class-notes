package core

import (
	"context"
)

// Server holds the setup for individual servers
type Server interface {
	Name() string
	Start(ctx context.Context) error
	Shutdown(ctx context.Context) error
}
