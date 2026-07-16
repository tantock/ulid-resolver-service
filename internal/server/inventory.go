package server

import (
	"context"
	"fmt"

	"github.com/oklog/ulid/v2"
	"github.com/tantock/ulid-resolver-service/internal/dto"
	"github.com/tantock/ulid-resolver-service/internal/inventory"
)

func (s *Server) idToULIDHandler(ctx context.Context, input *dto.ProductIdInput) (*dto.InventoryUlidOutput, error) {
	if input.Id == "" {
		return nil, fmt.Errorf("Empty id input")
	}
	var selectedUlid *inventory.InventoryUlid
	var dbErr error
	switch input.Type {
	case dto.IdUpc:
		selectedUlid, dbErr = s.db.SelectUlidFromUpc(input.Id)
		if dbErr != nil {
			return nil, dbErr
		}
	default:
		return nil, fmt.Errorf("unknown type specified: %v", input.Type)
	}

	var ulidStr = ""
	if selectedUlid != nil {
		ulidStr = selectedUlid.ULID
	} else {
		upcProduct := inventory.NewUpcProduct(input.Id, ulid.Make())
		err := s.db.InsertProduct(upcProduct)
		if err != nil {
			return nil, err
		}
	}
	resp := &dto.InventoryUlidOutput{
		Body: dto.InventoryUlidOutputBody{ULID: ulidStr},
	}
	return resp, nil
}
