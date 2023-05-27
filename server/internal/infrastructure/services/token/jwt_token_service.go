package token

import (
	"github.com/glopezep/arithmetic-calculator/internal/infrastructure/config"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type jwtTokenService struct {
	config config.Config
}

type Claims struct {
	jwt.RegisteredClaims
}

func (s *jwtTokenService) Sign(userId uuid.UUID) (string, error) {
	secret := []byte(s.config.Secret)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject: userId.String(),
		},
	})

	return token.SignedString(secret)
}

func (s *jwtTokenService) Verify(tokenStr string) (*Claims, error) {
	secret := []byte(s.config.Secret)

	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})

	claims, ok := token.Claims.(Claims)
	if !ok || !token.Valid {
		return nil, err
	}

	return &claims, nil
}

func NewJwtTokenService() *jwtTokenService {
	return &jwtTokenService{}
}
