package usecase_test

import (
	"testing"

	"github.com/minguu42/simoom/pkg/usecase"
	"github.com/stretchr/testify/assert"
)

func TestSignUpInput_Validate(t *testing.T) {
	tests := []struct {
		name     string
		in       usecase.SignUpInput
		hasError bool
	}{
		{
			name: "nameに空の文字列は指定できない",
			in: usecase.SignUpInput{
				Name:     "",
				Email:    "dummy@example.com",
				Password: "strong-password",
			},
			hasError: true,
		},
		{
			name: "nameに1文字の文字列は指定できる",
			in: usecase.SignUpInput{
				Name:     "A",
				Email:    "dummy@example.com",
				Password: "strong-password",
			},
			hasError: false,
		},
		{
			name: "nameに15文字の文字列は指定できる",
			in: usecase.SignUpInput{
				Name:     "あいうえおかきくけこさしすせそ",
				Email:    "dummy@example.com",
				Password: "strong-password",
			},
			hasError: false,
		},
		{
			name: "nameに16文字以上の文字列は指定できない",
			in: usecase.SignUpInput{
				Name:     "あいうえおかきくけこさしすせそA",
				Email:    "dummy@example.com",
				Password: "strong-password",
			},
			hasError: true,
		},
		{
			name: "emailに空の文字列は指定できない",
			in: usecase.SignUpInput{
				Name:     "テストユーザ",
				Email:    "",
				Password: "strong-password",
			},
			hasError: true,
		},
		{
			name: "emailに1文字の文字列は指定できる",
			in: usecase.SignUpInput{
				Name:     "テストユーザ",
				Email:    "a",
				Password: "strong-password",
			},
			hasError: false,
		},
		{
			name: "emailに254文字の文字列は指定できる",
			in: usecase.SignUpInput{
				Name:     "テストユーザ",
				Email:    "very-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-lo@example.com",
				Password: "strong-password",
			},
			hasError: false,
		},
		{
			name: "emailに255文字以上の文字列は指定できない",
			in: usecase.SignUpInput{
				Name:     "テストユーザ",
				Email:    "very-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-lon@example.com",
				Password: "strong-password",
			},
			hasError: true,
		},
		{
			name: "passwordに11文字以下の文字列は指定できない",
			in: usecase.SignUpInput{
				Name:     "テストユーザ",
				Email:    "dummy@example.com",
				Password: "short-pw123",
			},
			hasError: true,
		},
		{
			name: "passwordに12文字の文字列は指定できる",
			in: usecase.SignUpInput{
				Name:     "テストユーザ",
				Email:    "dummy@example.com",
				Password: "short-pw1234",
			},
			hasError: true,
		},
		{
			name: "passwordに20文字の文字列は指定できる",
			in: usecase.SignUpInput{
				Name:     "テストユーザ",
				Email:    "dummy@example.com",
				Password: "long-long-password12",
			},
			hasError: true,
		},
		{
			name: "passwordに21文字以上の文字列は指定できない",
			in: usecase.SignUpInput{
				Name:     "テストユーザ",
				Email:    "dummy@example.com",
				Password: "long-long-password123",
			},
			hasError: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.in.Validate()
			assert.Equal(t, tt.hasError, err != nil)
		})
	}
}

func TestSignInInput_Validate(t *testing.T) {
	tests := []struct {
		name     string
		in       usecase.SignInInput
		hasError bool
	}{
		{
			name: "emailに空の文字列は指定できない",
			in: usecase.SignInInput{
				Email:    "",
				Password: "strong-password",
			},
			hasError: true,
		},
		{
			name: "emailに1文字の文字列は指定できる",
			in: usecase.SignInInput{
				Email:    "a",
				Password: "strong-password",
			},
			hasError: false,
		},
		{
			name: "emailに254文字の文字列は指定できる",
			in: usecase.SignInInput{
				Email:    "very-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-lo@example.com",
				Password: "strong-password",
			},
			hasError: false,
		},
		{
			name: "emailに255文字以上の文字列は指定できない",
			in: usecase.SignInInput{
				Email:    "very-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-long-lon@example.com",
				Password: "strong-password",
			},
			hasError: true,
		},
		{
			name: "passwordに11文字以下の文字列は指定できない",
			in: usecase.SignInInput{
				Email:    "dummy@example.com",
				Password: "short-pw123",
			},
			hasError: true,
		},
		{
			name: "passwordに12文字の文字列は指定できる",
			in: usecase.SignInInput{
				Email:    "dummy@example.com",
				Password: "short-pw1234",
			},
			hasError: true,
		},
		{
			name: "passwordに20文字の文字列は指定できる",
			in: usecase.SignInInput{
				Email:    "dummy@example.com",
				Password: "long-long-password12",
			},
			hasError: true,
		},
		{
			name: "passwordに21文字以上の文字列は指定できない",
			in: usecase.SignInInput{
				Email:    "dummy@example.com",
				Password: "long-long-password123",
			},
			hasError: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.in.Validate()
			assert.Equal(t, tt.hasError, err != nil)
		})
	}
}

func TestRefreshAccessTokenInput_Validate(t *testing.T) {
	tests := []struct {
		name     string
		in       usecase.RefreshTokenInput
		hasError bool
	}{
		{
			name: "refresh_tokenに空の文字列は指定できない",
			in: usecase.RefreshTokenInput{
				RefreshToken: "",
			},
			hasError: true,
		},
		{
			name: "refresh_tokenに1文字は指定できる",
			in: usecase.RefreshTokenInput{
				RefreshToken: "a",
			},
			hasError: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.in.Validate()
			assert.Equal(t, tt.hasError, err != nil)
		})
	}
}
