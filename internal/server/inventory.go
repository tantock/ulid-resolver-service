package server

import (
	"context"

	"github.com/oklog/ulid/v2"
	"github.com/tantock/ulid-resolver-service/internal/dto"
	"github.com/tantock/ulid-resolver-service/internal/inventory"
)

func (s *Server) upcToULIDHandler(ctx context.Context, input *dto.InventoryUpcInput) (*dto.InventoryUlidOutput, error) {
	selectedUlid, err := s.db.SelectUlidFromUpc(input.Body.UPC)
	if err != nil {
		return nil, err
	}
	var ulidStr = ""
	if selectedUlid != nil {
		ulidStr = selectedUlid.ULID
	} else {
		ULID := ulid.Make()
		ulidStr = ULID.String()
		err = s.db.InsertUpc(inventory.UpcUlidPair{UPC: input.Body.UPC, ULID: ULID})
		if err != nil {
			return nil, err
		}
	}
	resp := &dto.InventoryUlidOutput{
		Body: dto.InventoryUlidOutputBody{ULID: ulidStr},
	}
	return resp, nil
}
