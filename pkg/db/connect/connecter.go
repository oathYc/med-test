package connect

import (
	"hello/pkg/db/config"
	"hello/pkg/db/wrapper"
)

type Connecter interface {
	Connect(config *config.Config) (*wrapper.Wrapper, error)
}