package auth

import (
	"errors"
	"fmt"
	"github.com/globalsign/mgo/bson"
	"github.com/golang-jwt/jwt/v5"
	"time"

	"github.com/google/uuid"
)

type TokenManager struct {
	signingKey string
}

func NewManager(signingKey string) TokenManager {
	return TokenManager{signingKey: signingKey}
}

func (m TokenManager) NewJWT(ID string, userID string, expiresAt time.Time) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": jwt.NewNumericDate(expiresAt),
		"sub": userID,
		"id":  ID,
	})

	return token.SignedString([]byte(m.signingKey))
}

func (m TokenManager) parse(accessToken string, options ...jwt.ParserOption) (jwt.MapClaims, error) {
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (i interface{}, err error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(m.signingKey), nil
	}, options...)

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("error get user claims from token")
	}

	return claims, nil
}

func (m TokenManager) NewRefreshToken() string {
	token := uuid.New().String()
	return token
}

func (m TokenManager) extractUserId(token string, options ...jwt.ParserOption) (string, error) {
	claims, err := m.extractClaims(token, options...)
	if err != nil {
		return "", err
	}

	sub, err := claims.GetSubject()
	if err != nil {
		return "", err
	}

	return sub, nil
}

func (m TokenManager) extractID(token string, options ...jwt.ParserOption) (bson.ObjectId, error) {
	claims, err := m.extractClaims(token, options...)
	if err != nil {
		return "", err
	}

	id, ok := claims["id"]
	if !ok {
		return "", errors.New("invalid claim")
	}

	objectId := bson.NewObjectId()
	err = objectId.UnmarshalText([]byte(id.(string)))
	return objectId, err
}

func (m TokenManager) ExtractUserIdWithoutClaimsValidation(token string) (string, error) {
	return m.extractUserId(token, jwt.WithoutClaimsValidation())
}

func (m TokenManager) ExtractIDWithoutClaimsValidation(token string) (bson.ObjectId, error) {
	return m.extractID(token, jwt.WithoutClaimsValidation())
}

func (m TokenManager) ExtractUserId(token string) (string, error) {
	return m.extractUserId(token)
}

func (m TokenManager) extractClaims(token string, options ...jwt.ParserOption) (jwt.MapClaims, error) {
	claims, err := m.parse(token, options...)
	if err != nil {
		return jwt.MapClaims{}, err
	}

	return claims, nil
}
