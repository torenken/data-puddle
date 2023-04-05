package main

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/ardanlabs/conf/v3"
	"github.com/torenken/data-puddle/foundation/encrypt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

var build = "develop"

type config struct {
	conf.Version
	Args conf.Args

	Url        string
	ApiKey     string
	EncryptKey string

	Auth struct {
		ClientId       string
		ClientSecretId string
		TokenUrl       string `conf:"default:https://datapuddle.auth.eu-central-1.amazoncognito.com/oauth2/token"`
	}
}

func main() {
	if err := run(); err != nil {
		if !errors.Is(err, conf.ErrHelpWanted) {
			fmt.Println("ERROR", err)
		}
		os.Exit(1)
	}
}

func run() error {
	fmt.Println("starting data export service 🤘")

	cfg, err := cmdConfig()
	if err != nil {
		return fmt.Errorf("prepare config: %w", err)
	}

	exportUrl, err := getPresignedURL(cfg)
	if err != nil {
		return fmt.Errorf("getting presigned url: %w", err)
	}

	fmt.Println("fetching data from presigned url 🤖")
	dataEncrypt, err := fetchData(exportUrl)
	if err != nil {
		return fmt.Errorf("fetching data: %w", err)
	}

	fmt.Println("decrypting data 🤫")
	data, err := decryptData(dataEncrypt, cfg)
	if err != nil {
		return fmt.Errorf("decrypt data: %w", err)
	}
	fmt.Printf("result 🥳 => %v\n", string(data))

	return nil
}

type ExportUrl struct {
	ExportUrl string `json:"exportUrl"`
}

func decryptData(data []byte, cfg config) ([]byte, error) {
	decodeKeyValue, err := base64.StdEncoding.DecodeString(cfg.EncryptKey)
	if err != nil {
		return nil, fmt.Errorf("base64 decoding: %w", err)
	}
	key := (*[32]byte)(decodeKeyValue)

	plaintext, err := encrypt.Decrypt(data, key)
	if err != nil {
		return nil, fmt.Errorf("encryption data: %w", err)
	}
	return plaintext, nil
}

func fetchData(exportUrl ExportUrl) ([]byte, error) {
	resp, err := http.Get(exportUrl.ExportUrl)
	if err != nil {
		return nil, fmt.Errorf("export data: %w", err)
	}
	buf := bytes.Buffer{}
	_, err = buf.ReadFrom(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("no response data: %w", err)
	}
	return buf.Bytes(), nil
}

func getPresignedURL(cfg config) (ExportUrl, error) {
	config := clientcredentials.Config{
		ClientID:     cfg.Auth.ClientId,
		ClientSecret: cfg.Auth.ClientSecretId,
		Scopes:       []string{"datapuddle/ticket-export-url"},
		TokenURL:     cfg.Auth.TokenUrl,
		AuthStyle:    oauth2.AuthStyleInParams,
	}
	client := config.Client(context.Background())
	client.Timeout = time.Second * 5
	req, err := http.NewRequest("GET", cfg.Url, nil)
	if err != nil {
		return ExportUrl{}, fmt.Errorf("configure http request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-api-key", cfg.ApiKey)

	resp, err := client.Do(req)
	if err != nil {
		return ExportUrl{}, fmt.Errorf("connecting to serice: %w", err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return ExportUrl{}, fmt.Errorf("connecting to serice: %w", err)
	}

	var exportUrl ExportUrl
	err = json.Unmarshal(body, &exportUrl)
	if err != nil {
		return ExportUrl{}, fmt.Errorf("connecting to serice: %w", err)
	}
	return exportUrl, nil
}

func cmdConfig() (config, error) {
	cfg := config{
		Version: conf.Version{
			Build: build,
		},
	}

	const prefix = "EXPORT"
	help, err := conf.Parse(prefix, &cfg)
	if err != nil {
		if errors.Is(err, conf.ErrHelpWanted) {
			fmt.Println(help)
			return config{}, nil
		}

		out, err := conf.String(&cfg)
		if err != nil {
			return config{}, fmt.Errorf("generating config for output: %w", err)
		}
		fmt.Printf("startup: %v", out)

		return config{}, fmt.Errorf("parsing config: %w", err)
	}
	return cfg, nil

}
