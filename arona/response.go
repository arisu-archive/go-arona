package arona

import "github.com/arisu-archive/arona-protos/protos"

type ResponseData struct {
	Protocol protos.Protocol `json:",omitempty,omitzero"`
	Packet   string          `json:",omitempty,omitzero"`
}
