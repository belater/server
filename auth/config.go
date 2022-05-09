package auth

import (
	"errors"
	"os"

	"github.com/dgrijalva/jwt-go"
)

var (
	key = os.Getenv("JWT_SECRET")
)

type AuthService interface {
	GenerateToken(UserId string, Role string) (string, error)
	ValidateToken(encodedToken string) (*jwt.Token, error)
}

type authService struct{}

func NewAuthService() *authService {
	return &authService{}
}
func (s *authService) GenerateToken(UserId string, Role string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": UserId,
		"role":    Role,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(key))

	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}

func (s *authService) ValidateToken(encodedToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(encodedToken, func(encodedToken *jwt.Token) (interface{}, error) {
		_, ok := encodedToken.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("invalid token")
		}

		return []byte(key), nil
	})

	if err != nil {
		return token, err
	}

	return token, nil
}
