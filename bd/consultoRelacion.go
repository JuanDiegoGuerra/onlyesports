package bd

import (
	"context"
	"time"

	"github.com/JuanDiegoGuerra/onlyesports/models"
	"go.mongodb.org/mongo-driver/bson"
)

/*ConsultoRelacion consulta la relacion entre 2 usuarios */
func ConsultoRelacion(t models.Relacion) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("onlyesports")
	col := db.Collection("relacion")

	condicion := bson.M{
		"usuarioid":         t.UsuarioID,
		"usuariorelacionid": t.UsuarioRelacionID,
	}

	var resultado models.Relacion
	//fmt.Println(resultado)
	err := col.FindOne(ctx, condicion).Decode(&resultado)
	if err != nil {
		return false, err
	}
	return true, nil
}
