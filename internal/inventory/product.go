package inventory

import (
	"github.com/oklog/ulid/v2"
)

type Product struct {
	ID   string
	ULID ulid.ULID
}
