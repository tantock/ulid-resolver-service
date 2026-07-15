package dto

type InventoryUlidOutput struct {
	Body InventoryUlidOutputBody
}

type InventoryUlidOutputBody struct {
	ULID string `json:"ulid" example:"01G65Z755AFWAKHE12NY0CQ9FH" doc:"Universally Unique Lexicographically Sortable Identifier"`
}
