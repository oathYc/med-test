package wrapper

import (
	"math/rand"
	"time"

	"hello/pkg/db/builder"
)

type Wrapper struct {
	Dsn builder.IBuilder
}

var random = rand.New(rand.NewSource(time.Now().Unix()))

func (db *Wrapper) Write() builder.IBuilder {
	return db.Dsn
}

func (db *Wrapper) Read() builder.IBuilder {
	return db.Dsn
}
