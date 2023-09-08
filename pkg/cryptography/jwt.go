package cryptography

import "github.com/golang-jwt/jwt/v4"

func GetJwtToken(secretKey string, iat, seconds int64, params map[string]any) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	for key, val := range params {
		claims[key] = val
	}
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
