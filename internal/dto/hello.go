package dto

type HelloWorldOutput struct {
	Body HelloWorldOutputBody
}

type HelloWorldOutputBody struct {
	Message string `json:"message" example:"Hello World!" doc:"Hello World!"`
}
