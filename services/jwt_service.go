package services

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

type jwtService struct {
	secretKey string
	issure    string
}

func NewJWTService() *jwtService {
	return &jwtService{
		secretKey: "secret-key",
		issure:    "book-api",
	}
}

type Claim struct {
	Sum uint `json:"sum"`
	jwt.StandardClaims
}

func (s *jwtService) GeraToken(id uint) (string, error) {
	claim := &Claim{
		id,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 2).Unix(),
			Issuer:    s.issure,
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	t, err := token.SignedString([]byte(s.secretKey))
	if err != nil {
		return "", err
	}

	return t, nil
}

func (s *jwtService) ValidaToken(token string) bool {
	_, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, isValid := t.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, fmt.Errorf("token inválido")
		}

		return []byte(s.secretKey), nil
	})

	return err == nil
}

/*func (s *jwtService) BuscaTokenPorID(t string) (int64, error) {
	token, err := jwt.Parse(t, func(token *jwt.Token) (interface{}, error) {
		if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
			return nil, fmt.Errorf("Token inválido: %v", t)
		}
		return []byte(config.GetConfig().JWTSecretKey), nil
	})
	if err != nil {
		return 0, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		id := claims["sum"].(string)
		val, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			return 0, err
		}

		return val, nil
	}

	return 0, err

}
*/
