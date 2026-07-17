package arona

import (
	"context"
	"fmt"

	"github.com/arisu-archive/arona-protos/protos"
)

type CafeService service

type CafeInteractRequestWrapper struct {
	*protos.CafeInteractWithCharacterRequest
}

func (w CafeInteractRequestWrapper) Packet() *protos.RequestPacket {
	return &w.RequestPacket
}

func (c *CafeService) Interact(ctx context.Context, session *UserSession, cafeID, characterID int64) (*protos.CafeInteractWithCharacterResponse, error) {
	param := CafeInteractRequestWrapper{
		&protos.CafeInteractWithCharacterRequest{
			CafeDBId:    cafeID,
			CharacterId: characterID,
		},
	}
	req, err := c.client.R().WithSession(session).Game(ctx, protos.Protocol_Cafe_Interact, param)
	if err != nil {
		return nil, fmt.Errorf("failed to create get friend detail request: %w", err)
	}
	result := new(protos.CafeInteractWithCharacterResponse)
	_, err = c.client.Do(ctx, req, result)
	if err != nil {
		return nil, fmt.Errorf("get friend detail request failed: %w", err)
	}
	return result, nil
}

type CafeGetInfoRequestWrapper struct {
	*protos.CafeGetInfoRequest
}

func (w CafeGetInfoRequestWrapper) Packet() *protos.RequestPacket {
	return &w.RequestPacket
}

func (c *CafeService) GetInfo(ctx context.Context, session *UserSession) (*protos.CafeGetInfoResponse, error) {
	param := CafeGetInfoRequestWrapper{
		&protos.CafeGetInfoRequest{
			AccountServerId: session.AccountServerId,
		},
	}
	req, err := c.client.R().WithSession(session).Game(ctx, protos.Protocol_Cafe_Get, param)
	if err != nil {
		return nil, fmt.Errorf("failed to create get cafe info request: %w", err)
	}
	result := new(protos.CafeGetInfoResponse)
	_, err = c.client.Do(ctx, req, result)
	if err != nil {
		return nil, fmt.Errorf("get cafe info request failed: %w", err)
	}
	return result, nil
}
