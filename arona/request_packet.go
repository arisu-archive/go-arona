package arona

import "github.com/arisu-archive/arona-protos/protos"

func newRequestPacket(protocol protos.Protocol, requestCount int64, session protos.SessionKey) *protos.RequestPacket {
	return &protos.RequestPacket{
		BasePacket: protos.BasePacket{
			Protocol:   protocol,
			SessionKey: session,
			AccountId:  session.AccountServerId,
		},
		Hash: requestCount | (int64(protocol) << 32),
	}
}
