package main

import (
	"github.com/tantock/ulid-resolver-service/internal/server"
)

func main() {

	server := server.NewServer()
	server.Run()
}
