# 🎮 Arona Client

[![Go Reference](https://pkg.go.dev/badge/github.com/arisu-archive/go-arona.svg)](https://pkg.go.dev/github.com/arisu-archive/go-arona) [![Go Report Card](https://goreportcard.com/badge/github.com/arisu-archive/go-arona)](https://goreportcard.com/report/github.com/arisu-archive/go-arona) [![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

**A robust Go client library for interacting with Blue Archive game servers**

[Features](#-features) • [Installation](#-installation) • [Quick Start](#-quick-start) • [Documentation](#-documentation) • [Contributing](#-contributing)


---

## 📖 Overview

Arona Client is a comprehensive Go library that provides a clean, idiomatic interface for communicating with Blue Archive game servers. It handles the complex protocol encoding, encryption, session management, and API interactions, allowing developers to focus on building applications rather than dealing with low-level protocol details.

### Key Highlights

- 🔐 **Secure Communication**: Built-in RSA encryption and AES session key management
- 🌍 **Multi-Region Support**: Supports Asia, Taiwan, North America, Europe, and Korea servers
- 🎯 **Type-Safe API**: Leverages Protocol Buffers and FlatBuffers for type safety
- 🧩 **Modular Design**: Clean service-oriented architecture with dedicated modules for each game feature
- ⚡ **Performance**: Efficient request processing with connection pooling and gzip compression
- 🧪 **Well-Tested**: Comprehensive test coverage using Ginkgo/Gomega

## ✨ Features

### Core Services

- **Account Management** - Authentication, Nexon account validation, session handling
- **Arena** - Competitive ranking information and opponent lists
- **Raid System** - Raid lobby access, opponent search, team formations, rankings
- **Eliminate Raid** - Specialized eliminate raid operations and rankings
- **Clan Operations** - Clan-related functionalities
- **Friend System** - Friend search, detailed info retrieval
- **Queuing System** - Ticket management for server queuing

### Advanced Capabilities

- Custom JSON serialization support
- Flexible request builder pattern
- Automatic protocol encoding with checksum validation
- Session-based encryption/decryption
- Multipart form data handling
- Comprehensive error handling with typed error responses

## 📦 Installation

```bash
go get github.com/arisu-archive/go-arona
```

### Requirements

- Go 1.24.3 or higher
- Dependencies are managed via Go modules

## 🚀 Quick Start

### Basic Setup

```go
package main

import (
    "context"
    "crypto/rsa"
    "net/url"
    
    "github.com/arisu-archive/go-arona/arona"
    "github.com/arisu-archive/arona-protos/protos"
)

func main() {
    // Initialize the client
    protocolEncoderURL, _ := url.Parse("https://your-protocol-encoder-service")
    publicKey := &rsa.PublicKey{...} // Your RSA public key
    
    client := arona.NewClient(
        arona.ServerAsia,
        protocolEncoderURL,
        publicKey,
        nil, // Use default HTTP client
    )
    
    ctx := context.Background()
    
    // Example: Authenticate
    session := &arona.UserSession{...}
    authResp, err := client.Account.Authenticate(ctx, session)
    if err != nil {
        panic(err)
    }
}
```

### Working with Different Servers

```go
// Create client for specific regions
asiaClient := arona.NewClient(arona.ServerAsia, encoderURL, pubKey, nil)
naClient := arona.NewClient(arona.ServerNorthAmerica, encoderURL, pubKey, nil)
euClient := arona.NewClient(arona.ServerEurope, encoderURL, pubKey, nil)

// Or switch servers dynamically
newClient := client.WithServer(arona.ServerKorea)
```

### Arena Operations

```go
// Get arena rankings
rankings, err := client.Arena.GetRanks(ctx, 1, 100)
if err != nil {
    log.Fatal(err)
}

for _, rank := range rankings.Ranks {
    fmt.Printf("Rank %d: %s\n", rank.Rank, rank.Nickname)
}
```

### Raid System

```go
// Search raid opponents by rank
opponents, err := client.Raid.WithOpponentRank(100).Search(ctx, session)
if err != nil {
    log.Fatal(err)
}

// Or search by score
opponents, err = client.Raid.WithOpponentScore(5000000).Search(ctx, session)

// Get raid lobby info
lobby, err := client.Raid.Lobby(ctx, session)

// Get best team for a player
team, err := client.Raid.GetBestTeam(ctx, session, accountID)
```

### Friend System

```go
// Search for friends
results, err := client.Friend.Search().
    ByNickname("PlayerName").
    Submit(ctx, session)

// Get detailed friend info
friendInfo, err := client.Friend.GetDetail(ctx, session, friendAccountID)
```

### Custom Request Builder

```go
// Build custom requests with headers
resp, err := client.R().
    WithSession(session).
    WithAuthToken("bearer-token").
    WithHeader("Custom-Header", "value").
    Game(ctx, protos.Protocol_Custom_Operation, requestPayload)
```

## 📚 Documentation

### Available Servers

| Server | Gateway API | Game API |
|--------|------------|----------|
| Asia | `https://nxm-th-bagl.nexon.com:5100/` | `https://nxm-th-bagl.nexon.com:5000/` |
| Taiwan | `https://nxm-tw-bagl.nexon.com:5100/` | `https://nxm-tw-bagl.nexon.com:5000/` |
| North America | `https://nxm-or-bagl.nexon.com:5100/` | `https://nxm-or-bagl.nexon.com:5000/` |
| Europe | `https://nxm-eu-bagl.nexon.com:5100/` | `https://nxm-eu-bagl.nexon.com:5000/` |
| Korea | `https://nxm-kr-bagl.nexon.com:5100/` | `https://nxm-kr-bagl.nexon.com:5000/` |

### Session Management

Sessions are managed through the `UserSession` type which maintains:
- Session keys from the game server
- Client and server AES key bundles for encryption
- Request counter for packet sequencing

```go
session := &arona.UserSession{
    SessionKey: protos.SessionKey{...},
    ClientKeyBundle: arona.AESKeyBundle{
        Key: clientKey,
        IV:  clientIV,
    },
    ServerKeyBundle: arona.AESKeyBundle{
        Key: serverKey,
        IV:  serverIV,
    },
    RequestCount: 0,
}
```

### Error Handling

The library provides typed errors for better error handling:

```go
resp, err := client.Account.Authenticate(ctx, session)
if err != nil {
    var sessionErr *arona.InvalidSessionError
    if errors.As(err, &sessionErr) {
        // Handle invalid session
        fmt.Printf("Session error: %s (code: %d)\n", sessionErr.Error(), sessionErr.Code())
    }
    
    var apiErr *arona.ErrWebAPIError
    if errors.As(err, &apiErr) {
        // Handle API error
        fmt.Printf("API error: code=%d, reason=%s\n", 
            apiErr.Code(), apiErr.Packet.Reason)
    }
}
```

### Custom JSON Serialization

Implement the `JSONSerializer` interface for custom serialization:

```go
type CustomSerializer struct{}

func (s *CustomSerializer) Serialize(v any, indent string) ([]byte, error) {
    // Custom serialization logic
}

func (s *CustomSerializer) Deserialize(data []byte, v any) error {
    // Custom deserialization logic
}

func (s *CustomSerializer) DeserializeReader(r io.Reader, v any) error {
    // Custom reader deserialization
}

// Use custom serializer
client.JSONSerializer = &CustomSerializer{}
```

## 🏗️ Project Structure

```
.
├── arona/
│   ├── account.go           # Account service
│   ├── arena.go             # Arena service
│   ├── arona.go             # Core client implementation
│   ├── clan.go              # Clan service
│   ├── eliminate_raid.go    # Eliminate raid service
│   ├── encoder.go           # Protocol encoding
│   ├── errors.go            # Error types
│   ├── friend.go            # Friend service
│   ├── processor.go         # Request/response processing
│   ├── queuing.go           # Queuing service
│   ├── raid.go              # Raid service
│   ├── request_packet.go    # Request packet handling
│   └── response.go          # Response handling
├── go.mod
├── go.sum
├── LICENSE
└── README.md
```

## 🧪 Testing

The project uses Ginkgo and Gomega for testing:

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run specific test suite
go test -v ./arona
```

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request. For major changes, please open an issue first to discuss what you would like to change.

### Development Setup

1. Clone the repository:
   ```bash
   git clone https://github.com/arisu-archive/go-arona.git
   cd go-arona
   ```

2. Install dependencies:
   ```bash
   go mod download
   ```

3. Run tests:
   ```bash
   go test ./...
   ```

### Guidelines

- Follow Go best practices and idiomatic code style
- Write tests for new features
- Update documentation for API changes
- Ensure all tests pass before submitting PR
- Use conventional commit messages

## 📋 Roadmap

- [ ] Add more comprehensive examples
- [ ] Implement rate limiting and retry logic
- [ ] Add support for WebSocket connections
- [ ] Expand test coverage
- [ ] Add benchmarks
- [ ] Create detailed API documentation site

## ⚠️ Disclaimer

This project is for educational purposes only. Use at your own risk. The authors are not responsible for any misuse or damage caused by this library. Please respect the game's terms of service.

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

```
Copyright (c) 2024 Arisu Archive

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction...
```

## 🙏 Acknowledgments

- Blue Archive game by Nexon
- Protocol Buffers and FlatBuffers communities
- All contributors to the Arisu Archive project

## 📞 Support

- 📫 Issues: [GitHub Issues](https://github.com/arisu-archive/go-arona/issues)
- 💬 Discussions: [GitHub Discussions](https://github.com/arisu-archive/go-arona/discussions)
- 🌟 Star this repo if you find it helpful!

---

<div align="center">

**Made with ❤️ by the Arisu Archive Team**

[⬆ Back to Top](#-arona-client)

</div>
