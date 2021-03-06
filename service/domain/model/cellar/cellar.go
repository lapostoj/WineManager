package cellar

import (
	"time"

	"cloud.google.com/go/datastore"
)

// Cellar defines the cellar object for our domain.
type Cellar struct {
	Key          *datastore.Key `datastore:"__key__"`
	Name         string
	AccountID    int
	CreationTime time.Time
}

// NewCellar creates a Cellar struct with default values.
func NewCellar(name string, accountID int) *Cellar {
	return &Cellar{
		Name:         name,
		AccountID:    accountID,
		CreationTime: time.Now().UTC(),
	}
}
