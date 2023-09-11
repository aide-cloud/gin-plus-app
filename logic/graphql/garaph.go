package graphql

import (
	"embed"
)

//go:embed sdl/*.graphql
var Sdl embed.FS

type Resolver struct {
}

func (r *Resolver) Ping() string {
	return "pong"
}

func NewResolver() *Resolver {
	return &Resolver{}
}
