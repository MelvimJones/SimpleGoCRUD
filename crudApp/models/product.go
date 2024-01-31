package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Product representa a estrutura de dados para um produto
type Product struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Descricao string             `json:"descricao" bson:"descricao"`
	Preco     float64            `json:"preco" bson:"preco"`
	Imagem    string             `json:"imagem" bson:"imagem"`
	Quant     int                `json:"quant" bson:"quant"`
}
