package cmdutil

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

func credentialsFilepath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("failed to get home dir: %w", err)
	}
	return filepath.Join(home, ".config", "simoom", "credentials.json"), nil
}

// ReadCredentials は認証ファイルから認証情報を読み込む
// 認証ファイルが存在しない場合は空の Credentials を返す
func ReadCredentials() (Credentials, error) {
	p, err := credentialsFilepath()
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
