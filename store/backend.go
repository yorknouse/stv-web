package store

import (
	"github.com/yorknouse/stv-web/storage"
)

type Backend interface {
	Read() (*storage.STV, error)
	Write(state *storage.STV) error
}
