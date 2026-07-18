package arona

import (
	"context"
	"fmt"

	"github.com/arisu-archive/arona-protos/protos"
)

type ShopService service

type ShopBuyMerchandiseRequestWrapper struct {
	*protos.ShopBuyMerchandiseRequest
}

func (w ShopBuyMerchandiseRequestWrapper) Packet() *protos.RequestPacket {
	return &w.RequestPacket
}

type ShopMerchandise struct {
	IsRefreshGoods bool
	ShopUniqueID   int64
	GoodsID        int64
	PurchaseCount  int64
}

func (s *ShopService) BuyMerchandise(ctx context.Context, session *UserSession, m ShopMerchandise) (*protos.ShopBuyMerchandiseResponse, error) {
	param := ShopBuyMerchandiseRequestWrapper{
		&protos.ShopBuyMerchandiseRequest{
			IsRefreshGoods: m.IsRefreshGoods,
			ShopUniqueId:   m.ShopUniqueID,
			GoodsId:        m.GoodsID,
			PurchaseCount:  m.PurchaseCount,
		},
	}
	req, err := s.client.R().WithSession(session).Game(ctx, protos.Protocol_Cafe_SummonCharacterTicketUse, param)
	if err != nil {
		return nil, fmt.Errorf("failed to create get cafe info request: %w", err)
	}
	result := new(protos.ShopBuyMerchandiseResponse)
	_, err = s.client.Do(ctx, req, result)
	if err != nil {
		return nil, fmt.Errorf("get cafe info request failed: %w", err)
	}
	return result, nil
}
