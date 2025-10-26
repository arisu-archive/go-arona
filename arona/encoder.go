package arona

import (
	"bytes"
	"io"
	"net/http"

	"github.com/arisu-archive/arona-protos/protos"
)

type encoderRequest struct {
	Server   string `json:"server"`
	Protocol uint64 `json:"protocol"`
	Crc32    uint64 `json:"crc32"`
}

func (c *Client) encodeProtocol(p protos.Protocol, crc32 uint32) (uint32, error) {
	data := encoderRequest{
		Server:   "global",
		Protocol: uint64(p),
		Crc32:    uint64(crc32),
	}
	payload, err := c.JSONSerializer.Serialize(data, "")
	if err != nil {
		return 0, err
	}

	u, err := c.ProtocolEncoderURL.Parse("/")
	if err != nil {
		return 0, err
	}
	req, err := http.NewRequest("POST", u.String(), bytes.NewBuffer(payload))
	if err != nil {
		return 0, err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	var encoderResp struct {
		Result uint64 `json:"result"`
	}
	responseData, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}
	if err := c.JSONSerializer.Deserialize(responseData, &encoderResp); err != nil {
		return 0, err
	}

	return uint32(encoderResp.Result), nil
}
