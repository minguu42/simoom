package api

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Credentials struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func filepathCredentials() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("failed to get home dir: %w", err)
	}
	return filepath.Join(home, ".config", "simoom", "credentials.json"), nil
}

// NewCredentials は Credentials を作成する
// 認証ファイルが存在する場合は認証ファイルを読み込み、 存在しない場合は空の Credentials を返す
func NewCredentials() (Credentials, error) {
	p, err := filepathCredentials()
	if err != nil {
		return Credentials{}, fmt.Errorf("failed to get credentials file path: %w", err)
	}
	f, err := os.Open(p)
	if err != nil {
		if os.IsNotExist(err) {
			return Credentials{}, nil
		}
		return Credentials{}, fmt.Errorf("failed to open file: %w", err)
	}
	defer f.Close()

	var c Credentials
	if err := json.NewDecoder(f).Decode(&c); err != nil {
		return Credentials{}, fmt.Errorf("failed to decode credentials: %w", err)
	}
	return c, nil
}

// WriteCredentials は認証ファイルに認証情報を書き込む
func WriteCredentials(accessToken, refreshToken string) error {
	p, err := filepathCredentials()
	if err != nil {
		return fmt.Errorf("failed to get credentials file path: %w", err)
	}
	if err := os.MkdirAll(filepath.Dir(p), 0755); err != nil {
		return fmt.Errorf("failed to make directories: %w", err)
	}
	f, err := os.Create(p)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer f.Close()

	if err := json.NewEncoder(f).Encode(Credentials{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}); err != nil {
		return fmt.Errorf("failed to encode credentials: %w", err)
	}

	if err := f.Sync(); err != nil {
		return fmt.Errorf("failed to sync file: %w", err)
	}
	return nil
}
