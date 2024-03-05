package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"slices"
)

type credentials struct {
	Profile      string `json:"profile"`
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

// newCredentials は credentials を作成する
// 認証ファイルが存在する場合は認証ファイルから読み込み、 存在しない場合は空の credentials を返す
func newCredentials(profile string) (credentials, error) {
	var cs []credentials
	if err := readCredentialsFile(&cs); err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			return credentials{}, nil
		}
		return credentials{}, fmt.Errorf("failed to read credentials file: %w", err)
	}

	if i := slices.IndexFunc(cs, func(c credentials) bool {
		return c.Profile == profile
	}); i != -1 {
		return cs[i], nil
	}
	return credentials{}, errors.New("no corresponding credentials for profile")
}

// readCredentialsFile は認証ファイルを読み込む
func readCredentialsFile(cs *[]credentials) error {
	p, err := filepathCredentials()
	if err != nil {
		return fmt.Errorf("failed to get credentials file path: %w", err)
	}

	f, err := os.Open(p)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer f.Close()

	if err := json.NewDecoder(f).Decode(cs); err != nil {
		return fmt.Errorf("failed to decode credentials: %w", err)
	}
	return nil
}

// SaveCredentials は認証情報を認証ファイルに保存する
func SaveCredentials(profile, accessToken, refreshToken string) error {
	p, err := filepathCredentials()
	if err != nil {
		return fmt.Errorf("failed to get credentials file path: %w", err)
	}

	var cs []credentials
	if err := readCredentialsFile(&cs); err != nil && !errors.Is(err, fs.ErrNotExist) {
		return fmt.Errorf("failed to read credentials file: %w", err)
	}
	c := credentials{
		Profile:      profile,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
	if i := slices.IndexFunc(cs, func(c credentials) bool {
		return c.Profile == profile
	}); i != -1 {
		cs[i] = c
	} else {
		cs = append(cs, c)
	}

	if err := os.MkdirAll(filepath.Dir(p), 0755); err != nil {
		return fmt.Errorf("failed to make directories: %w", err)
	}
	f, err := os.Create(p)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer f.Close()

	encoder := json.NewEncoder(f)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(cs); err != nil {
		return fmt.Errorf("failed to encode credentials: %w", err)
	}

	if err := f.Sync(); err != nil {
		return fmt.Errorf("failed to sync file: %w", err)
	}
	return nil
}
