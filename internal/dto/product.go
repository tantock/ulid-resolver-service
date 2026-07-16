package dto

import "github.com/danielgtaylor/huma/v2"

type IdType string

const (
	IdUpc IdType = "upc"
)

func (IdType) TransformSchema(r huma.Registry, s *huma.Schema) *huma.Schema {
	s.Enum = []any{
		string(IdUpc),
	}
	return s
}

type ProductIdInput struct {
	Id   string `path:"id" example:"8801052435022" doc:"Product Identifier"`
	Type IdType `query:"type" required:"true"`
}
