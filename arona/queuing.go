package arona

import (
	"context"
	"encoding/base64"
	"fmt"

	"github.com/arisu-archive/arona-protos/protos"
)

type QueuingService service

type GetCryptoKeysOptions struct {
	KeyBundle AESKeyBundle
}

type QueuingGetCryptoKeysRequestWrapper struct {
	*protos.QueuingGetCryptoKeysRequest
}

func (p QueuingGetCryptoKeysRequestWrapper) Packet() *protos.RequestPacket {
	return &p.RequestPacket
}

func (s *QueuingService) GetCryptoKeys(ctx context.Context, data GetCryptoKeysOptions) (*protos.QueuingGetCryptoKeysResponse, error) {
	param := QueuingGetCryptoKeysRequestWrapper{
		QueuingGetCryptoKeysRequest: &protos.QueuingGetCryptoKeysRequest{
			ClientGeneratedKey: base64.StdEncoding.EncodeToString(data.KeyBundle.Key),
			ClientGeneratedIV:  base64.StdEncoding.EncodeToString(data.KeyBundle.IV),
		},
	}
	req, err := s.client.R().Gateway(ctx, protos.Protocol_Queuing_GetCryptoKeys, param, WithHash(0))
	if err != nil {
		return nil, fmt.Errorf("failed to create get crypto keys request: %w", err)
	}
	result := new(protos.QueuingGetCryptoKeysResponse)
	_, err = s.client.Do(ctx, req, result)
	if err != nil {
		return nil, fmt.Errorf("get crypto keys request failed: %w", err)
	}
	return result, nil
}

type GetTicketOptions struct {
	ClientVersion string
	NpSN          int64
	NpToken       string
	NpaCode       string
	NgsmToken     string
}

type QueuingGetTicketRequestWrapper struct {
	*protos.QueuingGetTicketRequest
}

func (p QueuingGetTicketRequestWrapper) Packet() *protos.RequestPacket {
	return &p.RequestPacket
}

func (s *QueuingService) GetTicket(ctx context.Context, data GetTicketOptions) (*protos.QueuingGetTicketResponse, error) {
	param := QueuingGetTicketRequestWrapper{
		QueuingGetTicketRequest: &protos.QueuingGetTicketRequest{
			NpSN:          data.NpSN,
			NpToken:       data.NpToken,
			Npacode:       data.NpaCode,
			NgsmToken:     data.NgsmToken,
			ClientVersion: data.ClientVersion,
			OSType:        defaultFullOSType,
			AccessIP:      defaultAccessIP,
		},
	}
	req, err := s.client.R().WithGatewayBypass().Gateway(ctx, protos.Protocol_Queuing_GetTicket, param, WithHash(0))
	if err != nil {
		return nil, fmt.Errorf("failed to create get ticket request: %w", err)
	}
	result := new(protos.QueuingGetTicketResponse)
	_, err = s.client.Do(ctx, req, result)
	if err != nil {
		return nil, fmt.Errorf("get ticket request failed: %w", err)
	}
	return result, nil
}
