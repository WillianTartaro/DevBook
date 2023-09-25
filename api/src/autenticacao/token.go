package autenticacao

import (
	"DevBook/api/src/config"
	jwt "github.com/dgrijalva/jwt-go"
	"time"
)

// CriarToken retorna um token assinado com as permissões do usuario.
func CriarToken(usuarioID uint64) (string, error) {
	permissoes := jwt.MapClaims{}
	permissoes["authorized"] = true
	permissoes["exp"] = time.Now().Add(time.Hour * 6).Unix()
	permissoes["usuarioId"] = usuarioID
	token := jwt.NewWithClaims(jwt.SigningMethodES256, permissoes)
	return token.SignedString([]byte(config.SecretKey))
}
