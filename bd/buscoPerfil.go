package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/GITCART/development-twittor-go/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*BuscoPerfil busca perfil en la BD*/
func BuscoPerfil(ID string) (models.Usuario, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoCN.Database("twitter")
	col := db.Collection("usuarios")

	var perfil models.Usuario

	objtID, _ := primitive.ObjectIDFromHex(ID)

	condicion := bson.M{
		"_id": objtID,
	}

	err := col.FindOne(ctx, condicion).Decode(&perfil)
	perfil.Password = ""

	if err != nil {
		fmt.Println("Registro no encontrado" + err.Error())

		return perfil, err
	}

	return perfil, nil
}