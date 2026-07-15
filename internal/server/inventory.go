package server

import (
	"context"

	"github.com/oklog/ulid/v2"
	"github.com/tantock/ulid-resolver-service/internal/dto"
)

func (s *Server) upcToULIDHandler(ctx context.Context, input *dto.InventoryUpcInput) (*dto.InventoryUlidOutput, error) {
	ulidValue := ulid.Make()
	resp := &dto.InventoryUlidOutput{
		Body: dto.InventoryUlidOutputBody{ULID: ulidValue.String()},
	}
	return resp, nil
}
