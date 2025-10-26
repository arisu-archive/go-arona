package arona

import (
	"bytes"
	"crypto/rand"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"sync"

	"github.com/arisu-archive/arona-protos/protos"
)

const (
	Version          = "1.82.376243"
	defaultUserAgent = "BestHTTP/2 v2.4.0"
	defaultXorKey    = 0xD9
)

type Client struct {
	clientMu sync.Mutex
	client   *http.Client

	// XorEncryptionKey is the byte used to XOR the payload before sending.
	XorEncryptionKey byte

	// JSONSerializer is used to serialize and deserialize JSON payloads.
	JSONSerializer JSONSerializer

	// User agent used when communicating with the game API.
	UserAgent string

	// Base URL for protocol encoder service.
	ProtocolEncoderURL *url.URL

	Server     Server   // The server to which requests will be sent.
	GatewayURL *url.URL // Base URL for gateway requests. Defaults based on the provided server variable.
	GameURL    *url.URL // Base URL for game requests. Defaults based on the provided server variable.

	processor *Processor // Packet processor for handling encryption and payload building.

	common service // Reuse a single struct instead of allocating one for each service on the heap.
	// Services used for talking to different parts of the game API.
}

// service represents a service for interacting with a specific part of the game API.
type service struct {
	client *Client
}

// SessionKey holds keys and IVs used for encrypting and forging packets.
type SessionKey struct {
	AESKey    *[16]byte // Optional AES encryption key
	AESIV     *[16]byte // Optional AES initialization vector
	ServerKey []byte    // Server key to include in packet
	ServerIV  []byte    // Server IV to include in packet
}

// NewClient returns a new Arona API client. If a nil httpClient is
// provided, a new http.Client will be used.
func NewClient(server Server, encoder *url.URL, httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = &http.Client{}
	}
	httpClient2 := *httpClient
	c := &Client{
		client:             &httpClient2,
		Server:             server,
		ProtocolEncoderURL: encoder,
	}
	return c.initialize()
}

// NewRequest creates a new API request. A relative URL can be provided in urlStr,
// in which case it is resolved relative to the BaseURL of the Client.
// Relative URLs should always be specified without a preceding slash.
func (c *Client) NewRequest(protocol protos.Protocol, body any, key SessionKey) (*http.Request, error) {
	// Process payload through crypto pipeline
	payload, err := c.processor.Process(body, key)
	if err != nil {
		return nil, err
	}
	// Encode protocol with checksum
	checksum := computeHash(payload, 0)
	encodedProtocol, err := c.encodeProtocol(protocol, checksum)
	if err != nil {
		return nil, fmt.Errorf("protocol encoding failed: %w", err)
	}
	// Build final packet
	packetData := c.processor.BuildPacket(payload, checksum, encodedProtocol, key)
	// Create multipart form
	mw := &multipartWriter{}
	buf, contentType, err := mw.write(packetData)
	if err != nil {
		return nil, fmt.Errorf("multipart creation failed: %w", err)
	}

	// Build HTTP request
	req, err := c.buildHTTPRequest(buf, contentType)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// buildHTTPRequest creates the HTTP request with proper headers.
func (c *Client) buildHTTPRequest(body *bytes.Buffer, contentType string) (*http.Request, error) {
	u, err := c.GatewayURL.Parse("/api/gateway")
	if err != nil {
		return nil, fmt.Errorf("URL parse failed: %w", err)
	}

	req, err := http.NewRequest("POST", u.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", c.UserAgent)
	req.Header.Set("Content-Type", contentType)
	req.Header.Set("mx", "2")
	req.Header.Set("Accept-Encoding", "identity")

	return req, nil
}

// JSONSerializer defines methods for serializing and deserializing JSON data.
type JSONSerializer interface {
	Serialize(v any, indent string) ([]byte, error)
	Deserialize(data []byte, v any) error
}

// StdJSONSerializer is the default implementation of JSONSerializer using the encoding/json package.
type defaultJSONSerializer struct{}

// Serialize serializes a value into JSON.
func (s *defaultJSONSerializer) Serialize(v any, indent string) ([]byte, error) {
	if indent != "" {
		return json.MarshalIndent(v, "", indent)
	}
	return json.Marshal(v)
}

// Deserialize deserializes JSON data into a value.
func (s *defaultJSONSerializer) Deserialize(data []byte, v any) error {
	return json.Unmarshal(data, v)
}

// ErrInvalidServer is returned when an invalid server is specified.
var ErrInvalidServer = errors.New("invalid server specified")

// Server represents the different game servers available.
type Server int

const (
	ServerAsia Server = iota
	ServerTaiwan
	ServerNorthAmerica
	ServerEurope
	ServerKorea
)

// API mappings for each server.
type serverApiMapping struct {
	GatewayAPI string
	GameAPI    string
}

// Predefined server configurations.
var serverConfigs = map[Server]serverApiMapping{
	ServerAsia: {
		GatewayAPI: "https://nxm-th-bagl.nexon.com:5100/",
		GameAPI:    "https://nxm-th-bagl.nexon.com:5000/",
	},
	ServerTaiwan: {
		GatewayAPI: "https://nxm-tw-bagl.nexon.com:5100/",
		GameAPI:    "https://nxm-tw-bagl.nexon.com:5000/",
	},
	ServerNorthAmerica: {
		GatewayAPI: "https://nxm-or-bagl.nexon.com:5100/",
		GameAPI:    "https://nxm-or-bagl.nexon.com:5000/",
	},
	ServerEurope: {
		GatewayAPI: "https://nxm-eu-bagl.nexon.com:5100/",
		GameAPI:    "https://nxm-eu-bagl.nexon.com:5000/",
	},
	ServerKorea: {
		GatewayAPI: "https://nxm-kr-bagl.nexon.com:5100/",
		GameAPI:    "https://nxm-kr-bagl.nexon.com:5000/",
	},
}

// resolveGatewayURL returns the gateway URL for the specified server.
func resolveGatewayURL(server Server) (*url.URL, error) {
	config, ok := serverConfigs[server]
	if !ok {
		return nil, ErrInvalidServer
	}
	return url.Parse(config.GatewayAPI)
}

// resolveGameURL returns the game URL for the specified server.
func resolveGameURL(server Server) (*url.URL, error) {
	config, ok := serverConfigs[server]
	if !ok {
		return nil, ErrInvalidServer
	}
	return url.Parse(config.GameAPI)
}

// initialize sets up the client with default values.
func (c *Client) initialize() *Client {
	// Set default URLs based on the server
	if c.GatewayURL == nil {
		c.GatewayURL, _ = resolveGatewayURL(c.Server)
	}
	if c.GameURL == nil {
		c.GameURL, _ = resolveGameURL(c.Server)
	}
	if c.UserAgent == "" {
		c.UserAgent = defaultUserAgent
	}
	if c.XorEncryptionKey == 0 {
		c.XorEncryptionKey = defaultXorKey
	}
	if c.JSONSerializer == nil {
		c.JSONSerializer = &defaultJSONSerializer{}
	}
	c.processor = &Processor{
		xorKey:         c.XorEncryptionKey,
		jsonSerializer: c.JSONSerializer,
	}
	c.common.client = c
	return c
}

type multipartWriter struct {
	boundary string
}

func (mw *multipartWriter) write(packetData []byte) (*bytes.Buffer, string, error) {
	buf := &bytes.Buffer{}
	writer := multipart.NewWriter(buf)
	writer.SetBoundary(fmt.Sprintf("BestHTTP_HTTPMultiPartForm_%s", randomBoundary()))

	part, err := writer.CreateFormFile("mx", "mx.dat")
	if err != nil {
		return nil, "", err
	}
	if _, err := part.Write(packetData); err != nil {
		return nil, "", err
	}
	if err := writer.Close(); err != nil {
		return nil, "", err
	}
	return buf, writer.FormDataContentType(), nil
}

// randomBoundary generates a random boundary string for multipart form data.
func randomBoundary() string {
	var buf [4]byte
	_, err := io.ReadFull(rand.Reader, buf[:])
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("%X", buf[:])
}
