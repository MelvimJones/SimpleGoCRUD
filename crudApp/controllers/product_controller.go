package controllers

import (
	"context"
	"crudApp/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Conexão com o servidor MongoDB
func ConnectToMongoDB() (*mongo.Client, error) {
	mongoURI := os.Getenv("MONGODB_URI")
	clientOptions := options.Client().ApplyURI(mongoURI)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	return client, nil
}

// Cria um novo produto no banco de dados
func CreateProductInMongoDB(w http.ResponseWriter, r *http.Request) {
	client, err := ConnectToMongoDB()
	if err != nil {
		http.Error(w, "Erro ao conectar ao MongoDB", http.StatusInternalServerError)
		return
	}
	defer client.Disconnect(context.Background())

	database := client.Database("crud_GO")
	collection := database.Collection("products")

	var newProduct models.Product
	if err := json.NewDecoder(r.Body).Decode(&newProduct); err != nil {
		http.Error(w, "Erro ao decodificar o JSON", http.StatusBadRequest)
		return
	}

	result, err := collection.InsertOne(context.Background(), newProduct)
	if err != nil {
		http.Error(w, "Erro ao inserir o produto no banco de dados", http.StatusInternalServerError)
		return
	}

	insertedID := result.InsertedID

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Produto criado com sucesso. ID: %v", insertedID)
}

// Lê todos os produtos armazenados no banco.
func ListProductsFromMongoDB(w http.ResponseWriter, r *http.Request) {
	client, err := ConnectToMongoDB()
	if err != nil {
		http.Error(w, "Erro ao conectar ao MongoDB", http.StatusInternalServerError)
		return
	}
	defer client.Disconnect(context.Background())

	database := client.Database("crud_GO")
	collection := database.Collection("products")

	cursor, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		http.Error(w, "Erro ao consultar o banco de dados", http.StatusInternalServerError)
		return
	}
	defer cursor.Close(context.Background())

	var products []models.Product
	for cursor.Next(context.Background()) {
		var product models.Product
		if err := cursor.Decode(&product); err != nil {
			log.Printf("Erro ao decodificar produto: %v", err)
			continue
		}
		products = append(products, product)
	}

	jsonResponse, err := json.Marshal(products)
	if err != nil {
		http.Error(w, "Erro ao serializar resposta JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

// Atualiza as informações de um produto no banco de dados
func UpdateProductInMongoDB(w http.ResponseWriter, r *http.Request) {
	// Extrair o ID do produto a ser atualizado da query string
	productID := r.URL.Query().Get("id")
	if productID == "" {
		http.Error(w, "ID do produto não fornecido", http.StatusBadRequest)
		return
	}

	// Converter o ID para ObjectID
	objID, err := primitive.ObjectIDFromHex(productID)
	if err != nil {
		http.Error(w, "ID do produto inválido", http.StatusBadRequest)
		return
	}

	// Conectar ao MongoDB
	client, err := ConnectToMongoDB()
	if err != nil {
		http.Error(w, "Erro ao conectar ao MongoDB", http.StatusInternalServerError)
		return
	}
	defer client.Disconnect(context.Background())

	// Acessar a coleção de produtos
	database := client.Database("crud_GO") // Substitua 'seunovobanco' pelo nome do seu banco de dados
	collection := database.Collection("products")

	// Decodificar o JSON da solicitação para obter os dados atualizados do produto
	var updatedProduct models.Product
	if err := json.NewDecoder(r.Body).Decode(&updatedProduct); err != nil {
		http.Error(w, "Erro ao decodificar o JSON", http.StatusBadRequest)
		return
	}

	// Atualizar o produto com o ID fornecido
	filter := bson.M{"_id": objID}
	update := bson.M{"$set": updatedProduct}
	_, err = collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		http.Error(w, "Erro ao atualizar o produto", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Produto atualizado com sucesso. ID: %s", productID)
}

// Exclui um produto do banco de dados
func DeleteProductFromMongoDB(w http.ResponseWriter, r *http.Request) {
	// Extrair o ID do produto a ser deletado da query string
	productID := r.URL.Query().Get("id")
	if productID == "" {
		http.Error(w, "ID do produto não fornecido", http.StatusBadRequest)
		return
	}

	// Converter o ID para ObjectID
	objID, err := primitive.ObjectIDFromHex(productID)
	if err != nil {
		http.Error(w, "ID do produto inválido", http.StatusBadRequest)
		return
	}

	// Conectar ao MongoDB
	client, err := ConnectToMongoDB()
	if err != nil {
		http.Error(w, "Erro ao conectar ao MongoDB", http.StatusInternalServerError)
		return
	}
	defer client.Disconnect(context.Background())

	// Acessar a coleção de produtos
	database := client.Database("crud_GO") // Substitua 'seunovobanco' pelo nome do seu banco de dados
	collection := database.Collection("products")

	// Deletar o produto com o ID fornecido
	filter := bson.M{"_id": objID}
	_, err = collection.DeleteOne(context.Background(), filter)
	if err != nil {
		http.Error(w, "Erro ao deletar o produto", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Produto deletado com sucesso. ID: %s", productID)
}
