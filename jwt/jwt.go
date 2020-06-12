package jwt

import (
	"time"

	"github.com/GITCART/development-twittor-go/models"
	jwt "github.com/dgrijalva/jwt-go"
)

/*GeneroJWT genera el encriptado con JWT*/
func GeneroJWT(t models.Usuario) (string, error) {
	miClave := []byte("MasterdelDesarrollo_Facebook")

	payload := jwt.MapClaims{
		"email":            t.Email,
		"nombre":           t.Nombre,
		"apellido":         t.Apellidos,
		"fecha_nacimiento": t.FechaNacimiento,
		"biografgia":       t.Biografia,
		"ubicacion":        t.Ubicacion,
		"sitioweb":         t.SitoWeb,
		"_id":              t.ID.Hex(),
		"exp":              time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(miClave)

	if err != nil {
		return tokenStr, err
	}
	return tokenStr, nil
}
