package middlew

import (
	"net/http"

	"github.com/GITCART/development-twittor-go/bd"
)

/*ChequeoBD es el middlew que me permite conocer el estado de la BD*/
func ChequeoBD(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if bd.ChequeConnection() == 0 {
			http.Error(w, "Conexion perdida con la bd", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
