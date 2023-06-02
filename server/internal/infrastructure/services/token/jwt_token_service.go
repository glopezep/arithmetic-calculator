package token

import (
	"os"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type jwtTokenService struct{}

type Claims struct {
	jwt.RegisteredClaims
}

func (s *jwtTokenService) Sign(userId uuid.UUID) (string, error) {
	secret := []byte(os.Getenv("SECRET"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject: userId.String(),
		},
	})

	return token.SignedString(secret)
}

func (s *jwtTokenService) Verify(tokenStr string) (*Claims, error) {
	secret := []byte(os.Getenv("SECRET"))

	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, ErrInvalidToken
		}
		return secret, nil
	}

	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, keyFunc)
	if err != nil {
		return nil, ErrInvalidToken
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, ErrInvalidToken
	}

	return claims, nil
}

func NewJwtTokenService() *jwtTokenService {
	return &jwtTokenService{}
}
