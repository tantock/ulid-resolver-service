package dto

type InventoryUpcInput struct {
	UPC string `path:"upc" maxLength:"12" example:"8801052435022" doc:"Universal Product Code"`
}

type InventoryUlidOutput struct {
	Body InventoryUlidOutputBody
}

type InventoryUlidOutputBody struct {
	ULID string `json:"ulid" example:"01G65Z755AFWAKHE12NY0CQ9FH" doc:"Universally Unique Lexicographically Sortable Identifier"`
}
