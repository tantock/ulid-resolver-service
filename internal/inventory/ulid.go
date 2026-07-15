package inventory

import (
	"github.com/oklog/ulid/v2"
)

type InventoryUlid struct {
	ULID string
}

type UpcUlidPair struct {
	UPC  string
	ULID ulid.ULID
}
