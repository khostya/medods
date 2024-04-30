package model

import (
	"encoding/base64"
)

type Tokens struct {
	Access  string `json:"access"`
	Refresh string `json:"refresh"`
}

func NewTokens(accessToken, refreshToken string) Tokens {
	encodedRefreshToken := base64.StdEncoding.EncodeToString([]byte(refreshToken))
	return Tokens{
		Access:  accessToken,
		Refresh: encodedRefreshToken,
	}
}
