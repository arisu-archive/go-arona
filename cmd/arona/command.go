package main

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/hex"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/arisu-archive/go-arona/arona"
)

const (
	keysCommandName      = "keys"
	defaultPublicKeyPath = "./keys/key.pem"
)

const (
	serverNameAsia         = "asia"
	serverNameTaiwan       = "taiwan"
	serverNameNorthAmerica = "north-america"
	serverNameEurope       = "europe"
	serverNameKorea        = "korea"
)

type serverFlagValue struct {
	server *arona.Server
}

func (v *serverFlagValue) Set(value string) error {
	switch value {
	case serverNameAsia:
		*v.server = arona.ServerAsia
	case serverNameTaiwan:
		*v.server = arona.ServerTaiwan
	case serverNameNorthAmerica:
		*v.server = arona.ServerNorthAmerica
	case serverNameEurope:
		*v.server = arona.ServerEurope
	case serverNameKorea:
		*v.server = arona.ServerKorea
	default:
		return fmt.Errorf("invalid server %q", value)
	}
	return nil
}

func (v *serverFlagValue) String() string {
	switch *v.server {
	case arona.ServerAsia:
		return serverNameAsia
	case arona.ServerTaiwan:
		return serverNameTaiwan
	case arona.ServerNorthAmerica:
		return serverNameNorthAmerica
	case arona.ServerEurope:
		return serverNameEurope
	case arona.ServerKorea:
		return serverNameKorea
	default:
		return "unknown"
	}
}

func (*serverFlagValue) Type() string {
	return "server"
}

type rootOptions struct {
	publicKeyPath string
	server        arona.Server
	isJSONOutput  bool
}

type keysOutput struct {
	Key     string `json:"key"`
	License string `json:"license"`
}

func newRootCommand() *cobra.Command {
	options := rootOptions{
		publicKeyPath: defaultPublicKeyPath,
	}
	root := &cobra.Command{
		Use:           "arona",
		Short:         "arona command-line client",
		SilenceErrors: true,
		SilenceUsage:  true,
	}
	root.PersistentFlags().StringVar(&options.publicKeyPath, "key", options.publicKeyPath, "Path to the RSA public key")
	root.PersistentFlags().Var(&serverFlagValue{server: &options.server}, "server", "Game server")
	root.PersistentFlags().BoolVar(&options.isJSONOutput, "json", options.isJSONOutput, "Output decrypted keys as JSON")
	root.AddCommand(newKeysCommand(&options))
	if err := root.MarkPersistentFlagRequired("server"); err != nil {
		panic(err)
	}
	return root
}

func newKeysCommand(options *rootOptions) *cobra.Command {
	command := &cobra.Command{
		Use:   keysCommandName,
		Short: "Get SQLCipher keys",
		Args:  cobra.NoArgs,
		RunE: func(command *cobra.Command, _ []string) error {
			return retrieveKeys(command, *options)
		},
	}
	return command
}

func retrieveKeys(command *cobra.Command, options rootOptions) error {
	publicKey, err := loadRSAPublicKey(options.publicKeyPath)
	if err != nil {
		return err
	}
	bundle, err := newAESKeyBundle(rand.Reader)
	if err != nil {
		return err
	}
	client := arona.NewClient(publicKey, nil).WithServer(options.server)
	response, err := client.Queuing.GetCryptoKeys(command.Context(), bundle)
	if err != nil {
		return fmt.Errorf("get crypto keys: %w", err)
	}

	key, err := decryptServerValue(bundle.Key, bundle.IV, response.EncryptedSqlCipherKey)
	if err != nil {
		return fmt.Errorf("decrypt SQLCipher key: %w", err)
	}
	license, err := decryptServerValue(bundle.Key, bundle.IV, response.EncryptedSqlCipherLicense)
	if err != nil {
		return fmt.Errorf("decrypt SQLCipher license: %w", err)
	}
	output := keysOutput{
		Key:     hex.EncodeToString(key),
		License: string(license),
	}
	if options.isJSONOutput {
		if err := json.NewEncoder(command.OutOrStdout()).Encode(output); err != nil {
			return fmt.Errorf("write decrypted keys: %w", err)
		}
		return nil
	}
	if _, err := fmt.Fprintf(command.OutOrStdout(), "Key: %s\nLicense: %s\n", output.Key, output.License); err != nil {
		return fmt.Errorf("write decrypted keys: %w", err)
	}
	return nil
}

func loadRSAPublicKey(path string) (*rsa.PublicKey, error) {
	contents, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read RSA public key: %w", err)
	}

	contents = bytes.TrimSpace(contents)
	if !bytes.HasPrefix(contents, []byte("-----BEGIN PUBLIC KEY-----")) {
		return nil, errors.New("RSA public key file must contain one PUBLIC KEY PEM block")
	}
	block, remainder := pem.Decode(contents)
	if block == nil {
		return nil, errors.New("RSA public key file does not contain a PEM block")
	}
	if block.Type != "PUBLIC KEY" {
		return nil, errors.New("RSA public key PEM block must have type PUBLIC KEY")
	}
	if len(block.Headers) != 0 {
		return nil, errors.New("RSA public key PEM block must not contain headers")
	}
	if len(bytes.TrimSpace(remainder)) != 0 {
		return nil, errors.New("RSA public key file contains trailing data")
	}
	parsedKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("parse RSA public key: %w", err)
	}
	publicKey, ok := parsedKey.(*rsa.PublicKey)
	if !ok {
		return nil, errors.New("public key is not RSA")
	}
	return publicKey, nil
}
