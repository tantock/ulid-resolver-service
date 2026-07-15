package dto

type ProductIdInput struct {
	Id   string `path:"id" example:"8801052435022" doc:"Product Identifier"`
	Type string `query:"type"`
}
