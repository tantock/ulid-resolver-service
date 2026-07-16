package dto

type IdType string

const (
	IdUpc IdType = "upc"
)

type ProductIdInput struct {
	Id   string `path:"id" example:"8801052435022" doc:"Product Identifier"`
	Type IdType `query:"type" required:"true"`
}
