package inventory

import (
	"github.com/oklog/ulid/v2"
	"github.com/tantock/ulid-resolver-service/internal/dto"
)

type Product interface {
	Id() string
	IdType() dto.IdType
}

type product struct {
	id     string
	idType dto.IdType
	ulid   ulid.ULID
}

// Id implements [Product].
func (p *product) Id() string {
	return p.id
}

// IdType implements [Product].
func (p *product) IdType() dto.IdType {
	panic("unimplemented")
}

func NewUpcProduct(UPC string, ULID ulid.ULID) Product {
	product := product{id: UPC, ulid: ULID, idType: dto.IdUpc}
	return &product
}
